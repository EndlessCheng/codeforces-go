package main

func minimalSteps(grids []string) (ans int) {
	const inf int = 1e9
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	tsp := func(dist [][]int, st int) []int {
		n := len(dist)
		dp := make([][]int, 1<<n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = inf
			}
		}
		dp[1<<n-1][st] = 0
		for s := 1<<n - 2; s >= 0; s-- {
			for v := 0; v < n; v++ {
				for w := 0; w < n; w++ {
					if s>>w&1 == 0 {
						dp[s][v] = min(dp[s][v], dp[s|1<<w][w]+dist[v][w])
					}
				}
			}
		}
		return dp[0]
	}

	n, m := len(grids), len(grids[0])
	g := make([][]byte, n)
	for i := range g {
		g[i] = []byte(grids[i])
	}

	type point struct{ x, y int }
	find := func(tar byte) point {
		for i, row := range g {
			for j, b := range row {
				if b == tar {
					return point{i, j}
				}
			}
		}
		return point{-1, -1}
	}
	count := func(tar byte) (cnt int) {
		for _, row := range g {
			for _, b := range row {
				if b == tar {
					cnt++
				}
			}
		}
		return
	}
	type pair struct {
		point
		dep int
	}
	dir4 := [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	bfsDis := func(s, t point) int {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		for q := []pair{{s, 0}}; len(q) > 0; {
			p := q[0]
			q = q[1:]
			if p.point == t {
				return p.dep
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d[0], p.y+d[1]; xx >= 0 && xx < n && yy >= 0 && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, pair{point{xx, yy}, p.dep + 1})
				}
			}
		}
		return -1
	}
	findAllTarget := func(s point, tar byte) (ps []point) {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[s.x][s.y] = true
		for q := []point{s}; len(q) > 0; {
			p := q[0]
			q = q[1:]
			if g[p.x][p.y] == tar {
				ps = append(ps, p)
			}
			for _, d := range dir4 {
				if xx, yy := p.x+d[0], p.y+d[1]; xx >= 0 && xx < n && yy >= 0 && yy < m && !vis[xx][yy] && g[xx][yy] != '#' {
					vis[xx][yy] = true
					q = append(q, point{xx, yy})
				}
			}
		}
		return
	}

	// 特殊情况：无法到达终点
	s, t := find('S'), find('T')
	dst := bfsDis(s, t)
	if dst == -1 {
		return -1
	}

	// 特殊情况：没有机关
	c := count('M')
	if c == 0 {
		return dst
	}

	// 特殊情况：有些机关无法达到，或者无法达到任何石头
	mps, ops := findAllTarget(s, 'M'), findAllTarget(s, 'O')
	if len(mps) != c || len(ops) == 0 {
		return -1
	}

	// 一定能拿到宝藏
	mps = append(mps, s)
	cntM := len(mps)
	disMO := make([][]int, cntM)
	for i, mp := range mps {
		disMO[i] = make([]int, len(ops))
		for j, op := range ops {
			disMO[i][j] = bfsDis(mp, op)
		}
	}
	disMT := make([]int, cntM)
	for i, mp := range mps {
		disMT[i] = bfsDis(mp, t)
	}

	dist := make([][]int, cntM)
	for i := range dist {
		dist[i] = make([]int, cntM)
		for j := range dist {
			if j != i {
				dist[i][j] = inf
				for k := range ops {
					dist[i][j] = min(dist[i][j], disMO[i][k]+disMO[j][k])
				}
			}
		}
	}
	dp := tsp(dist, cntM-1)
	ans = inf
	for i, v := range dp[:cntM-1] {
		ans = min(ans, v+disMT[i])
	}
	return
}
