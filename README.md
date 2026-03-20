# scolor
A simple go library for color messages

# Usage 
See the example file [here](https://github.com/SnowsSky/scolor/blob/main/main.go), so you know how to use it.
```go
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

```

# Screenshot
![result](https://github.com/SnowsSky/scolor/blob/main/_screenshots/screenshot.png)


# Download
Run `go get github.com/SnowsSky/scolor`

# Contribute 
Create a fork of this repo.<br>
Clone your repo.<br>
Create a new branch with for example your name : `git switch -c <username>`.<br>
Make changes.<br>
Commit & push your changes : `git add . && git commit -m "" && git push`.<br>
Open a Pull Request.<br>

