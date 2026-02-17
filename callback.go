package logify

import (
	"fmt"
	"strconv"
)

// Callback is a flag that executes a function when set
type Callback struct {
	fn func()
}

// Set implements flag.Value interface
func (c *Callback) Set(value string) error {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("failed to parse callback flag")
	}
	if v && c.fn != nil {
		c.fn()
	}
	return nil
}

// String implements flag.Value interface
func (c *Callback) String() string {
	return "false"
}

// IsBoolFlag indicates this is a boolean flag
func (c *Callback) IsBoolFlag() bool {
	return true
}
