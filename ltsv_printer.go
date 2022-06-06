package jl

import (
	"io"
)

// DefaultLTSVPreferredFields is the set of fields that NewLTSVPrinter orders ahead of other fields.
var DefaultLTSVPreferredFields = []string{
	"timestamp",
	"time",
	"ts",
	"t",
	"lvl",
	"level",
	"log_level",
	"thread",
	"logger",
	"message",
	"msg",
	"exceptions",
}

// NewLTSVPrinter allocates and returns specialized LogFmtPrinter.
func NewLTSVPrinter(w io.Writer) *LogfmtPrinter {
	return &LogfmtPrinter{
		Out:             w,
		PreferredFields: DefaultLTSVPreferredFields,
		KVDelimiter:     ":",
		RecordDelimiter: "\t",
	}
}
