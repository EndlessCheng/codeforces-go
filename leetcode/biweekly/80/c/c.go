package main

// https://space.bilibili.com/206214/dynamic
func matchReplacement(s, sub string, mappings [][]byte) bool {
	mp := ['z' + 1]['z' + 1]bool{}
	for _, p := range mappings {
		mp[p[0]][p[1]] = true
	}
next:
	for i := len(sub); i <= len(s); i++ {
		for j, c := range s[i-len(sub) : i] {
			if byte(c) != sub[j] && !mp[sub[j]][c] {
				continue next
			}
		}
		return true
	}
	return false
}
