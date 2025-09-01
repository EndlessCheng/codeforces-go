package main

import "slices"

// https://space.bilibili.com/206214/dynamic
func fillCups(amount []int) int {
	s := amount[0] + amount[1] + amount[2]
	mx := slices.Max(amount)
	return max((s+1)/2, mx)
}
