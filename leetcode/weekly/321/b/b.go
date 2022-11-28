package main

// https://space.bilibili.com/206214
func appendCharacters(s, t string) int {
	j, m := 0, len(t)
	for _, c := range s {
		if byte(c) == t[j] { // s 的字符肯定匹配的是 t 的前缀
			j++
			if j == m {
				return 0
			}
		}
	}
	return m - j
}
