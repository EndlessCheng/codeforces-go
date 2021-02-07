package main

import (
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func minAbsDifference(a []int, goal int) (ans int) {
	ans = abs(goal)

	n := len(a)
	if n == 1 {
		return min(ans, abs(a[0]-goal))
	}

	m := n / 2
	b := a[:m]
	x := make([]int, 0, 1<<m)
	calc := func(sub uint) (res int) {
		for ; sub > 0; sub &= sub - 1 {
			res += b[bits.TrailingZeros(sub)]
		}
		return
	}
	for sub := uint(0); sub < 1<<m; sub++ {
		x = append(x, calc(sub))
	}
	sort.Ints(x)

	b = a[m:]
	m = n - m
	for sub := uint(0); sub < 1<<m; sub++ {
		res := calc(sub) - goal
		i := sort.SearchInts(x, -res)
		if i < len(x) {
			ans = min(ans, abs(res+x[i]))
		}
		if i > 0 {
			ans = min(ans, abs(res+x[i-1]))
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
