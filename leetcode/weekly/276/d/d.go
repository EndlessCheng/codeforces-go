package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxRunTime(n int, batteries []int) int64 {
	tot := 0
	for _, b := range batteries {
		tot += b
	}
	return int64(sort.Search(tot/n, func(x int) bool {
		x++
		sum := 0
		for _, b := range batteries {
			sum += min(b, x)
		}
		return n*x > sum
	}))
}

func min(a, b int) int { if a > b { return b }; return a }
