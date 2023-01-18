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

	var n, v, w, wt int
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	d := make([]int, n)
	for i := range d {
		Fscan(in, &d[i])
	}

	type result struct{ fi, i, se int }
	res := make([]result, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		cur := result{}
		for _, e := range g[v] {
			if w := e.to; w != fa {
				dfs(w, v)
				r := max(res[w].fi, d[w]) + e.wt
				if r > cur.fi {
					cur = result{r, w, cur.fi}
				} else if r > cur.se {
					cur.se = r
				}
			}
		}
		res[v] = cur
	}
	dfs(0, -1)

	ans := make([]int, n)
	var reroot func(int, int, int)
	reroot = func(v, fa, up int) {
		ans[v] = max(res[v].fi, up)
		up = max(up, d[v])
		for _, e := range g[v] {
			if w := e.to; w != fa {
				down := res[v].fi
				if w == res[v].i {
					down = res[v].se
				}
				reroot(w, v, max(up, down)+e.wt)
			}
		}
	}
	reroot(0, -1, 0)
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
