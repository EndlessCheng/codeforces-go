package main

import "sort"

// https://space.bilibili.com/206214
func dividePlayers(skill []int) (ans int64) {
	sort.Ints(skill)
	n := len(skill)
	sum := skill[0] + skill[n-1]
	for i := 0; i < n/2; i++ {
		x, y := skill[i], skill[n-1-i]
		if x+y != sum {
			return -1
		}
		ans += int64(x * y)
	}
	return
}
