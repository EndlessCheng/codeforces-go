package main

import "slices"

// https://space.bilibili.com/206214
func canPartitionGrid(grid [][]int) bool {
	total := 0
	for _, row := range grid {
		for _, x := range row {
			total += x
		}
	}

	// 能否水平分割
	check := func(a [][]int) bool {
		m, n := len(a), len(a[0])
		f := func() bool {
			has := map[int]bool{0: true} // 0 对应不删除数字
			s := 0
			for i, row := range a[:m-1] {
				for j, x := range row {
					s += x
					// 第一行，不能删除中间元素
					if i > 0 || j == 0 || j == n-1 {
						has[x] = true
					}
				}
				// 特殊处理只有一列的情况，此时只能删除第一个数或者分割线上那个数
				if n == 1 {
					if s*2 == total || s*2-total == a[0][0] || s*2-total == row[0] {
						return true
					}
					continue
				}
				if has[s*2-total] {
					return true
				}
				// 如果分割到更下面，那么可以删第一行的元素
				if i == 0 {
					for _, x := range row {
						has[x] = true
					}
				}
			}
			return false
		}
		// 删除上半部分中的一个数
		if f() {
			return true
		}
		slices.Reverse(a)
		// 删除下半部分中的一个数
		return f()
	}

	// 水平分割 or 垂直分割
	return check(grid) || check(rotate(grid))
}

// 顺时针旋转矩阵 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
