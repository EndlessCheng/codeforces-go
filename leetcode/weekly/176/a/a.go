package main

func countNegatives(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	i, j := 0, n-1 // 从右上角开始
	for i < m && j >= 0 { // 还有剩余元素
		if grid[i][j] < 0 {
			ans += m - i // 这一列剩余元素都是负数
			j--
		} else {
			i++ // 这一行剩余元素全都非负，排除
		}
	}
	return
}
