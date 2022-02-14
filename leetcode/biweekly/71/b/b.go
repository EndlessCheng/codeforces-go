package main

// 模拟

// github.com/EndlessCheng/codeforces-go
func pivotArray(nums []int, pivot int) []int {
	var a, b, c []int
	for _, v := range nums {
		if v < pivot {
			a = append(a, v)
		} else if v == pivot {
			b = append(b, v)
		} else {
			c = append(c, v)
		}
	}
	return append(append(a, b...), c...)
}
