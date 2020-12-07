package main

// github.com/EndlessCheng/codeforces-go
func minInsertions(s string) (ans int) {
	c, n := 0, len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			c++
		} else {
			if i+1 == n || s[i+1] != ')' {
				ans++
			} else {
				i++
			}
			if c > 0 {
				c--
			} else {
				ans++
			}
		}
	}
	ans += 2 * c
	return
}
