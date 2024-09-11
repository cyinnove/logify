package logify

import (
	"strings"
	"sync"
)

var stringBuilderPool = &sync.Pool{New: func() interface{} {
	return new(strings.Builder)
}}

// getLabel generates a label based on the level and wraps it in color if needed
func getLabel(level Level, label string, sb *strings.Builder) {
	if level == Silent || level == Misc {
		return
	}
	sb.WriteString("[")
	sb.WriteString(Colorize(labels[level], level))
	sb.WriteString("] ")
}
