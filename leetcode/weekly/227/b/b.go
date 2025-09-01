package main

// github.com/EndlessCheng/codeforces-go
func maximumScore(a, b, c int) int {
	s := a + b + c
	m := max(a, b, c)
	return min(s/2, s-m)
}
