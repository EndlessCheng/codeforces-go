package main

// https://space.bilibili.com/206214
func maximumNumberOfStringPairs(words []string) (ans int) {
	vis := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if vis[y][x] {
			ans++
		} else {
			vis[x][y] = true
		}
	}
	return
}
