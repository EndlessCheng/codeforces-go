package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimalKSum(a []int, k int) int64 {
	ans := 0
	a = append(a, 0, 2e9) // 加入两个哨兵
	sort.Ints(a)
	for i := 1; ; i++ {
		fill := a[i] - a[i-1] - 1 // 可以填充的数字个数
		if fill <= 0 { continue } // 没有可以填充的位置
		if fill >= k {
			ans += (a[i-1]*2 + 1 + k) * k / 2 // 等差数列求和
			return int64(ans)
		}
		ans += (a[i-1] + a[i]) * fill / 2 // 等差数列求和（+1 和 -1 抵消）
		k -= fill
	}
}
