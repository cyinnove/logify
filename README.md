## Overview

This tool provides a custom logging solution for Go applications. It utilizes the `logify` package to log messages with various levels of severity and custom formatting. The logger supports different log levels such as INFO, TEST, DEBUG, FATAL, ERROR, and WARN. Additionally, it allows the creation of custom log messages with specified colors and formats. The logger follows the Singleton pattern to ensure that the same logger instance is used throughout the entire application.

This logging package can be integrated into various types of applications, including CLI tools, security tools, web servers, and more. Its flexibility and customizability make it suitable for any application that requires robust logging capabilities.

## Features

- **Log Levels**: Supports multiple log levels including INFO, TEST, DEBUG, FATAL, ERROR, and WARN.
- **Custom Logger**: Allows for custom log messages with specified colors and formats.
- **Singleton Pattern**: Ensures a single instance of the logger is used throughout the application.

## Installation

To use this tool, you need to install the `logify` package. You can do this using `go get`:

```sh
go get github.com/cyinnove/logify@latest
```

## Usage

### Importing the Package

First, import the necessary package in your Go application:

```go
import (
    log "github.com/cyinnove/logify"
)
```

### Basic Logging

Here's an example of how to use the basic logging functionality:

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

### Custom Logger

You can create custom log messages with specified colors and formats:

```go
package main

import (
    log "github.com/cyinnove/logify"
)

func main() {
    path := "examples/example.go"
    // Default logger with warning level
    log.Msg().Warn(path)

    // Custom logger example
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
## Logging Example

![Logging Example](/static/logs.png)


## Log Levels

The logger supports the following log levels:
- **INFO**
- **TEST**
- **DEBUG**
- **FATAL**
- **ERROR**
- **WARN**

These levels can be used to categorize and filter log messages based on their severity.

## Singleton Pattern

The logger uses the Singleton pattern to ensure that a single instance of the logger is used across the entire application. This helps maintain consistency and avoids potential issues with multiple logger instances.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

