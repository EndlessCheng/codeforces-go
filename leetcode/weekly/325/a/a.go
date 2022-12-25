package main

// https://space.bilibili.com/206214
func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	ans := n
	for i, s := range words {
		if s == target {
			ans = min(ans, min(abs(i-startIndex), n-abs(i-startIndex)))
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
