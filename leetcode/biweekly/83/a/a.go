package main

import "bytes"

// https://space.bilibili.com/206214/dynamic
func bestHand(ranks []int, suits []byte) string {
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	cnt, pair := map[int]int{}, false
	for _, r := range ranks {
		cnt[r]++
		if cnt[r] == 3 {
			return "Three of a Kind"
		}
		if cnt[r] == 2 {
			pair = true
		}
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}
