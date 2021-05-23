package main

// github.com/EndlessCheng/codeforces-go
func stoneGameVIII(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	f := sum
	for i := len(a) - 1; i > 1; i-- {
		sum -= a[i]
		f = max(f, sum-f)
	}
	return f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
