package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cyinnove/logify"
)

// This example demonstrates the Var and VarP methods
// which allow you to bind flags directly to existing variables
func main() {
	flags := logify.New("Example of Var and VarP methods")

	// Declare your config variables
	var (
		// Input options
		inputFile  string
		targets    logify.StringSlice
		
		// Output options
		outputFile string
		verbose    bool
		silent     bool
		
		// Performance options
		threads    int
		timeout    time.Duration
		maxSize    logify.Size
	)

	// Use Var methods to bind flags directly to variables
	flags.CreateGroup("input", "Input Options", func() {
		// StringVarP - binds string flag with short and long name
		flags.StringVarP(&inputFile, "input", "i", "", "Input file path")
		
		// FileStringSliceVar - binds string slice with file support (long name only)
		flags.FileStringSliceVar(&targets, "target", []string{}, "Target URLs from file")
	})

	flags.CreateGroup("output", "Output Options", func() {
		// StringVar - binds string flag (long name only)
		flags.StringVar(&outputFile, "output", "results.txt", "Output file path")
		
		// BoolVarP - binds bool flag with short and long name
		flags.BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
		
		// BoolVar - binds bool flag (long name only)
		flags.BoolVar(&silent, "silent", false, "Silent mode (errors only)")
	})

	flags.CreateGroup("performance", "Performance Options", func() {
		// IntVarP - binds int flag with short and long name
		flags.IntVarP(&threads, "threads", "t", 10, "Number of concurrent threads")
		
		// DurationVar - binds duration flag (long name only)
		flags.DurationVar(&timeout, "timeout", 30*time.Second, "Request timeout")
		
		// SizeVarP - binds size flag with short and long name
		flags.SizeVarP(&maxSize, "max-size", "ms", "10mb", "Maximum response size")
	})

	// CallbackVar - execute function when flag is set
	flags.CreateGroup("utility", "Utility", func() {
		flags.CallbackVarP(func() {
			fmt.Println("MyApp v1.0.0")
			fmt.Println("Built with Logify")
		}, "version", "V", "Show version information")
	})

	// Parse command line arguments
	if err := flags.Parse(); err != nil {
		log.Fatal(err)
	}

	// Now you can use the variables directly - no need to dereference pointers!
	fmt.Println("=== Configuration ===")
	fmt.Println()

	fmt.Println("Input:")
	if inputFile != "" {
		fmt.Printf("  Input file: %s\n", inputFile)
	}
	if targets.Len() > 0 {
		fmt.Printf("  Targets: %v\n", targets.Get())
	}
	fmt.Println()

	fmt.Println("Output:")
	fmt.Printf("  Output file: %s\n", outputFile)
	fmt.Printf("  Verbose: %v\n", verbose)
	fmt.Printf("  Silent: %v\n", silent)
	fmt.Println()

	fmt.Println("Performance:")
	fmt.Printf("  Threads: %d\n", threads)
	fmt.Printf("  Timeout: %v\n", timeout)
	fmt.Printf("  Max size: %d bytes (%.2f MB)\n", maxSize.Bytes(), maxSize.MB())
	fmt.Println()

	if verbose {
		fmt.Println("=== Running in verbose mode ===")
	}
}
