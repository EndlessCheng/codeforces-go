package main

// github.com/EndlessCheng/codeforces-go
func check(nums []int) bool {
	n := len(nums)
	sorted := nums[0] >= nums[n-1]
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] { // 严格递减
			if !sorted { // 之前出现过严格递减，说明至少有三个递增段
				return false
			}
			sorted = false // 标记遇到了严格递减
		}
	}
	return true
}
