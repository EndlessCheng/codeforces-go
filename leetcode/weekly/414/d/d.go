package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func maxMoves(kx, ky int, positions [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	n := len(positions)
	// 计算马到兵的步数，等价于计算兵到其余格子的步数（马走日）
	dis := make([][50][50]int, n)
	for i, pos := range positions {
		d := &dis[i]
		for j := range d {
			for k := range d[j] {
				d[j][k] = -1
			}
		}
		px, py := pos[0], pos[1]
		d[px][py] = 0
		q := []pair{{px, py}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y
					if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
						d[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}

	positions = append(positions, []int{kx, ky})
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for mask := 1; mask < 1<<n; mask++ {
		for i, p := range positions {
			x, y := p[0], p[1]
			odd := bits.OnesCount(uint(u^mask))%2 > 0
			if odd {
				f[mask][i] = math.MaxInt
			}
			op := func(a, b int) int {
				if odd {
					return min(a, b)
				}
				return max(a, b)
			}
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				f[mask][i] = op(f[mask][i], f[mask^1<<j][j]+dis[j][x][y])
			}
		}
	}
	return f[u][n]
}

func maxMoves2(kx, ky int, positions [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	n := len(positions)
	// 计算马到兵的步数，等价于计算兵到其余格子的步数（马走日）
	dis := make([][50][50]int, n)
	for i, pos := range positions {
		d := &dis[i]
		for j := range d {
			for k := range d[j] {
				d[j][k] = -1
			}
		}
		px, py := pos[0], pos[1]
		d[px][py] = 0
		q := []pair{{px, py}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y
					if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
						d[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}

	positions = append(positions, []int{kx, ky})
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, 1<<n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	u := 1<<n - 1
	var dfs func(int, int) int
	dfs = func(i, mask int) int {
		if mask == 0 {
			return 0
		}
		p := &memo[i][mask]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		x, y := positions[i][0], positions[i][1]
		if bits.OnesCount(uint(u^mask))%2 == 0 { // Alice
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				res = max(res, dfs(j, mask^1<<j)+dis[j][x][y])
			}
		} else { // Bob
			res = math.MaxInt
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				res = min(res, dfs(j, mask^1<<j)+dis[j][x][y])
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n, u)
}
