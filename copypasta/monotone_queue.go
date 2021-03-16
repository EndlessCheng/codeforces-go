package copypasta

/* 单调队列 Monotone Queue
需要不断维护队列的单调性，时刻保证队列元素从大到小或从小到大
https://oi-wiki.org/ds/monotonous-queue/
https://oi-wiki.org/dp/opt/monotonous-queue-stack/
https://cp-algorithms.com/data_structures/stack_queue_modification.html
https://blog.csdn.net/weixin_43914593/article/details/105791217 算法竞赛专题解析（13）：DP优化(3)--单调队列优化
todo https://xyzl.blog.luogu.org/DQ-OP-DP

todo https://www.luogu.com.cn/problem/P2627
todo http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=1070
*/
type mqData struct {
	val int
	del int // 懒删除标记
}

type monotoneQueue struct {
	data []mqData // 初始化时可以 make([]mqData, 0, n) 来减少扩容的开销
	size int      // 单调队列对应的区间的长度
}

func (mq monotoneQueue) less(a, b mqData) bool {
	return a.val >= b.val // >= 维护区间最大值；<= 维护区间最小值
}

func (mq *monotoneQueue) push(v int) {
	mq.size++
	d := mqData{v, 1}
	for len(mq.data) > 0 && mq.less(d, mq.data[len(mq.data)-1]) {
		d.del += mq.data[len(mq.data)-1].del
		mq.data = mq.data[:len(mq.data)-1]
	}
	mq.data = append(mq.data, d)
}

func (mq *monotoneQueue) pop() {
	mq.size--
	if mq.data[0].del > 1 {
		mq.data[0].del--
	} else {
		mq.data = mq.data[1:]
	}
}

// 调用前需要判断 size > 0
func (mq monotoneQueue) top() int {
	return mq.data[0].val
}

func monotoneQueueCollections() {
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

	// 模板题 - 固定区间大小的区间最值
	// https://www.luogu.com.cn/problem/P1886 http://poj.org/problem?id=2823
	// https://codeforces.com/problemset/problem/940/E
	fixedSizeMax := func(a []int, fixedSize int) []int {
		n := len(a)
		q := monotoneQueue{} // 最大/最小由 less 来控制
		ans := make([]int, 0, n-fixedSize+1)
		for i, v := range a {
			q.push(v)
			if q.size > fixedSize {
				q.pop()
			}
			// 插入新元素并保证单调队列大小后，获取区间最值
			if i+1 >= fixedSize {
				ans = append(ans, q.top())
			}
		}
		return ans
	}

	// 模板题 - 最大子序和
	// https://www.acwing.com/problem/content/137/ https://ac.nowcoder.com/acm/contest/1006/D
	maxSubSumWithLimitSize := func(a []int, sizeLimit int) int {
		n := len(a)
		sum := make([]int, n+1) // int64
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		ans := int(-1e9)     // -1e18
		q := monotoneQueue{} // 维护区间最小值
		q.push(sum[0])
		for r := 1; r <= n; r++ {
			if q.size > sizeLimit {
				q.pop()
			}
			ans = max(ans, sum[r]-q.top())
			q.push(sum[r])
		}
		return ans
	}

	// 子数组和至少为 k 的最短非空子数组长度
	// 这题的关键在于，当右端点向右（枚举）时，左端点是绝对不会向左的（因为向左肯定会比当前求出的最短长度要长）
	// 想明白这一点就可以愉快地使用单调队列了
	// LC862 https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k/
	shortestSubSumAtLeastK := func(a []int, k int) int {
		n := len(a)
		ans := n + 1
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		q := monotoneQueue{} // 维护区间最小值
		q.push(sum[0])
		for r := 1; r <= n; r++ {
			for q.size > 0 && sum[r]-q.top() >= k {
				ans = min(ans, q.size)
				q.pop()
			}
			q.push(sum[r])
		}
		if ans > n {
			return -1
		}
		return ans
	}

	// 枚举区间左端点更为方便的情况 · 其一
	// 统计区间个数：区间最大值 >= 2*区间最小值
	// https://ac.nowcoder.com/acm/contest/6778/C
	//
	// 思路：转变成求「区间最大值 < 2*区间最小值」的区间个数
	// 随着左端点向右，右端点必然不会向左
	countSubArrayByMinMax := func(a []int) int {
		n := len(a)
		ans := n * (n + 1) / 2 // int64
		mx := monotoneQueue{}  // 维护区间最大值
		mi := monotoneQueue{}  // 维护区间最小值（需要新定义一个有不同 less 的 monotoneQueue）
		for i, j := 0, 0; i < n; i++ {
			// 确保符合条件再插入
			for ; j < n && (mx.size == 0 || mi.size == 0 || max(mx.top(), a[j]) < 2*min(mi.top(), a[j])); j++ {
				mx.push(a[j])
				mi.push(a[j])
			}
			sz := j - i
			ans -= sz
			// 若单调队列指向的区间的左端点为 i，则对应元素在下一次循环时将不再使用。故弹出之
			if mx.size == sz {
				mx.pop()
			}
			if mi.size == sz {
				mi.pop()
			}
		}
		return ans
	}

	// 枚举区间左端点更为方便的情况 · 其二
	// https://codeforces.com/problemset/problem/1237/D
	// 注意这题和 countSubArrayByMinMax 的不同之处：不满足要求的最小值一定要在最大值的右侧
	// 也可以枚举右端点，见 https://www.luogu.com.cn/blog/qianshang/solution-cf1237d
	balancedPlaylist := func(a []int, n int) (ans []int) {
		a = append(append(a, a...), a...)
		q := monotoneQueue{} // 维护区间最大值
		for i, j := 0, 0; i < n; i++ {
			// 不断扩大区间右端点 j 直至不满足题目要求
			for ; j < 3*n && (q.size == 0 || q.top() <= 2*a[j]); j++ {
				q.push(a[j])
			}
			sz := j - i
			if sz > 2*n {
				sz = -1
			}
			ans = append(ans, sz)
			if q.size == sz {
				q.pop()
			}
		}
		return
	}

	_ = []interface{}{
		fixedSizeMax, maxSubSumWithLimitSize, shortestSubSumAtLeastK,
		countSubArrayByMinMax, balancedPlaylist,
	}
}
