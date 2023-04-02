package main

import "sort"

// https://space.bilibili.com/206214
func miceAndCheese(reward1, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x
		reward1[i] -= x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}
