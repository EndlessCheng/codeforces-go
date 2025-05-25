package main

// https://space.bilibili.com/206214
func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n)
	for _, e := range hierarchy {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
	}

	var dfs func(int) [][2]int
	dfs = func(x int) [][2]int {
		// 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
		subF := make([][2]int, budget+1)
		for _, y := range g[x] {
			fy := dfs(y)
			for j := budget; j >= 0; j-- {
				for jy, p := range fy[:j+1] { // 枚举子树 y 的预算至多为 jy
					for k, resY := range p {
						subF[j][k] = max(subF[j][k], subF[j-jy][k]+resY)
					}
				}
			}
		}

		f := make([][2]int, budget+1)
		for j, p := range subF {
			for k := range 2 {
				cost := present[x] / (k + 1)
				if j >= cost {
					// 不买 x，转移来源是 subF[j][0]
					// 买 x，转移来源为 subF[j-cost][1]，因为对于子树来说，父节点一定买
					f[j][k] = max(p[0], subF[j-cost][1]+future[x]-cost)
				} else {
					f[j][k] = p[0]
				}
			}
		}
		return f
	}

	return dfs(0)[budget][0]
}
