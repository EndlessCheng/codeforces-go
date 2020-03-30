package copypasta

import (
	"sort"
)

/*
参考书籍推荐：
《算法竞赛进阶指南》- 介绍了大量且全面的 DP 内容，目前市面上讲解 DP 最好的一本书

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
重要技巧总结+套题！ https://codeforces.com/blog/entry/47764
AT 经典 DP 场 https://atcoder.jp/contests/dp 题解 https://www.cnblogs.com/shanxieng/p/10232228.html https://www.hamayanhamayan.com/entry/2019/01/12/163853

其他资料：
https://github.com/hzwer/shareOI/tree/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92
https://oi-wiki.org/dp/
https://cp-algorithms.com/dynamic_programming/divide-and-conquer-dp.html

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
	mapDP := func() {
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

	/* 线性 DP：前缀/后缀之间的转移
	数字三角形 https://www.luogu.com.cn/problem/P1216
	最长公共子序列 (LCS) https://leetcode-cn.com/problems/longest-common-subsequence/
	最长上升子序列 (LIS) https://leetcode-cn.com/problems/longest-increasing-subsequence/
	最长公共上升子序列 (LCIS) https://codeforces.com/problemset/problem/10/D

	两个排列的 LCS https://www.luogu.com.cn/problem/P1439
	*/

	// LCS
	// 最长公共子序列 (LCS) https://leetcode-cn.com/problems/longest-common-subsequence/
	// EXTRA: 最短公共超序列 (SCS) https://leetcode-cn.com/problems/shortest-common-supersequence/
	// todo 两个排列的 LCS https://www.luogu.com.cn/problem/P1439
	lcs := func(s1, s2 string) int {
		n, m := len(s1), len(s2)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		for i, b1 := range s1 {
			for j, b2 := range s2 {
				if b1 == b2 {
					dp[i+1][j+1] = dp[i][j] + 1
				} else {
					dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
				}
			}
		}
		return dp[n][m]
	}
	lcsPath := func(s1, s2 string) []byte {
		n, m := len(s1), len(s2)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		fa := make([][]int8, n+1)
		for i := range fa {
			fa[i] = make([]int8, m+1)
		}
		for i, b1 := range s1 {
			for j, b2 := range s2 {
				if b1 == b2 {
					dp[i+1][j+1] = dp[i][j] + 1
					fa[i+1][j+1] = 1
				} else {
					if dp[i][j+1] > dp[i+1][j] {
						dp[i+1][j+1] = dp[i][j+1]
						fa[i+1][j+1] = 2
					} else {
						dp[i+1][j+1] = dp[i+1][j]
						fa[i+1][j+1] = 3
					}
				}
			}
		}
		lcs := make([]byte, 0, dp[n][m])
		var makeLCS func(i, j int)
		makeLCS = func(i, j int) {
			if i == 0 || j == 0 {
				return
			}
			if fa[i][j] == 1 {
				makeLCS(i-1, j-1)
				lcs = append(lcs, s1[i-1])
			} else if fa[i][j] == 2 {
				makeLCS(i-1, j)
			} else {
				makeLCS(i, j-1)
			}
		}
		makeLCS(n, m)
		return lcs
	}

	// O(nlogn) LIS
	// https://oi-wiki.org/dp/basic/#_12
	// 最长上升子序列 (LIS) https://leetcode-cn.com/problems/longest-increasing-subsequence/
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

	// 本质不同子序列个数
	// 定义 dp[i][j] 表示前 i 个字符中长度为 j 的本质不同子序列个数
	// 转移 dp[i][j] = dp[i-1][j]（不选第 i 个字符）+ dp[i-1][j-1] - dp[prev[i]-1][j-1]（选第 i 个字符）
	// 其中 prev[i] 为 s[i] 的上一个相同字符位置
	// https://ac.nowcoder.com/discuss/394080 C 题
	// https://codeforces.com/problemset/problem/1183/H
	distinctSubsequence := func(s string) int64 {
		n := len(s)
		prev := [26]int{}
		dp := make([][]int64, n+1)
		for i := range dp {
			dp[i] = make([]int64, n+1)
		}
		dp[0][0] = 1
		for i := 1; i <= n; i++ {
			c := s[i-1] - 'a'
			dp[i][0] = 1
			for j := 1; j <= i; j++ {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
				if p := prev[c]; p > 0 {
					dp[i][j] -= dp[p-1][j-1]
				}
			}
			prev[c] = i
		}
		sum := int64(0)
		for _, cnt := range dp[n][1:] { // 不计入空字符串
			sum += cnt
		}
		return sum
	}

	// 无限物品：恰好装满背包至少需要多少个物品
	// 无法装满返回 -1
	minCoinChange := func(coins []int, amount int) int {
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
	https://codeforces.com/blog/entry/59606
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

	// 完全背包

	// 多重背包

	/* 区间 DP / 环形 DP：元素的位置影响结果
	最优三角剖分 https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/
	石子合并：相邻 k 堆 https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
	石子合并：环形，相邻 2 堆 https://www.luogu.com.cn/problem/P1880

	todo https://atcoder.jp/contests/abc159/tasks/abc159_f
	*/

	/* 状压 DP
	 */

	/* 数位 DP
	todo 讲解+套题 https://codeforces.com/blog/entry/53960
	入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
	入门题 https://atcoder.jp/contests/dp/tasks/dp_s
	好题 LC182D https://leetcode-cn.com/problems/find-all-good-strings/
	 */
	digitDP := func(lower, upper string) int {
		const mod int = 1e9 + 7
		n := len(upper) // len(lower) == len(upper)

		// <=s 的数目
		calc := func(s string) int {
			dp := make([][]int, n)
			for i := range dp {
				dp[i] = make([]int, n)
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, sum int, isUpper bool) int
			f = func(p, sum int, isUpper bool) (cnt int) {
				//if sum... { return 0 }
				if p >= n {
					return 1
				}
				dv := &dp[p][sum]
				if !isUpper && *dv >= 0 {
					return *dv
				}
				defer func() {
					if !isUpper {
						*dv = cnt
					}
				}()
				up := byte('9')
				if isUpper {
					up = s[p]
				}
				for digit := byte('0'); digit <= up; digit++ {
					tmp := sum
					// do tmp...
					c := f(p+1, tmp, isUpper && digit == up)
					cnt = (cnt + c) % mod
				}
				return
			}
			return f(0, 0, true)
		}
		ansLower := calc(lower)
		ansUpper := calc(upper)
		ans := ansUpper - ansLower + 1 // lower 算上
		ans = (ans%mod + mod) % mod
		return ans
	}

	// TODO: 单调队列/单调栈优化
	// https://oi-wiki.org/dp/opt/monotonous-queue-stack/

	// TODO: 斜率优化 / 凸包优化 (CHT)
	// https://oi-wiki.org/dp/opt/slope/
	// https://codeforces.com/blog/entry/63823
	// todo https://luckyglass.github.io/2019/19Dec21stArt1/

	// TODO: 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/
	// https://blog.csdn.net/weixin_43914593/article/details/105150937

	/* 树形 DP
	https://codeforces.com/blog/entry/20935
	https://codeforces.com/blog/entry/63257
	 */

	// todo 公司聚会/树上最大独立集
	// 入门经典 p.280
	// https://stackoverflow.com/questions/13544240/algorithm-to-find-max-independent-set-in-a-tree

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
		mapDP,
		lcs, lcsPath, lis, distinctSubsequence,
		minCoinChange,
		knapsack01,
		digitDP,
		maxMatchingOnTree,
	}
}
