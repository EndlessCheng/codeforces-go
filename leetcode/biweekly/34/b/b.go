package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func numWays(s string) (ans int) {
	const mod int = 1e9 + 7
	c1 := strings.Count(s, "1")
	if c1%3 > 0 {
		return
	}
	n := len(s)
	if c1 == 0 {
		return (n - 1) * (n - 2) / 2 % mod
	}
	l, r := 0, n-1
	for c := 0; 3*c < c1; l++ {
		c += int(s[l] & 15)
	}
	st := l
	for ; s[l] == '0'; l++ {
	}
	ans = l - st + 1
	for c := 0; 3*c < c1; r-- {
		c += int(s[r] & 15)
	}
	st = r
	for ; s[r] == '0'; r-- {
	}
	ans *= st - r + 1
	ans %= mod
	return
}
