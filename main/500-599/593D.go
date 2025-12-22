package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf593D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, v, w, wt int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt int }
	es := make([]edge, n-1)
	type nb struct{ to, i int }
	g := make([][]nb, n)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		es[i] = edge{v, w, wt}
		g[v] = append(g[v], nb{w, i})
		g[w] = append(g[w], nb{v, i})
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	paInfo := make([]nb, n)
	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, pa int) {
		for _, e := range g[v] {
			w := e.to
			if w == pa {
				continue
			}
			if es[e.i].wt == 1 {
				fa[w] = find(v)
			}
			paInfo[w] = nb{v, e.i}
			dep[w] = dep[v] + 1
			dfs(w, v)
		}
	}
	dfs(0, -1)

	for range m {
		Fscan(in, &op, &v, &w)
		if op == 1 {
			Fscan(in, &wt)
			v = find(v - 1)
			w = find(w - 1)
			for wt > 0 && v != w {
				if dep[v] > dep[w] {
					v, w = w, v
				}
				wt /= es[paInfo[w].i].wt
				w = find(paInfo[w].to)
			}
			Fprintln(out, wt)
		} else {
			e := &es[v-1]
			if w > 1 {
				e.wt = w
				continue
			}
			v, w := e.v, e.w
			if dep[v] > dep[w] {
				v, w = w, v
			}
			fa[w] = find(v)
		}
	}
}

//func main() { cf593D(bufio.NewReader(os.Stdin), os.Stdout) }
