package main

import "math"

// https://space.bilibili.com/206214
func maxScore(parent []int, op []int, score []int) int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	// 返回子树 x 的最大分数和
	var dfs func(int) int
	dfs = func(x int) int {
		sum := score[x]
		if g[x] == nil { // x 是叶子
			return sum
		}
		mn := math.MaxInt
		for _, y := range g[x] {
			s := dfs(y)     // 子树 y 的最大分数和
			sum += s        // 累加分数和
			mn = min(mn, s) // 维护最小的分数和
		}
		if mn < 0 || op[x] > 0 {
			sum -= mn // 去掉最小的分数和
		}
		return sum
	}

	return dfs(0)
}
