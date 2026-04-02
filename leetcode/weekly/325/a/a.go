package main

// https://space.bilibili.com/206214
func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	ans := n
	for i, word := range words {
		if word == target {
			d := abs(i - startIndex)
			ans = min(ans, d, n-d)
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
