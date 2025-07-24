package main

// https://space.bilibili.com/206214
func longestSubarray(nums []int) (ans int) {
	mx, cnt := 0, 0
	for _, x := range nums {
		if x > mx {
			// 重新统计所有内容
			mx = x
			cnt = 1
			ans = 1
		} else if x == mx {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 0
		}
	}
	return
}
