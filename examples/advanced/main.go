package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cyinnove/logify"
)

type Config struct {
	// Input options
	Target  *string
	Input   *logify.StringSlice
	
	// Output options
	Output  *string
	Verbose *bool
	Silent  *bool
	JSON    *bool
	
	// Filter options
	Severity *logify.StringSlice
	Status   *logify.StringSlice
	
	// Performance options
	Threads    *int
	Timeout    *time.Duration
	RateLimit  *int
	MaxSize    *logify.Size
	RetryDelay *time.Duration
}

func main() {
	// Create a new Logify instance
	flags := logify.New("Advanced security scanning tool with comprehensive options")
	
	config := &Config{}
	
	// Input group
	flags.CreateGroup("input", "Input Options", func() {
		config.Target = flags.String("target", "u", "", "Single target URL to scan")
		config.Input = flags.FileStringSlice("input", "i", "File containing list of targets (one per line)")
	})
	
	// Output group  
	flags.CreateGroup("output", "Output Options", func() {
		config.Output = flags.String("output", "o", "", "Output file path")
		config.Verbose = flags.Bool("verbose", "v", false, "Verbose output mode")
		config.Silent = flags.Bool("silent", "s", false, "Silent mode (errors only)")
		config.JSON = flags.Bool("json", "j", false, "Output results in JSON format")
	})
	
	// Filter group
	flags.CreateGroup("filter", "Filter Options", func() {
		config.Severity = flags.StringSlice("severity", "sev", "Filter by severity (low,medium,high,critical)")
		config.Status = flags.StringSlice("status", "st", "Filter by HTTP status codes (200,404,500)")
	})
	
	// Performance group
	flags.CreateGroup("performance", "Performance Tuning", func() {
		config.Threads = flags.Int("threads", "t", 10, "Number of concurrent threads")
		config.Timeout = flags.Duration("timeout", "to", 30*time.Second, "Request timeout duration")
		config.RateLimit = flags.Int("rate-limit", "rl", 150, "Maximum requests per second")
		config.MaxSize = flags.Size("max-size", "ms", "5mb", "Maximum response size to process")
		config.RetryDelay = flags.Duration("retry-delay", "rd", 1*time.Second, "Delay between retries")
	})
	
	// Utility group
	flags.CreateGroup("utility", "Utility Options", func() {
		flags.Callback("update", "up", "Check for tool updates", func() {
			fmt.Println("Checking for updates...")
			fmt.Println("Current version: v1.0.0")
			fmt.Println("Latest version: v1.0.0")
			fmt.Println("✓ You are using the latest version!")
		})
		
		flags.Callback("version", "", "Display version information", func() {
			fmt.Println("Advanced Security Scanner v1.0.0")
			fmt.Println("Built with Logify - https://github.com/cyinnove/logify")
		})
	})
	
	// Parse flags
	if err := flags.Parse(); err != nil {
		log.Fatal(err)
	}
	
	// Validate input
	if *config.Target == "" && len(config.Input.Get()) == 0 {
		fmt.Println("Error: Please provide either -target or -input")
		fmt.Println("Use -help for usage information")
		return
	}
	
	// Display configuration
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║        Advanced Security Scanner Configuration            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	
	// Input section
	fmt.Println("📥 INPUT OPTIONS")
	fmt.Println(strings.Repeat("─", 60))
	if *config.Target != "" {
		fmt.Printf("  Target URL:      %s\n", *config.Target)
	}
	if len(config.Input.Get()) > 0 {
		fmt.Printf("  Input files:     %s\n", strings.Join(config.Input.Get(), ", "))
		fmt.Printf("  Total targets:   %d\n", len(config.Input.Get()))
	}
	fmt.Println()
	
	// Output section
	fmt.Println("📤 OUTPUT OPTIONS")
	fmt.Println(strings.Repeat("─", 60))
	if *config.Output != "" {
		fmt.Printf("  Output file:     %s\n", *config.Output)
	} else {
		fmt.Printf("  Output file:     stdout\n")
	}
	fmt.Printf("  Verbose mode:    %v\n", *config.Verbose)
	fmt.Printf("  Silent mode:     %v\n", *config.Silent)
	fmt.Printf("  JSON format:     %v\n", *config.JSON)
	fmt.Println()
	
	// Filters
	if len(config.Severity.Get()) > 0 || len(config.Status.Get()) > 0 {
		fmt.Println("🔍 FILTER OPTIONS")
		fmt.Println(strings.Repeat("─", 60))
		if len(config.Severity.Get()) > 0 {
			fmt.Printf("  Severity:        %s\n", strings.Join(config.Severity.Get(), ", "))
		}
		if len(config.Status.Get()) > 0 {
			fmt.Printf("  Status codes:    %s\n", strings.Join(config.Status.Get(), ", "))
		}
		fmt.Println()
	}
	
	// Performance
	fmt.Println("⚡ PERFORMANCE TUNING")
	fmt.Println(strings.Repeat("─", 60))
	fmt.Printf("  Threads:         %d\n", *config.Threads)
	fmt.Printf("  Timeout:         %v\n", *config.Timeout)
	fmt.Printf("  Rate limit:      %d req/s\n", *config.RateLimit)
	fmt.Printf("  Max size:        %d bytes (%s)\n", config.MaxSize.Bytes(), formatSize(config.MaxSize.Bytes()))
	fmt.Printf("  Retry delay:     %v\n", *config.RetryDelay)
	fmt.Println()
	
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    Starting Scan...                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	
	// Your application logic here...
	if *config.Verbose {
		fmt.Println("[VERBOSE] Initializing scanner modules...")
		fmt.Println("[VERBOSE] Loading plugins...")
		fmt.Println("[VERBOSE] Ready to scan!")
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
