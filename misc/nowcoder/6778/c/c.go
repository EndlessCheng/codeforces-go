package main

import (
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func MaxMin(a []int) int64 {
	type pair struct{ x, y int }
	n := len(a)
	st := make([][17]pair, n)
	for i, v := range a {
		st[i][0] = pair{v, v}
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j].x = min(st[i][j-1].x, st[i+1<<(j-1)][j-1].x)
			st[i][j].y = max(st[i][j-1].y, st[i+1<<(j-1)][j-1].y)
		}
	}
	s := 0
	for r := 1; r <= n; r++ {
		s += sort.Search(r, func(l int) bool {
			k := bits.Len(uint(r-l)) - 1
			x, y := st[l][k], st[r-1<<k][k]
			return max(x.y, y.y) < min(x.x, y.x)<<1
		})
	}
	return int64(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
