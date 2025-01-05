package main

import "slices"

// https://space.bilibili.com/206214
func maxLength(nums []int) (ans int) {
	mx := slices.Max(nums)
	allLcm := 1
	for _, x := range nums {
		allLcm = lcm(allLcm, x)
	}

	for i := range nums {
		m, l, g := 1, 1, 0
		for j := i; j < len(nums) && m <= allLcm*mx; j++ {
			x := nums[j]
			m *= x
			l = lcm(l, x)
			g = gcd(g, x)
			if m == l*g {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
