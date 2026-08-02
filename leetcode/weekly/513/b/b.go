package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) add(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func countRatioSubarrays(nums []int, a, b int) (ans int) {
	sum := make([]int, len(nums)+1)
	val := [2]int{-b, a}
	for i, x := range nums {
		sum[i+1] = sum[i] + val[x&1]
	}

	sorted := slices.Clone(sum)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	t := make(fenwick, len(sorted)+1)
	for _, s := range sum {
		s = sort.SearchInts(sorted, s) + 1
		ans += t.pre(s)
		t.add(s)
	}
	return
}
