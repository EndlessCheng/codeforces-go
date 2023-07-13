package copypasta

import (
	"container/heap"
	"sort"
)

/*
可视化 https://visualgo.net/zh/heap

思维·转换
https://www.luogu.com.cn/problem/P2859
https://www.luogu.com.cn/problem/P4952 枚举中位数
LC857 https://leetcode.cn/problems/minimum-cost-to-hire-k-workers/
https://codeforces.com/contest/713/problem/C 使序列严格递增的最小操作次数 (+1/-1)
    https://codeforces.com/blog/entry/47094?#comment-315068
    https://codeforces.com/blog/entry/77298 Slope trick
https://codeforces.com/problemset/problem/884/D 从结果倒推（类似霍夫曼编码）
http://codeforces.com/problemset/problem/912/D 贡献
https://codeforces.com/problemset/problem/1251/E2
- 按 (mi,pi) 排序，然后把 (i,mi) 画在平面直角坐标系上
- 初始时，在 y=x 直线下方的点都可以视作是「免费」的，如果有不能免费的点，应考虑从最后一个不能免费的到末尾这段中的最小 pi，然后将 y=x 抬高成 y=x+1 继续比较
- 维护最小 pi 可以用最小堆

求前缀/后缀的最小的 k 个元素和（k 固定）https://www.luogu.com.cn/problem/P4952 https://www.luogu.com.cn/problem/P3963
滑动窗口中位数 LC480 https://leetcode-cn.com/problems/sliding-window-median/

第 k 小子序列和 https://codeforces.com/gym/101234/problem/G https://leetcode.cn/problems/find-the-k-sum-of-an-array/
- 思路见我的题解 https://leetcode.cn/problems/find-the-k-sum-of-an-array/solution/zhuan-huan-dui-by-endlesscheng-8yiq/

基于堆的反悔贪心（反悔堆）
https://www.cnblogs.com/nth-element/p/11768155.html
题单 https://www.luogu.com.cn/training/8793
https://codeforces.com/problemset/problem/1526/C2
JSOI07 建筑抢修 https://www.luogu.com.cn/problem/P4053 LC630 https://leetcode-cn.com/problems/course-schedule-iii/
用堆来不断修正最优决策 https://codeforces.com/problemset/problem/1428/E
股票买卖 https://codeforces.com/problemset/problem/865/D
https://atcoder.jp/contests/abc249/tasks/abc249_f
前缀和 https://codeforces.com/problemset/problem/1779/C 推荐

区间贪心相关
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

type hp struct{ sort.IntSlice }

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

// pushPop 将 v 入堆，然后弹出并返回堆顶
// 使用见下面的 dynamicMedians
func (h *hp) pushPop(v int) int {
	if h.Len() > 0 && v > h.IntSlice[0] { // 最大堆改成 v < h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

//

// 自定义类型（int64 可以替换成其余类型）
type hp64 []int64

func (h hp64) Len() int           { return len(h) }
func (h hp64) Less(i, j int) bool { return h[i] < h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v any)        { *h = append(*h, v.(int64)) }
func (h *hp64) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp64) push(v int64)      { heap.Push(h, v) }
func (h *hp64) pop() int64        { return heap.Pop(h).(int64) } // 稍微封装一下，方便使用

//

// 支持修改、删除指定元素的堆
// 用法：调用 push 会返回一个 *viPair 指针，记作 p
// 将 p 存于他处（如 slice 或 map），可直接在外部修改 p.v 后调用 fix(p.hi)，从而做到修改堆中指定元素
// 调用 remove(p.hi) 可以从堆中删除 p
// 例题 https://atcoder.jp/contests/abc170/tasks/abc170_e
// 模拟 multiset https://codeforces.com/problemset/problem/1106/E
type viPair struct {
	v  int64
	hi int // *viPair 在 mh 中的下标，可随着 Push Pop 等操作自动改变
}
type mh []*viPair // mh 指 modifiable heap

func (h mh) Len() int              { return len(h) }
func (h mh) Less(i, j int) bool    { return h[i].v < h[j].v } // > 为最大堆
func (h mh) Swap(i, j int)         { h[i], h[j] = h[j], h[i]; h[i].hi = i; h[j].hi = j }
func (h *mh) Push(v any)           { *h = append(*h, v.(*viPair)) }
func (h *mh) Pop() any             { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh) push(v int64) *viPair { p := &viPair{v, len(*h)}; heap.Push(h, p); return p }
func (h *mh) pop() *viPair         { return heap.Pop(h).(*viPair) }
func (h *mh) fix(i int)            { heap.Fix(h, i) }
func (h *mh) remove(i int) *viPair { return heap.Remove(h, i).(*viPair) }

// 对顶堆求动态中位数：medians[i] = a[:i+1] 的中位数
// https://www.luogu.com.cn/problem/P1168
// LC295 https://leetcode-cn.com/problems/find-median-from-data-stream/
// 与树状数组结合 https://leetcode-cn.com/contest/season/2020-fall/problems/5TxKeK/
func dynamicMedians(a []int) []int {
	n := len(a)
	medians := make([]int, 0, n)
	var big, small hp
	for _, v := range a {
		if len(big.IntSlice) == len(small.IntSlice) {
			big.push(-small.pushPop(-v))
		} else {
			small.push(-big.pushPop(v))
		}
		medians = append(medians, big.IntSlice[0])
	}
	return medians
}

// 下面是对顶堆模板
// 可以用来动态维护第 k 小 / 前 k 小的元素之和
// 还支持调整 k 的值
// 这里 k 就是 left 的大小
// 第 k 小 = left.a[0]
// 前 k 小的元素之和 = left.s
// 应用见 https://codeforces.com/contest/1374/problem/E2 https://codeforces.com/contest/1374/submission/193671570
type maxMinHeap struct {
	left  *maxHp
	right *minHp
}

// 向对顶堆中插入 v
// 保证 left 大小不变
func (h *maxMinHeap) push(v pair) {
	h.right.push(h.left.pushPop(v))
	//h.left.push(h.right.pushPop(v)) // 这样写就是插入 v 的同时扩大 left
}

// 缩小 left
func (h *maxMinHeap) l2r() {
	if h.left.Len() == 0 {
		panic("h.left is empty")
	}
	h.right.push(heap.Pop(h.left).(pair))
}

// 扩大 left
func (h *maxMinHeap) r2l() {
	if h.right.Len() == 0 {
		panic("h.right is empty")
	}
	h.left.push(heap.Pop(h.right).(pair))
}

type pair struct{ t, i int }
type minHp struct {
	a []pair
	s int // 维护堆中元素之和
}

func (h minHp) Len() int           { return len(h.a) }
func (h minHp) Less(i, j int) bool { return h.a[i].t < h.a[j].t }
func (h minHp) Swap(i, j int)      { h.a[i], h.a[j] = h.a[j], h.a[i] }
func (h *minHp) Push(v any)        { h.s += v.(pair).t; h.a = append(h.a, v.(pair)) }
func (h *minHp) Pop() any          { v := h.a[len(h.a)-1]; h.s -= v.t; h.a = h.a[:len(h.a)-1]; return v }
func (h *minHp) push(v pair)       { heap.Push(h, v) }
func (h *minHp) pushPop(v pair) pair {
	if h.Len() > 0 && v.t > h.a[0].t {
		h.s += v.t - h.a[0].t
		v, h.a[0] = h.a[0], v
		heap.Fix(h, 0)
	}
	return v
}

type maxHp struct {
	a []pair
	s int
}

func (h maxHp) Len() int           { return len(h.a) }
func (h maxHp) Less(i, j int) bool { return h.a[i].t > h.a[j].t }
func (h maxHp) Swap(i, j int)      { h.a[i], h.a[j] = h.a[j], h.a[i] }
func (h *maxHp) Push(v any)        { h.s += v.(pair).t; h.a = append(h.a, v.(pair)) }
func (h *maxHp) Pop() any          { v := h.a[len(h.a)-1]; h.s -= v.t; h.a = h.a[:len(h.a)-1]; return v }
func (h *maxHp) push(v pair)       { heap.Push(h, v) }
func (h *maxHp) pushPop(v pair) pair {
	if h.Len() > 0 && v.t < h.a[0].t {
		h.s += v.t - h.a[0].t
		v, h.a[0] = h.a[0], v
		heap.Fix(h, 0)
	}
	return v
}
