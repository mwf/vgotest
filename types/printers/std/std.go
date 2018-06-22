package std

import (
	"fmt"

	"github.com/mwf/vgotest/types"
)

type Std struct{}

var _ types.Printer = &Std{}

func (s *Std) Print(v ...interface{}) {
	fmt.Println(v...)
}
