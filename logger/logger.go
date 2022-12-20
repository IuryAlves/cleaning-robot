package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type BasicLogger struct{}

func (b *BasicLogger) Log(msg string, args ...any) {
	_, filepath, line, ok := runtime.Caller(1)
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
