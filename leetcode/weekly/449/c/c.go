package main

import "slices"

// https://space.bilibili.com/206214
func maxScore(n int, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var cycle, chain []int
	var cntV, cntE int
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		cntV++
		cntE += len(g[x])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if b {
			continue
		}
		cntV, cntE = 0, 0
		dfs(i)
		if cntV*2 == cntE { // 环
			cycle = append(cycle, cntV)
		} else if cntE > 0 { // 链，但不考虑孤立点
			chain = append(chain, cntV)
		}
	}

	ans := 0
	cur := n
	f := func(sz int, isCycle bool) {
		l, r := cur-sz+1, cur
		for i := l; i < r-1; i++ {
			ans += i * (i + 2)
		}
		ans += r * (r - 1)
		if isCycle {
			ans += l * (l + 1)
		}
		cur -= sz
	}

	slices.Sort(cycle)
	for _, sz := range cycle {
		f(sz, true)
	}

	slices.SortFunc(chain, func(a, b int) int { return b - a })
	for _, sz := range chain {
		f(sz, false)
	}

	return int64(ans)
}
