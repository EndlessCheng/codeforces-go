package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func closeStrings(s1, s2 string) (ans bool) {
	var c1, c2 [26]int
	for _, b := range s1 {
		c1[b-'a']++
	}
	for _, b := range s2 {
		c2[b-'a']++
	}
	for i, c := range c1[:] {
		if c > 0 != (c2[i] > 0) {
			return
		}
	}
	sort.Ints(c1[:])
	sort.Ints(c2[:])
	return c1 == c2
}
