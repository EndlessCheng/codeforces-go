package main

// github.com/EndlessCheng/codeforces-go
func largestMagicSquare1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowSum := make([][]int, m)    // → 前缀和
	colSum := make([][]int, m+1)  // ↓ 前缀和
	diagSum := make([][]int, m+1) // ↘ 前缀和
	antiSum := make([][]int, m+1) // ↙ 前缀和
	for i := range m + 1 {
		colSum[i] = make([]int, n)
		diagSum[i] = make([]int, n+1)
		antiSum[i] = make([]int, n+1)
	}
	for i, row := range grid {
		rowSum[i] = make([]int, n+1)
		for j, x := range row {
			rowSum[i][j+1] = rowSum[i][j] + x
			colSum[i+1][j] = colSum[i][j] + x
			diagSum[i+1][j+1] = diagSum[i][j] + x
			antiSum[i+1][j] = antiSum[i][j+1] + x
		}
	}

	// k×k 子矩阵的左上角为 (i−k, j−k)，右下角为 (i−1, j−1)
	for k := min(m, n); ; k-- {
		for i := k; i <= m; i++ {
		next:
			for j := k; j <= n; j++ {
				// 子矩阵主对角线的和
				sum := diagSum[i][j] - diagSum[i-k][j-k]

				// 子矩阵反对角线的和
				if antiSum[i][j-k]-antiSum[i-k][j] != sum {
					continue
				}

				// 子矩阵每行的和
				for _, rowS := range rowSum[i-k : i] {
					if rowS[j]-rowS[j-k] != sum {
						continue next
					}
				}

				// 子矩阵每列的和
				for c := j - k; c < j; c++ {
					if colSum[i][c]-colSum[i-k][c] != sum {
						continue next
					}
				}

				return k
			}
		}
	}
}

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowSum := make([][]int, m)    // → 前缀和
	colSum := make([][]int, m+1)  // ↓ 前缀和
	diagSum := make([][]int, m+1) // ↘ 前缀和
	antiSum := make([][]int, m+1) // ↙ 前缀和
	for i := range m + 1 {
		colSum[i] = make([]int, n)
		diagSum[i] = make([]int, n+1)
		antiSum[i] = make([]int, n+1)
	}
	for i, row := range grid {
		rowSum[i] = make([]int, n+1)
		for j, x := range row {
			rowSum[i][j+1] = rowSum[i][j] + x
			colSum[i+1][j] = colSum[i][j] + x
			diagSum[i+1][j+1] = diagSum[i][j] + x
			antiSum[i+1][j] = antiSum[i][j+1] + x
		}
	}

	// isSameColSum[i][j] 表示右下角为 (i, j) 的子矩形，每列元素和是否都相等
	isSameColSum := make([][]bool, m)
	for i := range isSameColSum {
		isSameColSum[i] = make([]bool, n)
	}
	for k := min(m, n); k > 1; k-- {
		for i := k; i <= m; i++ {
			// 想象有一个 k×k 的窗口在向右滑动
			sameCnt := 1
			for j := 1; j < n; j++ {
				if colSum[i][j]-colSum[i-k][j] == colSum[i][j-1]-colSum[i-k][j-1] {
					sameCnt++
				} else {
					sameCnt = 1
				}
				// 连续 k 列元素和是否都一样
				isSameColSum[i-1][j] = sameCnt >= k
			}
		}

		for j := k; j <= n; j++ {
			// 想象有一个 k×k 的窗口在向下滑动
			sum := rowSum[0][j] - rowSum[0][j-k]
			sameCnt := 1
			for i := 2; i <= m; i++ {
				rowS := rowSum[i-1][j] - rowSum[i-1][j-k]
				if rowS == sum {
					sameCnt++
					if sameCnt >= k && // 连续 k 行元素和都一样
						isSameColSum[i-1][j-1] && // 连续 k 列元素和都一样
						colSum[i][j-1]-colSum[i-k][j-1] == sum && // 列和 = 行和
						diagSum[i][j]-diagSum[i-k][j-k] == sum && // 主对角线和 = 行和
						antiSum[i][j-k]-antiSum[i-k][j] == sum {  // 反对角线和 = 行和
						return k
					}
				} else {
					sum = rowS
					sameCnt = 1
				}
			}
		}
	}

	return 1
}
