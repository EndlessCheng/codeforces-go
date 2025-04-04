package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf294E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type nb struct{ to, wt, i int }
	g := make([][]nb, n)
	for i := range n - 1 {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt, i})
		g[w] = append(g[w], nb{v, wt, i})
	}

	type tuple struct{ v, w, wt, i, sizeW int }
	es := make([]tuple, 0, n-1)
	var init func(int, int) int
	init = func(v, fa int) int {
		size := 1
		for _, e := range g[v] {
			w := e.to
			if w != fa {
				sz := init(w, v)
				es = append(es, tuple{v, w, e.wt, e.i, sz})
				size += sz
			}
		}
		return size
	}
	init(0, -1)

	find := func(st, del, n int) (centroid int) {
		minOfMaxSubSize := int(1e9)
		var f func(int, int) int
		f = func(v, fa int) int {
			size := 1
			maxSubSize := 0
			for _, e := range g[v] {
				w := e.to
				if w != fa && e.i != del {
					sz := f(w, v)
					maxSubSize = max(maxSubSize, sz)
					size += sz
				}
			}
			maxSubSize = max(maxSubSize, n-size)
			if maxSubSize < minOfMaxSubSize {
				minOfMaxSubSize = maxSubSize
				centroid = v
			}
			return size
		}
		f(st, -1)
		return
	}

	ans := int(1e18)
	for _, e := range es {
		del := e.i
		v := find(e.v, del, n-e.sizeW)
		w := find(e.w, del, e.sizeW)
		g[v] = append(g[v], nb{w, e.wt, -1})
		g[w] = append(g[w], nb{v, e.wt, -1})
		res := 0
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			size := 1
			for _, e := range g[v] {
				w := e.to
				if w != fa && e.i != del {
					sz := dfs(w, v)
					res += sz * (n - sz) * e.wt
					size += sz
				}
			}
			return size
		}
		dfs(0, -1)
		ans = min(ans, res)
		g[v] = g[v][:len(g[v])-1]
		g[w] = g[w][:len(g[w])-1]
	}
	Fprint(out, ans)
}

//func main() { cf294E(bufio.NewReader(os.Stdin), os.Stdout) }
