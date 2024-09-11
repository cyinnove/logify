package logify

type Level int

const (
	Null Level = iota
	Fatal
	Silent
	Label
	Misc
	Error
	Info
	Warning
	Debug
	Verbose
)

var labels = map[Level]string{
	Warning: "WRN",
	Error:   "ERR",
	Label:   "WRN",
	Fatal:   "FTL",
	Debug:   "DBG",
	Info:    "INF",
}
