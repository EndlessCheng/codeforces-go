package main

// https://space.bilibili.com/206214
var m = map[int]int{1: 1, 8: 6, 49: 35, 288: 204}

func pivotInteger(n int) int {
	if ans, ok := m[n]; ok {
		return ans
	}
	return -1
}
