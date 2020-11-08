package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxProfit(a []int, orders int) (ans int) {
	const mod int = 1e9 + 7
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	a = append(a, 0)
	for i := 0; i+1 < len(a); i++ {
		if a[i] == a[i+1] {
			continue
		}
		c := i + 1
		if (a[i]-a[i+1])*c > orders {
			m := orders / c
			ans += m * (2*a[i] - m + 1) / 2 * c
			ans += orders % c * (a[i] - m)
			break
		}
		ans += (a[i+1] + 1 + a[i]) * (a[i] - a[i+1]) / 2 * c
		orders -= (a[i] - a[i+1]) * c
	}
	ans %= mod
	return
}
