package main

func kLengthApart(nums []int, k int) bool {
	last1 := -k - 1
	for i, x := range nums {
		if x != 1 {
			continue
		}
		if i-last1 <= k {
			return false
		}
		last1 = i
	}
	return true
}
