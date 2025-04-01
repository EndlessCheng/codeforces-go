package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf87D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	type edge struct{ v, w, wt, dep, i int }
	es := make([]edge, n-1)
	g := make([][]int, n)
	for i := range es {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		es[i] = edge{v, w, wt, 0, i + 1}
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa {
				dep[w] = dep[v] + 1
				dfs(w, v)
			}
		}
	}
	dfs(0, -1)

	for i, e := range es {
		es[i].dep = min(dep[e.v], dep[e.w])
	}
	slices.SortFunc(es, func(a, b edge) int { return cmp.Or(a.wt-b.wt, b.dep-a.dep) })

	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
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

	mx := 0
	ans := []int{}
	tmp := make([]int, n)
	for i := 0; i < n-1; {
		st := i
		wt := es[st].wt
		for ; i < n-1 && es[i].wt == wt; i++ {
			v, w := es[i].v, es[i].w
			if dep[v] > dep[w] {
				v, w = w, v
			}
			tmp[i] = sz[w]
			fv := find(v)
			fa[w] = fv
			sz[fv] += sz[w]
		}
		for ; st < i; st++ {
			e := es[st]
			fv := find(e.v)
			s := tmp[st] * (sz[fv] - tmp[st])
			if s > mx {
				mx = s
				ans = []int{e.i}
			} else if s == mx {
				ans = append(ans, e.i)
			}
		}
	}
	Fprintln(out, mx*2, len(ans))
	slices.Sort(ans)
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { cf87D(bufio.NewReader(os.Stdin), os.Stdout) }
