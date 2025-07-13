package main

import "slices"

// https://space.bilibili.com/206214
func processStr(s string) string {
	ans := []byte{}
	for _, c := range s {
		if c == '*' {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else if c == '#' {
			ans = append(ans, ans...)
		} else if c == '%' {
			slices.Reverse(ans)
		} else {
			ans = append(ans, byte(c))
		}
	}
	return string(ans)
}
