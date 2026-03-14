package main

import scolor "github.com/SnowsSky/scolor/pkg"

func main() {
	scolor.DisplayText("blue", "Welcome ")
	scolor.DisplayText("red", "to")
	scolor.DisplayText("lightgreen", " scolors\n")

	scolor.White.DisplayText("A simple color printing library")
}
