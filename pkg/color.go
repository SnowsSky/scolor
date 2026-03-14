package pkg

import (
	"fmt"
	"strings"
)

type color struct {
	Color string
}

var Colors = map[string]string{
	"red":          "\033[31m",
	"green":        "\033[32m",
	"yellow":       "\033[33m",
	"blue":         "\033[34m",
	"magenta":      "\033[35m",
	"cyan":         "\033[36m",
	"gray":         "\033[37m",
	"white":        "\033[97m",
	"lightred":     "\033[1;31m",
	"lightgreen":   "\033[1;32m",
	"lightyellow":  "\033[1;33m",
	"lightblue":    "\033[1;34m",
	"lightmagenta": "\033[1;35m",
	"lightcyan":    "\033[1;36m",
	"lightgray":    "\033[1;37m",
	"lightwhite":   "\033[1;97m",
}

var Red = color{"red"}
var Green = color{"green"}
var Yellow = color{"green"}
var Blue = color{"blue"}
var Magenta = color{"magenta"}
var Cyan = color{"cyan"}
var Gray = color{"gray"}
var White = color{"white"}
var LightRed = color{"lightred"}
var LightGreen = color{"lightgreen"}
var LightYellow = color{"lightyellow"}
var LightBlue = color{"lightblue"}
var LightMagenta = color{"lightmagenta"}
var LightCyan = color{"lightcyan"}
var LightGray = color{"lightgray"}
var LightWhite = color{"lightwhite"}
var Reset = "\033[0m"

func DisplayText(color string, text ...string) {
	color = strings.ToLower(color)
	if _, ok := Colors[color]; !ok {
		panic("Color not found")
	}

	fmt.Printf("%s%s%s", Colors[color], text, Reset)

}

func (c color) DisplayText(text ...string) {
	fmt.Printf("%s%s%s", Colors[c.Color], text, Reset)
}

func ColorList() map[string]string {
	return Colors
}
