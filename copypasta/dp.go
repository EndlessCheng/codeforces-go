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

https://oi-wiki.org/dp/

若使用滚动数组，注意在下次复用时初始化第一排所有元素
但是实际情况是使用滚动数组仅降低了内存开销，算法运行效率与不使用滚动数组时无异

记忆化耗时大约是递推的 6 倍
*/
func dpCollections() {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	generalDP := func(a []int) (ans int) {
		n := len(a)
		dp := make([]int, n) // n+1
		for i := range dp {
			dp[i] = -1
		}
		var f func(int) int
		f = func(p int) (_ans int) {
			if p < 0 || p >= n {
				return 0
			}
			if dp[p] != -1 {
				return dp[p]
			}
			defer func() { dp[p] = _ans }()
			_ = a[p]

			return
		}
		//return f(0)
		for i := range a {
			fi := f(i)
			ans = max(ans, fi)
		}
		return
	}

	generalDP2 := func(a []int) int {
		n := len(a)
		dp := make([][]int, n) // n+1
		for i := range dp {
			dp[i] = make([]int, n) // n+1
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int
		f = func(l, r int) (_ans int) {
			if l >= r {
				return 0
			}
			if dp[l][r] != -1 {
				return dp[l][r]
			}
			defer func() { dp[l][r] = _ans }()

			return
		}
		return f(0, n-1)
	}

	// 由于数据范围的原因，采用 map 记忆化
	generalDPMap := func(x, y int) int {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(int, int) int
		f = func(x, y int) (_ans int) {
			// 边界检查
			// ...
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = _ans }()

			return
		}
		return f(x, y)
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
		generalDP, generalDP2, generalDPMap,
		lis, knapsack01,
		maxMatchingOnTree,
	}
}
