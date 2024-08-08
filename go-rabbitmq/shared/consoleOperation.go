package shared

import (
	"os"
	"strings"
)

func BodyFrom(args []string) string {
	if (len(args) < 2) || os.Args[1] == "" {
		return "Hello"
	}
	return strings.Join(args[1:], " ")
}
