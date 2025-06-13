package menu

import (
	"bufio"
	"github.com/oullin/pkg"
)

type Panel struct {
	Reader    *bufio.Reader
	Choice    *int
	Validator *pkg.Validator
}
