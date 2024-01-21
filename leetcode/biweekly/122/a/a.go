package main

import "math"

// https://space.bilibili.com/206214
func minimumCost(nums []int) int {
	fi, se := math.MaxInt, math.MaxInt
	for _, x := range nums[1:] {
		if x < fi {
			se = fi
			fi = x
		} else if x < se {
			se = x
		}
	}
	return nums[0] + fi + se
}
