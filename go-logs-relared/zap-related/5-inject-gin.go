package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"strings"
	"syscall"
	"time"
	"zap-related/global"
)

// 监听中断信号，中断后优雅退出gin
func gracefulShutdown(serv *http.Server) {
	// 监听程序退出信号(意外退出/ctrl + c)
	signalCtx, signalCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer signalCancel()
	// 阻塞 监听中断信号
	<-signalCtx.Done()
	// 给服务器5秒钟时间来中断服务，5秒钟后强制中断服务
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()
	if err := serv.Shutdown(timeoutCtx); err != nil {
		global.SugarLogger.Errorf("Shutdown Server Error：%s\n", err.Error())
	}
	global.SugarLogger.Debugln("Server is Graceful Shutdown")
}

func RunServer() {
	engine := gin.New()
	engine.Use(GinZapLogger(global.SugarLogger), GinZapRecovery(global.SugarLogger, true))
	engine.GET("/user", func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ip := ctx.ClientIP()
		global.SugarLogger.Debugf("sugarLogger startServer log，path is %s, query is %s, ip is %s\n", path, query, ip)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{"name": "张三"},
		})
	})
	engine.GET("/p", func(ctx *gin.Context) {
		panic("This is a deliberate panic for testing or demonstration purposes")
	})
	//serv := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", 6688),
	//	Handler: engine,
	//}
	//// 开启一个协程来启动服务，并检测服务其他中断错误，这样不会阻塞下面的优雅退出检测
	//go func() {
	//	if err := serv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		global.SugarLogger.Errorf("server error，the cause is %s\n", err.Error())
	//	}
	//}()
	//gracefulShutdown(serv)
	err := engine.Run(":6688")
	if err != nil {
		global.SugarLogger.Errorf("start server error，%s", err)
		return
	}
}

func initLogger5() {
	encoder := getEncoder5()
	writer := getWriter5()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	global.SugarLogger = logger.Sugar()
}

func getEncoder5() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time" // 设置日志时间的key为time
	// 格式化日志显示的时间格式
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.DateTime))
	}
	// 设置日志等级修改为大写 info => INFO debug => DEBUG
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriter5() zapcore.WriteSyncer {
	rootDir, _ := os.Getwd()
	separator := string(filepath.Separator)
	logDirPath := fmt.Sprintf("%s%s%s%s%s", rootDir, separator, "logs", separator, "gin-logs.log")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logDirPath, // 日志文件名称(路径)
		MaxSize:    20,         // 日志文件最大的尺寸(MB), 超限后开始自动分割
		MaxBackups: 30,         // 保留旧文件的最大个数
		MaxAge:     30,         // 保留旧文件的最大天数
		Compress:   false,      // 是否压缩/归档旧文件
		LocalTime:  false,      // 确定备份文件中时间戳格式化所使用的时间是否为计算机的本地时间。默认情况下，使用`UTC`时间。
	}
	// 在控制台输出日志，同时将日志记录到log文件中
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLogger), zapcore.AddSync(os.Stdout))
}

// GinZapLogger 实现gin logger 中间件
func GinZapLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		// Process request
		ctx.Next()
		cost := time.Since(start)
		ClientIp := ctx.ClientIP()
		Method := ctx.Request.Method
		StatusCode := ctx.Writer.Status()
		ErrMsg := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		if raw != "" {
			path = path + "?" + raw
		}
		logger.Infof(
			"Path:%s，Method:%s，StatusCode:%d，IP:%s，ErrMsg:%s，Cost:%d",
			path,
			Method,
			StatusCode,
			ClientIp,
			ErrMsg,
			cost,
		)
	}
}

// GinZapRecovery recover掉项目可能出现的panic
func GinZapRecovery(logger *zap.SugaredLogger, stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				headersToStr := strings.Join(headers, "\r\n")
				if brokenPipe {
					logger.Errorf("%s\n%s", err, headersToStr)
					// If the connection is dead, we can't write a status to it.
					ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}
				// 在日志中展示调用栈信息
				if stack {
					logger.Errorf(
						"[Recovery] panic recovered:\n%s\n%s%s",
						err,
						headersToStr,
						debug.Stack(),
					)
				} else {
					logger.Errorf(
						"[Recovery] panic recovered:\n%s\n%s",
						err,
						headersToStr,
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}

func main() {
	initLogger5()
	RunServer()
}
