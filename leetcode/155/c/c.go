package main

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
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
