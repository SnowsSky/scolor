package pkg

import "fmt"

type color struct {
	Color string
}

var Colors = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"gray":    "\033[37m",
	"white":   "\033[97m",
}
var Red = color{"red"}
var Green = color{"green"}
var Yellow = color{"green"}
var Blue = color{"blue"}
var Magenta = color{"magenta"}
var Cyan = color{"cyan"}
var Gray = color{"gray"}
var White = color{"white"}
var Reset = "\033[0m"

func DisplayText(color string, text string) {
	if _, ok := Colors[color]; !ok {
		panic("Color not found")
	}

	fmt.Printf("%s%s%s", Colors[color], text, Reset)

}

func (c color) DisplayText(text string) {
	fmt.Printf("%s%s%s", Colors[c.Color], text, Reset)
}
