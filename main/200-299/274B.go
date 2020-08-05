package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF274B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	g[0] = append(g[0], -1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	// 目标是得到将子树 v 清零所需的 -操作 和 +操作 的次数，这样就能够从下往上计算
	var f func(v, fa int) (mi, mx int64)
	f = func(v, fa int) (mi, mx int64) {
		x := a[v]
		if len(g[v]) == 1 {
			return min(0, x), max(0, x)
		}
		for _, w := range g[v] {
			if w != fa {
				x, y := f(w, v)
				mi = min(mi, x)
				mx = max(mx, y)
			}
		}
		x -= mi + mx
		if x < 0 {
			mi += x
		} else {
			mx += x
		}
		return
	}
	x, y := f(0, -1)
	Fprint(_w, y-x)
}

//func main() { CF274B(os.Stdin, os.Stdout) }
