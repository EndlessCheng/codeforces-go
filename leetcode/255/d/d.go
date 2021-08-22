package main

import (
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func recoverArray(n int, sums []int) []int {
	sort.Ints(sums)
	min := sums[0]
	if min < 0 {
		for i := range sums {
			sums[i] -= min
		}
	}

	skip := map[int]int{}
	ans := make([]int, 0, n)
	for j := 0; n > 0; n-- {
		for j++; skip[sums[j]] > 0; j++ {
			skip[sums[j]]--
		}
		s := sums[j]
		_s := make([]int, 1<<len(ans))
		for i, v := range ans {
			for m, b := 0, 1<<i; m < b; m++ {
				_s[b|m] = _s[m] + v
				skip[_s[b|m]+s]++
			}
		}
		ans = append(ans, s)
	}

	if min < 0 {
		_s := make([]int, 1<<len(ans))
		for i, v := range ans {
			b := 1 << i
			for j := 0; j < b; j++ {
				_s[b|j] = _s[j] + v
				if _s[b|j] == -min {
					for s := uint(b | j); s > 0; s &= s - 1 {
						p := bits.TrailingZeros(s)
						ans[p] = -ans[p]
					}
					return ans
				}
			}
		}
	}
	return ans
}
