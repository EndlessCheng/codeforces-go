package main

func countSquares(mat [][]int) int {
	ans := 0
	n, m := len(mat), len(mat[0])

	const mx = 305
	sumRow := [mx][mx]int{}
	for i, mi := range mat {
		for j, mij := range mi {
			sumRow[i][j+1] = sumRow[i][j] + mij
		}
	}

	sumCol := [mx][mx]int{}
	for j := range mat[0] {
		for i := range mat {
			sumCol[i+1][j] = sumCol[i][j] + mat[i][j]
		}
	}

	for i, mi := range mat {
		for j, mij := range mi {
			if mij == 0 {
				continue
			}
			ans++
			for sz := 1; i+sz < n && j+sz < m; sz++ {
				if sumRow[i+sz][j+sz+1]-sumRow[i+sz][j] == sz+1 &&
					sumCol[i+sz+1][j+sz]-sumCol[i][j+sz] == sz+1 {
					ans++
				} else {
					break
				}
			}
		}
	}
	return ans
}
