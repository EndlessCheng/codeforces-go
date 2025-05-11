package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func minDeletion(s string, k int) (ans int) {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	slices.Sort(cnt[:])
	for _, c := range cnt[:26-k] {
		ans += c
	}
	return
}
