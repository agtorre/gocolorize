package paints

import "github.com/agtorre/gocolorize"

func PaintColor(c gocolorize.FgColor, val ...interface{}) gocolorize.Colorize {
	color := gocolorize.Colorize{FgColor: c, BgColor: gocolorize.BgNone}
	return color.Paint(val...)
}

func PaintBgColor(b gocolorize.BgColor, val ...interface{}) gocolorize.Colorize {
	color := gocolorize.Colorize{FgColor: gocolorize.FgNone, BgColor: b}
	return color.Paint(val...)
}

func PaintColorAndBgColor(c gocolorize.FgColor, b gocolorize.BgColor, val ...interface{}) gocolorize.Colorize {
	color := gocolorize.Colorize{FgColor: c, BgColor: b}
	return color.Paint(val...)
}

func PaintBold(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgBold, val...)
}

func PaintBlack(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgBlack, val...)
}

func PaintRed(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgRed, val...)
}

func PaintGreen(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgGreen, val...)
}

func PaintYellow(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgYellow, val...)
}

func PaintBlue(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgBlue, val...)
}

func PaintMagenta(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgMagenta, val...)
}

func PaintCyan(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgCyan, val...)
}

func PaintWhite(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgWhite, val...)
}

func PaintLightGray(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgLightGray, val...)
}

func PaintGray(val ...interface{}) gocolorize.Colorize {
	return PaintColor(gocolorize.FgGray, val...)
}

func PaintBgBlack(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgBlack, val...)
}

func PaintBgRed(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgRed, val...)
}

func PaintBgGreen(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgGreen, val...)
}

func PaintBgYellow(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgYellow, val...)
}

func PaintBgBlue(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgBlue, val...)
}

func PaintBgMagenta(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgMagenta, val...)
}

func PaintBgCyan(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgCyan, val...)
}

func PaintBgWhite(val ...interface{}) gocolorize.Colorize {
	return PaintBgColor(gocolorize.BgWhite, val...)
}
