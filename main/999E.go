package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF999E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, st int
	Fscan(in, &n, &m, &st)
	st--
	g := make([][]int, n)
	type edge struct{ v, w int }
	_edges := make([]edge, m)
	for i := range _edges {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		_edges[i] = edge{v, w}
		g[v] = append(g[v], w)
	}

	vis := make([]bool, n)
	var f func(v int)
	f = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	f(st)
	edges := []edge{}
	for _, e := range _edges {
		if !vis[e.v] && !vis[e.w] {
			edges = append(edges, e)
		}
	}

	g = make([][]int, n)
	rg := make([][]int, n)
	for _, e := range edges {
		v, w := e.v, e.w
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	used := make([]bool, n)
	copy(used, vis)
	vs := make([]int, 0, n)
	var dfs func(int)
	dfs = func(v int) {
		used[v] = true
		for _, w := range g[v] {
			if !used[w] {
				dfs(w)
			}
		}
		vs = append(vs, v)
	}
	for v := range g {
		if !used[v] {
			dfs(v)
		}
	}

	used = make([]bool, n)
	copy(used, vis)
	var comp []int
	var rdfs func(int)
	rdfs = func(v int) {
		used[v] = true
		comp = append(comp, v)
		for _, w := range rg[v] {
			if !used[w] {
				rdfs(w)
			}
		}
	}
	comps := [][]int{}
	for i := len(vs) - 1; i >= 0; i-- {
		if v := vs[i]; !used[v] {
			comp = []int{}
			rdfs(v)
			comps = append(comps, comp)
		}
	}

	sccIDs := make([]int, n)
	for i, cp := range comps {
		for _, v := range cp {
			sccIDs[v] = i
		}
	}
	hasIn := make([]bool, len(comps))
	for _, e := range edges {
		if v, w := sccIDs[e.v], sccIDs[e.w]; v != w {
			hasIn[w] = true
		}
	}
	ans := 0
	for _, has := range hasIn {
		if !has {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF999E(os.Stdin, os.Stdout)
//}
