package main

import "math"

// https://space.bilibili.com/206214
func uniformArray(nums1 []int) bool {
	minOdd := math.MaxInt
	for _, x := range nums1 {
		if x%2 != 0 {
			minOdd = min(minOdd, x)
		}
	}

	// 没有奇数，都是偶数
	if minOdd == math.MaxInt {
		return true
	}

	for _, x := range nums1 {
		// 把偶数减去奇数，变成奇数，前提是偶数 > 奇数
		if x%2 == 0 && x < minOdd {
			return false
		}
	}

	return true
}
