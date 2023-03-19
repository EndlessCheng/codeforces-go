package main

import "sort"

// https://space.bilibili.com/206214
func findScore(nums []int) (ans int64) {
	type pair struct{ v, i int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, i + 1} // +1 保证下面 for 循环下标不越界
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.v < b.v || a.v == b.v && a.i < b.i
	})
	vis := make([]bool, len(nums)+2) // 保证下标不越界
	for _, p := range a {
		if !vis[p.i] {
			vis[p.i-1] = true
			vis[p.i+1] = true
			ans += int64(p.v)
		}
	}
	return
}
