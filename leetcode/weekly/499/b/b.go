package main

import (
	"slices"
)

// https://space.bilibili.com/206214
var mp = ['z' + 1]int{'a': 1, 'e': 2, 'i': 3, 'o': 4, 'u': 5}

func sortVowels(s string) string {
	cnt := [5]int{}
	vowels := []byte{} // 长度至多为 5
	for _, ch := range s {
		x := mp[ch] - 1
		if x < 0 {
			continue
		}
		if cnt[x] == 0 {
			vowels = append(vowels, byte(ch))
		}
		cnt[x]++
	}

	// 把 aeiou 按照出现次数从大到小排序
	slices.SortStableFunc(vowels, func(a, b byte) int { return cnt[mp[b]-1] - cnt[mp[a]-1] })

	t := []byte(s)
	j := 0
	for i, ch := range t {
		if mp[ch] == 0 {
			continue
		}
		t[i] = vowels[j]
		x := mp[t[i]] - 1
		cnt[x]--
		if cnt[x] == 0 {
			j++
		}
	}
	return string(t)
}
