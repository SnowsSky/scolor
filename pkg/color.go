package pkg

import (
	"fmt"
	"strings"
)

type color struct {
	Color string
}

var Colors = map[string]string{
	"red":         "\033[31m",
	"green":       "\033[32m",
	"yellow":      "\033[33m",
	"blue":        "\033[34m",
	"magenta":     "\033[35m",
	"cyan":        "\033[36m",
	"gray":        "\033[37m",
	"white":       "\033[97m",
	"boldred":     "\033[1;31m",
	"boldgreen":   "\033[1;32m",
	"boldyellow":  "\033[1;33m",
	"boldblue":    "\033[1;34m",
	"boldmagenta": "\033[1;35m",
	"boldcyan":    "\033[1;36m",
	"boldgray":    "\033[1;37m",
	"boldwhite":   "\033[1;97m",
}

var Red = color{"red"}
var Green = color{"green"}
var Yellow = color{"yellow"}
var Blue = color{"blue"}
var Magenta = color{"magenta"}
var Cyan = color{"cyan"}
var Gray = color{"gray"}
var White = color{"white"}
var BoldRed = color{"boldred"}
var BoldGreen = color{"boldgreen"}
var BoldYellow = color{"boldyellow"}
var BoldBlue = color{"boldblue"}
var BoldMagenta = color{"boldmagenta"}
var BoldCyan = color{"boldcyan"}
var BoldGray = color{"boldgray"}
var BoldWhite = color{"boldwhite"}
var Reset = "\033[0m"

func DisplayText(color string, args ...any) {
	color = strings.ToLower(color)
	if _, ok := Colors[color]; !ok {
		panic("Color not found")
	}
	for _, word := range args {
		fmt.Printf("%s%v%s", Colors[color], word, Reset)
	}

}

func DisplayTextln(color string, args ...any) {
	color = strings.ToLower(color)
	if _, ok := Colors[color]; !ok {
		panic("Color not found")
	}
	for _, word := range args {
		fmt.Printf("%s%v%s", Colors[color], word, Reset)
	}
	println()
}
func (c color) DisplayText(args ...any) {
	for _, word := range args {
		fmt.Printf("%s%v%s", Colors[c.Color], word, Reset)
	}
}

func (c color) DisplayTextln(args ...any) {
	for _, word := range args {
		fmt.Printf("%s%v%s", Colors[c.Color], word, Reset)
	}
	println()
}
