package main

// https://space.bilibili.com/206214
func removeOuterParentheses(s string) string {
	ans := []byte{}
	depth := 0
	for _, ch := range s {
		if ch == '(' {
			if depth > 0 {
				ans = append(ans, byte(ch))
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				ans = append(ans, byte(ch))
			}
		}
	}
	return string(ans)
}
