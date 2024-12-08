package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	S := n * 2
	T := S + 1

	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, T+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i, s := range strength {
		// 枚举这个锁是第几次开的
		for j := range n {
			x := 1 + k*j
			addEdge(i, n+j, 1, (s-1)/x+1)
		}
		addEdge(S, i, 1, 0)
	}
	for i := n; i < n*2; i++ {
		addEdge(i, T, 1, 0)
	}

	// 下面是最小费用最大流模板
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
		// 沿 st-end 的最短路尽量增广
		// 特别地，如果建图时所有边的容量都设为 1，那么 minF 必然为 1，下面第一个 for 循环可以省略
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
	return minCost
}

func findMinimumTime3(strength []int, k int) int {
	n := len(strength)
	m := 1 << n
	f := make([]int, m)
	for i := 1; i < m; i++ {
		x := 1 + k*(n-bits.OnesCount(uint(i)))
		f[i] = math.MaxInt
		for j, s := range strength {
			if i>>j&1 > 0 {
				f[i] = min(f[i], f[i^1<<j]+(s-1)/x+1)
			}
		}
	}
	return f[m-1]
}

func findMinimumTime2(strength []int, k int) int {
	n := len(strength)
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		x := 1 + k*(n-bits.OnesCount(uint(i)))
		res := math.MaxInt
		for j, s := range strength {
			if i>>j&1 > 0 {
				res = min(res, dfs(i^1<<j)+(s-1)/x+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(1<<n - 1)
}

func findMinimumTime1(strength []int, k int) int {
	ans := math.MaxInt
	n := len(strength)
	done := make([]bool, n)
	var dfs func(int, int)
	dfs = func(i, time int) {
		// 最优性剪枝：答案不可能变小
		if time >= ans {
			return
		}
		if i == n {
			ans = time
			return
		}
		x := 1 + k*i
		for j, d := range done {
			if !d {
				done[j] = true // 已开锁
				dfs(i+1, time+(strength[j]-1)/x+1)
				done[j] = false // 恢复现场
			}
		}
	}
	dfs(0, 0)
	return ans
}
