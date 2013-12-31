package gocolorize

import (
	"fmt"
	"log"
	"testing"
)

var (
	DEBUG *log.Logger
	WARN  *log.Logger
)

func TestStatefulColor(t *testing.T) {
	var red, blue Colorize

	//set some state
	red.SetColor(COLOR_RED)
	blue.SetColor(COLOR_BLUE)

	out_string := fmt.Sprintf("This %s a %s", red.Fmt("is"), blue.Fmt("test"))
	basis_string := "This \033[0;31mis\033[0m a \033[0;34mtest\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestStatefulBGColor(t *testing.T) {
	var white_red_bg Colorize

	//set color and background
	white_red_bg.SetColor(COLOR_WHITE)
	white_red_bg.SetBgColor(COLOR_BG_RED)

	out_string := fmt.Sprintf(white_red_bg.Fmt("Now with backgrounds!"))
	basis_string := "\033[1;37m\033[41mNow with backgrounds!\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestStatelessColor(t *testing.T) {
	var generic Colorize

	out_string := fmt.Sprintf(generic.FmtColor("on the fly", COLOR_GREEN))
	basis_string := "\033[0;32mon the fly\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestStatelessBackground(t *testing.T) {
	var generic Colorize

	out_string := fmt.Sprintf(generic.FmtBgColor("Yellow Background", COLOR_BG_YELLOW))
	basis_string := "\033[43mYellow Background\033[0m"

	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestStatelessColorAndBackground(t *testing.T) {
	var generic Colorize

	out_string := fmt.Sprintf(generic.FmtColorAndBgColor("Cyan on Magenta??", COLOR_CYAN, COLOR_BG_MAGENTA))
	basis_string := "\033[0;36m\033[45mCyan on Magenta??\033[0m"

	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestNew(t *testing.T) {
	black_on_white := New(COLOR_BLACK, COLOR_BG_WHITE)
	f := black_on_white.Fmt

	out_string := fmt.Sprintf(f("Now this is cool"))
	basis_string := "\033[0;30m\033[47mNow this is cool\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}
