package main

import "sort"

func minDays(bloomDay []int, m int, k int) (ans int) {
	ans = sort.Search(1e9+1, func(d int) bool {
		cnt, done := 0, 0
		for _, dd := range bloomDay {
			if dd > d {
				cnt = 0
			} else {
				cnt++
				if cnt == k {
					done++
					cnt = 0
				}
			}
		}
		return done >= m
	})
	if ans > 1e9 {
		ans = -1
	}
	return
}
