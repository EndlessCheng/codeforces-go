package main

import "math"

// https://space.bilibili.com/206214
func minOperations(s string) int {
	n := len(s)
	ans := math.MaxInt
	for rot := range n {
		op := rot
		for i := range n / 2 {
			d := abs(int(s[(rot+i)%n]) - int(s[(rot+n-1-i)%n]))
			op += min(d, 26-d)
		}
		ans = min(ans, op)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
