package main

import "math"

// https://space.bilibili.com/206214
func countCoveredBuildings(n int, buildings [][]int) (ans int) {
	type pair struct{ min, max int }
	row := make([]pair, n+1)
	col := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		row[i].min = math.MaxInt
		col[i].min = math.MaxInt
	}

	add := func(m []pair, x, y int) {
		m[y].min = min(m[y].min, x)
		m[y].max = max(m[y].max, x)
	}
	isInner := func(m []pair, x, y int) bool {
		return m[y].min < x && x < m[y].max
	}

	for _, p := range buildings {
		x, y := p[0], p[1]
		add(row, x, y) // x 加到 row[y] 中
		add(col, y, x) // y 加到 col[x] 中
	}

	for _, p := range buildings {
		x, y := p[0], p[1]
		if isInner(row, x, y) && isInner(col, y, x) {
			ans++
		}
	}
	return
}

func countCoveredBuildings1(_ int, buildings [][]int) (ans int) {
	type pair struct{ min, max int }
	row := map[int]pair{}
	col := map[int]pair{}
	add := func(m map[int]pair, x, y int) {
		p := m[y]
		if p.min == 0 {
			m[y] = pair{x, x}
		} else {
			m[y] = pair{min(p.min, x), max(p.max, x)}
		}
	}
	isInner := func(m map[int]pair, x, y int) bool {
		p := m[y]
		return p.min < x && x < p.max
	}

	for _, p := range buildings {
		x, y := p[0], p[1]
		add(row, x, y) // x 加到 row[y] 中
		add(col, y, x) // y 加到 col[x] 中
	}

	for _, p := range buildings {
		x, y := p[0], p[1]
		if isInner(row, x, y) && isInner(col, y, x) {
			ans++
		}
	}
	return
}
