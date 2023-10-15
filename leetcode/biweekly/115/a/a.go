package main

import "strconv"

// https://space.bilibili.com/206214
func lastVisitedIntegers(words []string) (ans []int) {
	nums := []int{}
	k := 0
	for _, s := range words {
		if s[0] != 'p' { // 不是 prev
			x, _ := strconv.Atoi(s)
			nums = append(nums, x)
			k = 0
		} else {
			k++
			if k > len(nums) {
				ans = append(ans, -1)
			} else {
				ans = append(ans, nums[len(nums)-k]) // 倒数第 k 个
			}
		}
	}
	return
}
