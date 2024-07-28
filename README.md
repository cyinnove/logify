## Overview

This tool provides a custom logging solution for Go applications using the `logify` package. It supports various log levels such as INFO, TEST, DEBUG, FATAL, ERROR, and WARN, and allows for custom log messages with specified colors and formats. The logger follows the Singleton pattern to ensure a consistent logging instance throughout the application.

This logging package is versatile, suitable for CLI tools, security tools, web servers, and more, offering robust and customizable logging capabilities.

## Features

- **Log Levels**: INFO, TEST, DEBUG, FATAL, ERROR, WARN.
- **Custom Logger**: Custom log messages with specified colors and formats.
- **Singleton Pattern**: Single logger instance throughout the application.

## Installation

Install the `logify` package using `go get`:

```sh
go get github.com/cyinnove/logify@latest
```

## How to Use

### Importing the Package

Import the package in your Go application:

```go
import (
    log "github.com/cyinnove/logify"
)
```

### Basic Logging

Example of basic logging functionality:

```go
package main

import (
    log "github.com/cyinnove/logify"
)

func main() {
    path := "docs/example.go"
    log.Msg().Warn(path)
}
```

### Custom Logging

Creating custom log messages with specified colors and formats:

```go
package main

import (
    log "github.com/cyinnove/logify"
)

func main() {
    path := "examples/example.go"
    log.Msg().Warn(path)

    CustomLogger(log.Red, "Custom", "This is a custom log message with color %s", "Red")
}

func CustomLogger(color log.Color, holder, message string, args ...interface{}) {
    formatter := log.Formatter{}
    formatter.SetHolder(holder)
    formatter.SetMessage(message, args...)
    formatter.SetColor(color)
    formatter.Log()
}
```

## Log Levels

The logger supports the following log levels:

- **INFO**
- **TEST**
- **DEBUG**
- **FATAL**
- **ERROR**
- **WARN**

These levels help categorize and filter log messages based on severity.

## Logging Example

![Logging Example](/static/logs.png)

## Singleton Pattern

The logger ensures a single instance across the entire application, maintaining consistency and avoiding issues with multiple instances.

## Contributing

Contributions are welcome! Submit a pull request or open an issue for suggestions or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
