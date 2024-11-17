package main

import "math"

// https://space.bilibili.com/206214
func minDifference(nums []int) (ans int) {
	n := len(nums)
	// 和空位相邻的最小数字 minL 和最大数字 maxR
	minL, maxR := math.MaxInt, 0
	for i, v := range nums {
		if v != -1 && (i > 0 && nums[i-1] == -1 || i < n-1 && nums[i+1] == -1) {
			minL = min(minL, v)
			maxR = max(maxR, v)
		}
	}

	updateAns := func(l, r int, big bool) {
		d := (min(r-minL, maxR-l) + 1) / 2
		if big {
			d = min(d, (maxR-minL+2)/3) // d 不能超过上界
		}
		ans = max(ans, d)
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
				updateAns(min(nums[preI], v), max(nums[preI], v), i-preI > 2)
			}
		} else if i > 0 {
			updateAns(v, v, false)
		}
		preI = i
	}
	if 0 <= preI && preI < n-1 {
		updateAns(nums[preI], nums[preI], false)
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
