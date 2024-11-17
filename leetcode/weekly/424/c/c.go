package main

import (
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	diff := make([]int, n+1)
	sumD, k := 0, 0
	for i, x := range nums {
		sumD += diff[i]
		for k < len(queries) && sumD < x { // 需要添加询问，把 x 减小
			q := queries[k]
			l, r, val := q[0], q[1], q[2]
			diff[l] += val
			diff[r+1] -= val
			if l <= i && i <= r { // x 在更新范围中
				sumD += val
			}
			k++
		}
		if sumD < x { // 无法更新
			return -1
		}
	}
	return k
}

type seg []struct {
	l, r, mx, todo int
}

func (t seg) do(o, v int) {
	t[o].mx -= v
	t[o].todo += v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg) maintain(o int) {
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mx = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, v int) {
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
	t.maintain(o)
}

func minZeroArray3(nums []int, queries [][]int) int {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	if t[1].mx <= 0 {
		return 0
	}
	for i, q := range queries {
		t.update(1, q[0], q[1], q[2])
		if t[1].mx <= 0 {
			return i + 1
		}
	}
	return -1
}

func minZeroArray2(nums []int, queries [][]int) int {
	q := len(queries)
	diff := make([]int, len(nums)+1)
	ans := sort.Search(q+1, func(k int) bool {
		clear(diff)
		for _, q := range queries[:k] {
			l, r, val := q[0], q[1], q[2]
			diff[l] += val
			diff[r+1] -= val
		}

		sumD := 0
		for i, x := range nums {
			sumD += diff[i]
			if x > sumD {
				return false
			}
		}
		return true
	})
	if ans > q {
		return -1
	}
	return ans
}
