package main

// https://space.bilibili.com/206214
func minimumCost(n int, edges, query [][]int) []int {
	type edge struct{ to, w int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	ids := make([]int, n)
	for i := range ids {
		ids[i] = -1
	}
	ccAnd := []int{}
	var dfs func(int) int
	dfs = func(x int) int {
		ids[x] = len(ccAnd)
		and := -1
		for _, e := range g[x] {
			and &= e.w
			if ids[e.to] < 0 {
				and &= dfs(e.to)
			}
		}
		return and
	}
	for i, id := range ids {
		if id < 0 {
			ccAnd = append(ccAnd, dfs(i))
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if s == t {
			continue
		}
		if ids[s] != ids[t] {
			ans[i] = -1
		} else {
			ans[i] = ccAnd[ids[s]]
		}
	}
	return ans
}

func minimumCost2(n int, edges, query [][]int) []int {
	fa := make([]int, n)
	and := make([]int, n)
	for i := range fa {
		fa[i] = i
		and[i] = -1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for _, e := range edges {
		x, y := find(e[0]), find(e[1])
		and[y] &= e[2]
		if x != y {
			and[y] &= and[x]
			fa[x] = y
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if s == t {
			continue
		}
		if find(s) != find(t) {
			ans[i] = -1
		} else {
			ans[i] = and[find(s)]
		}
	}
	return ans
}
