package main

import "sort"

// https://space.bilibili.com/206214
func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	lcm := d1 / gcd(d1, d2) * d2
	return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
		left1 := max(uniqueCnt1-x/d2+x/lcm, 0)
		left2 := max(uniqueCnt2-x/d1+x/lcm, 0)
		common := x - x/d1 - x/d2 + x/lcm
		return common >= left1+left2
	})
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
