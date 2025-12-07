package main

import "slices"

// https://space.bilibili.com/206214
func maxPoints(a, b []int, k int) (ans int64) {
	n := len(a)
	d := a[:0]
	for i, x := range a {
		ans += int64(x)
		v := b[i] - x
		if v > 0 {
			d = append(d, v)
		}
	}

	slices.SortFunc(d, func(a, b int) int { return b - a })
	for _, x := range d[:min(n-k, len(d))] {
		ans += int64(x)
	}
	return
}
