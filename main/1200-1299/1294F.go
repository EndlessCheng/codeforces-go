package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1294F(_r io.Reader, _w io.Writer) {
	var n, v, w, u, maxD, ans int
	in := bufio.NewReader(_r)
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD = d
			u = v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	maxD = -1
	f(0, -1, 0)
	dv := u
	maxD = -1
	f(dv, -1, 0)
	dw := u

	onPath := make([]bool, n)
	var f2 func(v, fa int) bool
	f2 = func(v, fa int) bool {
		if v == dw {
			onPath[v] = true
			return true
		}
		for _, w := range g[v] {
			if w != fa {
				if f2(w, v) {
					onPath[v] = true
					ans++
					return true
				}
			}
		}
		return false
	}
	f2(dv, -1)

	maxD = 0
	u = -1
	f = func(v, fa, d int) {
		if d > maxD {
			maxD = d
			u = v
		}
		for _, w := range g[v] {
			if w != fa {
				if onPath[w] {
					f(w, v, 0)
				} else {
					f(w, v, d+1)
				}
			}
		}
	}
	f(dv, -1, 0)
	if u == -1 {
		if dv > dw {
			dv, dw = dw, dv
		}
		u = 0
		if dv == 0 {
			u = 1
		}
		if dw == 1 {
			u = 2
		}
	}
	Fprintln(_w, ans+maxD)
	Fprintln(_w, dv+1, dw+1, u+1)
}

//func main() { CF1294F(os.Stdin, os.Stdout) }
