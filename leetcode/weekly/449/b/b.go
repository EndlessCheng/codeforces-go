package main

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
		s := 0
		for _, row := range a[:len(a)-1] {
			for _, x := range row {
				s += x
			}
			if s*2 == total {
				return true
			}
		}
		return false
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
