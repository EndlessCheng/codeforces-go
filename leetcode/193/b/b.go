package main

import "sort"

func findLeastNumOfUniqueInts(a []int, k int) (ans int) {
	cnts := map[int]int{}
	for _, v := range a {
		cnts[v]++
	}
	c := make([]int, 0, len(cnts))
	for _, v := range cnts {
		c = append(c, v)
	}
	sort.Ints(c)
	for len(c) > 0 && c[0] <= k {
		k -= c[0]
		c = c[1:]
	}
	return len(c)
}
