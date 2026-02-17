package logify

import (
	"os"
	"testing"
	"time"
)

// Test StringVar and StringVarP
func TestLogify_StringVar(t *testing.T) {
	flags := New("test")
	var name string

	flags.StringVar(&name, "name", "default", "Name")

	err := flags.ParseArgs([]string{"--name", "Alice"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if name != "Alice" {
		t.Errorf("Expected 'Alice', got '%s'", name)
	}
}

func TestLogify_StringVarP(t *testing.T) {
	flags := New("test")
	var name string

	flags.StringVarP(&name, "name", "n", "default", "Name")

	// Test with short flag
	err := flags.ParseArgs([]string{"-n", "Bob"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if name != "Bob" {
		t.Errorf("Expected 'Bob', got '%s'", name)
	}
}

// Test BoolVar and BoolVarP
func TestLogify_BoolVar(t *testing.T) {
	flags := New("test")
	var verbose bool

	flags.BoolVar(&verbose, "verbose", false, "Verbose mode")

	err := flags.ParseArgs([]string{"--verbose"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if !verbose {
		t.Errorf("Expected true, got false")
	}
}

func TestLogify_BoolVarP(t *testing.T) {
	flags := New("test")
	var debug bool

	flags.BoolVarP(&debug, "debug", "d", false, "Debug mode")

	// Test with short flag
	err := flags.ParseArgs([]string{"-d"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if !debug {
		t.Errorf("Expected true, got false")
	}
}

// Test IntVar and IntVarP
func TestLogify_IntVar(t *testing.T) {
	flags := New("test")
	var count int

	flags.IntVar(&count, "count", 10, "Count")

	err := flags.ParseArgs([]string{"--count", "42"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if count != 42 {
		t.Errorf("Expected 42, got %d", count)
	}
}

func TestLogify_IntVarP(t *testing.T) {
	flags := New("test")
	var threads int

	flags.IntVarP(&threads, "threads", "t", 5, "Threads")

	// Test with short flag
	err := flags.ParseArgs([]string{"-t", "20"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if threads != 20 {
		t.Errorf("Expected 20, got %d", threads)
	}
}

// Test DurationVar and DurationVarP
func TestLogify_DurationVar(t *testing.T) {
	flags := New("test")
	var timeout time.Duration

	flags.DurationVar(&timeout, "timeout", 30*time.Second, "Timeout")

	err := flags.ParseArgs([]string{"--timeout", "1m"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if timeout != 1*time.Minute {
		t.Errorf("Expected 1m, got %v", timeout)
	}
}

func TestLogify_DurationVarP(t *testing.T) {
	flags := New("test")
	var delay time.Duration

	flags.DurationVarP(&delay, "delay", "d", 10*time.Second, "Delay")

	// Test with short flag
	err := flags.ParseArgs([]string{"-d", "5s"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if delay != 5*time.Second {
		t.Errorf("Expected 5s, got %v", delay)
	}
}

// Test SizeVar and SizeVarP
func TestLogify_SizeVar(t *testing.T) {
	flags := New("test")
	var maxSize Size

	flags.SizeVar(&maxSize, "max-size", "1mb", "Max size")

	err := flags.ParseArgs([]string{"--max-size", "10mb"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	expected := int64(10 * 1024 * 1024)
	if maxSize.Bytes() != expected {
		t.Errorf("Expected %d, got %d", expected, maxSize.Bytes())
	}
}

func TestLogify_SizeVarP(t *testing.T) {
	flags := New("test")
	var fileSize Size

	flags.SizeVarP(&fileSize, "size", "s", "5mb", "File size")

	// Test with short flag
	err := flags.ParseArgs([]string{"-s", "2gb"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	expected := int64(2 * 1024 * 1024 * 1024)
	if fileSize.Bytes() != expected {
		t.Errorf("Expected %d, got %d", expected, fileSize.Bytes())
	}
}

// Test StringSliceVar and StringSliceVarP
func TestLogify_StringSliceVar(t *testing.T) {
	flags := New("test")
	tags := NewStringSlice(false)

	flags.StringSliceVar(tags, "tag", []string{"default"}, "Tags")

	err := flags.ParseArgs([]string{"--tag", "go", "--tag", "cli"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	values := tags.Get()
	// Should have default + new values
	if len(values) < 2 {
		t.Errorf("Expected at least 2 values, got %d: %v", len(values), values)
	}
}

func TestLogify_StringSliceVarP(t *testing.T) {
	flags := New("test")
	hosts := NewStringSlice(false)

	flags.StringSliceVarP(hosts, "host", "h", []string{}, "Hosts")

	// Test with short flag
	err := flags.ParseArgs([]string{"-h", "localhost", "-h", "example.com"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	values := hosts.Get()
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}
}

// Test FileStringSliceVar and FileStringSliceVarP
func TestLogify_FileStringSliceVar(t *testing.T) {
	// Create temporary test file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	content := "line1\nline2\nline3"
	tmpfile.Write([]byte(content))
	tmpfile.Close()

	flags := New("test")
	targets := NewStringSlice(false) // Will be enabled by FileStringSliceVar

	flags.FileStringSliceVar(targets, "target", []string{}, "Targets")

	err = flags.ParseArgs([]string{"--target", tmpfile.Name()})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	values := targets.Get()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}
}

func TestLogify_FileStringSliceVarP(t *testing.T) {
	// Create temporary test file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	content := "url1\nurl2"
	tmpfile.Write([]byte(content))
	tmpfile.Close()

	flags := New("test")
	urls := NewStringSlice(false)

	flags.FileStringSliceVarP(urls, "url", "u", []string{}, "URLs")

	// Test with short flag
	err = flags.ParseArgs([]string{"-u", tmpfile.Name()})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	values := urls.Get()
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}
}

// Test CallbackVar and CallbackVarP
func TestLogify_CallbackVar(t *testing.T) {
	flags := New("test")
	called := false

	flags.CallbackVar(func() {
		called = true
	}, "test", "Test callback")

	err := flags.ParseArgs([]string{"--test"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if !called {
		t.Errorf("Expected callback to be called")
	}
}

func TestLogify_CallbackVarP(t *testing.T) {
	flags := New("test")
	result := ""

	flags.CallbackVarP(func() {
		result = "executed"
	}, "exec", "e", "Execute")

	// Test with short flag
	err := flags.ParseArgs([]string{"-e"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if result != "executed" {
		t.Errorf("Expected 'executed', got '%s'", result)
	}
}

// Test mixed usage with groups
func TestLogify_VarMethodsWithGroups(t *testing.T) {
	flags := New("test")

	var (
		input   string
		output  string
		verbose bool
		threads int
	)

	flags.CreateGroup("input", "Input", func() {
		flags.StringVarP(&input, "input", "i", "", "Input file")
	})

	flags.CreateGroup("output", "Output", func() {
		flags.StringVarP(&output, "output", "o", "", "Output file")
		flags.BoolVarP(&verbose, "verbose", "v", false, "Verbose")
	})

	flags.CreateGroup("performance", "Performance", func() {
		flags.IntVarP(&threads, "threads", "t", 10, "Threads")
	})

	err := flags.ParseArgs([]string{
		"-i", "in.txt",
		"-o", "out.txt",
		"-v",
		"-t", "20",
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if input != "in.txt" {
		t.Errorf("input: expected 'in.txt', got '%s'", input)
	}
	if output != "out.txt" {
		t.Errorf("output: expected 'out.txt', got '%s'", output)
	}
	if !verbose {
		t.Errorf("verbose: expected true, got false")
	}
	if threads != 20 {
		t.Errorf("threads: expected 20, got %d", threads)
	}
}

// Test default values for Var methods
func TestLogify_VarMethodsDefaults(t *testing.T) {
	flags := New("test")

	var (
		name    string
		count   int
		enabled bool
		timeout time.Duration
	)

	flags.StringVar(&name, "name", "default-name", "Name")
	flags.IntVar(&count, "count", 42, "Count")
	flags.BoolVar(&enabled, "enabled", true, "Enabled")
	flags.DurationVar(&timeout, "timeout", 30*time.Second, "Timeout")

	// Parse with no arguments - should use defaults
	err := flags.ParseArgs([]string{})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if name != "default-name" {
		t.Errorf("name: expected 'default-name', got '%s'", name)
	}
	if count != 42 {
		t.Errorf("count: expected 42, got %d", count)
	}
	if !enabled {
		t.Errorf("enabled: expected true, got false")
	}
	if timeout != 30*time.Second {
		t.Errorf("timeout: expected 30s, got %v", timeout)
	}
}
