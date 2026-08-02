package main

import "math"

// https://space.bilibili.com/206214
func stoneGameII(piles []int) int {
	n := len(piles)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i, s := n-1, 0; i >= 0; i-- {
		s += piles[i]
		for m := 1; m <= i/2+1; m++ {
			if i+m*2 >= n { // 全拿
				f[i][m] = s
			} else {
				mn := math.MaxInt
				for x := 1; x <= m*2; x++ {
					mn = min(mn, f[i+x][max(m, x)])
				}
				f[i][m] = s - mn
			}
		}
	}
	return f[0][1]
}
