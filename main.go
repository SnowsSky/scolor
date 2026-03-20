package main

import (
	scolor "github.com/SnowsSky/scolor/v2/pkg" // import the color file
)

func main() {
	// If you don't want any background :
	scolor.Cyan.DisplayText("Welcome ") // Does not create a new line
	scolor.Magenta.DisplayTextln("to ", "scolor!")

	// If you want to have a background
	scolor.DisplayText(scolor.BgRed|scolor.FgBoldBlue, "Start printing")
	// You can also choose to don't use a bg
	scolor.DisplayTextln(scolor.FgYellow, " beautifully!") // Though none of them both are beautiful

	scolor.DisplayTextln(scolor.FgBoldWhite, "scolor@"+scolor.Version) // scolor.Version returns the installed scolor version.
}
