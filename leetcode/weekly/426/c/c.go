package main

// https://space.bilibili.com/206214
func calcTree(edges [][]int, k int) (diameter int, dp func() [][]int16) {
	n := len(edges) + 1
	g := make([][]int, n)
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

	dp = func() [][]int16 {
		cnt := make([][]int16, n)
		for i := range cnt {
			cnt[i] = make([]int16, k+1)
			cnt[i][0] = 1
		}

		var dfs func(int, int)
		dfs = func(x, fa int) {
			for _, y := range g[x] {
				if y == fa {
					continue
				}
				dfs(y, x)
				for i, c := range cnt[y][:k] {
					cnt[x][i+1] += c
				}
			}
		}
		dfs(0, -1)

		var reroot func(int, int)
		reroot = func(x, fa int) {
			for _, y := range g[x] {
				if y == fa {
					continue
				}
				for i := k - 1; i > 0; i-- {
					cnt[y][i+1] += cnt[x][i] - cnt[y][i-1]
				}
				if k > 0 {
					cnt[y][1]++ // x 到 y 的距离是 1
				}
				reroot(y, x)
			}
		}
		reroot(0, -1)

		return cnt
	}

	return
}

func maxTargetNodes(edges1, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	max2 := 0
	if k > 0 {
		diameter, dp := calcTree(edges2, k-1)
		if diameter < k {
			max2 = m // 第二棵树的每个节点都是目标节点
		} else {
			cnt := dp()
			for _, row := range cnt {
				s := int16(0)
				for _, c := range row {
					s += c
				}
				max2 = max(max2, int(s))
			}
		}
	}

	diameter, dp := calcTree(edges1, k)
	ans := make([]int, n)
	if diameter <= k {
		for i := range ans {
			ans[i] = n + max2 // 第一棵树的每个节点都是目标节点
		}
	} else {
		cnt := dp()
		for i, row := range cnt {
			s := int16(0)
			for _, c := range row {
				s += c
			}
			ans[i] = int(s) + max2
		}
	}
	return ans
}
