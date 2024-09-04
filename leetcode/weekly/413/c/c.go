package main

import (
	"math"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}
	slices.Sort(allNums) // 下面从大到小递归

	n := len(allNums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<len(grid))
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		// 枚举选第 k 行的 x
		// 如果循环结束后 res > 0，就不再递归不选的情况
		res := 0
		x := allNums[i]
		for t, lb := pos[x], 0; t > 0; t ^= lb {
			lb = t & -t    // lb = 1<<k，其中 k 是行号
			if j&lb == 0 { // 没选过第 k 行的数
				res = max(res, dfs(i-1, j|lb)+x)
			}
		}
		if res == 0 {
			// 不选 x
			res = dfs(i-1, j)
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, 0)
}

func maxScoreFlow(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	k := len(pos)
	m := len(grid)
	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, k+m+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	S := k + m
	T := k + m + 1
	i := 0
	for x, posMask := range pos {
		for t := uint(posMask); t > 0; t &= t - 1 {
			j := bits.TrailingZeros(t)
			addEdge(i, k+j, 1, 0)
		}
		addEdge(S, i, 1, -x)
		i++
	}
	for j := range grid {
		addEdge(k+j, T, 1, 0)
	}

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
		return dis[T] < math.MaxInt
	}

	minCost := 0
	for spfa() {
		minF := math.MaxInt
		for v := T; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := T; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[T] * minF
	}
	return -minCost
}

func maxScore1(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	f := make([]int, 1<<len(grid))
	for x, posMask := range pos {
		for j := range f {
			for t, lb := posMask, 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[j] = max(f[j], f[j|lb]+x)
				}
			}
		}
	}
	return f[0]
}

func maxScore2(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	f := make([][]int, len(allNums)+1)
	for i := range f {
		f[i] = make([]int, 1<<len(grid))
	}
	for i, x := range allNums {
		for j, v := range f[i] {
			f[i+1][j] = v // 不选 x
			for t, lb := pos[x], 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[i+1][j] = max(f[i+1][j], f[i][j|lb]+x) // 选第 k 行的 x
				}
			}
		}
	}
	return f[len(allNums)][0]
}

func maxScore3(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	n := len(allNums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<len(grid))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		// 不选 x
		res := dfs(i-1, j)
		// 枚举选第 k 行的 x
		x := allNums[i]
		for t, lb := pos[x], 0; t > 0; t ^= lb {
			lb = t & -t    // lb = 1<<k，其中 k 是行号
			if j&lb == 0 { // 没选过第 k 行的数
				res = max(res, dfs(i-1, j|lb)+x)
			}
		}
		*p = res
		return res
	}
	return dfs(n-1, 0)
}
