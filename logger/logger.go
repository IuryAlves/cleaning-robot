package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type BasicLogger struct{}

func (b *BasicLogger) log(msg string, args ...any) {
	_, filepath, line, ok := runtime.Caller(2)
	if ok {
		// Add file and line that called the logger
		_, filename := path.Split(filepath)
		msg = fmt.Sprintf("%s:%v %s", filename, line, msg)
	}
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	fmt.Printf(msg, args...)
}

func (b *BasicLogger) Error(msg string, args ...any) {
	msg = fmt.Sprintf("ERROR: %s", msg)
	b.log(msg, args)
}

func (b *BasicLogger) Info(msg string, args ...any) {
	msg = fmt.Sprintf("INFO: %s", msg)
	b.log(msg, args)
}
