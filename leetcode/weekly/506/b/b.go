package main

import (
	"maps"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func getLength1(nums []int) (ans int) {
	for i := range nums {
		cnt := map[int]int{}
		cc := map[int]int{}
		for j := i; j < len(nums); j++ {
			x := nums[j]
			if c := cnt[x]; c > 0 {
				cc[c]--
				if cc[c] == 0 {
					delete(cc, c) // 保证我们可以正确计算 cc 的大小
				}
			}
			cnt[x]++
			cc[cnt[x]]++

			if len(cnt) == 1 { // 子数组只有一种元素
				ans = max(ans, j-i+1)
			} else if len(cc) == 2 { // 子数组的元素出现次数恰好有两种
				c := slices.Sorted(maps.Keys(cc))
				if c[0]*2 == c[1] {
					ans = max(ans, j-i+1)
				}
			}
		}
	}
	return
}

func getLength(nums []int) (ans int) {
	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)
	// 直接改 nums
	for i, x := range nums {
		nums[i] = sort.SearchInts(sorted, x)
	}

	for i := range nums {
		cnt := make([]int, m+1)
		cc := make([]int, n+1)
		k := 0
		for j := i; j < len(nums); j++ {
			x := nums[j]
			if c := cnt[x]; c > 0 {
				cc[c]--
				if cc[c] == 0 {
					k--
				}
			}
			cnt[x]++
			if cc[cnt[x]] == 0 {
				k++
			}
			cc[cnt[x]]++

			if len(cnt) == 1 { // 子数组只有一种元素
				ans = max(ans, j-i+1)
			} else if len(cc) == 2 { // 子数组的元素出现次数恰好有两种
				c := slices.Sorted(maps.Keys(cc))
				if c[0]*2 == c[1] {
					ans = max(ans, j-i+1)
				}
			}
		}
	}
	return
}
