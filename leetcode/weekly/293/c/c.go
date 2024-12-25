package main

import (
	"math/bits"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func largestCombination(candidates []int) int {
	cnt := [24]int{}
	for _, x := range candidates {
		for s := uint(x); s > 0; s &= s - 1 {
			cnt[bits.TrailingZeros(s)]++
		}
	}
	return slices.Max(cnt[:])
}

func largestCombination2(candidates []int) (ans int) {
	m := bits.Len(uint(slices.Max(candidates)))
	for i := range m {
		cnt := 0
		for _, x := range candidates {
			cnt += x >> i & 1
		}
		ans = max(ans, cnt)
	}
	return
}
