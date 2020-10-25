package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func matrixRankTransform(a [][]int) (ans [][]int) {
	n, m := len(a), len(a[0])
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	allPos := map[int][][2]int{}
	for i, row := range a {
		for j, v := range row {
			allPos[v] = append(allPos[v], [2]int{i, j})
		}
	}
	type pair struct {
		v   int
		pos [][2]int
	}
	ps := []pair{}
	for v, pos := range allPos {
		np := len(pos)
		g := make([][]int, np)
		for i := 1; i < np; i++ {
			if pos[i][0] == pos[i-1][0] {
				g[i] = append(g[i], i-1)
				g[i-1] = append(g[i-1], i)
			}
		}
		pid := map[[2]int]int{}
		col := map[int][]int{}
		for i, p := range pos {
			pid[p] = i
			col[p[1]] = append(col[p[1]], p[0])
		}
		for j, is := range col {
			for k := 1; k < len(is); k++ {
				i := pid[[2]int{is[k-1], j}]
				i2 := pid[[2]int{is[k], j}]
				g[i] = append(g[i], i2)
				g[i2] = append(g[i2], i)
			}
		}
		var pp [][2]int
		vis := make([]bool, np)
		var f func(int)
		f = func(v int) {
			vis[v] = true
			pp = append(pp, pos[v])
			for _, w := range g[v] {
				if !vis[w] {
					f(w)
				}
			}
			return
		}
		for i, b := range vis {
			if !b {
				pp = nil
				f(i)
				ps = append(ps, pair{v, pp})
			}
		}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })

	type rkv struct{ rk, v int }
	const inf int = 1e18
	row := make([]rkv, n)
	for i := range row {
		row[i] = rkv{-inf, -inf}
	}
	col := make([]rkv, m)
	for i := range col {
		col[i] = rkv{-inf, -inf}
	}

	for _, vp := range ps {
		v, pos, maxRk := vp.v, vp.pos, 0
		for _, p := range pos {
			i, j := p[0], p[1]
			r, c := row[i], col[j]
			if r.v < v {
				r.rk++
			}
			if c.v < v {
				c.rk++
			}
			rk := max(1, max(r.rk, c.rk))
			maxRk = max(maxRk, rk)
		}
		for _, p := range pos {
			i, j := p[0], p[1]
			ans[i][j] = maxRk
			row[i] = rkv{maxRk, v}
			col[j] = rkv{maxRk, v}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
