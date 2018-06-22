package std

import (
	"fmt"

	"github.com/mwf/vgotest/types"
)

// New returns new printer, printing values to stdout
func New() types.Printer {
	return &std{}
}

type std struct{}

func (s *std) Print(v ...interface{}) {
	fmt.Println(v...)
}
