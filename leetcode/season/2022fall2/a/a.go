package main

// https://space.bilibili.com/206214
func minNumBooths(demand []string) (ans int) {
	maxCnt := [26]int{}
	for _, s := range demand {
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		for i, c := range cnt {
			maxCnt[i] = max(maxCnt[i], c)
		}
	}
	for _, c := range maxCnt {
		ans += c
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
