package main

func getLastMoment(n int, left []int, right []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, p := range left {
		ans = max(ans, p)
	}
	for _, p := range right {
		ans = max(ans, n-p)
	}
	return
}
