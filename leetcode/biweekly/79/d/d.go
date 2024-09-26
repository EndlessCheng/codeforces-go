package main

import "math/bits"

type seg []struct{ l, r, min, sum int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// 把下标 i 上的元素值增加 val
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].min += val
		t[o].sum += val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].min = min(lo.min, ro.min)
	t[o].sum = lo.sum + ro.sum
}

// 返回区间 [l,r] 内的元素和
func (t seg) querySum(o, l, r int) (sum int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		sum = t.querySum(o<<1, l, r)
	}
	if r > m {
		sum += t.querySum(o<<1|1, l, r)
	}
	return
}

// 返回区间 [0,r] 中 <= val 的最靠左的位置，不存在时返回 -1
func (t seg) findFirst(o, r, val int) int {
	if t[o].min > val {
		return -1 // 整个区间的元素值都大于 val
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	m := (t[o].l + t[o].r) / 2
	if t[o*2].min <= val {
		return t.findFirst(o*2, r, val)
	}
	if r > m {
		return t.findFirst(o*2+1, r, val)
	}
	return -1
}

type BookMyShow struct {
	seg
	n, m int
}

func Constructor(n, m int) BookMyShow {
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	return BookMyShow{t, n, m}
}

func (t *BookMyShow) Gather(k, maxRow int) []int {
	// 找第一个能倒入 k 升水的水桶
	r := t.findFirst(1, maxRow, t.m-k)
	if r < 0 { // 没有这样的水桶
		return nil
	}
	c := t.querySum(1, r, r)
	t.update(1, r, k) // 倒水
	return []int{r, c}
}

func (t *BookMyShow) Scatter(k, maxRow int) bool {
	// [0,maxRow] 的接水量之和
	s := t.querySum(1, 0, maxRow)
	if s > t.m*(maxRow+1)-k {
		return false // 水桶已经装了太多的水
	}
	// 从第一个没有装满的水桶开始
	i := t.findFirst(1, maxRow, t.m-1)
	for k > 0 {
		left := min(t.m-t.querySum(1, i, i), k)
		t.update(1, i, left) // 倒水
		k -= left
		i++
	}
	return true
}
