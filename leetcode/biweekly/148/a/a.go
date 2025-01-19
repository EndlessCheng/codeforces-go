package main

// https://space.bilibili.com/206214
func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	ans := abs(nums[0] - nums[n-1])
	for i := 1; i < n; i++ {
		ans = max(ans, abs(nums[i]-nums[i-1]))
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }