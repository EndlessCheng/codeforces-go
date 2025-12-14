package main

import (
	"slices"
	"strings"
)

// https://space.bilibili.com/206214
func countVowel(s string) (vowel int) {
	for _, c := range s {
		if strings.IndexRune("aeiou", c) >= 0 {
			vowel++
		}
	}
	return
}

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	cnt0 := countVowel(a[0])
	for i := 1; i < len(a); i++ {
		if countVowel(a[i]) == cnt0 {
			t := []byte(a[i])
			slices.Reverse(t)
			a[i] = string(t)
		}
	}
	return strings.Join(a, " ")
}
