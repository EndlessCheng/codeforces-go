package main

// https://space.bilibili.com/206214
func calcTree(edges [][]int, k int) (diameter int, dfs func(int, int, int) int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfsDiameter func(int, int) int
	dfsDiameter = func(x, fa int) (maxLen int) {
		for _, y := range g[x] {
			if y != fa {
				subLen := dfsDiameter(y, x) + 1
				diameter = max(diameter, maxLen+subLen)
				maxLen = max(maxLen, subLen)
			}
		}
		return
	}
	dfsDiameter(0, -1)

	dfs = func(x, fa, d int) int {
		if d > k {
			return 0
		}
		cnt := 1
		for _, y := range g[x] {
			if y != fa {
				cnt += dfs(y, x, d+1)
			}
		}
		return cnt
	}

	return
}

func maxTargetNodes(edges1, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	max2 := 0
	if k > 0 {
		diameter, dfs := calcTree(edges2, k-1)
		if diameter < k {
			max2 = m // 第二棵树的每个节点都是目标节点
		} else {
			for i := range m {
				max2 = max(max2, dfs(i, -1, 0))
			}
		}
	}

	diameter, dfs := calcTree(edges1, k)
	ans := make([]int, n)
	if diameter <= k {
		for i := range ans {
			ans[i] = n + max2 // 第一棵树的每个节点都是目标节点
		}
	} else {
		for i := range ans {
			ans[i] = dfs(i, -1, 0) + max2
		}
	}
	return ans
}
