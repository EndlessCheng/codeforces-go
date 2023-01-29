package main

import "sort"

// https://space.bilibili.com/206214
func putMarbles(wt []int, k int) (ans int64) {
	for i, w := range wt[1:] {
		wt[i] += w
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	for _, w := range wt[len(wt)-k+1:] {
		ans += int64(w)
	}
	for _, w := range wt[:k-1] {
		ans -= int64(w)
	}
	return
}
