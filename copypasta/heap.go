package copypasta

import (
	"container/heap"
	"sort"
)

/*
本页面力扣题目已整理至【题单】常用数据结构
https://leetcode.cn/circle/discuss/mOr1u6/

可视化 https://visualgo.net/zh/heap
【证明】堆化的时间复杂度为 O(n) https://leetcode.cn/problems/take-gifts-from-the-richest-pile/solution/yuan-di-dui-hua-o1-kong-jian-fu-ti-dan-p-fzdh/

【疑问】对于所有 1~n 的排列，heapify 后是有序数组的排列，有多少个？

动态维护最大的 k 个数用最小堆，动态维护最小的 k 个数用最大堆
https://codeforces.com/problemset/problem/1969/D 1900

#### 第 K 小/大（值/和）
- [703. 数据流中的第 K 大元素](https://leetcode.cn/problems/kth-largest-element-in-a-stream/)
- [2558. 从数量最多的堆取走礼物](https://leetcode.cn/problems/take-gifts-from-the-richest-pile/) 1277
- [2530. 执行 K 次操作后的最大分数](https://leetcode.cn/problems/maximal-score-after-applying-k-operations/) 1386
- [1962. 移除石子使总数最小](https://leetcode.cn/problems/remove-stones-to-minimize-the-total/) 1419
- [2208. 将数组和减半的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-halve-array-sum/) 1550
- [2233. K 次增加后的最大乘积](https://leetcode.cn/problems/maximum-product-after-k-increments/) 1686
- [2542. 最大子序列的分数](https://leetcode.cn/problems/maximum-subsequence-score/) 2056
- [1383. 最大的团队表现值](https://leetcode.cn/problems/maximum-performance-of-a-team/) 2091
- [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/)
    题解 https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/solution/jiang-qing-chu-wei-shi-yao-yi-kai-shi-ya-i0dj/
- [1439. 有序矩阵中的第 k 个最小数组和](https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/) 2134
    题解 https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/solution/san-chong-suan-fa-bao-li-er-fen-da-an-du-k1vd/
https://atcoder.jp/contests/abc297/tasks/abc297_e
https://codeforces.com/problemset/problem/1106/D 1500
https://codeforces.com/problemset/problem/1140/C 1600
- https://atcoder.jp/contests/abc376/tasks/abc376_e
https://codeforces.com/problemset/problem/1862/E 1600
https://codeforces.com/problemset/problem/1862/E 1600
https://codeforces.com/problemset/problem/1935/C 1800
https://codeforces.com/problemset/problem/1196/F 2200

#### 懒删除
本页面搜索【懒删除堆】

#### 模拟
- [2532. 过桥的时间](https://leetcode.cn/problems/time-to-cross-a-bridge/) 2589
https://codeforces.com/problemset/problem/1985/F
https://codeforces.com/problemset/problem/1945/G 2500

#### 思维·转换
- [2054. 两个最好的不重叠活动](https://leetcode.cn/problems/two-best-non-overlapping-events/) 1883
- [264. 丑数 II](https://leetcode.cn/problems/ugly-number-ii/)
- [313. 超级丑数](https://leetcode.cn/problems/super-ugly-number/)
https://www.luogu.com.cn/problem/P5930
- 3D 接雨水 LC407 https://leetcode.cn/problems/trapping-rain-water-ii/
https://www.luogu.com.cn/problem/P2859
https://www.luogu.com.cn/problem/P4952 枚举中位数
LC857 https://leetcode.cn/problems/minimum-cost-to-hire-k-workers/
https://codeforces.com/contest/713/problem/C 使序列严格递增的最小操作次数 (+1/-1)
    https://codeforces.com/blog/entry/47094?#comment-315068
    https://codeforces.com/blog/entry/77298 Slope trick
https://codeforces.com/problemset/problem/884/D 从结果倒推（类似霍夫曼编码）
https://codeforces.com/problemset/problem/912/D 贡献
https://codeforces.com/problemset/problem/1251/E2
- 按 (mi,pi) 排序，然后把 (i,mi) 画在平面直角坐标系上
- 初始时，在 y=x 直线下方的点都可以视作是「免费」的，如果有不能免费的点，应考虑从最后一个不能免费的到末尾这段中的最小 pi，然后将 y=x 抬高成 y=x+1 继续比较
- 维护最小 pi 可以用最小堆
https://atcoder.jp/contests/agc057/tasks/agc057_b
https://ac.nowcoder.com/acm/contest/65157/C
https://www.luogu.com.cn/problem/P7840
- https://atcoder.jp/contests/abc359/tasks/abc359_f
https://atcoder.jp/contests/arc051/tasks/arc051_c 1898=CF2147

第 k 小子序列和 https://codeforces.com/gym/101234/problem/G https://leetcode.cn/problems/find-the-k-sum-of-an-array/
- 思路见我的题解 https://leetcode.cn/problems/find-the-k-sum-of-an-array/solution/zhuan-huan-dui-by-endlesscheng-8yiq/

#### 会议室
https://codeforces.com/problemset/problem/845/C
https://leetcode.cn/problems/meeting-rooms-ii/
https://leetcode.cn/problems/meeting-rooms-iii/
https://leetcode.cn/problems/t3fKg1/
https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/

#### 基于堆的反悔贪心（反悔堆）
- [LCP 30. 魔塔游戏](https://leetcode.cn/problems/p0NxJO/)
- [1642. 可以到达的最远建筑](https://leetcode.cn/problems/furthest-building-you-can-reach/) 1962
- [630. 课程表 III](https://leetcode.cn/problems/course-schedule-iii/)
   - JSOI07 建筑抢修 https://www.luogu.com.cn/problem/P4053
- [871. 最低加油次数](https://leetcode.cn/problems/minimum-number-of-refueling-stops/) 2074
- [2813. 子序列最大优雅度](https://leetcode.cn/problems/maximum-elegance-of-a-k-length-subsequence/) 2582 也可以用栈
- [3049. 标记所有下标的最早秒数 II](https://leetcode.cn/problems/earliest-second-to-mark-indices-ii/) 3111
- [2463. 最小移动总距离](https://leetcode.cn/problems/minimum-total-distance-traveled/) 做到 O((n+m)log(n+m))  模拟费用流
   - https://codeforces.com/problemset/problem/797/F 2600
   - https://www.cnblogs.com/Scarab/p/17672813.html
- [2599. 使前缀和数组非负](https://leetcode.cn/problems/make-the-prefix-sum-non-negative/)（会员题）
题单 https://www.luogu.com.cn/training/8793
https://www.luogu.com.cn/problem/P2949 经典题
https://www.luogu.com.cn/problem/P3045
https://www.luogu.com.cn/problem/P11457
https://codeforces.com/problemset/problem/1526/C2 1600
https://codeforces.com/problemset/problem/1779/C 1600 前缀和 推荐
https://codeforces.com/problemset/problem/730/I 2000
- 加强版 https://atcoder.jp/contests/agc018/tasks/agc018_c
https://codeforces.com/problemset/problem/1974/G 2000
https://atcoder.jp/contests/abc249/tasks/abc249_f 1786=CF2062
https://codeforces.com/problemset/problem/1428/E 2200 用堆来不断修正最优决策
https://atcoder.jp/contests/aising2020/tasks/aising2020_e 2133=CF2324 洛谷 P2949 的加强版
https://codeforces.com/problemset/problem/865/D 2400 股票买卖
https://codeforces.com/problemset/problem/3/D 2600 难度虚高
https://www.cnblogs.com/nth-element/p/11768155.html

#### 模拟费用流
https://www.luogu.com/article/wn7c3auk 模拟费用流小记 by command_block
https://www.cnblogs.com/Call-me-Eric/p/17878027.html
https://blog.csdn.net/wyy603/article/details/105038432
https://www.luogu.com.cn/problem/P4694
https://uoj.ac/problem/455
https://www.luogu.com.cn/problem/P1484
https://codeforces.com/problemset/problem/730/I 2000
https://www.luogu.com.cn/problem/P6122

#### 区间贪心相关
最小不相交区间划分数
- https://www.acwing.com/problem/content/113/
- https://www.acwing.com/problem/content/908/
- https://codeforces.com/problemset/problem/845/C
https://codeforces.com/problemset/problem/555/B
区间最大交集 https://codeforces.com/problemset/problem/754/D
https://codeforces.com/problemset/problem/1701/D
区间放球 https://atcoder.jp/contests/abc214/tasks/abc214_e
*/

