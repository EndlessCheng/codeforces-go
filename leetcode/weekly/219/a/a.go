package main

// github.com/EndlessCheng/codeforces-go
func numberOfMatches(n int) (ans int) {
	for ; n > 1; n = (n + 1) / 2 {
		ans += n / 2
	}
	return
}
