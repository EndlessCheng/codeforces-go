package main

import "math"

// github.com/EndlessCheng/codeforces-go
func maxMatrixSum(matrix [][]int) int64 {
	total, negCnt, mn := 0, 0, math.MaxInt
	for _, row := range matrix {
		for _, x := range row {
			if x < 0 {
				negCnt++
				x = -x // 先把负数都变成正数
			}
			mn = min(mn, x)
			total += x
		}
	}

	if negCnt%2 > 0 { // 必须有一个负数
		total -= mn * 2 // 最小的数添加负号
	}
	return int64(total)
}
