package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
var k int

type data struct {
	mul int
	cnt []int
}

type seg []data

func mergeData(a, b data) data {
	cnt := slices.Clone(a.cnt)
	for rx, c := range b.cnt {
		cnt[a.mul*rx%k] += c
	}
	return data{a.mul * b.mul % k, cnt}
}

func newData(val int) data {
	mul := val % k
	cnt := make([]int, k)
	cnt[mul] = 1
	return data{mul, cnt}
}

func (t seg) maintain(o int) {
	t[o] = mergeData(t[o<<1], t[o<<1|1])
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = newData(a[l])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o] = newData(val)
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

func (t seg) query(o, l, r, ql, qr int) data {
	if ql <= l && r <= qr {
		return t[o]
	}
	m := (l + r) / 2
	if qr <= m {
		return t.query(o*2, l, m, ql, qr)
	}
	if ql > m {
		return t.query(o*2+1, m+1, r, ql, qr)
	}
	lRes := t.query(o*2, l, m, ql, qr)
	rRes := t.query(o*2+1, m+1, r, ql, qr)
	return mergeData(lRes, rRes)
}

func newSegmentTreeWithArray(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func resultArray(nums []int, K int, queries [][]int) []int {
	k = K
	t := newSegmentTreeWithArray(nums)
	n := len(nums)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		t.update(1, 0, n-1, q[0], q[1])
		res := t.query(1, 0, n-1, q[2], n-1)
		ans[qi] = res.cnt[q[3]]
	}
	return ans
}
