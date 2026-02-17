package logify

import (
	"os"
	"testing"
	"time"
)

func TestFileExists(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"existing file", tmpfile.Name(), true},
		{"non-existing file", "/non/existing/file.txt", false},
		{"empty path", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fileExists(tt.path)
			if got != tt.expected {
				t.Errorf("fileExists(%q) = %v, want %v", tt.path, got, tt.expected)
			}
		})
	}
}

func TestReadLines(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
		wantErr  bool
	}{
		{
			name:     "simple lines",
			content:  "line1\nline2\nline3",
			expected: []string{"line1", "line2", "line3"},
			wantErr:  false,
		},
		{
			name:     "with comments",
			content:  "line1\n# comment\nline2\n#another comment\nline3",
			expected: []string{"line1", "line2", "line3"},
			wantErr:  false,
		},
		{
			name:     "with empty lines",
			content:  "line1\n\n\nline2\n\nline3",
			expected: []string{"line1", "line2", "line3"},
			wantErr:  false,
		},
		{
			name:     "with whitespace",
			content:  "  line1  \n\t\nline2\n   \nline3  ",
			expected: []string{"line1", "line2", "line3"},
			wantErr:  false,
		},
		{
			name:     "mixed",
			content:  "line1\n\n# comment\n  line2  \n\nline3\n# end",
			expected: []string{"line1", "line2", "line3"},
			wantErr:  false,
		},
		{
			name:     "empty file",
			content:  "",
			expected: []string{},
			wantErr:  false,
		},
		{
			name:     "only comments",
			content:  "# comment1\n# comment2\n# comment3",
			expected: []string{},
			wantErr:  false,
		},
		{
			name:     "only empty lines",
			content:  "\n\n\n\n",
			expected: []string{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpfile, err := os.CreateTemp("", "test-*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tt.content)); err != nil {
				t.Fatal(err)
			}
			tmpfile.Close()

			got, err := readLines(tmpfile.Name())

			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("readLines() returned %d lines, want %d", len(got), len(tt.expected))
				return
			}

			for i, line := range got {
				if line != tt.expected[i] {
					t.Errorf("line[%d] = %q, want %q", i, line, tt.expected[i])
				}
			}
		})
	}
}

func TestReadLines_NonExistentFile(t *testing.T) {
	_, err := readLines("/non/existent/file.txt")
	if err == nil {
		t.Error("readLines() should return error for non-existent file")
	}
}

func TestIsZeroValue(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"empty string", "", true},
		{"non-empty string", "hello", false},
		{"zero int", 0, true},
		{"non-zero int", 42, false},
		{"false bool", false, true},
		{"true bool", true, false},
		{"zero duration", time.Duration(0), true},
		{"non-zero duration", 5 * time.Second, false},
		{"nil", nil, false},
		{"empty slice", []string{}, false}, // Not handled, returns false
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isZeroValue(tt.value)
			if got != tt.expected {
				t.Errorf("isZeroValue(%v) = %v, want %v", tt.value, got, tt.expected)
			}
		})
	}
}

func TestReadLines_LargeFile(t *testing.T) {
	// Create a file with many lines
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	lineCount := 1000
	for i := 0; i < lineCount; i++ {
		tmpfile.WriteString("line\n")
	}
	tmpfile.Close()

	got, err := readLines(tmpfile.Name())
	if err != nil {
		t.Fatalf("readLines() error = %v", err)
	}

	if len(got) != lineCount {
		t.Errorf("readLines() returned %d lines, want %d", len(got), lineCount)
	}
}
