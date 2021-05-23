package main

// github.com/EndlessCheng/codeforces-go
func stoneGameVIII(a []int) int {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	f0, f1 := sum[n], sum[n]
	for i := n - 1; i > 1; i-- {
		f0, f1 = max(f0, sum[i]-f1), max(f1, sum[i]-f0)
	}
	return f0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
