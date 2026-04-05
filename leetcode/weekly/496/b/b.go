package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
var goodIntegers []int // 1554 个

func init() {
	const mx = 1_000_000_000
	cnt := map[int]int{}
	for a := 1; a*a*a <= mx/2; a++ {
		for b := a; a*a*a+b*b*b <= mx; b++ {
			cnt[a*a*a+b*b*b]++
		}
	}

	for x, c := range cnt {
		if c > 1 {
			goodIntegers = append(goodIntegers, x)
		}
	}

	slices.Sort(goodIntegers)
}

func findGoodIntegers(n int) []int {
	i := sort.SearchInts(goodIntegers, n+1)
	return goodIntegers[:i]
}
