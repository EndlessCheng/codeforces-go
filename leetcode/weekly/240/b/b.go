package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxDistance(x, y []int) (ans int) {
	for j, v := range y {
		i := sort.Search(min(len(x), j), func(i int) bool { return x[i] <= v })
		if i < min(len(x), j) && j-i > ans {
			ans = j - i
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
