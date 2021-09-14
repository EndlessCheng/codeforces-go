package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func check(a []int) (ans bool) {
	for i := 0; i < len(a); i++ {
		if sort.IntsAreSorted(a) {
			return true
		}
		a = append(a[1:], a[:1]...)
	}
	return
}
