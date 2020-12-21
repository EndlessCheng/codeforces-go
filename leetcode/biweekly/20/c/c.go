package main

// github.com/EndlessCheng/codeforces-go
func numberOfSubstrings(s string) (ans int) {
	c := ['d']int{}
	l := 0
	for _, b := range s {
		c[b]++
		for c['a'] > 0 && c['b'] > 0 && c['c'] > 0 {
			c[s[l]]--
			l++
		}
		ans += l
	}
	return
}
