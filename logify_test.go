package logify

import (
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	description := "Test application"
	l := New(description)

	if l.description != description {
		t.Errorf("description = %q, want %q", l.description, description)
	}

	if l.flagSet == nil {
		t.Error("flagSet should not be nil")
	}

	if l.flagData == nil {
		t.Error("flagData should not be nil")
	}

	if len(l.groups) != 0 {
		t.Errorf("groups length = %d, want 0", len(l.groups))
	}
}

func TestLogify_String(t *testing.T) {
	l := New("test")
	name := l.String("name", "n", "default", "test name")

	if name == nil {
		t.Fatal("String() returned nil")
	}

	if *name != "default" {
		t.Errorf("default value = %q, want \"default\"", *name)
	}

	err := l.ParseArgs([]string{"--name", "John"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *name != "John" {
		t.Errorf("parsed value = %q, want \"John\"", *name)
	}
}

func TestLogify_StringShort(t *testing.T) {
	l := New("test")
	name := l.String("name", "n", "default", "test name")

	err := l.ParseArgs([]string{"-n", "Jane"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *name != "Jane" {
		t.Errorf("parsed value = %q, want \"Jane\"", *name)
	}
}

func TestLogify_Bool(t *testing.T) {
	l := New("test")
	verbose := l.Bool("verbose", "v", false, "verbose mode")

	err := l.ParseArgs([]string{"--verbose"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if !*verbose {
		t.Error("verbose should be true")
	}
}

func TestLogify_Int(t *testing.T) {
	l := New("test")
	count := l.Int("count", "c", 1, "count")

	err := l.ParseArgs([]string{"--count", "42"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *count != 42 {
		t.Errorf("count = %d, want 42", *count)
	}
}

func TestLogify_Duration(t *testing.T) {
	l := New("test")
	timeout := l.Duration("timeout", "t", 10*time.Second, "timeout")

	err := l.ParseArgs([]string{"--timeout", "30s"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *timeout != 30*time.Second {
		t.Errorf("timeout = %v, want 30s", *timeout)
	}
}

func TestLogify_Size(t *testing.T) {
	l := New("test")
	maxSize := l.Size("max-size", "ms", "1mb", "max size")

	err := l.ParseArgs([]string{"--max-size", "5mb"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	expected := int64(5 * 1024 * 1024)
	if maxSize.Bytes() != expected {
		t.Errorf("max-size = %d, want %d", maxSize.Bytes(), expected)
	}
}

func TestLogify_StringSlice(t *testing.T) {
	l := New("test")
	tags := l.StringSlice("tag", "t", "tags")

	err := l.ParseArgs([]string{"--tag", "go", "--tag", "cli"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	values := tags.Get()
	if len(values) != 2 {
		t.Errorf("got %d tags, want 2", len(values))
	}

	if values[0] != "go" || values[1] != "cli" {
		t.Errorf("tags = %v, want [go cli]", values)
	}
}

func TestLogify_FileStringSlice(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	content := "target1\ntarget2\ntarget3\n# comment\ntarget4"
	tmpfile.Write([]byte(content))
	tmpfile.Close()

	l := New("test")
	targets := l.FileStringSlice("target", "t", "targets")

	err = l.ParseArgs([]string{"--target", tmpfile.Name()})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	values := targets.Get()
	expected := []string{"target1", "target2", "target3", "target4"}

	if len(values) != len(expected) {
		t.Fatalf("got %d targets, want %d", len(values), len(expected))
	}

	for i, v := range values {
		if v != expected[i] {
			t.Errorf("target[%d] = %q, want %q", i, v, expected[i])
		}
	}
}

func TestLogify_Callback(t *testing.T) {
	l := New("test")
	called := false

	l.Callback("test", "t", "test callback", func() {
		called = true
	})

	err := l.ParseArgs([]string{"--test"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if !called {
		t.Error("callback should have been called")
	}
}

func TestLogify_CreateGroup(t *testing.T) {
	l := New("test with groups")

	var inputFlag *string
	var outputFlag *string

	l.CreateGroup("input", "Input Options", func() {
		inputFlag = l.String("input", "i", "", "input file")
	})

	l.CreateGroup("output", "Output Options", func() {
		outputFlag = l.String("output", "o", "", "output file")
	})

	if len(l.groups) != 2 {
		t.Errorf("got %d groups, want 2", len(l.groups))
	}

	err := l.ParseArgs([]string{"--input", "in.txt", "--output", "out.txt"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *inputFlag != "in.txt" {
		t.Errorf("input = %q, want \"in.txt\"", *inputFlag)
	}

	if *outputFlag != "out.txt" {
		t.Errorf("output = %q, want \"out.txt\"", *outputFlag)
	}
}

func TestLogify_MultipleFlags(t *testing.T) {
	l := New("test")
	name := l.String("name", "n", "", "name")
	verbose := l.Bool("verbose", "v", false, "verbose")
	count := l.Int("count", "c", 1, "count")
	tags := l.StringSlice("tag", "t", "tags")

	err := l.ParseArgs([]string{
		"--name", "Alice",
		"-v",
		"--count", "5",
		"--tag", "go,cli",
		"--tag", "tool",
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *name != "Alice" {
		t.Errorf("name = %q, want \"Alice\"", *name)
	}

	if !*verbose {
		t.Error("verbose should be true")
	}

	if *count != 5 {
		t.Errorf("count = %d, want 5", *count)
	}

	tagValues := tags.Get()
	if len(tagValues) != 3 {
		t.Errorf("got %d tags, want 3", len(tagValues))
	}
}

func TestLogify_DefaultValues(t *testing.T) {
	l := New("test")
	name := l.String("name", "n", "John", "test name")
	count := l.Int("count", "c", 10, "count")
	verbose := l.Bool("verbose", "v", true, "verbose")
	timeout := l.Duration("timeout", "t", 30*time.Second, "timeout")

	err := l.ParseArgs([]string{})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *name != "John" {
		t.Errorf("name = %q, want \"John\"", *name)
	}

	if *count != 10 {
		t.Errorf("count = %d, want 10", *count)
	}

	if !*verbose {
		t.Error("verbose should be true (default)")
	}

	if *timeout != 30*time.Second {
		t.Errorf("timeout = %v, want 30s", *timeout)
	}
}

func TestLogify_Args(t *testing.T) {
	l := New("test")
	l.String("name", "n", "", "name")

	err := l.ParseArgs([]string{"--name", "test", "arg1", "arg2", "arg3"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	args := l.Args()
	if len(args) != 3 {
		t.Errorf("got %d args, want 3", len(args))
	}

	expected := []string{"arg1", "arg2", "arg3"}
	for i, arg := range args {
		if arg != expected[i] {
			t.Errorf("arg[%d] = %q, want %q", i, arg, expected[i])
		}
	}

	if l.NArg() != 3 {
		t.Errorf("NArg() = %d, want 3", l.NArg())
	}
}

func TestLogify_Parsed(t *testing.T) {
	l := New("test")
	l.String("name", "n", "", "name")

	if l.Parsed() {
		t.Error("Parsed() should be false before Parse()")
	}

	l.ParseArgs([]string{"--name", "test"})

	if !l.Parsed() {
		t.Error("Parsed() should be true after Parse()")
	}
}

func TestLogify_NoShortFlag(t *testing.T) {
	l := New("test")
	name := l.String("name", "", "default", "name without short flag")

	err := l.ParseArgs([]string{"--name", "value"})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *name != "value" {
		t.Errorf("name = %q, want \"value\"", *name)
	}
}

func TestLogify_MixedShortAndLong(t *testing.T) {
	l := New("test")
	name := l.String("name", "n", "", "name")
	verbose := l.Bool("verbose", "v", false, "verbose")
	count := l.Int("count", "c", 0, "count")

	err := l.ParseArgs([]string{
		"--name", "Alice",
		"-v",
		"-c", "10",
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *name != "Alice" {
		t.Errorf("name = %q, want \"Alice\"", *name)
	}

	if !*verbose {
		t.Error("verbose should be true")
	}

	if *count != 10 {
		t.Errorf("count = %d, want 10", *count)
	}
}

func TestLogify_GroupsWithUngroupedFlags(t *testing.T) {
	l := New("test")

	var groupedFlag *string
	var ungroupedFlag *string

	l.CreateGroup("test", "Test Group", func() {
		groupedFlag = l.String("grouped", "g", "", "grouped flag")
	})

	ungroupedFlag = l.String("ungrouped", "u", "", "ungrouped flag")

	err := l.ParseArgs([]string{
		"--grouped", "value1",
		"--ungrouped", "value2",
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if *groupedFlag != "value1" {
		t.Errorf("grouped = %q, want \"value1\"", *groupedFlag)
	}

	if *ungroupedFlag != "value2" {
		t.Errorf("ungrouped = %q, want \"value2\"", *ungroupedFlag)
	}
}

func TestLogify_ComplexScenario(t *testing.T) {
	l := New("complex test application")

	type Config struct {
		Input   *StringSlice
		Output  *string
		Verbose *bool
		Threads *int
		Timeout *time.Duration
		MaxSize *Size
	}

	config := &Config{}

	l.CreateGroup("input", "Input", func() {
		config.Input = l.FileStringSlice("input", "i", "input files")
	})

	l.CreateGroup("output", "Output", func() {
		config.Output = l.String("output", "o", "result.txt", "output file")
		config.Verbose = l.Bool("verbose", "v", false, "verbose mode")
	})

	l.CreateGroup("performance", "Performance", func() {
		config.Threads = l.Int("threads", "t", 10, "threads")
		config.Timeout = l.Duration("timeout", "to", 30*time.Second, "timeout")
		config.MaxSize = l.Size("max-size", "ms", "10mb", "max size")
	})

	// Create test file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Write([]byte("item1\nitem2\nitem3"))
	tmpfile.Close()

	err = l.ParseArgs([]string{
		"--input", tmpfile.Name(),
		"--output", "custom.txt",
		"--verbose",
		"--threads", "20",
		"--timeout", "1m",
		"--max-size", "50mb",
		"extra1", "extra2",
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Validate all values
	if len(config.Input.Get()) != 3 {
		t.Errorf("input items = %d, want 3", len(config.Input.Get()))
	}

	if *config.Output != "custom.txt" {
		t.Errorf("output = %q, want \"custom.txt\"", *config.Output)
	}

	if !*config.Verbose {
		t.Error("verbose should be true")
	}

	if *config.Threads != 20 {
		t.Errorf("threads = %d, want 20", *config.Threads)
	}

	if *config.Timeout != time.Minute {
		t.Errorf("timeout = %v, want 1m", *config.Timeout)
	}

	if config.MaxSize.Bytes() != 50*1024*1024 {
		t.Errorf("max-size = %d, want %d", config.MaxSize.Bytes(), 50*1024*1024)
	}

	if l.NArg() != 2 {
		t.Errorf("NArg() = %d, want 2", l.NArg())
	}
}
