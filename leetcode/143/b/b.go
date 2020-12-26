package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	return append(pathInZigZagTree(1<<(bits.Len(uint(label))-2)*3-1-label/2), label)
}
