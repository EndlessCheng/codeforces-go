package main

import "slices"

// https://space.bilibili.com/206214
func minDamage(power int, damage, health []int) (ans int64) {
	type pair struct{ k, d int }
	a := make([]pair, len(health))
	for i, h := range health {
		a[i] = pair{(h-1)/power + 1, damage[i]}
	}
	slices.SortFunc(a, func(p, q pair) int { return p.k*q.d - q.k*p.d })

	s := 0
	for _, p := range a {
		s += p.k
		ans += int64(s) * int64(p.d)
	}
	return
}
