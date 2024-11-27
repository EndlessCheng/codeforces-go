package copypasta

/* 单调队列 Monotone Queue

视频讲解 
https://www.bilibili.com/video/BV1bM411X72E/

两张图秒懂单调队列
https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/solution/liang-zhang-tu-miao-dong-dan-diao-dui-li-9fvh/

需要不断维护队列的单调性，时刻保证队列元素从大到小或从小到大
前置知识：双指针
以固定窗口大小的区间最大值为例（此时维护的是一个从大到小的单调队列）：
每次向右移动一格左指针，在移动前，如果左指针指向的元素在队首左侧，说明左指针指向的元素小于队首，移动左指针不会改变区间最大值，直接移动左指针即可；
如果左指针指向的就是队首，那么移动左指针会使区间最大值变小（变为单调队列队首之后的那个元素），我们要弹出队首。
这样无论是何种情况，都保证了在移动左指针后，单调队列的队首始终为当前区间的最大值。
https://oi-wiki.org/ds/monotonous-queue/
https://oi-wiki.org/dp/opt/monotonous-queue-stack/
https://cp-algorithms.com/data_structures/stack_queue_modification.html
https://blog.csdn.net/weixin_43914593/article/details/105791217 算法竞赛专题解析（13）：DP优化(3)--单调队列优化
todo https://xyzl.blog.luogu.org/DQ-OP-DP

单调队列求最值等
- [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)
- [LCR 184. 设计自助结算系统](https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/)
- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/) 1672
- [2762. 不间断子数组](https://leetcode.cn/problems/continuous-subarrays/) 同 1438 题
- [2398. 预算内的最多机器人数目](https://leetcode.cn/problems/maximum-number-of-robots-within-budget/) 1917
- [862. 和至少为 K 的最短子数组](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/) 2307
- [1499. 满足不等式的最大值](https://leetcode.cn/problems/max-value-of-equation/) 2456
https://codeforces.com/problemset/problem/6/E 1900

注：如果要维护区间次小值（次大值），可以用 last 记录上一个被 pop 的元素，然后分类讨论：
- 如果队列中至少有两个元素，那么取队首后面的元素为次小值
- 否则，次小值为 last
- 注意 last 的位置在队首左边，需要注意是否小于左边界（或者保证区间长度至少为 2）

单调双端队列
https://codeforces.com/problemset/problem/1584/E 2300 

单调队列优化 DP
一般用来维护区间最值
0. 前提：区间右端点变大时，左端点也在变大（同滑动窗口）
1. 转移前，去掉队首无用数据
2. 计算转移（直接从队首转移）
3. 把 f[i] 插入队尾前，去掉队尾无用数据

- [2944. 购买水果需要的最少金币数](https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/) 1709 可以用单调队列优化到 $\mathcal{O}(n)$
- [1696. 跳跃游戏 VI](https://leetcode.cn/problems/jump-game-vi/) 1954
- [1425. 带限制的子序列和](https://leetcode.cn/problems/constrained-subsequence-sum/) 2032
- [375. 猜数字大小 II](https://leetcode.cn/problems/guess-number-higher-or-lower-ii/) 可以用单调队列优化到 $\mathcal{O}(n^2)$
      https://leetcode.cn/problems/guess-number-higher-or-lower-ii/solution/cong-ji-yi-hua-sou-suo-on3-dao-dong-tai-q13g9/
- [1687. 从仓库到码头运输箱子](https://leetcode.cn/problems/delivering-boxes-from-storage-to-ports/) 2610
- [3117. 划分数组得到最小的值之和](https://leetcode.cn/problems/minimum-sum-of-values-by-dividing-array/) 2735
- [2945. 找到最大非递减数组的长度](https://leetcode.cn/problems/find-maximum-non-decreasing-array-length/) 2943
      https://www.luogu.com.cn/problem/P4954 https://www.luogu.com.cn/problem/P2300
      https://iai.sh.cn/problem/565
   - 反过来 https://codeforces.com/problemset/problem/1898/B 1500
https://codeforces.com/problemset/problem/1736/C1 1300
https://codeforces.com/problemset/problem/1918/D 1900
https://codeforces.com/problemset/problem/487/B 2000
https://codeforces.com/problemset/problem/940/E 2000
https://codeforces.com/problemset/problem/1237/D 2000
https://codeforces.com/problemset/problem/1077/F2 2100
https://codeforces.com/problemset/problem/1216/F 2100
https://codeforces.com/problemset/problem/797/F 2600
https://www.luogu.com.cn/problem/P3572
https://www.luogu.com.cn/problem/P3594
https://www.luogu.com.cn/problem/P3957
https://www.luogu.com.cn/problem/P1419
todo https://www.luogu.com.cn/problem/P2627
 http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=1070
 老鼠进洞 https://codeforces.com/problemset/problem/797/F
*/
type MqData struct {
	Val int
	Del int // 懒删除标记
}

