package logify

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

// Logify is a simple flag parser with common options
type Logify struct {
	flagSet      *flag.FlagSet
	description  string
	groups       []flagGroup
	flagData     map[string]*FlagData
	currentGroup string
}

type flagGroup struct {
	name        string
	description string
}

// FlagData holds metadata about a flag
type FlagData struct {
	short        string
	long         string
	usage        string
	group        string
	defaultValue interface{}
}

// New creates a new Logify instance
func New(description string) *Logify {
	return &Logify{
		flagSet:     flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		description: description,
		groups:      []flagGroup{},
		flagData:    make(map[string]*FlagData),
	}
}

// CreateGroup creates a group and executes the setup function
func (l *Logify) CreateGroup(name, description string, setup func()) {
	l.groups = append(l.groups, flagGroup{
		name:        name,
		description: description,
	})
	l.currentGroup = name
	if setup != nil {
		setup()
	}
	l.currentGroup = ""
}

// StringVar adds a string flag with a longname
func (l *Logify) StringVar(field *string, long, defaultValue, usage string) *FlagData {
	return l.StringVarP(field, long, "", defaultValue, usage)
}

// StringVarP adds a string flag with a shortname and longname
func (l *Logify) StringVarP(field *string, long, short, defaultValue, usage string) *FlagData {
	*field = defaultValue

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.StringVar(field, short, defaultValue, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.StringVar(field, long, defaultValue, usage)
	l.flagData[long] = flagData

	return flagData
}

// String adds a string flag (legacy method for compatibility)
func (l *Logify) String(long, short, defaultValue, usage string) *string {
	value := new(string)
	l.StringVarP(value, long, short, defaultValue, usage)
	return value
}

// BoolVar adds a boolean flag with a longname
func (l *Logify) BoolVar(field *bool, long string, defaultValue bool, usage string) *FlagData {
	return l.BoolVarP(field, long, "", defaultValue, usage)
}

// BoolVarP adds a boolean flag with a shortname and longname
func (l *Logify) BoolVarP(field *bool, long, short string, defaultValue bool, usage string) *FlagData {
	*field = defaultValue

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.BoolVar(field, short, defaultValue, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.BoolVar(field, long, defaultValue, usage)
	l.flagData[long] = flagData

	return flagData
}

// Bool adds a boolean flag (legacy method for compatibility)
func (l *Logify) Bool(long, short string, defaultValue bool, usage string) *bool {
	value := new(bool)
	l.BoolVarP(value, long, short, defaultValue, usage)
	return value
}

// IntVar adds an integer flag with a longname
func (l *Logify) IntVar(field *int, long string, defaultValue int, usage string) *FlagData {
	return l.IntVarP(field, long, "", defaultValue, usage)
}

// IntVarP adds an integer flag with a shortname and longname
func (l *Logify) IntVarP(field *int, long, short string, defaultValue int, usage string) *FlagData {
	*field = defaultValue

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.IntVar(field, short, defaultValue, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.IntVar(field, long, defaultValue, usage)
	l.flagData[long] = flagData

	return flagData
}

// Int adds an integer flag (legacy method for compatibility)
func (l *Logify) Int(long, short string, defaultValue int, usage string) *int {
	value := new(int)
	l.IntVarP(value, long, short, defaultValue, usage)
	return value
}

// DurationVar adds a duration flag with a longname
func (l *Logify) DurationVar(field *time.Duration, long string, defaultValue time.Duration, usage string) *FlagData {
	return l.DurationVarP(field, long, "", defaultValue, usage)
}

// DurationVarP adds a duration flag with a shortname and longname
func (l *Logify) DurationVarP(field *time.Duration, long, short string, defaultValue time.Duration, usage string) *FlagData {
	*field = defaultValue

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.DurationVar(field, short, defaultValue, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.DurationVar(field, long, defaultValue, usage)
	l.flagData[long] = flagData

	return flagData
}

// Duration adds a duration flag (legacy method for compatibility)
func (l *Logify) Duration(long, short string, defaultValue time.Duration, usage string) *time.Duration {
	value := new(time.Duration)
	l.DurationVarP(value, long, short, defaultValue, usage)
	return value
}

// StringSliceVar adds a string slice flag with a longname
func (l *Logify) StringSliceVar(field *StringSlice, long string, defaultValue []string, usage string) *FlagData {
	return l.StringSliceVarP(field, long, "", defaultValue, usage)
}

// StringSliceVarP adds a string slice flag with a shortname and longname
func (l *Logify) StringSliceVarP(field *StringSlice, long, short string, defaultValue []string, usage string) *FlagData {
	// Set default values
	for _, val := range defaultValue {
		field.Set(val)
	}

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.Var(field, short, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.Var(field, long, usage)
	l.flagData[long] = flagData

	return flagData
}

// FileStringSliceVar adds a string slice flag with file support and a longname
func (l *Logify) FileStringSliceVar(field *StringSlice, long string, defaultValue []string, usage string) *FlagData {
	return l.FileStringSliceVarP(field, long, "", defaultValue, usage)
}

// FileStringSliceVarP adds a string slice flag with file support, shortname and longname
func (l *Logify) FileStringSliceVarP(field *StringSlice, long, short string, defaultValue []string, usage string) *FlagData {
	// Enable file support
	field.supportFile = true

	// Set default values
	for _, val := range defaultValue {
		field.Set(val)
	}

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.Var(field, short, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.Var(field, long, usage)
	l.flagData[long] = flagData

	return flagData
}

// StringSlice adds a string slice flag (basic, no file support) - legacy method
func (l *Logify) StringSlice(long, short string, usage string) *StringSlice {
	value := NewStringSlice(false)
	l.StringSliceVarP(value, long, short, []string{}, usage)
	return value
}

// FileStringSlice adds a string slice flag with file support - legacy method
func (l *Logify) FileStringSlice(long, short string, usage string) *StringSlice {
	value := NewStringSlice(true)
	l.FileStringSliceVarP(value, long, short, []string{}, usage)
	return value
}

// SizeVar adds a size flag with a longname
func (l *Logify) SizeVar(field *Size, long string, defaultValue string, usage string) *FlagData {
	return l.SizeVarP(field, long, "", defaultValue, usage)
}

// SizeVarP adds a size flag with a shortname and longname
func (l *Logify) SizeVarP(field *Size, long, short string, defaultValue string, usage string) *FlagData {
	if defaultValue != "" {
		_ = field.Set(defaultValue)
	}

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: defaultValue,
	}

	if short != "" {
		l.flagSet.Var(field, short, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.Var(field, long, usage)
	l.flagData[long] = flagData

	return flagData
}

// Size adds a size flag (supports kb, mb, gb, tb) - legacy method
func (l *Logify) Size(long, short string, defaultValue string, usage string) *Size {
	value := new(Size)
	l.SizeVarP(value, long, short, defaultValue, usage)
	return value
}

// CallbackVar adds a callback flag with a longname that executes fn when set
func (l *Logify) CallbackVar(fn func(), long string, usage string) *FlagData {
	return l.CallbackVarP(fn, long, "", usage)
}

// CallbackVarP adds a callback flag with a shortname and longname that executes fn when set
func (l *Logify) CallbackVarP(fn func(), long, short string, usage string) *FlagData {
	callback := &Callback{fn: fn}

	flagData := &FlagData{
		short:        short,
		long:         long,
		usage:        usage,
		group:        l.currentGroup,
		defaultValue: false,
	}

	if short != "" {
		l.flagSet.Var(callback, short, usage)
		l.flagData[short] = flagData
	}
	l.flagSet.Var(callback, long, usage)
	l.flagData[long] = flagData

	return flagData
}

// Callback adds a callback flag that executes fn when the flag is set - legacy method
func (l *Logify) Callback(long, short string, usage string, fn func()) *FlagData {
	return l.CallbackVarP(fn, long, short, usage)
}

// Parse parses the command line arguments
func (l *Logify) Parse() error {
	l.flagSet.Usage = l.printUsage
	return l.flagSet.Parse(os.Args[1:])
}

// ParseArgs parses specific arguments (useful for testing)
func (l *Logify) ParseArgs(args []string) error {
	l.flagSet.Usage = l.printUsage
	return l.flagSet.Parse(args)
}

// printUsage prints the usage information with groups
func (l *Logify) printUsage() {
	output := os.Stderr

	if l.description != "" {
		fmt.Fprintf(output, "%s\n\n", l.description)
	}

	fmt.Fprintf(output, "Usage: %s [options]\n\n", os.Args[0])

	if len(l.groups) > 0 {
		// Print grouped flags
		for _, group := range l.groups {
			fmt.Fprintf(output, "%s:\n", strings.ToUpper(group.description))

			writer := tabwriter.NewWriter(output, 0, 0, 3, ' ', 0)

			// Track seen flags to avoid duplicates
			seen := make(map[string]bool)

			for _, data := range l.flagData {
				if data.group == group.name && !seen[data.long] {
					seen[data.long] = true
					l.printFlag(writer, data)
				}
			}

			writer.Flush()
			fmt.Fprintln(output)
		}

		// Print ungrouped flags
		writer := tabwriter.NewWriter(output, 0, 0, 3, ' ', 0)
		hasUngrouped := false
		seen := make(map[string]bool)

		for _, data := range l.flagData {
			if data.group == "" && !seen[data.long] {
				if !hasUngrouped {
					fmt.Fprintf(output, "OTHER OPTIONS:\n")
					hasUngrouped = true
				}
				seen[data.long] = true
				l.printFlag(writer, data)
			}
		}

		if hasUngrouped {
			writer.Flush()
			fmt.Fprintln(output)
		}
	} else {
		// Print all flags without grouping
		fmt.Fprintf(output, "Options:\n")
		writer := tabwriter.NewWriter(output, 0, 0, 3, ' ', 0)
		seen := make(map[string]bool)

		for _, data := range l.flagData {
			if !seen[data.long] {
				seen[data.long] = true
				l.printFlag(writer, data)
			}
		}

		writer.Flush()
		fmt.Fprintln(output)
	}
}

// printFlag prints a single flag
func (l *Logify) printFlag(w *tabwriter.Writer, data *FlagData) {
	var names []string
	
	if data.short != "" {
		names = append(names, "-"+data.short)
	}
	if data.long != "" {
		names = append(names, "-"+data.long)
	}
	
	flagNames := "  " + strings.Join(names, ", ")

	// Get the flag to check its type
	fl := l.flagSet.Lookup(data.long)
	if fl != nil {
		flagType := ""

		// Determine type based on default value
		switch data.defaultValue.(type) {
		case string:
			if data.defaultValue != "" {
				flagType = " string"
			}
		case int:
			flagType = " int"
		case time.Duration:
			flagType = " duration"
		case bool:
			flagType = ""
		default:
			flagType = " value"
		}

		fmt.Fprintf(w, "%s%s\t%s", flagNames, flagType, data.usage)

		// Show default value if not empty/zero
		if !isZeroValue(data.defaultValue) {
			fmt.Fprintf(w, " (default: %v)", data.defaultValue)
		}

		fmt.Fprintln(w)
	}
}

// Parsed returns true if Parse() has been called
func (l *Logify) Parsed() bool {
	return l.flagSet.Parsed()
}

// Args returns the non-flag arguments
func (l *Logify) Args() []string {
	return l.flagSet.Args()
}

// NArg returns the number of non-flag arguments
func (l *Logify) NArg() int {
	return l.flagSet.NArg()
}
