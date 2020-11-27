package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func frequencySort(a []int) (ans []int) {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	sort.Slice(a, func(i, j int) bool {
		v, w := a[i], a[j]
		return cnt[v] < cnt[w] || cnt[v] == cnt[w] && v > w
	})
	return a
}
