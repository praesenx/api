package menu

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strconv"
	"strings"
)

func (p Panel) CaptureInput() (*int, error) {
	fmt.Print(ColorYellow + "Select an option: " + ColorReset)
	input, err := p.Reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("%s error reading input: %v %s", ColorRed, err, ColorReset)
	}

	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)

	if err != nil {
		return nil, fmt.Errorf("%s Please enter a valid number. %s", ColorRed, ColorReset)
	}

	return &choice, nil
}

func (p Panel) PrintMenu() {
	// Try to get the terminal width; default to 80 if it fails
	width, _, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil || width < 20 {
		width = 80
	}

	inner := width - 2 // space between the two border chars

	// Build box pieces
	border := "╔" + strings.Repeat("═", inner) + "╗"
	title := "║" + p.CenterText(" Main Menu ", inner) + "║"
	divider := "╠" + strings.Repeat("═", inner) + "╣"
	footer := "╚" + strings.Repeat("═", inner) + "╝"

	// Print in color
	fmt.Println()
	fmt.Println(ColorCyan + border)
	fmt.Println(title)
	fmt.Println(divider)

	p.PrintOption("1) Say Hello", inner)
	p.PrintOption("2) Show Time", inner)
	p.PrintOption("3) Do Something Else", inner)
	p.PrintOption("0) Exit", inner)

	fmt.Println(footer + ColorReset)
}

// PrintOption left-pads a space, writes the text, then fills to the full inner width.
func (p Panel) PrintOption(text string, inner int) {
	content := " " + text
	if len(content) > inner {
		content = content[:inner]
	}
	padding := inner - len(content)
	fmt.Printf("║%s%s║\n", content, strings.Repeat(" ", padding))
}

// CenterText centers s within width, padding with spaces.
func (p Panel) CenterText(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}

	pad := width - len(s)
	left := pad / 2
	right := pad - left

	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
