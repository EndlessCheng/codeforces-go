package main

import "sort"

func sortItems(n int, m int, group []int, beforeItems [][]int) (ans []int) {
	topSort := func(g [][]int, inDeg []int) (orders []int) {
		q := []int{}
		for i, deg := range inDeg {
			if deg == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			var v int
			v, q = q[0], q[1:]
			orders = append(orders, v)
			for _, w := range g[v] {
				inDeg[w]--
				if inDeg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		return
	}

	blocks := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i
		}
		blocks[group[i]] = append(blocks[group[i]], i)
	}

	g := make([][]int, n)
	inDeg := make([]int, n)
	g2 := make([][]int, m+n)
	inDeg2 := make([]int, m+n)
	for w, items := range beforeItems {
		for _, v := range items {
			g[v] = append(g[v], w)
			inDeg[w]++
			if gv, gw := group[v], group[w]; gv != gw {
				g2[gv] = append(g2[gv], gw)
				inDeg2[gw]++
			}
		}
	}

	orders, orders2 := topSort(g, inDeg), topSort(g2, inDeg2)
	if len(orders) < len(g) || len(orders2) < len(g2) {
		return
	}
	pos := make([]int, n)
	for i, v := range orders {
		pos[v] = i
	}
	for _, b := range blocks {
		sort.Slice(b, func(i, j int) bool { return pos[b[i]] < pos[b[j]] })
	}
	for _, ord := range orders2 {
		ans = append(ans, blocks[ord]...)
	}
	return
}
