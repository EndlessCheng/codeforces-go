package main

// https://space.bilibili.com/206214
func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for j := n - 2; j >= 0; j-- {
		for i, row := range grid {
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if grid[k][j+1] > row[j] {
					f[i][j] = max(f[i][j], f[k][j+1]+1)
				}
			}
		}
	}
	for _, r := range f {
		ans = max(ans, r[0])
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }