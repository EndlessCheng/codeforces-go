package main

// https://space.bilibili.com/206214
func longestAlternating(nums []int) (ans int) {
	n := len(nums)
	lastEnd, lastLen := 0, 0
	for i := 0; i < n; {
		start := i
		for i++; i < n; i++ {
			if nums[i] == nums[i-1] ||
				i > 1 && (nums[i-2] == nums[i-1] || (nums[i-2] < nums[i-1]) == (nums[i-1] < nums[i])) {
				break
			}
		}
		ans = max(ans, i-start)

		// 删除 nums[start] 后可以拼接
		if lastEnd+2 == start {

			ans = max(ans, i-start+lastEnd)
		}

		lastEnd = i - 1
		lastLen = i-start
	}
	return ans
}
