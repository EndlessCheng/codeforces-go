package main

// github.com/EndlessCheng/codeforces-go
func sumOfThree(n int64) []int64 {
	if n%3 == 0 {
		n /= 3
		return []int64{n - 1, n, n + 1}
	}
	return nil
}
