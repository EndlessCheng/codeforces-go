package main

func minPushBox(g [][]byte) int {
	n, m := len(g), len(g[0])
	type point struct{ x, y int }
	type pair struct {
		point
		di, dep int
	}

	findOne := func(tar byte) point {
		for i, row := range g {
			for j, b := range row {
				if b == tar {
					return point{i, j}
				}
			}
		}
		panic(1)
	}
	s, t, b := findOne('S'), findOne('T'), findOne('B')

	valid := func(p point) bool { return p.x >= 0 && p.x < n && p.y >= 0 && p.y < m && g[p.x][p.y] != '#' }
	dir4 := [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	reachable := func(s, t point) bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		for q := []point{s}; len(q) > 0; {
			p := q[0]
			q = q[1:]
			if p == t {
				return true
			}
			for _, d := range dir4 {
				if pp := (point{p.x + d[0], p.y + d[1]}); valid(pp) && !vis[pp.x][pp.y] {
					vis[pp.x][pp.y] = true
					q = append(q, pp)
				}
			}
		}
		return false
	}

	vis := make([][][4]bool, n)
	for i := range vis {
		vis[i] = make([][4]bool, m)
	}
	q := []pair{}
	g[b.x][b.y] = '#'
	for i, d := range dir4 {
		if man, nb := (point{b.x + d[0], b.y + d[1]}), (point{b.x - d[0], b.y - d[1]}); valid(man) && valid(nb) && reachable(s, man) {
			vis[b.x][b.y][i] = true
			q = append(q, pair{b, i, 1})
		}
	}
	g[b.x][b.y] = 'b'
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		curBox, dir, dep := p.point, dir4[p.di], p.dep
		b = point{curBox.x - dir[0], curBox.y - dir[1]}
		if b == t {
			return p.dep
		}
		g[b.x][b.y] = '#'
		for i, d := range dir4 {
			if !vis[b.x][b.y][i] {
				if man, nb := (point{b.x + d[0], b.y + d[1]}), (point{b.x - d[0], b.y - d[1]}); valid(man) && valid(nb) && reachable(curBox, man) {
					vis[b.x][b.y][i] = true
					q = append(q, pair{b, i, dep + 1})
				}
			}
		}
		g[b.x][b.y] = 'b'
	}
	return -1
}
