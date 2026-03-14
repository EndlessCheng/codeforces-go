package main

import "slices"

// github.com/EndlessCheng/codeforces-go
// 48. 旋转图像
func rotate(matrix [][]int) {
	n := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < n; j++ { // 遍历对角线上方元素，做转置
			row[j], matrix[j][i] = matrix[j][i], row[j]
		}
		slices.Reverse(row) // 行翻转
	}
}

func findRotation1(mat, target [][]int) bool {
	for range 4 {
		if slices.EqualFunc(mat, target, slices.Equal[[]int]) {
			return true
		}
		rotate(mat)
	}
	return false
}

func findRotation(mat, target [][]int) bool {
	n := len(mat)
	ok := 1<<4 - 1 // ok := []bool{true, true, true, true}
	for i, row := range mat {
		for j, x := range row {
			if x != target[i][j] {
				ok &^= 1 << 0 // ok[0] = false
			}
			if x != target[j][n-1-i] {
				ok &^= 1 << 1 // ok[1] = false
			}
			if x != target[n-1-i][n-1-j] {
				ok &^= 1 << 2 // ok[2] = false
			}
			if x != target[n-1-j][i] {
				ok &^= 1 << 3 // ok[3] = false
			}
			if ok == 0 { // 所有的 ok[i] 都是 false
				return false
			}
		}
	}
	return true // 至少有一个 ok[i] 是 true
}
