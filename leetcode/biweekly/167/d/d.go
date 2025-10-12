package main

import "sort"

// https://space.bilibili.com/206214

// 原理见 785. 判断二分图
func isBipartite(points [][]int, low int) bool {
	colors := make([]int8, len(points))

	var dfs func(int, int8) bool
	dfs = func(x int, c int8) bool {
		colors[x] = c
		p := points[x]
		for y, q := range points {
			if y == x || abs(p[0]-q[0])+abs(p[1]-q[1]) >= low { // 符合要求
				continue
			}
			if colors[y] == c || colors[y] == 0 && !dfs(y, -c) {
				return false // 不是二分图
			}
		}
		return true
	}

	// 可能有多个连通块
	for i, c := range colors {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	if n == 2 {
		return 0
	}

	// 不想算的话可以写 maxDis = int(2e8)
	maxDis := 0
	for i, p := range points {
		for _, q := range points[:i] {
			maxDis = max(maxDis, abs(p[0]-q[0])+abs(p[1]-q[1]))
		}
	}

	return sort.Search(maxDis, func(low int) bool {
		// 二分最小的不满足要求的 low+1，就可以得到最大的满足要求的 low
		return !isBipartite(points, low+1)
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
