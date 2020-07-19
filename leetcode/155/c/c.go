package main

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
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
	for _, p := range pairs {
		fa[find(p[0])] = find(p[1])
	}

	type pair struct {
		b []byte
		i []int
	}
	group := make([]pair, n)
	for i := range s {
		f := find(i)
		group[f].b = append(group[f].b, s[i])
		group[f].i = append(group[f].i, i)
	}

	ans := make([]byte, n)
	for _, p := range group {
		bs, ps := p.b, p.i
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		sort.Ints(ps)
		for i, v := range bs {
			ans[ps[i]] = v
		}
	}
	return string(ans)
}

func smallestStringWithSwapsCC(s string, pairs [][]int) string {
	n := len(s)
	ans := make([]byte, n)
	g := make([][]int, n)
	for _, p := range pairs {
		v, w := p[0], p[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make([]bool, n)
	var comp []int
	var f func(int)
	f = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i, vi := range vis {
		if !vi {
			comp = []int{}
			f(i)
			sort.Ints(comp)
			cs := []byte{}
			for _, v := range comp {
				cs = append(cs, s[v])
			}
			sort.Slice(cs, func(i, j int) bool { return cs[i] < cs[j] })
			for j, c := range cs {
				ans[comp[j]] = c
			}
		}
	}
	return string(ans)
}
