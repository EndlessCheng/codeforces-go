package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func waysToSplit(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	sn := sum[n]
	for r := 2; r < n && 3*sum[r] <= 2*sn; r++ {
		l1 := sort.SearchInts(sum[1:r], 2*sum[r]-sn) + 1
		l2 := sort.Search(r, func(l int) bool { return 2*sum[l] > sum[r] })
		ans += l2 - l1
	}
	return ans % (1e9 + 7)
}
