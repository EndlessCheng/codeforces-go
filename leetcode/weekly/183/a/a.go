package main

import "sort"

func minSubsequence(a []int) (ans []int) {
	s := 0
	for _, v := range a {
		s += v
	}
	sort.Ints(a)
	s2 := 0
	for i := len(a) - 1; i >= 0; i-- {
		v := a[i]
		ans = append(ans, v)
		s2 += v
		s -= v
		if s2 > s {
			break
		}
	}
	return
}
