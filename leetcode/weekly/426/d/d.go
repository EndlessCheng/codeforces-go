package main

// https://space.bilibili.com/206214
func count(edges [][]int) (g [][]int, cnt [2]int) {
	g = make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int, int)
	dfs = func(x, fa, d int) {
		cnt[d]++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, d^1)
			}
		}
	}
	dfs(0, -1, 0)
	return
}

func maxTargetNodes(edges1, edges2 [][]int) []int {
	_, cnt2 := count(edges2)
	max2 := max(cnt2[0], cnt2[1])

	g, cnt1 := count(edges1)
	ans := make([]int, len(g))
	var dfs func(int, int, int)
	dfs = func(x, fa, d int) {
		ans[x] = cnt1[d] + max2
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, d^1)
			}
		}
	}
	dfs(0, -1, 0)
	return ans
}
