package main

import "math"

// https://space.bilibili.com/206214
func minLength(nums []int, k int) int {
	cnt := map[int]int{}
	sum := 0
	left := 0
	ans := math.MaxInt

	for i, x := range nums {
		// 1. 入
		cnt[x]++
		if cnt[x] == 1 {
			sum += x
		}

		for sum >= k {
			// 2. 更新答案
			ans = min(ans, i-left+1)

			// 3. 出
			out := nums[left]
			cnt[out]--
			if cnt[out] == 0 {
				sum -= out
			}
			left++
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
