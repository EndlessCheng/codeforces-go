package main

// github.com/EndlessCheng/codeforces-go
func replaceElements(a []int) []int {
	mx := -1
	for i := len(a) - 1; i >= 0; i-- {
		mx, a[i] = max(mx, a[i]), mx
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
