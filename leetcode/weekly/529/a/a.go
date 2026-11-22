package main

// https://space.bilibili.com/206214
func countDominantElements(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return 1
	}
	if nums[0] >= nums[1] {
		ans = 1
	}
	for i := 1; i < n-1; i++ {
		if nums[i]*2 >= nums[i-1]+nums[i+1] {
			ans++
		}
	}
	if nums[n-1] >= nums[n-2] {
		ans++
	}
	return
}
