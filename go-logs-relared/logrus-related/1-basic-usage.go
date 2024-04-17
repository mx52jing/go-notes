package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func showLogs() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Errorln("logrus => Error")
	logrus.Warningln("logrus => Warn")
	logrus.Infoln("logrus => Info")
	logrus.Debugln("logrus => Debug")
	fmt.Println(logrus.GetLevel())
}

func withFields() {
	logrus.SetLevel(logrus.DebugLevel)
	//logInstance := logrus.WithField("name", "如燕")
	logInstance := logrus.WithFields(logrus.Fields{
		"name": "如燕",
		"age":  18,
	})
	logInstance.Errorln("logInstance => Error")
	logInstance.Warningln("logInstance => Warn")
	logInstance.Infoln("logInstance => Info")
	logInstance.Debugln("logInstance => Debug")
}

func setFormatter() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	})
	logInstance := logrus.WithFields(logrus.Fields{
		"name": "如燕",
		"age":  18,
	})
	logInstance.Infoln("info")
}

func outputFile() {
	file, _ := os.OpenFile("logs/a.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	writes := []io.Writer{
		file,
		os.Stdout,
	}
	outputWrite := io.MultiWriter(writes...)
	logrus.SetOutput(outputWrite)
	//logrus.SetFormatter(&logrus.TextFormatter{
	//	DisableColors: false,
	//	//DisableTimestamp: false,
	//	FullTimestamp: true,
	//	//TimestampFormat: "2006-01-02 15:04:05",
	//})
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// 只取后两级目录
			filePath := path.Join(path.Base(path.Dir(frame.File)), path.Base(frame.File))
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%s", filePath, strconv.Itoa(frame.Line))
		},
	})
	logInstance := logrus.WithFields(logrus.Fields{
		"name": "狄仁杰",
		"age":  70,
	})
	logInstance.Infoln("info")
}

func showCallerLine() {
	logrus.SetReportCaller(true)
	logrus.Errorln("logrus => Error")
	logrus.Warningln("logrus => Warn")
	logrus.Infoln("logrus => Info")
	logrus.Debugln("logrus => Debug")
}

func printColor() {
	fmt.Println("\033[30m 黑色文本 \033[0m")
	fmt.Println("\033[31m 红色文本 \033[0m")
	fmt.Println("\x1b[32m 绿色文本 \x1b[0m")
	fmt.Println("\033[33m 黄色文本 \033[0m")
	fmt.Println("\033[34m 蓝色文本 \033[0m")
	fmt.Println("\x1b[35m 紫色文本 \x1b[0m")
	fmt.Println("\033[36m 青色文本 \033[0m")
	fmt.Println("\033[37m 灰色文本 \033[0m")
	// 背景色
	fmt.Println("\033[40m 黑色背景 \033[0m")
	fmt.Println("\033[41m 红色背景 \033[0m")
	fmt.Println("\033[42m 绿色背景 \033[0m")
	fmt.Println("\033[43m 黄色背景 \033[0m")
	fmt.Println("\033[44m 蓝色背景 \033[0m")
	fmt.Println("\033[45m 紫色背景 \033[0m")
	fmt.Println("\x1b[46m 青色背景 \x1b[0m")
	fmt.Println("\x1b[47m 灰色背景 \x1b[0m")
}

// 设置一套颜色
const (
	CRed    = 31
	CYellow = 33
	CCyan   = 36
	CPurple = 35
	CGray   = 37
)

type MFormat struct{}

func (m MFormat) Format(entry *logrus.Entry) ([]byte, error) {
	// 设置日志颜色
	var levelColor int
	switch entry.Level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		levelColor = CRed
	case logrus.WarnLevel:
		levelColor = CYellow
	case logrus.InfoLevel:
		levelColor = CCyan
	case logrus.DebugLevel:
		levelColor = CPurple
	default:
		levelColor = CGray
	}
	var buffer *bytes.Buffer
	if buffer = entry.Buffer; buffer == nil {
		buffer = &bytes.Buffer{}
	}
	fmt.Fprintf(buffer, "\x1b[%dm %s [%s]\x1b[0m\n", levelColor, entry.Level, entry.Message)
	return buffer.Bytes(), nil
}

func customLogFormat() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&MFormat{})
	log.Errorln("自定义颜色 => Error输出")
	log.Warningln("自定义颜色 => Warn输出")
	log.Infoln("自定义颜色 => Info输出")
	log.Debugln("自定义颜色 => Debug输出")
}

func main() {
	//showLogs()
	//withFields()
	//setFormatter()
	//outputFile()
	//showCallerLine()
	//printColor()
	customLogFormat()
}
