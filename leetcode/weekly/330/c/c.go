package main

import "slices"

// https://space.bilibili.com/206214
func putMarbles(weights []int, k int) (ans int64) {
	for i, w := range weights[1:] {
		weights[i] += w
	}
	weights = weights[:len(weights)-1]
	slices.Sort(weights)
	for _, w := range weights[len(weights)-k+1:] {
		ans += int64(w)
	}
	for _, w := range weights[:k-1] {
		ans -= int64(w)
	}
	return
}
