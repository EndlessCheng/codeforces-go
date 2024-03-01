package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1822F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, wt, c, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &wt, &c)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ds := make([]struct{ w, fi, se int }, n)
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				mx := dfs(w, v) + wt
				if mx > ds[v].fi {
					ds[v].se = ds[v].fi
					ds[v].fi = mx
					ds[v].w = w
				} else if mx > ds[v].se {
					ds[v].se = mx
				}
			}
			return ds[v].fi
		}
		ans := dfs(0, -1)

		var reroot func(int, int, int, int)
		reroot = func(v, fa, mxFa, cost int) {
			ans = max(ans, max(mxFa, ds[v].fi)-cost)
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				if w != ds[v].w {
					reroot(w, v, max(mxFa, ds[v].fi)+wt, cost+c)
				} else {
					reroot(w, v, max(mxFa, ds[v].se)+wt, cost+c)
				}
			}
		}
		reroot(0, -1, 0, 0)
		Fprintln(out, ans)
	}
}

//func main() { cf1822F(os.Stdin, os.Stdout) }
