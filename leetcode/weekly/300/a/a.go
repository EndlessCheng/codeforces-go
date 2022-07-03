package main

// https://space.bilibili.com/206214/dynamic
func decodeMessage(key, message string) string {
	mp := ['z' + 1]byte{' ': ' '} // 空格映射到空格上
	cur := byte('a')
	for _, c := range key {
		if mp[c] == 0 {
			mp[c] = cur
			cur++
		}
	}
	s := []byte(message)
	for i, c := range s {
		s[i] = mp[c]
	}
	return string(s)
}
