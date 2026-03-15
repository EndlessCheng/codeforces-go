package main

import "slices"

// https://space.bilibili.com/206214
func gcdSum(nums []int) (ans int64) {
	n := len(nums)
	pre := make([]int, n)
	mx := 0
	for i, x := range nums {
		mx = max(mx, x)
		pre[i] = gcd(x, mx)
	}

	slices.Sort(pre)
	for i := range n / 2 {
		ans += int64(gcd(pre[i], pre[n-1-i]))
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
