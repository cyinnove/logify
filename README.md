<<<<<<< HEAD
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
=======
# Logify

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Test Coverage](https://img.shields.io/badge/coverage-71.3%25-brightgreen.svg)](.)

A simple, lightweight, and feature-rich command-line flag parsing library for Go. Logify is inspired by [goflags](https://github.com/projectdiscovery/goflags) from ProjectDiscovery, focusing on the most important features with a cleaner, more maintainable codebase.

## Features

### Core Features
- ✨ **Simple API** - Clean and intuitive API design
- 🎯 **Rich Flag Types** - String, Bool, Int, Duration, Size, StringSlice, and Callback
- 📁 **File Input Support** - Load values from files (StringSlice)
- 🔤 **Short & Long Flags** - Support for both `-v` and `--verbose` style
- 📦 **Flag Groups** - Organize flags into logical groups
- 📝 **Auto Help** - Beautiful, automatically generated help messages
- 🎨 **Clean Output** - Well-formatted, grouped help display
- ✅ **Well Tested** - 71.3% test coverage with 65+ unit tests
- 🚀 **Zero External Dependencies** - Only uses Go standard library

### Advanced Flag Types
- **Duration**: Parse time durations (e.g., `30s`, `5m`, `1h`)
- **Size**: Parse file sizes with units (e.g., `10mb`, `5gb`, `1tb`)
- **StringSlice**: Support comma-separated values and file input
- **Callback**: Execute functions when flags are set

## Installation

```bash
go get -u github.com/cyinnove/logify
```

## Quick Start
>>>>>>> 4a5d0ef (first commit)

```go
package main

import (
<<<<<<< HEAD
=======
    "fmt"
    "log"
>>>>>>> 4a5d0ef (first commit)
    "github.com/cyinnove/logify"
)

func main() {
<<<<<<< HEAD
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
=======
    flags := logify.New("My CLI Tool")

    name := flags.String("name", "n", "", "Your name")
    verbose := flags.Bool("verbose", "v", false, "Verbose mode")
    
    if err := flags.Parse(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello, %s!\n", *name)
}
```

## Package Structure

The package is organized into modular files for maintainability:

```
logify/
├── logify.go              # Main package and core flag types
├── string_slice.go        # StringSlice implementation
├── size.go                # Size type with unit support
├── callback.go            # Callback flag implementation
├── utils.go               # Utility functions
├── logify_test.go         # Integration tests
├── string_slice_test.go   # StringSlice unit tests
├── size_test.go           # Size unit tests
├── callback_test.go       # Callback unit tests
└── utils_test.go          # Utils unit tests
```

## Usage Examples

### Basic Flags

```go
flags := logify.New("Basic example")

// String flag
output := flags.String("output", "o", "result.txt", "Output file")

// Boolean flag
debug := flags.Bool("debug", "d", false, "Debug mode")

// Integer flag
threads := flags.Int("threads", "t", 5, "Worker threads")

flags.Parse()
```

**Command line:**
```bash
./app --output=out.txt --debug --threads=10
# or with short flags
./app -o out.txt -d -t 10
```

### Duration Flags

```go
import "time"

flags := logify.New("Duration example")

timeout := flags.Duration("timeout", "t", 30*time.Second, "Request timeout")

flags.Parse()

fmt.Printf("Timeout: %v\n", *timeout)
```

**Command line:**
```bash
./app --timeout=5m     # 5 minutes
./app --timeout=30s    # 30 seconds
./app --timeout=1h30m  # 1 hour 30 minutes
```

### Size Flags

```go
flags := logify.New("Size example")

maxSize := flags.Size("max-size", "ms", "10mb", "Maximum file size")

flags.Parse()

fmt.Printf("Max size: %d bytes (%.2f MB)\n", 
    maxSize.Bytes(), maxSize.MB())
```

**Command line:**
```bash
./app --max-size=50mb   # 50 megabytes
./app --max-size=2gb    # 2 gigabytes
./app --max-size=1tb    # 1 terabyte
```

### String Slice Flags

String slices support multiple values through repeated flags or comma-separated values:

```go
flags := logify.New("Slice example")

// Basic string slice
tags := flags.StringSlice("tag", "t", "Tags")

// String slice with file support
targets := flags.FileStringSlice("target", "u", "Target URLs")

flags.Parse()
```

**Command line:**
```bash
# Multiple flags
./app --tag=go --tag=cli --tag=tool

# Comma-separated
./app --tag=go,cli,tool

# File input (with FileStringSlice)
./app --target=urls.txt
```

**File format (urls.txt):**
```
https://example.com
https://test.com
# Comments are ignored
https://demo.com
```

### Flag Groups

Organize flags into logical groups for better help output:

```go
flags := logify.New("Advanced tool")

var (
    input    *string
    output   *string
    verbose  *bool
    threads  *int
)

// Input group
flags.CreateGroup("input", "Input Options", func() {
    input = flags.String("input", "i", "", "Input file")
})

// Output group
flags.CreateGroup("output", "Output Options", func() {
    output = flags.String("output", "o", "", "Output file")
    verbose = flags.Bool("verbose", "v", false, "Verbose output")
})

// Performance group
flags.CreateGroup("performance", "Performance", func() {
    threads = flags.Int("threads", "t", 10, "Worker threads")
})

flags.Parse()
```

**Help output:**
```
Advanced tool

Usage: ./app [options]

INPUT OPTIONS:
  -i, --input string   Input file

OUTPUT OPTIONS:
  -o, --output string   Output file
  -v, --verbose         Verbose output

PERFORMANCE:
  -t, --threads int     Worker threads (default: 10)
```

### Callback Flags

Execute functions when flags are set:

```go
flags := logify.New("Callback example")

flags.Callback("update", "u", "Check for updates", func() {
    fmt.Println("Checking for updates...")
    fmt.Println("You are using the latest version!")
})

flags.Callback("version", "", "Show version", func() {
    fmt.Println("MyApp v1.0.0")
})

flags.Parse()
```

**Command line:**
```bash
./app --update    # Executes the update function
./app --version   # Executes the version function
```

### Complete Example
>>>>>>> 4a5d0ef (first commit)

```go
package main

import (
<<<<<<< HEAD
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

=======
    "fmt"
    "log"
    "time"
    "github.com/cyinnove/logify"
)

type Config struct {
    Input    *logify.StringSlice
    Output   *string
    Verbose  *bool
    Threads  *int
    Timeout  *time.Duration
    MaxSize  *logify.Size
}

func main() {
    flags := logify.New("Complete CLI Example")
    config := &Config{}

    // Input group
    flags.CreateGroup("input", "Input Options", func() {
        config.Input = flags.FileStringSlice("input", "i", 
            "Input files (supports file lists)")
    })

    // Output group
    flags.CreateGroup("output", "Output Options", func() {
        config.Output = flags.String("output", "o", "", "Output file")
        config.Verbose = flags.Bool("verbose", "v", false, "Verbose mode")
    })

    // Performance group
    flags.CreateGroup("performance", "Performance", func() {
        config.Threads = flags.Int("threads", "t", 10, "Concurrent threads")
        config.Timeout = flags.Duration("timeout", "", 30*time.Second, "Timeout")
        config.MaxSize = flags.Size("max-size", "ms", "10mb", "Max size")
    })

    // Utility
    flags.CreateGroup("utility", "Utility", func() {
        flags.Callback("version", "", "Show version", func() {
            fmt.Println("MyApp v1.0.0")
        })
    })

    if err := flags.Parse(); err != nil {
        log.Fatal(err)
    }

    // Use configuration
    runApp(config)
}
```

## API Reference

### Creating a Parser

```go
flags := logify.New(description string) *Logify
```

### Flag Methods

Logify provides two styles of flag methods:

1. **Simple methods** (e.g., `String()`, `Int()`) - Return pointers to new variables
2. **Var methods** (e.g., `StringVar()`, `IntVar()`) - Bind to existing variables

#### String Flags

```go
// Simple method - returns pointer to new string
value := flags.String(long, short, defaultValue, usage string) *string

// Var method - binds to existing variable (long name only)
flags.StringVar(field *string, long, defaultValue, usage string) *FlagData

// VarP method - binds to existing variable (short + long name)
flags.StringVarP(field *string, long, short, defaultValue, usage string) *FlagData
```

**Example:**
```go
// Using simple method
name := flags.String("name", "n", "", "Your name")
fmt.Println(*name) // Dereference pointer

// Using Var method
var name string
flags.StringVarP(&name, "name", "n", "", "Your name")
fmt.Println(name) // Use directly, no dereference needed
```

#### Bool Flags

```go
// Simple method
value := flags.Bool(long, short string, defaultValue bool, usage string) *bool

// Var method (long only)
flags.BoolVar(field *bool, long string, defaultValue bool, usage string) *FlagData

// VarP method (short + long)
flags.BoolVarP(field *bool, long, short string, defaultValue bool, usage string) *FlagData
```

#### Int Flags

```go
// Simple method
value := flags.Int(long, short string, defaultValue int, usage string) *int

// Var method (long only)
flags.IntVar(field *int, long string, defaultValue int, usage string) *FlagData

// VarP method (short + long)
flags.IntVarP(field *int, long, short string, defaultValue int, usage string) *FlagData
```

#### Duration Flags

```go
// Simple method
value := flags.Duration(long, short string, defaultValue time.Duration, usage string) *time.Duration

// Var method (long only)
flags.DurationVar(field *time.Duration, long string, defaultValue time.Duration, usage string) *FlagData

// VarP method (short + long)
flags.DurationVarP(field *time.Duration, long, short string, defaultValue time.Duration, usage string) *FlagData
```

#### Size Flags

```go
// Simple method
value := flags.Size(long, short string, defaultValue string, usage string) *logify.Size

// Var method (long only)
flags.SizeVar(field *logify.Size, long string, defaultValue string, usage string) *FlagData

// VarP method (short + long)
flags.SizeVarP(field *logify.Size, long, short string, defaultValue string, usage string) *FlagData
```

Size methods:
- `Bytes() int64` - Get size in bytes
- `KB() float64` - Get size in kilobytes
- `MB() float64` - Get size in megabytes
- `GB() float64` - Get size in gigabytes
- `TB() float64` - Get size in terabytes

#### StringSlice Flags

```go
// Simple methods
value := flags.StringSlice(long, short, usage string) *logify.StringSlice
value := flags.FileStringSlice(long, short, usage string) *logify.StringSlice

// Var methods (long only)
flags.StringSliceVar(field *logify.StringSlice, long string, defaultValue []string, usage string) *FlagData
flags.FileStringSliceVar(field *logify.StringSlice, long string, defaultValue []string, usage string) *FlagData

// VarP methods (short + long)
flags.StringSliceVarP(field *logify.StringSlice, long, short string, defaultValue []string, usage string) *FlagData
flags.FileStringSliceVarP(field *logify.StringSlice, long, short string, defaultValue []string, usage string) *FlagData
```

StringSlice methods:
- `Get() []string` - Get all values
- `Len() int` - Get number of values
- `Reset()` - Clear all values
- `String() string` - Get comma-separated string

#### Callback Flags

```go
// Simple method
flags.Callback(long, short, usage string, fn func()) *logify.FlagData

// Var method (long only)
flags.CallbackVar(fn func(), long string, usage string) *FlagData

// VarP method (short + long)
flags.CallbackVarP(fn func(), long, short string, usage string) *FlagData
```

### Group Management

```go
flags.CreateGroup(name, description string, setup func())
```

### Parsing

```go
// Parse os.Args[1:]
err := flags.Parse()

// Parse custom arguments (useful for testing)
err := flags.ParseArgs([]string{"--flag", "value"})
```

### Utility Methods

```go
// Check if parsed
parsed := flags.Parsed() bool

// Get non-flag arguments
args := flags.Args() []string

// Get count of non-flag arguments
count := flags.NArg() int
```

## Testing

Run all tests:
```bash
go test -v ./...
```

Run with coverage:
```bash
go test -cover
```

Generate coverage report:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Examples

Check out the [examples](./examples) directory:

- [Basic Example](./examples/basic) - All flag types with groups
- [Advanced Example](./examples/advanced) - Complex security scanner configuration
- [Var Methods Example](./examples/var-methods) - Using Var/VarP methods to bind flags directly

Run examples:
```bash
cd examples/basic
go run main.go --help

cd examples/advanced
go run main.go --input=../testdata/targets.txt --verbose
```

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch
3. Write tests for new features
4. Ensure all tests pass: `go test ./...`
5. Format code: `go fmt ./...`
6. Submit a pull request

## License

MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

Logify is inspired by [goflags](https://github.com/projectdiscovery/goflags) by ProjectDiscovery. We appreciate their excellent work on the original library and the Go security community.

## Support

- 📫 Issues: [GitHub Issues](https://github.com/cyinnove/logify/issues)
- 📖 Documentation: This README and code examples
- 💬 Discussions: [GitHub Discussions](https://github.com/cyinnove/logify/discussions)

---

Made with ❤️ by [cyinnove](https://github.com/cyinnove)
>>>>>>> 4a5d0ef (first commit)
