package main

import "slices"

// https://space.bilibili.com/206214
func f(a []int) (mx int) {
	slices.Sort(a)
	for i, n := 0, len(a); i < n; {
		st := i
		for i++; i < n && a[i]-a[i-1] == 1; i++ {}
		mx = max(mx, i-st+1)
	}
	return
}

func maximizeSquareHoleArea(_, _ int, hBars, vBars []int) int {
	mn := min(f(hBars), f(vBars))
	return mn * mn
}
