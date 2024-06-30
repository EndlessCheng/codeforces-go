package main

import "math"

// https://space.bilibili.com/206214
func f(n, m int) int {
	odd := int(math.Sqrt(float64(n)))
	even := int((math.Sqrt(float64(m*4+1)) - 1) / 2)
	if odd > even {
		return even*2 + 1
	}
	return odd * 2
}

func maxHeightOfTriangle(red, blue int) int {
	return max(f(red, blue), f(blue, red))
}
