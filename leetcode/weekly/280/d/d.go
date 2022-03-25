package main

// github.com/EndlessCheng/codeforces-go
func maximumANDSum(nums []int, numSlots int) (ans int) {
	const inf int = 1e9

	// 集合 A 和 B 的大小
	n, m := len(nums), numSlots

	// 建图
	type neighbor struct{ to, rid, cap, cost int } // 相邻节点、反向边下标、容量、费用
	g := make([][]neighbor, n+m+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	start := n + m   // 超级源点
	end := start + 1 // 超级汇点
	for i, num := range nums {
		addEdge(start, i, 1, 0)
		for j := 1; j <= m; j++ {
			addEdge(i, n+j-1, inf, -(num & j))
		}
	}
	for i := 0; i < m; i++ {
		addEdge(n+i, end, 2, 0)
	}

	// 下面为最小费用最大流模板
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		inQ := make([]bool, len(g))
		inQ[start] = true
		q := []int{start}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[end] < inf
	}
	for spfa() {
		// 沿 start-end 的最短路尽量增广
		minFlow := inf
		for v := end; v != start; {
			p := fa[v]
			if c := g[p.v][p.i].cap; c < minFlow {
				minFlow = c
			}
			v = p.v
		}
		for v := end; v != start; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minFlow
			g[v][e.rid].cap += minFlow
			v = p.v
		}
		ans -= dist[end] * minFlow
	}
	return
}
