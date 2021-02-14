package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimumSize(a []int, maxOperations int) int {
	return sort.Search(1e9, func(mx int) bool {
		if mx == 0 {
			return false
		}
		c := 0
		for _, v := range a {
			c += (v - 1) / mx
		}
		return c <= maxOperations
	})
}
