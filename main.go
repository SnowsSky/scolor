package main

import (
	scolor "github.com/SnowsSky/scolor/pkg" // import the color file
)

func main() {
	// Old API:
	scolor.Cyan.DisplayText("Welcome ") // Does not create a new line
	scolor.Magenta.DisplayTextln("to ", "scolor!")

	// New API:
	scolor.DisplayText(scolor.BgRed|scolor.FgBoldBlue, "Start printing")
	scolor.DisplayTextln(scolor.FgYellow, " beautifully!") // Though none of them both are beautiful
}
