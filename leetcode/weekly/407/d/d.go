package main

// https://space.bilibili.com/206214
func minimumOperations(nums, target []int) int64 {
	posSum, negSum := 0, 0
	d := target[0] - nums[0]
	if d > 0 {
		posSum = d
	} else {
		negSum = -d
	}
	for i := 1; i < len(nums); i++ {
		d := (target[i] - nums[i]) - (target[i-1] - nums[i-1])
		if d > 0 {
			posSum += d
		} else {
			negSum -= d
		}
	}
	return int64(max(posSum, negSum))
}
