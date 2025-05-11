package main

// https://space.bilibili.com/206214
func maxWeight1(n int, edges [][]int, k int, t int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
	}

	ans := -1
	type tuple struct{ x, i, s int }
	vis := map[tuple]bool{}
	var dfs func(int, int, int)
	dfs = func(x, i, s int) {
		if i == k {
			ans = max(ans, s)
			return
		}
		args := tuple{x, i, s}
		if vis[args] {
			return
		}
		vis[args] = true
		for _, e := range g[x] {
			if s+e.wt < t {
				dfs(e.to, i+1, s+e.wt)
			}
		}
	}
	for x := range n {
		dfs(x, 0, 0)
	}
	return ans
}

func maxWeight2(n int, edges [][]int, k int, t int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
	}

	ans := -1
	f := make([][]map[int]struct{}, n)
	for i := range f {
		f[i] = make([]map[int]struct{}, k+1)
		for j := range f[i] {
			f[i][j] = map[int]struct{}{}
		}
	}
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		f[x][0][0] = struct{}{}  // x 单独一个点，路径边权和为 0
		for s := range f[x][k] { // 恰好 k 条边
			ans = max(ans, s)
		}
		for _, e := range g[x] {
			y, wt := e.to, e.wt
			for i, st := range f[x][:k] {
				for s := range st {
					if s+wt < t {
						f[y][i+1][s+wt] = struct{}{}
					}
				}
			}
			deg[y]--
			if deg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	return ans
}

func maxWeight(n int, edges [][]int, k int, t int) int {
	f := make([][]map[int]struct{}, k+1)
	for i := range f {
		f[i] = make([]map[int]struct{}, n)
		for j := range f[i] {
			f[i][j] = map[int]struct{}{}
		}
	}
	for i := range f[0] {
		f[0][i][0] = struct{}{}
	}
	for i, sets := range f[:k] {
		for _, e := range edges {
			x, y, wt := e[0], e[1], e[2]
			for s := range sets[x] {
				if s+wt < t {
					f[i+1][y][s+wt] = struct{}{}
				}
			}
		}
	}

	ans := -1
	for _, set := range f[k] {
		for s := range set {
			ans = max(ans, s)
		}
	}
	return ans
}