// 下面这些都是最小堆

type hp struct{ sort.IntSlice } // 继承 sort.IntSlice 的 Len Less Swap，这样就只需要实现 Push 和 Pop

//func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 加上这行变成最大堆
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) } // 稍微封装一下，方便使用

// EXTRA: 参考 Python，引入下面两个效率更高的方法（相比调用 push + pop）
// replace 弹出并返回堆顶，同时将 v 入堆
// 需保证 h 非空
func (h *hp) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}

// pushPop 先将 v 入堆，然后弹出并返回堆顶
// 使用见下面的 dynamicMedians
func (h *hp) pushPop(v int) int {
	if h.Len() > 0 && v > h.IntSlice[0] { // 最大堆改成 v < h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

// 对顶堆：前缀中位数
// 返回数组 medians，其中 medians[i] 等于前缀 a[:i+1] 的中位数
// 如果前缀长度是偶数，取大的那个作为中位数
// LC295 https://leetcode.cn/problems/find-median-from-data-stream/
// - https://www.luogu.com.cn/problem/P1168
func dynamicMedians(a []int) []int {
	medians := make([]int, len(a))
	l := hp{} // 大根堆，元素取反
	r := hp{} // 小根堆
	for i, v := range a {
		if l.Len() == r.Len() {
			r.push(-l.pushPop(-v))
		} else {
			l.push(-r.pushPop(v))
		}
		// 如果 i 是奇数，另一个中位数是 -l.IntSlice[0]
		medians[i] = r.IntSlice[0]
	}
	return medians
}

////

// 自定义类型（int32 可以替换成其余类型）
type hp32 []int32

func (h hp32) Len() int           { return len(h) }
func (h hp32) Less(i, j int) bool { return h[i] < h[j] } // > 为最大堆
func (h hp32) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp32) Push(v any)        { *h = append(*h, v.(int32)) }
func (h *hp32) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp32) push(v int32)      { heap.Push(h, v) }
func (h *hp32) pop() int32        { return heap.Pop(h).(int32) } // 稍微封装一下，方便使用

