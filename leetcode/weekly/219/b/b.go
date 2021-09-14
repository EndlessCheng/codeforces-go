package main

// github.com/EndlessCheng/codeforces-go
func minPartitions(n string) (ans int) {
	for _, b := range n {
		ans = max(ans, int(b&15))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
