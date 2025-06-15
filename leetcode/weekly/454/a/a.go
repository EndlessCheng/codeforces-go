package main

import (
	"strings"
	"unicode"
)

// https://space.bilibili.com/206214
func generateTag(caption string) string {
	ans := []byte{'#'}
	for i, s := range strings.Fields(caption) {
		s = strings.ToLower(s)
		if i > 0 { // 不是第一个单词，首字母大写
			s = strings.Title(s)
		}
		ans = append(ans, s...)
		if len(ans) >= 100 {
			ans = ans[:100]
			break
		}
	}
	return string(ans)
}

func generateTag2(caption string) string {
	s := strings.ToLower(caption)
	s = strings.Title(s) // 所有单词首字母大写
	s = strings.ReplaceAll(s, " ", "")
	if s == "" {
		return "#"
	}
	s = "#" + string(unicode.ToLower(rune(s[0]))) + s[1:]
	if len(s) >= 100 {
		s = s[:100]
	}
	return s
}
