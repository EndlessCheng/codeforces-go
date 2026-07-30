package main

import (
	"math"
)

// https://space.bilibili.com/206214
func largestInteger(n, s int) int {
	if s > n*9 {
		return -1
	}

	ans := 0
	for range n {
		d := min(s, 9)
		ans = ans*10 + d
		s -= d
	}
	return ans
}

func largestInteger2(n, s int) int {
	if s > n*9 {
		return -1
	}

	ans := int(math.Pow10(s/9)) - 1 // 填 9
	if s%9 > 0 {
		ans = ans*10 + s%9 // 填 s%9
		n--
	}
	return ans * int(math.Pow10(n-s/9)) // 填 0
}
