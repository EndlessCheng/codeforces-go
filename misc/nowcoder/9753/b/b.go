package main

// github.com/EndlessCheng/codeforces-go
func Gameresults(n, p, q int) int {
	if p > q || p < q && n <= p || p == q && n%(p+1) > 0 {
		return 1
	}
	return -1
}
