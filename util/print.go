package util

import (
	"log"
	"runtime"
	"strings"
)

func Debug(format string, v ...interface{}) {
	_, fileName, line, _ := runtime.Caller(1)

	parts := strings.Split(fileName, "/nbanitama/chat/")
	if len(parts) == 2 {
		fileName = parts[1]
	} else {
		fileName = ""
	}

	v = append([]interface{}{fileName, line}, v...)
	log.Printf("[%s:%d] "+format, v...)
}
