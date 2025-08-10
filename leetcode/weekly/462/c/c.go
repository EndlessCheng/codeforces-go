package main

import "slices"

// https://space.bilibili.com/206214
func maxTotal(value, limit []int) (ans int64) {
	n := len(value)
	groups := make([][]int, n+1)
	for i, lim := range limit {
		groups[lim] = append(groups[lim], value[i])
	}
	for lim, a := range groups {
		// 取最大的 lim 个数
		slices.SortFunc(a, func(a, b int) int { return b - a })
		if len(a) > lim {
			a = a[:lim]
		}
		for _, x := range a {
			ans += int64(x)
		}
	}
	return
}
