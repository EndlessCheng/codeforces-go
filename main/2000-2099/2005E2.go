package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2005E2(in io.Reader, out io.Writer) {
	buf := make([]byte, 4096)
	_i, _n := 0, 0
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	for range rd() {
		t, n, m := rd(), rd(), rd()
		b := make([]int, t+1)
		p := make([]int, n*m+2)
		for i := 1; i <= t; i++ {
			b[i] = rd()
		}
		for i := t; i > 0; i-- {
			p[b[i]] = i
		}
		a := make([][]int, n+1)
		for i := 1; i <= n; i++ {
			a[i] = make([]int, m+1)
			for j := 1; j <= m; j++ {
				a[i][j] = rd()
			}
		}

		f := make([]int, t+2)
		for i := n; i > 0; i-- {
			k := p[a[i][1]]
			for j := 1; j <= m; j++ {
				k = p[a[i][j]]
				if k > 0 && f[k+1] <= j {
					f[k] = max(f[k], j)
				}
			}
		}

		if f[1] > 0 {
			Fprintln(out, "T")
		} else {
			Fprintln(out, "N")
		}
	}
}

//func main() { cf2005E2(os.Stdin, os.Stdout) }
