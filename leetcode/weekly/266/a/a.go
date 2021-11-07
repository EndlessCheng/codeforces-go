package main

import "strings"

// 双指针 O(n)

// github.com/EndlessCheng/codeforces-go
func countVowelSubstrings(word string) (ans int) {
	for _, s := range strings.FieldsFunc(word, func(r rune) bool { return !strings.ContainsRune("aeiou", r) }) { // 分割出仅包含元音的字符串
		cnt := ['v']int{}
		l := 0
		for _, ch := range s {
			cnt[ch]++
			for cnt[s[l]] > 1 { // 双指针，仅当该元音个数不止一个时才移动左指针
				cnt[s[l]]--
				l++
			}
			if cnt['a'] > 0 && cnt['e'] > 0 && cnt['i'] > 0 && cnt['o'] > 0 && cnt['u'] > 0 { // 必须包含全部五种元音
				ans += l + 1
			}
		}
	}
	return
}
