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

func base(colorN int, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorN, s)
}

func Black(s string) string {
	return base(Bk, s)
}

func Red(s string) string {
	return base(R, s)
}

func Green(s string) string {
	return base(G, s)
}

func Yellow(s string) string {
	return base(Y, s)
}

func Blue(s string) string {
	return base(Bl, s)
}

func Magenta(s string) string {
	return base(M, s)
}

func Cyan(s string) string {
	return base(C, s)
}

func White(s string) string {
	return base(W, s)
}
