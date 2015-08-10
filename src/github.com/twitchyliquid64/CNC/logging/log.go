package logging

import (
	"log"
	"fmt"
	"strings"
	"os"
)


func Info(module string, content ...interface{}) {
	writeLogLine(formatLogPrefix(module, "I", content...))
}

func Warning(module string, content ...interface{}) {
	writeLogLine(formatLogPrefix(module, "W", content...))
}

func Error(module string, content ...interface{}) {
	writeLogLine(formatLogPrefix(module, "E", content...))
}

func Fatal(module string, content ...interface{}) {
	writeLogLine(formatLogPrefix(module, "F", content...))
	os.Exit(1)
}


func formatLogPrefix(module, messagePrefix string, content ...interface{})string {
	c := fmt.Sprint(content...)
	module = strings.ToUpper(module)
	if module != "" {
		return "[" + messagePrefix + "] [" + module + "] " + c
	} else {
		return "[" + messagePrefix + "] " + c
	}
}

func writeLogLine(inp string) {
	log.Println(inp)
}
