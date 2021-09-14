package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumRemovals(s, p string, id []int) int {
	return sort.Search(len(id), func(k int) bool {
		del := make([]bool, len(s))
		for _, i := range id[:k+1] {
			del[i] = true
		}
		j := 0
		for i := range s {
			if !del[i] && s[i] == p[j] {
				if j++; j == len(p) {
					return false
				}
			}
		}
		return true
	})
}
