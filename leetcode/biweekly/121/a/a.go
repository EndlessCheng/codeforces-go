package main

// https://space.bilibili.com/206214
func missingInteger(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		sum += nums[i]
	}

	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}
	for has[sum] { // 至多循环 n 次
		sum++
	}
	return sum
}
