package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf678C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }
	var n, a, b, p, q int
	Fscan(in, &n, &a, &b, &p, &q)
	Fprint(out, n/a*p+n/b*q-n/lcm(a, b)*min(p, q))
}

//func main() { cf678C(os.Stdin, os.Stdout) }
