package main

import "math/bits"

// https://space.bilibili.com/206214/dynamic
func countExcellentPairs(nums []int, k int) (ans int64) {
	const U = 30
	vis := map[int]bool{}
	cnt := [U]int{}
	for _, x := range nums {
		if !vis[x] {
			vis[x] = true
			cnt[bits.OnesCount(uint(x))]++
		}
	}
	s := 0
	for i := k; i < U; i++ {
		s += cnt[i]
	}
	for cx, ccx := range cnt {
		ans += int64(ccx) * int64(s)
		cy := k - 1 - cx
		if 0 <= cy && cy < U {
			s += cnt[cy]
		}
	}
	return
}
