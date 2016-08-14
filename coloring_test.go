package coloring

import (
	"fmt"
	"testing"
)

func basefunc(f func(string) string) {
	fmt.Printf("Hello, %s!\n", f("@CreatorQsF"))
}

func TestBlack(t *testing.T) {
	basefunc(Black)
}

func TestRed(t *testing.T) {
	basefunc(Red)
}

func TestGreen(t *testing.T) {
	basefunc(Green)
}

func TestYellow(t *testing.T) {
	basefunc(Yellow)
}

func TestBlue(t *testing.T) {
	basefunc(Blue)
}

func TestMagenta(t *testing.T) {
	basefunc(Magenta)
}

func TestCyan(t *testing.T) {
	basefunc(Cyan)
}

func TestWhite(t *testing.T) {
	basefunc(White)
}
