package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxIceCream(a []int, coins int) int {
	sort.Ints(a)
	for i, v := range a {
		if coins < v {
			return i
		}
		coins -= v
	}
	return len(a)
}
