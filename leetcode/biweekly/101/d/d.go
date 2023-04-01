package main

import "math"

// https://space.bilibili.com/206214
func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	ans := math.MaxInt
	dis := make([]int, n) // dis[i] 表示从 start 到 i 的最短路长度
next:
	for start := 0; start < n; start++ { // 枚举每个起点跑 BFS
		for j := range dis {
			dis[j] = -1
		}
		dis[start] = 0
		type pair struct{ x, fa int }
		q := []pair{{start, -1}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, fa := p.x, p.fa
			for _, y := range g[x] {
				if dis[y] < 0 { // 第一次遇到
					dis[y] = dis[x] + 1
					q = append(q, pair{y, x})
				} else if y != fa { // 第二次遇到
					ans = min(ans, dis[x]+dis[y]+1)
					continue next // 由于是 BFS，后面不会遇到更短的环了，直接枚举下一个 start
				}
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
