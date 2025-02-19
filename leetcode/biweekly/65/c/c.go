package main

import (
	"slices"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] })
	k := 0
	for _, item := range items[1:] {
		if item[1] > items[k][1] {
			k++
			items[k] = item
		}
	}

	for i, q := range queries {
		j := sort.Search(k+1, func(i int) bool { return items[i][0] > q })
		if j > 0 {
			queries[i] = items[j-1][1]
		} else {
			queries[i] = 0
		}
	}
	return queries
}

func maximumBeauty1(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] })
	idx := make([]int, len(queries))
	for i := range queries {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return queries[i] - queries[j] })

	ans := make([]int, len(queries))
	maxBeauty, j := 0, 0
	for _, i := range idx {
		q := queries[i]
		// 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
		for j < len(items) && items[j][0] <= q {
			maxBeauty = max(maxBeauty, items[j][1])
			j++
		}
		ans[i] = maxBeauty
	}
	return ans
}
