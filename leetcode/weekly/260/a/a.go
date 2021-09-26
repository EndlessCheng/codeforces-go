package main

// github.com/EndlessCheng/codeforces-go
func maximumDifference(nums []int) (ans int) {
	preMin := nums[0]
	for j := 1; j < len(nums); j++ {
		ans = max(ans, nums[j]-preMin)
		preMin = min(preMin, nums[j])
	}
	if ans == 0 {
		ans--
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
