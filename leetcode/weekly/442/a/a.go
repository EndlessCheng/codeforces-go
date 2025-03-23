package main

// https://space.bilibili.com/206214
func maxContainers(n, w, maxWeight int) int {
	return min(maxWeight/w, n*n)
}
