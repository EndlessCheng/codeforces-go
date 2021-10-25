package main

import (
	"strings"
	"unicode"
)

// Go 模拟，附详细注释

// github.com/EndlessCheng/codeforces-go
func countValidWords(sentence string) (ans int) {
next:
	for _, s := range strings.Fields(sentence) { // 按照空格分割
		if strings.ContainsAny(s, "0123456789") { // 不能包含数字
			continue
		}
		i := strings.IndexByte(s, '-')
		if i >= 0 && (strings.Contains(s[i+1:], "-") || // 应至多包含一个连字符
			i == 0 || i == len(s)-1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1]))) { // 连字符左右应均为小写字母
			continue
		}
		cnt := 0
		for _, ch := range "!.," {
			if i := strings.IndexRune(s, ch); i >= 0 {
				cnt++
				if cnt > 1 || i != len(s)-1 { // 应包含至多一个标点符号且标点符号应当位于末尾
					continue next
				}
			}
		}
		ans++ // 单词是有效的
	}
	return
}
