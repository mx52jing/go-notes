package shared

import (
	"os"
	"strings"
)

// BodyFrom 从命令行输入中获取文案
func BodyFrom(args []string) string {
	if (len(args) < 3) || os.Args[2] == "" {
		return "Hello MQ"
	}
	return strings.Join(args[2:], ",")
}

// GetLogSeverity 从命令行参数中获取日志等级
func GetLogSeverity(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "info"
	}
	return os.Args[1]
}
