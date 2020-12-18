package main

// github.com/EndlessCheng/codeforces-go
func Sum(n int) int64 {
	if n == 1 {
		return 1
	}
	return int64(n) + 2*Sum(n/2)
}
