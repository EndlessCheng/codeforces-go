package main

import "math"

// https://space.bilibili.com/206214
func maximumValueSum(board [][]int) int64 {
	m, n := len(board), len(board[0])
	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, m+n+3)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	R := m + n
	C := m + n + 1
	S := m + n + 2
	for i, row := range board {
		for j, x := range row {
			addEdge(i, m+j, 1, -x)
		}
		addEdge(R, i, 1, 0)
	}
	for j := range board[0] {
		addEdge(m+j, C, 1, 0)
	}
	addEdge(S, R, 3, 0) // 把 3 改成 k 可以支持 k 个车

	// 下面是费用流模板
	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[C] < math.MaxInt
	}

	minCost := 0
	for spfa() {
		minF := math.MaxInt
		for v := C; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := C; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[C] * minF
	}
	return int64(-minCost)
}

func maximumValueSum2(board [][]int) int64 {
	m := len(board)
	type pair struct{ x, j int }
	suf := make([][3]pair, m)
	p := [3]pair{} // 最大、次大、第三大
	for i := range p {
		p[i].x = math.MinInt
	}
	update := func(row []int) {
		for j, x := range row {
			if x > p[0].x {
				if p[0].j != j { // 如果相等，仅更新最大
					if p[1].j != j { // 如果相等，仅更新最大和次大
						p[2] = p[1]
					}
					p[1] = p[0]
				}
				p[0] = pair{x, j}
			} else if x > p[1].x && j != p[0].j {
				if p[1].j != j { // 如果相等，仅更新次大
					p[2] = p[1]
				}
				p[1] = pair{x, j}
			} else if x > p[2].x && j != p[0].j && j != p[1].j {
				p[2] = pair{x, j}
			}
		}
	}
	for i := m - 1; i > 1; i-- {
		update(board[i])
		suf[i] = p
	}

	ans := math.MinInt
	for i := range p {
		p[i].x = math.MinInt // 重置，计算 pre
	}
	for i, row := range board[:m-2] {
		update(row)
		for j, x := range board[i+1] { // 第二个车
			for _, p := range p { // 第一个车
				if p.j == j {
					continue
				}
				for _, q := range suf[i+2] { // 第三个车
					if q.j != j && q.j != p.j { // 没有同列的车
						ans = max(ans, p.x+x+q.x)
						break
					}
				}
			}
		}
	}
	return int64(ans)
}
