package main

// github.com/EndlessCheng/codeforces-go
func maxValueAfterReverse(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		ans += abs(nums[i] - nums[i+1])
	}
	topIntervalBottomLine := int(-1e9)
	bottomIntervalTopLine := int(1e9)
	for i := 0; i < n-1; i++ {
		topIntervalBottomLine = max(topIntervalBottomLine, min(nums[i], nums[i+1]))
		bottomIntervalTopLine = min(bottomIntervalTopLine, max(nums[i], nums[i+1]))
	}
	diff := max(0, (topIntervalBottomLine-bottomIntervalTopLine)*2)

	// Edge case 1: subarray starts at index 0
	for i := 1; i < n-1; i++ {
		diff = max(diff, abs(nums[0]-nums[i+1])-abs(nums[i]-nums[i+1]))
	}
	// Edge case w: subarray ends at index n - 1
	for i := 0; i < n-1; i++ {
		diff = max(diff, abs(nums[n-1]-nums[i])-abs(nums[i+1]-nums[i]))
	}
	return ans + diff
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
