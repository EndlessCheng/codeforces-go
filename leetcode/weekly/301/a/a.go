package main

import "sort"

// https://space.bilibili.com/206214/dynamic
func fillCups(a []int) (ans int) {
	for sort.Ints(a); a[1] > 0; sort.Ints(a) {
		ans++
		a[1]--
		a[2]--
	}
	return ans + a[2]
}
