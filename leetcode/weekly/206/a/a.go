package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func numSpecial(mat [][]int) (ans int) {
	for _, row := range mat {
		rowSum := 0
		for _, x := range row {
			rowSum += x
		}
		if rowSum != 1 {
			continue
		}

		j := slices.Index(row, 1)
		// 计算 j 列的元素和，必须恰好是 1
		colSum := 0
		for _, r := range mat {
			colSum += r[j]
		}
		if colSum == 1 {
			ans++
		}
	}
	return
}

func numSpecial2(mat [][]int) (ans int) {
next:
	for _, row := range mat {
		col := -1
		for j, x := range row {
			if x == 0 {
				continue
			}
			if col >= 0 { // 这一行有多个 1
				continue next // 检查下一行
			}
			col = j // 记录 1 的列号
		}
		if col < 0 {
			continue
		}

		seen1 := false
		for _, r := range mat {
			if r[col] == 0 {
				continue
			}
			if seen1 { // 这一列有多个 1
				continue next // 检查下一行
			}
			seen1 = true // 这一列有 1
		}
		ans++ // 由于这一列一定有 1，循环结束时 seen1 一定是 true
	}
	return
}
