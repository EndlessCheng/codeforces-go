package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func largestSubmatrix1(matrix [][]int) (ans int) {
	n := len(matrix[0])
	heights := make([]int, n)

	for _, row := range matrix {
		for j, x := range row {
			if x == 0 {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}

		hs := slices.Clone(heights)
		slices.Sort(hs)
		for i, h := range hs { // 把 hs[i:] 作为子数组
			ans = max(ans, (n-i)*h) // 子数组长为 n-i，最小值为 h，对应的子矩形面积为 (n-i)*h
		}
	}

	return
}

func largestSubmatrix(matrix [][]int) (ans int) {
	n := len(matrix[0])
	heights := make([]int, n)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	_nonZeros := make([]int, n) // 避免在循环内反复申请内存

	for _, row := range matrix {
		zeros := idx[:0]
		nonZeros := _nonZeros[:0]
		for _, j := range idx {
			if row[j] == 0 {
				heights[j] = 0
				zeros = append(zeros, j)
			} else {
				heights[j]++
				nonZeros = append(nonZeros, j)
			}
		}
		idx = append(zeros, nonZeros...)

		// heights[idx[i]] 是递增的
		for i := len(zeros); i < n; i++ { // 高度 0 无需计算
			ans = max(ans, (n-i)*heights[idx[i]])
		}
	}

	return
}
