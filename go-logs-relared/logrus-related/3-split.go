package main

import (
	"compress/gzip"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SplitHook struct {
	logPath       string // 存放日志的目录
	logDateFormat string // log按照日期分隔 日期格式化
	fileMaxSize   int64  // 日志文件最大默认，超过会压缩为gzip
}

var (
	debugCurDate string // 当前时间(精确到天)
	errCurDate   string // 当前时间(精确到天)
	debugFile    *os.File
	errFile      *os.File
	debugFileIdx int // debug 日志文件压缩文件索引
	errorFileIdx int // error 日志文件压缩文件索引
	rwLock       sync.RWMutex
)

func (splitHook *SplitHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.ErrorLevel,
	}
}

func (splitHook *SplitHook) Fire(entry *logrus.Entry) error {
	err := splitHook.checkMakeLogDir(entry) // 检测是否需要创建日志等级目录
	if err != nil {
		return err
	}
	err = splitHook.checkMakeLogFile(entry) // 检测是否需要创建新文件
	if err != nil {
		return err
	}
	err = splitHook.checkCreateNewLog(entry) // 检测文件大小是否超限 超限就压缩
	if err != nil {
		return err
	}
	str, _ := entry.String()
	switch entry.Level {
	case logrus.DebugLevel:
		if _, debugWriteErr := debugFile.Write([]byte(str)); debugWriteErr != nil {
			logrus.Error(debugWriteErr, "debugWriteErr", debugFile)
			return debugWriteErr
		}
	case logrus.ErrorLevel:
		if _, errWriteErr := errFile.Write([]byte(str)); errWriteErr != nil {
			return errWriteErr
		}
	}
	return nil
}

// 生成当前日期
func (splitHook *SplitHook) generateCurDate(entry *logrus.Entry) string {
	dateFormat := splitHook.logDateFormat
	if len(dateFormat) == 0 {
		dateFormat = "20060102"
	}
	return entry.Time.Format(dateFormat)
}

// 生成当前日志目录
func (splitHook *SplitHook) generateCurLogDir(entry *logrus.Entry, date string) string {
	// log文件路径
	curLogDir := path.Join(splitHook.logPath, entry.Level.String(), date)
	return curLogDir
}

// 检测是否需要创建新的日志
func (splitHook *SplitHook) checkMakeLogFile(entry *logrus.Entry) error {
	rwLock.Lock()
	defer rwLock.Unlock()
	curLogDate := splitHook.generateCurDate(entry)
	var compareDate string
	switch entry.Level {
	case logrus.DebugLevel:
		compareDate = debugCurDate
	case logrus.ErrorLevel:
		compareDate = errCurDate
	}
	if curLogDate <= compareDate {
		return nil
	}
	switch entry.Level {
	case logrus.DebugLevel:
		debugCurDate = curLogDate
	case logrus.ErrorLevel:
		errCurDate = curLogDate
	}
	curLogDir := splitHook.generateCurLogDir(entry, curLogDate)
	// log文件路径
	curLogPath := path.Join(curLogDir, fmt.Sprintf("%s-%s.log", curLogDate, entry.Level.String()))
	logFile, createFileErr := os.OpenFile(curLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if createFileErr != nil {
		return createFileErr
	}
	switch entry.Level {
	case logrus.DebugLevel:
		debugFile.Close()
		debugFile = logFile
		debugCurDate = curLogDate
	case logrus.ErrorLevel:
		errFile.Close()
		errFile = logFile
		errCurDate = curLogDate
	}
	return nil
}

// 检测是否需要压缩文件并创建新的文件
func (splitHook *SplitHook) checkCreateNewLog(entry *logrus.Entry) error {
	rwLock.Lock()
	defer rwLock.Unlock()
	// log文件路径
	var compareDate string
	switch entry.Level {
	case logrus.DebugLevel:
		compareDate = debugCurDate
	case logrus.ErrorLevel:
		compareDate = errCurDate
	}
	curLogDir := splitHook.generateCurLogDir(entry, compareDate)
	curLogPath := path.Join(curLogDir, fmt.Sprintf("%s-%s.log", compareDate, entry.Level.String()))
	fileInfo, err := os.Stat(curLogPath) // 文件不存在就忽略
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	maxSize := splitHook.fileMaxSize
	if maxSize == 0 {
		maxSize = 50 * 1024 * 1024 // 最大50M
	}
	fileSize := fileInfo.Size()
	if fileSize >= maxSize {
		idx := 0
		// 压缩当前日志文件
		switch entry.Level {
		case logrus.DebugLevel:
			idx = debugFileIdx
			debugFileIdx++
			fmt.Println("DEBUG超过最大限制", idx)
		case logrus.ErrorLevel:
			idx = errorFileIdx
			errorFileIdx++
			fmt.Println("ERROR超过最大限制", idx)
		}
		baseName := path.Base(curLogPath)
		dirName := path.Dir(curLogPath)
		nameAndExt := strings.Split(baseName, ".")
		gzippedLogFilePath := fmt.Sprintf("%s.gz", path.Join(dirName, fmt.Sprintf("%s-%d.%s", nameAndExt[0], idx, nameAndExt[1])))
		err = compressFile(curLogPath, gzippedLogFilePath)
		if err != nil {
			return err
		}
		// 清空原始的日志文件内容
		truncFile, truncErr := os.OpenFile(curLogPath, os.O_TRUNC, 0666)
		defer truncFile.Close()
		if truncErr != nil {
			return truncErr
		}
	}
	return nil
}

// 压缩文件
func compressFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	gzWriter := gzip.NewWriter(destinationFile)
	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, sourceFile)
	return err
}

// 检测是否需要创建新的日志目录
func (splitHook *SplitHook) checkMakeLogDir(entry *logrus.Entry) error {
	rwLock.Lock()
	defer rwLock.Unlock()
	curLogDate := splitHook.generateCurDate(entry)
	var compareDate string
	switch entry.Level {
	case logrus.DebugLevel:
		compareDate = debugCurDate
	case logrus.ErrorLevel:
		compareDate = errCurDate
	}
	if curLogDate <= compareDate {
		return nil
	}
	// 需要更新日期
	logDirPath := splitHook.generateCurLogDir(entry, curLogDate)
	debugFileIdx = 0
	errorFileIdx = 0
	if mkErr := os.MkdirAll(logDirPath, os.ModePerm); mkErr != nil { // 创建目录
		return mkErr
	}
	return nil
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// 只取后两级目录
			filePath := path.Join(path.Base(path.Dir(frame.File)), path.Base(frame.File))
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%s", filePath, strconv.Itoa(frame.Line))
		},
	})
	logrus.SetReportCaller(true)
	logrus.AddHook(&SplitHook{
		logPath:       "logs",
		logDateFormat: "200601021504",
		fileMaxSize:   1 * 1024,
	})
	logrus.SetLevel(logrus.DebugLevel)
	ticker := time.NewTicker(time.Second * 1)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				logrus.Debugln("打印一个Debug")
				logrus.Errorln("打印一个Error")
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
	time.Sleep(time.Second * 60 * 2)
	done <- true
}

func init() {
	rwLock = sync.RWMutex{}
}
