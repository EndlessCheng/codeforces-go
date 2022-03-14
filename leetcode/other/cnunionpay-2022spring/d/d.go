package main

import (
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func coopDevelop(a [][]int) int {
	ans := len(a) * (len(a) - 1) / 2
	sort.Slice(a, func(i, j int) bool { return len(a[i]) > len(a[j]) }) // 按照数组大小降序排序
	cnt := map[int]int{}
	for _, skill := range a {
		s := 0
		for _, v := range skill {
			s = s<<10 | v
		}
		ans -= cnt[s]
		// 枚举 skill 的所有非空子集
		for i := 1; i < 1<<len(skill); i++ {
			s := 0
			for j, v := range skill {
				if i>>j&1 > 0 {
					s = s<<10 | v
				}
			}
			cnt[s]++ // 加到哈希表中
		}
	}
	return ans % (1e9 + 7)
}
