# logify

`logify` is a customizable logging solution for Go applications that supports various log levels and custom formatting. This package offers flexibility in logging messages with different levels of severity, colors, and formats. It uses a simple and efficient approach to logging by providing a single instance of the logger across the application.

## Features

- **Log Levels**: Supports multiple log levels including `INFO`, `DEBUG`, `ERROR`, `FATAL`, `WARNING`, and `SILENT`.
- **Custom Colors**: Allows custom log messages with specified colors.
- **Singleton Pattern**: Ensures a single logger instance is used throughout the application for consistency.

## Installation

To use `logify`, install the package via `go get`:

```sh
go get github.com/cyinnove/logify
```

## Usage

### Importing the Package

Import the `logify` package in your Go application:

```go
import (
    "github.com/cyinnove/logify"
)
```

### Basic Logging

Here's an example of basic logging:

```go
package main

import (
    "github.com/cyinnove/logify"
)

func main() {
    logify.UseColors = true
    logify.MaxLevel = logify.Debug

    logify.Infof("This is an %s message", "info")
    logify.Warningf("This is a %s message", "warning")
    logify.Errorf("This is an %s message", "error")
    logify.Debugf("This is a %s message", "debug")
    logify.Verbosef("This is a verbose message with a label", "LABEL")
    logify.Silentf("This is a silent message")
    
    // Uncomment to test Fatalf
    // logify.Fatalf("This is a fatal message, the program will exit")
}
```

### Custom Logger

Create custom log messages with specified colors:

```go
package main

import (
    "github.com/cyinnove/logify"
)

func main() {
    logify.UseColors = true
    logify.MaxLevel = logify.Debug

    // Default logging
    logify.Warningf("Default warning message")

    // Custom logging
    CustomLogger(logify.Red, "CustomLabel", "This is a custom log message with color %s", "Red")
}

func CustomLogger(color logify.Color, label, format string, args ...interface{}) {
    logify.UseColors = true
    logify.MaxLevel = logify.Debug
    logify.Printf(format, args...)
}
```

## Log Levels

The logger supports the following levels:

- **INFO**: Informational messages.
- **DEBUG**: Debugging messages.
- **ERROR**: Error messages.
- **FATAL**: Fatal errors that cause application exit.
- **WARNING**: Warning messages.
- **SILENT**: Messages with no label.

## Singleton Pattern

The logger follows the Singleton pattern to maintain a single instance throughout the application, ensuring consistent logging behavior and avoiding multiple instances.

## Contributing

Contributions are welcome! To contribute, please submit a pull request or open an issue with your suggestions or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

