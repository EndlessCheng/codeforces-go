package main

import (
	"strconv"
)

// https://space.bilibili.com/206214
func isFascinating2(n int) bool {
	if n < 123 || n > 329 {
		return false
	}
	mask := 0
	for _, c := range strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3) {
		mask |= 1 << (c - '0')
	}
	return mask == 1<<10-2
}

func isFascinating(n int) bool {
	return n == 192 || n == 219 || n == 273 || n == 327
}
