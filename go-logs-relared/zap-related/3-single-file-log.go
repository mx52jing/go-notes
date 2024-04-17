package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"zap-related/global"
)

func initSugarLogger3() {
	encoder := initLogEncoder3()
	debugWriter := initDebugWriter()
	errorWriter := initErrorWriter()
	// 记录debug和更高级别的日志
	c1 := zapcore.NewCore(encoder, debugWriter, zapcore.DebugLevel)
	// 记录error日志
	c2 := zapcore.NewCore(encoder, errorWriter, zapcore.ErrorLevel)
	// 使用zapcore.NewTee方法将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	global.SugarLogger = logger.Sugar()
}

// 初始化写入日志的格式
//
//	func initLogEncoder() zapcore.Encoder {
//		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//	}
func initLogEncoder3() zapcore.Encoder {
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
func initDebugWriter() zapcore.WriteSyncer {
	rootDir, _ := os.Getwd()
	separator := string(filepath.Separator)
	logDirPath := fmt.Sprintf("%s%s%s%s%s", rootDir, separator, "logs", separator, "logs.log")
	file, err := os.OpenFile(logDirPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		fmt.Printf("创建日志文件失败: %s\n", err)
	}
	// 在控制台输出日志，同时将日志记录到log文件中
	return zapcore.NewMultiWriteSyncer(file, zapcore.AddSync(os.Stdout))
}

func initErrorWriter() zapcore.WriteSyncer {
	rootDir, _ := os.Getwd()
	separator := string(filepath.Separator)
	logDirPath := fmt.Sprintf("%s%s%s%s%s", rootDir, separator, "logs", separator, "logs.error.log")
	file, err := os.OpenFile(logDirPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		fmt.Printf("创建日志文件失败: %s\n", err)
	}
	// 在控制台输出日志，同时将日志记录到log文件中
	return zapcore.NewMultiWriteSyncer(file, zapcore.AddSync(os.Stdout))
}

func multiReqLog3(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		global.SugarLogger.Errorf("sugarLogger error, url is %s\n", url)
		return
	}
	global.SugarLogger.Debugf("sugarLogger success, url is %s, statusCode is %d", url, resp.StatusCode)
}

func main() {
	initSugarLogger3()
	defer global.SugarLogger.Sync()
	multiReqLog3("https://www.baidu.com")
	multiReqLog3("http://192.168.1.9")
	global.SugarLogger.Errorln("测试单独创建error.log")
	global.SugarLogger.Infoln("end")
}
