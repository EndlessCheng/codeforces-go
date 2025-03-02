package main

import "slices"

// https://space.bilibili.com/206214
func largestInteger(nums []int, k int) int {
	n := len(nums)
	if k == n {
		return slices.Max(nums)
	}
	if k == 1 {
		cnt := map[int]int{}
		for _, x := range nums {
			cnt[x]++
		}
		ans := -1
		for x, c := range cnt {
			if c == 1 {
				ans = max(ans, x)
			}
		}
		return ans
	}
	return max(f(nums[1:], nums[0]), f(nums[:n-1], nums[n-1]))
}

func f(a []int, x int) int {
	if slices.Contains(a, x) {
		return -1
	}
	return x
}
