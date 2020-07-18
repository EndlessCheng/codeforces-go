package main

import . "sort"

// github.com/EndlessCheng/codeforces-go
func solve(n, a, b int) int {
	return Search(a+b, func(x int) bool { return x > 0 && a/x+b/x < n }) - 1
}
