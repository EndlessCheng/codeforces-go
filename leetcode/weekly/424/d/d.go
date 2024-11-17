package main

import "math"

// https://space.bilibili.com/206214
func minDifference(nums []int) (ans int) {
	// 和空位相邻的最小数字 minL 和最大数字 maxR
	minL, maxR := math.MaxInt, 0
	for i, v := range nums {
		if v != -1 && (i > 0 && nums[i-1] == -1 || i < len(nums)-1 && nums[i+1] == -1) {
			minL = min(minL, v)
			maxR = max(maxR, v)
		}
	}

	preI := -1
	for i, v := range nums {
		if v == -1 {
			continue
		}
		if preI >= 0 {
			if i-preI == 1 {
				ans = max(ans, abs(v-nums[preI]))
			} else {
				l, r := min(nums[preI], v), max(nums[preI], v)
				d := (min(r-minL, maxR-l) + 1) / 2
				if i-preI > 2 {
					d = min(d, (maxR-minL+2)/3) // d 不能超过上界
				}
				ans = max(ans, d)
			}
		}
		preI = i
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
