package main

// https://space.bilibili.com/206214
type seg []struct{ l, r, max int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) modify(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].max = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.modify(o<<1, i, val)
	} else {
		t.modify(o<<1|1, i, val)
	}
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

// 返回区间 [l,r] 内的最大值
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func lengthOfLIS(nums []int, k int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	t := make(seg, mx*4)
	t.build(1, 1, mx)
	for _, x := range nums {
		if x == 1 {
			t.modify(1, 1, 1)
		} else {
			t.modify(1, x, 1+t.query(1, max(x-k, 1), x-1))
		}
	}
	return t[1].max
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
