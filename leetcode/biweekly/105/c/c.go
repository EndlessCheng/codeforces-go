package main

import "math/bits"

// https://space.bilibili.com/206214
func maxStrength(a []int) int64 {
	ans := int(-1e18)

	calc := func(sub int) (res int) {
		res = 1
		for _s := uint(sub); _s > 0; _s &= _s - 1 {
			p1 := bits.TrailingZeros(_s)
			v := a[p1]
			res *= v
		}
		return
	}
	for sub := 1; sub < 1<<len(a); sub++ {
		res := calc(sub)
		ans = max(ans, res)
	}

	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
