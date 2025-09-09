package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func canBeTypedWords1(text, brokenLetters string) (ans int) {
	for _, word := range strings.Split(text, " ") {
		if !strings.ContainsAny(word, brokenLetters) {
			ans++
		}
	}
	return
}

func canBeTypedWords(text, brokenLetters string) (ans int) {
	brokenMask := 0
	for _, c := range brokenLetters {
		brokenMask |= 1 << (c - 'a')
	}

	ok := 1
	for _, c := range text {
		if c == ' ' { // 上一个单词遍历完毕
			ans += ok
			ok = 1
		} else if brokenMask>>(c-'a')&1 > 0 { // c 在 brokenLetters 中
			ok = 0
		}
	}
	ans += ok // 最后一个单词
	return
}
