package main

import "slices"

// https://space.bilibili.com/206214
func smallestPalindrome(s string) string {
	n := len(s)
	t := []byte(s[:n/2])
	slices.Sort(t)

	ans := string(t)
	if n%2 > 0 {
		ans += string(s[n/2])
	}
	slices.Reverse(t)
	return ans + string(t)
}
