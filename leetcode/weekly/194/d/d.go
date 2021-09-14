package main

import "sort"

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) (ans [][]int) {
	var fa []int
	initFa := func() {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	isConn := func() bool {
		c := 0
		for i := range fa {
			if find(i) == i {
				c++
			}
		}
		return c == 1
	}

	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	sum := 0
	initFa()
	for _, e := range edges {
		if from, to := find(e[0]), find(e[1]); from != to {
			sum += e[2]
			fa[from] = to
		}
	}

	var ans1, ans2 []int
	for _, ei := range edges {
		eid := ei[3]

		s := 0
		initFa()
		for _, e := range edges {
			if e[3] == eid {
				continue
			}
			if from, to := find(e[0]), find(e[1]); from != to {
				s += e[2]
				fa[from] = to
			}
		}
		if s > sum || !isConn() {
			ans1 = append(ans1, eid)
			continue
		}

		s = ei[2]
		initFa()
		fa[ei[0]] = ei[1]
		for _, e := range edges {
			if from, to := find(e[0]), find(e[1]); from != to {
				s += e[2]
				fa[from] = to
			}
		}
		if s == sum && isConn() {
			ans2 = append(ans2, eid)
		}
	}

	return [][]int{ans1, ans2}
}
