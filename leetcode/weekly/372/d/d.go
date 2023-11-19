package main

import "container/heap"

// https://space.bilibili.com/206214
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	left := make([][]pair, len(heights))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i // 保证 i <= j
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j // i 直接跳到 j
		} else {
			left[j] = append(left[j], pair{heights[i], qi}) // 离线
		}
	}

	h := hp{}
	for i, x := range heights { // 从小到大枚举下标 i
		for h.Len() > 0 && h[0].h < x {
			ans[heap.Pop(&h).(pair).qi] = i // 可以跳到 i（此时 i 是最小的）
		}
		for _, p := range left[i] {
			heap.Push(&h, p) // 后面再回答
		}
	}
	return ans
}

type pair struct{ h, qi int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 下面是方法二

type seg []int

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = max(t[o<<1], t[o<<1|1])
}

func (t seg) query(o, l, r, L, v int) int {
	if v >= t[o] {
		return 0
	}
	if l == r {
		return l
	}
	m := (l + r) >> 1
	if L <= m {
		pos := t.query(o<<1, l, m, L, v)
		if pos > 0 {
			return pos
		}
	}
	return t.query(o<<1|1, m+1, r, L, v)
}

func leftmostBuildingQueriesSeg(heights []int, queries [][]int) []int {
	n := len(heights)
	t := make(seg, n*4)
	t.build(heights, 1, 1, n)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j
		} else {
			pos := t.query(1, 1, n, j+1, heights[i])
			ans[qi] = pos - 1 // 不存在时刚好得到 -1
		}
	}
	return ans
}
