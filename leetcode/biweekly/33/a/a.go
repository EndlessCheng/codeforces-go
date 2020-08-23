package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func thousandSeparator(n int) string {
	s := []byte(strconv.Itoa(n))
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	ans := []byte{}
	for i, b := range s {
		if i > 0 && i%3 == 0 {
			ans = append(ans, '.')
		}
		ans = append(ans, b)
	}
	for i, j := 0, len(ans)-1; i < j; i++ {
		ans[i], ans[j] = ans[j], ans[i]
		j--
	}
	return string(ans)
}
