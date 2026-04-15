package main

// https://space.bilibili.com/206214
func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	for k := range n/2 + 1 {
		if words[(startIndex-k+n)%n] == target || words[(startIndex+k)%n] == target {
			return k
		}
	}
	return -1
}
