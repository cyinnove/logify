package logify

import (
	"fmt"
	"strconv"
	"strings"
)

// Size represents a file size in bytes
type Size int64

// Set implements flag.Value interface
func (s *Size) Set(value string) error {
	size, err := parseSize(value)
	if err != nil {
		return err
	}
	*s = Size(size)
	return nil
}

// String implements flag.Value interface
func (s *Size) String() string {
	return strconv.FormatInt(int64(*s), 10)
}

// Bytes returns the size in bytes
func (s *Size) Bytes() int64 {
	return int64(*s)
}

// KB returns the size in kilobytes
func (s *Size) KB() float64 {
	return float64(*s) / 1024
}

// MB returns the size in megabytes
func (s *Size) MB() float64 {
	return float64(*s) / (1024 * 1024)
}

// GB returns the size in gigabytes
func (s *Size) GB() float64 {
	return float64(*s) / (1024 * 1024 * 1024)
}

// TB returns the size in terabytes
func (s *Size) TB() float64 {
	return float64(*s) / (1024 * 1024 * 1024 * 1024)
}

// parseSize parses a size string like "10mb", "5gb" into bytes
func parseSize(s string) (int64, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return 0, nil
	}

	// Check for suffixes in order from longest to shortest
	suffixes := []struct {
		suffix     string
		multiplier int64
	}{
		{"tb", 1024 * 1024 * 1024 * 1024},
		{"gb", 1024 * 1024 * 1024},
		{"mb", 1024 * 1024},
		{"kb", 1024},
		{"b", 1},
	}

	for _, sf := range suffixes {
		if strings.HasSuffix(s, sf.suffix) {
			numStr := strings.TrimSuffix(s, sf.suffix)
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid size format: %s", s)
			}
			return int64(num * float64(sf.multiplier)), nil
		}
	}

	// No suffix, assume bytes
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid size format: %s", s)
	}
	return num, nil
}
