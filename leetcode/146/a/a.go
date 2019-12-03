package main

import "sort"

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := [81]int{}
	for _, d := range dominoes {
		sort.Ints(d)
		cnt[9*(d[0]-1)+d[1]-1]++
	}
	ans := 0
	for _, c := range cnt {
		if c > 1 {
			ans += c * (c - 1) / 2
		}
	}
	return ans
}
