package main

import "sort"

// https://space.bilibili.com/206214
func minimumRightShifts(a []int) int {
	i, n := 1, len(a)
	for i < n && a[i-1] < a[i] {
		i++
	}
	if i == n {
		return 0
	}
	if a[0] < a[n-1] {
		return -1
	}
	mid := i
	i++
	for i < n && a[i-1] < a[i] {
		i++
	}
	if i < n {
		return -1
	}
	return n - mid
}

func minimumRightShifts2(a []int) int {
	for i := 0; i < len(a); i++ {
		if sort.IntsAreSorted(a) {
			return i
		}
		a = append(a[len(a)-1:], a[:len(a)-1]...)
	}
	return -1
}
