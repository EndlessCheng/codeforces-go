package main

import "math"

// https://space.bilibili.com/206214
func maxDifference(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	max1, min0 := 0, math.MaxInt
	for _, c := range cnt {
		if c%2 > 0 {
			max1 = max(max1, c)
		} else if c > 0 {
			min0 = min(min0, c)
		}
	}
	return max1 - min0
}
