package main

// https://space.bilibili.com/206214
func sumOfGoodNumbers(nums []int, k int) (ans int) {
	for i, x := range nums {
		if (i < k || x > nums[i-k]) && (i+k >= len(nums) || x > nums[i+k]) {
			ans += x
		}
	}
	return
}
