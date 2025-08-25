package main

import "cmp"

// https://space.bilibili.com/206214
func findClosest(x, y, z int) int {
	a := abs(x - z)
	b := abs(y - z)
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return 2
}

var state = [3]int{1, 0, 2}

func findClosest2(x, y, z int) int {
	a := abs(x - z)
	b := abs(y - z)
	return state[cmp.Compare(a, b)+1]
}

func abs(x int) int { if x < 0 { return -x }; return x }
