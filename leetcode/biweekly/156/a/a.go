package main

// https://space.bilibili.com/206214
func maxFreqSum1(s string) int {
	cnt := [26]int{}
	maxVowelCnt := 0
	maxConsonantCnt := 0
	for _, ch := range s {
		cnt[ch-'a']++
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			maxVowelCnt = max(maxVowelCnt, cnt[ch-'a'])
		} else {
			maxConsonantCnt = max(maxConsonantCnt, cnt[ch-'a'])
		}
	}
	return maxVowelCnt + maxConsonantCnt
}

func maxFreqSum(s string) int {
	const vowelMask = 0x104111
	cnt := [26]int{}
	maxCnt := [2]int{}
	for _, ch := range s {
		ch -= 'a'
		bit := vowelMask >> ch & 1
		cnt[ch]++
		maxCnt[bit] = max(maxCnt[bit], cnt[ch])
	}
	return maxCnt[0] + maxCnt[1]
}
