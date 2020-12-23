package main

// github.com/EndlessCheng/codeforces-go
func countLetters(s string) (ans int) {
	for i, n := 0, len(s); i < n; {
		st := i
		v := s[st]
		for ; i < n && s[i] == v; i++ {
		}
		ans += (i - st + 1) * (i - st) / 2
	}
	return
}
