package main

import (
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func minimumTimeRequired(a []int, k int) int {
	m := 1 << len(a)
	sum := make([]int, m)
	for i := range sum {
		for s := uint(i); s > 0; s &= s - 1 {
			sum[i] += a[bits.TrailingZeros(s)]
		}
	}
	return sort.Search(12e7, func(mx int) bool {
		dp := make([][]int8, m)
		for i := range dp {
			dp[i] = make([]int8, k+1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int8) int8
		f = func(set int, left int8) (res int8) {
			if set == 0 {
				return 1
			}
			if left == 0 {
				return
			}
			dv := &dp[set][left]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
			for sub, ok := set, true; ok; ok = sub != set {
				if sum[sub] <= mx && f(set^sub, left-1) == 1 {
					return 1
				}
				sub = (sub - 1) & set
			}
			return
		}
		return f(m-1, int8(k)) == 1
	})
}
