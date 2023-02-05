package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	mn := math.MaxInt
	var a, b []int
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		if c > 0 {
			for i := 0; i < c/2; i++ {
				a = append(a, x)
			}
		} else {
			for i := 0; i < -c/2; i++ {
				b = append(b, x)
			}
		}
	}
	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	for i, x := range a {
		ans += int64(min(min(x, b[i]), mn*2))
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
