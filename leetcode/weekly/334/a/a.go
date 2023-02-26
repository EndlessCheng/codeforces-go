package main

// https://space.bilibili.com/206214
func leftRigthDifference(nums []int) []int {
	n := len(nums)
	rightSum := make([]int, n)
	for i := n - 1; i > 0; i-- {
		rightSum[i-1] = rightSum[i] + nums[i]
	}

	ans := make([]int, n)
	leftSum := 0
	for i, x := range nums {
		ans[i] = abs(leftSum - rightSum[i])
		leftSum += x
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
