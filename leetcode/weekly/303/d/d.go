package main

import "math/bits"

// https://space.bilibili.com/206214/dynamic
func countExcellentPairs(nums []int, k int) (ans int64) {
	vis := map[int]bool{}
	cnt := map[int]int{}
	for _, x := range nums {
		if !vis[x] {
			vis[x] = true
			cnt[bits.OnesCount(uint(x))]++
		}
	}
	for cx, ccx := range cnt {
		for cy, ccy := range cnt {
			if cx+cy >= k { // (x,y) 是优质数对
				ans += int64(ccx) * int64(ccy)
			}
		}
	}
	return
}
