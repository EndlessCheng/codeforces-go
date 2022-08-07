package main

import "sort"

// https://space.bilibili.com/206214
func mergeSimilarItems(item1 [][]int, item2 [][]int) (ans [][]int) {
	m := map[int]int{}
	for _, p := range item1 {
		m[p[0]] += p[1]
	}
	for _, p := range item2 {
		m[p[0]] += p[1]
	}
	for v, w := range m {
		ans = append(ans, []int{v, w})
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i][0] < ans[j][0] })
	return
}
