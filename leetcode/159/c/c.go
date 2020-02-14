package main

import "sort"

func balancedString(s string) int {
	n := len(s)
	sum := make([][4]int, n+1)
	for i := range s {
		for j := 0; j < 4; j++ {
			sum[i+1][j] = sum[i][j]
			if s[i] == "QWER"[j] {
				sum[i+1][j]++
			}
		}
	}
	return sort.Search(n, func(sz int) bool {
	outer:
		for i := 0; i+sz <= n; i++ {
			for j := 0; j < 4; j++ {
				if sum[i][j]+sum[n][j]-sum[i+sz][j] > n/4 {
					continue outer
				}
			}
			return true
		}
		return false
	})
}
