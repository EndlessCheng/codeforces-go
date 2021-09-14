package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func f(st, sz int) int {
	if sz == 0 {
		return 0
	}
	if sz <= st {
		return (2*st + 1 - sz) * sz / 2
	}
	return (st+1)*st/2 + sz - st
}

func maxValue(n, i, maxSum int) (ans int) {
	return sort.Search(maxSum+1, func(v int) bool { return f(v-1, i)+v+f(v-1, n-1-i) > maxSum }) - 1
}
