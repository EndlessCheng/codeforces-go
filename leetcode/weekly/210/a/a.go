package main

// github.com/EndlessCheng/codeforces-go
func maxDepth(s string) (ans int) {
	c := 0
	for _, b := range s {
		if b == '(' {
			c++
			if c > ans {
				ans = c
			}
		} else if b == ')' {
			c--
		}
	}
	return
}
