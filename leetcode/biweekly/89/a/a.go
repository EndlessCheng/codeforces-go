package main

import "fmt"

// https://space.bilibili.com/206214
func count(time string, limit int) (ans int) {
next:
	for i := 0; i < limit; i++ {
		for j, c := range fmt.Sprintf("%02d", i) {
			if time[j] != '?' && byte(c) != time[j] {
				continue next
			}
		}
		ans++
	}
	return
}

func countTime(time string) int {
	return count(time[:2], 24) * count(time[3:], 60)
}
