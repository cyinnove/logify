package logify

import "fmt"

// Formatter struct represents the log message formatter including the color.
type Formatter struct {
	LogHolder string
	Message   string
	Args      []interface{}
	Color     string // Holds the color for the holder
}

// Format formats the log message using the specified color and log holder.
// If no color is specified, it defaults to no color.
// It returns the formatted log message as a string.
func (f *Formatter) Format() string {
	color := f.Color
	if color == "" {
		color = Colors[Reset] // Default to no color if none is specified
	}
	return fmt.Sprintf("[%s%s%s] %s", color, f.LogHolder, Colors[Reset], fmt.Sprintf(f.Message, f.Args...))
}

// Log prints the formatted log message to the console.
func (f *Formatter) Log() {
	fmt.Println(f.Format())
}

// SetHolder sets the log holder with an optional color.
func (f *Formatter) SetHolder(holder string) {
	f.LogHolder = holder
}

// SetMessage sets the log message and arguments.
func (f *Formatter) SetMessage(msg string, args ...interface{}) {
	f.Message = fmt.Sprintf(msg, args...)
	f.Args = nil
}

// SetColor sets the color of the formatter.
// It takes a Color parameter and assigns the corresponding color value from the Colors map to the Formatter's Color field.
func (f *Formatter) SetColor(color Color) {
	f.Color = Colors[color]
}
