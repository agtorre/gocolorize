// Copyright (c) 2013 Aaron Torres. All rights reserved.
package main

import (
	"github.com/agtorre/gocolorize"
	"log"
	"os"
)

var (
	INFO     *log.Logger
	WARNING  *log.Logger
	CRITICAL *log.Logger
)

type MyError struct {
	What string
}

func main() {
	//Revel Example

	//first set some color information
	var info, warning gocolorize.Colorize
	info.SetColor(gocolorize.FgGreen)
	warning.SetColor(gocolorize.FgYellow)

	critical := gocolorize.Colorize{FgColor: gocolorize.FgBlack, BgColor: gocolorize.BgRed}

	//helper functions to shorten code
	i := info.PaintString
	w := warning.PaintString
	c := critical.PaintString

	//Define the look/feel of the INFO logger
	INFO = log.New(os.Stdout, i("INFO "), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	WARNING = log.New(os.Stdout, w("WARNING "), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	CRITICAL = log.New(os.Stdout, c("CRITICAL "), log.Ldate|log.Lmicroseconds|log.Lshortfile)

	//print out some messages, note the i wrappers for yellow text on the actual info string
	INFO.Println(i("Loaded module x"))
	INFO.Println(i("Loaded module y"))
	WARNING.Println(w("Failed to load module z"))

	e := MyError{What: "Failed"}

	CRITICAL.Println(c("Multiple arguments", "and complex types", e))

}
