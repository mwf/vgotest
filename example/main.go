package main

import (
	"github.com/mwf/vgotest/types"
	"github.com/mwf/vgotest/types/printers/stderr"
	"github.com/mwf/vgotest/types/printers/stdout"
)

func print(p types.Printer, v ...interface{}) {
	p.Print(v...)
}

func main() {
	print(stderr.New(), "It's stderr:", "foo", "bar", 1)
	print(stdout.New(), "It's stdout:", 1, 2, 3, 4, 5)
}