////

// 支持修改、删除指定元素的堆
// 用法：调用 push 会返回一个 *viPair 指针，记作 p
// 将 p 存于他处（如 slice 或 map），可直接在外部修改 p.v 后调用 fix(p.hi)，从而做到修改堆中指定元素
// 调用 remove(p.hi) 可以从堆中删除 p
// 例题 https://atcoder.jp/contests/abc170/tasks/abc170_e
// 模拟 multiset https://codeforces.com/problemset/problem/1106/E
type viPair struct {
	v  int
	hi int // *viPair 在 mh 中的下标，可随着 Push Pop 等操作自动改变
}
type mh []*viPair // mh 指 modifiable heap

func (h mh) Len() int              { return len(h) }
func (h mh) Less(i, j int) bool    { return h[i].v < h[j].v } // > 为最大堆
func (h mh) Swap(i, j int)         { h[i], h[j] = h[j], h[i]; h[i].hi = i; h[j].hi = j }
func (h *mh) Push(v any)           { *h = append(*h, v.(*viPair)) }
func (h *mh) Pop() any             { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh) push(v int) *viPair   { p := &viPair{v, len(*h)}; heap.Push(h, p); return p }
func (h *mh) pop() *viPair         { return heap.Pop(h).(*viPair) }
func (h *mh) fix(i int)            { heap.Fix(h, i) }
func (h *mh) remove(i int) *viPair { return heap.Remove(h, i).(*viPair) }

////

// 懒删除堆
// LC716 https://leetcode.cn/problems/max-stack/
// LC3092 https://leetcode.cn/problems/most-frequent-ids/
// https://codeforces.com/problemset/problem/1883/D 1500
// https://codeforces.com/problemset/problem/796/C 1900
// https://codeforces.com/problemset/problem/2009/G2 2200
// https://codeforces.com/problemset/problem/1732/D2 2400 简化版懒删除堆
type lazyHeap struct {
	sort.IntSlice
	todo map[int]int
	size int // 实际大小
	sum  int // 实际元素和（可选）
}

