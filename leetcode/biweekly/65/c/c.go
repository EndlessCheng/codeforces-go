package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool { return items[i][0] < items[j][0] }) // 按价格排序
	for i, q := range queries {
		queries[i] = q<<32 | i // 这样排序时可以保留查询的下标
	}
	sort.Ints(queries)

	ans := make([]int, len(queries))
	maxBeauty, i := 0, 0
	for _, q := range queries {
		for ; i < len(items) && items[i][0] <= q>>32; i++ {
			if items[i][1] > maxBeauty {
				maxBeauty = items[i][1]
			}
		}
		ans[q&(1<<32-1)] = maxBeauty
	}
	return ans
}
