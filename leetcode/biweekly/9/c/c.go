package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func smallestCommonElement(a [][]int) (ans int) {
	mp := map[int]bool{}
	for _, v := range a[0] {
		mp[v] = true
	}
	for _, r := range a[1:] {
		mp2 := map[int]bool{}
		for _, v := range r {
			if mp[v] {
				mp2[v] = true
			}
		}
		mp = mp2
	}
	if len(mp) == 0 {
		return -1
	}
	keys := make([]int, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[0]
}
