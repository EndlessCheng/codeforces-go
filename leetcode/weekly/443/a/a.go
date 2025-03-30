package main

// https://space.bilibili.com/206214
func minCosts(cost []int) []int {
	for i := 1; i < len(cost); i++ {
		cost[i] = min(cost[i], cost[i-1])
	}
	return cost
}
