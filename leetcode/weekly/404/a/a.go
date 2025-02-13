package main

import "math"

// https://space.bilibili.com/206214
func maxHeightOfTriangle1(red, blue int) int {
	cnt := [2]int{}
	for i := 1; ; i++ {
		cnt[i%2] += i
		if (cnt[0] > red || cnt[1] > blue) && (cnt[0] > blue || cnt[1] > red) {
			return i - 1
		}
	}
}

func f(n, m int) int {
	odd := int(math.Sqrt(float64(n)))
	even := (int(math.Sqrt(float64(m*4+1))) - 1) / 2
	if odd > even {
		return even*2 + 1
	}
	return odd * 2
}

func maxHeightOfTriangle(red, blue int) int {
	return max(f(red, blue), f(blue, red))
}
