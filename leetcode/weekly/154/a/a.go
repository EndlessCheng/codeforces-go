package main

// https://space.bilibili.com/206214
func maxNumberOfBalloons(text string) int {
	cnt := ['z' + 1]int{}
	for _, ch := range text {
		cnt[ch]++
	}
	return min(cnt['a'], cnt['b'], cnt['l']/2, cnt['n'], cnt['o']/2)
}
