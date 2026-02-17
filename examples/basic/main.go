package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cyinnove/logify"
)

func main() {
	// Create a new Logify instance with description
	flags := logify.New("A comprehensive example demonstrating all logify features")

	var (
		// Basic flags
		name    *string
		verbose *bool
		threads *int
		
		// Advanced flags
		timeout  *time.Duration
		maxSize  *logify.Size
		targets  *logify.StringSlice
		exclude  *logify.StringSlice
	)

	// Create organized groups
	flags.CreateGroup("input", "Input Options", func() {
		name = flags.String("name", "n", "", "Target name")
		targets = flags.FileStringSlice("target", "t", "Target URLs (supports file input)")
		exclude = flags.FileStringSlice("exclude", "e", "URLs to exclude (supports file input)")
	})

	flags.CreateGroup("config", "Configuration", func() {
		verbose = flags.Bool("verbose", "v", false, "Enable verbose output")
		threads = flags.Int("threads", "c", 10, "Number of concurrent threads")
		timeout = flags.Duration("timeout", "to", 30*time.Second, "Request timeout")
		maxSize = flags.Size("max-size", "ms", "10mb", "Maximum response size")
	})

	flags.CreateGroup("utility", "Utility", func() {
		flags.Callback("update", "u", "Check for updates", func() {
			fmt.Println("Checking for updates...")
			fmt.Println("You are using the latest version!")
		})
	})

	// Parse the flags
	if err := flags.Parse(); err != nil {
		log.Fatal(err)
	}

	// Display parsed configuration
	fmt.Println("=== Configuration ===")
	fmt.Println()

	// Input options
	fmt.Println("Input:")
	if *name != "" {
		fmt.Printf("  Name: %s\n", *name)
	}
	if len(targets.Get()) > 0 {
		fmt.Printf("  Targets: %v\n", targets.Get())
	}
	if len(exclude.Get()) > 0 {
		fmt.Printf("  Exclude: %v\n", exclude.Get())
	}
	fmt.Println()

	// Configuration
	fmt.Println("Configuration:")
	fmt.Printf("  Verbose: %v\n", *verbose)
	fmt.Printf("  Threads: %d\n", *threads)
	fmt.Printf("  Timeout: %v\n", *timeout)
	fmt.Printf("  Max Size: %d bytes (%s)\n", maxSize.Bytes(), formatSize(maxSize.Bytes()))
	fmt.Println()

	// Show remaining non-flag arguments
	if flags.NArg() > 0 {
		fmt.Printf("Additional arguments: %v\n", flags.Args())
	}

	if *verbose {
		fmt.Println("\n=== Starting application in verbose mode ===")
	}
}

func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGT"[exp])
}
