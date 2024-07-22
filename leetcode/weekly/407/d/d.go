package main

// https://space.bilibili.com/206214
func minimumOperations(nums, target []int) int64 {
	for i ,x := range nums {
		target[i] -= x
	}
	ans := target[0] - nums[0]
	for i := 1; i < len(nums); i++ {
		k := (target[i] - target[i-1]) - (nums[i] - nums[i-1])
		ans += max(k, 0)
	}
	return int64(ans)
}
