package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF913C(in io.Reader, out io.Writer) {
	search := func(l, r int64, f func(int64) bool) int64 {
		i, j := l, r
		for i < j {
			h := (i + j) >> 1
			if f(h) {
				j = h
			} else {
				i = h + 1
			}
		}
		return i
	}
	var n, l int
	Fscan(in, &n, &l)
	c := [31]int64{}
	for i := 0; i < n; i++ {
		Fscan(in, &c[i])
		if i > 0 && c[i] > 2*c[i-1] {
			c[i] = 2 * c[i-1]
		}
	}
	for i := n; i < 31; i++ {
		c[i] = c[i-1] << 1
	}
	ans := search(1, c[30], func(m int64) bool {
		s := 0
		for i := 29; i >= 0; i-- {
			if c[i] <= m {
				m -= c[i]
				s |= 1 << i
			}
		}
		return s >= l
	})
	Fprint(out, ans)
}

//func main() { CF913C(os.Stdin, os.Stdout) }
