package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func highestPeak(g [][]int) (ans [][]int) {
	n, m := len(g), len(g[0])
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	q := []pair{}
	for i, r := range g {
		for j, v := range r {
			if v > 0 {
				ans[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}
	for h := 1; len(q) > 0; h++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && ans[x][y] == -1 {
					ans[x][y] = h
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return
}
