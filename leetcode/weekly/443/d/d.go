package main

import (
	"container/heap"
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func minOperations1(nums []int, x, k int) int64 {
	n := len(nums)
	dis := medianSlidingWindow(nums, x)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= k; i++ {
		f[i][i*x-1] = math.MaxInt
		for j := i * x; j <= n-(k-i)*x; j++ { // 左右留出足够空间给其他子数组
			f[i][j] = min(f[i][j-1], f[i-1][j-x]+dis[j-x]) // j-x 为子数组左端点
		}
	}
	return int64(f[k][n])
}

func minOperations(nums []int, x, k int) int64 {
	n := len(nums)
	dis := medianSlidingWindow(nums, x)
	f := make([]int, n+1)
	g := make([]int, n+1) // 滚动数组
	for i := 1; i <= k; i++ {
		g[i*x-1] = math.MaxInt
		for j := i * x; j <= n-(k-i)*x; j++ { // 左右留出足够空间给其他子数组
			g[j] = min(g[j-1], f[j-x]+dis[j-x]) // j-x 为子数组左端点
		}
		f, g = g, f
	}
	return int64(f[n])
}

// 480. 滑动窗口中位数（有改动）
// 返回 nums 的所有长为 k 的子数组的（到子数组中位数的）距离和
func medianSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	left := newLazyHeap()  // 最大堆（元素取反）
	right := newLazyHeap() // 最小堆

	for i, in := range nums {
		// 1. 进入窗口
		if left.size == right.size {
			left.push(-right.pushPop(in))
		} else {
			right.push(-left.pushPop(-in))
		}

		l := i + 1 - k
		if l < 0 { // 窗口大小不足 k
			continue
		}

		// 2. 计算答案
		v := -left.top()
		s1 := v*left.size + left.sum // sum 取反
		s2 := right.sum - v*right.size
		ans[l] = s1 + s2

		// 3. 离开窗口
		out := nums[l]
		if out <= -left.top() {
			left.remove(-out)
			if left.size < right.size {
				left.push(-right.pop()) // 平衡两个堆的大小
			}
		} else {
			right.remove(out)
			if left.size > right.size+1 {
				right.push(-left.pop()) // 平衡两个堆的大小
			}
		}
	}

	return ans
}

func newLazyHeap() *lazyHeap {
	return &lazyHeap{removeCnt: map[int]int{}}
}

// 懒删除堆
type lazyHeap struct {
	sort.IntSlice
	removeCnt map[int]int // 每个元素剩余需要删除的次数
	size      int         // 实际大小
	sum       int         // 堆中元素总和
}

// 必须实现的两个接口
func (h *lazyHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 删除
func (h *lazyHeap) remove(v int) {
	h.removeCnt[v]++ // 懒删除
	h.size--
	h.sum -= v
}

// 正式执行删除操作
func (h *lazyHeap) applyRemove() {
	for h.removeCnt[h.IntSlice[0]] > 0 {
		h.removeCnt[h.IntSlice[0]]--
		heap.Pop(h)
	}
}

// 查看堆顶
func (h *lazyHeap) top() int {
	h.applyRemove()
	return h.IntSlice[0]
}

// 出堆
func (h *lazyHeap) pop() int {
	h.applyRemove()
	h.size--
	h.sum -= h.IntSlice[0]
	return heap.Pop(h).(int)
}

// 入堆
func (h *lazyHeap) push(v int) {
	if h.removeCnt[v] > 0 {
		h.removeCnt[v]-- // 抵消之前的删除
	} else {
		heap.Push(h, v)
	}
	h.size++
	h.sum += v
}

// push(v) 然后 pop()
func (h *lazyHeap) pushPop(v int) int {
	if h.size > 0 && v > h.top() { // 最小堆，v 比堆顶大就替换堆顶
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
