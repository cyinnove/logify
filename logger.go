package logify

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	UseColors = true
	MaxLevel  = Info
	mutex     = &sync.Mutex{}
)

// log logs the actual message to the screen
func log(level Level, label string, format string, args ...interface{}) {
	if level == Null || level > MaxLevel {
		return
	}

	sb := stringBuilderPool.Get().(*strings.Builder)
	defer stringBuilderPool.Put(sb)

	getLabel(level, label, sb)

	message := fmt.Sprintf(format, args...)
	sb.WriteString(message)
	if !strings.HasSuffix(message, "\n") {
		sb.WriteString("\n")
	}

	mutex.Lock()
	defer mutex.Unlock()
	if level == Silent {
		fmt.Fprint(os.Stdout, sb.String())
	} else {
		fmt.Fprint(os.Stderr, sb.String())
	}
}
