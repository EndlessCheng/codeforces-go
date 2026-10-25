package main

// https://space.bilibili.com/206214
func countSubarrays(nums []int, minK, maxK int) (ans int64) {
	end := -1
	for i, x := range nums {
		if x < minK || x > maxK {
			end = i
		} else {
			ans += int64(i - end)
		}
	}
	return
}
