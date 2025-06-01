package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math/bits"
)

// https://space.bilibili.com/206214
const mx int = 1e5

var np = [mx + 1]bool{true, true}

func init() {
	for i := 2; i <= mx; i++ {
		if !np[i] {
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}
}

type lazySeg []struct {
	l, r int
	mx   int
	todo int
}

func mergeInfo(a, b int) int {
	return max(a, b)
}

const todoInit = 0

func mergeTodo(f, old int) int {
	return f + old
}

func (t lazySeg) apply(o int, f int) {
	cur := &t[o]
	cur.mx += f
	cur.todo = mergeTodo(f, cur.todo)
}

func (t lazySeg) maintain(o int) {
	t[o].mx = mergeInfo(t[o<<1].mx, t[o<<1|1].mx)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t lazySeg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].mx = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazySeg) update(o, l, r int, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t lazySeg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func newLazySegmentTreeWithArray(a []int) lazySeg {
	n := len(a)
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func newLazySegmentTree(n int, initVal int) lazySeg {
	a := make([]int, n)
	for i := range a {
		a[i] = initVal
	}
	return newLazySegmentTreeWithArray(a)
}

func maximumCount(nums []int, queries [][]int) (ans []int) {
	n := len(nums)
	pos := map[int]*redblacktree.Tree[int, struct{}]{}
	for i, x := range nums {
		if np[x] {
			continue
		}
		if _, ok := pos[x]; !ok {
			pos[x] = redblacktree.New[int, struct{}]()
		}
		pos[x].Put(i, struct{}{})
	}

	t := newLazySegmentTree(n, 0)
	for _, ps := range pos {
		if ps.Size() > 1 {
			t.update(1, ps.Left().Key, ps.Right().Key, 1)
		}
	}

	for _, q := range queries {
		i, v := q[0], q[1]
		old := nums[i]
		nums[i] = v

		if !np[old] {
			ps := pos[old]
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, -1)
			}
			ps.Remove(i)
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, 1)
			} else if ps.Empty() {
				delete(pos, old)
			}
		}

		if !np[v] {
			if _, ok := pos[v]; !ok {
				pos[v] = redblacktree.New[int, struct{}]()
			}
			ps := pos[v]
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, -1)
			}
			ps.Put(i, struct{}{})
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, 1)
			}
		}

		ans = append(ans, len(pos)+t.query(1, 0, n-1))
	}

	return
}
