package main

import "slices"

// https://space.bilibili.com/206214
func maximumHappinessSum(happiness []int, k int) (ans int64) {
	slices.SortFunc(happiness, func(a, b int) int { return b - a })
	for i, x := range happiness[:k] {
		if x <= i {
			break
		}
		ans += int64(x - i)
	}
	return
}
