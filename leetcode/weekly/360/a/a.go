package main

import "strings"

// https://space.bilibili.com/206214
func furthestDistanceFromOrigin(moves string) int {
	cntR := strings.Count(moves, "R")
	cntL := strings.Count(moves, "L")
	return abs(cntR-cntL) + len(moves) - cntR - cntL
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
