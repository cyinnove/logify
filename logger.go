package logger

import (
	"fmt"
	"sync"
)

// Logger interface defines the methods that any logger implementation should have.
type Logger interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

// LoggerOptions struct represents the default logger implementation.
type LoggerOptions struct {
	Formatter      Formatter
	ColorEnabled   bool // Controls color output
	DebugEnabled   bool // Controls debug log visibility
	InfoEnabled    bool // Controls info log visibility
	WarningEnabled bool // Controls warning log visibility
	ErrorEnabled   bool // Controls error log visibility
	FatalEnabled   bool // Controls fatal log visibility
}

var (
	once     sync.Once
	instance *LoggerOptions
)

// GetLogger returns the singleton instance of LoggerOptions.
// Msg returns an instance of the Logger.
// The Logger is a singleton object that provides logging functionality.
// It initializes the LoggerOptions with default values if it hasn't been initialized before.
// The LoggerOptions controls various logging options such as formatter, color enablement, and log level visibility.
func Msg() Logger {
	once.Do(func() {
		instance = &LoggerOptions{
			Formatter:      Formatter{},
			ColorEnabled:   true, // Colors are enabled by default
			DebugEnabled:   true, // All log levels are visible by default
			InfoEnabled:    true, // All log levels are visible by default
			WarningEnabled: true, // All log levels are visible by default
			ErrorEnabled:   true, // All log levels are visible by default
			FatalEnabled:   true, // All log levels are visible by default
		}
	})
	return instance
}

func (l *LoggerOptions) colorize(holder, color string) {
	if l.ColorEnabled {
		l.Formatter.SetHolder(fmt.Sprintf("%s%s%s", color, holder, Colors[Reset]))
	} else {
		l.Formatter.SetHolder(holder)
	}
}

// Info logs an info message with the specified message and arguments.
func (l *LoggerOptions) Info(msg string, args ...interface{}) {
	if !l.InfoEnabled {
		return
	}
	l.colorize(Info.String(), Colors[Blue])
	l.Formatter.SetMessage(msg, args...)
	l.Formatter.Log()
}

// Debug logs a debug message with the specified message and arguments.
func (l *LoggerOptions) Debug(msg string, args ...interface{}) {
	if !l.DebugEnabled {
		return
	}

	l.colorize(Debug.String(), Colors[Purple])
	l.Formatter.SetMessage(msg, args...)
	l.Formatter.Log()
}

// Warn logs a warning message with the specified message and arguments.
func (l *LoggerOptions) Warn(msg string, args ...interface{}) {
	if !l.WarningEnabled {
		return
	}
	l.colorize(Warn.String(), Colors[Yellow])
	l.Formatter.SetMessage(msg, args...)
	l.Formatter.Log()
}

// Error logs an error message with the specified message and arguments.
func (l *LoggerOptions) Error(msg string, args ...interface{}) {
	if !l.ErrorEnabled {
		return
	}
	l.colorize(Error.String(), Colors[Red])
	l.Formatter.SetMessage(msg, args...)
	l.Formatter.Log()
}

// Fatal logs a fatal message with the specified message and arguments.
func (l *LoggerOptions) Fatal(msg string, args ...interface{}) {
	if !l.FatalEnabled {
		return
	}
	l.colorize(Fatal.String(), Colors[Orange])
	l.Formatter.SetMessage(msg, args...)
	l.Formatter.Log()
}


