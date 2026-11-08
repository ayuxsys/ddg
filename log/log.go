package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	Info    = log.New(os.Stderr, color.CyanString("[info] "), 0)
	Error   = log.New(os.Stderr, color.RedString("[error] "), 0)
	Success = log.New(os.Stderr, color.GreenString("[ok] "), 0)
	Debug   = log.New(os.Stderr, color.BlueString("[debug] "), 0)
	Warn    = log.New(os.Stderr, color.YellowString("[warn] "), 0)
)

func Setup(verbose bool) {
	if verbose {
		Info.SetFlags(log.Ltime | log.Lshortfile)
		Error.SetFlags(log.Ltime | log.Lshortfile)
		Success.SetFlags(log.Ltime | log.Lshortfile)
		Warn.SetFlags(log.Ltime | log.Lshortfile)
	}
}
