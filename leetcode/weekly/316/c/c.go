package main

import "sort"

// https://space.bilibili.com/206214
func minCost(nums, cost []int) int64 {
	type pair struct{ x, c int }
	a := make([]pair, len(nums))
	sumCost := 0
	for i, c := range cost {
		a[i] = pair{nums[i], c}
		sumCost += c
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x < b.x })

	ans, s := 0, 0
	for _, p := range a {
		s += p.c
		if s >= sumCost/2 {
			// 把所有数变成 p.x
			for _, q := range a {
				ans += abs(q.x-p.x) * q.c
			}
			break
		}
	}
	return int64(ans)
}

func abs(x int) int { if x < 0 { return -x }; return x }
