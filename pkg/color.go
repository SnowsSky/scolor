package pkg

import (
	"fmt"
	"reflect"
)

func printSlice(s any) {
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		for _, v := range s.([]any) {
			if reflect.TypeOf(s).Kind() == reflect.Slice {
				printSlice(v)
			} else {
				fmt.Print(v)
			}
		}
	} else {
		fmt.Print(s)
	}
}

type Flag uint

type color struct {
	Color string
}

var Colors = map[string]string{
	"red":           "\033[31m",
	"green":         "\033[32m",
	"yellow":        "\033[33m",
	"blue":          "\033[34m",
	"magenta":       "\033[35m",
	"cyan":          "\033[36m",
	"gray":          "\033[37m",
	"white":         "\033[97m",
	"boldred":       "\033[1;31m",
	"boldgreen":     "\033[1;32m",
	"boldyellow":    "\033[1;33m",
	"boldblue":      "\033[1;34m",
	"boldmagenta":   "\033[1;35m",
	"boldcyan":      "\033[1;36m",
	"boldgray":      "\033[1;37m",
	"boldwhite":     "\033[1;97m",
	"bgred":         "\033[41m",
	"bggreen":       "\033[42m",
	"bgyellow":      "\033[43m",
	"bgblue":        "\033[44m",
	"bgmagenta":     "\033[45m",
	"bgcyan":        "\033[46m",
	"bggray":        "\033[47m",
	"bgwhite":       "\033[107m",
	"bgboldred":     "\033[1;41m",
	"bgboldgreen":   "\033[1;42m",
	"bgboldyellow":  "\033[1;43m",
	"bgboldblue":    "\033[1;44m",
	"bgboldmagenta": "\033[1;45m",
	"bgboldcyan":    "\033[1;46m",
	"bgboldgray":    "\033[1;47m",
	"bgboldwhite":   "\033[1;107m",
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

const (
	FgRed Flag = 1 << iota
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgGray
	FgWhite
	FgBoldRed
	FgBoldGreen
	FgBoldYellow
	FgBoldBlue
	FgBoldMagenta
	FgBoldCyan
	FgBoldGray
	FgBoldWhite
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgGray
	BgWhite
	BgBoldRed
	BgBoldGreen
	BgBoldYellow
	BgBoldBlue
	BgBoldMagenta
	BgBoldCyan
	BgBoldGray
	BgBoldWhite
	FgGrey     = FgGray
	FgBoldGrey = FgBoldGray
	BgGrey     = BgGray
	BgBoldGrey = BgBoldGray
)

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

func DisplayText(flags Flag, args ...any) {
	var color string
	if (flags & FgRed) != 0 {
		color = Colors["red"]
	} else if (flags & FgGreen) != 0 {
		color = Colors["green"]
	} else if (flags & FgYellow) != 0 {
		color = Colors["yellow"]
	} else if (flags & FgBlue) != 0 {
		color = Colors["blue"]
	} else if (flags & FgMagenta) != 0 {
		color = Colors["magenta"]
	} else if (flags & FgCyan) != 0 {
		color = Colors["cyan"]
	} else if (flags & FgWhite) != 0 {
		color = Colors["white"]
	} else if (flags&FgGray) != 0 || (flags&FgGrey) != 0 {
		color = Colors["gray"]
	} else if (flags & FgBoldRed) != 0 {
		color = Colors["red"]
	} else if (flags & FgBoldGreen) != 0 {
		color = Colors["green"]
	} else if (flags & FgBoldYellow) != 0 {
		color = Colors["yellow"]
	} else if (flags & FgBoldBlue) != 0 {
		color = Colors["blue"]
	} else if (flags & FgBoldMagenta) != 0 {
		color = Colors["magenta"]
	} else if (flags & FgBoldCyan) != 0 {
		color = Colors["cyan"]
	} else if (flags & FgBoldWhite) != 0 {
		color = Colors["white"]
	} else if (flags&FgBoldGray) != 0 || (flags&FgBoldGrey) != 0 {
		color = Colors["gray"]
	} else {
		color = ""
	}
	fmt.Print(color)
	if (flags & BgRed) != 0 {
		color = Colors["bgred"]
	} else if (flags & BgGreen) != 0 {
		color = Colors["bggreen"]
	} else if (flags & BgYellow) != 0 {
		color = Colors["bgyellow"]
	} else if (flags & BgBlue) != 0 {
		color = Colors["bgblue"]
	} else if (flags & BgMagenta) != 0 {
		color = Colors["bgmagenta"]
	} else if (flags & BgCyan) != 0 {
		color = Colors["bgcyan"]
	} else if (flags & BgWhite) != 0 {
		color = Colors["bgwhite"]
	} else if (flags&BgGray) != 0 || (flags&BgGrey) != 0 {
		color = Colors["bggray"]
	} else if (flags & BgBoldRed) != 0 {
		color = Colors["bgred"]
	} else if (flags & BgBoldGreen) != 0 {
		color = Colors["bggreen"]
	} else if (flags & BgBoldYellow) != 0 {
		color = Colors["bgyellow"]
	} else if (flags & BgBoldBlue) != 0 {
		color = Colors["bgblue"]
	} else if (flags & BgBoldMagenta) != 0 {
		color = Colors["bgmagenta"]
	} else if (flags & BgBoldCyan) != 0 {
		color = Colors["bgcyan"]
	} else if (flags & BgBoldWhite) != 0 {
		color = Colors["bgwhite"]
	} else if (flags&BgBoldGray) != 0 || (flags&BgBoldGrey) != 0 {
		color = Colors["bggray"]
	} else {
		color = ""
	}
	fmt.Print(color)
	for _, arg := range args {
		printSlice(arg)
	}
	fmt.Print(Reset)
}

func DisplayTextln(flags Flag, args ...any) {
	DisplayText(flags, args)
	println()
}
