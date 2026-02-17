package logify

import (
	"os"
	"testing"
)

func TestStringSlice_Set(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "single value",
			input:    []string{"value1"},
			expected: []string{"value1"},
		},
		{
			name:     "multiple values",
			input:    []string{"value1", "value2", "value3"},
			expected: []string{"value1", "value2", "value3"},
		},
		{
			name:     "comma separated",
			input:    []string{"value1,value2,value3"},
			expected: []string{"value1", "value2", "value3"},
		},
		{
			name:     "mixed comma and multiple",
			input:    []string{"value1,value2", "value3"},
			expected: []string{"value1", "value2", "value3"},
		},
		{
			name:     "with spaces",
			input:    []string{" value1 ", "value2 ", " value3"},
			expected: []string{"value1", "value2", "value3"},
		},
		{
			name:     "comma with spaces",
			input:    []string{"value1, value2 , value3"},
			expected: []string{"value1", "value2", "value3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice := NewStringSlice(false)
			for _, val := range tt.input {
				if err := slice.Set(val); err != nil {
					t.Fatalf("Set() error = %v", err)
				}
			}

			got := slice.Get()
			if len(got) != len(tt.expected) {
				t.Errorf("got %d values, want %d", len(got), len(tt.expected))
			}

			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("got[%d] = %q, want %q", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestStringSlice_FileSupport(t *testing.T) {
	// Create temporary test file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	content := "line1\nline2\n# comment\nline3\n\nline4"
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	slice := NewStringSlice(true)
	if err := slice.Set(tmpfile.Name()); err != nil {
		t.Fatalf("Set() error = %v", err)
	}

	expected := []string{"line1", "line2", "line3", "line4"}
	got := slice.Get()

	if len(got) != len(expected) {
		t.Fatalf("got %d values, want %d", len(got), len(expected))
	}

	for i, v := range got {
		if v != expected[i] {
			t.Errorf("got[%d] = %q, want %q", i, v, expected[i])
		}
	}
}

func TestStringSlice_NoFileSupport(t *testing.T) {
	// Create temporary test file
	tmpfile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tmpfile.Write([]byte("line1\nline2"))
	tmpfile.Close()

	// Without file support, should treat filename as regular value
	slice := NewStringSlice(false)
	if err := slice.Set(tmpfile.Name()); err != nil {
		t.Fatalf("Set() error = %v", err)
	}

	got := slice.Get()
	if len(got) != 1 {
		t.Fatalf("got %d values, want 1", len(got))
	}

	if got[0] != tmpfile.Name() {
		t.Errorf("got %q, want %q", got[0], tmpfile.Name())
	}
}

func TestStringSlice_String(t *testing.T) {
	slice := NewStringSlice(false)
	slice.Set("value1")
	slice.Set("value2")
	slice.Set("value3")

	expected := "value1,value2,value3"
	got := slice.String()

	if got != expected {
		t.Errorf("String() = %q, want %q", got, expected)
	}
}

func TestStringSlice_Len(t *testing.T) {
	slice := NewStringSlice(false)

	if slice.Len() != 0 {
		t.Errorf("Len() = %d, want 0", slice.Len())
	}

	slice.Set("value1")
	slice.Set("value2,value3")

	if slice.Len() != 3 {
		t.Errorf("Len() = %d, want 3", slice.Len())
	}
}

func TestStringSlice_Reset(t *testing.T) {
	slice := NewStringSlice(false)
	slice.Set("value1")
	slice.Set("value2")

	if slice.Len() != 2 {
		t.Errorf("Len() = %d, want 2", slice.Len())
	}

	slice.Reset()

	if slice.Len() != 0 {
		t.Errorf("After Reset(), Len() = %d, want 0", slice.Len())
	}

	if len(slice.Get()) != 0 {
		t.Errorf("After Reset(), Get() = %v, want empty slice", slice.Get())
	}
}

func TestStringSlice_EmptyValues(t *testing.T) {
	slice := NewStringSlice(false)

	// Empty strings should be ignored
	slice.Set("")
	slice.Set("  ")
	slice.Set(",,,")

	if slice.Len() != 0 {
		t.Errorf("Len() = %d, want 0 (empty values should be ignored)", slice.Len())
	}

	// Mixed empty and valid
	slice.Set("value1,,value2")
	expected := []string{"value1", "value2"}
	got := slice.Get()

	if len(got) != len(expected) {
		t.Fatalf("got %d values, want %d", len(got), len(expected))
	}
}
