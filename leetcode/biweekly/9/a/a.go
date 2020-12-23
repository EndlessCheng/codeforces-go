package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxNumberOfApples(a []int) (ans int) {
	sort.Ints(a)
	s := 0
	for _, v := range a {
		s += v
		if s <= 5000 {
			ans++
		}
	}
	return
}
