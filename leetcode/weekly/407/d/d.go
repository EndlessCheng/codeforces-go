package main

// https://space.bilibili.com/206214
func minimumOperations(nums, target []int) int64 {
	n := len(nums)
	ans := max(target[0]-nums[0], 0)
	for i := 1; i < n; i++ {
		ans += max((target[i]-nums[i])-(target[i-1]-nums[i-1]), 0)
	}
	ans += max(-(target[n-1] - nums[n-1]), 0)
	return int64(ans)
}
