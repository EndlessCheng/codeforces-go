package main

// https://space.bilibili.com/206214
func missingMultiple(nums []int, k int) int {
	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}
	for x := k; ; x += k {
		if !has[x] {
			return x
		}
	}
}
