package main

import "sort"

func getStrongest(a []int, k int) (ans []int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	sort.Ints(a)
	m := a[(len(a)-1)/2]
	sort.Slice(a, func(i, j int) bool { x, y := a[i], a[j]; return abs(x-m) > abs(y-m) || abs(x-m) == abs(y-m) && x > y })
	return a[:k]
}
