package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func minimumDeletions(word string, k int) int {
	cnt := make([]int, 26)
	for _, b := range word {
		cnt[b-'a']++
	}
	slices.Sort(cnt)

	maxSave := 0
	for i, base := range cnt {
		sum := 0
		for _, c := range cnt[i:] {
			sum += min(c, base+k)
		}
		maxSave = max(maxSave, sum)
	}
	return len(word) - maxSave
}
