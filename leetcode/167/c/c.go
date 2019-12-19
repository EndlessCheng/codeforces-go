package main

func maxSideLength(mat [][]int, threshold int) int {
	n, m := len(mat), len(mat[0])
	var sum2d [][]int
	initSum2D := func(mat [][]int) {
		sum2d = make([][]int, len(mat)+1)
		sum2d[0] = make([]int, len(mat[0])+1)
		for i, mi := range mat {
			sum2d[i+1] = make([]int, len(mi)+1)
			for j, mij := range mi {
				sum2d[i+1][j+1] = sum2d[i+1][j] + sum2d[i][j+1] - sum2d[i][j] + mij
			}
		}
	}
	// r1<=r<=r2 && c1<=c<=c2
	querySum2D := func(r1, c1, r2, c2 int) int {
		r2++
		c2++
		return sum2d[r2][c2] - sum2d[r2][c1] - sum2d[r1][c2] + sum2d[r1][c1]
	}

	initSum2D(mat)
	ans := 0
	for i, mi := range mat {
		for j := range mi {
			for sz := 0; i+sz < n && j+sz < m; sz++ {
				if querySum2D(i, j, i+sz, j+sz) <= threshold {
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
