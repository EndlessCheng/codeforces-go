package main

import "math"

// https://space.bilibili.com/206214
func minXor(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt
	}
	for i := 1; i <= k; i++ {
		for j := n - (k - i); j >= i; j-- {
			res := math.MaxInt
			s := 0
			for l := j - 1; l >= i-1; l-- {
				s ^= nums[l]
				res = min(res, max(f[l], s))
			}
			f[j] = res
		}
	}
	return f[n]
}
