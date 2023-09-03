package main

// https://space.bilibili.com/206214
func canBeEqual(s1, s2 string) bool {
	var cnt1, cnt2 [2][26]int
	for i, c := range s1 {
		cnt1[i%2][c-'a']++
		cnt2[i%2][s2[i]-'a']++
	}
	return cnt1 == cnt2
}
