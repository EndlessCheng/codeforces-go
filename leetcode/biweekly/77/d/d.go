package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	bfs := func(q []pair) (int, int, int) {
		time := make([][]int, m)
		for i := range time {
			time[i] = make([]int, n)
			for j := range time[i] {
				time[i][j] = -1
			}
		}
		for _, p := range q {
			time[p.x][p.y] = 0
		}
		for t := 1; len(q) > 0; t++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0 {
						time[x][y] = t
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return time[m-1][n-1], time[m-1][n-2], time[m-2][n-1]
	}

	manToHouseTime, m1, m2 := bfs([]pair{{}})
	if manToHouseTime < 0 {
		return -1 // 人无法到终点
	}

	fires := []pair{}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				fires = append(fires, pair{i, j})
			}
		}
	}
	fireToHouseTime, f1, f2 := bfs(fires)
	if fireToHouseTime < 0 {
		return 1e9 // 火无法到终点
	}

	ans := fireToHouseTime - manToHouseTime
	if ans < 0 {
		return -1 // 火比人先到终点
	}
	if m1 < 0 || m2 < 0 || f1-m1 == ans && f2-m2 == ans {
		return ans - 1 // 火只会跟在人的后面，在到达终点前，人和火不能重合
	}
	return ans // 人和火可以同时到终点
}
