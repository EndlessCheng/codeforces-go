package main

import (
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	preXor := func(n int) int {
		switch n % 4 {
		case 0:
			return n
		case 1:
			return 1
		case 2:
			return n + 1
		default:
			return 0
		}
	}
	var l, r int
	Fscan(in, &l, &r)
	Fprint(out, preXor(r)^preXor(l-1))
}

func main() { run(os.Stdin, os.Stdout) }
