package colorize

import "strings"

//A list of the basic unix shell colors
type Color string

const (
	COLOR_NONE    Color = ""
	COLOR_RESET   Color = "\033[0m"
	COLOR_BOLD    Color = "\033[1m"
	COLOR_BLACK   Color = "\033[0;30m"
	COLOR_RED     Color = "\033[0;31m"
	COLOR_GREEN   Color = "\033[0;32m"
	COLOR_YELLOW  Color = "\033[0;33m"
	COLOR_BLUE    Color = "\033[0;34m"
	COLOR_MAGENTA Color = "\033[0;35m"
	COLOR_CYAN    Color = "\033[0;36m"

	COLOR_LIGHT_GRAY Color = "\033[0;37m"
	COLOR_GRAY       Color = "\033[1;30m"

	COLOR_BOLD_RED     Color = "\033[1;31m"
	COLOR_BOLD_GREEN   Color = "\033[1;32m"
	COLOR_BOLD_YELLOW  Color = "\033[1;33m"
	COLOR_BOLD_BLUE    Color = "\033[1;34m"
	COLOR_BOLD_MAGENTA Color = "\033[1;35m"
	COLOR_BOLD_CYAN    Color = "\033[1;36m"
	COLOR_WHITE        Color = "\033[1;37m"
)

type BgColor string

const (
	COLOR_BG_NONE    BgColor = ""
	COLOR_BG_RED     BgColor = "\033[41m"
	COLOR_BG_GREEN   BgColor = "\033[42m"
	COLOR_BG_YELLOW  BgColor = "\033[43m"
	COLOR_BG_BLUE    BgColor = "\033[44m"
	COLOR_BG_MAGENTA BgColor = "\033[45m"
	COLOR_BG_CYAN    BgColor = "\033[46m"
	COLOR_BG_WHITE   BgColor = "\033[47m"
)

type Colorize struct {
	Color
	BgColor
}

func New(c Color, b BgColor) Colorize {
	var result Colorize
	result.SetColor(c)
	result.SetBgColor(b)
	return result
}

func (C *Colorize) SetColor(c Color) {
	C.Color = c
}

func (C *Colorize) SetBgColor(b BgColor) {
	C.BgColor = b
}

func (C *Colorize) Fmt(val string) string {
	var s []string
	s = append(s, string(C.Color))
	if C.BgColor != COLOR_BG_NONE {
		s = append(s, string(C.BgColor))
	}
	s = append(s, val)
	s = append(s, string(COLOR_RESET))

	result := strings.Join(s, "")
	return result
}

func (C *Colorize) FmtColor(val string, c Color) string {
	var orig_color = C.Color
	C.Color = c
	result := C.Fmt(val)
	C.Color = orig_color
	return result
}

func (C *Colorize) FmtBgColor(val string, b BgColor) string {
	var orig_bg_color = C.BgColor
	C.BgColor = b
	result := C.Fmt(val)
	C.BgColor = orig_bg_color
	return result
}

func (C *Colorize) FmtColorAndBgColor(val string, c Color, b BgColor) string {
	var orig_color = C.Color
	var orig_bg_color = C.BgColor
	C.Color = c
	C.BgColor = b
	result := C.Fmt(val)
	C.Color = orig_color
	C.BgColor = orig_bg_color
	return result
}
