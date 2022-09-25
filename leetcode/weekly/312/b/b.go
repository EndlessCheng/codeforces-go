package main

// https://space.bilibili.com/206214
func longestSubarray(nums []int) (ans int) {
	max, cnt := 0, 0
	for _, x := range nums {
		if x > max {
			max, ans, cnt = x, 1, 1
		} else if x < max {
			cnt = 0
		} else if cnt++; cnt > ans {
			ans = cnt
		}
	}
	return
}
