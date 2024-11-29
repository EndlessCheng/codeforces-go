package main

import (
	"maps"
	"slices"
)

func minSetSize(arr []int) int {
	cnt := make([]int, slices.Max(arr)+1)
	for _, x := range arr {
		cnt[x]++
	}
	slices.SortFunc(cnt, func(a, b int) int { return b - a })

	s := 0
	for i, c := range cnt {
		s += c
		if s >= len(arr)/2 {
			return i + 1
		}
	}
	panic("impossible")
}

func minSetSize2(arr []int) int {
	freq := map[int]int{}
	for _, x := range arr {
		freq[x]++
	}

	cnt := slices.SortedFunc(maps.Values(freq), func(a, b int) int { return b - a })

	s := 0
	for i, c := range cnt {
		s += c
		if s >= len(arr)/2 {
			return i + 1
		}
	}
	panic("impossible")
}
