package main

import (
	scolor "github.com/SnowsSky/scolor/pkg" // import the color file
)

func main() {
	// you can print with scolor.DisplayText(color, text); color is a string.
	scolor.DisplayText("blue", "Welcome ") // Spaces are not done automatically, so make sure you don't forget to.
	// you can also use capital letters for the color :
	scolor.DisplayText("RED", "to")               // or scolor.DisplayText("rEd", "to")....
	scolor.DisplayText("boldgreen", " scolors\n") // Same, '\n' are not done automatically.

	// Or if you want, you can print with scolor.White.DisplayText(text);
	scolor.White.DisplayText("A simple go color printing library.\n")
	// Don't use scolor.Reset.DisplayText(text); Because it will do nothing.

	// Since text is a ...string, you can also do
	scolor.Magenta.DisplayText("have ", "fun ") // Spaces are not done automatically, so make sure you don't forget to.
	// Or you can do.
	scolor.DisplayText("cyan", "with ", "it", ".") // Same.
}
