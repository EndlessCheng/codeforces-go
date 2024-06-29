package main

import "slices"

// https://space.bilibili.com/206214
func miceAndCheese(r1, r2 []int, k int) (ans int) {
	for i, x := range r2 {
		ans += x // 先全部给第二只老鼠
		r1[i] -= x
	}
	slices.SortFunc(r1, func(a, b int) int { return b - a })
	for _, x := range r1[:k] {
		ans += x
	}
	return
}
