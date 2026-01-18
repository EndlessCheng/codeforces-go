package main

// https://space.bilibili.com/206214
func minOperations(nums, target []int) int {
	set := map[int]struct{}{}
	for i, x := range nums {
		if x != target[i] {
			set[x] = struct{}{}
		}
	}
	return len(set)
}
