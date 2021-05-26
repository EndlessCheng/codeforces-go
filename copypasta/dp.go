package copypasta

import (
	"container/heap"
	"math/bits"
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
   如何定义状态：
      https://codeforces.com/problemset/problem/553/A
      https://codeforces.com/problemset/problem/687/C
      LC956/周赛114D https://leetcode-cn.com/problems/tallest-billboard/ https://leetcode-cn.com/contest/weekly-contest-114/
      涉及到相邻状态先后关系的 DP（喂兔子）https://codeforces.com/problemset/problem/358/D
      戳气球 LC312 https://leetcode-cn.com/problems/burst-balloons/
      消消乐 LC546/周赛25D https://leetcode-cn.com/problems/remove-boxes/ https://leetcode.com/contest/leetcode-weekly-contest-25
   谁来当 DP 对象 LC1434/双周赛25D https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/ https://leetcode-cn.com/contest/biweekly-contest-25/
   扔蛋问题 LC887/周赛97D https://leetcode-cn.com/problems/super-egg-drop/ https://www.bilibili.com/video/BV1KE41137PK https://leetcode-cn.com/contest/weekly-contest-97/
   LC920* https://leetcode-cn.com/problems/number-of-music-playlists/ 注：官方题解给出了一种生成函数的做法
   状态优化 https://codeforces.com/problemset/problem/838/E
  「排序」题的转换 https://codeforces.com/problemset/problem/1223/D

NOTE: 若状态转移不构成 DAG，请尝试建图+BFS，见：
	https://ac.nowcoder.com/acm/contest/6218/B
	https://codeforces.com/problemset/problem/283/B 活用 012 染色
NOTE: 若使用滚动数组，注意复用时可能要初始化
NOTE:（区间 DP）正向计算不易时，试着反向计算
TIPS: 若转移是若干相邻项之和，可以考虑 f(p) - f(p-1) 的值，用滑动窗口来维护区间和，从而优化转移
      例题 LC837 https://leetcode-cn.com/problems/new-21-game/
递归打印路径：https://codeforces.com/problemset/problem/2/B
需要补充额外的状态 https://codeforces.com/problemset/problem/682/D

参考书籍推荐：
《算法竞赛进阶指南》- 介绍了大量且全面的 DP 内容，是目前市面上讲解 DP 最好的一本书

视频讲解：
https://www.bilibili.com/video/BV1gf4y1i78H 动态规划的套路 by wisdompeak
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
《挑战程序设计竞赛》上的练习题（均为 POJ）
2.3 节
3176 https://www.luogu.com.cn/problem/P1216 数字三角形
2229 https://www.luogu.com.cn/problem/P6065 将 n 分拆为若干个 2 的次幂的和的方法数 https://oeis.org/A018819
2385 https://www.luogu.com.cn/problem/P2690 dp[i分钟][j移动次数] = max(dp[i-1][j], dp[i-1][j-1]) + 当前分钟是否有苹果落在 j 次移动后的位置   最后答案为 max{dp[n-1]}
3616 https://www.luogu.com.cn/problem/P2889 DAG 最长路
3280 https://www.luogu.com.cn/problem/P2890 增删取 min，跑区间 DP
1742 http://acm.hdu.edu.cn/showproblem.php?pid=2844 多重背包
3046 http://poj.org/problem?id=3046 todo
3181 https://www.luogu.com.cn/problem/P6205 完全背包
1065 http://acm.hdu.edu.cn/showproblem.php?pid=1051 n 轮 LIS
1631 http://acm.hdu.edu.cn/showproblem.php?pid=1950 转换成 LIS
3666 https://www.luogu.com.cn/problem/P2893
     https://codeforces.com/problemset/problem/13/C
     https://codeforces.com/problemset/problem/713/C
     https://www.luogu.com.cn/problem/P4597 加强版
2392 https://www.luogu.com.cn/problem/P6771 多重背包，按高度限制排序。高度既是价值也是体积
2184 https://www.luogu.com.cn/problem/P2340 把 IQ 看成体积，EQ 看成价值，注意把负数偏移到非负数，以及负数的转移写法
todo 3.4 节
2686 https://www.luogu.com.cn/problem/SP1700
1769 https://www.luogu.com.cn/problem/SP90 https://www.luogu.com.cn/problem/UVA1322
2441
3254 https://www.luogu.com.cn/problem/P1879
2836
1795 https://www.luogu.com.cn/problem/SP1776
3411 https://www.luogu.com.cn/problem/SP3953
3420
3735
3171 https://www.luogu.com.cn/problem/P4644 见 graph.shortestPathDijkstra
Non-trivial DP Tricks and Techniques https://codeforces.com/blog/entry/47764
SOS Dynamic Programming https://codeforces.com/blog/entry/45223
CSES DP section editorial https://codeforces.com/blog/entry/70018
CF 全部 DP 题  https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp
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
todo 题目推荐 https://www.luogu.com.cn/blog/wyy2020/dp-qian-tan
  https://www.cnblogs.com/flashhu/p/9480669.html

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

	// https://codeforces.com/problemset/problem/510/D
	// 由于数据范围的原因，采用 map 记忆化         dpMap
	mapDP := func(n int) {
		{
			dp := map[int]int{}
			var f func(int) int
			f = func(x int) (res int) {
				//if x == 0 {
				//	return
				//}
				if v, ok := dp[x]; ok {
					return v
				}
				defer func() { dp[x] = res }()

				return
			}
			f(n)
		}

		{
			type pair struct{ x, y int }
			dp := map[pair]int{}
			var f func(int, int) int
			f = func(x, y int) (res int) {
				//if x == n {
				//  return
				//}
				p := pair{x, y}
				if v, ok := dp[p]; ok {
					return v
				}
				defer func() { dp[p] = res }()

				return
			}
			f(0, 0)
		}
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
	双周赛38D https://leetcode-cn.com/contest/biweekly-contest-38/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
	③ 一些题目
	最大整除子集 LC368 https://leetcode-cn.com/problems/largest-divisible-subset/
	编辑距离 LC72 https://leetcode-cn.com/problems/edit-distance/
	最高的广告牌 LC956/周赛114D https://leetcode-cn.com/problems/tallest-billboard/ https://leetcode-cn.com/contest/weekly-contest-114/
	数字三角形 https://www.luogu.com.cn/problem/P1216
	贪心+abs https://atcoder.jp/contests/abc163/tasks/abc163_e
	LC1477/双周赛28C https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
	看起来是区间 DP，仔细分析后是线性 DP https://leetcode-cn.com/contest/weekly-contest-199/problems/string-compression-ii/
	好题：涉及到相邻状态先后关系的 DP（喂兔子） https://codeforces.com/problemset/problem/358/D
	期望 DP https://codeforces.com/problemset/problem/235/B
	期望 DP https://codeforces.com/problemset/problem/1097/D
	*/

	// 最大子段和 https://www.luogu.com.cn/problem/P1115
	// 算法导论 练习4.1-5
	// [题型总结] 关于最大子段和及其变式 https://www.luogu.com.cn/blog/wey-yzyl/zui-tai-zi-duan-hu-ji-ji-bian-shi-di-qi-shi
	// 子段长度有上限的最大子段和：见单调队列，题目为 https://ac.nowcoder.com/acm/contest/1006/D
	// 子段长度有下限的最大子段和：转换为前缀和之差 sum[i]-sum[j]，i-j>=K，维护 mi=min(sum[j])，同时更新 sum[i]-mi 的最大值（题目见 sort.go 中的 0-1 分数规划）
	// 子段和有上限的最大字段和：转换为前缀和之差 sum[i]-sum[j]<=K，在平衡树上二分 sum[j] LC363 https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/
	// 最大两段子段和：求每个位置上的前缀最大字段和和后缀最大子段和 https://www.luogu.com.cn/problem/P2642
	// 最大 m 段子段和 https://acm.hdu.edu.cn/showproblem.php?pid=1024
	// 环状最大子段和：转换为 max(最大子段和, 总和减去最小子段和) LC918 https://leetcode-cn.com/problems/maximum-sum-circular-subarray/
	// 环状最大两段子段和：思路类似，注意取反后需要传入 a[1:n-1] https://www.luogu.com.cn/problem/P1121 https://ac.nowcoder.com/acm/contest/7738/B
	// 变体 https://codeforces.com/problemset/problem/1155/D
	// 变体 https://codeforces.com/problemset/problem/1373/D
	maxSubArraySum := func(a []int) int {
		if len(a) == 0 {
			return 0
		}
		curSum, maxSum := a[0], a[0] // int64
		for _, v := range a[1:] {
			curSum = max(curSum+v, v)
			maxSum = max(maxSum, curSum)
		}
		return max(maxSum, 0) // 若不允许非空，返回 maxSum
	}

	// 最大两段子段和（两段必须间隔至少 gap 个数）
	maxTwoSubArraySum := func(a []int, gap int) int {
		// 注意下界
		n := len(a)
		suf := make([]int, n) // int64
		suf[n-1] = a[n-1]
		curSum := a[n-1]
		for i := n - 2; i >= 0; i-- {
			v := a[i]
			curSum = max(curSum+v, v)
			suf[i] = max(suf[i+1], curSum)
		}
		curSum, pre := a[0], a[0]
		ans := pre + suf[1+gap]
		for i := 1; i < n-1-gap; i++ {
			v := a[i]
			curSum = max(curSum+v, v)
			pre = max(pre, curSum)
			ans = max(ans, pre+suf[i+1+gap])
		}
		return ans
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

	// 修改序列为非降或非增的最小修改次数
	// 单次修改可以把某个数 +1 或 -1
	// https://www.luogu.com.cn/problem/solution/P4597
	// 通过一个例子来解释这个基于堆的算法：1 5 10 4 2 2 2 2
	// 假设当前维护的是非降序列，前三个数直接插入，不需要任何修改
	// 插入 4 的时候，可以修改为 1 5 5 5，或 1 5 6 6，或... 1 5 10 10，修改次数均为 6
	// 但我们也可以把修改后的序列视作 1 5 4 4，虽然序列不为非降序列，但修改的次数仍然为 6
	// 接下来插入 2，基于 1 5 5 5 的话，修改后的序列就是 1 5 5 5 5，总的修改次数为 9
	// 但我们也可以把修改后的序列视作 1 2 4 4 2，总的修改次数仍然为 9
	// 接下来插入 2，如果基于 1 5 5 5 5 变成 1 5 5 5 5 5，会得到错误的修改次数 12
	// 但是实际上有更优的修改 1 4 4 4 4 4，总的修改次数为 11
	// 同上，把这个序列视作 1 2 2 4 2 2，总的修改次数仍然为 11
	// ...
	// https://www.luogu.com.cn/problem/P2893 http://poj.org/problem?id=3666
	// https://codeforces.com/problemset/problem/13/C
	// https://codeforces.com/problemset/problem/713/C 严格单调递增 https://codeforces.com/blog/entry/47094?#comment-315161
	//     这道题做了一个 a[i]-=i 的操作（i 从 1 开始），把严格单调递增变成了非降的情况，从而可以应用该算法
	//     这一技巧的原理是，对于整数来说，单调递增的最小情况是 y=x+C，减去这一函数，就得到了非降序列的最小情况 y=C
	// https://www.luogu.com.cn/problem/P4597 (加强版)
	minCostSorted := func(a []int) int64 {
		h := hp{} // 大根堆
		ans := int64(0)
		for _, v := range a {
			h.push(v)
			if d := h.IntSlice[0] - v; d > 0 {
				ans += int64(d)
				h.IntSlice[0] = v
				heap.Fix(&h, 0)
			}
		}
		return ans
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
	//     其中一个改为子串 https://codeforces.com/problemset/problem/163/A
	//     https://codeforces.com/problemset/problem/1446/B
	// 若其中一个序列无重复元素，可以转换成 LIS https://www.luogu.com.cn/problem/P1439 LC1713/周赛222D https://leetcode-cn.com/contest/weekly-contest-222/problems/minimum-operations-to-make-a-subsequence/
	lcs := func(s, t []byte) int {
		// dp[i][j] = LCS(s[:i], t[:j])
		n, m := len(s), len(t)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		for i, v := range s {
			for j, w := range t {
				if v == w {
					// ignore values from dp[i][j+1] and dp[i+1][j]
					dp[i+1][j+1] = dp[i][j] + 1
				} else {
					dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
				}
			}
		}

		{
			// EXTRA: 某些 dp 非单调性的题目需要计算全局最值
			allMax := 0
			for _, row := range dp {
				for _, v := range row {
					allMax = max(allMax, v)
				}
			}
		}

		return dp[n][m]
	}
	lcsPath := func(s, t []byte) []byte {
		n, m := len(s), len(t)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		fa := make([][]int8, n+1)
		for i := range fa {
			fa[i] = make([]int8, m+1)
		}
		for i, v := range s {
			for j, w := range t {
				if v == w {
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
				lcs = append(lcs, s[i-1])
			} else if fa[i][j] == 2 {
				makeLCS(i-1, j)
			} else {
				makeLCS(i, j-1)
			}
		}
		makeLCS(n, m)
		return lcs
	}

	// 最长回文子序列 (LPS)
	// LC516 https://leetcode-cn.com/problems/longest-palindromic-subsequence/
	// LC1216/双周赛10D https://leetcode-cn.com/contest/biweekly-contest-10/problems/valid-palindrome-iii/
	longestPalindromeSubsequence := func(s string) int {
		n := len(s)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
		}
		for i := n - 1; i >= 0; i-- {
			dp[i][i] = 1
			for j := i + 1; j < n; j++ {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
		}
		return dp[0][n-1]
	}

	// 最长上升子序列 (LIS)
	// O(n^2) - 定义 dp[i] 为以 a[i] 为末尾的 LIS 的长度
	//          可以把此问题想象成一个「跳跃游戏」，任选一个初始位置向右跳跃，每次只能跳到比当前位置更高的位置，问最多能跳多少次（最后答案加一）
	//          这样能更容易地看出转移的顺序，然后变成一个 DAG 上求最长路的问题
	// 转换 http://acm.hdu.edu.cn/showproblem.php?pid=1950
	// 变体 https://codeforces.com/problemset/problem/1350/B
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
	// 最小划分数 Dilworth's theorem https://en.wikipedia.org/wiki/Dilworth%27s_theorem
	// 随机排列 LIS 的长度期望 https://www.zhihu.com/question/266958886
	//
	// 例题 导弹拦截 https://www.luogu.com.cn/problem/P1020
	// 例题 LC300 https://leetcode-cn.com/problems/longest-increasing-subsequence/
	// 建模 https://codeforces.com/problemset/problem/269/B
	// 方案数 LC673 https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
	//       https://www.zhihu.com/question/34905638
	// 合唱队形 https://www.luogu.com.cn/problem/P1091
	// 合唱队形（至少有升有降）https://leetcode-cn.com/contest/biweekly-contest-40/problems/minimum-number-of-removals-to-make-mountain-array/
	// LC354 俄罗斯套娃信封问题 https://leetcode-cn.com/problems/russian-doll-envelopes/
	// 重复 T 次的 LIS 问题 https://codeforces.com/problemset/problem/582/B
	// 若其中一个序列无重复元素，LCS 可以转换成 LIS https://www.luogu.com.cn/problem/P1439 LC1713/周赛222D https://leetcode-cn.com/contest/weekly-contest-222/problems/minimum-operations-to-make-a-subsequence/
	// 二维 LIS：在一维 LIS 的基础上，a[i] 可以从多个数中选一个，问 LIS 最长可以多长
	//          思路：将各个 a[i] 的可选项从大到小排序，然后拼接成一个序列，求 LIS 即可（关键：从大到小排序避免了在同一个可选项中选择多个元素）
	lis := func(a []int) int {
		dp := []int{}
		for _, v := range a {
			if p := sort.SearchInts(dp, v); p < len(dp) { // 改成 v+1 为非降
				dp[p] = v
			} else {
				dp = append(dp, v)
			}
		}
		return len(dp)
	}
	// 每个前缀的 LIS
	lisAll := func(a []int) []int {
		n := len(a)
		lis := make([]int, n)
		dp := make([]int, 0, n)
		for i, v := range a {
			p := sort.SearchInts(dp, v)
			if p < len(dp) { // 改成 v+1 为非降
				dp[p] = v
			} else {
				dp = append(dp, v)
			}
			lis[i] = p + 1
		}
		return lis
	}

	// LIS 相关构造题
	// https://codeforces.com/problemset/problem/1304/D
	// https://atcoder.jp/contests/arc091/tasks/arc091_c

	// 最长公共上升子序列 (LCIS)
	// https://www.acwing.com/problem/content/274/
	// https://codeforces.com/problemset/problem/10/D
	lcis := func(a, b []int) int {
		n, m := len(a), len(b)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m)
		}
		for i, v := range a {
			mx := 0
			for j, w := range b {
				if v == w {
					dp[i+1][j] = mx + 1
				} else {
					dp[i+1][j] = dp[i][j]
				}
				if w < v {
					mx = max(mx, dp[i][j])
				}
			}
		}
		ans := 0
		for _, v := range dp[n] {
			ans = max(ans, v)
		}
		return ans
	}

	// LCIS 打印方案
	lcisPath := func(a, b []int) (ans int, lcis []int) {
		n, m := len(a), len(b)
		dp := make([][]int, n+1)
		fa := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m)
			fa[i] = make([]int, m)
		}
		for i, v := range a {
			mx, k := 0, -1
			for j, w := range b {
				if v == w {
					dp[i+1][j] = mx + 1
					fa[i+1][j] = k // k < j
				} else {
					dp[i+1][j] = dp[i][j]
					fa[i+1][j] = j
				}
				if w < v && dp[i][j] > mx {
					mx, k = dp[i][j], j
				}
			}
		}
		ansJ := 0
		for j, dv := range dp[n] {
			if dv > dp[n][ansJ] {
				ansJ = j
			}
		}
		ans = dp[n][ansJ]
		var getLCIS func(i, j int)
		getLCIS = func(i, j int) {
			if i == 0 || j < 0 {
				return
			}
			getLCIS(i-1, fa[i][j])
			if fa[i][j] < j {
				lcis = append(lcis, b[j])
			}
		}
		getLCIS(n, ansJ)
		return
	}

	// 长度为 m 的 LIS 个数
	// 赤壁之战 https://www.acwing.com/problem/content/299/
	// 定义 dp[i][j] 表示 a[:j+1] 的长度为 i 且以 a[j] 结尾的 LIS
	// 则有 dp[i][j] = ∑dp[i-1][k]  (k<j && a[k]<a[j])
	// 注意到当 j 增加 1 时，只多了 k=j 这一个新决策，这样可以用树状数组来维护
	// 复杂度 O(mnlogn)
	countLIS := func(a []int, m int) int {
		// 将 a 离散化成从 2 开始的序列
		b := append([]int(nil), a...)
		sort.Ints(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + 2
		}

		n := len(a)
		const mod int = 1e9 + 7
		tree := make([]int, n+2)
		add := func(i, val int) {
			for ; i < n+2; i += i & -i {
				tree[i] = (tree[i] + val) % mod
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res = (res + tree[i]) % mod
			}
			return
		}

		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n)
		}
		for i := 1; i <= m; i++ {
			tree = make([]int, n+2)
			if i == 1 {
				add(1, 1)
			}
			for j, v := range a {
				dp[i][j] = sum(v - 1)
				add(v, dp[i-1][j])
			}
		}
		ans := 0
		for _, v := range dp[m] {
			ans = (ans + v) % mod
		}
		return ans
	}

	// 本质不同子序列个数
	// 定义 dp[i][j] 表示前 i 个字符中长度为 j 的本质不同子序列个数
	// 转移 dp[i][j] = dp[i-1][j]（不选第 i 个字符）+ dp[i-1][j-1] - dp[prev[i]-1][j-1]（选第 i 个字符）
	// 其中 prev[i] 为 s[i] 的上一个相同字符位置
	// https://ac.nowcoder.com/acm/contest/4853/C 题解 https://ac.nowcoder.com/discuss/394080
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

	// 回文串最小分割次数
	// LC132 https://leetcode-cn.com/problems/palindrome-partitioning-ii/
	minPalindromeCut := func(s string) int {
		n := len(s)
		g := make([][]bool, n)
		for i := range g {
			g[i] = make([]bool, n)
			for j := range g[i] {
				g[i][j] = true
			}
		}
		for i := n - 1; i >= 0; i-- {
			for j := i + 1; j < n; j++ {
				g[i][j] = s[i] == s[j] && g[i+1][j-1]
			}
		}

		f := make([]int, n)
		for i := range f {
			if g[0][i] { // f[i] = 0
				continue
			}
			f[i] = int(1e9)
			for j := 0; j < i; j++ {
				if g[j+1][i] {
					f[i] = min(f[i], f[j]+1)
				}
			}
		}
		return f[n-1]
	}

	/* 背包问题
	这类问题可以从物品选择次序的无后效性入手
	子区间 -> 前缀和
	子序列 -> 背包
	https://en.wikipedia.org/wiki/Knapsack_problem
	https://codeforces.com/blog/entry/59606
	浅谈 ZKP 问题 https://www.luogu.com.cn/blog/xww666/qian-tan-zkp-wen-ti-gai-post
	另见 math_ntt.go 中的生成函数

	NOTE: 若求能否凑成 1,2,3,...,M，只需判断 dp[i] 是否为正 LC1049 https://leetcode-cn.com/problems/last-stone-weight-ii/
	套题 https://www.acwing.com/problem/
	混合背包 https://www.luogu.com.cn/problem/P1833
	*/

	// 0-1 背包 (n 个物品，背包容量为 maxW)
	// 状态：从前 i 个物品中选择若干个，当容量限制为 j 时能获得的最大价值和  i∈[0,n-1], j∈[0,maxW]
	// 初始值：f(0,j)=0  j∈[0,maxW]
	// 除开初始状态，每个状态有两个来源，决策为 max：
	// - 不选第 i 个物品：f(i-1,j) -> f(i,j)
	// - 选第 i 个物品：f(i-1,j-wi)+vi -> f(i,j)   j≥wi
	// 最优解为 f(n-1,maxW)
	// https://oi-wiki.org/dp/knapsack/
	// 模板题 https://www.luogu.com.cn/problem/P1048 https://atcoder.jp/contests/dp/tasks/dp_d
	// 转换 LC1049 https://leetcode-cn.com/problems/last-stone-weight-ii/
	// 转换 https://codeforces.com/problemset/problem/1381/B
	// 打印方案 https://codeforces.com/problemset/problem/864/E
	// EXTRA: 恰好装满（相当于方案数不为 0）LC416 https://leetcode-cn.com/problems/partition-equal-subset-sum/
	// EXTRA: 背包容量为 0 https://codeforces.com/problemset/problem/366/C
	// EXTRA: 二维费用 https://www.acwing.com/problem/content/8/ https://www.luogu.com.cn/problem/P1507 LC474 https://leetcode-cn.com/problems/ones-and-zeroes/
	// EXTRA: 把一个维度转换成 DP 的定义 https://codeforces.com/problemset/problem/837/D
	// EXTRA: 离散化背包 https://codeforces.com/contest/366/submission/61452111
	zeroOneKnapsack := func(values, weights []int, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for i, v := range values {
			w := weights[i]
			for j := maxW; j >= w; j-- {
				dp[j] = max(dp[j], dp[j-w]+v)
			}
		}
		return dp[maxW]
	}

	// 0-1 背包 EXTRA: 至少装满 https://www.luogu.com.cn/problem/P4377
	// 二维费用的情况+价值最小 https://ac.nowcoder.com/acm/contest/6218/C
	zeroOneKnapsackAtLeastFillUp := func(values, weights []int, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for i := range dp {
			dp[i] = -1e18 // 价值最小改成 1e18
		}
		dp[0] = 0
		for i, v := range values {
			w := weights[i]
			for j := maxW; j >= 0; j-- {
				dp[j] = max(dp[j], dp[max(j-w, 0)]+v) // max(j-w, 0) 蕴含了「至少」
			}
		}

		{
			// 另一种写法
			for i, v := range values {
				w := weights[i]
				for j := maxW; j >= 0; j-- {
					k := min(j+w, maxW)
					dp[k] = max(dp[k], dp[j]+v)
				}
			}
		}

		return dp[maxW]
	}

	// 0-1 背包 EXTRA: 从序列 a 中选若干个数，使其总和为 sum 的方案数
	// NOTE: 1,1,1,...1(32个1),s-32,s-31,...,s 可以让方案数恰好为 2^32
	// 转换 LC494 https://leetcode-cn.com/problems/target-sum/
	// 二维+上限+下限 LC879/周赛95D https://leetcode-cn.com/contest/weekly-contest-95/problems/profitable-schemes/
	// 隐藏的 0-1 背包 LC1434 https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/
	// 建模转换 https://atcoder.jp/contests/abc169/tasks/abc169_f
	// 由于顺序不同也算方案，所以这题需要正序递推 LC377 https://leetcode-cn.com/problems/combination-sum-iv/
	zeroOneWaysToSum := func(a []int, sum int) int {
		dp := make([]int, sum+1) // int64
		dp[0] = 1
		for _, v := range a {
			for s := sum; s >= v; s-- {
				dp[s] += dp[s-v] // % mod
			}
		}
		return dp[sum]
	}

	// 0-1 背包 EXTRA: 打印字典序最小的方案
	// 倒序遍历物品，同时用 fa 数组记录转移来源，这样跑完 DP 后，从第一个物品开始即可得到字典序最小的方案
	// https://www.acwing.com/problem/content/description/12/
	zeroOneKnapsackLexicographicallySmallestResult := func(values, weights []int, maxW int) (ans []int) {
		n := len(values)
		dp := make([]int, maxW+1) // int64  fill
		//dp[0] = 0
		fa := make([][]int, n)
		for i := n - 1; i >= 0; i-- {
			fa[i] = make([]int, maxW+1)
			for j := range fa[i] {
				fa[i][j] = j
			}
			v, w := values[i], weights[i]
			for j := maxW; j >= w; j-- {
				if dp[j-w]+v >= dp[j] { // 注意这里要取等号，从而保证尽可能地从字典序最小的方案转移过来
					dp[j] = dp[j-w] + v
					fa[i][j] = j - w
				}
			}
		}
		for i, j := 0, maxW; i < n; {
			if fa[i][j] == j {
				i++
			} else {
				ans = append(ans, i+1) // 下标从 1 开始
				j = fa[i][j]
				i++ // 完全背包的情况，这行去掉
			}
		}
		return
	}

	// 0-1 背包 EXTRA: 价值主导的 0-1 背包
	// 把重量看成价值，价值看成重量，求同等价值下能得到的最小重量

	// 完全背包
	// 转换 LC322 https://leetcode-cn.com/problems/coin-change/
	// EXTRA: 打印方案 LC1449/双周赛26D https://leetcode-cn.com/contest/biweekly-contest-26/problems/form-largest-integer-with-digits-that-add-up-to-target/
	unboundedKnapsack := func(values, weights []int, maxW int) int {
		dp := make([]int, maxW+1) // int64  fill
		//dp[0] = 0
		for i, v := range values {
			w := weights[i]
			for j := w; j <= maxW; j++ {
				dp[j] = max(dp[j], dp[j-w]+v)
			}
		}
		return dp[maxW]
	}

	// 完全背包 EXTRA: 方案数
	// LC518 https://leetcode-cn.com/problems/coin-change-2/
	// https://www.luogu.com.cn/problem/P1832
	// https://www.luogu.com.cn/problem/P6205（需要高精）
	// 类似完全背包但是枚举的思路不一样 LC377 https://leetcode-cn.com/problems/combination-sum-iv/
	unboundedWaysToSum := func(a []int, sum int) int {
		dp := make([]int, sum+1) // int64
		dp[0] = 1
		for _, v := range a {
			for s := v; s <= sum; s++ {
				dp[s] += dp[s-v] // % mod
			}
		}
		return dp[sum]
	}

	// 完全背包 EXTRA: 二维费用方案数
	// 注意：「恰好使用 m 个物品」这个条件要当成一种费用来看待
	// https://codeforces.com/problemset/problem/543/A

	// 多重背包 - 未优化
	boundedKnapsack := func(values, stocks, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1) // int64
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

	// 多重背包 - 优化 1 - 二进制优化
	// 模板题 https://codeforces.com/problemset/problem/106/C
	// todo 多重背包+完全背包 https://www.luogu.com.cn/problem/P1782 https://www.luogu.com.cn/problem/P1833 https://www.luogu.com.cn/problem/P2851
	// http://acm.hdu.edu.cn/showproblem.php?pid=2844 http://poj.org/problem?id=1742
	// https://www.luogu.com.cn/problem/P6771 http://poj.org/problem?id=2392
	// https://codeforces.com/contest/999/problem/F
	boundedKnapsackBinary := func(values, stocks, weights []int, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for i, v := range values {
			num, w := stocks[i], weights[i]
			for k := 1; num > 0; k <<= 1 {
				K := min(k, num)
				for j := maxW; j >= K*w; j-- {
					dp[j] = max(dp[j], dp[j-K*w]+K*v)
				}
				num -= K
			}
		}
		return dp[maxW]
	}

	// 多重背包 - 优化 2 - 单调队列优化
	// todo 挑战 P340

	// 分组背包
	// https://www.acwing.com/problem/content/9/
	// https://www.luogu.com.cn/problem/P1757
	type item struct{ v, w int }
	groupKnapsack := func(groups [][]item, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for _, g := range groups {
			for j := maxW; j >= 0; j-- {
				for _, it := range g {
					if v, w := it.v, it.w; w <= j {
						dp[j] = max(dp[j], dp[j-w]+v)
					}
				}
			}
		}
		return dp[maxW]
	}

	// 树上背包/树形背包/依赖背包
	// todo 树上背包的上下界优化 https://ouuan.gitee.io/post/%E6%A0%91%E4%B8%8A%E8%83%8C%E5%8C%85%E7%9A%84%E4%B8%8A%E4%B8%8B%E7%95%8C%E4%BC%98%E5%8C%96/
	//   子树合并背包的复杂度证明 https://blog.csdn.net/lyd_7_29/article/details/79854245
	//   复杂度优化 https://loj.ac/d/3144
	//   https://zhuanlan.zhihu.com/p/103813542
	//
	// todo https://loj.ac/p/160
	//   https://www.luogu.com.cn/problem/P2014 https://www.acwing.com/problem/content/10/ https://www.acwing.com/problem/content/288/
	//   加强版 https://www.luogu.com.cn/problem/U53204
	//   https://www.luogu.com.cn/problem/P1272
	//   加强版 https://www.luogu.com.cn/problem/U53878
	//   https://www.luogu.com.cn/problem/P3177
	treeKnapsack := func(g [][]int, items []item, root, maxW int) int {
		var f func(int) []int
		f = func(v int) []int {
			it := items[v]
			dp := make([]int, maxW+1) // int64
			for i := it.w; i <= maxW; i++ {
				dp[i] = it.v // 根节点必须选
			}
			for _, to := range g[v] {
				dt := f(to)
				for j := maxW; j >= it.w; j-- {
					// 类似分组背包，枚举分给子树 to 的容量 w，对应的子树的最大价值为 dt[w]
					// w 不可超过 j-it.w，否则无法选择根节点
					for w := 0; w <= j-it.w; w++ {
						dp[j] = max(dp[j], dp[j-w]+dt[w])
					}
				}
			}
			return dp
		}
		return f(root)[maxW]
	}

	/* 区间 DP
	一般来说转移是合并区间或者分解区间
	套路 https://www.luogu.com.cn/blog/BreakPlus/ou-jian-dp-zong-jie-ti-xie
	① 将序列分成 K 个连续区间，求解这些区间的某个最优性质
	一般定义 dp[i][k] 表示将 a[:i] 分成 k 个连续区间得到的最优解
	此时可以枚举最后一个区间的左端点 j，从 dp[j-1][k-1] 转移到 dp[i][k]，转移时考虑 a[j:i] 对最优解的影响
	力扣题目 1278,813,410,1335
	② 求解关于某个序列的最优性质，要求大区间的最优解可以依赖于小区间的最优解
	一般定义 dp[i][j] 表示 a[i:j] 的最优解
	此时可以枚举区间大小和区间左端点，从小区间转移到大区间
	力扣题目 516,312,375,1246
	戳气球（好题） LC312 https://leetcode-cn.com/problems/burst-balloons/
	移除盒子（状态定义和转移的好题） LC546/周赛25D https://leetcode-cn.com/problems/remove-boxes/ https://leetcode.com/contest/leetcode-weekly-contest-25
	打印机（好题） LC664 https://leetcode-cn.com/problems/strange-printer/
	最优三角剖分 LC1039 https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/
	删除回文子数组 LC1246/双周赛12D https://leetcode-cn.com/contest/biweekly-contest-12/problems/palindrome-removal/
	同色消除【套路】 https://www.luogu.com.cn/problem/P4170 https://codeforces.com/problemset/problem/1132/F
	③ 一些题目
	https://blog.csdn.net/weixin_43914593/article/details/106163859 算法竞赛专题解析（14）：DP应用--区间DP
	todo https://atcoder.jp/contests/abc159/tasks/abc159_f
	     https://codeforces.com/problemset/problem/245/H
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

	/* 环形 DP
	两种解题策略：
	一种是假设在 0 处断开，求一遍 DP，然后强制让 0 和 n-1 上是连通的，再求一遍 DP，取二者最值
	另一种是倍增链
	休息时间 https://www.luogu.com.cn/problem/P6064 https://www.acwing.com/problem/content/290/
	环路运输 https://www.acwing.com/problem/content/291/
	https://www.luogu.com.cn/problem/P1453
	*/

	/* 博弈类 DP
	转移：让「自己与对手的分差」最大
	图上博弈 https://codeforces.com/problemset/problem/917/B
	LC877 https://leetcode-cn.com/problems/stone-game/ https://nanti.jisuanke.com/t/48
	LC1140 https://leetcode-cn.com/problems/stone-game-ii/
	LC1406 https://leetcode-cn.com/problems/stone-game-iii/
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp%2Cgames
	*/

	/* 状压 DP
	常用于处理包含排列的问题等
	NOTE: 若问题无法划分成小问题，必须考虑各种可能的情况，则可能是 NP 完全问题
	浅谈状压 DP https://www.luogu.com.cn/blog/yijan/zhuang-ya-dp
	https://blog.csdn.net/weixin_43914593/article/details/106432695 算法竞赛专题解析（15）：DP应用--状态压缩DP

	https://www.luogu.com.cn/problem/P1879
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp%2Cbitmasks

	todo 题单 https://ac.nowcoder.com/acm/problem/collection/808
	     题单 https://ac.nowcoder.com/acm/problem/collection/810
	todo LC691  https://leetcode-cn.com/problems/stickers-to-spell-word/
	     LC1125 https://leetcode-cn.com/problems/smallest-sufficient-team/
	     LC943  https://leetcode-cn.com/problems/find-the-shortest-superstring/

	todo 汉密尔顿路径/回路 Hamiltonian path
	https://en.wikipedia.org/wiki/Hamiltonian_path
	https://en.wikipedia.org/wiki/Hamiltonian_path_problem
	*/

	// 旅行商问题 (TSP)
	// 返回一个 ans 数组，ans[i] 表示从 st 出发，访问完所有位置且最后停在 i 的最短路径（注意可能要特判 i==st 的情况）
	// 做法：定义 dp[s][v] 表示已访问的集合为 s，最后一个访问的位置是 v 时的最小花费
	//      则有 dp[s|1<<w][w] = min(dp[s|1<<w][w], ds[v]+dist[v][w])
	//      枚举 v 和 w 时可以用 TrailingZeros 来直接枚举每个 1 和 0 的位置
	// https://en.wikipedia.org/wiki/Travelling_salesman_problem
	// 模板题 https://www.luogu.com.cn/problem/P1171 https://www.luogu.com.cn/problem/P1433 https://www.acwing.com/problem/content/93/
	// 恰好访问 m 个点 https://codeforces.com/contest/580/problem/D
	// 建模转换题 LC943 https://leetcode-cn.com/problems/find-the-shortest-superstring/
	//          LCP13 https://leetcode-cn.com/problems/xun-bao/
	// EXTRA: 固定起点终点的问题，视问题情况有两种方法：
	//        添加一个节点 https://stackoverflow.com/questions/14527815/how-to-fix-the-start-and-end-points-in-travelling-salesmen-problem
	//        设置距离 https://stackoverflow.com/questions/36086406/traveling-salesman-tsp-with-set-start-and-end-point
	tsp := func(dist [][]int, st int) []int {
		n := len(dist)
		const inf int = 1e9 // 1e18
		dp := make([][]int, 1<<n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = inf
			}
		}
		dp[1<<st][st] = 0 // 多个起点的话就设置多个 dp[1<<st[i]][st[i]] = 0
		for s, dv := range dp {
			// 利用位运算快速求出 s 中 1 的位置 v，以及 s 中 0 的位置 w（通过 s 的补集中的 1 的位置求出）
			for S := uint(s); S > 0; S &= S - 1 {
				v := bits.TrailingZeros(S)
				for C := uint(s) ^ (1<<n - 1); C > 0; C &= C - 1 {
					w := bits.TrailingZeros(C)
					dp[s|1<<w][w] = min(dp[s|1<<w][w], dv[v]+dist[v][w])
				}
			}
		}
		return dp[1<<n-1]
	}

	// 枚举子集的子集，复杂度 O(3^m) (元素个数为 k 的集合有 C(m,k) 个，其子集有 2^k 个，∑C(m,k)*2^k = (2+1)^m = 3^m)
	// 例如：dp[set] = min{dp[set^sub] + cost of sub} for all valid sub
	// LC1494/双周赛29D https://leetcode-cn.com/contest/biweekly-contest-29/problems/parallel-courses-ii/
	// LC1654/双周赛39D https://leetcode-cn.com/contest/biweekly-contest-39/problems/distribute-repeating-integers/
	subsubDP := func(a, cost []int, limit int) int {
		n, m := len(a), len(cost)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, 1<<m)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		const inf int = 1e9 // 1e18
		var f func(p, set int) int
		f = func(p, set int) (res int) {
			if p == n {
				if set > 0 {
					return inf
				}
				return
			}
			dv := &dp[p][set]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
			res = inf

			// 所有子集
			for sub, ok := set, true; ok; ok = sub != set {
				s := 0
				for mask := uint(sub); mask > 0; mask &= mask - 1 {
					s += cost[bits.TrailingZeros(mask)]
				}
				r := f(p+1, set^sub)
				res = min(res, r+s)
				sub = (sub - 1) & set
			}

			// 所有非空子集
			for sub := set; sub > 0; sub = (sub - 1) & set {

			}

			return
		}
		return f(0, 1<<m-1)
	}

	/* 插头 DP / 轮廓线 DP / Broken Profile DP
	《训练指南》6.1
	todo https://oi-wiki.org/dp/plug/
	https://cp-algorithms.com/dynamic_programming/profile-dynamics.html
	https://www.luogu.com.cn/blog/efforts-will-pay-off/du-liu-dong-gui-cha-tou-dp

	入门题 https://www.luogu.com.cn/problem/P3272
	*/

	/* 数位 DP
	https://zhuanlan.zhihu.com/p/348851463

	入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
	      https://atcoder.jp/contests/dp/tasks/dp_s
	      https://codeforces.com/problemset/problem/1036/C
	含有某个数字的数字个数
	LC233 https://leetcode-cn.com/problems/number-of-digit-one/
	      https://leetcode-cn.com/problems/number-of-2s-in-range-lcci/
	      http://acm.hdu.edu.cn/showproblem.php?pid=3555
	      http://acm.hdu.edu.cn/showproblem.php?pid=2089
	LC600 二进制不含连续 1 的数字个数 https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/
	LC902/周赛101C 最大为 N 的数字组合 https://leetcode-cn.com/contest/weekly-contest-101/problems/numbers-at-most-n-given-digit-set/
	LC1012/周赛128D 有重复数字的数字个数 https://leetcode-cn.com/contest/weekly-contest-128/problems/numbers-with-repeated-digits/
	LC1067/双周赛1D 字符 d 出现次数 https://leetcode-cn.com/contest/biweekly-contest-1/problems/digit-count-in-range/
	LC1397/周赛182D 与 KMP 结合 https://leetcode-cn.com/contest/weekly-contest-182/problems/find-all-good-strings/
	digsum(n)|n 的数的个数 https://www.luogu.com.cn/problem/P4127 https://www.acwing.com/problem/content/313/
	http://acm.hdu.edu.cn/showproblem.php?pid=3886
	http://acm.hdu.edu.cn/showproblem.php?pid=6796
	注：一些第 k 小的题目需要与二分结合，或者用试填法（见后面的 kth666）
	todo 套题 https://www.luogu.com.cn/blog/s-r-f/oi-bi-ji-shuo-wei-dp-ge-ji-dui-shuo-wei-dp-di-yi-dian-li-xie
	todo 套题 https://codeforces.com/blog/entry/53960
	*/
	digitDP := func(lower, upper string) int {
		const mod int = 1e9 + 7

		// 返回 <=s 的符合要求的字符串数目
		// TIPS: 某些情况下思考补集会更加容易，即求不符合要求的字符串数目
		// TIPS: 对于需要判断/禁止前导零的情况，可以加一个额外的维度 hasD 表示是否有非零数字（意为「真正填了数字」），最后 p>=n 的时候根据情况返回 1 或者 0
		calc := func(s string) int {
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
			var f func(p, sum int, limitUp bool) int
			f = func(p, sum int, limitUp bool) (res int) {
				if p == n {
					return 1
				} // sum
				if !limitUp {
					dv := &dp[p][sum]
					if *dv >= 0 {
						return *dv
					} // *dv + sum*int64(math.Pow10(n-p))
					defer func() { *dv = res }()
				}
				up := upperC
				if limitUp {
					up = s[p]
				}
				for d := lowerC; d <= up; d++ {
					tmp := sum

					c := f(p+1, tmp, limitUp && d == up)
					res = (res + c) % mod
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

	// 试填法
	// 第 k 个包含 3 个连续的 6 的数 https://www.acwing.com/problem/content/312/
	kth666 := func(k int) (ans []byte) {
		// dp[i][3] 表示由 i 位数字构成的魔鬼数的个数
		// dp[i][j] (j<3) 表示 i 位数字构成的、开头有连续 j 个 6 的非魔鬼数的个数
		const mx = 30  // 长度上限
		const cont = 3 // 连续 3 个数才算符合要求
		dp := [mx][cont + 1]int{}
		dp[0][0] = 1
		for i := 1; i < mx; i++ {
			for j := 0; j < cont; j++ {
				dp[i][0] += dp[i-1][j] * 9 // 开头无 6，直接转移（0-9 中除去 6 共 9 个数）
				dp[i][j+1] = dp[i-1][j]    // 开头有 j+1 个 6，下一个有 j 个 6
			}
			dp[i][cont] += dp[i-1][cont] * 10
		}

		const tarDigit byte = '6'
		n := 1
		for ; dp[n][cont] < k; n++ {
		}
		has := 0
		for i := 1; i <= n; i++ {
			for digit := byte('0'); digit <= '9'; digit++ { // 试填当前位
				need := cont
				if has == cont {
					need = 0
				} else if digit == tarDigit {
					need = cont - 1 - has
				}
				sum := 0
				for j := need; j <= cont; j++ {
					sum += dp[n-i][j]
				}
				if sum >= k { // 填入
					ans = append(ans, digit)
					if has < cont {
						if digit == tarDigit {
							has++
						} else {
							has = 0
						}
					}
					break
				}
				k -= sum
			}
		}
		return
	}

	/* 倍增优化 DP

	 */

	/* 数据结构优化 DP
	https://codeforces.com/problemset?order=BY_RATING_ASC&tags=data+structures%2Cdp

	动态 DP - 线段树维护转移矩阵
	todo https://www.cnblogs.com/Miracevin/p/9124511.html
	1. 先不考虑修改，不考虑区间，直接列出整个区间的 DP 转移。
	2. 列出转移矩阵。由于有很多修改操作，我们将数据集中在一起处理，还可以利用矩阵结合律，并且区间比较好提取（找一段矩阵就好了），修改也方便。
	3. 线段树维护矩阵。对于修改，我们就是在矩阵上进行修改。对于不同的题目，我们要用不同的修改方式和记录手段，但都是用线段树维护对应区间内的矩阵信息。如矩阵乘积，矩阵和等等。线段树的区间操作可以应对区间修改问题。
	https://codeforces.com/problemset/problem/1380/F
	https://codeforces.com/problemset/problem/718/C
	https://codeforces.com/problemset/problem/750/E
	*/

	// 单调队列优化
	// 见 monotone_queue.go

	// 斜率优化 / 凸包优化 (Convex Hull Trick, CHT)
	// https://oi-wiki.org/dp/opt/slope/
	// https://cp-algorithms.com/geometry/convex_hull_trick.html
	// https://www.luogu.com.cn/blog/ChenXingLing/post-xue-xi-bi-ji-dong-tai-gui-hua-xie-shuai-you-hua-dp-chao-yang-x
	// https://blog.csdn.net/weixin_43914593/article/details/105560357 算法竞赛专题解析（12）：DP优化(2)--斜率(凸壳)优化
	// https://zhuanlan.zhihu.com/p/363772434
	// https://codeforces.com/blog/entry/63823
	//
	// https://www.luogu.com.cn/problem/P2365 https://www.luogu.com.cn/problem/P5785 http://poj.org/problem?id=1180
	// todo https://www.luogu.com.cn/problem/P2900
	//  https://www.luogu.com.cn/problem/P3195 https://loj.ac/p/10188
	//  http://poj.org/problem?id=3709
	cht := func(n int, s int64, sumT, sumC []int64) int64 {
		// 以 https://www.luogu.com.cn/problem/P5785 为例
		dp := make([]int64, n+1)
		Y := func(i int) int64 { return dp[i] }
		X := func(i int) int64 { return sumC[i] }
		// 下凸包：i0-i1 的斜率 < i1-i2 的斜率
		// 上凸包：i0-i1 的斜率 > i1-i2 的斜率
		less := func(i0, i1, i2 int) bool {
			y0, y1, y2 := Y(i0), Y(i1), Y(i2)
			x0, x1, x2 := X(i0), X(i1), X(i2)
			return (y1-y0)*(x2-x1) < (y2-y1)*(x1-x0) // 注意 < 还是 > 以及是否会爆 int64
		}
		q := []int{0} //
		push := func(i int) {
			for len(q) > 1 && !less(q[len(q)-2], q[len(q)-1], i) {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
		for i := 1; i <= n; i++ {
			k := s + sumT[i]
			// 在队列中二分第一个斜率大于（小于）k 的位置
			// 下凸包：>
			// 上凸包：<
			j := sort.Search(len(q)-1, func(j int) bool { return Y(q[j+1])-Y(q[j]) > k*(X(q[j+1])-X(q[j])) })
			dp[i] = Y(q[j]) - k*X(q[j]) + sumT[i]*sumC[i] + s*sumC[n]
			push(i)
		}
		return dp[n]
	}

	// 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/
	// todo https://blog.csdn.net/weixin_43914593/article/details/105150937 算法竞赛专题解析（10）：DP优化(1)--四边形不等式
	//      决策单调性优化讲解 https://www.luogu.com.cn/blog/83547/zong-dong-tai-gui-hua-di-ben-zhi-kan-si-bian-xing-fou-deng-shi-you-hua
	// 扔蛋问题 LC887 https://leetcode-cn.com/problems/super-egg-drop/

	/* 树形 DP
	https://blog.csdn.net/weixin_43914593/article/details/107145592
	https://codeforces.com/blog/entry/20935
	https://codeforces.com/blog/entry/63257

	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp%2Ctrees
	todo 题单 https://ac.nowcoder.com/acm/problem/collection/807
	     题单 https://ac.nowcoder.com/acm/problem/collection/809
	https://codeforces.com/problemset/problem/982/C
	好题 https://codeforces.com/problemset/problem/1453/E
	*/

	// 树的直径（两遍 DFS 求法另见 graph_tree.go 中的 diameter）
	// LC1245 https://leetcode-cn.com/problems/tree-diameter/
	diameter := func(st int, g [][]int) (diameter int) {
		var f func(v, fa int) int
		f = func(v, fa int) (mxDep int) {
			for _, w := range g[v] {
				if w != fa {
					dep := f(w, v) + 1
					diameter = max(diameter, mxDep+dep)
					mxDep = max(mxDep, dep)
				}
			}
			return
		}
		f(st, -1)
		return
	}

	// 树的直径及其个数
	// http://acm.hdu.edu.cn/showproblem.php?pid=3534
	// https://ac.nowcoder.com/acm/contest/view-submission?submissionId=45988692
	countDiameter := func(st int, g [][]int) (diameter, diameterCnt int) {
		var f func(v, fa int) (int, int)
		f = func(v, fa int) (int, int) {
			mxDep, cnt := 0, 1
			for _, w := range g[v] {
				if w != fa {
					d, c := f(w, v)
					if l := mxDep + d; l > diameter {
						diameter, diameterCnt = l, cnt*c
					} else if l == diameter {
						diameterCnt += cnt * c
					}
					if d > mxDep {
						mxDep, cnt = d, c
					} else if d == mxDep {
						cnt += c
					}
				}
			}
			return mxDep + 1, cnt
		}
		f(st, -1)
		return
	}

	// 树的直径及在直径上的节点个数
	// https://ac.nowcoder.com/acm/contest/view-submission?submissionId=45987468
	// 注意这里的 cnt 初始化与 countDiameter 的不同之处
	countVerticesOnDiameter := func(st int, g [][]int) (diameter, verticesCnt int) {
		var f func(v, fa int) (int, int)
		f = func(v, fa int) (int, int) {
			mxDep, cnt := 0, 0
			for _, w := range g[v] {
				if w != fa {
					d, c := f(w, v)
					if l := mxDep + d; l > diameter {
						diameter, verticesCnt = l, cnt+c+1 // 最长的链 + 当前链 + 当前节点
					} else if l == diameter {
						verticesCnt += c
					}
					if d > mxDep {
						mxDep, cnt = d, c
					} else if d == mxDep {
						cnt += c
					}
				}
			}
			return mxDep + 1, cnt + 1
		}
		f(st, -1)
		return
	}

	// 树上最大独立集
	// 返回最大点权和（最大独立集的情形即所有点权均为一）
	// 每个点有选和不选两种决策，接受子树转移时，选的决策只能加上不选子树，而不选的决策可以加上 max{不选子树, 选子树}
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	// https://stackoverflow.com/questions/13544240/algorithm-to-find-max-independent-set-in-a-tree
	// 经典题：没有上司的舞会 https://www.luogu.com.cn/problem/P1352 https://ac.nowcoder.com/acm/problem/51178
	maxIndependentSetOfTree := func(n int, g [][]int, a []int) int { // 无根树
		var f func(int, int) (notChosen, chosen int)
		f = func(v, fa int) (notChosen, chosen int) { // int64
			chosen = a[v]
			for _, w := range g[v] {
				if w != fa {
					nc, c := f(w, v)
					notChosen += max(nc, c)
					chosen += nc
				}
			}
			return
		}
		nc, c := f(0, -1)
		return max(nc, c)
	}

	// 树上最小边覆盖
	// 代码和树上最大独立集类似
	// 经典题：战略游戏 https://www.luogu.com.cn/problem/P2016
	minEdgeCoverOfTree := func(n int, g [][]int, a []int) int { // 无根树
		var f func(int, int) (notChosen, chosen int)
		f = func(v, fa int) (notChosen, chosen int) { // int64
			chosen = a[v]
			for _, w := range g[v] {
				if w != fa {
					nc, c := f(w, v)
					notChosen += c
					chosen += min(nc, c)
				}
			}
			return
		}
		nc, c := f(0, -1)
		return min(nc, c)
	}

	// 树上最小支配集/最小顶点覆盖
	// 返回最小点权和（最小支配集的情形即所有点权均为一）
	// 下面的定义省去了（……时的最小支配集的元素个数）   w 为 i 的儿子
	// dp[i][0]：i 属于支配集 = a[i]+∑min(dp[w][0],dp[w][1],dp[w][2])
	// dp[i][1]：i 不属于支配集，且被儿子支配 = ∑min(dp[w][0],dp[w][1]) + 如果全选 dp[w][1] 则补上 min{dp[w][0]-dp[w][1]}
	// dp[i][2]：i 不属于支配集，且被父亲支配 = ∑min(dp[w][0],dp[w][1])
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	// 经典题：保安站岗 https://www.luogu.com.cn/problem/P2458 手机网络 https://www.luogu.com.cn/problem/P2899 https://ac.nowcoder.com/acm/problem/24953
	// 监控二叉树 LC968 https://leetcode-cn.com/problems/binary-tree-cameras/
	// todo EXTRA: 消防局的设立（支配距离为 2） https://www.luogu.com.cn/problem/P2279
	// todo EXTRA: 将军令（支配距离为 k） https://www.luogu.com.cn/problem/P3942
	//                                https://atcoder.jp/contests/arc116/tasks/arc116_e
	minDominatingSetOfTree := func(n int, g [][]int, a []int) int { // 无根树
		const inf int = 1e9 // 1e18
		var f func(int, int) (chosen, bySon, byFa int)
		f = func(v, fa int) (chosen, bySon, byFa int) { // int64
			chosen = a[v]
			extra := inf
			for _, w := range g[v] {
				if w != fa {
					c, bs, bf := f(w, v)
					m := min(c, bs)
					chosen += min(m, bf)
					bySon += m
					byFa += m
					extra = min(extra, c-bs)
				}
			}
			if extra > 0 {
				bySon += extra
			}
			return
		}
		chosen, bySon, _ := f(0, -1)
		return min(chosen, bySon)
	}

	// 树上最大匹配
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	// https://codeforces.com/blog/entry/2059
	// https://blog.csdn.net/lycheng1215/article/details/78368002
	// https://vijos.org/p/1892
	maxMatchingOfTree := func(n int, g [][]int) int { // 无根树
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

	// todo 给一棵树和树上的一些关键节点，选 m 个点，使得关键节点到这些点中距离的最小值的最大值最小，求这个值
	//      https://www.luogu.com.cn/problem/P3523

	// 换根 DP
	// 进阶指南 p.292-295
	// https://codeforces.com/blog/entry/20935
	//
	// https://www.luogu.com.cn/problem/P3478
	// https://codeforces.com/problemset/problem/1092/F
	// https://www.luogu.com.cn/problem/P2986
	// https://codeforces.com/problemset/problem/219/D
	// https://codeforces.com/problemset/problem/337/D
	// LC834 树中距离之和 https://leetcode-cn.com/problems/sum-of-distances-in-tree
	// 下面的代码来自 积蓄程度 https://www.acwing.com/problem/content/289/ http://poj.org/problem?id=3585
	rerootDP := func(n int) { // 无根树
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

	// 树上所有路径的位运算与(&)的和
	// 单个点也算路径
	// 解法：对每一位，统计仅含 1 的路径个数
	// a[i] <= 2^20
	// https://ac.nowcoder.com/acm/contest/10167/C
	andPathSum := func(g [][]int, a []int) int64 {
		const mx = 21
		ans := int64(0)
		for i := 0; i < mx; i++ {
			cntOnePath := int64(0)
			var f func(v, fa int) int64
			f = func(v, fa int) int64 {
				one := int64(a[v] >> i & 1)
				cntOnePath += one
				for _, w := range g[v] {
					if w != fa {
						o := f(w, v)
						if one > 0 {
							cntOnePath += one * o
							one += o
						}
					}
				}
				return one
			}
			f(0, -1)
			ans += 1 << i * cntOnePath
		}

		{
			// 另一种做法是对每一位，用并查集求出 1 组成的连通分量，每个连通分量对答案的贡献是 sz*(sz+1)/2
			n := len(a)
			fa := make([]int, n)
			var find func(int) int
			find = func(x int) int {
				if fa[x] != x {
					fa[x] = find(fa[x])
				}
				return fa[x]
			}
			merge := func(from, to int) { fa[find(from)] = find(to) }

			ans := int64(0)
			for i := 0; i < mx; i++ {
				for j := range fa {
					fa[j] = j
				}
				sz := make([]int, n)
				for v, vs := range g {
					for _, w := range vs {
						if a[v]>>i&1 > 0 && a[w]>>i&1 > 0 {
							merge(v, w)
						}
					}
				}
				for j := 0; j < n; j++ {
					sz[find(j)]++
				}
				for j, f := range fa {
					if j == f && a[j]>>i&1 > 0 {
						ans += 1 << i * int64(sz[j]) * int64(sz[j]+1) / 2
					}
				}
			}
		}
		return ans
	}

	// 树上所有路径的位运算或(|)的和
	// 单个点也算路径
	// 做法和上面类似，求出仅含 0 的路径的个数，然后用路径总数 n*(n+1) 减去该个数就得到了包含至少一个 1 的路径个数
	// 也可以用并查集求出 0 组成的连通分量

	// 树上所有路径的位运算异或(^)的和
	// 原题失效了，只找到几个题解可以参考 https://www.cnblogs.com/kuronekonano/p/11135742.html https://blog.csdn.net/qq_36876305/article/details/80060491
	// 上面链接是边权，这里改成点权，且路径至少有两个点
	// 解法：由于任意路径异或和可以用从根节点出发的路径异或和表示
	// 对每一位，统计从根节点出发的路径异或和在该位上的 0 的个数和 1 的个数，
	// 只有当 0 与 1 异或时才对答案有贡献，所以贡献即为这两个个数之积
	xorPathSum := func(g [][]int, a []int) int64 {
		n := len(a)
		const mx = 30
		cnt := [mx]int{}
		xor := 0
		var f func(v, fa int)
		f = func(v, fa int) {
			xor ^= a[v]
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
				}
			}
			for i := 0; i < mx; i++ {
				cnt[i] += xor >> i & 1
			}
			xor ^= a[v]
		}
		f(0, -1)
		ans := int64(0)
		for i, c := range cnt {
			ans += 1 << i * int64(c) * int64(n-c)
		}
		return ans
	}

	// 树上所有路径的位运算异或(^)的异或和
	// 这里的路径至少有两个点
	// 方法是考虑每个点出现在多少条路径上，若数目为奇数则对答案有贡献
	// 路径分两种情况，一种是没有父节点参与的，树形 DP 一下就行了；另一种是父节点参与的，个数就是 子树*(n-子树)
	// https://ac.nowcoder.com/acm/contest/272/B
	xorPathXorSum := func(g [][]int, a []int) int {
		n := len(a)
		ans := 0
		var f func(v, fa int) int64
		f = func(v, fa int) int64 {
			cnt := int64(0)
			sz := int64(1)
			for _, w := range g[v] {
				if w != fa {
					s := f(w, v)
					cnt += sz * s
					sz += s
				}
			}
			cnt += sz * (int64(n) - sz)
			// 若一个点也算路径，那就再加一。或者在递归结束后把 ans^=a[0]^...^a[n-1]
			if cnt&1 > 0 {
				ans ^= a[v]
			}
			return sz
		}
		f(0, -1)
		return ans
	}

	_ = []interface{}{
		prefixSumDP, mapDP,
		maxSubArraySum, maxTwoSubArraySum, maxSubArrayAbsSum,
		minCostSorted,
		lcs, lcsPath, longestPalindromeSubsequence,
		lisSlow, lis, lisAll, lcis, lcisPath, countLIS, distinctSubsequence, minPalindromeCut,

		zeroOneKnapsack, zeroOneKnapsackAtLeastFillUp, zeroOneWaysToSum, zeroOneKnapsackLexicographicallySmallestResult,
		unboundedKnapsack, unboundedWaysToSum,
		boundedKnapsack, boundedKnapsackBinary,
		groupKnapsack,
		treeKnapsack,

		mergeStones,

		tsp,
		subsubDP,

		digitDP,
		kth666,

		cht,

		diameter, countDiameter, countVerticesOnDiameter,
		maxIndependentSetOfTree, minEdgeCoverOfTree, minDominatingSetOfTree, maxMatchingOfTree,
		rerootDP,
		andPathSum, xorPathSum, xorPathXorSum,
	}
}
