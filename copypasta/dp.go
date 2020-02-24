package copypasta

import "sort"

/*
参考书籍推荐：《算法竞赛进阶指南》

基础 DP 视频讲解：
https://www.bilibili.com/video/av70148899 DP 入门，01 背包，完全背包，多重背包
https://www.bilibili.com/video/av77393700 LCS LIS
https://www.bilibili.com/video/av83939419 区间 DP
https://www.bilibili.com/video/av85636122 动态规划 · 零 - Introduction
https://www.bilibili.com/video/av86983419 动态规划 · 一 - 序列型
https://www.bilibili.com/video/av89052674 动态规划 · 二 - 坐标、双序列、划分 & 状态压缩

基础 DP 题目推荐/类型总结：
https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
https://codeforces.com/blog/entry/70018
https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md
https://zxi.mytechroad.com/blog/leetcode-problem-categories/

其他资料：
https://codeforces.com/blog/entry/45223 SOS Dynamic Programming
https://github.com/hzwer/shareOI/tree/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92
https://oi-wiki.org/dp/

线性 DP 经典题：数字三角形 https://www.luogu.com.cn/problem/P1216

NOTE: 若使用滚动数组，复用时可能要初始化
NOTE: 实际情况是使用滚动数组仅降低了内存开销，算法运行效率与不使用滚动数组时无异

记忆化耗时大约是递推的 6 倍？
*/
func dpCollections() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 由于数据范围的原因，采用 map 记忆化
	generalDPMap := func() {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(int, int) int
		f = func(x, y int) (_ans int) {
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = _ans }()

			return
		}
		_ = f
	}

	// O(nlogn) LIS
	// https://oi-wiki.org/dp/basic/#_12
	lis := func(arr []int) int {
		dp := make([]int, 0, len(arr))
		for _, v := range arr {
			if i := sort.SearchInts(dp, v); i < len(dp) {
				dp[i] = v
			} else {
				dp = append(dp, v)
			}
		}
		return len(dp)
	}

	// 无限物品：恰好装满背包至少需要多少个物品
	// 无法装满返回 -1
	coinChange := func(coins []int, amount int) int {
		const inf int = 1e9
		dp := make([]int, amount+1)
		for i := range dp {
			dp[i] = inf
		}
		dp[0] = 0
		for cur := range dp {
			for _, c := range coins {
				if c <= cur {
					dp[cur] = min(dp[cur], dp[cur-c]+1)
				}
			}
		}
		if dp[amount] < inf {
			return dp[amount]
		}
		return -1
	}

	/* 背包问题
	https://en.wikipedia.org/wiki/Knapsack_problem
	*/

	// 01背包
	// https://oi-wiki.org/dp/knapsack/
	knapsack01 := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij
				} else {
					dp[i+1][j] = max(dpij, dp[i][j-wi]+vi)
				}
			}
		}
		return dp[n][maxW]
	}

	// TODO: 单调队列/单调栈优化
	// https://oi-wiki.org/dp/opt/monotonous-queue-stack/

	// TODO: 斜率优化
	// https://oi-wiki.org/dp/opt/slope/

	// TODO: 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/

	// 树上最大匹配
	// https://codeforces.com/blog/entry/2059
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	maxMatchingOnTree := func(n int, g [][]int) int {
		cover, nonCover := make([]int, n), make([]int, n)
		var f func(int, int)
		f = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
					nonCover[v] += max(cover[w], nonCover[w])
				}
			}
			for _, w := range g[v] {
				cover[v] = max(cover[v], 1+nonCover[w]+nonCover[v]-max(cover[w], nonCover[w]))
			}
		}
		f(0, -1)
		return max(cover[0], nonCover[0])
	}

	_ = []interface{}{
		generalDPMap,
		lis, coinChange,
		knapsack01,
		maxMatchingOnTree,
	}
}
