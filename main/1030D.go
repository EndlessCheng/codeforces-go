package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1030D(in io.Reader, out io.Writer) {
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, m, k int64
	Fscan(in, &n, &m, &k)
	if 2*n*m%k != 0 {
		Fprint(out, "NO")
		return
	}
	g := gcd(2*n, k)
	x, y := 2*n/g, m*g/k
	if x > n {
		x /= 2
		y *= 2
	}
	Fprintf(out, "YES\n0 0\n%d 0\n0 %d", x, y)
}

//func main() { CF1030D(os.Stdin, os.Stdout) }
