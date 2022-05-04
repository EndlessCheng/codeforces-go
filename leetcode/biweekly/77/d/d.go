package main

import (
	"fmt"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := sort.Search(m*n+1, func(t int) bool {
		fire := make([][]bool, m)
		for i := range fire {
			fire[i] = make([]bool, n)
		}
		f := []pair{}
		for i, row := range grid {
			for j, v := range row {
				if v == 1 {
					fire[i][j] = true
					f = append(f, pair{i, j})
				}
			}
		}
		spreadFire := func() {
			tmp := f
			f = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] != 2 {
						fire[x][y] = true
						f = append(f, pair{x, y})
					}
				}
			}
		}
		fmt.Println(t)
		for ; t > 0 && len(f) > 0; t-- {
			spreadFire() // 扩充至多 t 分钟的火势
		}
		if fire[0][0] { // 起点着火，寄
			return true
		}

		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		q := []pair{{}}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, p := range tmp {
				if !fire[p.x][p.y] {
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && !fire[x][y] && grid[x][y] != 2 {
							if x == m-1 && y == n-1 { // 我们安全了…暂时。
								return false
							}
							vis[x][y] = true
							q = append(q, pair{x, y})
						}
					}
				}
			}
			fmt.Println(len(q))
			spreadFire() // 扩充 1 分钟的火势
		}
		fmt.Println()
		return true
	}) - 1
	return 111
	if ans < m*n {
		return ans
	}
	return 1e9
}
