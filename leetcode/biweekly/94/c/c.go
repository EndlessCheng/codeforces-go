package main

import "sort"

// https://space.bilibili.com/206214
func minimizeSet1(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	l := lcm(d1, d2)
	return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
		left1 := max(uniqueCnt1-x/d2+x/l, 0)
		left2 := max(uniqueCnt2-x/d1+x/l, 0)
		common := x - x/d1 - x/d2 + x/l
		return common >= left1+left2
	})
}

func f(cnt, d int) int {
	return cnt + (cnt-1)/(d-1)
}

func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	l := lcm(d1, d2)
	return max(f(uniqueCnt1, d1), f(uniqueCnt2, d2), f(uniqueCnt1+uniqueCnt2, l))
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
