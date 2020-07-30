package main

import (
	//. "nc_tools"
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// github.com/EndlessCheng/codeforces-go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

func (t seg) update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].max += v
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func BoomKill(n, m int, players []*Point) (ans int) {
	m--
	gx := make([][]int, n+1)
	for _, p := range players {
		gx[p.X] = append(gx[p.X], p.Y)
	}
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for i := 1; i <= m; i++ {
		for _, y := range gx[i] {
			t.update(1, y, 1)
		}
	}
	for i := 1; i <= n; i++ {
		if i-m-1 > 0 {
			for _, y := range gx[i-m-1] {
				t.update(1, y, -1)
			}
		}
		if i+m <= n {
			for _, y := range gx[i+m] {
				t.update(1, y, 1)
			}
		}
		for _, y := range gx[i] {
			t.update(1, y, -1)
		}
		ans = max(ans, len(gx[i])+t[1].max)
		for _, y := range gx[i] {
			t.update(1, y, 1)
		}
	}
	return
}
