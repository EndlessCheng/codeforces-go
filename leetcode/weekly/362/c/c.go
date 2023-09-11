package main

// https://space.bilibili.com/206214
func minimumMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	src := m * n   // 超级源点
	dst := src + 1 // 超级汇点
	type edge struct{ to, rid, cap, cost int }
	g := make([][]edge, m*n+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], edge{to, len(g[to]), cap, cost})
		g[to] = append(g[to], edge{from, len(g[from]) - 1, 0, -cost})
	}
	for x, row := range grid {
		for y, v := range row {
			if v > 1 {
				addEdge(src, x*n+y, v-1, 0)
				for i, r := range grid {
					for j, w := range r {
						if w == 0 {
							addEdge(x*n+y, i*n+j, 1, abs(x-i)+abs(y-j))
						}
					}
				}
			} else if v == 0 {
				addEdge(x*n+y, dst, 1, 0)
			}
		}
	}

	// 下面是最小费用最大流模板
	const inf int = 1e9
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]int, len(g))
	timestamp := 0
	spfa := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[src] = 0
		timestamp++
		inQ[src] = timestamp
		q := []int{src}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = 0
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if inQ[w] != timestamp {
						inQ[w] = timestamp
						q = append(q, w)
					}
				}
			}
		}
		return dist[dst] < inf
	}
	ek := func() (maxFlow, minCost int) {
		for spfa() {
			// 沿 st-end 的最短路尽量增广
			minF := inf
			for v := dst; v != src; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := dst; v != src; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dist[dst] * minF
		}
		return
	}
	_, cost := ek()
	return cost
}

func abs(x int) int { if x < 0 { return -x }; return x }
