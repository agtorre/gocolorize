package gocolorize

import (
	"bytes"
	"fmt"
)

//A list of the basic unix shell colors
const reset = "\033[0m"

type FgColor string

const (
	FgNone        FgColor = ""
	FgBold        FgColor = "\033[1m"
	FgBlack       FgColor = "\033[0;30m"
	FgRed         FgColor = "\033[0;31m"
	FgGreen       FgColor = "\033[0;32m"
	FgYellow      FgColor = "\033[0;33m"
	FgBlue        FgColor = "\033[0;34m"
	FgMagenta     FgColor = "\033[0;35m"
	FgCyan        FgColor = "\033[0;36m"
	FgLightGray   FgColor = "\033[0;37m"
	FgGray        FgColor = "\033[1;30m"
	FgBoldRed     FgColor = "\033[1;31m"
	FgBoldGreen   FgColor = "\033[1;32m"
	FgBoldYellow  FgColor = "\033[1;33m"
	FgBoldBlue    FgColor = "\033[1;34m"
	FgBoldMagenta FgColor = "\033[1;35m"
	FgBoldCyan    FgColor = "\033[1;36m"
	FgWhite       FgColor = "\033[1;37m"
)

type BgColor string

const (
	BgNone    BgColor = ""
	BgBlack   BgColor = "\033[40m"
	BgRed     BgColor = "\033[41m"
	BgGreen   BgColor = "\033[42m"
	BgYellow  BgColor = "\033[43m"
	BgBlue    BgColor = "\033[44m"
	BgMagenta BgColor = "\033[45m"
	BgCyan    BgColor = "\033[46m"
	BgWhite   BgColor = "\033[47m"
)

type Colorize struct {
	Value []interface{}
	FgColor
	BgColor
}

func (c Colorize) Paint(v ...interface{}) Colorize {
	return Colorize{Value: v, FgColor: c.FgColor, BgColor: c.BgColor}
}

func (c Colorize) PaintString(v ...interface{}) string {
	r := Colorize{Value: v, FgColor: c.FgColor, BgColor: c.BgColor}
	return fmt.Sprint(r)
}

// Format allows ColorText to satisfy the fmt.Formatter interface. The format
// behaviour is the same as for fmt.Print.
func (ct Colorize) Format(fs fmt.State, c rune) {
	fmt.Fprint(fs, ct.FgColor)
	fmt.Fprint(fs, ct.BgColor)
	w, wOk := fs.Width()
	p, pOk := fs.Precision()
	var b bytes.Buffer
	for i, _ := range ct.Value {
		b.WriteByte('%')
		for _, f := range "+-# 0" {
			if fs.Flag(int(f)) {
				b.WriteRune(f)
			}
		}
		if wOk {
			fmt.Fprint(&b, w)
		}
		if pOk {
			b.WriteByte('.')
			fmt.Fprint(&b, p)
		}
		b.WriteRune(c)
		if i < len(ct.Value)-1 {
			b.WriteByte(' ')
		}
	}
	fmt.Fprintf(fs, b.String(), ct.Value...)
	fmt.Fprint(fs, reset)
}

func (C *Colorize) SetColor(c FgColor) {
	C.FgColor = c
}

func (C *Colorize) SetBgColor(b BgColor) {
	C.BgColor = b
}
