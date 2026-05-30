package main

// https://space.bilibili.com/206214
func leftRightDifference(nums []int) []int {
	total := 0
	for _, x := range nums {
		total += x
	}

	leftSum := 0
	for i, x := range nums {
		nums[i] = abs(leftSum*2 + x - total)
		leftSum += x
	}
	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
