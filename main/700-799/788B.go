package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF788B(_r io.Reader, out io.Writer) {
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m := r(), r()
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	cnt := int64(0)
	d := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < m; i++ {
		v, w := r()-1, r()-1
		if fv, fw := f(v), f(w); fv != fw {
			fa[fv] = fw
			n--
		}
		if v == w {
			cnt++
		} else {
			d[v]++
			d[w]++
		}
		vis[v] = true
		vis[w] = true
	}
	for _, b := range vis {
		if !b {
			n--
		}
	}
	if n > 1 {
		Fprint(out, 0)
		return
	}
	ans := cnt*(int64(m)-cnt) + cnt*(cnt-1)/2
	for _, c := range d {
		ans += int64(c) * int64(c-1) / 2
	}
	Fprint(out, ans)
}

//func main() { CF788B(os.Stdin, os.Stdout) }
