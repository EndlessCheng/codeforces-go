package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func challengeOfTheKeeper(maze []string) int {
	m, n := len(maze), len(maze[0])

	// 1. 找到起点终点坐标
	var sx, sy, tx, ty int
	for i, row := range maze {
		for j, c := range row {
			if c == 'S' {
				sx, sy = i, j
			} else if c == 'T' {
				tx, ty = i, j
			}
		}
	}

	// 2. BFS 计算终点到其余点的最短距离
	disFromT := make([][]int, m)
	for i := range disFromT {
		disFromT[i] = make([]int, n)
		for j := range disFromT[i] {
			disFromT[i][j] = math.MaxInt
		}
	}
	disFromT[tx][ty] = 0
	q := []pair{{tx, ty}}
	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && maze[x][y] != '#' && disFromT[x][y] == math.MaxInt {
					disFromT[x][y] = step
					q = append(q, pair{x, y})
				}
			}
		}
	}

	// 3. 剪枝：如果 S 无法到达 T，直接返回 -1
	if disFromT[sx][sy] == math.MaxInt {
		return -1
	}

	// 4. 二分答案
	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
	}
	ans := sort.Search(m*n+1, func(maxDis int) bool {
		// DFS，看能否在「附加负面效果」的情况下，移动不超过 maxDis 步到达终点
		var dfs func(int, int) bool
		dfs = func(i, j int) bool {
			if i < 0 || i >= m || j < 0 || j >= n || vis[i][j] == maxDis+1 || maze[i][j] == '#' {
				return false
			}
			if maze[i][j] == 'T' { // 到达终点
				return true
			}
			vis[i][j] = maxDis + 1 // 避免反复创建 vis，用一个每次二分都不一样的数来标记
			if maze[i][j] == '.' {
				// 守护者使用卷轴传送小扣，如果小扣无法在 maxDis 步内到达终点，则返回 false
				if x, y := i, n-1-j; maze[x][y] != '#' && disFromT[x][y] > maxDis {
					return false
				}
				if x, y := m-1-i, j; maze[x][y] != '#' && disFromT[x][y] > maxDis {
					return false
				}
			}
			// 枚举四个方向
			for _, d := range dirs {
				if dfs(i+d.x, j+d.y) { // 到达终点
					return true
				}
			}
			return false // 无法到达终点
		}
		return dfs(sx, sy)
	})
	if ans > m*n { // 守护者使用卷轴传送小扣，可以把小扣传送到一个无法到达终点的位置
		return -1
	}
	return ans
}
