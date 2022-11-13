package main

import "math"

// https://space.bilibili.com/206214
// 评价：读题题。非常标准的树上 DFS 题目，想清楚了就是两遍 DFS 模拟的事情。
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	bobTime := make([]int, n) // bobTime[x] 表示 bob 访问节点 x 的时间
	for i := range bobTime {
		bobTime[i] = n // 也可以初始化成 inf
	}
	var dfsBob func(int, int, int) bool
	dfsBob = func(x, fa, t int) bool {
		if x == 0 {
			bobTime[x] = t
			return true
		}
		found0 := false
		for _, y := range g[x] {
			if y != fa && dfsBob(y, x, t+1) {
				found0 = true
			}
		}
		if found0 {
			bobTime[x] = t // 只有可以到达 0 才标记访问时间
		}
		return found0
	}
	dfsBob(bob, -1, 0)

	g[0] = append(g[0], -1) // 防止把根节点当作叶子
	ans := math.MinInt32
	var dfsAlice func(int, int, int, int)
	dfsAlice = func(x, fa, aliceTime, sum int) {
		if aliceTime < bobTime[x] {
			sum += amount[x]
		} else if aliceTime == bobTime[x] {
			sum += amount[x] / 2
		}
		if len(g[x]) == 1 { // 叶子
			ans = max(ans, sum) // 更新答案
			return
		}
		for _, y := range g[x] {
			if y != fa {
				dfsAlice(y, x, aliceTime+1, sum)
			}
		}
	}
	dfsAlice(0, -1, 0, 0)
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
