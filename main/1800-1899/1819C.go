package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1819C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	maxD, u := -1, 0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 0)
	dv := u
	maxD = -1
	f(u, -1, 0)
	dw := u

	path := []int{}
	var findDiameterPath func(v, fa int) bool
	findDiameterPath = func(v, fa int) bool {
		if v == dw {
			path = append(path, v)
			return true
		}
		for _, w := range g[v] {
			if w != fa && findDiameterPath(w, v) {
				path = append(path, v)
				return true
			}
		}
		return false
	}
	findDiameterPath(dv, -1)

	ans := make([]int, 0, n)
	m := len(path)
	for i := 0; i < m; i += 2 {
		ans = append(ans, path[i])
		if i+2 < m {
			for _, w := range g[path[i+1]] {
				if w != path[i] && w != path[i+2] {
					ans = append(ans, w)
				}
			}
		}
	}
	for i := m - 1 - m%2; i > 0; i -= 2 {
		ans = append(ans, path[i])
		if i > 1 {
			for _, w := range g[path[i-1]] {
				if w != path[i] && w != path[i-2] {
					ans = append(ans, w)
				}
			}
		}
	}
	if len(ans) < n {
		Fprint(out, "No")
		return
	}
	Fprintln(out, "Yes")
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF1819C(os.Stdin, os.Stdout) }
