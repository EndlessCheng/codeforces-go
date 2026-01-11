package main

import (
	"slices"
	"sort"
)

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

func maxPartitionFactor1(points [][]int) int {
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

//

type unionFind struct {
	fa  []int
	dis []int8 // dis[x] 表示 x 到其代表元的距离
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	dis := make([]int8, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, dis}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		rt := u.find(u.fa[x])
		u.dis[x] ^= u.dis[u.fa[x]] // 更新 x 到其代表元的距离
		u.fa[x] = rt
	}
	return u.fa[x]
}

// 合并两个互斥的点
// 如果已经合并，返回是否与已知条件矛盾
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不合并
		return u.dis[from] != u.dis[to] // 是否与已知信息矛盾
	}
	//    2 ------ 4
	//   /        /
	//  1 ------ 3
	// 如果知道 1->2 的距离和 3->4 的距离，现在合并 1 和 3，并传入 1->3 的距离（本题等于 1）
	// 由于 1->3->4 和 1->2->4 的距离相等
	// 所以 2->4 的距离为 (1->3) + (3->4) - (1->2)
	u.dis[x] = 1 ^ u.dis[to] ^ u.dis[from]
	u.fa[x] = y
	return true
}

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	type tuple struct{ dis, x, y int }
	manhattanTuples := make([]tuple, 0, n*(n-1)/2) // 预分配空间
	for i, p := range points {
		for j, q := range points[:i] {
			manhattanTuples = append(manhattanTuples, tuple{abs(p[0]-q[0]) + abs(p[1]-q[1]), i, j})
		}
	}
	slices.SortFunc(manhattanTuples, func(a, b tuple) int { return a.dis - b.dis })

	uf := newUnionFind(n)
	for _, t := range manhattanTuples {
		if !uf.merge(t.x, t.y) {
			return t.dis // t.x 和 t.y 必须在同一个集合，t.dis 就是这一划分的最小划分因子
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
