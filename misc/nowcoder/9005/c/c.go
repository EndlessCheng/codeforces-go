package main

// github.com/EndlessCheng/codeforces-go
const m int = 1e9 + 7

func cowModCount(N int64) int64 {
	n := int(N)
	ans := 0
	for l, r := 1, 0; l <= n; l = r + 1 {
		h := n / l
		r = n / h
		w := r - l + 1
		ans += h * w * (2*n - h*(l+r)) % m
	}
	return int64(ans % m * (m + 1) / 2 % m)
}
