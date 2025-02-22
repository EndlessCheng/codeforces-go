package main

import (
	"maps"
	"slices"
)

// https://space.bilibili.com/206214
func beautifulSubsets(nums []int, k int) int {
	groups := map[int]map[int]int{}
	for _, x := range nums {
		// 模 k 同余的数分到同一组，记录元素 x 及其出现次数
		if groups[x%k] == nil {
			groups[x%k] = map[int]int{}
		}
		groups[x%k][x]++
	}

	ans := 1
	for _, cnt := range groups {
		// 计算这一组的方案数
		a := slices.Sorted(maps.Keys(cnt))
		f0, f1, newF := 1, 1<<cnt[a[0]], 0
		for i := 1; i < len(a); i++ {
			c := cnt[a[i]]
			if a[i]-a[i-1] == k {
				newF = f1 + f0*(1<<c-1)
			} else {
				newF = f1 << c
			}
			f0 = f1
			f1 = newF
		}
		ans *= f1 // 每组方案数相乘
	}
	return ans - 1 // 去掉空集
}
