package main

// https://space.bilibili.com/206214
func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			nums[i] ^= (x + 1) &^ x >> 1
		}
	}
	return nums
}

func minBitwiseArray2(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			t := ^x
			nums[i] ^= t & -t >> 1
		}
	}
	return nums
}
