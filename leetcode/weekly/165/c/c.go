package main

func countSquares0(matrix [][]int) (ans int) {
	sum := make([][]int, len(matrix)+1)
	sum[0] = make([]int, len(matrix[0])+1)
	for i, mi := range matrix {
		sum[i+1] = make([]int, len(mi)+1)
		for j, mij := range mi {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + mij
		}
	}
	query := func(r1, c1, r2, c2 int) int {
		return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
	}
	for i, mi := range matrix {
		for j := range mi {
			for sz := 1; i+sz <= len(matrix) && j+sz <= len(mi) && query(i, j, i+sz, j+sz) == sz*sz; sz++ {
				ans++
			}
		}
	}
	return ans
}

func countSquares1(matrix [][]int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	for top := range m { // 枚举上边界
		a := make([]int, n)
		for bottom := top; bottom < m; bottom++ { // 枚举下边界
			h := bottom - top + 1 // 高
			// 2348. 全 h 子数组的数目
			last := -1
			for j := range n {
				a[j] += matrix[bottom][j] // 把 bottom 这一行的值加到 a 中
				if a[j] != h {
					last = j // 记录上一个非 h 元素的位置
				} else if j-last >= h { // 右端点为 j 的长为 h 的子数组全是 h
					ans++
				}
			}
		}
	}
	return
}

func countSquares(matrix [][]int) (ans int) {
	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && i > 0 && j > 0 {
				row[j] += min(matrix[i-1][j], matrix[i-1][j-1], row[j-1])
			}
			ans += row[j]
		}
	}
	return
}
