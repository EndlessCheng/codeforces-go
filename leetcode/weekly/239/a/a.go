package main

import "math"

// github.com/EndlessCheng/codeforces-go
func getMinDistance1(nums []int, target int, start int) int {
	ans := math.MaxInt
	for i, x := range nums {
		if x == target {
			ans = min(ans, abs(i-start))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMinDistance(nums []int, target int, start int) int {
	for k := 0; ; k++ {
		if start >= k && nums[start-k] == target ||
			start+k < len(nums) && nums[start+k] == target {
			return k
		}
	}
}
