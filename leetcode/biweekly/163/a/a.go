package main

// https://space.bilibili.com/206214
func minSensors(n, m, k int) int {
	size := k*2 + 1
	return ((n-1)/size + 1) * ((m-1)/size + 1)
}
