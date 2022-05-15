package main

// github.com/EndlessCheng/codeforces-go
func waysToSplitArray(nums []int) (ans int) {
	sum := 0
	for _, v := range nums { sum += v }
	for i, s := 0, 0; i < len(nums)-1; i++ {
		s += nums[i]
		if s*2 >= sum { ans++ }
	}
	return
}
