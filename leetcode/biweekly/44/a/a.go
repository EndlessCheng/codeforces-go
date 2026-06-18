package main

// github.com/EndlessCheng/codeforces-go
func largestAltitude(gain []int) (ans int) {
	h := 0
	for _, d := range gain {
		h += d
		ans = max(ans, h)
	}
	return
}
