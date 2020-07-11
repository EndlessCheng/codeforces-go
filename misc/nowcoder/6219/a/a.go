package main

// github.com/EndlessCheng/codeforces-go
func Orderofpoker(s string) (ans string) {
	for s != "" {
		n := len(s) / 2
		if n == 2 || n == 3 || n == 5 || n == 7 {
			ans += s[:2]
			s = s[2:]
		} else {
			n *= 2
			ans += s[n-2:]
			s = s[:n-2]
		}
	}
	return
}
