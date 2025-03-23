package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf238C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type nb struct{ to, inv int }
	g := make([][]nb, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], nb{w, 0})
		g[w] = append(g[w], nb{v, 1})
	}

	invCnt := make([]int, n)
	var dfs0 func(int, int)
	dfs0 = func(v, fa int) {
		for _, e := range g[v] {
			w := e.to
			if w != fa {
				dfs0(w, v)
				invCnt[0] += e.inv
			}
		}
	}
	dfs0(0, -1)

	var reroot func(int, int)
	reroot = func(v, fa int) {
		for _, e := range g[v] {
			w := e.to
			if w != fa {
				invCnt[w] = invCnt[v] + 1 - e.inv*2
				reroot(w, v)
			}
		}
	}
	reroot(0, -1)

	ans := int(1e9)
	for i, invI := range invCnt {
		var dfs func(int, int, int, int, int)
		dfs = func(v, fa, dep, f0, f1 int) {
			ans = min(ans, (invI+invCnt[v]-dep)/2+f1)
			for _, e := range g[v] {
				w := e.to
				if w != fa {
					dfs(w, v, dep+1, f0+e.inv, min(f1, f0)+1-e.inv)
				}
			}
		}
		dfs(i, -1, 0, 0, 0)
	}
	Fprint(out, ans)
}

//func main() { cf238C(bufio.NewReader(os.Stdin), os.Stdout) }
