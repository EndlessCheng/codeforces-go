package main

// https://space.bilibili.com/206214
func transformArray(nums []int) []int {
	cnt := [2]int{}
	for _, x := range nums {
		cnt[x%2]++
	}
	for i := range nums {
		if i < cnt[0] {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}
	return nums
}
