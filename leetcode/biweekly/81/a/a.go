package main

import "strings"

// https://space.bilibili.com/206214/dynamic
func countAsterisks(s string) (ans int) {
	sp := strings.Split(s, "|")
	for i := 0; i < len(sp); i += 2 {
		ans += strings.Count(sp[i], "*")
	}
	return
}
