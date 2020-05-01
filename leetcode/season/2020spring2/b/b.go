package main

import "sort"

func minTime(a []int, m int) (ans int) {
	n := len(a)
	return sort.Search(1e9, func(t int) bool {
		cnt := 0
		for i := 0; i < n; {
			s, max := 0, 0
			for ; i < n; i++ {
				if a[i] > max {
					max = a[i]
				}
				s += a[i]
				if s-max > t {
					break
				}
			}
			cnt++
		}
		return cnt <= m
	})
}
