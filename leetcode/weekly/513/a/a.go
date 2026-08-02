package main

import "slices"

// https://space.bilibili.com/206214
func maxPairStrength(nums []int) int64 {
	slices.SortFunc(nums, func(a, b int) int { return b - a })

	allGcd := 0
	for _, x := range nums {
		allGcd = gcd(allGcd, x)
	}
	allGcd2 := allGcd * allGcd

	ans := 0
	for i, x := range nums {
		for _, y := range nums[:i] {
			mul := x * y
			if mul/allGcd2 <= ans {
				break
			}
			g := gcd(x, y)
			ans = max(ans, mul/(g*g))
		}
	}
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
