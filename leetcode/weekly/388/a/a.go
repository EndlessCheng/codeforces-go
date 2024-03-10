package main

import "slices"

// https://space.bilibili.com/206214
func minimumBoxes(apple, capacity []int) int {
	s := 0
	for _, x := range apple {
		s += x
	}
	slices.SortFunc(capacity, func(a, b int) int { return b - a })
	for i, c := range capacity {
		s -= c
		if s <= 0 {
			return i + 1
		}
	}
	return -1
}
