package main

// https://space.bilibili.com/206214
func resultGrid(a [][]int, threshold int) [][]int {
	m, n := len(a), len(a[0])
	result := make([][]int, m)
	cnt := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}
	for i := 2; i < m; i++ {
	next:
		for j := 2; j < n; j++ {
			// 检查左右相邻格子
			for _, row := range a[i-2 : i+1] {
				if abs(row[j-2]-row[j-1]) > threshold || abs(row[j-1]-row[j]) > threshold {
					continue next // 不合法，下一个
				}
			}

			// 检查上下相邻格子
			for y := j - 2; y <= j; y++ {
				if abs(a[i-2][y]-a[i-1][y]) > threshold || abs(a[i-1][y]-a[i][y]) > threshold {
					continue next // 不合法，下一个
				}
			}

			// 合法，计算 3x3 区域的平均值
			sum := 0
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					sum += a[x][y]
				}
			}
			sum /= 9

			// 更新区域内的 result
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					result[x][y] += sum // 先累加，最后再求平均值
					cnt[x][y]++
				}
			}
		}
	}

	for i, row := range cnt {
		for j, c := range row {
			if c == 0 { // (i,j) 不属于任何区域
				result[i][j] = a[i][j]
			} else {
				result[i][j] /= c // 求平均值
			}
		}
	}
	return result
}

func abs(x int) int { if x < 0 { return -x }; return x }
