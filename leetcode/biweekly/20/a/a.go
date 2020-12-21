package main

import (
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func sortByBits(a []int) (ans []int) {
	sort.Ints(a)
	sort.SliceStable(a, func(i, j int) bool { return bits.OnesCount(uint(a[i])) < bits.OnesCount(uint(a[j])) })
	return a
}
