package main

// https://space.bilibili.com/206214/dynamic
func digitCount(num string) bool {
	cnt := [10]int{}
	for _, ch := range num {
		cnt[ch&15]++
	}
	for i, ch := range num {
		if cnt[i] != int(ch&15) {
			return false
		}
	}
	return true
}
