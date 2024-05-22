package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func prim(g [][]int) (mst int) {
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
		mst += minW[v]
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
	var n, m, xor int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		xor ^= wt
		es[i] = edge{v - 1, w - 1, wt}
	}

	// 反图可能是生成树
	if m >= (n-2)*(n-1)/2 {
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, n)
		}
		for _, e := range es {
			g[e.v][e.w] = e.wt
			g[e.w][e.v] = e.wt
		}
		ans := int(1e18)
		for i, r := range g {
			for j, wt := range r[:i] {
				if wt == 0 { // 枚举是 xor 的边
					g[i][j] = xor
					g[j][i] = xor
					ans = min(ans, prim(g))
					g[i][j] = 0
					g[j][i] = 0
				}
			}
		}
		Fprint(out, ans)
		return
	}

	// 反图一定不是生成树，用一条不在生成树上的边放 xor
	// 所以只需要求反图连通块，再用 Kruskal 补上原图的边
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	g := make([][]int, n)
	for _, e := range es {
		g[e.v] = append(g[e.v], e.w)
		g[e.w] = append(g[e.w], e.v)
	}
	maxV := 0
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
				fa[find(i)] = find(v)
			}
		}
	}
	mergeInv(maxV)
	for v := range g {
		if find(v) != find(maxV) {
			mergeInv(v)
		}
	}

	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })
	ans := 0
	for _, e := range es {
		v, w := find(e.v), find(e.w)
		if v != w {
			fa[v] = w
			ans += e.wt
		}
	}
	Fprint(out, ans)
}

//func main() { CF1508C(os.Stdin, os.Stdout) }