func (h lazyHeap) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *lazyHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *lazyHeap) del(v int)         { h.todo[v]++; h.size--; h.sum -= v } // 懒删除
func (h *lazyHeap) push(v int) {
	if h.todo[v] > 0 {
		h.todo[v]--
	} else {
		heap.Push(h, v)
	}
	h.size++
	h.sum += v
}
func (h *lazyHeap) _do() {
	for h.Len() > 0 && h.todo[h.IntSlice[0]] > 0 {
		h.todo[h.IntSlice[0]]--
		heap.Pop(h)
	}
}
func (h *lazyHeap) pop() int    { h._do(); h.size--; v := heap.Pop(h).(int); h.sum -= v; return v }
func (h *lazyHeap) top() int    { h._do(); return h.IntSlice[0] }
func (h *lazyHeap) empty() bool { return h.size == 0 }
func (h *lazyHeap) pushPop(v int) int {
	if h.size > 0 && v < h.top() { // 最大堆，v 比堆顶小就替换堆顶
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

// 对顶堆：滑动窗口前 k 小元素和
// 保证 1 <= k <= windowSize <= n
// 返回数组 kthSum，其中 kthSum[i] 为 a[i:i+windowSize] 的前 k 小元素和
// - [3013. 将数组分成最小总代价的子数组 II](https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/) 2540
// - https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii/
// 另见 treap_kthsum.go
func slidingWindowKthSum(a []int, windowSize, k int) []int {
	h := newKthHeap()
	// 注：也可以 copy 一份 a[:k] 然后堆化
	for _, v := range a[:k] {
		h.l.push(v)
	}
	for _, v := range a[k:windowSize] {
		h.add(v)
	}
	kthSum := make([]int, len(a)-windowSize+1)
	kthSum[0] = h.l.sum
	for r := windowSize; r < len(a); r++ {
		l := r - windowSize // 前一个窗口的左端点
		h.add(a[r])
		h.del(a[l]) // 先加再删（注意 windowSize=1 的情况）
		kthSum[l+1] = h.l.sum
	}
	return kthSum
}

type kthHeap struct {
	l *lazyHeap // 最大堆
	r *lazyHeap // 最小堆，所有元素取反
}

func newKthHeap() *kthHeap {
	return &kthHeap{&lazyHeap{todo: map[int]int{}}, &lazyHeap{todo: map[int]int{}}}
}

func (h *kthHeap) empty() bool {
	return h.l.size == 0 && h.r.size == 0
}

func (h *kthHeap) size() int {
	return h.l.size + h.r.size
}

func (h *kthHeap) l2r() {
	if h.l.size == 0 {
		panic("h.l is empty")
	}
	h.r.push(-h.l.pop())
}

func (h *kthHeap) r2l() {
	if h.r.size == 0 {
		panic("h.r is empty")
	}
	h.l.push(-h.r.pop())
}

// 保证 h.l 大小不变
func (h *kthHeap) add(v int) {
	h.r.push(-h.l.pushPop(v))
}

// 保证 h.l 大小不变
func (h *kthHeap) del(v int) {
	if v <= h.l.top() {
		h.l.del(v)
		h.r2l()
	} else {
		h.r.del(-v)
	}
}

// 把 h.l 的大小调整为 k
func (h *kthHeap) balance(k int) {
	for h.l.size > k {
		h.l2r()
	}
	for h.l.size < k {
		h.r2l()
	}
}

// 其它题目
// 求前缀/后缀的最小的 k 个元素和（k 固定）https://www.luogu.com.cn/problem/P4952 https://www.luogu.com.cn/problem/P3963
// - https://www.codechef.com/problems/OKLAMA
// LC480 滑动窗口中位数 https://leetcode.cn/problems/sliding-window-median/
// https://codeforces.com/contest/1374/problem/E2 代码 https://codeforces.com/contest/1374/submission/193671570

// 如果值域比较小，可以用分桶法做到 O(n+U)
// 特别地，如果 U<=n，则时间复杂度为 O(n)
// https://leetcode.cn/problems/smallest-substring-with-identical-characters-ii/solutions/3027031/er-fen-da-an-tan-xin-gou-zao-pythonjavac-3i4f/
