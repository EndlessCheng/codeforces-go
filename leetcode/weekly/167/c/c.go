package main

func maxSideLength(mat [][]int, threshold int) int {
	n, m := len(mat), len(mat[0])
	var sum [][]int
	initS := func(mat [][]int) {
		sum = make([][]int, len(mat)+1)
		sum[0] = make([]int, len(mat[0])+1)
		for i, mi := range mat {
			sum[i+1] = make([]int, len(mi)+1)
			for j, mij := range mi {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + mij
			}
		}
	}
	// r1<=r<=r2 && c1<=c<=c2
	query := func(r1, c1, r2, c2 int) int {
		r2++
		c2++
		return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
	}

	initS(mat)
	ans := 0
	for i, mi := range mat {
		for j := range mi {
			for sz := 0; i+sz < n && j+sz < m; sz++ {
				if query(i, j, i+sz, j+sz) <= threshold {
					if sz+1 > ans {
						ans = sz + 1
					}
				} else {
					break
				}
			}
		}
	}
	return ans
}
