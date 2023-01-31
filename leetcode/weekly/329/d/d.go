package main

import "fmt"

// https://space.bilibili.com/206214
type seg []struct{ l, r, min, todo int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o, v int) {
	t[o].min += v
	t[o].todo += v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
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
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func minCost(nums []int, k int) (ans int) {
	n := len(nums)
	last := make([]int, n)
	last2 := make([]int, n)
	t := make(seg, n*4)
	t.build(1, 1, n)
	for i, x := range nums {
		i++
		t.update(1, i, i, ans) // 相当于设置 f[i+1] 的值
		for _, p := range t {
			fmt.Print(p.min, ", ")
		}
		fmt.Println()
		t.update(1, last[x]+1, i, -1)
		for _, p := range t {
			fmt.Print(p.min, ", ")
		}
		fmt.Println()
		if last[x] > 0 {
			t.update(1, last2[x]+1, last[x], 1)
			for _, p := range t {
				fmt.Print(p.min, ", ")
			}
			fmt.Println()
		}
		ans = k + t.query(1, 1, i)
		fmt.Print(i,ans," ")
		for _, p := range t {
			fmt.Print(p.min, ", ")
		}
		fmt.Println()
		last2[x] = last[x]
		last[x] = i
	}
	return
}

func min(a, b int) int { if a > b { return b }; return a }

