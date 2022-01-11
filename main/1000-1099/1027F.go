package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1027F(_r io.Reader, out io.Writer) {
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
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
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	n := r()
	g := make(map[int][]int, n*2)
	for ; n > 0; n-- {
		v, w := r(), r()
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make(map[int]bool, len(g))
	var cntE, cntV, mx, mx2, ans int
	var f func(int)
	f = func(v int) {
		vis[v] = true
		cntV++
		if v > mx {
			mx, mx2 = v, mx
		} else if v > mx2 {
			mx2 = v
		}
		for _, w := range g[v] {
			cntE++
			if !vis[w] {
				f(w)
			}
		}
	}
	for v := range g {
		if !vis[v] {
			cntE, cntV, mx, mx2 = 0, 0, 0, 0
			f(v)
			cntE /= 2
			if cntE > cntV {
				Fprint(out, -1)
				return
			}
			if cntE == cntV {
				ans = max(ans, mx)
			} else {
				ans = max(ans, mx2)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1027F(os.Stdin, os.Stdout) }
