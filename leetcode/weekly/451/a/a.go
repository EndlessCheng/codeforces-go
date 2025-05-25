package main

// https://space.bilibili.com/206214
func minCuttingCost(n, m, k int) int64 {
	return int64(max(k*(max(n, m)-k), 0))
}
