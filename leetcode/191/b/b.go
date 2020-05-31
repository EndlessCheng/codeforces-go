package main

import "sort"

func maxArea(h int, w int, a []int, b []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	a = append(append([]int{0}, a...), h)
	b = append(append([]int{0}, b...), w)
	f := func(a []int) (mx int) {
		sort.Ints(a)
		for i := 1; i < len(a); i++ {
			mx = max(mx, a[i]-a[i-1])
		}
		return
	}
	ans = f(a) * f(b)
	ans %= 1e9 + 7
	return
}
