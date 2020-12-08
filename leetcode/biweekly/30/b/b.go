package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func rangeSum(a []int, n, left, right int) (ans int) {
	b := make([]int, 0, n*(n+1)/2)
	for i, s := range a {
		b = append(b, s)
		for _, v := range a[i+1:] {
			s += v
			b = append(b, s)
		}
	}
	sort.Ints(b)
	for _, v := range b[left-1 : right] {
		ans += v
	}
	return ans % (1e9 + 7)
}
