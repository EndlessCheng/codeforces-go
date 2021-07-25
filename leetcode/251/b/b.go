package main

// Go 贪心

// github.com/EndlessCheng/codeforces-go
func maximumNumber(num string, change []int) string {
	s := []byte(num)
	i, n := 0, len(s)
	for ; i < n && byte(change[s[i]&15]) <= s[i]&15; i++ {
	}
	for ; i < n && byte(change[s[i]&15]) >= s[i]&15; i++ {
		s[i] = '0' + byte(change[s[i]&15])
	}
	return string(s)
}
