package menu

import (
	"bufio"
)

const (
	ColorReset  = "\033[0m"
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
)

type Panel struct {
	Reader *bufio.Reader
}
