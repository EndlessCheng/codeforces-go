package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf161C(in io.Reader, out io.Writer) {
	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)
	var solve func(int, int, int, int, int) int
	solve = func(a, b, c, d, now int) int {
		if a > b || c > d {
			return 0
		}
		l := 1<<now
		for {
			ok := false
			for c > l {
				c -= l
				d -= l
				ok = true
			}
			for a > l {
				a -= l
				b -= l
				ok = true
			}
			for l > max(b, d) {
				now--
				l >>= 1
				ok = true
			}
			if a > c || a == c && d > b {
				a, c = c, a
				b, d = d, b
			}
			if !ok {
				break
			}
		}
		if b >= d {
			return d - c + 1
		}
		return max(solve(a, b, c, l-1, now), solve(1, d-l, a, b, now), b-c+1)
	}
	Fprint(out, solve(a, b, c, d, 30))
}

//func main() { cf161C(os.Stdin, os.Stdout) }
