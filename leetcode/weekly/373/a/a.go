package main

import "slices"

// https://space.bilibili.com/206214
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k %= n
	if k == 0 {
		return true
	}
	for _, r := range mat {
		if !slices.Equal(r, append(r[k:], r[:k]...)) {
			return false
		}
	}
	return true
}
