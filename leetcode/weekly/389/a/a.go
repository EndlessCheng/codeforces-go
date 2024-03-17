package main

// https://space.bilibili.com/206214
func isSubstringPresent(s string) bool {
	vis := [26]int{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x] |= 1 << y
		if vis[y]>>x&1 > 0 {
			return true
		}
	}
	return false
}
