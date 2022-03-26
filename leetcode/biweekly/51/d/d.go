package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func closestRoom(a, qs [][]int) []int {
	ans := make([]int, len(qs))
	sort.Slice(a, func(i, j int) bool { return a[i][1] > a[j][1] })
	for i := range qs {
		qs[i] = append(qs[i], i)
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i][1] > qs[j][1] })

	t := redblacktree.NewWithIntComparator()
	i, n := 0, len(a)
	for _, q := range qs {
		for ; i < n && a[i][1] >= q[1]; i++ {
			t.Put(a[i][0], nil)
		}
		tar := q[0]
		res := -1
		if o, _ := t.Floor(tar - 1); o != nil {
			res = o.Key.(int)
		}
		if o, _ := t.Ceiling(tar); o != nil && (res < 0 || o.Key.(int)-tar < tar-res) {
			res = o.Key.(int)
		}
		ans[q[2]] = res
	}
	return ans
}
