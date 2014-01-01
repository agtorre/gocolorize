// Copyright (c) 2013 Aaron Torres. All rights reserved.
package main

import (
	"fmt"
	"github.com/agtorre/gocolorize"
	"github.com/agtorre/gocolorize/paints"
)

func main() {
	// one way to define a stateful colorizer
	var green gocolorize.Colorize
	green.SetColor(gocolorize.FgGreen)
	g := green.Paint

	// Another way to do it
	red := gocolorize.Colorize{FgColor: gocolorize.FgRed}
	r := red.Paint

	//Some tips for using a generic stateless colorizer
	b := paints.PaintBlue

	//lets checkout out an advanced one
	//this could also be chained like this, but it's a lot less efficient:
	//paints.PaintBgBlack(paints.PaintFgBlack("What we want to print"))
	//in our case though, we'll use a stateful version
	c := gocolorize.Colorize{FgColor: gocolorize.FgWhite, BgColor: gocolorize.BgBlack}.Paint

	//r := gocolorize.COLOR_RED

	fmt.Println(b("On the twelfth day of Christmas"))
	fmt.Println(b("my true love sent to me:"))
	fmt.Println(g("Twelve"), c("Drummers"), r("Drumming"))
	fmt.Println(g("Eleven"), c("Pipers"), r("Piping"))
	fmt.Println(g("Ten"), c("Lords"), r("a Leaping"))
	fmt.Println(g("Nine"), c("Ladies"), r("Dancing"))
	fmt.Println(g("Eight"), c("Maids"), r("a Milking"))
	fmt.Println(g("Seven"), c("Swans"), r("a Swimming"))
	fmt.Println(g("Six"), c("Geese"), r("a Laying"))
	fmt.Println(g("Five"), c("Golden"), r("Rings"))
	fmt.Println(g("Four"), c("Calling"), r("Birds"))
	fmt.Println(g("Three"), c("French"), r("Hens"))
	fmt.Println(g("Two"), c("Turtle"), r("Doves"))
	fmt.Println(b("and a Partridge in a Pear Tree"))

}
