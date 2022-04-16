package main

import "math"

// github.com/EndlessCheng/codeforces-go
func findClosestNumber(nums []int) (ans int) {
	min := math.MaxInt32
	for _, x := range nums {
		if y := abs(x); y < min || y == min && x > ans {
			min, ans = y, x
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
