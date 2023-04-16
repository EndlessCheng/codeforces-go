package main

import "math"

// https://space.bilibili.com/206214
type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt // 初始化为无穷大，表示 i 到 j 没有边
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边）
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（输入保证这条边之前不存在）
}

func (g Graph) ShortestPath(node1, node2 int) int {
	ans := g.dijkstra(node1)[node2]
	if ans < math.MaxInt {
		return ans
	}
	return -1
}

// Dijkstra 邻接矩阵版本
// 返回从 start 出发，到各个点的最短路，如果不存在则为无穷大
func (g Graph) dijkstra(start int) []int {
	n := len(g)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		// 找到当前最短路，去更新它的邻居的最短路，根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == math.MaxInt { // 所有从 start 能到达的点都被更新了
			return dis
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			if w < math.MaxInt && dis[x]+w < dis[y] {
				dis[y] = dis[x] + w // 更新最短路长度
			}
		}
	}
}
