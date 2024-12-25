package main

// github.com/EndlessCheng/codeforces-go
func waysToSplitArray(nums []int) (ans int) {
	total := 0
	for _, x := range nums {
		total += x
	}

	s := 0
	for _, x := range nums[:len(nums)-1] {
		s += x
		if s*2 >= total {
			ans++
		}
	}
	return
}
