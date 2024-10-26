package logify

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	UseColors      = true
	MaxLevel       = Info
	mutex          = &sync.Mutex{}
	lastLogMessage string
	lastLogTime    time.Time
	duplicationGap = 100 * time.Millisecond // time gap to allow between duplicate logs
)

// log logs the actual message to the screen
func log(level Level, label string, format string, args ...interface{}) {
	if level == Null || level > MaxLevel {
		return
	}

	// Lock immediately to prevent concurrent calls
	mutex.Lock()
	defer mutex.Unlock()

	sb := stringBuilderPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		stringBuilderPool.Put(sb)
	}()

	// Build the message
	getLabel(level, label, sb)
	message := fmt.Sprintf(format, args...)
	sb.WriteString(message)
	if !strings.HasSuffix(message, "\n") {
		sb.WriteString("\n")
	}

	finalMessage := sb.String()

	// Deduplicate consecutive identical log entries within a short time window
	if finalMessage == lastLogMessage && time.Since(lastLogTime) < duplicationGap {
		return
	}
	lastLogMessage = finalMessage
	lastLogTime = time.Now()

	// Output to appropriate stream
	if level == Silent {
		fmt.Fprint(os.Stdout, finalMessage)
	} else {
		fmt.Fprint(os.Stderr, finalMessage)
	}
}
