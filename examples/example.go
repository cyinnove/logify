package main

import (
	log "github.com/cyinnove/logify"
)

func main() {
	//host := "https://example.com"
	path := "docs/example.go"

	log.Msg().Test(path)

	CustomLogger(log.Red, "Custom", "This is a custom log message with color %s", "Red")

}

func CustomLogger(color log.Color, holder, message string, args ...interface{}) {
	formatter := log.Formatter{}

	formatter.SetHolder(holder)
	formatter.SetMessage(message, args...)	
	formatter.SetColor(color)
	formatter.Log()

}
