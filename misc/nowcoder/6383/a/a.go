package main

// github.com/EndlessCheng/codeforces-go
func search(n int, a []int) int {
	s := 0
	for _, v := range a {
		s ^= v
	}
	for i := 1; i <= n; i++ {
		s ^= i
	}
	return s
}
