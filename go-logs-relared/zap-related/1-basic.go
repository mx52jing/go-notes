package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"zap-related/global"
)

func intLoggerConfig() {
	logger, _ := zap.NewProduction()
	global.Logger = logger
}

func initSugarLogger() {
	logger, _ := zap.NewProduction()
	global.SugarLogger = logger.Sugar()
}

func reqGetLog(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		global.Logger.Error("logger error", zap.String("url", url))
		global.SugarLogger.Errorf("sugarLogger error, url is %s\n", url)
		return
	}
	global.Logger.Info(
		"logger success",
		zap.String("url", url),
		zap.Int("statusCode", resp.StatusCode),
	)
	global.SugarLogger.Infof("sugarLogger success, url is %s, statusCode is %d", url, resp.StatusCode)
}

func main() {
	intLoggerConfig()
	initSugarLogger()
	defer global.Logger.Sync()
	defer global.SugarLogger.Sync()
	reqGetLog("https://www.baidu.com")
	reqGetLog("http://192.198.1.1")
	fmt.Printf("end")
}
