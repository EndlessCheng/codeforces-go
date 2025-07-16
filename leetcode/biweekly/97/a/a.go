package main

import "slices"

// https://space.bilibili.com/206214
func separateDigits(nums []int) (ans []int) {
	for _, x := range slices.Backward(nums) {
		for ; x > 0; x /= 10 {
			ans = append(ans, x%10)
		}
	}
	slices.Reverse(ans)
	return
}
