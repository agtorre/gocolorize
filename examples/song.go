// Copyright (c) 2013 Aaron Torres. All rights reserved.
package main

import (
	"fmt"
	"gocolorize"
)

func main() {
	// one way to define a stateful colorizer
	var green gocolorize.Colorize
	green.SetColor(gocolorize.COLOR_GREEN)
    g := green.Fmt

	// Another way to do it
	red := gocolorize.New(gocolorize.COLOR_RED, gocolorize.COLOR_BG_NONE)
    r := red.Fmt

	//Some tips for using a generic stateless colorizer
	var generic gocolorize.Colorize
	f := generic.FmtColor
	y := gocolorize.COLOR_YELLOW

	//r := gocolorize.COLOR_RED

	fmt.Println(f("On the twelfth day of Christmas", y))
	fmt.Println(f("my true love sent to me:", y))
	fmt.Println(g("Twelve"), "Drummers", r("Drumming"))
	fmt.Println(g("Eleven"), "Pipers", r("Piping"))
	fmt.Println(g("Ten"), "Lords a", r("Leaping"))
	fmt.Println(g("Nine"), "Ladies", r("Dancing"))
	fmt.Println(g("Eight"), "Maids a", r("Milking"))
	fmt.Println(g("Seven"), "Swans a", r("Swimming"))
	fmt.Println(g("Six"), "Geese a", r("Laying"))
	fmt.Println(g("Five"), "Golden", r("Rings"))
	fmt.Println(g("Four"), "Calling", r("Birds"))
	fmt.Println(g("Three"), "French", r("Hens"))
	fmt.Println(g("Two"), "Turtle", r("Doves"))
	fmt.Println(f("and a Partridge in a Pear Tree", y))

}
