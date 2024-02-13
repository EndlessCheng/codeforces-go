package main

import (
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func cf30A(in io.Reader, out io.Writer) {
	var a, b, n int
	Fscan(in, &a, &b, &n)
	if b == 0 {
		Fprint(out, 0)
		return
	}
	if a == 0 || b%a != 0 {
		Fprint(out, "No solution")
		return
	}
	b /= a
	neg := b < 0
	if neg {
		if n%2 == 0 {
			Fprint(out, "No solution")
			return
		}
		b = -b
	}
	f := math.Pow(float64(b), 1/float64(n))
	x := int(math.Round(f))
	s := x
	for i := 1; i < n; i++ {
		s *= x
	}
	if s != b {
		Fprint(out, "No solution")
		return
	}
	if neg {
		x = -x
	}
	Fprint(out, x)
}

//func main() { cf30A(os.Stdin, os.Stdout) }
