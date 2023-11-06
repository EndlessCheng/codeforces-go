package main

import "math"

// https://space.bilibili.com/206214
const inf = math.MaxInt / 3 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	d := make([][]int, n) // 邻接矩阵
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			if j != i {
				d[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
			}
		}
	}
	for _, e := range edges {
		d[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边和自环）
	}
	for k := range d {
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	return d
}

func (d Graph) AddEdge(e []int) {
	x, y, w := e[0], e[1], e[2]
	if w >= d[x][y] {
		return
	}
	for i := range d {
		for j := range d {
			d[i][j] = min(d[i][j], d[i][x]+w+d[y][j])
		}
	}
}

func (d Graph) ShortestPath(start, end int) int {
	ans := d[start][end]
	if ans == inf {
		return -1
	}
	return ans
}
