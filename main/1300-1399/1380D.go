package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1380D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, r int
	var x, k, y, s int64
	Fscan(in, &n, &m, &x, &k, &y)
	a := make([]int, n)
	pos := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = i
	}

	l := -1
	f := func() bool {
		cnt := int64(r - l - 1)
		if cnt < 0 {
			return true
		}
		max := 0
		for _, v := range a[l+1 : r] {
			if v > max {
				max = v
			}
		}
		edge := 0
		if l >= 0 {
			edge = a[l]
		}
		if r < n && a[r] > edge {
			edge = a[r]
		}
		if cnt < k {
			if max > edge {
				return true
			}
			s += cnt * y
		} else if k*y < x {
			if max < edge {
				s += cnt * y
			} else {
				s += (cnt-k)*y + x
			}
		} else {
			s += cnt%k*y + cnt/k*x
		}
		return false
	}
	for ; m > 0; m-- {
		Fscan(in, &r)
		r = pos[r]
		if f() {
			Fprint(out, -1)
			return
		}
		l = r
	}
	r = n
	if f() {
		Fprint(out, -1)
		return
	}
	Fprint(out, s)
}

//func main() { CF1380D(os.Stdin, os.Stdout) }
