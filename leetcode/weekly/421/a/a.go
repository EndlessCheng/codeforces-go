package main

import "slices"

// https://space.bilibili.com/206214
func maxScore(nums []int) int64 {
	n := len(nums)
	sufGcd := make([]int, n+1)
	sufLcm := make([]int, n+1)
	sufLcm[n] = 1
	for i, x := range slices.Backward(nums) {
		sufGcd[i] = gcd(sufGcd[i+1], x)
		sufLcm[i] = lcm(sufLcm[i+1], x)
	}

	ans := sufGcd[0] * sufLcm[0] // 不移除元素
	preGcd, preLcm := 0, 1
	for i, x := range nums { // 枚举移除 nums[i]
		ans = max(ans, gcd(preGcd, sufGcd[i+1])*lcm(preLcm, sufLcm[i+1]))
		preGcd = gcd(preGcd, x)
		preLcm = lcm(preLcm, x)
	}
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
