package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf579A(in io.Reader, out io.Writer) {
	var n uint
	Fscan(in, &n)
	Fprint(out, bits.OnesCount(n))
}

//func main() { cf579A(os.Stdin, os.Stdout) }
