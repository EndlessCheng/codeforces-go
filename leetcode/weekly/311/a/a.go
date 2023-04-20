package main

// https://space.bilibili.com/206214
func smallestEvenMultiple(n int) int {
	return n << (n & 1)
}
