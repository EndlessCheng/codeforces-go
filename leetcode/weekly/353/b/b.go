package main

import (
	"math"
)

// https://space.bilibili.com/206214
type node struct {
	lo, ro   *node
	l, r, mx int
}

func (o *node) get() int {
	if o != nil {
		return o.mx
	}
	return math.MinInt
}

func (o *node) update(i, val int) {
	if o.l == o.r {
		o.mx = max(o.mx, val)
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == nil {
			o.lo = &node{l: o.l, r: m, mx: math.MinInt}
		}
		o.lo.update(i, val)
	} else {
		if o.ro == nil {
			o.ro = &node{l: m + 1, r: o.r, mx: math.MinInt}
		}
		o.ro.update(i, val)
	}
	o.mx = max(o.lo.get(), o.ro.get())
}

func (o *node) query(l, r int) int {
	if o == nil || l > o.r || r < o.l {
		return math.MinInt
	}
	if l <= o.l && o.r <= r {
		return o.mx
	}
	return max(o.lo.query(l, r), o.ro.query(l, r))
}

func maximumJumps(nums []int, target int) int {
	n, mx := len(nums), 0
	rt := &node{l: -3e9, r: 3e9, mx: math.MinInt}
	rt.update(nums[0], 0)
	for i := 1; i<n; i++ {
		mx = rt.query(nums[i]-target, nums[i]+target) + 1
		rt.update(nums[i], mx)
	}
	if mx < 0 {
		return -1
	}
	return mx
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
