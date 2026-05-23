package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf585C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var x, y int
	Fscan(in, &x, &y)
	if gcd(x, y) > 1 {
		Fprint(out, "Impossible")
		return
	}

	w := 0
	for x != y {
		if x < y {
			x, y = y, x
			w ^= 1
		}
		if w > 0 {
			Fprint(out, (x-1)/y, "B")
		} else {
			Fprint(out, (x-1)/y, "A")
		}
		x = 1 + (x-1)%y
	}
}

//func main() { cf585C(os.Stdin, os.Stdout) }