type MonotoneQueue struct {
	Data []MqData
	Size int // 单调队列对应的区间的长度
}

func (mq MonotoneQueue) Less(a, b MqData) bool {
	return a.Val >= b.Val // >= 维护区间最大值；<= 维护区间最小值
}

func (mq *MonotoneQueue) Push(v int) {
	mq.Size++
	d := MqData{v, 1}
	for len(mq.Data) > 0 && mq.Less(d, mq.Data[len(mq.Data)-1]) {
		d.Del += mq.Data[len(mq.Data)-1].Del
		mq.Data = mq.Data[:len(mq.Data)-1]
	}
	mq.Data = append(mq.Data, d)
}

func (mq *MonotoneQueue) Pop() {
	mq.Size--
	if mq.Data[0].Del > 1 {
		mq.Data[0].Del--
	} else {
		mq.Data = mq.Data[1:]
	}
}

// 返回区间最值
// 调用前需保证 mq.size > 0
func (mq MonotoneQueue) Top() int {
	return mq.Data[0].Val
}

// 滑动窗口最值（固定区间大小的区间最值）
// LC239 https://leetcode.cn/problems/sliding-window-maximum/
// https://www.luogu.com.cn/problem/P1886 http://poj.org/problem?id=2823
// https://codeforces.com/problemset/problem/940/E
// https://codeforces.com/problemset/problem/372/C（另一种做法是用堆）
// 贡献+差分数组 https://codeforces.com/problemset/problem/1208/E
func fixedSizeMax(a []int, fixedSize int) []int {
	n := len(a)
	q := MonotoneQueue{} // 最大/最小由 less 来控制
	ans := make([]int, 0, n-fixedSize+1)
	for i, v := range a {
		q.Push(v)
		if q.Size > fixedSize {
			q.Pop()
		}
		// 插入新元素并保证单调队列大小后，获取区间最值
		if i+1 >= fixedSize {
			ans = append(ans, q.Top())
		}
	}
	return ans
}

// 子数组长度不超过 sizeLimit 的最大子数组和
// 用单调队列维护前缀和的最小值，循环时保证单调队列对应的区间长度不超过 sizeLimit
// https://www.acwing.com/problem/content/137/ https://ac.nowcoder.com/acm/contest/1006/D
func maxSubSumWithLimitSize(a []int, sizeLimit int) int {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	ans := int(-1e9)     // -1e18
	q := MonotoneQueue{} // 维护区间最小值
	q.Push(sum[0])
	for r := 1; r <= n; r++ {
		if q.Size > sizeLimit {
			q.Pop()
		}
		ans = max(ans, sum[r]-q.Top())
		q.Push(sum[r])
	}
	return ans
}

// 子数组和至少为 k 的最短非空子数组长度
// 转换成两个前缀和的差至少为 k
// 这题的关键在于，当右端点向右（枚举）时，左端点是绝对不会向左的（因为向左肯定会比当前求出的最短长度要长）
// 想明白这一点就可以愉快地使用单调队列了
// LC862 https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
func shortestSubSumAtLeastK(a []int, k int) int {
	n := len(a)
	ans := n + 1
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	q := MonotoneQueue{} // 维护区间最小值
	q.Push(sum[0])
	for r := 1; r <= n; r++ {
		for q.Size > 0 && sum[r]-q.Top() >= k {
			ans = min(ans, q.Size)
			q.Pop()
		}
		q.Push(sum[r])
	}
	if ans > n {
		return -1
	}
	return ans
}

