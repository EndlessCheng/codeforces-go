package main

import (
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func minWastedSpace(a []int, boxes [][]int) int {
	sort.Ints(a)
	ans := math.MaxInt64
	for _, b := range boxes {
		sort.Ints(b)
		if b[len(b)-1] < a[len(a)-1] {
			continue
		}
		s, l := 0, 0
		for _, v := range b {
			r := sort.SearchInts(a, v+1)
			s += (r - l) * v
			l = r
		}
		ans = min(ans, s)
	}
	if ans == math.MaxInt64 {
		return -1
	}
	for _, v := range a {
		ans -= v
	}
	return ans % (1e9 + 7)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
