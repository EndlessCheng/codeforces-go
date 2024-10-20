package main

// https://space.bilibili.com/206214
func stringSequence(target string) (ans []string) {
	s := make([]byte, len(target))
	for i, c := range target {
		for j := byte('a'); j <= byte(c); j++ {
			s[i] = j
			ans = append(ans, string(s[:i+1]))
		}
	}
	return
}
