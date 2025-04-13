package main

import (
	"bytes"
	"slices"
)

// https://space.bilibili.com/206214
func smallestPalindrome(s string) string {
	n := len(s)
	cnt := [26]int{}
	for _, b := range s[:n/2] {
		cnt[b-'a']++
	}

	ans := make([]byte, 0, n) // 预分配空间
	for i, c := range cnt {
		ans = append(ans, bytes.Repeat([]byte{'a' + byte(i)}, c)...)
	}
	t := slices.Clone(ans)
	if n%2 > 0 {
		ans = append(ans, s[n/2])
	}
	slices.Reverse(t)
	return string(append(ans, t...))
}

func smallestPalindrome1(s string) string {
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
