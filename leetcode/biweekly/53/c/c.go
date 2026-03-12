package main

// github.com/EndlessCheng/codeforces-go
func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	diagSum := make([][]int, m+1) // ↘ 前缀和
	antiSum := make([][]int, m+1) // ↙ 前缀和
	for i := range diagSum {
		diagSum[i] = make([]int, n+1)
		antiSum[i] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, v := range row {
			diagSum[i+1][j+1] = diagSum[i][j] + v
			antiSum[i+1][j] = antiSum[i][j+1] + v
		}
	}

	// 从 (x,y) 开始，向 ↘，连续 k 个数的和
	queryDiagonal := func(x, y, k int) int { return diagSum[x+k][y+k] - diagSum[x][y] }

	// 从 (x,y) 开始，向 ↙，连续 k 个数的和
	queryAntiDiagonal := func(x, y, k int) int { return antiSum[x+k][y+1-k] - antiSum[x][y+1] }

	var x, y, z int // 最大，次大，第三大
	update := func(v int) {
		if v > x {
			x, y, z = v, x, y
		} else if v < x && v > y {
			y, z = v, y
		} else if v < y && v > z {
			z = v
		}
	}

	// 枚举菱形正中心 (i,j)
	for i, row := range grid {
		for j, v := range row {
			update(v) // 一个数也算菱形
			// 枚举菱形顶点到正中心的距离 k，注意菱形顶点不能出界
			// i-k >= 0 且 i+k <= m-1，所以 k <= min(i, m-1-i)，对于 j 同理
			mx := min(i, m-1-i, j, n-1-j)
			for k := 1; k <= mx; k++ {
				a := queryDiagonal(i-k, j, k)           // 菱形右上的边
				b := queryDiagonal(i, j-k, k)           // 菱形左下的边
				c := queryAntiDiagonal(i-k+1, j-1, k-1) // 菱形左上的边
				d := queryAntiDiagonal(i, j+k, k+1)     // 菱形右下的边
				update(a + b + c + d)
			}
		}
	}

	ans := []int{x, y, z}
	for ans[len(ans)-1] == 0 { // 不同的和少于三个
		ans = ans[:len(ans)-1]
	}
	return ans
}
