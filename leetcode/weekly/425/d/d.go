package main

import "slices"

// https://space.bilibili.com/206214
func maximizeSumOfWeights(edges [][]int, k int) int64 {
	type edge struct{ to, wt int }
	g := make([][]edge, len(edges)+1)
	sumWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
		sumWt += wt
	}

	// 优化
	simple := true
	for _, to := range g {
		if len(to) > k {
			simple = false
			break
		}
	}
	if simple {
		return int64(sumWt)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notChoose := 0
		inc := []int{}
		for _, e := range g[x] {
			y := e.to
			if y == fa {
				continue
			}
			nc, c := dfs(y, x)
			notChoose += nc // 先都不选
			if d := c + e.wt - nc; d > 0 {
				inc = append(inc, d)
			}
		}

		// 再选增量最大的 k 个或者 k-1 个
		slices.SortFunc(inc, func(a, b int) int { return b - a })
		for i := range min(len(inc), k-1) {
			notChoose += inc[i]
		}
		choose := notChoose
		if len(inc) >= k {
			notChoose += inc[k-1]
		}
		return notChoose, choose
	}
	nc, _ := dfs(0, -1) // notChoose >= choose
	return int64(nc)
}
