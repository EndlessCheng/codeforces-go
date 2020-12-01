package main

// github.com/EndlessCheng/codeforces-go
func Maximumlength(s string) (ans int) {
	c := ['z' + 1]int{}
	l := 0
	for r := range s {
		c[s[r]]++
		for c['n']*c['p']*c['y'] > 0 {
			c[s[l]]--
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return
}
