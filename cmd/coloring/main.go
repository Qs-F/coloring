// Copyright 2016 de-liKeR CreatorQsF under MIT License.
// This package provide general command of `github.com/Qs-F/coloring`.
// Colorize stdin string and output colorized msg.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/Qs-F/coloring"
	_ "golang.org/x/crypto/ssh/terminal"
)

const (
	unknownColor_ = coloring.R
)

type output struct {
	Color   int
	Message string
}

func flagManage() *output {
	c := output{
		Color:   unknownColor_,
		Message: "",
	}
	// must parses:
	// coloring black long message
	// hoge | coloring black
	// coloring -bk long message
	// hoge | coloring -bk
	// coloring long message
	// hoge | coloring
	// var color string
	var colorN int
	var err error
	msgP := 2
	if len(os.Args) >= 2 {
		colorN, err = colorChooser(os.Args[1])
		msgP = 2 // coloring hoge FROM HERE
		if err != nil {
			colorN = unknownColor_
			msgP = 1
		}
	} else {
		colorN = unknownColor_
	}
	// }
	message := ""
	if len(os.Args) >= msgP+1 {
		for _, v := range os.Args[msgP:] {
			message += v
		}
	}
	c.Color = colorN
	if message != "" {
		message += "\n"
	}
	c.Message = message
	return &c
}

func colorChooser(color string) (int, error) {
	switch color {
	case "black", "-bk":
		return coloring.Bk, nil
	case "red", "-r":
		return coloring.R, nil
	case "green", "-g":
		return coloring.G, nil
	case "yellow", "-y":
		return coloring.Y, nil
	case "blue", "-bl":
		return coloring.Bl, nil
	case "magenta", "-m":
		return coloring.M, nil
	case "cyan", "-c":
		return coloring.C, nil
	case "white", "-w":
		return coloring.W, nil
	default:
		return 0, errors.New("no color")
	}
	// ERR_CODE_001
	return 0, errors.New("unexpected error(err_code: 001)")
}

func main() {
	m := flagManage()
	if m.Message == "" {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			m.Message += s.Text() + "\n"
		}
	}
	fmt.Fprintf(os.Stdout, coloring.Base(m.Color, m.Message))
}
