package logify

import "fmt"

const (
	Red     = "\033[31m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Reset   = "\033[0m"
	Bold    = "\033[1m"
)

// Colorize adds color based on the log level
func Colorize(text string, level Level) string {
	if !UseColors {
		return text
	}

	switch level {
	case Fatal:
		return fmt.Sprintf("%s%s%s%s", Bold, Red, text, Reset)
	case Error:
		return fmt.Sprintf("%s%s%s", Red, text, Reset)
	case Warning, Label:
		return fmt.Sprintf("%s%s%s", Yellow, text, Reset)
	case Debug:
		return fmt.Sprintf("%s%s%s", Magenta, text, Reset)
	case Info, Verbose:
		return fmt.Sprintf("%s%s%s", Blue, text, Reset)
	default:
		return text
	}
}
