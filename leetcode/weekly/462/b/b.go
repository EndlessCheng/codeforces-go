package main

// https://space.bilibili.com/206214
func sortPermutation(nums []int) int {
	ans := -1 // 二进制全为 1
	for i, x := range nums {
		if i != x {
			ans &= x
		}
	}
	return max(ans, 0)
}