// 对每个右端点，求最远的左端点，满足这一区间内的最大值减最小值不超过 limit
// 求这个的同时，用单调队列维护 DP https://codeforces.com/problemset/problem/487/B
// 完整代码（传入 less 的写法）https://codeforces.com/contest/487/submission/121388184
func leftPosInDiffLimit(a []int, limit int) []int {
	posL := make([]int, len(a))
	small := MonotoneQueue{} // 最小值
	big := MonotoneQueue{}   // 最大值
	for i, v := range a {
		small.Push(v)
		big.Push(v)
		for big.Top()-small.Top() > limit {
			small.Pop()
			big.Pop()
		}
		posL[i] = i + 1 - small.Size // 通过 size 求出左端点位置
	}
	return posL
}

// 枚举区间左端点更为方便的情况 · 其一
// 统计区间个数：区间最大值 >= 2*区间最小值
// https://ac.nowcoder.com/acm/contest/6778/C
//
// 思路：转变成求「区间最大值 < 2*区间最小值」的区间个数
// 随着左端点向右，右端点必然不会向左
func countSubarrayByMinMax(a []int) int {
	n := len(a)
	ans := n * (n + 1) / 2
	mx := MonotoneQueue{} // 维护区间最大值
	mn := MonotoneQueue{} // 维护区间最小值（需要新定义一个有不同 less 的 monotoneQueue）
	for i, j := 0, 0; i < n; i++ {
		// 确保符合条件再插入
		for ; j < n && (mx.Size == 0 || mn.Size == 0 || max(mx.Top(), a[j]) < 2*min(mn.Top(), a[j])); j++ {
			mx.Push(a[j])
			mn.Push(a[j])
		}
		sz := j - i
		ans -= sz
		// 若单调队列指向的区间的左端点为 i，则对应元素在下一次循环时将不再使用。故弹出之
		if mx.Size == sz {
			mx.Pop()
		}
		if mn.Size == sz {
			mn.Pop()
		}
	}
	return ans
}

// 枚举区间左端点更为方便的情况 · 其二
// https://codeforces.com/problemset/problem/1237/D
// 注意这题和 countSubarrayByMinMax 的不同之处：不满足要求的最小值一定要在最大值的右侧
// 也可以枚举右端点，见 https://www.luogu.com.cn/blog/qianshang/solution-cf1237d
func balancedPlaylist(a []int, n int) (ans []int) {
	a = append(append(a, a...), a...)
	q := MonotoneQueue{} // 维护区间最大值
	for i, j := 0, 0; i < n; i++ {
		// 不断扩大区间右端点 j 直至不满足题目要求
		for ; j < 3*n && (q.Size == 0 || q.Top() <= 2*a[j]); j++ {
			q.Push(a[j])
		}
		sz := j - i
		if sz > 2*n {
			sz = -1
		}
		ans = append(ans, sz)
		if q.Size == sz {
			q.Pop()
		}
	}
	return
}

// 二维单调队列
// 输入：一个 n 行 m 列的矩阵 mat
// 输入：高 h 宽 w 的窗口大小
// 返回：一个 n-h+1 行 m-w+1 列的矩阵 areaMax，其中 areaMax[i][j] 表示窗口左上角位于矩阵 (i,j) 时的窗口中元素的最大值
// https://www.luogu.com.cn/problem/P2216
// 图解：https://cdn.acwing.com/media/article/image/2021/06/29/52559_7d7b27ced8-1.png
// https://codeforces.com/problemset/problem/1195/E 2100
// https://www.lanqiao.cn/problems/17002/learning/?contest_id=174
func fixedSizeAreaMax(mat [][]int, h, w int) [][]int {
	n, m := len(mat), len(mat[0])

	// 每行求一遍滑窗最值，窗口大小为 w
	rowMax := make([][]int, n)
	for i := range rowMax {
		rowMax[i] = make([]int, m-w+1)
	}
	for i, row := range mat {
		q := MonotoneQueue{} // 维护区间最大值
		for j, v := range row {
			q.Push(v)
			if q.Size > w {
				q.Pop()
			}
			if j+1 >= w {
				rowMax[i][j+1-w] = q.Top()
			}
		}
	}

	// 对上面的结果再求一遍滑窗最值，窗口大小为 h
	areaMax := make([][]int, n-h+1)
	for i := range areaMax {
		areaMax[i] = make([]int, m-w+1)
	}
	for j := 0; j+w <= m; j++ {
		q := MonotoneQueue{} // 维护区间最大值
		for i, mx := range rowMax {
			q.Push(mx[j])
			if q.Size > h {
				q.Pop()
			}
			if i+1 >= h {
				areaMax[i+1-h][j] = q.Top()
			}
		}
	}
	return areaMax
}
