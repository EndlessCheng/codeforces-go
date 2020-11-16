package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func longestSubarray(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + 1 - v
	}
	for i, v := range a {
		if v == 1 && (i+1 == n || a[i+1] == 0) {
			j := sort.Search(i+1, func(j int) bool { return sum[i+1]-sum[j] <= 1 })
			l := i - j
			// WA 了一发：漏了 [1,0] 这种情况！
			if sum[i+1]-sum[j] == 0 && i+1 < n {
				l++
			}
			if l > ans {
				ans = l
			}
		}
	}
	return
}
