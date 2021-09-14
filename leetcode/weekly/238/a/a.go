package main

// github.com/EndlessCheng/codeforces-go
func sumBase(n, k int) (ans int) {
	for ; n > 0; n /= k {
		ans += n % k
	}
	return
}
