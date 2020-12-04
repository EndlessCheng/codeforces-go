package main

// github.com/EndlessCheng/codeforces-go
func Sum(n int) int64 {
	s := n
	n--
	for l := 1; l <= n; {
		h := n / l
		r := n / h
		s += h * (r - l + 1)
		l = r + 1
	}
	return int64(s)
}
