// Copyright (c) 2013 Aaron Torres. All rights reserved.
package main

import (
    "gocolorize"
    "log"
    "os"
)


var (
    INFO *log.Logger
    WARNING *log.Logger
)


func main(){
    //Revel Example

    //first set some color information
    var info, warning gocolorize.Colorize
    info.SetColor(gocolorize.COLOR_YELLOW)
    warning.SetColor(gocolorize.COLOR_RED)

    //helper functions to shorten code
    i := info.Fmt
    w := warning.Fmt

    //Define the look/feel of the INFO logger
    INFO = log.New(os.Stdout, i("INFO "), log.Ldate|log.Lmicroseconds|log.Lshortfile)
    WARNING = log.New(os.Stdout, w("WARNING "), log.Ldate|log.Lmicroseconds|log.Lshortfile)

    //print out some messages, note the i wrappers for yellow text on the actual info string
    INFO.Println(i("Loaded module x"))
    INFO.Println(i("Loaded module y"))
    WARNING.Println(w("Failed to load module z"))
}

