package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"zap-related/global"
)

func initSugarLogger4() {
	encoder := initLogEncoder4()
	writer := initLogWriter4()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	global.SugarLogger = logger.Sugar()
}

// 初始化写入日志的格式
//
//	func initLogEncoder() zapcore.Encoder {
//		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//	}
func initLogEncoder4() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "time" // 设置日志时间的key为time
	// 格式化日志显示的时间格式
	config.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.DateTime))
	}
	// 设置日志等级修改为大写 info => INFO debug => DEBUG
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(config)
}

// 初始化写入日志的位置/路径
func initLogWriter4() zapcore.WriteSyncer {
	rootDir, _ := os.Getwd()
	separator := string(filepath.Separator)
	logDirPath := fmt.Sprintf("%s%s%s%s%s", rootDir, separator, "logs", separator, "logs.log")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logDirPath, // 日志文件名称(路径)
		MaxSize:    20,         // 日志文件最大的尺寸(MB), 超限后开始自动分割
		MaxBackups: 30,         // 保留旧文件的最大个数
		MaxAge:     30,         // 保留旧文件的最大天数
		Compress:   false,      // 是否压缩/归档旧文件
	}
	// 在控制台输出日志，同时将日志记录到log文件中
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLogger), zapcore.AddSync(os.Stdout))
}

func multiReqLog4(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		global.SugarLogger.Errorf("sugarLogger error, url is %s\n", url)
		return
	}
	global.SugarLogger.Debugf("sugarLogger success, url is %s, statusCode is %d", url, resp.StatusCode)
}

func main() {
	initSugarLogger4()
	defer global.SugarLogger.Sync()
	multiReqLog4("https://www.baidu.com")
	multiReqLog4("http://192.168.1.9")
	global.SugarLogger.Errorln("test error")
	global.SugarLogger.Debugln("end")
}
