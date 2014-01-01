package paints

import (
	"fmt"
	"github.com/agtorre/gocolorize"
	"testing"
)

func TestPaintColor(t *testing.T) {
	out_string := fmt.Sprint(PaintColor(gocolorize.FgGreen, "on the fly"))
	basis_string := "\033[0;32mon the fly\033[0m"
	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintBackground(t *testing.T) {
	out_string := fmt.Sprint(PaintBgColor(gocolorize.BgYellow, "Yellow Background"))
	basis_string := "\033[43mYellow Background\033[0m"

	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintColorAndBackground(t *testing.T) {
	out_string := fmt.Sprint(PaintColorAndBgColor(gocolorize.FgCyan, gocolorize.BgMagenta, "Cyan on Magenta??"))
	basis_string := "\033[0;36m\033[45mCyan on Magenta??\033[0m"

	if out_string != basis_string {
		t.Errorf("Error: string '%s' does not match '%s'\n", out_string, basis_string)
	} else {
		fmt.Printf("Success: string: '%s' matches '%s'\n", out_string, basis_string)
	}
}

func TestPaintColors(t *testing.T) {
	fmt.Println("Painting all colors:")
	fmt.Println(PaintBold("test1"))
	fmt.Println(PaintBlack("test2"))
	fmt.Println(PaintRed("test3"))
	fmt.Println(PaintGreen("test4"))
	fmt.Println(PaintYellow("test5"))
	fmt.Println(PaintBlue("test6"))
	fmt.Println(PaintMagenta("test7"))
	fmt.Println(PaintCyan("test8"))
	fmt.Println(PaintWhite("test9"))
	fmt.Println(PaintLightGray("test10"))
	fmt.Println(PaintGray("test11"))

	fmt.Println(PaintBgBlack("BgTest1"))
	fmt.Println(PaintBgRed("BgTest2"))
	fmt.Println(PaintBgGreen("BgTest3"))
	fmt.Println(PaintBgYellow("BgTest4"))
	fmt.Println(PaintBgBlue("BgTest5"))
	fmt.Println(PaintBgMagenta("BgTest6"))
	fmt.Println(PaintBgCyan("BgTest7"))
	fmt.Println(PaintBgWhite("BgTest8"))

	fmt.Println(PaintWhite(PaintBgBlack("Combined test??")))

}
