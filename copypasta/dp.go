package copypasta

import (
	"sort"
)

/* 动态规划

首先请透彻理解何为问题的「状态空间」，见 search.go 开头的注释

思考过程：
1.1 如何把问题形式化为状态空间？（可以从边界、子集的角度思考）
1.2 子问题是如何重叠的？
1.3 子问题是怎么逐层递进的？（题目描述、隐含的顺序）
2.1 如何定义状态？需要用几个维度表示？
2.2 状态的范围是多少？起点状态和终点状态是什么？
2.3 哪些状态是相邻的？（即通过一次转移就能得到）
2.4 状态转移时要计算哪些内容？
2.5 对于转移来的相邻状态（入边），怎么决策？（简单的有取最值取和，复杂的有组合决策）
3.1 若复杂度过高，如何优化决策？
*  状态不好确定时，尝试转化问题模型、逆序思考、增加维度等等
*  对于计数问题或概率问题来说，状态定义和状态转移要做到不重不漏
   推荐 https://codeforces.ml/blog/entry/47764
   戳气球 LC312 https://leetcode-cn.com/problems/burst-balloons/
   消消乐 LC546/周赛25D https://leetcode-cn.com/problems/remove-boxes/ https://leetcode.com/contest/leetcode-weekly-contest-25
   如何定义状态 https://codeforces.ml/problemset/problem/553/A
   如何定义状态 LC956/周赛114D https://leetcode-cn.com/problems/tallest-billboard/ https://leetcode-cn.com/contest/weekly-contest-114/
   谁来当 DP 对象 LC1434/双周赛25D https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/ https://leetcode-cn.com/contest/biweekly-contest-25/
   扔蛋问题 LC887/周赛97D https://leetcode-cn.com/problems/super-egg-drop/ https://www.bilibili.com/video/BV1KE41137PK https://leetcode-cn.com/contest/weekly-contest-97/
   LC920* https://leetcode-cn.com/problems/number-of-music-playlists/ 注：官方题解给出了一种生成函数的做法
   状态优化 https://codeforces.ml/problemset/problem/838/E

NOTE: 若使用滚动数组，注意复用时可能要初始化
NOTE:（区间 DP）正向计算不易时，试着反向计算
TIPS: 若转移是若干相邻项之和，可以考虑 f(p) - f(p-1) 的值，用滑动窗口来维护区间和，从而优化转移
      例题 LC837 https://leetcode-cn.com/problems/new-21-game/

参考书籍推荐：
《算法竞赛进阶指南》- 介绍了大量且全面的 DP 内容，是目前市面上讲解 DP 最好的一本书

视频讲解：
https://www.bilibili.com/video/av70148899 DP 入门，01 背包，完全背包，多重背包
https://www.bilibili.com/video/av77393700 LCS LIS
https://www.bilibili.com/video/av83939419 区间 DP
https://www.bilibili.com/video/av93356551 状态压缩 DP
https://www.bilibili.com/video/av98090640 树形 DP
https://www.bilibili.com/video/BV1MT4y1376C 数位 DP
https://www.bilibili.com/video/av85636122 动态规划 · 零 - Introduction
https://www.bilibili.com/video/av86983419 动态规划 · 一 - 序列型
https://www.bilibili.com/video/av89052674 动态规划 · 二 - 坐标、双序列、划分 & 状态压缩

套题/总结：
Non-trivial DP Tricks and Techniques https://codeforces.ml/blog/entry/47764
SOS Dynamic Programming https://codeforces.ml/blog/entry/45223
CSES DP section editorial https://codeforces.ml/blog/entry/70018
CF 全部 DP 题  https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp
力扣上的 DP 问题
    分类汇总 https://zhuanlan.zhihu.com/p/126546914
    https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
    https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md
    https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems
    https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-w-5/
    https://leetcode-cn.com/tag/dynamic-programming/
AT 经典 DP 场 https://atcoder.jp/contests/dp
    题解 https://www.cnblogs.com/shanxieng/p/10232228.html
    题解（日语）https://www.hamayanhamayan.com/entry/2019/01/12/163853
信息学奥赛一本通 第二部分 基础算法 --> 第九章 动态规划 http://ybt.ssoier.cn:8088/index.php
算法竞赛专题解析（11）：DP概述和常见DP面试题 https://blog.csdn.net/weixin_43914593/article/details/105444090

其他资料：
https://github.com/hzwer/shareOI/tree/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92
https://oi-wiki.org/dp/
https://cp-algorithms.com/dynamic_programming/divide-and-conquer-dp.html
https://wenku.baidu.com/view/7c9de809581b6bd97f19ea72.html 算法合集之《从《鹰蛋》一题浅析对动态规划算法的优化》
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
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	// 涉及到前缀和/子数组和的问题
	// 定义 dp[i] 表示前缀 a[:i] 中子数组和为 targetSum 的最短子数组长度
	// 下面的代码来自 LC1477/双周赛28C https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
	prefixSumDP := func(a []int, targetSum int) int {
		n := len(a)
		const inf int = 1e9

		ans := inf
		dp := make([]int, n+1)
		for _i := range dp {
			dp[_i] = inf
		}
		preSumPos := map[int]int{0: -1} // int64
		sum := 0
		for i, v := range a {
			dp[i+1] = dp[i]
			sum += v
			if p, ok := preSumPos[sum-targetSum]; ok {
				// sum_[p+1,i] == targetSum
				l := i - p
				if dp[p+1] < inf {
					ans = min(ans, dp[p+1]+l)
				}
				dp[i+1] = min(dp[i+1], l)
			}
			preSumPos[sum] = i
		}
		if ans == inf {
			ans = -1
		}
		return ans
	}

	// https://codeforces.ml/problemset/problem/510/D
	// 由于数据范围的原因，采用 map 记忆化
	mapDP := func() {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(int, int) int
		f = func(x, y int) (res int) {
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = res }()

			return
		}
		f(0, 0)
	}

	/* 线性 DP
	① 前缀/后缀之间的转移，例如从 dp[i-1] 转移到 dp[i]，或者从 dp[j] 转移到 dp[i] (j<i)，这里 dp[i] 可以表示一个状态或一组状态等
	力扣上有大量这类题目，例如：
	198,213,123,309,376,276,931 (从dp[i-1] 转移到 dp[i])
	487,1186 (从 dp[i-1] 转移到 dp[i]，带一个额外的决策维度，长度一般是 2-4)
	300,368,1105* (从 dp[j] 转移到 dp[i])
	903/周赛101D https://leetcode-cn.com/problems/valid-permutations-for-di-sequence/ https://leetcode-cn.com/contest/weekly-contest-101/
	② 双序列问题，一般定义 dp[i][j] 表示对子问题 (s1[:i],s2[:j]) 的求解结果
	力扣题目 1143,1092,72,97,115,727,583,712,1035,1216,1312
	983/周赛121C https://leetcode-cn.com/problems/minimum-cost-for-tickets/ https://leetcode-cn.com/contest/weekly-contest-121/
	③ 一些题目
	最大整除子集 LC368 https://leetcode-cn.com/problems/largest-divisible-subset/
	编辑距离 LC72 https://leetcode-cn.com/problems/edit-distance/
	最高的广告牌 LC956/周赛114D https://leetcode-cn.com/problems/tallest-billboard/ https://leetcode-cn.com/contest/weekly-contest-114/
	数字三角形 https://www.luogu.com.cn/problem/P1216
	todo 最长公共上升子序列 (LCIS) https://codeforces.ml/problemset/problem/10/D
	todo 两个排列的 LCS https://www.luogu.com.cn/problem/P1439
	贪心+abs https://atcoder.jp/contests/abc163/tasks/abc163_e
	LC1477/双周赛28C https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
	*/

	// 最大子段和
	// 算法导论 练习4.1-5
	// 变体 https://codeforces.com/problemset/problem/1373/D
	maxSubArraySum := func(a []int) int {
		if len(a) == 0 {
			return 0
		}
		curSum, maxSum := a[0], a[0]
		for _, v := range a[1:] {
			curSum = max(curSum+v, v)
			maxSum = max(maxSum, curSum)
		}
		return max(maxSum, 0) // 若不允许非空，返回 maxSum
	}

	maxSubArrayAbsSum := func(a []int) int {
		if len(a) == 0 {
			return 0
		}
		//min, max, abs := math.Min, math.Max, math.Abs
		curMaxSum, maxSum := a[0], a[0]
		curMinSum, minSum := a[0], a[0]
		for _, v := range a[1:] {
			curMaxSum = max(curMaxSum+v, v)
			maxSum = max(maxSum, curMaxSum)
			curMinSum = min(curMinSum+v, v)
			minSum = min(minSum, curMinSum)
		}
		return max(abs(maxSum), abs(minSum))
	}

	// 最长公共子序列 (LCS)
	// 有向无环图：s1[i] == s2[j] (i-1,j-1) -> (i,j) $ 1
	//           s1[i] != s2[j] (i-1,j) -> (i,j) $ 0
	//                          (i,j-1) -> (i,j) $ 0
	// 例题 LC1143 https://leetcode-cn.com/problems/longest-common-subsequence/
	// EXTRA: 最短公共超序列 (SCS) LC1092 https://leetcode-cn.com/problems/shortest-common-supersequence/
	// 变种 LC97   https://leetcode-cn.com/problems/interleaving-string/
	//     LC115  https://leetcode-cn.com/problems/distinct-subsequences/
	//     LC583  https://leetcode-cn.com/problems/delete-operation-for-two-strings/
	//     LC712  https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/
	//     LC1035 https://leetcode-cn.com/problems/uncrossed-lines/
	//     LC1312 https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
	lcs := func(s1, s2 string) int {
		n, m := len(s1), len(s2)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		for i, b1 := range s1 {
			for j, b2 := range s2 {
				if b1 == b2 {
					// ignore values from dp[i][j+1] and dp[i+1][j]
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

	// 最长上升子序列 (LIS)
	// O(n^2) - 定义 dp[i] 为以 a[i] 为末尾的 LIS 的长度
	//          可以把此问题想象成一个「跳跃游戏」，任选一个初始位置向右跳跃，每次只能跳到比当前位置更高的位置，问最多能跳多少次（最后答案加一）
	//          这样能更容易地看出转移的顺序，然后变成一个 DAG 上求最长路的问题
	// 变体 https://codeforces.ml/problemset/problem/1350/B
	lisSlow := func(a []int) (ans int) {
		n := len(a)
		dp := make([]int, n)
		for i, v := range a {
			dp[i] = 1
			for j, w := range a[:i] {
				if w < v { // 改成 <= 为非降
					dp[i] = max(dp[i], dp[j]+1)
				}
			}
			ans = max(ans, dp[i])
		}
		return
	}

	// 最长上升子序列 (LIS)
	// O(nlogn) - 定义 dp[i] 为长度为 i+1 的 LIS 末尾元素的最小值
	// 求下降，可以考虑把序列元素去相反数
	// https://oi-wiki.org/dp/basic/#_12
	// 例题 LC300 https://leetcode-cn.com/problems/longest-increasing-subsequence/
	// 建模 https://codeforces.ml/problemset/problem/269/B
	// 方案数 LC673 https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
	//       https://www.zhihu.com/question/34905638
	// LC354 俄罗斯套娃信封问题 https://leetcode-cn.com/problems/russian-doll-envelopes/
	lis := func(a []int) int {
		n := len(a)
		dp := make([]int, 0, n)
		for _, v := range a {
			if i := sort.SearchInts(dp, v); i < len(dp) { // 改成 v+1 为非降
				dp[i] = v
			} else {
				dp = append(dp, v)
			}
		}
		return len(dp)
	}

	// LIS 相关构造题
	// https://codeforces.ml/problemset/problem/1304/D
	// https://atcoder.jp/contests/arc091/tasks/arc091_c

	// 本质不同子序列个数
	// 定义 dp[i][j] 表示前 i 个字符中长度为 j 的本质不同子序列个数
	// 转移 dp[i][j] = dp[i-1][j]（不选第 i 个字符）+ dp[i-1][j-1] - dp[prev[i]-1][j-1]（选第 i 个字符）
	// 其中 prev[i] 为 s[i] 的上一个相同字符位置
	// https://ac.nowcoder.com/acm/contest/4853/C 题解 https://ac.nowcoder.com/discuss/394080
	// https://codeforces.ml/problemset/problem/1183/H
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

	/* 背包问题
	这类问题可以从物品选择次序的无后效性入手
	子区间 -> 前缀和
	子序列 -> 背包
	https://en.wikipedia.org/wiki/Knapsack_problem
	https://codeforces.ml/blog/entry/59606
	NOTE: 若求能否凑成 1,2,3,...,M，只需判断 dp[i] 是否为正 https://leetcode-cn.com/problems/last-stone-weight-ii/
	套题 https://www.acwing.com/problem/
	*/

	// 0-1 背包 (n 个物品，背包容量为 maxW)
	// 基本状态：前 i 个物品  i∈[0,n]
	// 附加状态：容量(上限)为 j  j∈[0,maxW]
	// 点权：最大价值
	//     初始值：(0,j)=0  j∈[0,maxW]
	// 有向无环图：不选第 i 个物品，对各个容量 j，连一条横边，即 (i-1,j) -> (i,j) $ 0
	//             选第 i 个物品，对各个容量 j (j≥wi)，连边 (i-1,j-wi) -> (i,j) $ vi
	//     起点：(0,j)  j∈[0,maxW]
	//     终点：(n,maxW)
	// 核心函数：最大价值（最长路），即 max
	// https://oi-wiki.org/dp/knapsack/
	// 模板题 https://atcoder.jp/contests/dp/tasks/dp_d
	// 建模转换 LC1049 https://leetcode-cn.com/problems/last-stone-weight-ii/
	// EXTRA: 能否恰好装满（其实就是 waysToSum 的方案数是否为 0）LC416 https://leetcode-cn.com/problems/partition-equal-subset-sum/
	// EXTRA: 二维费用 LC474 https://leetcode-cn.com/problems/ones-and-zeroes/
	zeroOneKnapsack := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij // 入度为 1，直接转移
				} else {
					dp[i+1][j] = max(dpij, dp[i][j-wi]+vi) // 入度为 2，取核心函数转移
				}
			}
		}
		return dp[n][maxW]
	}

	// todo 价值主导的 0-1 背包

	// 从 a 中选出若干个数，总和为 sum 的方案数
	// 基本状态：前 i 个数  i∈[0,n]
	// 附加状态：和为 j  j∈[0,sum]
	// 点权：方案数
	//     初始值：(0,0)=1
	// 有向无环图：不选第 i 个数，对各个和 j，连一条横边，即 (i-1,j) -> (i,j)
	//             选第 i 个数，对各个和 j (j≥ai)，连边 (i-1,j-ai) -> (i,j)
	//     起点：(0,j)  j∈[0,sum]
	//     终点：(n,sum)
	// 核心函数：方案数（点权汇合），即 +
	// 例题 LC879 https://leetcode-cn.com/problems/profitable-schemes/
	// 例题（需要转换）LC494 https://leetcode-cn.com/problems/target-sum/
	// 隐藏的 0-1 背包 LC1434 https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/
	// 建模转换 https://atcoder.jp/contests/abc169/tasks/abc169_f
	waysToSum := func(a []int, sum int) int {
		n := len(a)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, sum+1)
		}
		dp[0][0] = 1
		for i, v := range a {
			for s, dv := range dp[i] {
				dp[i+1][s] = dv
				if s >= v {
					dp[i+1][s] += dp[i][s-v]
				}
			}
		}
		return dp[n][sum]
	}

	// 完全背包
	unboundedKnapsack := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij // 入度为 1，直接转移
				} else {
					dp[i+1][j] = max(dpij, dp[i+1][j-wi]+vi) // 入度为 2，取核心函数转移
				}
			}
		}
		return dp[n][maxW]
	}

	// 恰好装满背包至少需要多少个物品，物品无限。无法装满时返回 -1
	// 基本状态：容量 i  i∈[0,amount]
	// 点权：最少物品数
	//     初始值：0=0
	// 有向无环图：i-wj (wj≤i) -> i $ 1
	//     起点：0
	//     终点：amount
	// 核心函数：最少物品数（最短路），即 min
	// LC322 https://leetcode-cn.com/problems/coin-change/
	minCoinChange := func(coins []int, amount int) int {
		const inf int = 1e9
		dp := make([]int, amount+1)
		for i := range dp {
			dp[i] = inf
		}
		dp[0] = 0
		// 按容量遍历以满足拓扑序
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

	// EXTRA: 完全背包 - 求方案数
	// LC518 https://leetcode-cn.com/problems/coin-change-2/

	// EXTRA: 二维费用完全背包 - 求方案数
	// 注意：「恰好使用 m 个物品」这个条件要当成一种费用来看待
	// https://codeforces.ml/problemset/problem/543/A

	// 多重背包 - 未优化
	// 模板题 https://codeforces.ml/problemset/problem/106/C
	boundedKnapsack := func(values, stocks, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			si, wi := stocks[i], weights[i]
			for j := range dp[i] {
				for k := 0; k <= si && k*wi <= j; k++ {
					dp[i+1][j] = max(dp[i+1][j], dp[i][j-k*wi]+k*vi)
				}
			}
		}
		return dp[n][maxW]
	}

	// todo 多重背包 - 优化 1 - 二进制优化

	/* 区间 DP / 环形 DP
	一般来说转移是合并区间或者分解区间
	① 将序列分成 K 个连续区间，求解这些区间的某个最优性质
	一般定义 dp[i][k] 表示将 a[:i] 分成 k 个连续区间得到的最优解
	此时可以枚举最后一个区间的左端点 j，从 dp[j-1][k-1] 转移到 dp[i][k]，转移时考虑 a[j:i] 对最优解的影响
	力扣题目 1278,813,410,1335
	② 求解关于某个序列的最优性质，要求大区间的最优解可以依赖于小区间的最优解
	一般定义 dp[i][j] 表示 a[i:j] 的最优解
	此时可以枚举区间大小和区间左端点，从小区间转移到大区间
	力扣题目 516,312,375,1246
	移除盒子 LC546/周赛25D https://leetcode-cn.com/problems/remove-boxes/ https://leetcode.com/contest/leetcode-weekly-contest-25
	③ 一些题目
	https://blog.csdn.net/weixin_43914593/article/details/106163859 算法竞赛专题解析（14）：DP应用--区间DP
	最优三角剖分 LC1039 https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/
	戳气球 LC312 https://leetcode-cn.com/problems/burst-balloons/
	打印机 LC664 https://leetcode-cn.com/problems/strange-printer/
	安排邮筒 LC1478/双周赛28D https://leetcode-cn.com/problems/allocate-mailboxes/
	todo https://atcoder.jp/contests/abc159/tasks/abc159_f
	     https://codeforces.ml/problemset/problem/245/H
	*/

	// 石子合并
	// https://ac.nowcoder.com/acm/contest/1043/A https://ac.nowcoder.com/acm/problem/51170
	// 环形的情况 https://www.luogu.com.cn/problem/P1880
	// 相邻 k 堆的情况（综合①②）LC1000 https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
	mergeStones := func(a []int) int {
		n := len(a)
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = 1e9
			}
			dp[i][i] = 0
		}
		for sz := 2; sz <= n; sz++ {
			for l := 0; l+sz <= n; l++ {
				r := l + sz - 1
				for i := l; i < r; i++ {
					dp[l][r] = min(dp[l][r], dp[l][i]+dp[i+1][r])
				}
				dp[l][r] += sum[r+1] - sum[l]
			}
		}
		return dp[0][n-1]
	}

	/* 博弈类 DP
	转移：让「自己与对手的分差」最大
	LC877 https://leetcode-cn.com/problems/stone-game/ https://nanti.jisuanke.com/t/48
	LC1140 https://leetcode-cn.com/problems/stone-game-ii/
	LC1406 https://leetcode-cn.com/problems/stone-game-iii/
	CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp%2Cgames
	*/

	/* 状压 DP
	NOTE: 若问题无法划分成小问题，必须考虑各种可能的情况，则可能是 NP 完全问题
	浅谈状压 DP https://www.luogu.com.cn/blog/yijan/zhuang-ya-dp
	https://blog.csdn.net/weixin_43914593/article/details/106432695 算法竞赛专题解析（15）：DP应用--状态压缩DP
	CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp%2Cbitmasks
	todo LC691  https://leetcode-cn.com/problems/stickers-to-spell-word/
	     LC1125 https://leetcode-cn.com/problems/smallest-sufficient-team/
	     LC943  https://leetcode-cn.com/problems/find-the-shortest-superstring/

	枚举子集，复杂度 O(3^n)
		dp[s] = min(valid dp[s^sub]) + 1
		todo LC1494/双周赛29D https://leetcode-cn.com/problems/parallel-courses-ii/
	todo 汉密尔顿路径/回路 Hamiltonian path
	*/

	// 旅行商问题 (TSP)
	// 返回一个 ans 数组，ans[i] 表示从 st 出发，访问完所有位置且最后停在 i 的最短路径（注意可能要特判 i==st 的情况）
	// https://en.wikipedia.org/wiki/Travelling_salesman_problem
	// 模板题 https://www.luogu.com.cn/problem/P1171 https://www.luogu.com.cn/problem/P1433
	// 建模转换题 LC943 https://leetcode-cn.com/problems/find-the-shortest-superstring/
	//          LCP13 https://leetcode-cn.com/problems/xun-bao/
	// EXTRA: 固定起点终点的问题，视问题情况有两种方法：
	//        添加一个节点 https://stackoverflow.com/questions/14527815/how-to-fix-the-start-and-end-points-in-travelling-salesmen-problem
	//        设置距离 https://stackoverflow.com/questions/36086406/traveling-salesman-tsp-with-set-start-and-end-point
	tsp := func(dist [][]int, st int) []int {
		// 记忆化：已经访问的集合 s，当前位置 v
		n := len(dist)
		const inf int = 1e9 // 1e18
		dp := make([][]int, 1<<n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = inf
			}
		}
		dp[1<<n-1][st] = 0 // 访问了所有节点并回到了 st（多个起点的话就设置多个 dp[1<<n-1][st[i]] = 0）
		for s := 1<<n - 2; s >= 0; s-- {
			for v := 0; v < n; v++ {
				for w := 0; w < n; w++ {
					if s>>w&1 == 0 {
						dp[s][v] = min(dp[s][v], dp[s|1<<w][w]+dist[v][w])
					}
				}
			}
		}
		return dp[0]
	}

	{
		// 由于 s 的特性，在单起点的情况下，有很多状态是没有访问到的
		_ = func(dist [][]int, st int) int {
			n := len(dist)
			dp := make([][]int, 1<<n)
			for i := range dp {
				dp[i] = make([]int, n)
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			const inf int = 1e9 // 1e18
			// 记忆化：已经访问的集合 s，当前位置 v
			var f func(s, v int) int
			f = func(s, v int) (res int) {
				dv := &dp[s][v]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
				if s == 1<<n-1 && v == st {
					return
				} // 访问了所有节点并回到了 st
				res = inf
				for w := 0; w < n; w++ {
					if s>>w&1 == 0 {
						res = min(res, f(s|1<<w, w)+dist[v][w])
					}
				}
				return
			}
			return f(0, st)
		}
	}

	/* 插头 DP / 轮廓线 DP / Broken Profile DP
	《训练指南》6.1
	https://oi-wiki.org/dp/plug/
	https://cp-algorithms.com/dynamic_programming/profile-dynamics.html
	https://www.luogu.com.cn/blog/efforts-will-pay-off/du-liu-dong-gui-cha-tou-dp
	*/

	/* 数位 DP
	入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
	      https://atcoder.jp/contests/dp/tasks/dp_s
	      https://codeforces.com/problemset/problem/1036/C
	LC233 https://leetcode-cn.com/problems/number-of-digit-one/
	LC902 https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/
	好题 LC182D https://leetcode-cn.com/problems/find-all-good-strings/
	todo 套题 https://codeforces.ml/blog/entry/53960
	*/
	digitDP := func(lower, upper string) int {
		const mod int = 1e9 + 7

		// <=s 的符合要求的字符串数目
		calc := func(s string) int {
			// 有些题 lowerC 要从 1 开始，而 0 的部分单独计算（由于 0 后面可以填所有数字，这部分可以用 ∑_p>0 f(p, false) 来算）
			const lowerC, upperC byte = '0', '9'
			n := len(s)
			sumUpper := n
			dp := make([][]int, n)
			for i := range dp {
				dp[i] = make([]int, sumUpper+1)
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, sum int, isUpper bool) int
			f = func(p, sum int, isUpper bool) (cnt int) {
				//if sum... { return 0 }
				if p >= n {
					return 1 // 0
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
				up := upperC
				if isUpper {
					up = s[p]
				}
				for digit := lowerC; digit <= up; digit++ {
					tmp := sum
					// do tmp...
					c := f(p+1, tmp, isUpper && digit == up)
					// do c...
					cnt = (cnt + c) % mod
				}
				return
			}
			return f(0, 0, true)
		}
		ansLower := calc(lower) // lower-1
		ansUpper := calc(upper)
		ans := ansUpper - ansLower
		// lower 是否算上
		//if lowerIsValid {
		//	ans++
		//}
		ans = (ans%mod + mod) % mod
		return ans
	}

	/* 数据结构优化 DP
	https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=data+structures%2Cdp
	*/

	// 单调队列/单调栈优化
	// https://oi-wiki.org/dp/opt/monotonous-queue-stack/
	// https://blog.csdn.net/weixin_43914593/article/details/105791217 算法竞赛专题解析（13）：DP优化(3)--单调队列优化

	// 斜率优化 / 凸包优化 (CHT)  李超树
	// https://oi-wiki.org/dp/opt/slope/
	// https://cp-algorithms.com/geometry/convex_hull_trick.html
	// https://codeforces.ml/blog/entry/63823
	// todo https://blog.csdn.net/weixin_43914593/article/details/105560357 算法竞赛专题解析（12）：DP优化(2)--斜率(凸壳)优化
	// todo https://luckyglass.github.io/2019/19Dec21stArt1/
	//      浅谈斜率优化 https://www.luogu.com.cn/blog/duyi/xie-lv-you-hua
	// 一类单调问题的求解(宋新波) http://www.doc88.com/p-2953873379975.html
	// 题目 https://qiita.com/drken/items/9b311d553aa434bb26e4#%E4%BE%8B%E9%A1%8C-4-4-4k-anonymous-sequence-poj-no3709

	// 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/
	// todo https://blog.csdn.net/weixin_43914593/article/details/105150937 算法竞赛专题解析（10）：DP优化(1)--四边形不等式
	//      决策单调性优化讲解 https://www.luogu.com.cn/blog/83547/zong-dong-tai-gui-hua-di-ben-zhi-kan-si-bian-xing-fou-deng-shi-you-hua
	// 扔蛋问题 LC887 https://leetcode-cn.com/problems/super-egg-drop/

	/* 树形 DP
	https://codeforces.ml/blog/entry/20935
	https://codeforces.ml/blog/entry/63257
	CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp%2Ctrees
	https://codeforces.ml/problemset/problem/982/C
	*/

	// todo 公司聚会/树上最大独立集
	// 进阶指南 p.289-290
	// https://stackoverflow.com/questions/13544240/algorithm-to-find-max-independent-set-in-a-tree

	// 树上最大匹配
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	// https://codeforces.ml/blog/entry/2059
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

	// 换根 DP
	// 进阶指南 p.292-295
	// https://codeforces.ml/blog/entry/20935
	// 例题 https://codeforces.ml/problemset/problem/219/D
	// 下面的代码来自 http://poj.org/problem?id=3585
	rerootDP := func(n int) {
		type edge struct{ to, cap int }
		g := make([][]edge, n)
		// read...

		subCap := make([]int, n)
		var f func(v, fa int) int
		f = func(v, fa int) (c int) {
			for _, e := range g[v] {
				if w := e.to; w != fa {
					if len(g[w]) == 1 {
						c += e.cap
					} else {
						c += min(e.cap, f(w, v))
					}
				}
			}
			subCap[v] = c
			return
		}
		f(0, -1)

		ans := make([]int, n)
		var reroot func(v, fa, ansV int)
		reroot = func(v, fa, ansV int) {
			ans[v] = ansV
			for _, e := range g[v] {
				if w, c := e.to, e.cap; w != fa {
					if sc := subCap[w]; len(g[v]) == 1 {
						reroot(w, v, sc+c)
					} else {
						reroot(w, v, sc+min(c, ansV-min(sc, c)))
					}
				}
			}
		}
		reroot(0, -1, subCap[0])
	}

	_ = []interface{}{
		prefixSumDP, mapDP,
		maxSubArraySum, maxSubArrayAbsSum,
		lcs, lcsPath, lisSlow, lis, distinctSubsequence,
		zeroOneKnapsack, waysToSum, unboundedKnapsack, minCoinChange, boundedKnapsack,
		mergeStones,
		tsp,
		digitDP,
		maxMatchingOnTree, rerootDP,
	}
}
