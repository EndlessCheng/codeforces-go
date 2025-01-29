package main

import (
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	preXor := func(n int) int {
		b := n >> 1 & 1
		if n&1 > 0 {
			return b ^ 1
		}
		return n | b
	}
	var l, r int
	Fscan(in, &l, &r)
	Fprint(out, preXor(r)^preXor(l-1))
}

func main() { run(os.Stdin, os.Stdout) }
