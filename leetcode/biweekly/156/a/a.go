package main

import "slices"

// https://space.bilibili.com/206214
func maxFreqSum(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	maxVowel := 0
	for _, b := range "aeiou" {
		maxVowel = max(maxVowel, cnt[b-'a'])
		cnt[b-'a'] = 0 // 这样下面计算的一定是辅音出现次数的最大值
	}

	maxConsonant := slices.Max(cnt[:])
	return maxVowel + maxConsonant
}
