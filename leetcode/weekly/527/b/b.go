package main

import "slices"

// https://space.bilibili.com/206214
func countSuffix(nums []int, target []int) int {
	// 大前提：每种元素的出现次数都相等
	mx := max(slices.Max(nums), slices.Max(target))
	diff := make([]int, mx+1) // slice 比 map 快
	for i, x := range nums {
		diff[x]++
		diff[target[i]]--
	}
	for _, d := range diff {
		if d != 0 {
			return 0
		}
	}

	n := len(nums)
	maxI := n - 1
	for i, x := range nums {
		if x != target[i] {
			maxI = i
			break
		}
	}

	minI := 0
	for i := n - 1; i > 0; i-- {
		if target[i-1] < target[i] {
			minI = i
			break
		}
	}

	return max(maxI-minI+1, 0)
}
