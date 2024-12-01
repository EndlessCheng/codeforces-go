package main

import "math"

// https://space.bilibili.com/206214
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, y := range nums {
		t := total - y*2
		if cnt[t] > 1 || cnt[t] > 0 && t != y {
			ans = max(ans, t)
		}
	}
	return ans
}

func getLargestOutlier2(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, x := range nums {
		if (total-x)%2 == 0 {
			y := (total - x) / 2
			if cnt[y] > 1 || cnt[y] > 0 && y != x {
				ans = max(ans, x)
			}
		}
	}
	return ans
}
