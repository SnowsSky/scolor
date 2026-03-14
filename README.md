# scolor
A simple go library for color messages

# Usage 
See the example file [here](https://github.com/SnowsSky/scolor/blob/main/main.go), so you know how to use it.
```go
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

```

# Screenshot
![result](https://github.com/SnowsSky/scolor/blob/main/screenshots/ss.png)


# Download
Run `go get github.com/SnowsSky/scolor`

# Contribute 
Create a fork of this repo.<br>
Clone your repo.<br>
Create a new branch with for example your name : `git switch -c <username>`.<br>
Make changes.<br>
Commit & push your changes : `git add . && git commit -m "" && git push`.<br>
Open a Pull Request.<br>