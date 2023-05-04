package main

import "strings"

// https://space.bilibili.com/206214
func adventureCamp(a []string) int {
	vis := map[string]bool{}
	for _, s := range strings.Split(a[0], "->") {
		vis[s] = true
	}
	maxCnt, ans := 0, -1
	for i := 1; i < len(a); i++ {
		if a[i] == "" {
			continue
		}
		cnt := 0
		for _, s := range strings.Split(a[i], "->") {
			if !vis[s] {
				vis[s] = true
				cnt++
			}
		}
		if cnt > maxCnt {
			maxCnt, ans = cnt, i
		}
	}
	return ans
}
