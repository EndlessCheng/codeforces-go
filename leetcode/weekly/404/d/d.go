package main

// https://space.bilibili.com/206214
func diameter(edges [][]int) (res int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) (maxLen int) {
		for _, y := range g[x] {
			if y != fa {
				subLen := dfs(y, x) + 1
				res = max(res, maxLen+subLen)
				maxLen = max(maxLen, subLen)
			}
		}
		return
	}
	dfs(0, -1)
	return
}

func minimumDiameterAfterMerge(edges1, edges2 [][]int) int {
	d1 := diameter(edges1)
	d2 := diameter(edges2)
	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
}
