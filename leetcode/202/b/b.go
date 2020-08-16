package main

// github.com/EndlessCheng/codeforces-go
func minOperations(n int) (ans int) {
	for i := 1; i < n; i += 2 {
		ans += n - i
	}
	return
}
