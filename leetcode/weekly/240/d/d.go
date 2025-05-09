package main

// github.com/EndlessCheng/codeforces-go
func largestPathValue(colors string, edges [][]int) (ans int) {
	n := len(colors)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		if x == y { // 自环
			return -1
		}
		g[x] = append(g[x], y)
		deg[y]++
	}

	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 0 {
			q = append(q, i) // 入度为 0 的点入队
		}
	}

	f := make([][26]int, n)
	for len(q) > 0 {
		x := q[0] // x 的所有转移来源都计算完毕，也都更新到 f[x] 中
		q = q[1:]
		ch := colors[x] - 'a'
		f[x][ch]++
		ans = max(ans, f[x][ch])
		for _, y := range g[x] {
			for i, cnt := range f[x] {
				f[y][i] = max(f[y][i], cnt) // 刷表法，更新邻居的最大值
			}
			deg[y]--
			if deg[y] == 0 {
				q = append(q, y)
			}
		}
	}

	if cap(q) > 0 { // 有节点没入队，说明有环
		return -1
	}
	return
}

func largestPathValue1(colors string, edges [][]int) (ans int) {
	n := len(colors)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		if x == y { // 自环
			return -1
		}
		g[x] = append(g[x], y)
	}

	memo := make([][]int, n)
	var dfs func(int) []int
	dfs = func(x int) []int {
		if memo[x] != nil { // x 计算中或者计算过
			return memo[x] // 如果 memo[x] 是空列表，返回空列表，表示有环
		}
		memo[x] = []int{} // 用空列表表示计算中
		res := make([]int, 26)
		for _, y := range g[x] {
			cy := dfs(y)
			if len(cy) == 0 { // 有环
				return cy
			}
			for i, c := range cy {
				res[i] = max(res[i], c)
			}
		}
		res[colors[x]-'a']++
		memo[x] = res // 记忆化，同时也表示 x 计算完毕
		return res
	}

	for x, ch := range colors {
		res := dfs(x)
		if len(res) == 0 { // 有环
			return -1
		}
		ans = max(ans, res[ch-'a'])
	}
	return
}
