package main

// https://space.bilibili.com/206214
func findValidPair(s string) (ans string) {
	cnt := [10]int{}
	for _, b := range s {
		cnt[b-'0']++
	}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'0', s[i]-'0'
		if x != y && cnt[x] == int(x) && cnt[y] == int(y) {
			return s[i-1 : i+1]
		}
	}
	return
}
