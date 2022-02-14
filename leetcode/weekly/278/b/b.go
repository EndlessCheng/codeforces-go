package main

// github.com/EndlessCheng/codeforces-go
func maxScoreIndices(nums []int) []int {
	left, right := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		right += nums[i] // 先算出整个数组的 1 的个数
	}
	maxSum := right
	ans := []int{0}
	for i, v := range nums {
		left += v ^ 1 // 不断增加左边 0 的个数
		right -= v // 减少右边 1 的个数
		if left+right > maxSum { // 更新答案
			maxSum = left + right
			ans = []int{i + 1}
		} else if left+right == maxSum {
			ans = append(ans, i+1)
		}
	}
	return ans
}
