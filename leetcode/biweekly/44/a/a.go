package main

// github.com/EndlessCheng/codeforces-go
func largestAltitude(a []int) (ans int) {
	sum := 0
	for _, v := range a {
		sum += v
		ans = max(ans, sum)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
