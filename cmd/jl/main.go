package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/miihael/jl"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Usage = func() {
		fmt.Printf(`Usage of %s:

    %s [filename]

If [filename] is omitted, it reads from standard input.

`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	formatFlag := flag.String("format", "compact", `Formatter for logs. The options are "compact","logfmt","ltsv"`)
	flag.StringVar(formatFlag, "f", "compact", "alias for -format")
	color := flag.String("color", "auto", `Sets the color mode. The options are "auto", "yes", and "no". "auto" disables color if stdout is not a tty`)
	flag.StringVar(color, "c", "auto", "alias for -color")
	truncate := flag.Bool("truncate", true, "Whether to truncate strings in the compact formatter")
	flag.Parse()

	disableColor := false
	switch *color {
	case "auto":
		if !isatty.IsTerminal(os.Stdout.Fd()) {
			disableColor = true
		}
	case "yes":
		disableColor = false
	case "no":
		disableColor = true
	default:
		return fmt.Errorf("invalid -color=%s", *color)
	}

	var printer jl.EntryPrinter
	switch *formatFlag {
	case "logfmt":
		lp := jl.NewLogfmtPrinter(os.Stdout)
		lp.DisableColor = disableColor
		printer = lp
	case "compact":
		cp := jl.NewCompactPrinter(os.Stdout)
		cp.DisableColor = disableColor
		cp.DisableTruncate = !*truncate
		printer = cp
	case "ltsv":
		lp := jl.NewLTSVPrinter(os.Stdout)
		lp.DisableColor = disableColor
		printer = lp
	}

	fileArg := flag.Arg(0)
	inFile := os.Stdin
	if fileArg != "" {
		f, err := os.Open(fileArg)
		if err != nil {
			return err
		}
		defer f.Close()
		inFile = f
	}
	return jl.NewParser(inFile, printer).Consume()
}
