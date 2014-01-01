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
	red.SetColor(FgRed)
	blue.SetColor(FgBlue)

	out_string := fmt.Sprintf("This %s a %s", red.Paint("is"), blue.Paint("test"))
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
	white_red_bg.SetColor(FgWhite)
	white_red_bg.SetBgColor(BgRed)

	out_string := fmt.Sprint(white_red_bg.Paint("Now with backgrounds!"))
	basis_string := "\033[1;37m\033[41mNow with backgrounds!\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestStatefulMultipleInterface(t *testing.T) {
	blue := Colorize{FgColor: FgBlue}
	out_string := fmt.Sprint(blue.Paint("Multiple types of args:", 1, 1.24))
	basis_string := "\033[0;34mMultiple types of args: 1 1.24\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestStatefulComplexType(t *testing.T) {
	green := Colorize{BgColor: BgGreen}
	out_string := fmt.Sprint(green.Paint(
		struct {
			int
			string
		}{}))
	basis_string := fmt.Sprintf("\033[42m%v\033[0m", struct {
		int
		string
	}{})
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestInitialize(t *testing.T) {
	black_on_white := Colorize{FgColor: FgBlack, BgColor: BgWhite}
	f := black_on_white.Paint

	out_string := fmt.Sprint(f("Now this is cool"))
	basis_string := "\033[0;30m\033[47mNow this is cool\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}
