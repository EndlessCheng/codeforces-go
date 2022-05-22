package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimumLines(a [][]int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	for i, preDY, preDX := 1, 1, 0; i < len(a); i++ {
		dy, dx := a[i][1]-a[i-1][1], a[i][0]-a[i-1][0]
		if dy*preDX != preDY*dx {
			ans++
			preDY, preDX = dy, dx
		}
	}
	return
}
