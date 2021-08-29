package main

import "sort"

// Go 自定义排序

// github.com/EndlessCheng/codeforces-go
func kthLargestNumber(a []string, k int) string {
	sort.Slice(a, func(i, j int) bool {
		s, t := a[i], a[j]
		if len(s) != len(t) {
			return len(s) > len(t)
		}
		return s > t
	})
	return a[k-1]
}
