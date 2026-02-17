package logify

import (
	"testing"
)

func TestSize_Set(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		wantErr  bool
	}{
		{"bytes", "100b", 100, false},
		{"bytes no suffix", "100", 100, false},
		{"kilobytes", "1kb", 1024, false},
		{"megabytes", "2mb", 2 * 1024 * 1024, false},
		{"gigabytes", "1gb", 1024 * 1024 * 1024, false},
		{"terabytes", "1tb", 1024 * 1024 * 1024 * 1024, false},
		{"decimal kb", "1.5kb", 1536, false},
		{"decimal mb", "2.5mb", 2621440, false},
		{"zero", "0", 0, false},
		{"empty", "", 0, false},
		{"uppercase", "5MB", 5 * 1024 * 1024, false},
		{"mixed case", "10Mb", 10 * 1024 * 1024, false},
		{"with spaces", " 5 mb ", 5 * 1024 * 1024, false},
		{"invalid format", "abc", 0, true},
		{"invalid number", "xmb", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var size Size
			err := size.Set(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && size.Bytes() != tt.expected {
				t.Errorf("Bytes() = %d, want %d", size.Bytes(), tt.expected)
			}
		})
	}
}

func TestSize_String(t *testing.T) {
	var size Size
	size.Set("1024")

	got := size.String()
	expected := "1024"

	if got != expected {
		t.Errorf("String() = %q, want %q", got, expected)
	}
}

func TestSize_Conversions(t *testing.T) {
	var size Size
	size.Set("1024kb") // 1 MB

	tests := []struct {
		name     string
		got      float64
		expected float64
	}{
		{"Bytes", float64(size.Bytes()), 1024 * 1024},
		{"KB", size.KB(), 1024},
		{"MB", size.MB(), 1},
		{"GB", size.GB(), 1.0 / 1024},
		{"TB", size.TB(), 1.0 / (1024 * 1024)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %f, want %f", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestSize_LargeValues(t *testing.T) {
	tests := []struct {
		name  string
		input string
		mb    float64
	}{
		{"100mb", "100mb", 100},
		{"1gb", "1gb", 1024},
		{"5gb", "5gb", 5 * 1024},
		{"1tb", "1tb", 1024 * 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var size Size
			if err := size.Set(tt.input); err != nil {
				t.Fatalf("Set() error = %v", err)
			}

			got := size.MB()
			if got != tt.mb {
				t.Errorf("MB() = %f, want %f", got, tt.mb)
			}
		})
	}
}

func TestSize_ZeroValue(t *testing.T) {
	var size Size

	if size.Bytes() != 0 {
		t.Errorf("Zero value Bytes() = %d, want 0", size.Bytes())
	}

	if size.KB() != 0 {
		t.Errorf("Zero value KB() = %f, want 0", size.KB())
	}

	if size.String() != "0" {
		t.Errorf("Zero value String() = %q, want \"0\"", size.String())
	}
}

func TestParseSize_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		wantErr  bool
	}{
		{"zero bytes", "0b", 0, false},
		{"zero kb", "0kb", 0, false},
		{"fractional bytes", "100.5b", 100, false},
		{"very small decimal", "0.001kb", 1, false},
		{"invalid suffix", "10xy", 0, true},
		{"only suffix", "mb", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSize(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.expected {
				t.Errorf("parseSize() = %d, want %d", got, tt.expected)
			}
		})
	}
}
