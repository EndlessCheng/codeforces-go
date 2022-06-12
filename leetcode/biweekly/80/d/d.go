package main

// https://space.bilibili.com/206214/dynamic
func countSubarrays(nums []int, k int64) (ans int64) {
	sum, left := int64(0), 0
	for right, num := range nums {
		sum += int64(num)
		for sum*int64(right-left+1) >= k {
			sum -= int64(nums[left])
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
