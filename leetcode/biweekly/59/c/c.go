package main

import "math"

// https://space.bilibili.com/206214
func countPaths(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2 // 防止溢出
		}
	}
	for _, r := range roads {
		x, y, d := r[0], r[1], r[2]
		g[x][y] = d
		g[y][x] = d
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt / 2
	}
	f := make([]int, n)
	f[0] = 1
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x == n-1 {
			// 不可能找到比 dis[n-1] 更短，或者一样短的最短路了（注意本题边权都是正数）
			return f[n-1]
		}
		done[x] = true
		for y, d := range g[x] {
			newDis := dis[x] + d
			if newDis < dis[y] { 
				// 就目前来说，最短路必须经过 x
				dis[y] = newDis
				f[y] = f[x]
			} else if newDis == dis[y] {
				// 和之前求的最短路一样长
				f[y] = (f[y] + f[x]) % 1_000_000_007
			}
		}
	}
}
