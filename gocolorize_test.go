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

func TestPaint(t *testing.T) {
	var blue Colorize

	//set some state
	blue.SetColor(FgBlue)

	out_string := fmt.Sprintf("Testing %s", blue.Paint("paint"))
	basis_string := "Testing \033[0;34mpaint\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintString(t *testing.T){
	var red Colorize

	//set some state
	red.SetColor(FgRed)

	out_string := red.PaintString("Returning a string")
	basis_string := "\033[0;31mReturning a string\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
    }
}

func TestSetColorSetBgColor(t *testing.T) {
	var white_red_bg Colorize

	//set color and background
	white_red_bg.SetColor(FgWhite)
	white_red_bg.SetBgColor(BgRed)

	out_string := fmt.Sprint(white_red_bg.Paint("Setting a foreground and background color!"))
	basis_string := "\033[1;37m\033[41mSetting a foreground and background color!\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintMultipleInterface(t *testing.T) {
	blue := Colorize{FgColor: FgBlue}
	out_string := fmt.Sprint(blue.Paint("Multiple types of args:", 1, 1.24))
	basis_string := "\033[0;34mMultiple types of args: 1 1.24\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintComplexType(t *testing.T) {
	green := Colorize{BgColor: BgGreen}
    out_string := fmt.Sprint(green.Paint("Complex types:",
		struct {
			int
			string
		}{}))
        basis_string := fmt.Sprintf("\033[42mComplex types: %v\033[0m", struct {
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

func TestWidthPrecisionAndFlag(t *testing.T){
    f := Colorize{FgColor: FgYellow}
    out_string := fmt.Sprintf("%+v %1.2v", f.Paint("1234"), f.Paint("1234"))
	basis_string := "\033[0;33m1234\033[0m \033[0;33m12\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}

}
