package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func Maximumlength(s string) int {
	return 3*sort.Search(len(s)/3+1, func(n int) bool {
		for b, i := byte('a'), 0; b < 'd'; b++ {
			c := 0
			for ; i < len(s) && c < n; i++ {
				if s[i] == b {
					c++
				}
			}
			if c < n {
				return true
			}
		}
		return false
	}) - 3
}
