package main

import "slices"

// https://space.bilibili.com/206214
func toggleLightBulbs(bulbs []int) (ans []int) {
	cnt := map[int]int{}
	for _, i := range bulbs {
		cnt[i] ^= 1
	}
	for i, c := range cnt {
		if c > 0 {
			ans = append(ans, i)
		}
	}
	slices.Sort(ans)
	return
}
