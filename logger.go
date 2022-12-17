package cleaning_robot

import (
	"fmt"
	"strings"
)

type BasicLogger struct{}

func (b *BasicLogger) Log(msg string, args ...any) {
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	fmt.Printf(msg, args...)
}
