package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumUnits(a [][]int, truckSize int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return a[i][1] > a[j][1] })
	for _, p := range a {
		if p[0] < truckSize {
			truckSize -= p[0]
			ans += p[0] * p[1]
		} else {
			ans += truckSize * p[1]
			break
		}
	}
	return
}
