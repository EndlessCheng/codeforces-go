package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumScore(scores []int, edges [][]int) int {
	type nb struct{ to, s int }
	g := make([][]nb, len(scores))
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], nb{y, scores[y]})
		g[y] = append(g[y], nb{x, scores[x]})
	}
	for i, vs := range g {
		sort.Slice(vs, func(i, j int) bool { return vs[i].s > vs[j].s })
		if len(vs) > 3 {
			vs = vs[:3]
		}
		g[i] = vs
	}

	ans := -1
	for _, e := range edges {
		x, y := e[0], e[1]
		for _, p := range g[x] {
			for _, q := range g[y] {
				if p.to != y && q.to != x && p.to != q.to {
					ans = max(ans, p.s+scores[x]+scores[y]+q.s)
				}
			}
		}
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
