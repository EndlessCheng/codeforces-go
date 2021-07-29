package main

import (
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF679B(in io.Reader, out io.Writer) {
	var m int64
	Fscan(in, &m)
	type pair struct{ n, s int64 }
	var f func(m, n, s int64) pair
	f = func(m, n, s int64) pair {
		if m == 0 {
			return pair{n, s}
		}
		k := int64(math.Cbrt(float64(m)))
		x, y := k*k*k, (k-1)*(k-1)*(k-1)
		p, q := f(m-x, n+1, s+x), f(x-1-y, n+1, s+y)
		if p.n > q.n || p.n == q.n && p.s > q.s {
			return p
		}
		return q
	}
	ans := f(m, 0, 0)
	Fprint(out, ans.n, ans.s)
}

//func main() { CF679B(os.Stdin, os.Stdout) }
