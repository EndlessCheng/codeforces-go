package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1468I(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, a, b, c, d int
	Fscan(in, &n, &a, &b, &c, &d)
	for c != 0 {
		b -= a / c * d
		a %= c
		a, c = c, a
		b, d = d, b
	}

	if abs(a*d) != n {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i := range abs(a) {
		for j := range abs(d) {
			Fprintln(out, i, j)
		}
	}
}

//func main() { cf1468I(os.Stdin, os.Stdout) }
