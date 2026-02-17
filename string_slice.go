package logify

import (
	"strings"
)

// StringSlice is a custom type for string slices with file support
type StringSlice struct {
	values      []string
	supportFile bool
}

// NewStringSlice creates a new StringSlice
func NewStringSlice(supportFile bool) *StringSlice {
	return &StringSlice{
		values:      []string{},
		supportFile: supportFile,
	}
}

// String implements flag.Value interface
func (s *StringSlice) String() string {
	return strings.Join(s.values, ",")
}

// Set implements flag.Value interface
func (s *StringSlice) Set(value string) error {
	// Check if it's a file
	if s.supportFile && fileExists(value) {
		lines, err := readLines(value)
		if err != nil {
			return err
		}
		s.values = append(s.values, lines...)
		return nil
	}

	// Support comma-separated values
	if strings.Contains(value, ",") {
		parts := strings.Split(value, ",")
		for _, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				s.values = append(s.values, trimmed)
			}
		}
	} else {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			s.values = append(s.values, trimmed)
		}
	}
	return nil
}

// Get returns the slice values
func (s *StringSlice) Get() []string {
	return s.values
}

// Len returns the number of elements
func (s *StringSlice) Len() int {
	return len(s.values)
}

// Reset clears all values
func (s *StringSlice) Reset() {
	s.values = []string{}
}
