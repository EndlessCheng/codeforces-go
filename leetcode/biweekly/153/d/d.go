package main

import (
	"math/bits"
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

// 查询区间最大值，[l,r) 左闭右开
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
	belong := make([]int, n) // 每个 0 所属的区间下标，每个 1 右边最近的 0 区间下标
	a := []pair{{-1, -1}}
	start := 0
	for i, b := range s {
		belong[i] = len(a)
		if i == n-1 || byte(b) != s[i+1] {
			if s[i] == '1' {
				total1 += i - start + 1
			} else {
				a = append(a, pair{start, i + 1})
			}
			start = i + 1
		}
	}
	a = append(a, pair{n + 1, n + 1})

	merge := func(x, y int) int {
		if x > 0 && y > 0 {
			return x + y
		}
		return 0
	}

	st := newST(a)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		ql, qr := q[0], q[1]

		i := belong[ql]
		if ql > 0 && s[ql] == '0' && s[ql-1] == '0' {
			i++ // i 在残缺区间中
		}
		j := belong[qr] - 1
		if qr+1 < n && s[qr] == '0' && s[qr+1] == '1' {
			j++ // j 刚好在完整区间的右端点，无需减一
		}
		qr++

		mx := 0
		if i <= j {
			mx = max(
				st.query(i, j),
				merge(a[i-1].r-ql, a[i].r-a[i].l),
				merge(qr-a[j+1].l, a[j].r-a[j].l),
			)
		} else if i == j+1 {
			mx = merge(a[i-1].r-ql, qr-a[j+1].l)
		}
		ans[qi] = total1 + mx
	}
	return ans
}
