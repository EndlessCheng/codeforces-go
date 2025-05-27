package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n)
	for _, e := range hierarchy {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
	}

	var dfs func(int) [2][]int
	dfs = func(x int) [2][]int {
		// 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
		subF := [2][]int{make([]int, budget+1), make([]int, budget+1)}
		for i := 1; i <= budget; i++ {
			subF[0][i] = math.MinInt / 2 // 表示不存在对应的花费总和
			subF[1][i] = math.MinInt / 2
		}
		for _, y := range g[x] {
			fy := dfs(y)
			for k, fyk := range fy {
				nf := make([]int, budget+1)
				for i := 1; i <= budget; i++ {
					nf[i] = math.MinInt / 2
				}
				for jy, resY := range fyk {
					if resY < 0 { // 重要优化：物品价值为负数，一定不选
						continue
					}
					for j := jy; j <= budget; j++ {
						nf[j] = max(nf[j], subF[k][j-jy]+resY)
					}
				}
				subF[k] = nf
			}
		}

		f := [2][]int{}
		for k := range 2 {
			// 不买 x，转移来源为 subF[0]，因为对于子树来说，父节点一定不买
			f[k] = slices.Clone(subF[0])
			cost := present[x] / (k + 1)
			// 买 x，转移来源为 subF[1]，因为对于子树来说，父节点一定买
			for j := cost; j <= budget; j++ {
				f[k][j] = max(f[k][j], subF[1][j-cost]+future[x]-cost)
			}
		}
		return f
	}

	return slices.Max(dfs(0)[0])
}
