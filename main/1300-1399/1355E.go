package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1355E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var n int
	var add, rm, mv, d1, d2 int64
	Fscan(in, &n, &add, &rm, &mv)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	mi, mx := a[0], a[n-1]
	for _, v := range a {
		d1 += int64(v - mi)
		d2 += int64(mx - v)
	}
	ans := min(d1*rm, d2*add)
	f := func(h int) (s int64) {
		if mv < add+rm {
			j := n - 1
			b := append([]int(nil), a...)
			for i, v := range b {
				if d := h - v; d > 0 {
					for ; j > i && b[j] > h; j-- {
						if b[j]-h > d {
							s += int64(d) * mv
							b[j] -= d
							d = 0
							break
						}
						s += int64(b[j]-h) * mv
						d -= b[j] - h
						b[j] = h
					}
					s += int64(d) * add
				} else {
					s += int64(-d) * rm
				}
			}
		} else {
			for _, v := range a {
				if d := int64(h - v); d > 0 {
					s += d * add
				} else {
					s += -d * rm
				}
			}
		}
		return
	}
	l, r := mi+1, mx-1
	for r-l > 4 {
		m1 := l + (r-l)/3
		m2 := r - (r-l)/3
		v1, v2 := f(m1), f(m2)
		if v1 < v2 {
			r = m2
		} else {
			l = m1
		}
	}
	for i := l; i <= r; i++ {
		ans = min(ans, f(i))
	}
	Fprint(out, ans)
}

//func main() { CF1355E(os.Stdin, os.Stdout) }
