package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func prim8(g [][]int) (mst int64) {
	n := len(g)
	minW := make([]int, n)
	for i := range minW {
		minW[i] = 2e9
	}
	minW[0] = 0
	used := make([]bool, n)
	for {
		v := -1
		for i, u := range used {
			if !u && (v < 0 || minW[i] < minW[v]) {
				v = i
			}
		}
		if v < 0 {
			break
		}
		used[v] = true
		mst += int64(minW[v])
		for w, wt := range g[v] {
			if wt < minW[w] {
				minW[w] = wt
			}
		}
	}
	return
}

func CF1508C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, wt, xor, maxV int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		xor ^= wt
		es[i] = edge{v - 1, w - 1, wt}
	}

	if n < 1e3 && m >= (n-2)*(n-1)/2 {
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, n)
		}
		for _, e := range es {
			g[e.v][e.w] = e.wt
			g[e.w][e.v] = e.wt
		}
		ans := int64(1e18)
		for i, r := range g {
			for j, wt := range r[:i] {
				if wt == 0 {
					g[i][j] = xor
					g[j][i] = xor
					if mst := prim8(g); mst < ans {
						ans = mst
					}
					g[i][j] = 0
					g[j][i] = 0
				}
			}
		}
		Fprint(out, ans)
		return
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	g := make([][]int, n)
	for _, e := range es {
		g[e.v] = append(g[e.v], e.w)
		g[e.w] = append(g[e.w], e.v)
	}
	for v, ws := range g {
		if len(ws) < len(g[maxV]) {
			maxV = v
		}
	}
	mergeInv := func(v int) {
		has := map[int]bool{v: true}
		for _, w := range g[v] {
			has[w] = true
		}
		for i := range g {
			if !has[i] {
				fa[f(i)] = f(v)
			}
		}
	}
	mergeInv(maxV)
	for v := range g {
		if f(v) != f(maxV) {
			mergeInv(v)
		}
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })
	ans := int64(0)
	for _, e := range es {
		if v, w := f(e.v), f(e.w); v != w {
			fa[v] = w
			ans += int64(e.wt)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1508C(os.Stdin, os.Stdout) }
