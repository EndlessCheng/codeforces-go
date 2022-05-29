package main

import "strings"

// https://space.bilibili.com/206214/dynamic
func largestWordCount(messages, senders []string) (ans string) {
	cnt := map[string]int{}
	for i, msg := range messages {
		cnt[senders[i]] += strings.Count(msg, " ") + 1
	}
	for s, c := range cnt {
		if c > cnt[ans] || c == cnt[ans] && s > ans {
			ans = s
		}
	}
	return
}
