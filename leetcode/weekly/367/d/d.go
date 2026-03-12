package main

// https://space.bilibili.com/206214
func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	n, m := len(grid), len(grid[0])
	p := make([][]int, n)
	suf := 1 // 后缀乘积
	for i := n - 1; i >= 0; i-- {
		p[i] = make([]int, m)
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suf // p[i][j] 先初始化成后缀乘积
			suf = suf * grid[i][j] % mod
		}
	}

	pre := 1 // 前缀乘积
	for i, row := range grid {
		for j, x := range row {
			p[i][j] = p[i][j] * pre % mod // 乘上前缀乘积
			pre = pre * x % mod
		}
	}
	return p
}
