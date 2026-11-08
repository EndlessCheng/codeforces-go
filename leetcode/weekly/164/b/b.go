package main

func countServers(grid [][]int) (ans int) {
	// 统计每行每列的服务器个数
	rowCnt := make([]int, len(grid))
	colCnt := make([]int, len(grid[0]))
	for i, row := range grid {
		for j, x := range row {
			if x > 0 {
				rowCnt[i]++
				colCnt[j]++
			}
		}
	}

	// 判断每台服务器是否与另一台服务器同行或同列
	for i, row := range grid {
		for j, x := range row {
			if x > 0 && (rowCnt[i] > 1 || colCnt[j] > 1) {
				ans++
			}
		}
	}
	return
}
