package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf31E(in io.Reader, out io.Writer) {
	var n, mx, mask int
	var s []byte
	Fscan(in, &n, &s)
	type pair struct{ s, mask int }
	f := func(s []byte, mul bool) []pair {
		g := make([]pair, n+1)
		for i := range 1 << n {
			var a, b, c int
			for j, v := range s {
				if i>>j&1 > 0 {
					a = a*10 + int(v-'0')
					c++
				} else {
					b = b*10 + int(v-'0')
				}
			}
			res := a + b
			if mul {
				res = a*int(math.Pow10(n-c)) + b*int(math.Pow10(c))
			}
			if res > g[c].s {
				g[c] = pair{res, i}
			}
		}
		return g
	}
	a := f(s[:n], true)
	b := f(s[n:], false)

	for i, p := range a {
		sum := p.s + b[n-i].s
		if sum > mx {
			mx = sum
			mask = p.mask<<n | b[n-i].mask
		}
	}
	for i := range n * 2 {
		Fprintf(out, "%c", "MH"[mask>>i&1])
	}
}

//func main() { cf31E(os.Stdin, os.Stdout) }
