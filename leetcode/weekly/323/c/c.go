package main

import "math/bits"

// https://space.bilibili.com/206214
type segTree []struct {
	l, r, pre0, suf0, max0, todo int
}

func newSegTree(n int) segTree {
	t := make(segTree, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	return t
}

func (t segTree) do(i, v int) {
	o := &t[i]
	size := 0
	if v <= 0 {
		size = o.r - o.l + 1
	}
	o.pre0 = size // 区间前缀连续 0 的个数
	o.suf0 = size // 区间后缀连续 0 的个数
	o.max0 = size // 区间最长连续 0 的个数
	o.todo = v
}

func (t segTree) spread(o int) {
	v := t[o].todo
	if v != -1 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = -1
	}
}

// 初始化线段树
func (t segTree) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t.do(o, -1)
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// 把 [l,r] 都置为 v
func (t segTree) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}

	lo, ro := t[o<<1], t[o<<1|1]
	t[o].pre0 = lo.pre0
	if lo.pre0 == m-t[o].l+1 {
		t[o].pre0 += ro.pre0 // 和右子树的 pre0 拼起来
	}
	t[o].suf0 = ro.suf0
	if ro.suf0 == t[o].r-m {
		t[o].suf0 += lo.suf0 // 和左子树的 suf0 拼起来
	}
	t[o].max0 = max(lo.max0, ro.max0, lo.suf0+ro.pre0)
}

// 线段树二分，找最左边的区间左端点，满足区间全为 0 且长度 >= size
// 如果不存在这样的区间，返回 -1
func (t segTree) findFirst(o, size int) int {
	if t[o].max0 < size {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findFirst(o<<1, size) // 递归左子树
	if idx < 0 {
		// 左子树的后缀 0 个数 + 右子树的前缀 0 个数 >= size
		if t[o<<1].suf0+t[o<<1|1].pre0 >= size {
			m := (t[o].l + t[o].r) >> 1
			return m - t[o<<1].suf0 + 1
		}
		idx = t.findFirst(o<<1|1, size) // 递归右子树
	}
	return idx
}

// 上面为线段树代码

type interval struct {
	l, r int
}

type Allocator struct {
	tree   segTree
	blocks map[int][]interval
}

func Constructor(n int) Allocator {
	return Allocator{
		tree:   newSegTree(n),
		blocks: map[int][]interval{},
	}
}

func (a Allocator) Allocate(size, mID int) int {
	i := a.tree.findFirst(1, size)
	if i < 0 { // 无法分配内存
		return -1
	}
	a.blocks[mID] = append(a.blocks[mID], interval{i, i + size - 1})
	a.tree.update(1, i, i+size-1, 1) // 分配内存 [i,i+size-1]
	return i
}

func (a Allocator) FreeMemory(mID int) (ans int) {
	for _, p := range a.blocks[mID] {
		ans += p.r - p.l + 1
		a.tree.update(1, p.l, p.r, 0) // 释放内存
	}
	delete(a.blocks, mID)
	return
}
