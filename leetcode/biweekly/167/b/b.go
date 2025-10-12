package main

// https://space.bilibili.com/206214
func longestSubarray1(nums []int) int {
	ans := 2
	for i := range nums { // 枚举斐波那契子数组的左端点
		j := i + 2
		for j < len(nums) && nums[j] == nums[j-1]+nums[j-2] {
			j++
		}
		ans = max(ans, j-i) // [i,j-1] 是斐波那契子数组
	}
	return ans
}

func longestSubarray(nums []int) int {
	n := len(nums)
	ans := 2
	start := 0
	for i := 2; i < n; i++ {
		if nums[i] != nums[i-1]+nums[i-2] {
			ans = max(ans, i-start) // [start,i-1] 是斐波那契子数组
			start = i - 1
		}
	}
	return max(ans, n-start) // [start,n-1] 是斐波那契子数组
}
