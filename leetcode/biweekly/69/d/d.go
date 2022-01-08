package main

// github.com/EndlessCheng/codeforces-go
func possibleToStamp(grid [][]int, stampHeight, stampWidth int) bool {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	diff := make([][]int, m+1)
	diff[0] = make([]int, n+1)
	for i, row := range grid {
		sum[i+1] = make([]int, n+1)
		for j, v := range row { // grid 的二维前缀和
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
		diff[i+1] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				x, y := i+stampHeight, j+stampWidth // 注意这是矩形右下角横纵坐标都 +1 后的位置
				if x <= m && y <= n && sum[x][y]-sum[x][j]-sum[i][y]+sum[i][j] == 0 {
					diff[i][j]++
					diff[i][y]--
					diff[x][j]--
					diff[x][y]++ // 更新二维差分
				}
			}
		}
	}

	// 还原二维差分矩阵对应的计数矩阵，这里用滚动数组实现
	cnt := make([]int, n+1)
	pre := make([]int, n+1)
	for i, row := range grid {
		for j, v := range row {
			cnt[j+1] = cnt[j] + pre[j+1] - pre[j] + diff[i][j]
			if v == 0 && cnt[j+1] == 0 {
				return false
			}
		}
		cnt, pre = pre, cnt
	}
	return true
}
