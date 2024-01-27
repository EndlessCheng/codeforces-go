package main

import (
	"container/heap"
	"github.com/emirpasic/gods/trees/redblacktree"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minimumCost(nums []int, k int, dist int) int64 {
	res := slidingWindowKthSum(nums[1:], dist+1, k-1)
	return int64(slices.Min(res) + nums[0])
}

func slidingWindowKthSum(a []int, windowSize, k int) []int {
	n := len(a)
	kthSum := make([]int, n-windowSize+1)
	h := &kthHeap{&lazyHeap{todo: map[int]int{}}, &lazyHeap{todo: map[int]int{}}}
	for _, v := range a[:k] {
		h.l.push(v)
	}
	for _, v := range a[k:windowSize] {
		h.add(v)
	}
	kthSum[0] = h.l.sum
	for r := windowSize; r < n; r++ {
		l := r - windowSize // 前一个窗口的左端点
		h.add(a[r])
		h.del(a[l]) // 一定要先加再删
		kthSum[l+1] = h.l.sum
	}
	return kthSum
}

type kthHeap struct {
	l *lazyHeap // 最大堆
	r *lazyHeap // 最小堆，所有元素取反
}

func (h *kthHeap) l2r() {
	h.r.push(-h.l.pop())
}
func (h *kthHeap) r2l() {
	h.l.push(-h.r.pop())
}
func (h *kthHeap) add(v int) {
	h.r.push(-h.l.pushPop(v)) // 保证 h.l 大小不变
}
func (h *kthHeap) del(v int) {
	if v <= h.l.top() {
		h.l.del(v)
		h.r2l() // 保证 h.l 大小不变
	} else {
		h.r.del(-v)
	}
}

type lazyHeap struct {
	sort.IntSlice
	todo map[int]int
	size int // 实际大小
	sum  int // 实际元素和
}

func (h lazyHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
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
func (h *lazyHeap) pop() int    { h._do(); h.size--; v := heap.Pop(h).(int); h.sum -= v; return v }
func (h *lazyHeap) top() int    { h._do(); return h.IntSlice[0] }
func (h *lazyHeap) pushPop(v int) int {
	if h.size > 0 && v < h.top() { // 最大堆，v 比堆顶小就替换堆顶
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
func (h *lazyHeap) _do() {
	for h.Len() > 0 && h.todo[h.IntSlice[0]] > 0 {
		h.todo[h.IntSlice[0]]--
		heap.Pop(h)
	}
}

//

func minimumCost2(nums []int, k, dist int) int64 {
	k--
	L := redblacktree.NewWithIntComparator()
	R := redblacktree.NewWithIntComparator()
	add := func(t *redblacktree.Tree, x int) {
		c, ok := t.Get(x)
		if ok {
			t.Put(x, c.(int)+1)
		} else {
			t.Put(x, 1)
		}
	}
	del := func(t *redblacktree.Tree, x int) {
		c, _ := t.Get(x)
		if c.(int) > 1 {
			t.Put(x, c.(int)-1)
		} else {
			t.Remove(x)
		}
	}

	sumL := nums[0]
	for _, x := range nums[1 : dist+2] {
		sumL += x
		add(L, x)
	}
	sizeL := dist + 1

	l2r := func() {
		x := L.Right().Key.(int)
		sumL -= x
		sizeL--
		del(L, x)
		add(R, x)
	}
	r2l := func() {
		x := R.Left().Key.(int)
		sumL += x
		sizeL++
		del(R, x)
		add(L, x)
	}
	for sizeL > k {
		l2r()
	}

	ans := sumL
	for i := dist + 2; i < len(nums); i++ {
		// 移除 out
		out := nums[i-dist-1]
		if _, ok := L.Get(out); ok {
			sumL -= out
			sizeL--
			del(L, out)
		} else {
			del(R, out)
		}

		// 添加 in
		in := nums[i]
		if in < L.Right().Key.(int) {
			sumL += in
			sizeL++
			add(L, in)
		} else {
			add(R, in)
		}

		// 维护大小
		if sizeL == k-1 {
			r2l()
		} else if sizeL == k+1 {
			l2r()
		}

		ans = min(ans, sumL)
	}
	return int64(ans)
}
