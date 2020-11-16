package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func kthFactor(n, k int) (ans int) {
	ds := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ds = append(ds, i)
			if i*i < n {
				ds = append(ds, n/i)
			}
		}
	}
	sort.Ints(ds)
	if len(ds) < k {
		return -1
	}
	return ds[k-1]
}
