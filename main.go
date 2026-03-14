package main

import (
	"fmt"

	scolor "github.com/SnowsSky/scolor/pkg"
) // import the color file

func main() {
	// you can print with scolor.DisplayText(color, text); color is a string and text is a ...string.
	scolor.DisplayText("blue", "Welcome ") // Spaces are not done automatically, so make sure you don't forget to.
	scolor.DisplayText("red", "to")
	scolor.DisplayText("lightgreen", " scolors\n") // Same, '\n' are not done automatically

	// or if you want, you can print with scolor.White.DisplayText(text); text is a ...string.
	scolor.White.DisplayText("A simple go color printing library")

	// list all colors available
	fmt.Println(scolor.ColorList())
}
