package main

import "sort"

// github.com/EndlessCheng/codeforces-go
var ds = []int{0, 1, 6, 8, 9}
var mp = [...]int{1: 1, 6: 9, 8: 8, 9: 6}
var a []int

func f(cur, rev, p10 int) {
	if cur > 1e9 {
		return
	}
	if cur != rev {
		a = append(a, cur)
	}
	for _, d := range ds {
		if cur == 0 && d == 0 {
			continue
		}
		f(cur*10+d, mp[d]*p10+rev, p10*10)
	}
}

func init() {
	f(0, 0, 1)
	sort.Ints(a)
}

func confusingNumberII(n int) int {
	return sort.SearchInts(a, n+1)
}
