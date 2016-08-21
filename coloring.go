package coloring

import (
	"fmt"
)

const (
	Bk = 30
	R  = 31
	G  = 32
	Y  = 33
	Bl = 34
	M  = 35
	C  = 36
	W  = 37
)

func Base(colorN int, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorN, s)
}

func Black(s string) string {
	return Base(Bk, s)
}

func Red(s string) string {
	return Base(R, s)
}

func Green(s string) string {
	return Base(G, s)
}

func Yellow(s string) string {
	return Base(Y, s)
}

func Blue(s string) string {
	return Base(Bl, s)
}

func Magenta(s string) string {
	return Base(M, s)
}

func Cyan(s string) string {
	return Base(C, s)
}

func White(s string) string {
	return Base(W, s)
}
