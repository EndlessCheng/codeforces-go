package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minCost(nums, cost []int) (ans int64) {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	sumCost := int64(0)
	for i, c := range cost {
		a[i] = pair{nums[i], c}
		sumCost += int64(c)
	}
	slices.SortFunc(a, func(p, q pair) int { return p.x - q.x })

	s, mid := int64(0), (sumCost+1)/2
	for _, p := range a {
		s += int64(p.c)
		if s >= mid {
			for _, q := range a {
				ans += int64(abs(q.x-p.x)) * int64(q.c)
			}
			break
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }

func minCost2(nums, cost []int) int64 {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, cost[i]}
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x < b.x })

	var total, sumCost int64
	for _, p := range a {
		total += int64(p.c) * int64(p.x-a[0].x)
		sumCost += int64(p.c)
	}
	ans := total
	for i := 1; i < len(a); i++ {
		sumCost -= int64(a[i-1].c * 2)
		total -= sumCost * int64(a[i].x-a[i-1].x)
		ans = min(ans, total)
	}
	return ans
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
