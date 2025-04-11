package cli

import (
	"errors"
	"fmt"
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

var colours []string = []string{Reset, Red, Green, Yellow, Blue, Magenta, Cyan, Gray, White}

type TextColour struct {
	text    string
	color   string
	padding bool
}

func MakeTextColour(text string, colour string) (*TextColour, error) {
	if isInvalidValidColor(colour) {
		return nil, errors.New("the given colour is invalid")
	}

	return &TextColour{
		text:    text,
		padding: false,
		color:   colour,
	}, nil
}

func MakePaddedTextColour(text string, colour string) (*TextColour, error) {
	if isInvalidValidColor(colour) {
		return nil, errors.New("the given colour is invalid")
	}

	textColour, err := MakeTextColour(text, colour)

	if err != nil {
		return nil, err
	}

	textColour.padding = true

	return textColour, nil
}

func (t TextColour) Get() string {
	if t.padding == false {
		return t.color + t.text + Reset
	}

	return fmt.Sprintf("\n     ----- %s%s%s -----     \n\n", t.color, t.text, Reset)
}

func isInvalidValidColor(seed string) bool {
	return !slices.Contains(colours, seed)
}
