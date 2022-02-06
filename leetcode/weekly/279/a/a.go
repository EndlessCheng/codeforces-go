package main

import "sort"

// 模拟

// github.com/EndlessCheng/codeforces-go
func sortEvenOdd(nums []int) []int {
	a := [2]sort.IntSlice{}
	for i, v := range nums {
		a[i&1] = append(a[i&1], v)
	}
	sort.Sort(a[0])
	sort.Sort(sort.Reverse(a[1]))
	nums = nums[:0]
	for i, v := range a[0] {
		nums = append(nums, v)
		if i < len(a[1]) {
			nums = append(nums, a[1][i])
		}
	}
	return nums
}
