package copypasta

import (
	"container/heap"
	"math"
	"math/bits"
	"sort"
)

/* 动态规划

首先请透彻理解何为问题的「状态空间」，见 search.go 开头的注释

思考过程：
可以参考我的这篇题解：https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/solution/by-endlesscheng-pa3v/
1.1 如何把问题形式化为状态空间？（可以从边界、子集的角度思考）
1.2 子问题是如何重叠的？
1.3 子问题是怎么逐层递进的？（题目描述、隐含的顺序）
2.1 如何定义状态？需要用几个维度表示？
2.2 状态的范围是多少？起点状态和终点状态是什么？
2.3 哪些状态是相邻的？（即通过一次转移就能得到）
2.4 状态转移时要计算哪些内容？
2.5 对于转移来的相邻状态（入边），怎么决策？（简单的有取最值取和，复杂的有组合决策）
3.1 若复杂度过高，如何优化决策？
* 状态不好确定时，尝试转化问题模型、逆序思考、增加维度等等
* 对于计数问题或概率问题来说，状态定义和状态转移要做到不重不漏

如何设计状态
https://codeforces.com/problemset/problem/360/B
https://codeforces.com/problemset/problem/461/B
https://codeforces.com/problemset/problem/553/A
https://codeforces.com/problemset/problem/687/C
https://codeforces.com/problemset/problem/744/C
https://codeforces.com/problemset/problem/1012/C
https://codeforces.com/problemset/problem/1025/D
https://codeforces.com/problemset/problem/1027/E
https://codeforces.com/problemset/problem/1286/A
https://codeforces.com/problemset/problem/1408/D
https://atcoder.jp/contests/abc237/tasks/abc237_f
SEERC05，紫书例题 9-3，UVa 1347 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=446&page=show_problem&problem=4093
Daejeon11，紫书例题 9-8，UVa 1625 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4500
LC956 https://leetcode-cn.com/problems/tallest-billboard/ https://leetcode-cn.com/contest/weekly-contest-114/
涉及到相邻状态先后关系的 DP（喂兔子）https://codeforces.com/problemset/problem/358/D
戳气球 LC312 https://leetcode-cn.com/problems/burst-balloons/
消消乐 LC546 https://leetcode-cn.com/problems/remove-boxes/ https://leetcode.com/contest/leetcode-weekly-contest-25
混合逆序对 https://atcoder.jp/contests/arc097/tasks/arc097_c
https://codeforces.com/contest/1579/problem/G
todo https://atcoder.jp/contests/abc200/tasks/abc200_e
DI 序列的有效排列 LC903 https://leetcode.cn/problems/valid-permutations-for-di-sequence/

决策单调性
https://codeforces.com/problemset/problem/229/D

思维转换
谁来当 DP 对象 LC1434 https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/ https://leetcode-cn.com/contest/biweekly-contest-25/
扔蛋问题 LC887 https://leetcode-cn.com/problems/super-egg-drop/ https://www.bilibili.com/video/BV1KE41137PK https://leetcode-cn.com/contest/weekly-contest-97/
LC920* https://leetcode-cn.com/problems/number-of-music-playlists/ 注：官方题解给出了一种生成函数的做法
状态优化 https://codeforces.com/problemset/problem/838/E
「排序」题的转换 https://codeforces.com/problemset/problem/1223/D
https://codeforces.com/problemset/problem/1542/D
https://codeforces.com/problemset/problem/520/E
https://codeforces.com/problemset/problem/883/I
路径计数+推箱子 https://codeforces.com/problemset/problem/1225/E
找关键元素+状态机DP https://codeforces.com/problemset/problem/623/B

NOTE: 无后效性是指当前的决策只与过去的结果有关，而与过去的决策无关
NOTE: 若状态转移不构成 DAG，请尝试建图+BFS，见：
	https://ac.nowcoder.com/acm/contest/6218/B
	https://codeforces.com/problemset/problem/283/B 活用 012 染色
NOTE: 若使用滚动数组，注意复用时可能要初始化
NOTE:（区间 DP）正向计算不易时，试着反向计算
TIPS: 若转移是若干相邻项之和，可以考虑 f(p) - f(p-1) 的值，用滑动窗口来维护区间和，从而优化转移
      例题 LC837 https://leetcode-cn.com/problems/new-21-game/
递归打印路径：https://codeforces.com/problemset/problem/2/B
需要补充额外的状态 https://codeforces.com/problemset/problem/682/D

todo Non-trivial DP Tricks and Techniques https://codeforces.com/blog/entry/47764

贪心优化 DP
https://codeforces.com/problemset/problem/864/E

参考书籍推荐：
《算法竞赛进阶指南》- 介绍了大量且全面的 DP 内容，是目前市面上讲解 DP 最好的一本书

视频讲解：
https://www.bilibili.com/video/BV1gf4y1i78H 动态规划的套路 by wisdompeak
https://www.bilibili.com/video/av70148899 DP 入门，01 背包，完全背包，多重背包
https://www.bilibili.com/video/av77393700 LCS LIS
https://www.bilibili.com/video/av83939419 区间 DP
https://www.bilibili.com/video/av93356551 状态压缩 DP
https://www.bilibili.com/video/av98090640 树形 DP
https://www.bilibili.com/video/av85636122 动态规划 · 零 - Introduction
https://www.bilibili.com/video/av86983419 动态规划 · 一 - 序列型
https://www.bilibili.com/video/av89052674 动态规划 · 二 - 坐标、双序列、划分 & 状态压缩

套题/总结：
推荐 AtCoder 上的经典 DP 场 https://atcoder.jp/contests/dp
    题解 https://www.cnblogs.com/shanxieng/p/10232228.html
        https://codeforces.com/blog/entry/92170
        https://www.hamayanhamayan.com/entry/2019/01/12/163853
    讨论 https://codeforces.com/blog/entry/64250
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
CSES DP section editorial https://codeforces.com/blog/entry/70018
力扣上的 DP 问题
    分类汇总 https://zhuanlan.zhihu.com/p/126546914
    https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
    https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md
    https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems
    https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-w-5/
    https://leetcode-cn.com/tag/dynamic-programming/
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
func _(min, max func(int, int) int, abs func(int) int) {
	// 涉及到前缀和/子数组和的问题
	// 定义 dp[i] 表示前缀 a[:i] 中子数组和为 targetSum 的最短子数组长度
	// 下面的代码来自 LC1477 https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
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
			// 一维（多维见下）
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
			// 多维
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
	① 前缀/后缀之间的转移，例如从 dp[i-1] 转移到 dp[i]，或者从 dp[j] 转移到 dp[i]
	LC198 https://leetcode.cn/problems/house-robber/
	- 变形：恰好选 floor(n/2) 个 https://atcoder.jp/contests/abc162/tasks/abc162_f
	LC213 https://leetcode.cn/problems/house-robber-ii/
	- 相似题目 https://atcoder.jp/contests/abc251/tasks/abc251_e
	LC276 https://leetcode.cn/problems/paint-fence/
	LC368 https://leetcode.cn/problems/largest-divisible-subset/
	LC1105 https://leetcode.cn/problems/filling-bookcase-shelves/
	LC2369 https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/
	另见 LIS

	② 双序列问题，一般定义 dp[i][j] 表示对子问题 (s1[:i],s2[:j]) 的求解结果
	LC727 https://leetcode.cn/problems/minimum-window-subsequence/
	LC983 https://leetcode.cn/problems/minimum-cost-for-tickets/
	LC1639 https://leetcode.cn/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
	另见 LCS LPS

	③ 多维 / 额外状态
	LC123 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
	LC309 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
	LC920 https://leetcode-cn.com/problems/number-of-music-playlists/
	LC956 https://leetcode-cn.com/problems/tallest-billboard/ https://leetcode-cn.com/contest/weekly-contest-114/
	LC1186 https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
	LC1477 https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
	LC1531 看起来是区间 DP，仔细分析后是线性 DP https://leetcode-cn.com/contest/weekly-contest-199/problems/string-compression-ii/

	贪心+abs https://atcoder.jp/contests/abc163/tasks/abc163_e
	由 n 个值互不相同的点组成的高度不小于 h 的 BST 有多少个 https://codeforces.com/problemset/problem/9/D
	https://codeforces.com/problemset/problem/38/E
	好题：涉及到相邻状态先后关系的 DP（喂兔子） https://codeforces.com/problemset/problem/358/D
	https://codeforces.com/problemset/problem/446/A
	https://codeforces.com/problemset/problem/603/A
	处理区间元素不能在区间外面的技巧 https://codeforces.com/problemset/problem/811/C https://codeforces.com/contest/811/submission/174568255
	https://codeforces.com/problemset/problem/1120/C
	与 KMP 结合 https://codeforces.com/problemset/problem/1163/D
	https://codeforces.com/problemset/problem/1168/C
	https://codeforces.com/problemset/problem/1542/D
	*/

	// 网格路径问题
	// LC62 https://leetcode.cn/problems/unique-paths/
	// LC63 https://leetcode.cn/problems/unique-paths-ii/
	// LC64 https://leetcode.cn/problems/minimum-path-sum/
	// - 变形：连续性 & 上下界思想 https://codeforces.com/contest/1695/problem/C
	// LC120 https://leetcode.cn/problems/triangle/ https://www.luogu.com.cn/problem/P1216
	// LC931 https://leetcode.cn/problems/minimum-falling-path-sum/
	// LC2435 https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/

	// 最大子段和 https://www.luogu.com.cn/problem/P1115
	// 有两种思路
	// 1. 定义状态 dp[i] 表示以 a[i] 结尾的最大子段和，则有状态转移方程 dp[i]=max(dp[i−1],0)+a[i]，答案为 max(dp)
	// 2. 遍历 a 的同时维护前缀和的最小值，则遍历到 a[i] 时，当前最大子段和为 sum[i]-min(sum[j]), j<i
	// 算法导论 练习4.1-5
	// [题型总结] 关于最大子段和及其变式 https://www.luogu.com.cn/blog/wey-yzyl/zui-tai-zi-duan-hu-ji-ji-bian-shi-di-qi-shi
	// 子段长度有上限的最大子段和：见单调队列，题目为 https://ac.nowcoder.com/acm/contest/1006/D
	// 子段长度有下限的最大子段和：转换为前缀和之差 sum[i]-sum[j]，i-j>=K，维护 mi=min(sum[j])，同时更新 sum[i]-mi 的最大值（题目见 sort.go 中的 0-1 分数规划）
	// 子段和有上限的最大子段和：转换为前缀和之差 sum[i]-sum[j]<=K，在平衡树上二分 sum[j] LC363 https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/
	// 最大两段子段和：求每个位置上的前缀最大子段和和后缀最大子段和 https://www.luogu.com.cn/problem/P2642
	// - 等价题目：允许翻转一段子区间的最大子段和
	// 删除至多一个元素后的最大子段和 LC1186 https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
	// 最大 m 段子段和 http://acm.hdu.edu.cn/showproblem.php?pid=1024
	// 环状最大子段和：转换为 max(最大子段和, 总和减去最小子段和) LC918 https://leetcode-cn.com/problems/maximum-sum-circular-subarray/
	// 环状最大两段子段和：思路类似，注意取反后需要传入 a[1:n-1] https://www.luogu.com.cn/problem/P1121 https://ac.nowcoder.com/acm/contest/7738/B
	// 变形题 https://codeforces.com/problemset/problem/33/C
	//       https://codeforces.com/problemset/problem/788/A
	//       https://codeforces.com/problemset/problem/1155/D
	//       https://codeforces.com/problemset/problem/1197/D 思路 https://docs.qq.com/sheet/DWGFoRGVZRmxNaXFz 里面搜本题链接
	//       https://codeforces.com/problemset/problem/1373/D
	//       需要一些转换技巧 https://codeforces.com/problemset/problem/1082/E
	// 多个小数组合并 https://codeforces.com/problemset/problem/75/D
	//    这题做法需要用到上面说到的第二种思路
	maxSubarraySum := func(a []int) int {
		if len(a) == 0 { // 根据题意返回
			return 0
		}
		maxS, sum := a[0], a[0] // int64
		for _, v := range a[1:] {
			sum = max(sum, 0) + v
			maxS = max(maxS, sum)
		}
		if maxS < 0 { // 根据题意返回
			//return 0
		}
		return maxS
	}

	// 除了返回最大子段和外，还返回最大子段和对应的子段 [l,r]
	// https://codeforces.com/contest/1692/problem/H
	maxSubarraySumWithRange := func(a []int) (maxS, l, r int) {
		if len(a) == 0 { // 根据题意返回
			return 0, -1, -1
		}
		// int64
		maxS = a[0] // 注意 l 和 r 默认为 0，即 a[:1]
		for i, sum, st := 1, a[0], 0; i < len(a); i++ {
			if sum < 0 {
				sum, st = 0, i // 重新开始
			}
			sum += a[i]
			if sum > maxS {
				maxS, l, r = sum, st, i
			}
		}
		if maxS < 0 { // 根据题意返回
			//return 0, -1, -1
		}
		return
	}

	// 维护前缀和的最小值的写法
	// https://codeforces.com/contest/1692/problem/H
	maxSubarraySumWithRange = func(a []int) (maxS, l, r int) {
		if len(a) == 0 { // 根据题意返回
			return 0, -1, -1
		}
		// int64
		maxS = a[0] // 注意 l 和 r 默认为 0，即 a[:1]
		sum, minS, minI := 0, 0, -1
		for i, v := range a {
			sum += v
			if sum-minS > maxS {
				maxS, l, r = sum-minS, minI+1, i
			}
			if sum < minS {
				minS, minI = sum, i
			}
		}
		if maxS < 0 { // 根据题意返回
			//return 0, -1, -1
		}
		return
	}

	// 最大两段子段和（两段必须间隔至少 gap 个数）
	maxTwoSubarraySum := func(a []int, gap int) int {
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

	maxSubarrayAbsSum := func(a []int) int {
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

	// 最大子序列交替和（买卖股票）
	// 有两种思路：
	// 1. 动态规划，具体见我的题解 https://leetcode-cn.com/problems/maximum-alternating-subsequence-sum/solution/dong-tai-gui-hua-by-endlesscheng-d92a/
	// 2. 贪心，由于第一个值需要取正，将开头补上 0，就变成买卖股票问题了，只需关心波峰和波谷的值，即 ∑max(0,a[i+1]-a[i])
	// LC1911 https://leetcode-cn.com/problems/maximum-alternating-subsequence-sum/
	// LC122 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
	// 扩展：O(1) 回答交换其中两个元素后的最大子序列交替和 https://codeforces.com/problemset/problem/1420/C2
	maxAlternatingSumDP := func(a []int) int {
		dp := [2]int{0, -1e9} // int64
		for _, v := range a {
			dp = [2]int{max(dp[0], dp[1]-v), max(dp[1], dp[0]+v)}
		}
		return dp[1]
	}

	maxAlternatingSumGreedy := func(a []int) (ans int) {
		a = append([]int{0}, a...)
		for i := 1; i < len(a); i++ {
			ans += max(0, a[i]-a[i-1]) // int64
		}
		return
	}

	// 修改序列为非降或非增的最小修改次数
	// - 单次修改可以把某个数 +1 或 -1
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
	// 其他解释见 https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing/solution/by-gittauros-6x9v/
	// https://codeforces.com/problemset/problem/13/C
	// https://www.luogu.com.cn/problem/P4597
	// LC2263 https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing/
	// https://www.luogu.com.cn/problem/P2893
	// http://poj.org/problem?id=3666
	// https://codeforces.com/problemset/problem/713/C 严格单调递增 https://codeforces.com/blog/entry/47094?#comment-315161
	//     这道题做了一个 a[i]-=i 的操作（i 从 1 开始），把严格单调递增变成了非降的情况，从而可以应用该算法
	//     这一技巧的原理是，对于整数来说，单调递增的最小情况是 y=x+C，减去这一函数，就得到了非降序列的最小情况 y=C
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
	// 更快的做法（位运算）见 SPOJ LCS0 https://www.luogu.com.cn/problem/SP12076
	//
	// 模板题 LC1143 https://leetcode-cn.com/problems/longest-common-subsequence/
	// EXTRA: 最短公共超序列 (SCS) LC1092 https://leetcode-cn.com/problems/shortest-common-supersequence/
	// 变种 LC72   https://leetcode-cn.com/problems/edit-distance/
	//     LC97   https://leetcode-cn.com/problems/interleaving-string/
	//     LC115  https://leetcode-cn.com/problems/distinct-subsequences/
	//     LC583  https://leetcode-cn.com/problems/delete-operation-for-two-strings/
	//     LC712  https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/
	//     LC1035 https://leetcode-cn.com/problems/uncrossed-lines/
	//     LC1312 https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/ https://www.luogu.com.cn/problem/P1435
	//     https://atcoder.jp/contests/abc185/tasks/abc185_e
	//     其中一个改为子串 https://codeforces.com/problemset/problem/163/A
	//     https://codeforces.com/problemset/problem/1446/B
	// 多个排列的 LCS（转化成 DAG 最长路）https://codeforces.com/problemset/problem/463/D
	// 转换【巧妙】https://codeforces.com/problemset/problem/1114/D
	// 20多校第二场 https://acm.hdu.edu.cn/showproblem.php?pid=6774
	// 与 KMP 结合 https://codeforces.com/problemset/problem/346/B
	// 若其中一个序列无重复元素，可以转换成 LIS https://www.luogu.com.cn/problem/P1439 LC1713 https://leetcode-cn.com/contest/weekly-contest-222/problems/minimum-operations-to-make-a-subsequence/
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
	// 即 LCS(s, reverse(s))
	// LC516 https://leetcode-cn.com/problems/longest-palindromic-subsequence/
	// LC1216 https://leetcode-cn.com/problems/valid-palindrome-iii/
	// LC1246 https://leetcode.cn/problems/palindrome-removal/
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
	// 这种写法适用于一些定义比较复杂的变形题
	// O(n^2) - 定义 dp[i] 为以 a[i] 为末尾的 LIS 的长度
	//          可以把此问题想象成一个「跳跃游戏」，任选一个初始位置向右跳跃，每次只能跳到比当前位置更高的位置，问最多能跳多少次（最后答案加一）
	//          这样能更容易地看出转移的顺序，然后变成一个 DAG 上求最长路的问题
	// 转换 http://acm.hdu.edu.cn/showproblem.php?pid=1950
	// 转换 https://codeforces.com/problemset/problem/1562/E
	// 变体 https://codeforces.com/problemset/problem/1350/B
	//【网络流 24 题】能取出多少个长为 len(LIS) 的不相交子序列 https://loj.ac/p/6005 https://www.luogu.com.cn/problem/P2766
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
	// 方法一：二分
	// O(nlogn) - 定义 dp[i] 为长度为 i+1 的 LIS 末尾元素的最小值
	// 求下降，可以考虑取相反数
	// https://oi-wiki.org/dp/basic/#_12
	// 最小划分数 / 狄尔沃斯定理（Dilworth's theorem）https://en.wikipedia.org/wiki/Dilworth%27s_theorem
	//    偏序集的最少反链划分数等于最长链的长度
	// 随机排列 LIS 的长度期望 https://www.zhihu.com/question/266958886
	//
	// 例题 LC300 https://leetcode-cn.com/problems/longest-increasing-subsequence/
	// 建模 https://codeforces.com/problemset/problem/269/B
	// 经典转换（最多相交问题） https://codeforces.com/problemset/problem/67/D
	// 最小划分数（导弹拦截）https://www.luogu.com.cn/problem/P1020
	// 转化成最小划分数+打印划分方案 https://codeforces.com/problemset/problem/1296/E2
	// 合唱队形 https://www.luogu.com.cn/problem/P1091
	// 合唱队形（至少有升有降）https://leetcode-cn.com/contest/biweekly-contest-40/problems/minimum-number-of-removals-to-make-mountain-array/
	// 二维 LIS LC354 https://leetcode-cn.com/problems/russian-doll-envelopes/
	// 二维 LIS + 打印方案 http://codeforces.com/problemset/problem/4/D
	// 将所有元素分成三类：不在任何 LIS / 在至少一个 LIS / 在所有 LIS https://codeforces.com/problemset/problem/486/E
	// 重复 T 次的 LIS 问题 https://codeforces.com/problemset/problem/582/B
	// 若其中一个序列无重复元素，LCS 可以转换成 LIS https://www.luogu.com.cn/problem/P1439 LC1713 https://leetcode-cn.com/contest/weekly-contest-222/problems/minimum-operations-to-make-a-subsequence/
	// 在一维 LIS 的基础上，a[i] 可以从多个数中选一个，问 LIS 最长可以多长
	// - 思路：将各个 a[i] 的可选项从大到小排序，然后拼接成一个序列，求 LIS 即可（关键：从大到小排序避免了在同一个可选项中选择多个元素）
	// 图上的路径的 LIS https://codeforces.com/problemset/problem/960/F
	// LaIS 与单调栈结合 https://codeforces.com/problemset/problem/1468/A
	// 状态设计 LIS 计数 https://atcoder.jp/contests/abc237/tasks/abc237_f
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

	// 方法二：线段树优化 DP
	// https://www.acwing.com/problem/content/description/3665/
	// https://leetcode.cn/problems/longest-increasing-subsequence-ii/

	// 方法三：平衡树
	// todo 参考 https://leetcode.cn/problems/longest-increasing-subsequence-ii/solution/jianjie-by-xing-chen-26-ydqp/

	// 方法四：分治 + 单调队列
	// todo 参考 https://leetcode.cn/problems/longest-increasing-subsequence-ii/solution/fen-zhi-by-heltion-h31y/

	// 每个前缀的 LIS
	// https://leetcode-cn.com/contest/weekly-contest-253/problems/find-the-longest-valid-obstacle-course-at-each-position/
	lisAll := func(a []int) []int {
		n := len(a)
		lis := make([]int, n)
		dp := []int{}
		for i, v := range a {
			p := sort.SearchInts(dp, v) // 改成 v+1 为非降
			if p < len(dp) {
				dp[p] = v
			} else {
				dp = append(dp, v)
			}
			lis[i] = p + 1
		}
		return lis
	}

	// LIS 方案数 O(nlogn)
	// 原理见下面这题官方题解的方法二
	// LC673 https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
	cntLis := func(a []int) int {
		dp := [][]int{}  // 保留所有历史信息
		cnt := [][]int{} // 个数前缀和
		for _, v := range a {
			p := sort.Search(len(dp), func(i int) bool { return dp[i][len(dp[i])-1] >= v })
			c := 1
			if p > 0 {
				// 根据 dp[p-1] 来计算 cnt
				i := sort.Search(len(dp[p-1]), func(i int) bool { return dp[p-1][i] < v })
				c = cnt[p-1][len(cnt[p-1])-1] - cnt[p-1][i]
			}
			if p == len(dp) {
				dp = append(dp, []int{v})
				cnt = append(cnt, []int{0, c})
			} else {
				dp[p] = append(dp[p], v)
				cnt[p] = append(cnt[p], cnt[p][len(cnt[p])-1]+c)
			}
		}
		c := cnt[len(cnt)-1]
		return c[len(c)-1]
	}

	// LIS 相关构造题
	// https://codeforces.com/problemset/problem/1304/D
	// https://atcoder.jp/contests/arc091/tasks/arc091_c

	// 最大上升子序列和
	// 按值从小到大排序，值相同的下标从大到小排序
	// 然后用树状数组或线段树：单点更新，维护前缀最大值
	// https://www.acwing.com/problem/content/3665/

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

	// 本质不同非空子序列个数
	// 详细讲解见 https://leetcode.cn/problems/distinct-subsequences-ii/solution/xi-fen-wen-ti-fu-za-du-you-hua-pythonjav-1ihu/
	// 模板题 LC940 https://leetcode-cn.com/problems/distinct-subsequences-ii/
	// 倒序遍历即可 LC1987 https://leetcode-cn.com/problems/number-of-unique-good-subsequences/
	// 需要一点构造能力 https://codeforces.com/problemset/problem/645/E
	distinctSubsequence := func(s string) int {
		const mod int = 1e9 + 7
		f := [26]int{}
		sumF := 0
		for _, b := range s {
			b -= 'a'
			tmp := (sumF + mod - f[b]) % mod
			f[b] = (sumF + 1) % mod
			sumF = (tmp + f[b]) % mod
		}
		// 把空的也算上
		//sumF = (sumF + 1) % mod
		return sumF
	}

	// 滚动数组写法
	distinctSubsequence = func(s string) int {
		const mod int = 1e9 + 7
		last := make([]int, 26)
		dp := 1
		for _, v := range s {
			v -= 'a'
			res := dp - last[v]
			if res < 0 {
				res += mod
			}
			dp = (dp + res) % mod
			last[v] = (last[v] + res) % mod
		}
		return (dp + mod - 1) % mod // 去掉空序列
	}

	// 多重组合
	// todo 挑战

	// 多重排列
	// dp[i][j] 表示前 i 类数字组成长为 j 的排列个数
	// dp[i][j] = ∑dp[i-1][k]*C(j,k), 0<=k<=min(j,cnt[i])
	// 边界 dp[0][0] = 1
	// todo https://atcoder.jp/contests/abc234/tasks/abc234_f

	// 回文串最小分割次数
	// 紫书例题 9-7，UVa 11584 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=27&page=show_problem&problem=2631
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
	// 转换 LC494 https://leetcode.cn/problems/target-sum/
	//            https://atcoder.jp/contests/abc274/tasks/abc274_d
	// 转换 LC1049 https://leetcode-cn.com/problems/last-stone-weight-ii/
	// 转换 https://codeforces.com/problemset/problem/1381/B
	// 转换 https://atcoder.jp/contests/dp/tasks/dp_x
	// 排序+转换 https://codeforces.com/problemset/problem/1203/F2
	// 转移对象是下标 https://codeforces.com/edu/course/2/lesson/9/3/practice/contest/307094/problem/I
	// - dp[i][j] 表示前 i 个数，凑成 j 的所有方案中，最小下标的最大值
	// 转移对象是下标 https://codeforces.com/problemset/problem/981/E
	// 打印方案 https://codeforces.com/problemset/problem/864/E
	// NOIP06·提高 金明的预算方案（也可以用树上背包做）https://www.luogu.com.cn/problem/P1064
	// EXTRA: 恰好装满（相当于方案数不为 0）LC416 https://leetcode-cn.com/problems/partition-equal-subset-sum/
	//        必须定义成恰好装满（紫书例题 9-5，UVa 12563）https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=441&page=show_problem&problem=4008
	// EXTRA: 背包容量为 0 https://codeforces.com/problemset/problem/366/C
	// EXTRA: 二维费用 https://www.acwing.com/problem/content/8/ https://www.luogu.com.cn/problem/P1507 LC474 https://leetcode-cn.com/problems/ones-and-zeroes/
	// EXTRA: 把一个维度转换成 DP 的定义 https://codeforces.com/problemset/problem/837/D
	// EXTRA: 离散化背包 https://codeforces.com/contest/366/submission/61452111
	zeroOneKnapsack := func(values, weights []int, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for i, w := range weights {
			v := values[i]
			// 这里 j 的初始值可以优化成前 i 个物品的重量之和（但不能超过 maxW）
			for j := maxW; j >= w; j-- {
				dp[j] = max(dp[j], dp[j-w]+v)
			}
		}
		return dp[maxW]
	}

	// 0-1 背包 EXTRA: 恰好装满
	// https://leetcode.cn/contest/sf-tech/problems/cINqyA/
	zeroOneKnapsackExactlyFull := func(values, weights []int, maxW int) {
		dp := make([]int, maxW+1) // int64
		for i := range dp {
			dp[i] = -1e9 // -1e18
		}
		dp[0] = 0
		for i, w := range weights {
			v := values[i]
			for j := maxW; j >= w; j-- {
				dp[j] = max(dp[j], dp[j-w]+v)
			}
		}
		for i := maxW; i >= 0; i-- {
			if dp[i] >= 0 { // 能恰好装满 i，此时背包物品价值和的最大值是 dp[i]
				// ...
			}
		}
	}

	// 0-1 背包 EXTRA: 至少装入重量和为 maxW 的物品，求价值和的最小值 https://www.luogu.com.cn/problem/P4377
	// 需要一点转换 https://codeforces.com/problemset/problem/19/B
	// 二维费用的情况+价值最小 https://ac.nowcoder.com/acm/contest/6218/C
	zeroOneKnapsackAtLeastFillUp := func(values, weights []int, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for i := range dp {
			dp[i] = 1e9 // 1e18
		}
		dp[0] = 0
		for i, v := range values {
			w := weights[i]
			for j := maxW; j >= 0; j-- {
				dp[j] = min(dp[j], dp[max(j-w, 0)]+v) // max(j-w, 0) 蕴含了「至少」
			}
		}

		{
			// 另一种写法
			for i, v := range values {
				w := weights[i]
				for j := maxW; j >= 0; j-- {
					k := min(j+w, maxW)
					dp[k] = min(dp[k], dp[j]+v)
				}
			}
		}

		return dp[maxW]
	}

	// 0-1 背包 EXTRA: 从序列 a 中选若干个数，使其总和为 sum 的方案数
	// NOTE: 1,1,1,...1(32个1),s-32,s-31,...,s 可以让方案数恰好为 2^32
	// 二维+上限+下限 LC879 https://leetcode-cn.com/contest/weekly-contest-95/problems/profitable-schemes/
	// 转换 https://atcoder.jp/contests/abc169/tasks/abc169_f
	// 转换 https://codeforces.com/problemset/problem/478/D
	// 转换 LC494 https://leetcode-cn.com/problems/target-sum/
	// 转换 LC1434 https://leetcode-cn.com/problems/number-of-ways-to-wear-different-hats-to-each-other/
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
				fa[i][j] = j // 注意：<w 的转移来源也要标上！
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
			if fa[i][j] == j { // &&  weights[i] > 0      考虑重量为 0 的情况，必须都选上
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
	// 适用于背包容量很大，但是物品价值不高的情况
	// 把重量看成价值，价值看成重量，求同等价值下能得到的最小重量，若该最小重量不超过背包容量，则该价值合法。所有合法价值的最大值即为答案
	// 时间复杂度 O(n * sum(values)) 或 O(n^2 * maxV)
	// https://atcoder.jp/contests/dp/tasks/dp_e
	// https://codeforces.com/contest/1650/problem/F
	zeroOneKnapsackByValue := func(values, weights []int, maxW int) int {
		totValue := 0
		for _, v := range values {
			totValue += v
		}
		dp := make([]int, totValue+1) // int64
		for i := range dp {
			dp[i] = 1e18
		}
		dp[0] = 0
		totValue = 0
		for i, v := range values {
			w := weights[i]
			totValue += v
			for j := totValue; j >= v; j-- {
				dp[j] = min(dp[j], dp[j-v]+w)
			}
		}
		for i := totValue; ; i-- {
			if dp[i] <= maxW {
				return i
			}
		}
	}

	// 完全背包
	// 转换 LC322 https://leetcode-cn.com/problems/coin-change/
	// EXTRA: 恰好装满+打印方案 LC1449 https://leetcode-cn.com/contest/biweekly-contest-26/problems/form-largest-integer-with-digits-that-add-up-to-target/
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

	// 多重背包计数（可以用前缀和优化）
	// https://www.luogu.com.cn/problem/P1077

	// 多重背包 - 未优化
	// 转换（价值主导）（由于要取 min 所以不能用二进制优化）https://codeforces.com/problemset/problem/922/E
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
	// todo 打印方案
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

	// 分组背包·每组至多选一个（恰好选一个见后面）
	// https://www.acwing.com/problem/content/9/
	// https://www.luogu.com.cn/problem/P1757
	// LC2218 https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/
	// https://codeforces.com/problemset/problem/148/E
	// todo 进一步优化 https://codeforces.com/problemset/problem/1442/D
	type item struct{ v, w int }
	groupKnapsack := func(groups [][]item, maxW int) int {
		dp := make([]int, maxW+1) // int64
		for _, g := range groups {
			// 这里 j 的初始值可以优化成前 i 个组的每组最大重量之和（但不能超过 maxW）
			for j := maxW; j >= 0; j-- {
				for _, it := range g {
					if v, w := it.v, it.w; w <= j {
						dp[j] = max(dp[j], dp[j-w]+v) // 如果 it.w 可能为 0 则需要用 dp[2][] 来滚动（或者保证每组至多一个 0 且 0 在该组最前面）
					}
				}
			}
		}
		return dp[maxW]
	}

	// 分组背包·每组恰好选一个
	// 允许物品重量为 0
	// https://atcoder.jp/contests/abc240/tasks/abc240_c
	// LC1981 https://leetcode-cn.com/problems/minimize-the-difference-between-target-and-chosen-elements/
	// 与二分图染色结合 https://codeforces.com/problemset/problem/1354/E
	// 转换 https://codeforces.com/problemset/problem/1637/D
	groupKnapsackFill := func(groups [][]int, maxW int) []bool {
		dp := make([]bool, maxW+1) // dp[i][j] 表示能否从前 i 组物品中选出重量恰好为 j 的，且每组都恰好选一个物品
		dp[0] = true
		for _, g := range groups {
		next:
			for j := maxW; j >= 0; j-- { // 这里 j 的初始值可以优化至前 i 组的最大元素值之和
				for _, w := range g {
					if w <= j && dp[j-w] {
						dp[j] = true
						continue next
					}
				}
				dp[j] = false // 由于我们是滚动数组的写法，dp[i][j] 无法满足时要标记成 false
			}
		}
		return dp // dp[j] 表示从每组恰好选一个，能否凑成重量 j
	}

	// 树上背包/树形背包/依赖背包
	// todo 树上背包的上下界优化 https://ouuan.github.io/post/%E6%A0%91%E4%B8%8A%E8%83%8C%E5%8C%85%E7%9A%84%E4%B8%8A%E4%B8%8B%E7%95%8C%E4%BC%98%E5%8C%96/
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
	// NOIP06·提高 金明的预算方案 https://www.luogu.com.cn/problem/P1064
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
	LC410 https://leetcode.cn/problems/split-array-largest-sum/
	LC813 https://leetcode.cn/problems/largest-sum-of-averages/
	LC1278 https://leetcode.cn/problems/palindrome-partitioning-iii/
	LC1335 https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/

	② 求解关于某个序列的最优性质，要求大区间的最优解可以依赖于小区间的最优解
	一般定义 dp[i][j] 表示 a[i:j] 的最优解
	此时可以枚举区间大小和区间左端点，从小区间转移到大区间
	LC375 https://leetcode.cn/problems/guess-number-higher-or-lower-ii/
	戳气球（好题）LC312 https://leetcode.cn/problems/burst-balloons/
	移除盒子（状态定义和转移的好题）LC546 https://leetcode.cn/problems/remove-boxes/
	打印机（好题）LC664 https://leetcode.cn/problems/strange-printer/
	最优三角剖分 LC1039 https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/
	插入形成回文 LC1312 https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/ https://www.luogu.com.cn/problem/P1435
	另见 LPS

	染色【套路】https://codeforces.com/problemset/problem/1114/D
	同色消除【套路】https://www.luogu.com.cn/problem/P4170 https://codeforces.com/problemset/problem/1132/F
	回文消除 https://codeforces.com/problemset/problem/607/B
	二维区间 DP https://codeforces.com/problemset/problem/1198/D
	状态设计的好题 https://codeforces.com/problemset/problem/1025/D
	https://blog.csdn.net/weixin_43914593/article/details/106163859 算法竞赛专题解析（14）：DP应用--区间DP
	todo https://atcoder.jp/contests/abc159/tasks/abc159_f
	     https://codeforces.com/problemset/problem/245/H
	*/

	// 石子合并
	// https://atcoder.jp/contests/dp/tasks/dp_n
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

	// 统计区间内回文串个数
	// 返回一个二维数组 dp, dp[i][j] 表示 [i,j] 内的回文串的个数
	// https://codeforces.com/problemset/problem/245/H
	countPalindromes := func(s string) [][]int {
		n := len(s)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
			dp[i][i] = 1
			if i+1 < n && s[i] == s[i+1] {
				dp[i][i+1] = 1
			}
		}
		for i := n - 3; i >= 0; i-- {
			for j := i + 2; j < n; j++ {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
		// 到这里为止，dp[i][j] = 1 表示 s[i:j+1] 是回文串
		for i := n - 2; i >= 0; i-- {
			for j := i + 1; j < n; j++ {
				dp[i][j] += dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1] // 容斥
			}
		}
		return dp
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

	/* 概率 DP / 期望 DP
	https://oi-wiki.org/dp/probability/
	https://en.wikipedia.org/wiki/Probability
	https://en.wikipedia.org/wiki/Expected_value
	https://en.wikipedia.org/wiki/Variance
	https://en.wikipedia.org/wiki/Optional_stopping_theorem
	todo https://codeforces.com/blog/entry/62690
	     https://codeforces.com/blog/entry/62792
	 https://www.luogu.com.cn/blog/Troverld/gai-shuai-ji-wang-xue-xi-bi-ji
	 一类概率期望问题的杀器：势函数和鞅的停时定理 https://www.cnblogs.com/TinyWong/p/12887591.html https://codeforces.com/blog/entry/87598 最后一题
	 鞅与停时定理学习笔记 https://www.luogu.com.cn/blog/gxy001/yang-yu-ting-shi-ding-li-xue-xi-bi-ji

	期望的可加性
	https://zhidao.baidu.com/question/259203053.html

	马尔可夫链 Markov chain https://en.wikipedia.org/wiki/Markov_chain
	吸收马尔可夫链 Absorbing Markov chain https://en.wikipedia.org/wiki/Absorbing_Markov_chain
	https://www.bilibili.com/video/BV14y4y1S7ve

	一个比较有用的公式（应用：CF1623D）
	E(x) = ∑i*P(x=i) = ∑P(x>=i)

	方差
	σ²(x) = sum(x²)/n - (sum(x)/n)²

	概率
	http://codeforces.com/problemset/problem/16/E
	https://codeforces.com/problemset/problem/540/D
	https://codeforces.com/problemset/problem/678/E

	期望
	https://codeforces.com/problemset/problem/235/B
	https://codeforces.com/problemset/problem/908/D
	https://codeforces.com/problemset/problem/1097/D
	https://codeforces.com/problemset/problem/1623/D
	https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4dea
	https://leetcode.cn/contest/ubiquant2022/problems/I3Gm2h/
	*/

	/* 状压 DP
	常用于处理包含排列的问题等
	NOTE: 若问题无法划分成小问题，必须考虑各种可能的情况，则可能是 NP 完全问题
	浅谈状压 DP https://www.luogu.com.cn/blog/yijan/zhuang-ya-dp
	https://blog.csdn.net/weixin_43914593/article/details/106432695 算法竞赛专题解析（15）：DP应用--状态压缩DP

	todo 题单 https://ac.nowcoder.com/acm/problem/collection/808
	     题单 https://ac.nowcoder.com/acm/problem/collection/810
	 LC691  https://leetcode-cn.com/problems/stickers-to-spell-word/
	 LC1125 https://leetcode-cn.com/problems/smallest-sufficient-team/
	 LC943  https://leetcode-cn.com/problems/find-the-shortest-superstring/
	https://www.luogu.com.cn/problem/P1879
	循环移位 https://codeforces.com/contest/1209/problem/E2
	https://codeforces.com/problemset/problem/401/D
	与质因子分解结合 https://codeforces.com/problemset/problem/453/B
	与排序贪心结合 https://codeforces.com/problemset/problem/1316/E
	与概率 DP 结合 https://codeforces.com/problemset/problem/16/E

	todo 汉密尔顿路径/回路 Hamiltonian path
	https://en.wikipedia.org/wiki/Hamiltonian_path
	https://en.wikipedia.org/wiki/Hamiltonian_path_problem

	求最大团/最大独立集的记忆化写法见 graph.go
	*/

	// 任意排列 DP
	// https://atcoder.jp/contests/dp/tasks/dp_o
	// https://atcoder.jp/contests/abc199/tasks/abc199_e
	// https://codeforces.com/problemset/problem/1215/E
	// 状态设计 https://codeforces.com/problemset/problem/743/E
	// 状态设计 https://codeforces.com/problemset/problem/744/C
	// 枚举来源 https://codeforces.com/problemset/problem/377/C
	// 卡常优化 https://codeforces.com/problemset/problem/327/E 另一种做法是折半枚举
	// LC1879 https://leetcode-cn.com/problems/minimum-xor-sum-of-two-arrays/
	// LC2172 https://leetcode-cn.com/problems/maximum-and-sum-of-array/
	permDP := func(a []int) int {
		n := len(a)
		m := 1 << n
		dp := make([]int, m) // int64
		dp[0] = 1
		for s, dv := range dp {
			i := bits.OnesCount(uint(s))
			v := a[i]
			for t, lb := m-1^s, 0; t > 0; t ^= lb {
				lb = t & -t
				ss := s | lb
				j := bits.TrailingZeros(uint(lb))
				w := a[j]
				_ = v + w
				// dp[ss] <- dv
				dp[ss] += dv // mod
			}
		}
		return dp[m-1]
	}

	// 旅行商问题 (TSP)
	// 图论中的一个等价形式是：给定一个加权完全图（顶点表示城市，边表示道路，权重是道路的距离），求一权值和最小的哈密尔顿回路。
	// 返回一个 ans 数组，ans[i] 表示从 st 出发，访问完所有位置且最后停在 i 的最短路径（注意可能要特判 i==st 的情况）
	// 做法：定义 dp[s][i] 表示已访问的集合为 s，最后一个访问的位置是 i 时的最小花费
	//      则有 dp[s|1<<j][j] = min(dp[s|1<<j][j], dp[s][i]+dist[i][j])
	//      枚举 i 和 j 时可以用 TrailingZeros 来直接枚举每个 1 和 0 的位置
	// https://en.wikipedia.org/wiki/Travelling_salesman_problem
	// https://en.wikipedia.org/wiki/Hamiltonian_path HCP
	// 模板题 https://www.luogu.com.cn/problem/P1171 https://www.luogu.com.cn/problem/P1433 https://www.acwing.com/problem/content/93/
	// 略微变形 吃加速药水 https://atcoder.jp/contests/abc274/tasks/abc274_e
	// https://codeforces.com/problemset/problem/1185/G1
	// LC847 https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/
	// 恰好访问 m 个点 https://codeforces.com/contest/580/problem/D
	// 建模转换题 LC943 https://leetcode-cn.com/problems/find-the-shortest-superstring/
	//          LCP13 https://leetcode-cn.com/problems/xun-bao/
	// 变体+打印路径 https://codeforces.com/problemset/problem/8/C
	// 矩阵交换行问题 https://codeforces.com/problemset/problem/1102/F
	// EXTRA: 固定起点终点的问题，视问题情况有两种方法：
	//        添加一个节点 https://stackoverflow.com/questions/14527815/how-to-fix-the-start-and-end-points-in-travelling-salesmen-problem
	//        设置距离 https://stackoverflow.com/questions/36086406/traveling-salesman-tsp-with-set-start-and-end-point
	tsp := func(dist [][]int, st int) []int {
		const inf int = 1e9 // 1e18
		dp := make([][]int, 1<<len(dist))
		for i := range dp {
			dp[i] = make([]int, len(dist))
			for j := range dp[i] {
				dp[i][j] = inf
			}
		}
		dp[1<<st][st] = 0 // 多个起点的话就设置多个 dp[1<<st[i]][st[i]] = 0
		for s, dr := range dp {
			// 利用位运算快速求出 s 中 1 的位置 i，以及 s 中 0 的位置 j（通过 s 的补集中的 1 的位置求出）
			for ss := uint(s); ss > 0; ss &= ss - 1 {
				i := bits.TrailingZeros(ss)
				for t, lb := len(dp)-1^s, 0; t > 0; t ^= lb {
					lb = t & -t
					j := bits.TrailingZeros(uint(lb))
					dp[s|lb][j] = min(dp[s|lb][j], dr[i]+dist[i][j])
				}
			}
		}
		return dp[len(dp)-1]
	}

	// 无向图简单环数量
	// https://blog.csdn.net/fangzhenpeng/article/details/49078233
	// https://codeforces.com/problemset/problem/11/D
	countCycle := func(g [][]int, n, m int) int64 {
		ans := int64(0)
		// 取集合 s 的最小值作为起点
		dp := make([][]int64, 1<<n)
		for i := range dp {
			dp[i] = make([]int64, n)
		}
		for i := 0; i < n; i++ {
			dp[1<<i][i] = 1
		}
		for s := range dp {
			for v, dv := range dp[s] {
				if dv == 0 {
					continue
				}
				for _, w := range g[v] {
					if 1<<w < s&-s {
						continue
					}
					if 1<<w&s == 0 {
						dp[s|1<<w][w] += dv
					} else if 1<<w == s&-s {
						ans += dv
					}
				}
			}
		}
		return ans - int64(m)/2
	}

	// 枚举子集的子集
	// 复杂度 O(3^n)，证明：元素个数为 k 的集合有 C(n,k) 个，其子集有 2^k 个，故有 ∑C(n,k)*2^k = (2+1)^n = 3^n
	// 例如：dp[set] = max{dp[set^sub] + sum of sub} for all valid sub
	// 模板题 https://atcoder.jp/contests/dp/tasks/dp_u
	// 模板题 LC1986 https://leetcode-cn.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/
	// LC1494 https://leetcode-cn.com/contest/biweekly-contest-29/problems/parallel-courses-ii/
	// LC1654 https://leetcode-cn.com/contest/biweekly-contest-39/problems/distribute-repeating-integers/
	// 训练指南第一章例题 29，UVa11825 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=226&page=show_problem&problem=2925
	// - 将 n 个集合分成尽量多组，使得对于每组，组内所有集合的并集等于全集
	// 训练指南第一章例题 32，WF10，UVa1099 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=245&page=show_problem&problem=3540
	subsubDP := func(a []int) int {
		n := len(a)
		m := 1 << n
		// 预处理每个子集的子集和
		sum := make([]int, m)
		for i := range sum {
			for s := uint(i); s > 0; s &= s - 1 {
				sum[i] += a[bits.TrailingZeros(s)]
			}
		}
		dp := make([]int, m)
		for s, dv := range dp {
			t := m - 1 ^ s
			// 枚举补集的非空子集
			for sub := t; sub > 0; sub = (sub - 1) & t {
				ss := s | sub
				dp[ss] = max(dp[ss], dv+sum[sub])
			}
		}
		return dp[m-1]
	}

	// 上面的记忆化写法
	subsubDPMemo := func(a []int) int {
		n := len(a)
		m := 1 << n
		// 预处理每个子集的子集和
		sum := make([]int, m)
		for i := range sum {
			for s := uint(i); s > 0; s &= s - 1 {
				sum[i] += a[bits.TrailingZeros(s)]
			}
		}
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
				r := f(p+1, set^sub)
				res = min(res, r+sum[sub])
				sub = (sub - 1) & set
			}

			// 所有非空子集
			for sub := set; sub > 0; sub = (sub - 1) & set {

			}

			return
		}
		return f(0, 1<<m-1)
	}

	// 高维前缀和 SOS DP (Sum over Subsets)
	// 给一个集合，对该集合的所有子集，计算该子集的所有子集之和（这个「和」不一定是加法，可以是其它的满足合并性质的统计量）
	// https://codeforces.com/blog/entry/45223
	// 大量习题 https://blog.csdn.net/weixin_38686780/article/details/100109753
	//
	// https://codeforces.com/problemset/problem/1234/F
	//    求满足 ai&aj=0 的 ai|aj 的二进制 1 的个数的最大值
	//    思路是转换成求每个 ai 的补集的 SOS，维护子集二进制 1 的个数的最大值
	// https://www.hackerearth.com/zh/problem/algorithm/special-pairs-5-3ee6b3fe-3d8a1606/
	//    求 ai&aj=0 的 (i,j) 对数，0<=ai<=1e6
	//    思路和上面类似，转换成求每个 ai 的补集的 SOS
	//    注：另一种解法是求 FWT(cnt)[0]
	// 转换成求集合中最大次大 https://atcoder.jp/contests/arc100/tasks/arc100_c
	// 求下标最大次大，且不需要在乎 k 的上限的写法 https://codeforces.com/problemset/problem/1554/B
	// https://codeforces.com/problemset/problem/165/E
	// 容斥 https://codeforces.com/problemset/problem/449/D
	// todo https://codeforces.com/problemset/problem/1208/F
	//  https://codeforces.com/problemset/problem/800/D
	//  https://codeforces.com/problemset/problem/383/E
	//  https://www.luogu.com.cn/problem/P6442
	// https://codeforces.com/problemset/problem/1523/D
	sos := func(a []int) []int {
		// 从子集转移的写法
		const mx = 20 // bits.Len(uint(max(a))
		dp := make([]int, 1<<mx)
		for _, v := range a {
			dp[v]++
		}
		for i := 0; i < mx; i++ {
			for s := 0; s < 1<<mx; s++ {
				s |= 1 << i
				// 将 s 的子集 s^1<<i 的统计量合并到 s 中
				dp[s] += dp[s^1<<i]
			}
		}

		{
			// 从超集转移的写法
			for i := 0; i < mx; i++ {
				for s := 1<<mx - 1; s >= 0; s-- {
					if s>>i&1 == 0 {
						dp[s] += dp[s|1<<i]
					}
				}
			}
		}

		{
			// 维护集合最大和次大的写法
			type pair struct{ fi, se int }
			dp := make([]pair, 1<<mx)
			for i := 0; i < mx; i++ {
				for s := 0; s < 1<<mx; s++ {
					s |= 1 << i
					p, q := dp[s], dp[s^1<<i]
					if q.se > p.fi {
						dp[s] = q
					} else if q.fi > p.fi {
						dp[s] = pair{q.fi, p.fi}
					} else if q.fi > p.se {
						dp[s].se = q.fi
					}
				}
			}
		}

		return dp
	}

	/* 插头 DP（Plug DP）/ 轮廓线 DP（Broken Profile DP）
	轮廓线：已决策格子和未决策格子的分界线
	陈丹琦《基于连通性状态压缩的动态规划问题》
	训练指南 6.1
	https://oi-wiki.org/dp/plug/（花絮 https://zhuanlan.zhihu.com/p/133761303）
	https://cp-algorithms.com/dynamic_programming/profile-dynamics.html
	图解 https://www.luogu.com.cn/blog/GNAQ/solution-p5056
	An Introduction to Plug DP https://codeforces.com/blog/entry/90841
	todo 大量题目 https://www.luogu.com.cn/blog/efforts-will-pay-off/du-liu-dong-gui-cha-tou-dp
	模板题 https://www.luogu.com.cn/problem/P5056
	https://www.luogu.com.cn/problem/P3272
	https://www.luogu.com.cn/problem/P5074
	https://www.luogu.com.cn/problem/P1933
	*/
	plugDP := func(grids [][]byte) int {
		n, m := len(grids), len(grids[0])
		var endX, endY int
		g := make([][]bool, n)
		for i, row := range grids {
			g[i] = make([]bool, m)
			for j, b := range row {
				if b == '.' {
					g[i][j] = true
					endX, endY = i, j
				}
			}
		}

		// 四进制状态
		get := func(s, k int) int { return s >> (k << 1) & 3 }
		set := func(k, v int) int { return 1 << (k << 1) * v }

		ans := 0 // int64
		dp := map[int]int{0: 1}
		for i, row := range g {
			tmp := dp
			dp = make(map[int]int, len(tmp))
			for s, dv := range tmp {
				dp[s<<2] = dv // 轮廓线移动到当前行
			}
			for j, empty := range row {
				tmp := dp
				dp = make(map[int]int, len(tmp))
				for s, dv := range tmp {
					switch x, y := get(s, j), get(s, j+1); {
					case !empty: // 障碍格
						if x == 0 && y == 0 { // 空
							dp[s] += dv
						}
					case x == 0 && y == 0: // ┌ 单独形成一对括号
						if j+1 < m && row[j+1] && i+1 < n && g[i+1][j] {
							dp[s|set(j, 1)|set(j+1, 2)] += dv
						}
					case x == 0 && y > 0:
						if j+1 < m && row[j+1] { // └
							dp[s] += dv
						}
						if i+1 < n && g[i+1][j] { // │
							dp[s|set(j, y)^set(j+1, y)] += dv
						}
					case x > 0 && y == 0:
						if j+1 < m && row[j+1] { // ─
							dp[s^set(j, x)|set(j+1, x)] += dv
						}
						if i+1 < n && g[i+1][j] { // ┐
							dp[s] += dv
						}
					case x == 1 && y == 1: // ┘ 消去 x 和 y，并找到和 y 匹配的右括号，将其改成左括号
						// 注：这里和下边的 k 的位置可以事先预处理出来
						for k, c := j+2, 1; ; k++ {
							if t := get(s, k); t == 1 {
								c++
							} else if t == 2 {
								if c--; c == 0 {
									dp[s^set(j, x)^set(j+1, y)^set(k, 3)] += dv // 将 2 改成 1 要异或 3
									break
								}
							}
						}
					case x == 2 && y == 2: // ┘ 消去 x 和 y，并找到和 x 匹配的左括号，将其改成右括号
						for k, c := j-1, 1; ; k-- {
							if t := get(s, k); t == 2 {
								c++
							} else if t == 1 {
								if c--; c == 0 {
									dp[s^set(j, x)^set(j+1, y)^set(k, 3)] += dv // 将 1 改成 2 要异或 3
									break
								}
							}
						}
					case x == 2 && y == 1: // ┘ 消去右括号和左括号，连接两个插头
						dp[s^set(j, x)^set(j+1, y)] += dv
					default: // ┘ x == 1 && y == 2
						// 此时封闭整个路径，这只应当发生在最后一个合法格子上
						if i == endX && j == endY {
							ans += dv
						}
					}
				}
			}
		}
		// 若需要取模则中间记得取模（若结果不大可以仅在循环结束时取模）
		return ans
	}

	/* 数位 DP

	一般用 dp[i][j] 表示当前在第 i 位，前面维护了一个为 j 的值，且后面的位数可以随便选时的数字个数
	在解释状态的含义时，网上的很多文章都漏了「后面的位数可以随便选」这个约束，只有加上这个约束，我们才能根据「是否紧贴上界」来完成相对应的代码逻辑

	视频讲解，从 19:30 开始 https://www.bilibili.com/video/BV1rS4y1s721
	https://zhuanlan.zhihu.com/p/348851463
	https://www.bilibili.com/video/BV1MT4y1376C
	https://www.bilibili.com/video/BV1yT4y1u7jW

	入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
	      https://atcoder.jp/contests/dp/tasks/dp_s
	      https://codeforces.com/problemset/problem/1036/C
	二进制 1 的个数恰为 k 的数字个数 https://codeforces.com/problemset/problem/431/D https://www.acwing.com/problem/content/1083/
	是 m 的倍数且偶数位为 d 且奇数位不为 d 的数字个数 https://codeforces.com/problemset/problem/628/D
	所有数字均出现偶数次的数字个数 https://codeforces.com/problemset/problem/855/E
	相邻数字约束 SC09 https://www.luogu.com.cn/problem/P2657
	数位统计
	LC233 https://leetcode-cn.com/problems/number-of-digit-one/
	      https://leetcode-cn.com/problems/number-of-2s-in-range-lcci/
	      http://acm.hdu.edu.cn/showproblem.php?pid=3555
	      http://acm.hdu.edu.cn/showproblem.php?pid=2089
	      LC1067 https://leetcode.cn/problems/digit-count-in-range/
	LC600 二进制不含连续 1 的数字个数 https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/
	LC788 含有某些特殊数字 https://leetcode.cn/problems/rotated-digits/
	LC902 最大为 N 的数字组合 https://leetcode-cn.com/contest/weekly-contest-101/problems/numbers-at-most-n-given-digit-set/
	LC1012 有重复数字的数字个数 https://leetcode-cn.com/contest/weekly-contest-128/problems/numbers-with-repeated-digits/
	- LC2376 互补问题 无重复数字的数字个数 https://leetcode.cn/contest/weekly-contest-306/problems/count-special-integers/
	LC1397 与 KMP 结合 https://leetcode-cn.com/contest/weekly-contest-182/problems/find-all-good-strings/
	                          https://www.luogu.com.cn/problem/P3193
	digsum(n)|n 的数的个数 https://www.luogu.com.cn/problem/P4127 https://www.acwing.com/problem/content/313/
	http://acm.hdu.edu.cn/showproblem.php?pid=3886
	http://acm.hdu.edu.cn/showproblem.php?pid=6796
	todo LC248 中心对称数个数 https://leetcode.cn/problems/strobogrammatic-number-iii/
	     LC1088 互补 https://leetcode.cn/problems/confusing-number-ii/
	注：一些第 k 小的题目需要与二分结合，或者用试填法（见后面的 kth666）
	todo 套题 https://www.luogu.com.cn/blog/s-r-f/oi-bi-ji-shuo-wei-dp-ge-ji-dui-shuo-wei-dp-di-yi-dian-li-xie
	todo 套题 https://codeforces.com/blog/entry/53960
	*/
	digitDP := func(lower, upper string, sumUpper int) int64 {
		const mod int64 = 1e9 + 7

		// 返回 <=s 的符合要求的字符串数目
		// TIPS: 某些情况下思考补集会更加容易，即求不符合要求的字符串数目
		calc := func(s string) int64 {
			const lowerC, upperC byte = '0', '9'
			dp := make([][]int64, len(s))
			for i := range dp {
				dp[i] = make([]int64, sumUpper+1)
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, sum int, limitUp bool) int64
			f = func(p, sum int, limitUp bool) (res int64) {
				if p == len(s) {
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
				for ch := lowerC; ch <= up; ch++ {
					tmp := sum

					cnt := f(p+1, tmp, limitUp && ch == up)
					res = (res + cnt) % mod
				}
				return
			}
			res := f(0, 0, true)
			return res
		}
		ansLower := calc(lower) // lower-1
		ansUpper := calc(upper)
		ans := ansUpper - ansLower
		// lower 是否算上
		//if lowerIsValid {
		//	ans++
		//}
		ans = (ans%mod + mod) % mod

		// TIPS: 对于需要判断/禁止前导零的情况，可以加一个额外的维度 fill，表示已经填入了数字（没有前导零的合法状态），最后 p>=n 的时候可以根据情况返回 1 或者 0
		// 例如 https://codeforces.com/contest/855/submission/125651587
		// 以下代码以 https://www.luogu.com.cn/problem/P2657 为例
		calc = func(s string) int64 {
			dp := make([][10]int64, len(s))
			for i := range dp {
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, pre int, limitUp, fill bool) int64
			f = func(p, pre int, limitUp, fill bool) (res int64) {
				if p == len(s) {
					return 1
				}
				if !limitUp && fill { // 注意这里的判断
					dv := &dp[p][pre]
					if *dv >= 0 {
						return *dv
					}
					defer func() { *dv = res }()
				}
				up := 9
				if limitUp {
					up = int(s[p] & 15)
				}
				for d := 0; d <= up; d++ {
					if !fill || abs(d-pre) > 1 {
						res += f(p+1, d, limitUp && d == up, fill || d > 0)
					}
				}
				return
			}
			return f(0, 0, true, false)
		}

		// 若需要计算的不是合法数字个数，而是合法数字之和，则需要在计算时考虑单个数位的贡献
		// 以下代码以 https://codeforces.com/problemset/problem/1073/E 为例
		calcSum := func(s string, k int) int64 {
			n := len(s)
			type pair struct{ cnt, sum int64 }
			dp := make([][1 << 10]pair, n)
			for i := range dp {
				for j := range dp[i] {
					dp[i][j] = pair{-1, -1}
				}
			}
			var f func(int, uint16, bool, bool) pair
			f = func(p int, mask uint16, limitUp, fill bool) (res pair) {
				if p == n {
					if !fill {
						return
					}
					return pair{1, 0}
				}
				if !limitUp && fill {
					dv := &dp[p][mask]
					if dv.cnt >= 0 {
						return *dv
					}
					defer func() { *dv = res }()
				}
				up := 9
				if limitUp {
					up = int(s[p] & 15)
				}
				for d := 0; d <= up; d++ {
					tmp := mask
					if fill || d > 0 {
						tmp |= 1 << d
					}
					if bits.OnesCount16(tmp) <= k {
						pr := f(p+1, tmp, limitUp && d == up, fill || d > 0)
						res.cnt = (res.cnt + pr.cnt) % mod
						res.sum = (res.sum + int64(math.Pow10(n-1-p))%mod*pr.cnt%mod*int64(d) + pr.sum) % mod
					}
				}
				return
			}
			return f(0, 0, true, false).sum
		}
		_ = calcSum

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
	模板题 https://codeforces.com/problemset/problem/1175/E
	      https://codeforces.com/problemset/problem/1516/D
	      https://atcoder.jp/contests/arc060/tasks/arc060_c
	开车旅行 https://www.luogu.com.cn/problem/P1081
	计算重复 https://www.acwing.com/problem/content/296/
	*/
	binaryLifting := func(segs, qs []struct{ l, r int }) []int {
		// 以 CF1175E 为例
		const mx = 19
		f := make([][mx]int, 5e5+1)
		for _, s := range segs {
			l, r := s.l, s.r
			f[l][0] = max(f[l][0], r)
		}
		// 前缀最大值（最右）
		for i := 1; i < len(f); i++ {
			f[i][0] = max(f[i][0], f[i-1][0])
		}
		// 倍增
		for i := 0; i+1 < mx; i++ {
			for p := range f {
				f[p][i+1] = f[f[p][i]][i]
			}
		}

		ans := make([]int, len(qs))
		for qi, q := range qs {
			l, r := q.l, q.r
			res := 0
			for i := mx - 1; i >= 0; i-- {
				if f[l][i] < r {
					l = f[l][i]
					res |= 1 << i
				}
			}
			if f[l][0] >= r {
				ans[qi] = res + 1
			} else {
				ans[qi] = -1
			}
		}
		return ans
	}

	/* 数据结构优化 DP
	长为 k 的上升子序列个数 https://codeforces.com/problemset/problem/597/C
	Lazy 线段树 https://atcoder.jp/contests/dp/tasks/dp_w
	https://atcoder.jp/contests/arc073/tasks/arc073_d https://www.luogu.com.cn/problem/T190609?contestId=48376 https://www.luogu.com.cn/blog/abruce-home/ti-xie-nao-zhong
	https://codeforces.com/problemset/problem/66/E
	https://codeforces.com/problemset?order=BY_RATING_ASC&tags=data+structures%2Cdp

	动态 DP
	https://oi-wiki.org/dp/dynamic/
	线段树维护转移矩阵 https://www.cnblogs.com/Miracevin/p/9124511.html
	1. 先不考虑修改，不考虑区间，直接列出整个区间的 DP 转移。
	2. 列出转移矩阵。由于有很多修改操作，我们将数据集中在一起处理，还可以利用矩阵结合律，并且区间比较好提取（找一段矩阵就好了），修改也方便。
	3. 线段树维护矩阵。对于修改，我们就是在矩阵上进行修改。对于不同的题目，我们要用不同的修改方式和记录手段，但都是用线段树维护对应区间内的矩阵信息。如矩阵乘积，矩阵和等等。线段树的区间操作可以应对区间修改问题。
	https://codeforces.com/problemset/problem/1380/F
	https://codeforces.com/problemset/problem/718/C
	https://codeforces.com/problemset/problem/750/E
	https://codeforces.com/problemset/problem/1149/B
	*/

	// 单调队列优化
	// 见 monotone_queue.go

	// 斜率优化 / 凸包优化 (Convex Hull Trick, CHT)
	//
	// 若状态转移方程具有类似于 dp[i] = min{dp[j]-a[i]*b[j]}, j<i 的形式，方程中包含一个 i 和 j 的乘积项，且序列 a 和 b 均单调递增
	// 若将 (b[j],dp[j]) 看作二维平面上的点，则 dp[i] 就是所有斜率为 a[i] 且过其中一点的直线中，与 y 轴的最小截距
	// 我们可以用一个单调队列来维护 (b[j],dp[j]) 的相邻点所构成的下凸包
	// 对于斜率 a[i]，我们需要在队列中寻找一个位置 k，其左侧斜率小于 a[i]，右侧斜率大于 a[i]，此时经过点 (b[k],dp[k]) 能取到最小截距
	//
	// 具体到实现，设两转移来源的下标为 j 和 k，若 k < j 且 dp[k]-a[i]*b[k] < dp[j]-a[i]*b[j]
	// 则有 (dp[j]-dp[k])/(b[j]-b[k]) > a[i]
	// 据此式，用单调队列来维护斜率（下凸包）
	// 转移前，在单调队列中找到斜率 a[i] 的对应位置，然后代入转移方程，求出 dp[i]
	// 转移后，将点 (b[i],dp[i]) 加入单调队列中
	//
	// https://oi-wiki.org/dp/opt/slope/
	// https://cp-algorithms.com/geometry/convex_hull_trick.html
	// https://www.luogu.com.cn/blog/ChenXingLing/post-xue-xi-bi-ji-dong-tai-gui-hua-xie-shuai-you-hua-dp-chao-yang-x
	// https://www.luogu.com.cn/blog/ningago-lsh/xie-lv-you-hua-dp
	// https://blog.csdn.net/weixin_43914593/article/details/105560357 算法竞赛专题解析（12）：DP优化(2)--斜率(凸壳)优化
	// https://zhuanlan.zhihu.com/p/558522044
	// https://zhuanlan.zhihu.com/p/363772434
	// https://codeforces.com/blog/entry/63823
	//
	// https://codeforces.com/problemset/problem/319/C
	// https://www.luogu.com.cn/problem/P2365 https://www.luogu.com.cn/problem/P5785 http://poj.org/problem?id=1180
	// todo https://atcoder.jp/contests/dp/tasks/dp_z
	// todo https://www.luogu.com.cn/problem/P2900
	//  https://www.luogu.com.cn/problem/P3195 https://loj.ac/p/10188
	//  http://poj.org/problem?id=3709
	//  https://codeforces.com/problemset/problem/311/B
	//  https://codeforces.com/problemset/problem/1715/E
	cht := func(a, b []int64) int64 {
		n := len(a)
		dp := make([]int64, n)
		// 计算两点间的斜率，若分子分母均在 32 位整数范围内，可以去掉浮点，改用乘法
		slope := func(i, j int) float64 {
			if b[i] == b[j] { // 若保证不相等则去掉
				if dp[j] > dp[i] {
					return 1e99
				}
				return -1e99
			}
			return float64(dp[j]-dp[i]) / float64(b[j]-b[i])
		}
		q := []int{0}
		for i := 1; i < n; i++ {
			k := a[i]
			// 若斜率 k 随 i 单调递增，则可以直接将单调队列中小于 k 的斜率弹出
			for len(q) > 1 && slope(q[0], q[1]) < float64(k) {
				q = q[1:]
			}
			// ……之后，队首 q[0] 就是最优决策的下标
			j := q[0]
			// ……否则，需要在单调队列中二分得到最优决策的位置
			j = sort.Search(len(q)-1, func(j int) bool { return slope(j, j+1) > float64(k) })

			// 转移
			dp[i] = dp[j] - a[i]*b[j]

			// 然后，将点 (b[i],dp[i]) 加入单调队列中
			for len(q) > 1 && slope(q[len(q)-1], i) < slope(q[len(q)-2], q[len(q)-1]) {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
		return dp[n-1]
	}

	// 凸优化 DP / 带权二分 / WQS 二分
	// 《浅析一类二分方法》
	// 把强制选 k 个物品的问题转换成选任意个物品的问题
	// todo https://www.luogu.com.cn/blog/daniu/wqs-er-fen
	//      https://www.luogu.com.cn/blog/Flying2018/wqs-er-fen-min-ke-fu-si-ji-hu-xue-xi-bi-ji
	// todo https://www.cnblogs.com/CreeperLKF/p/9045491.html
	// todo https://www.luogu.com.cn/blog/juruoforever/wqs-er-fen-qian-xi
	// todo https://taodaling.github.io/blog/2020/07/31/WQS%E4%BA%8C%E5%88%86/
	//
	// todo https://leetcode-cn.com/problems/minimum-white-tiles-after-covering-with-carpets/solution/wqs-er-fen-on-log-n-by-zerotrac2-cp7j/
	//
	// http://codeforces.com/problemset/problem/739/E（这题还可以费用流）
	// IOI00 邮局 https://www.luogu.com.cn/problem/P4767
	// LC188 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/yi-chong-ji-yu-wqs-er-fen-de-you-xiu-zuo-x36r/
	// https://www.luogu.com.cn/problem/U72600
	// https://www.luogu.com.cn/training/3495#problems
	// 单度限制最小生成树（恰好）https://codeforces.com/problemset/problem/125/E

	// 四边形不等式优化 Knuth's Optimization
	// https://oi-wiki.org/dp/opt/quadrangle/
	// https://jeffreyxiao.me/blog/knuths-optimization
	// todo https://blog.csdn.net/weixin_43914593/article/details/105150937 算法竞赛专题解析（10）：DP优化(1)--四边形不等式
	//      决策单调性优化讲解 https://www.luogu.com.cn/blog/83547/zong-dong-tai-gui-hua-di-ben-zhi-kan-si-bian-xing-fou-deng-shi-you-hua
	// 扔蛋问题 LC887 https://leetcode-cn.com/problems/super-egg-drop/

	/* 树形 DP
	一般是从自底向上计算的，也就是根据子树返回值来计算父节点的值
	也有自顶向下的写法，见后面

	https://blog.csdn.net/weixin_43914593/article/details/107145592
	https://codeforces.com/blog/entry/20935
	https://codeforces.com/blog/entry/63257

	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp%2Ctrees
	todo 题单 https://ac.nowcoder.com/acm/problem/collection/807
	     题单 https://ac.nowcoder.com/acm/problem/collection/809
	https://codeforces.com/problemset/problem/743/D
	https://codeforces.com/problemset/problem/855/C
	https://codeforces.com/problemset/problem/982/C
	https://codeforces.com/problemset/problem/1083/A
	好题 http://codeforces.com/problemset/problem/1249/F
	好题 https://codeforces.com/problemset/problem/1453/E
	如何定义状态 https://codeforces.com/problemset/problem/461/B
	**如何转移 https://codeforces.com/problemset/problem/538/E
	可以重复走 https://codeforces.com/problemset/problem/1220/E
	巧妙的转换 https://codeforces.com/problemset/problem/734/E
	https://codeforces.com/problemset/problem/1292/C

	自顶向下
	https://leetcode.cn/problems/U7WvvU/ 题解 https://leetcode.cn/problems/U7WvvU/solution/shu-xing-dp-by-endlesscheng-isuo/
	*/

	// 树的直径（两遍 DFS 求法另见 graph_tree.go 中的 diameter）
	// LC1245 https://leetcode-cn.com/problems/tree-diameter/
	// 变形 LC2246 https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/
	// 变形 https://codeforces.com/problemset/problem/1238/F
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

	// 最大路径点权和
	maxPathSum := func(st int, g [][]int, a []int) (maxPathSum int) {
		var f func(v, fa int) int
		f = func(v, fa int) (maxSum int) {
			rootVal := a[v]
			maxPathSum = max(maxPathSum, rootVal)
			for _, w := range g[v] {
				if w != fa {
					sum := f(w, v)
					maxPathSum = max(maxPathSum, maxSum+sum+rootVal)
					maxSum = max(maxSum, sum)
				}
			}
			return maxSum + rootVal
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
	// 方案是否唯一 Tehran06，紫书例题 9-13，UVa 1220 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=247&page=show_problem&problem=3661
	maxIndependentSetOfTree := func(n int, g [][]int, a []int) int { // 无根树
		var f func(int, int) (notChosen, chosen int)
		f = func(v, fa int) (notChosen, chosen int) { // int64
			chosen = a[v] // 1
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

	// 树上最小顶点覆盖
	// 代码和树上最大独立集类似
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	// 经典题：战略游戏 https://www.luogu.com.cn/problem/P2016
	// 训练指南第一章例题 30，UVa10859 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1800
	// - 求最小顶点覆盖，以及所有最小顶点覆盖中，两端点都被覆盖的边的最大个数
	// 构造 https://codeforces.com/problemset/problem/959/C
	minVertexCoverOfTree := func(n int, g [][]int, a []int) int { // 无根树
		var f func(int, int) (notChosen, chosen int)
		f = func(v, fa int) (notChosen, chosen int) { // int64
			chosen = a[v] // 1
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

	// 树上最小支配集
	// 返回最小点权和（最小支配集的情形即所有点权均为一）
	// 下面的定义省去了（……时的最小支配集的元素个数）   w 为 i 的儿子
	// dp[i][0]：i 属于支配集 = a[i]+∑min(dp[w][0],dp[w][1],dp[w][2])
	// dp[i][1]：i 不属于支配集，且被儿子支配 = ∑min(dp[w][0],dp[w][1]) + 如果全选 dp[w][1] 则补上 min{dp[w][0]-dp[w][1]}
	// dp[i][2]：i 不属于支配集，且被父亲支配 = ∑min(dp[w][0],dp[w][1])
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	//
	// 保安站岗 https://www.luogu.com.cn/problem/P2458
	// 手机网络 https://www.luogu.com.cn/problem/P2899
	// https://ac.nowcoder.com/acm/problem/24953
	// 监控二叉树 LC968 https://leetcode-cn.com/problems/binary-tree-cameras/
	// todo EXTRA: 消防局的设立（支配距离为 2） https://www.luogu.com.cn/problem/P2279
	// todo EXTRA: 将军令（支配距离为 k） https://www.luogu.com.cn/problem/P3942
	//                                https://atcoder.jp/contests/arc116/tasks/arc116_e
	minDominatingSetOfTree := func(n int, g [][]int, a []int) int { // 无根树
		const inf int = 1e9 // 1e18
		var f func(int, int) (chosen, bySon, byFa int)
		f = func(v, fa int) (chosen, bySon, byFa int) { // int64
			chosen = a[v] // 1
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

	// EXTRA: 每个被支配的点，仅被一个点支配
	// Kaoshiung06，紫书例题 9-14，UVa 1218 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=247&page=show_problem&problem=3659

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

	// 树上完美匹配子集个数
	// https://codeforces.com/problemset/problem/1032/F

	// todo 给一棵树和树上的一些关键节点，选 m 个点，使得关键节点到这些点中距离的最小值的最大值最小，求这个值
	//      https://www.luogu.com.cn/problem/P3523

	// 换根 DP / 二次扫描法
	// 进阶指南 p.292-295
	// https://codeforces.com/blog/entry/20935
	//
	// https://www.luogu.com.cn/problem/P3478
	// https://www.luogu.com.cn/problem/P2986
	// https://codeforces.com/problemset/problem/763/A（有更巧妙的做法）
	// https://codeforces.com/problemset/problem/1092/F
	// https://codeforces.com/problemset/problem/219/D
	// https://codeforces.com/problemset/problem/337/D
	// 注意不存在逆元的情形 https://codeforces.com/problemset/problem/543/D
	// https://codeforces.com/problemset/problem/1626/E

	// 给一颗无根树
	// 返回每个点到其余点的距离之和
	// LC834 https://leetcode-cn.com/problems/sum-of-distances-in-tree
	// 任意两点距离除以 k 的上取整之和 https://codeforces.com/problemset/problem/791/D
	sumOfDistancesInTree := func(g [][]int) []int {
		n := len(g)
		size := make([]int, n)
		var f func(int, int) int // int64
		f = func(v, fa int) (sum int) { // sum 表示以 0 为根时的子树 v 中的节点到 v 的距离之和
			size[v] = 1
			for _, w := range g[v] {
				if w != fa {
					sum += f(w, v) + size[w] // 子树 w 的每个节点都要经过 w-v，因此这条边对 sum 产生的贡献为 size[w]
					size[v] += size[w]
				}
			}
			return
		}
		sum0 := f(0, -1)

		ans := make([]int, n)
		var reroot func(v, fa, sum int)
		reroot = func(v, fa, sum int) {
			ans[v] = sum
			for _, w := range g[v] {
				if w != fa {
					// 换根后，离子树 w 中的所有节点近了 1，又离不在子树 w 中的节点远了 1
					// 所以要减去 sz[w]，并加上 n-size[w]
					reroot(w, v, sum+n-size[w]*2)
				}
			}
		}
		reroot(0, -1, sum0)
		return ans
	}

	// 积蓄程度 https://www.acwing.com/problem/content/289/ http://poj.org/problem?id=3585
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
		maxSubarraySum, maxSubarraySumWithRange, maxTwoSubarraySum, maxSubarrayAbsSum,
		maxAlternatingSumDP, maxAlternatingSumGreedy,
		minCostSorted,
		lcs, lcsPath, longestPalindromeSubsequence,
		lisSlow, lis, lisAll, cntLis, lcis, lcisPath, countLIS,
		distinctSubsequence,
		minPalindromeCut,

		zeroOneKnapsack, zeroOneKnapsackExactlyFull, zeroOneKnapsackAtLeastFillUp, zeroOneWaysToSum, zeroOneKnapsackLexicographicallySmallestResult, zeroOneKnapsackByValue,
		unboundedKnapsack, unboundedWaysToSum,
		boundedKnapsack, boundedKnapsackBinary,
		groupKnapsack, groupKnapsackFill,
		treeKnapsack,

		mergeStones, countPalindromes,

		permDP, tsp, countCycle, subsubDP, subsubDPMemo, sos, plugDP,

		digitDP, kth666,

		binaryLifting,

		cht,

		diameter, countDiameter, countVerticesOnDiameter, maxPathSum,
		maxIndependentSetOfTree, minVertexCoverOfTree, minDominatingSetOfTree, maxMatchingOfTree,
		sumOfDistancesInTree, rerootDP,
		andPathSum, xorPathSum, xorPathXorSum,
	}
}
