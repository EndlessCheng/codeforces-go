package main

// github.com/EndlessCheng/codeforces-go
var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func latestDayToCross(row, col int, cells [][]int) int {
	top := row * col
	bottom := top + 1
	uf := newUnionFind(bottom + 1)
	land := make([][]bool, row)
	for i := range land {
		land[i] = make([]bool, col)
	}
	// 倒序遍历天数，如果最上和最下连通了，这一天就是答案
	for day := len(cells) - 1; ; day-- {
		p := cells[day]
		r, c := p[0]-1, p[1]-1
		v := r*col + c
		for _, d := range dir4 {
			if x, y := r+d.x, c+d.y; 0 <= x && x < row && 0 <= y && y < col && land[x][y] { // 与四周的陆地相连
				uf.merge(v, x*col+y)
			}
		}
		land[r][c] = true // 将该位置标记为陆地
		if r == 0 {
			uf.merge(v, top) // 与最上面相连
		}
		if r == row-1 {
			uf.merge(v, bottom) // 与最下面相连
		}
		if uf.same(top, bottom) {
			return day // 最上和最下连通了，返回答案
		}
	}
}

// 并查集模板
type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x != y {
		u.fa[x] = y
	}
}

func (u uf) same(x, y int) bool {
	return u.find(x) == u.find(y)
}
