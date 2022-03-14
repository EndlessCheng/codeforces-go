package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func leastMinutes(n int) int {
	return bits.Len(uint(n-1)) + 1
}
