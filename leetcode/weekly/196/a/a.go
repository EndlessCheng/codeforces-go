package main

import "sort"

func canMakeArithmeticProgression(a []int) (ans bool) {
	sort.Ints(a)
	for i := 2; i < len(a); i++ {
		if a[i]-a[i-1] != a[1]-a[0] {
			return
		}
	}
	return true
}
