package main

import "math/bits"

// https://space.bilibili.com/206214
type data struct {
	f00 int // 第一个数一定不选，最后一个数一定不选
	f01 int // 第一个数一定不选，最后一个数可选可不选
	f10 int // 第一个数可选可不选，最后一个数一定不选
	f11 int // 第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制
}

type seg []data

func (t seg) maintain(o int) {
	a, b := t[o<<1], t[o<<1|1]
	t[o] = data{
		max(a.f00+b.f10, a.f01+b.f00),
		max(a.f00+b.f11, a.f01+b.f01),
		max(a.f10+b.f10, a.f11+b.f00),
		max(a.f10+b.f11, a.f11+b.f01),
	}
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o].f11 = max(a[l], 0)
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o].f11 = max(val, 0)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t.maintain(o)
}

func maximumSumSubsequence(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)

	for _, q := range queries {
		t.update(1, 0, n-1, q[0], q[1])
		ans += t[1].f11 // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
	}
	return ans % 1_000_000_007
}
