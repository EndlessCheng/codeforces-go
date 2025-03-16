package main

// github.com/EndlessCheng/codeforces-go
func subsetXORSum(nums []int) int {
	or := 0
	for _, x := range nums {
		or |= x
	}
	return or << (len(nums) - 1)
}
