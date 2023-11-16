package main

// https://space.bilibili.com/206214
func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 {
			i++ // 直接跳过
			continue
		}
		start := i // 记录这一组的开始位置
		i++ // 开始位置已经满足要求，从下一个位置开始判断
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		// 从 start 到 i-1 是满足题目要求的子数组
		ans = max(ans, i-start)
	}
	return
}
