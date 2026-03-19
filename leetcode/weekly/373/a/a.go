package main

import "slices"

// https://space.bilibili.com/206214
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	for _, row := range mat {
		for j, x := range row {
			if x != row[(j+k)%n] {
				return false
			}
		}
	}
	return true
}

func areSimilar2(mat [][]int, k int) bool {
	k %= len(mat[0])
	for _, row := range mat {
		if !slices.Equal(row, append(row[k:], row[:k]...)) {
			return false
		}
	}
	return true
}
