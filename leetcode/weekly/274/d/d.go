package main

// github.com/EndlessCheng/codeforces-go
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	g := make([][]int, n)
	rg := make([][]int, n) // 图 g 的反图
	deg := make([]int, n)  // 图 g 上每个节点的入度
	for v, w := range favorite {
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
		deg[w]++
	}

	// 拓扑排序，剪掉图 g 上的所有树枝
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	// 寻找图 g 上的基环
	comp := []int{}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}

	// 通过反图 rg 寻找树枝上最深的链
	maxDepth := 0
	var rdfs func(int, int, int)
	rdfs = func(v, fa, depth int) {
		maxDepth = max(maxDepth, depth)
		for _, w := range rg[v] {
			if w != fa {
				rdfs(w, v, depth+1)
			}
		}
	}

	sumListSize, maxCycleSize := 0, 0
	for i, b := range vis {
		if !b && deg[i] > 0 { // 遍历基环上的点（拓扑排序后入度不为 0）
			comp = []int{}
			dfs(i)
			if len(comp) == 2 { // 基环大小为 2
				v, w := comp[0], comp[1]
				maxDepth = 0
				rdfs(v, w, 1)
				sumListSize += maxDepth // 累加 v 这一侧的最长链的长度
				maxDepth = 0
				rdfs(w, v, 1)
				sumListSize += maxDepth // 累加 w 这一侧的最长链的长度
			} else {
				maxCycleSize = max(maxCycleSize, len(comp)) // 取所有基环的最大值
			}
		}
	}
	return max(maxCycleSize, sumListSize)
}

func max(a, b int) int { if b > a { return b }; return a }
