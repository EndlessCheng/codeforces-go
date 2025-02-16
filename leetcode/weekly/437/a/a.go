package main

// https://space.bilibili.com/206214
func hasSpecialSubstring(s string, k int) bool {
	cnt := 0
	for i := range s {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt == k {
				return true
			}
			cnt = 0
		}
	}
	return false
}
