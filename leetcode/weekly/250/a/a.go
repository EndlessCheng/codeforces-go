package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func canBeTypedWords(text, brokenLetters string) (ans int) {
	for _, word := range strings.Split(text, " ") {
		if !strings.ContainsAny(word, brokenLetters) {
			ans++
		}
	}
	return
}
