package main

import (
	"github.com/cyinnove/logify"
)

func main() {
	// Set logging configuration
	logify.UseColors = true
	

	// Example logging
	logify.Infof("This is an %s message", "info")
	logify.Infof("This is new info msg")
	logify.Warningf("This is a %s message", "warning")
	logify.Errorf("This is an %s message", "error")
	logify.Debugf("This is a %s message", "debug")
	logify.Verbosef("This is a verbose message with a label", "LABEL")
	logify.Silentf("This is a silent message")
	
	// Fatal error example (uncomment to test)
	// logify.Fatalf("This is a fatal message, the program will exit")
}
