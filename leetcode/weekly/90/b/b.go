package main

// github.com/EndlessCheng/codeforces-go
func scoreOfParentheses(s string) (ans int) {
	depth := 0
	for i, ch := range s {
		if ch == '(' {
			depth++
		} else {
			depth--
			if s[i-1] == '(' {
				ans += 1 << depth
			}
		}
	}
	return
}
