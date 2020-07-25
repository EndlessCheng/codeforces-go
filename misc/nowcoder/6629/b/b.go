package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func solve(n int, a []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Ints(a)
	ans = max(ans, a[1]-a[0])
	for i := 2; i < n; i++ {
		ans = max(ans, a[i]-a[i-2])
	}
	return
}

// 我的憨憨写法
func solve2(n int, a []int) int {
	sort.Ints(a)
	return sort.Search(1e9, func(d int) bool {
		pre := a[0]
		vis := make([]bool, n)
		vis[0] = true
		for i := 1; i < n; i++ {
			st := i
			for ; i < n && a[i]-pre <= d; i++ {
			}
			if i == st {
				return false
			}
			i--
			vis[i] = true
			pre = a[i]
		}
		b := []int{}
		for i, v := range vis {
			if !v {
				b = append(b, a[i])
			}
		}
		for i := len(b) - 1; i >= 0; i-- {
			st := i
			for ; i >= 0 && pre-b[i] <= d; i-- {
			}
			if i == st {
				return false
			}
			i++
			pre = b[i]
		}
		return pre-a[0] <= d
	})
}
