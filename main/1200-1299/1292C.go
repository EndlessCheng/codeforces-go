package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1292C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

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

	pa := make([][]int, n)
	sz := make([][]int, n)
	for i := range pa {
		pa[i] = make([]int, n)
		sz[i] = make([]int, n)
		var f func(int, int)
		f = func(v, fa int) {
			pa[i][v] = fa
			sz[i][v] = 1
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
					sz[i][v] += sz[i][w]
				}
			}
		}
		f(i, -1)
	}

	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	var f func(int, int) int64
	f = func(v, w int) int64 {
		if v == w {
			return 0
		}
		if dp[v][w] == 0 {
			dp[v][w] = max(f(pa[w][v], w), f(v, pa[v][w])) + int64(sz[v][w]*sz[w][v])
		}
		return dp[v][w]
	}
	ans := int64(0)
	for v := 0; v < n; v++ {
		for w := v + 1; w < n; w++ {
			ans = max(ans, f(v, w))
		}
	}
	Fprint(out, ans)
}

//func main() { CF1292C(os.Stdin, os.Stdout) }
