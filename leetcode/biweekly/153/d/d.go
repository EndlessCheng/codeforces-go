package main

import (
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
type pair struct{ l, r int } // 左闭右开

type ST [][]int

func newST(a []pair) ST {
	n := len(a) - 1
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, p := range a[:n] {
		st[i] = make([]int, sz)
		st[i][0] = p.r - p.l + a[i+1].r - a[i+1].l
	}
	for j := 1; j < sz; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = max(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// [l,r) 左闭右开
func (st ST) query(l, r int) int {
	if l >= r {
		return 0
	}
	k := bits.Len(uint(r-l)) - 1
	return max(st[l][k], st[r-1<<k][k])
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)
	total1 := 0
	// 统计连续 0 段对应的区间（左闭右开）
	a := []pair{{-1, -1}} // 哨兵
	start := 0
	for i := range n {
		if i == n-1 || s[i] != s[i+1] {
			if s[i] == '1' {
				total1 += i - start + 1
			} else {
				a = append(a, pair{start, i + 1}) // 左闭右开
			}
			start = i + 1
		}
	}
	a = append(a, pair{n + 1, n + 1}) // 哨兵

	calc := func(x, y int) int {
		if x > 0 && y > 0 {
			return x + y
		}
		return 0
	}

	st := newST(a)
	m := len(a)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		ql, qr := q[0], q[1]+1 // 左闭右开
		i := sort.Search(m, func(i int) bool { return a[i].l >= ql })
		j := sort.Search(m, func(i int) bool { return a[i].r > qr }) - 1
		mx := 0
		if i <= j { // [ql,qr) 中有完整的区间
			mx = max(
				st.query(i, j),                   // 相邻完整区间的长度之和的最大值
				calc(a[i-1].r-ql, a[i].r-a[i].l), // i-1 残缺区间 + i
				calc(qr-a[j+1].l, a[j].r-a[j].l), // j+1 残缺区间 + j
			)
		} else if i == j+1 { // [ql,qr) 中有两个相邻的残缺区间
			mx = calc(a[i-1].r-ql, qr-a[j+1].l) // i-1 残缺区间 + j+1 残缺区间
		}
		ans[qi] = total1 + mx
	}
	return ans
}
