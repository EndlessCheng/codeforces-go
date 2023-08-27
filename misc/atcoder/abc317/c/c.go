package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 0; i < m; i++ {
		var v, w int
		var wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	ans := 0
	vis := make([]bool, len(g))
	var dfs func(v, s int)
	dfs = func(v, s int) {
		vis[v] = true
		ans = max(ans, s)
		for _, e := range g[v] {
			if !vis[e.to] {

				dfs(e.to, s+e.wt)

			}
		}
		vis[v] = false
	}
	for st := range vis {
		dfs(st, 0)
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
