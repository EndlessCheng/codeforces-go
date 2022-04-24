package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func intersection(nums [][]int) (ans []int) {
	cnt := map[int]int{}
	for i, r := range nums {
		for j, v := range r {
			_, _ = i, j
			cnt[v]++
		}
	}
	for v, c := range cnt {
		if c == len(nums) {
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return
}
