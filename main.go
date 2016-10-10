package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Qs-F/cliset"
	"github.com/fatih/color"
)

func main() {
	var c color.Attribute
	var bg color.Attribute
	fc := flag.String("c", "", "Character color")
	fbg := flag.String("bg", "", "Character background color")
	msg := flag.String("m", "", "Body message")
	flag.Parse()
	// Char color
	switch *fc {
	//Low color
	case "black", "bk":
		c = color.FgBlack
	case "red", "r":
		c = color.FgRed
	case "green", "g":
		c = color.FgGreen
	case "yellow", "y":
		c = color.FgYellow
	case "blue", "bl":
		c = color.FgBlue
	case "magenta", "m":
		c = color.FgMagenta
	case "cyan", "c":
		c = color.FgCyan
	case "white", "w":
		c = color.FgWhite
	// Hi color
	case "hi-black", "hbk":
		c = color.FgHiBlack
	case "hi-red", "hr":
		c = color.FgHiRed
	case "hi-green", "hg":
		c = color.FgHiGreen
	case "hi-yellow", "hy":
		c = color.FgHiYellow
	case "hi-blue", "hbl":
		c = color.FgHiBlue
	case "hi-magenta", "hm":
		c = color.FgHiMagenta
	case "hi-cyan", "hc":
		c = color.FgHiCyan
	case "hi-white", "hw":
		c = color.FgHiWhite
	default:
	}

	//Background color
	switch *fbg {
	// Low color
	case "black", "bk":
		bg = color.BgBlack
	case "red", "r":
		bg = color.BgRed
	case "green", "g":
		bg = color.BgGreen
	case "yellow", "y":
		bg = color.BgYellow
	case "blue", "bl":
		bg = color.BgBlue
	case "magenta", "m":
		bg = color.BgMagenta
	case "cyan", "c":
		bg = color.BgCyan
	case "white", "w":
		bg = color.BgWhite
	// Hi color
	case "hi-black", "hbk":
		bg = color.BgHiBlack
	case "hi-red", "hr":
		bg = color.BgHiRed
	case "hi-green", "hg":
		bg = color.BgHiGreen
	case "hi-yellow", "hy":
		bg = color.BgHiYellow
	case "hi-blue", "hbl":
		bg = color.BgHiBlue
	case "hi-magenta", "hm":
		bg = color.BgHiMagenta
	case "hi-cyan", "hc":
		bg = color.BgHiCyan
	case "hi-white", "hw":
		bg = color.BgHiWhite
	default:
	}

	defer color.Unset()
	if *msg == "" {
		v, err := cliset.PipeValue()
		if err != nil {
			*msg = ""
		}
		*msg = v
	}
	if c == 0 && bg == 0 {
		Output(*msg)
	} else if c != 0 && bg == 0 {
		color.Set(c)
		Output(*msg)
	} else if c == 0 && bg != 0 {
		color.Set(bg)
		Output(*msg)
	} else if c != 0 && bg != 0 {
		color.Set(c, bg)
		Output(*msg)
	}
	return
}

func Output(msg string) {
	fmt.Fprintf(os.Stdout, msg)
}
