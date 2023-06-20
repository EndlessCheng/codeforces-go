package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF354A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	
	var n, l, r, ql, qr, s int
	Fscan(in, &n, &l, &r, &ql, &qr)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	s *= r
	ans := s + (n-1)*qr
	for i, v := range a {
		s += v * (l - r)
		cost := s
		cl, cr := i+1, n-1-i
		if cl+1 < cr {
			cost += (cr - cl - 1) * qr
		} else if cr+1 < cl {
			cost += (cl - cr - 1) * ql
		}
		ans = min(ans, cost)
	}
	Fprint(out, ans)
}

//func main() { CF354A(os.Stdin, os.Stdout) }
