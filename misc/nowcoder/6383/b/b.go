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

	type node struct{ dfn, size int }
	nodes := make([]node, n+1)
	dfn := 0
	var build func(v, fa int) int
	build = func(v, fa int) int {
		dfn++
		nodes[v].dfn = dfn
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

	tree := make([]int, n+1)
	add := func(i, val int) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	addRange := func(l, r, val int) { add(l, val); add(r+1, -val) }
	for _, q := range queries {
		o := nodes[q.X]
		addRange(o.dfn, o.dfn+o.size-1, q.Y)
	}

	ans := make([]int64, 0, n)
	for i := 1; i <= n; i++ {
		ans = append(ans, int64(sum(nodes[i].dfn)))
	}
	return ans
}
