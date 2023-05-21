package main

// https://space.bilibili.com/206214
func makeSmallestPalindrome(S string) string {
	s := []byte(S)
	for i, n := 0, len(s); i < n/2; i++ {
		x, y := s[i], s[n-1-i]
		if x > y {
			s[i] = y
		} else {
			s[n-1-i] = x
		}
	}
	return string(s)
}
