package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minBitFlips(start, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}
