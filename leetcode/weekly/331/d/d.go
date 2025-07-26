package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	a := []int{}
	mn := math.MaxInt
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		for range abs(c) / 2 {
			a = append(a, x)
		}
	}

	slices.Sort(a)

	for _, x := range a[:len(a)/2] {
		ans += int64(min(x, mn*2))
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
