package main

import "sort"

// https://space.bilibili.com/206214
func sumOfPower(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	s := 0
	for _, x := range nums {
		ans = (ans + x*x%mod*(x+s)) % mod
		s = (s*2 + x) % mod
	}
	return
}
