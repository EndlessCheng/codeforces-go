package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return b.x - a.x })
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}
	slices.SortFunc(qid, func(i, j int) int { return queries[j][0] - queries[i][0] })

	ans := make([]int, len(queries))
	type data struct{ y, s int }
	st := []data{}
	j := 0
	for _, i := range qid {
		x, y := queries[i][0], queries[i][1]
		for ; j < len(a) && a[j].x >= x; j++ { // 下面只需关心 a[j].y
			for len(st) > 0 && st[len(st)-1].s <= a[j].x+a[j].y { // a[j].y >= st[len(st)-1].y
				st = st[:len(st)-1]
			}
			if len(st) == 0 || st[len(st)-1].y < a[j].y {
				st = append(st, data{a[j].y, a[j].x + a[j].y})
			}
		}
		p := sort.Search(len(st), func(i int) bool { return st[i].y >= y })
		if p < len(st) {
			ans[i] = st[p].s
		} else {
			ans[i] = -1
		}
	}
	return ans
}
