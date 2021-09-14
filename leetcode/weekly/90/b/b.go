package main

// github.com/EndlessCheng/codeforces-go
func scoreOfParentheses(s string) (ans int) {
	var f func(string) int
	f = func(s string) (res int) {
		c := 0
		for i, b := range s {
			if b == '(' {
				c++
			} else {
				c--
				if c == 0 {
					if i == 1 {
						return 1 + f(s[2:])
					}
					return 2*f(s[1:i]) + f(s[i+1:])
				}
			}
		}
		return 0
	}
	return f(s)
}
