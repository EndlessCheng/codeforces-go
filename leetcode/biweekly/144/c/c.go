package main

import (
	"container/heap"
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxRemoval(nums []int, queries [][]int) int {
	slices.SortFunc(queries, func(a, b []int) int { return a[0] - b[0] })
	h := hp{}
	diff := make([]int, len(nums)+1)
	sumD, j := 0, 0
	for i, x := range nums {
		sumD += diff[i]
		for ; j < len(queries) && queries[j][0] <= i; j++ {
			heap.Push(&h, queries[j][1])
		}
		for sumD < x && h.Len() > 0 && h.IntSlice[0] >= i {
			sumD++
			diff[heap.Pop(&h).(int)+1]--
		}
		if sumD < x {
			return -1
		}
	}
	return h.Len()
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

//

type SegmentTree struct {
	tree, lazy []int
	n          int
}

func NewSegmentTree(size int) *SegmentTree {
	tree := make([]int, size*4)
	lazy := make([]int, size*4)
	return &SegmentTree{
		tree: tree,
		lazy: lazy,
		n:    size,
	}
}

func (st *SegmentTree) pushDown(node, start, end int) {
	if st.lazy[node] != 0 {
		st.tree[node] += st.lazy[node]
		if start != end {
			st.lazy[node*2] += st.lazy[node]
			st.lazy[node*2+1] += st.lazy[node]
		}
		st.lazy[node] = 0
	}
}

func (st *SegmentTree) updateRange(node, start, end, l, r, val int) {
	st.pushDown(node, start, end)
	if start > r || end < l {
		return
	}
	if start >= l && end <= r {
		st.lazy[node] += val
		st.pushDown(node, start, end)
		return
	}
	mid := (start + end) / 2
	st.updateRange(node*2, start, mid, l, r, val)
	st.updateRange(node*2+1, mid+1, end, l, r, val)
	st.tree[node] = int(math.Min(float64(st.tree[node*2]), float64(st.tree[node*2+1])))
}

func (st *SegmentTree) queryRange(node, start, end, l, r int) int {
	st.pushDown(node, start, end)
	if start > r || end < l {
		return math.MaxInt32
	}
	if start >= l && end <= r {
		return st.tree[node]
	}
	mid := (start + end) / 2
	left := st.queryRange(node*2, start, mid, l, r)
	right := st.queryRange(node*2+1, mid+1, end, l, r)
	return int(math.Min(float64(left), float64(right)))
}

func (st *SegmentTree) Update(l, r, val int) {
	st.updateRange(1, 0, st.n-1, l, r, val)
}

func (st *SegmentTree) Query(l, r int) int {
	return st.queryRange(1, 0, st.n-1, l, r)
}

func maxRemovalWA(nums []int, queries [][]int) int {
	n := len(nums)
	diff := make([]int, n+1)

	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l]++
		if r+1 < n {
			diff[r+1]--
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	for i := 0; i < n; i++ {
		if nums[i] > diff[i] {
			return -1
		}
		diff[i] -= nums[i]
	}

	segTree := NewSegmentTree(n)
	for i := 0; i < n; i++ {
		segTree.Update(i, i, diff[i])
	}

	ans := 0

	sort.Slice(queries, func(i, j int) bool {
		a, b := queries[i], queries[j]
		if a[1]-a[0] == b[1]-b[0] {
			return a[0] < b[0]
		}
		return (a[1] - a[0]) < (b[1] - b[0])
	})

	for _, q := range queries {
		l, r := q[0], q[1]
		if segTree.Query(l, r) > 0 {
			segTree.Update(l, r, -1)
			ans++
		}
	}

	return ans
}
