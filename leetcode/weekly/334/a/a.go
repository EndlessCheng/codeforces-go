package main

// https://space.bilibili.com/206214
func leftRigthDifference(nums []int) []int {
	rightSum := 0
	for _, x := range nums {
		rightSum += x
	}
	leftSum := 0
	for i, x := range nums {
		rightSum -= x
		nums[i] = abs(leftSum - rightSum)
		leftSum += x
	}
	return nums
}

func abs(x int) int { if x < 0 { return -x }; return x }
