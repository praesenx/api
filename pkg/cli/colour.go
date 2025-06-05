package cli

import (
	"fmt"
	"os"
	"slices"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var colours = []string{Reset, Red, Green, Yellow, Blue, Magenta, Cyan, Gray, White}

type TextColour struct {
	text    string
	colour  string
	padding bool
}

func MakeTextColour(text string, colour string) TextColour {
	if !slices.Contains(colours, colour) {
		colour = White
	}

	return TextColour{
		text:    text,
		colour:  colour,
		padding: true,
	}
}

func (t TextColour) Print() {
	_, err := fmt.Print(t.String())

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func (t TextColour) Println() {
	_, err := fmt.Println(t.String())

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func (t TextColour) String() string {
	return fmt.Sprintf("%s > %s %s\n", t.colour, t.text, Reset)
}
