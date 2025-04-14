package main

import (
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var a, b, xor, lb, x, y int
	Fscan(in, &a, &b, &xor)

	for t := xor; t > 0; t ^= lb {
		lb = t & -t
		if a > b {
			x |= lb
			a--
		} else {
			y |= lb
			b--
		}
	}
	if a < 0 || a != b {
		Fprint(out, -1)
		return
	}

	for t := 1<<60 - 1 ^ xor; t > 0 && a > 0; t ^= lb {
		lb = t & -t
		x |= lb
		y |= lb
		a--
	}
	if a > 0 {
		Fprint(out, -1)
		return
	}

	Fprint(out, x, y)
}

func main() { run(os.Stdin, os.Stdout) }
