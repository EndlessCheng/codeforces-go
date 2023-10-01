package main

import (
	"math"
)

// https://space.bilibili.com/206214
func minSizeSubarray(nums []int, target int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	ans := math.MaxInt
	left, sum, n := 0, 0, len(nums)
	for right := 0; right < n*2; right++ {
		sum += nums[right%n]
		for sum > target%total {
			sum -= nums[left%n]
			left++
		}
		if sum == target%total {
			ans = min(ans, right-left+1)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + target/total*n
}

func min(a, b int) int { if b < a { return b }; return a }
