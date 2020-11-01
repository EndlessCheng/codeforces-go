package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func furthestBuilding(h []int, bricks int, ladders int) (ans int) {
	return sort.Search(len(h), func(r int) bool {
		a := []int{}
		for i := 1; i <= r; i++ {
			if h[i] > h[i-1] {
				a = append(a, h[i]-h[i-1])
			}
		}
		sort.Ints(a)
		s := 0
		for i := 0; i < len(a)-ladders; i++ {
			s += a[i]
		}
		return s > bricks
	}) - 1
}
