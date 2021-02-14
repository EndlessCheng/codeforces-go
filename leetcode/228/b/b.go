package main

// github.com/EndlessCheng/codeforces-go
func countHomogenous(s string) (ans int) {
	for i, n := 0, len(s); i < n; {
		st := i
		v := s[st]
		for ; i < n && s[i] == v; i++ {
		}
		sz := i - st
		ans += (sz + 1) * sz / 2
	}
	return ans % (1e9 + 7)
}
