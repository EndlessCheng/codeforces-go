package main

// https://space.bilibili.com/206214
func minOperations(nums []int, k int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	return s % k
}
