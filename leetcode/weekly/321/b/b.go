package main

// https://space.bilibili.com/206214
func appendCharacters(s, t string) int {
	i, n := 0, len(s)
	for j := range t {
		for i < n && s[i] != t[j] {
			i++
		}
		if i == n {
			return len(t) - j
		}
		i++
	}
	return 0
}
