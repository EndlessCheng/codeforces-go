package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func numberOfSets(n, maxDistance int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i { // g[i][i] = 0（也可以不加这个 if，下面判断 maxDistance 时要保证 j != i）
				g[i][j] = math.MaxInt / 2 // 防止加法溢出
			}
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	ans := 1 // s=0 一定满足要求
	f := make([][][]int, 1<<n)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
			for k := range f[i][j] {
				f[i][j][k] = math.MaxInt / 2
			}
		}
	}
	f[0] = g
	for s := uint(1); s < 1<<n; s++ {
		k := bits.TrailingZeros(s)
		t := s ^ (1 << k)
		ok := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[s][i][j] = min(f[t][i][j], f[t][i][k]+f[t][k][j])
				if ok && j < i && s>>i&1 != 0 && s>>j&1 != 0 && f[s][i][j] > maxDistance {
					ok = false
				}
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}

func numberOfSets2(n, maxDistance int, roads [][]int) (ans int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i { // g[i][i] = 0
				g[i][j] = math.MaxInt / 2 // 防止加法溢出
			}
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
next:
	for s := 0; s < 1<<n; s++ { // 枚举子集
		for i, row := range g {
			if s>>i&1 == 0 {
				continue
			}
			copy(f[i], row)
		}

		// Floyd
		for k := range f {
			if s>>k&1 == 0 {
				continue
			}
			for i := range f {
				if s>>i&1 == 0 {
					continue
				}
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}

		for i, di := range f {
			if s>>i&1 == 0 {
				continue
			}
			for j, dij := range di {
				if s>>j&1 > 0 && dij > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return
}
