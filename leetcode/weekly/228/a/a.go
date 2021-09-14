package main

// github.com/EndlessCheng/codeforces-go
func minOperations(s string) (ans int) {
	c := 0
	for i, b := range s {
		if int(b&1) != i&1 {
			c++
		}
	}
	return min(c, len(s)-c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
