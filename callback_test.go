package logify

import (
	"testing"
)

func TestCallback_Set(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		shouldCall  bool
		expectError bool
	}{
		{"true", "true", true, false},
		{"1", "1", true, false},
		{"false", "false", false, false},
		{"0", "0", false, false},
		{"invalid", "invalid", false, true},
		{"empty", "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			callback := &Callback{
				fn: func() {
					called = true
				},
			}

			err := callback.Set(tt.value)

			if (err != nil) != tt.expectError {
				t.Errorf("Set() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError && called != tt.shouldCall {
				t.Errorf("callback called = %v, want %v", called, tt.shouldCall)
			}
		})
	}
}

func TestCallback_String(t *testing.T) {
	callback := &Callback{fn: func() {}}

	if callback.String() != "false" {
		t.Errorf("String() = %q, want \"false\"", callback.String())
	}
}

func TestCallback_IsBoolFlag(t *testing.T) {
	callback := &Callback{fn: func() {}}

	if !callback.IsBoolFlag() {
		t.Error("IsBoolFlag() = false, want true")
	}
}

func TestCallback_NilFunction(t *testing.T) {
	callback := &Callback{fn: nil}

	// Should not panic even with nil function
	err := callback.Set("true")
	if err != nil {
		t.Errorf("Set() with nil function error = %v, want nil", err)
	}
}

func TestCallback_MultipleCalls(t *testing.T) {
	callCount := 0
	callback := &Callback{
		fn: func() {
			callCount++
		},
	}

	// Call multiple times
	callback.Set("true")
	callback.Set("false") // Should not increment
	callback.Set("true")

	if callCount != 2 {
		t.Errorf("function called %d times, want 2", callCount)
	}
}

func TestCallback_WithState(t *testing.T) {
	var result string
	callback := &Callback{
		fn: func() {
			result = "callback executed"
		},
	}

	if result != "" {
		t.Error("result should be empty before callback")
	}

	callback.Set("true")

	if result != "callback executed" {
		t.Errorf("result = %q, want \"callback executed\"", result)
	}
}
