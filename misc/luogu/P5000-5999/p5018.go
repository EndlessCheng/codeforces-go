package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p5018(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	g := make([][3]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &g[i][0])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &g[i][1], &g[i][2])
	}
	sz := make([]int, n+1)
	var dfs func(int) int
	dfs = func(v int) int {
		if v < 0 {
			return 0
		}
		sz[v] = dfs(g[v][1]) + dfs(g[v][2]) + 1
		return sz[v]
	}
	dfs(1)

	var f func(int, int) bool
	f = func(p, q int) bool {
		if p < 0 || q < 0 {
			return p == q
		}
		return g[p][0] == g[q][0] && f(g[p][1], g[q][2]) && f(g[p][2], g[q][1])
	}
	for i := 1; i <= n; i++ {
		if f(g[i][1], g[i][2]) {
			ans = max(ans, sz[i])
		}
	}
	Fprint(out, ans)
}

//func main() { p5018(bufio.NewReader(os.Stdin), os.Stdout) }
