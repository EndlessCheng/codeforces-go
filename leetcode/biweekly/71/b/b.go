package main

import "cmp"

// github.com/EndlessCheng/codeforces-go
func pivotArray(nums []int, pivot int) []int {
	var a, b, c []int
	for _, x := range nums {
		if x < pivot {
			a = append(a, x)
		} else if x == pivot {
			b = append(b, x)
		} else {
			c = append(c, x)
		}
	}
	return append(append(a, b...), c...)
}

func pivotArray2(nums []int, pivot int) []int {
	a := [3][]int{}
	for _, x := range nums {
		cmp := cmp.Compare(x, pivot)
		a[cmp+1] = append(a[cmp+1], x)
	}
	return append(append(a[0], a[1]...), a[2]...)
}
