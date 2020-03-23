package copypasta

import "sort"

/*
参考书籍推荐：《算法竞赛进阶指南》

DP 视频讲解：
https://www.bilibili.com/video/av70148899 DP 入门，01 背包，完全背包，多重背包
https://www.bilibili.com/video/av77393700 LCS LIS
https://www.bilibili.com/video/av83939419 区间 DP
https://www.bilibili.com/video/av93356551 状态压缩 DP
https://www.bilibili.com/video/av98090640 树形 DP
https://www.bilibili.com/video/av85636122 动态规划 · 零 - Introduction
https://www.bilibili.com/video/av86983419 动态规划 · 一 - 序列型
https://www.bilibili.com/video/av89052674 动态规划 · 二 - 坐标、双序列、划分 & 状态压缩

套题/总结：
线性 DP 和区间 DP https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
按照相似题目分类 https://zxi.mytechroad.com/blog/leetcode-problem-categories/
https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-w-5/
状压 DP https://codeforces.com/blog/entry/45223
CSES DP section editorial https://codeforces.com/blog/entry/70018
LC 全部 DP 题 https://leetcode-cn.com/tag/dynamic-programming/

其他资料：
https://github.com/hzwer/shareOI/tree/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92
https://oi-wiki.org/dp/

NOTE: 若使用滚动数组，复用时可能要初始化
NOTE: 实际情况是使用滚动数组仅降低了内存开销，整体运行效率与不使用滚动数组时无异
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

	/* 线性 DP
	数字三角形 https://www.luogu.com.cn/problem/P1216
	最长公共子序列 (LCS) https://leetcode-cn.com/problems/longest-common-subsequence/
	最长上升子序列 (LIS) https://leetcode-cn.com/problems/longest-increasing-subsequence/
	最长公共上升子序列 (LCIS) https://codeforces.com/problemset/problem/10/D

	两个排列的 LCS https://www.luogu.com.cn/problem/P1439
	*/

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

	/* 区间 DP
	最优三角剖分 https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/
	石子合并：相邻 k 堆 https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
	石子合并：环形，相邻 2 堆 https://www.luogu.com.cn/problem/P1880
	*/

	/* 状压 DP
	*/

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

	// 插头 DP / 轮廓线动态规划
	//《训练指南》6.1

	_ = []interface{}{
		generalDPMap,
		lis, coinChange,
		knapsack01,
		maxMatchingOnTree,
	}
}
