package main

import "slices"

// https://space.bilibili.com/206214
func maxTotal(value, limit []int) (ans int64) {
	groups := make([][]int, len(value)+1)
	for i, lim := range limit {
		groups[lim] = append(groups[lim], value[i])
	}

	for lim, a := range groups {
		if lim < len(a) {
			// 只取最大的 lim 个数
			slices.Sort(a)
			a = a[len(a)-lim:]
		}
		for _, x := range a {
			ans += int64(x)
		}
	}
	return
}
