package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func canDistribute(a []int, quantity []int) (ans bool) {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	n := len(cnt)
	a = make([]int, 0, n)
	for _, v := range cnt {
		a = append(a, v)
	}
	m := len(quantity)

	dp := make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, 1<<m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int8
	f = func(p, subset int) (res int8) {
		if subset == 0 {
			return 1
		}
		if p == n {
			return
		}
		dv := &dp[p][subset]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		sub := subset
		for ok := true; ok; ok = sub != subset {
			s := 0
			for mask := uint(sub); mask > 0; mask &= mask - 1 {
				s += quantity[bits.TrailingZeros(mask)]
			}
			if s <= a[p] && f(p+1, subset^sub) > 0 {
				return 1
			}
			sub = (sub - 1) & subset
		}
		return
	}
	return f(0, 1<<m-1) > 0
}
