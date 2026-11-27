package main

import "math"

// github.com/EndlessCheng/codeforces-go
func minimumAverageDifference(nums []int) (ans int) {
	n := len(nums)
	total := 0
	for _, x := range nums {
		total += x
	}

	s := 0
	minDiff := math.MaxInt
	for i, x := range nums[:n-1] {
		s += x
		diff := abs(s/(i+1) - (total-s)/(n-1-i))
		if diff < minDiff {
			minDiff = diff
			ans = i
		}
	}

	// 特殊处理 i = n-1 的情况
	diff := total / n
	if diff < minDiff {
		ans = n - 1
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
