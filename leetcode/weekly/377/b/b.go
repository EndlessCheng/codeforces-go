package main

import "slices"

// https://space.bilibili.com/206214
func f(a []int, mx int) map[int]bool {
	a = append(a, 1, mx)
	slices.Sort(a)
	set := map[int]bool{}
	for i, x := range a {
		for _, y := range a[i+1:] {
			set[y-x] = true
		}
	}
	return set
}

func maximizeSquareArea(m, n int, hFences, vFences []int) int {
	h := f(hFences, m)
	v := f(vFences, n)
	ans := 0
	for x := range h {
		if v[x] {
			ans = max(ans, x)
		}
	}
	if ans == 0 {
		return -1
	}
	return ans * ans % 1_000_000_007
}
