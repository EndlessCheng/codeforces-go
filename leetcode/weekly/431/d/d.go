package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maximumWeight(intervals [][]int) []int {
	type tuple struct{ l, r, weight, i int }
	a := make([]tuple, len(intervals))
	for i, interval := range intervals {
		a[i] = tuple{interval[0], interval[1], interval[2], i}
	}
	slices.SortFunc(a, func(a, b tuple) int { return a.r - b.r })

	n := len(intervals)
	type pair struct {
		sum int
		id  []int
	}
	f := make([][5]pair, n+1)
	for i, t := range a {
		k := sort.Search(i, func(k int) bool { return a[k].r >= t.l })
		for j := 1; j < 5; j++ {
			s1 := f[i][j].sum
			// 为什么是 f[k] 不是 f[k+1]：上面算的是 >= t.l，-1 后得到 < t.l，但由于还要 +1，抵消了
			s2 := f[k][j-1].sum + t.weight
			if s1 > s2 {
				f[i+1][j] = f[i][j]
				continue
			}
			newId := slices.Clone(f[k][j-1].id)
			newId = append(newId, t.i)
			slices.Sort(newId)
			if s1 == s2 && slices.Compare(f[i][j].id, newId) < 0 {
				newId = f[i][j].id
			}
			f[i+1][j] = pair{s2, newId}
		}
	}
	return f[n][4].id
}
