package main

// github.com/EndlessCheng/codeforces-go
func solve(s string) int {
	n := len(s)
	f := make([]int, n)
	for i, c := 1, 0; i < n; i++ {
		b := s[i]
		for c > 0 && s[c] != b {
			c = f[c-1]
		}
		if s[c] == b {
			c++
		}
		f[i] = c
	}
	ans := f[n-1]
	if ans == 0 {
		ans--
	}
	return ans
}
