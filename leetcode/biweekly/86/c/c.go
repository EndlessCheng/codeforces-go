package main

import "math/bits"

// https://space.bilibili.com/206214
func maximumRows(mat [][]int, cols int) (ans int) {
	m, n := len(mat), len(mat[0])
	mask := make([]int, m)
	for i, row := range mat {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	for set := 1<<cols - 1; set < 1<<n; {
		cnt := 0
		for _, row := range mask {
			if row&set == row { // row 是 set 的子集，所有 1 都被覆盖
				cnt++
			}
		}
		ans = max(ans, cnt)
		lb := set & -set
		x := set + lb
		// 下式等价于 set = (set^x)/lb>>2 | x
		set = (set^x)>>bits.TrailingZeros(uint(lb))>>2 | x
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
