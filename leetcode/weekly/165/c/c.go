package main

func countSquares(mat [][]int) (ans int) {
	sum := make([][]int, len(mat)+1)
	sum[0] = make([]int, len(mat[0])+1)
	for i, mi := range mat {
		sum[i+1] = make([]int, len(mi)+1)
		for j, mij := range mi {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + mij
		}
	}
	query := func(r1, c1, r2, c2 int) int {
		return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
	}
	for i, mi := range mat {
		for j := range mi {
			for sz := 1; i+sz <= len(mat) && j+sz <= len(mi) && query(i, j, i+sz, j+sz) == sz*sz; sz++ {
				ans++
			}
		}
	}
	return ans
}
