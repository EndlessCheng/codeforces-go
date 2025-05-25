package main

// https://space.bilibili.com/206214
func minCuttingCost(n, m, k int) int64 {
	return int64(max(k*(n-k), 0) + max(k*(m-k), 0))
}
