package stderr

import (
	"fmt"
	"os"

	"github.com/mwf/vgotest/types"
)

// New returns new printer, printing values to stderr
func New() types.Printer {
	return &std{}
}

type std struct{}

func (s *std) Print(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}
