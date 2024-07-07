package main

// https://space.bilibili.com/206214
func numberOfSubmatrices(grid [][]byte) (ans int) {
	colCnt := make([][2]int, len(grid[0]))
	for _, row := range grid {
		s0, s1 := 0, 0
		for j, c := range row {
			if c != '.' {
				colCnt[j][c&1]++
			}
			s0 += colCnt[j][0]
			s1 += colCnt[j][1]
			if s0 > 0 && s0 == s1 {
				ans++
			}
		}
	}
	return
}

func numberOfSubmatrices2(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	sum := make([][][2]int, m+1)
	for i := range sum {
		sum[i] = make([][2]int, n+1)
	}
	for i, row := range grid {
		for j, c := range row {
			sum[i+1][j+1][0] = sum[i+1][j][0] + sum[i][j+1][0] - sum[i][j][0]
			sum[i+1][j+1][1] = sum[i+1][j][1] + sum[i][j+1][1] - sum[i][j][1]
			if c != '.' {
				sum[i+1][j+1][c&1]++
			}
			if sum[i+1][j+1][0] > 0 && sum[i+1][j+1][0] == sum[i+1][j+1][1] {
				ans++
			}
		}
	}
	return
}
