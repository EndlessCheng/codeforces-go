package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1092E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, u, o, mx, maxO int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n+1)
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		vis[v] = true
		if d > mx {
			mx, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	var f2 func(v, fa, d int) bool
	f2 = func(v, fa, d int) bool {
		if v == u {
			if d == mx/2 {
				o = v
			}
			return true
		}
		for _, w := range g[v] {
			if w != fa && f2(w, v, d+1) {
				if d == mx/2 {
					o = v
				}
				return true
			}
		}
		return false
	}
	os := []int{}
	maxD := -1
	for i := 1; i <= n; i++ {
		if !vis[i] {
			mx = -1
			f(i, 0, 0)
			mx = -1
			x := u
			f(x, 0, 0)
			o = 0
			f2(x, 0, 0)
			os = append(os, o)
			if mx > maxD {
				maxD = mx
				maxO = o
			}
		}
	}
	for _, o := range os {
		if o != maxO {
			g[maxO] = append(g[maxO], o)
			g[o] = append(g[o], maxO)
		}
	}
	mx = -1
	f(1, 0, 0)
	mx = -1
	f(u, 0, 0)
	Fprintln(out, mx)
	for _, o := range os {
		if o != maxO {
			Fprintln(out, maxO, o)
		}
	}
}

//func main() { CF1092E(os.Stdin, os.Stdout) }
