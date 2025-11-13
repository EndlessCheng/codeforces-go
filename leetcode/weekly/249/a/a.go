package main

// github.com/EndlessCheng/codeforces-go
func getConcatenation1(nums []int) []int {
	return append(nums, nums...)
}

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i, x := range nums {
		ans[i] = x
		ans[i+n] = x
	}
	return ans
}
