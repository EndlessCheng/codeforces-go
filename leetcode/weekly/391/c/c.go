package main

// https://space.bilibili.com/206214
func countAlternatingSubarrays(nums []int) (ans int64) {
	cnt := 0
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			cnt = 1
		} else {
			cnt++
		}
		ans += int64(cnt)
	}
	return
}
