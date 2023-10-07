package main

// github.com/EndlessCheng/codeforces-go
func maximumInvitations(g []int) int { // favorite 就是内向基环森林 g
	n := len(g)
	deg := make([]int, n) // g 上每个节点的入度
	for _, w := range g {
		deg[w]++
	}

	maxDepth := make([]int, n) // maxDepth[i] 表示在 rg 上节点 i 到其所在子树的叶节点的最大距离
	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 拓扑排序，剪掉 g 上的所有树枝
		v := q[0]
		q = q[1:]
		w := g[v] // v 只有一条出边
		maxDepth[v]++
		maxDepth[w] = max(maxDepth[w], maxDepth[v])
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range deg {
		if d <= 0 {
			continue
		}
		// 遍历基环上的点（拓扑排序后入度大于 0）
		deg[i] = 0
		ringSize := 1
		for v := g[i]; v != i; v = g[v] {
			deg[v] = -1 // 将基环上的点的入度标记为 -1，避免重复访问
			ringSize++
		}
		if ringSize == 2 { // 基环大小为 2
			sumChainSize += maxDepth[i] + maxDepth[g[i]] + 2 // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 取所有基环的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}

func max(a, b int) int { if b > a { return b }; return a }
