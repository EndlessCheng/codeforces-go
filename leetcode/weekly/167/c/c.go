package main

func maxSideLength(mat [][]int, threshold int) (ans int) {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range mat {
		sum[i+1] = make([]int, n+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
		}
	}

	// 返回左上角在 (r1, c1)，右下角在 (r2, c2) 的子矩阵元素和
	query := func(r1, c1, r2, c2 int) int {
		return sum[r2+1][c2+1] - sum[r2+1][c1] - sum[r1][c2+1] + sum[r1][c1]
	}

	for i := range m {
		for j := range n {
			// 边长为 ans+1 的正方形，左上角在 (i, j)，右下角在 (i+ans, j+ans)
			for i+ans < m && j+ans < n && query(i, j, i+ans, j+ans) <= threshold {
				ans++
			}
		}
	}
	return
}
