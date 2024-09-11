package logify

import "os"

// Infof writes an info message
func Infof(format string, args ...interface{}) {
	log(Info, "", format, args...)
}

// Warningf writes a warning message
func Warningf(format string, args ...interface{}) {
	log(Warning, "", format, args...)
}

// Errorf writes an error message
func Errorf(format string, args ...interface{}) {
	log(Error, "", format, args...)
}

// Debugf writes a debug message
func Debugf(format string, args ...interface{}) {
	log(Debug, "", format, args...)
}

// Verbosef writes a verbose message
func Verbosef(format string, label string, args ...interface{}) {
	log(Verbose, label, format, args...)
}

// Silentf writes a message with no label
func Silentf(format string, args ...interface{}) {
	log(Silent, "", format, args...)
}

// Fatalf exits the program after logging a fatal message
func Fatalf(format string, args ...interface{}) {
	log(Fatal, "", format, args...)
	os.Exit(1)
}
