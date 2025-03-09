package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
type seg []int

func (t seg) maintain(o int) {
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 初始化线段树
func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 找区间内的第一个 >= x 的数，并更新为 -1，返回这个数的下标（没有则返回 -1）
func (t seg) findFirstAndUpdate(o, l, r, x int) int {
	if t[o] < x { // 区间没有 >= x 的数
		return -1
	}
	if l == r {
		t[o] = -1 // 更新为 -1，表示不能放水果
		return l
	}
	m := (l + r) >> 1
	i := t.findFirstAndUpdate(o<<1, l, m, x) // 先递归左子树
	if i < 0 { // 左子树没找到
		i = t.findFirstAndUpdate(o<<1|1, m+1, r, x) // 再递归右子树
	}
	t.maintain(o)
	return i
}

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func numOfUnplacedFruits(fruits, baskets []int) (ans int) {
	t := newSegmentTree(baskets)
	for _, x := range fruits {
		if t.findFirstAndUpdate(1, 0, len(baskets)-1, x) < 0 {
			ans++
		}
	}
	return
}
