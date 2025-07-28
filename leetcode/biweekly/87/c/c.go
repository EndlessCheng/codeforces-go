package main

// https://space.bilibili.com/206214
func smallestSubarrays(nums []int) []int {
	ans := make([]int, len(nums))
	for i, x := range nums { // 计算右端点为 i 的子数组的或值
		ans[i] = 1 // 子数组的长度至少是 1
		// 循环直到 nums[j] 无法增大，其左侧元素也无法增大
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
			ans[j] = i - j + 1 // nums[j] 最后一次增大时的子数组长度就是答案
		}
	}
	return ans
}
