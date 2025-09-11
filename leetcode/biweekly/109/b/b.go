package main

import (
	"slices"
	"strings"
	"unicode"
)

// https://space.bilibili.com/206214
func sortVowels1(s string) string {
	vowels := []byte{}
	for _, ch := range s {
		c := unicode.ToLower(ch)
		if strings.ContainsRune("aeiou", c) {
			vowels = append(vowels, byte(ch))
		}
	}

	slices.Sort(vowels)

	t := []byte(s)
	j := 0
	for i, ch := range t {
		c := unicode.ToLower(rune(ch))
		if strings.ContainsRune("aeiou", c) {
			t[i] = vowels[j]
			j++
		}
	}
	return string(t)
}

func sortVowels2(s string) string {
	const vowelMask = 0x208222
	vowels := []byte{}
	for _, ch := range s {
		if vowelMask>>(ch&31)&1 > 0 { // ch 是元音
			vowels = append(vowels, byte(ch))
		}
	}
	slices.Sort(vowels)

	t := []byte(s)
	j := 0
	for i, ch := range t {
		if vowelMask>>(ch&31)&1 > 0 { // ch 是元音
			t[i] = vowels[j]
			j++
		}
	}
	return string(t)
}

func sortVowels(s string) string {
	const vowelMask = 0x208222
	cnt := ['u' + 1]int{}
	for _, ch := range s {
		if vowelMask>>(ch&31)&1 > 0 {
			cnt[ch]++
		}
	}

	t := []byte(s)
	j := byte('A')
	for i, ch := range t {
		if vowelMask>>(ch&31)&1 == 0 {
			continue
		}
		// 找下一个出现次数大于 0 的元音字母
		for cnt[j] == 0 {
			if j == 'Z' {
				j = 'a'
			} else {
				j++
			}
		}
		t[i] = j
		cnt[j]--
	}
	return string(t)
}
