package main

func countSquares1(matrix [][]int) (ans int) {
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
