package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func getKth(lo, hi, k int) (ans int) {
	type pair struct{ x, y int }
	a := []pair{}
	for i := lo; i <= hi; i++ {
		c := 0
		for v := i; v > 1; {
			c++
			if v&1 > 0 {
				v = 3*v + 1
			} else {
				v /= 2
			}
		}
		a = append(a, pair{c, i})
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
	return a[k-1].y
}
