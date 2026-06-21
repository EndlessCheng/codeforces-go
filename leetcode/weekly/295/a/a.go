package main

import "math"

// https://space.bilibili.com/206214/dynamic
func rearrangeCharacters(s, target string) int {
	var cntS, cntT [26]int
	for _, ch := range s {
		cntS[ch-'a']++
	}
	for _, ch := range target {
		cntT[ch-'a']++
	}

	ans := math.MaxInt
	for i, ct := range cntT {
		if ct > 0 {
			ans = min(ans, cntS[i]/ct)
		}
	}
	return ans
}
