package main

// github.com/EndlessCheng/codeforces-go
func getBiggestThree(a [][]int) []int {
	n, m := len(a), len(a[0])
	ds := make([][]int, n+1) // 主对角线前缀和
	as := make([][]int, n+1) // 反对角线前缀和
	for i := range ds {
		ds[i] = make([]int, m+1)
		as[i] = make([]int, m+1)
	}
	for i, r := range a {
		for j, v := range r {
			ds[i+1][j+1] = ds[i][j] + v // ↘
			as[i+1][j] = as[i][j+1] + v // ↙
		}
	}
	// 从 x,y 开始，向 ↘，连续的 k 个数的和
	queryDiagonal := func(x, y, k int) int { return ds[x+k][y+k] - ds[x][y] }
	// 从 x,y 开始，向 ↙，连续的 k 个数的和
	queryAntiDiagonal := func(x, y, k int) int { return as[x+k][y+1-k] - as[x][y+1] }

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

	for i, r := range a {
		for j, v := range r {
			update(v)
			for k := 1; k <= i && i+k < n && k <= j && j+k < m; k++ {
				a := queryDiagonal(i-k, j, k)           // 菱形右上
				b := queryAntiDiagonal(i-k+1, j-1, k-1) // 菱形左上
				c := queryDiagonal(i, j-k, k)           // 菱形左下
				d := queryAntiDiagonal(i, j+k, k+1)     // 菱形右下
				update(a + b + c + d)
			}
		}
	}

	ans := []int{x, y, z}
	for ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	return ans
}
