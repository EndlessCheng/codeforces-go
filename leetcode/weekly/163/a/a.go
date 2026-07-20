package main

func shiftGrid1(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	size := m * n
	for i, row := range grid {
		for j, x := range row {
			p := (i*n + j + k) % size
			ans[p/n][p%n] = x
		}
	}
	return ans
}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])

	reverse := func(l, r int) {
		for l < r {
			x1, y1 := l/n, l%n
			x2, y2 := r/n, r%n
			grid[x1][y1], grid[x2][y2] = grid[x2][y2], grid[x1][y1]
			l++
			r--
		}
	}

	// 189. 轮转数组
	size := m * n
	k %= size // 轮转 k 次等于轮转 k % size 次
	reverse(0, size-1)
	reverse(0, k-1)
	reverse(k, size-1)
	return grid
}
