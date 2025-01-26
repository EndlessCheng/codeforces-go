package main

// https://space.bilibili.com/206214
func countPartitions(nums []int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 == 0 {
		return len(nums) - 1
	}
	return 0
}
