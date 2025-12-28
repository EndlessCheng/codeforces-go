package main

// github.com/EndlessCheng/codeforces-go
type unionFind struct {
	fa []int // 代表元
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 判断 x 和 y 是否在同一个集合
func (u unionFind) same(x, y int) bool {
	// 如果 x 的代表元和 y 的代表元相同，那么 x 和 y 就在同一个集合
	// 这就是代表元的作用：用来快速判断两个元素是否在同一个集合
	return u.find(x) == u.find(y)
}

// 把 from 所在集合合并到 to 所在集合中
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
}

var dirs = []struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func latestDayToCross(m, n int, cells [][]int) int {
	left := m * n
	right := m*n + 1
	uf := newUnionFind(m*n + 2)

	water := make([][]bool, m)
	for i := range water {
		water[i] = make([]bool, n)
	}

	for i, cell := range cells {
		r, c := cell[0]-1, cell[1]-1 // 改成从 0 开始的下标
		v := r*n + c
		water[r][c] = true // 变成水

		if c == 0 {
			uf.merge(v, left) // 与最左边相连
		}

		if c == n-1 {
			uf.merge(v, right) // 与最右边相连
		}

		for _, d := range dirs {
			x, y := r+d.x, c+d.y
			if 0 <= x && x < m && 0 <= y && y < n && water[x][y] {
				uf.merge(v, x*n+y) // 与八方向的水相连
			}
		}

		if uf.same(left, right) { // 最左边和最右边连通
			return i
		}
	}
	return -1
}

func latestDayToCross22(m, n int, cells [][]int) int {
	// 0：陆地
	// 1：水（未被感染）
	// 2：水（已被感染）
	state := make([][]int8, m)
	for i := range state {
		state[i] = make([]int8, n)
	}

	// 能否从第一列到达 (r, c)
	canReachFromLeft := func(r, c int) bool {
		if c == 0 { // 已经是第一列
			return true
		}
		for _, d := range dirs {
			x, y := r+d.x, c+d.y
			if 0 <= x && x < m && 0 <= y && y < n && state[x][y] == 2 {
				return true
			}
		}
		return false
	}

	// 从 (r, c) 出发，能否到达最后一列
	var dfs func(int, int) bool
	dfs = func(r, c int) bool {
		if c == n-1 {
			return true
		}
		state[r][c] = 2 // 感染
		for _, d := range dirs {
			x, y := r+d.x, c+d.y
			// 传播病毒到未被感染的水
			if 0 <= x && x < m && 0 <= y && y < n && state[x][y] == 1 && dfs(x, y) {
				return true
			}
		}
		return false
	}

	for i, cell := range cells {
		r, c := cell[0]-1, cell[1]-1 // 改成从 0 开始的下标
		state[r][c] = 1              // 未被感染的水
		if canReachFromLeft(r, c) && dfs(r, c) {
			return i
		}
	}
	return -1
}
