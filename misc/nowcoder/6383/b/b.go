package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func solve(n int, edges []*Point, _ int, queries []*Point) []int64 {
	g := make([][]int, n+1)
	for _, e := range edges {
		v, w := e.X, e.Y
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	// 另一种写法是直接在节点上打 lazy tag
	type node struct{ dfn, size int }
	nodes := make([]node, n+1)
	dfn := 0
	ids := make([]int, n+1)
	var build func(v, fa int) int
	build = func(v, fa int) int {
		dfn++
		nodes[v].dfn = dfn
		ids[dfn] = v
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += build(w, v)
			}
		}
		nodes[v].size = sz
		return sz
	}
	build(1, 0)

	diff := make([]int, n+2)
	for _, q := range queries {
		o := nodes[q.X]
		diff[o.dfn] += q.Y
		diff[o.dfn+o.size] -= q.Y
	}

	ans := make([]int64, n)
	s := 0
	for i := 1; i <= n; i++ {
		s += diff[i]
		ans[ids[i]-1] = int64(s)
	}
	return ans
}
