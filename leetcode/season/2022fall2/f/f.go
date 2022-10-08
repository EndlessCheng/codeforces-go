package main

// https://space.bilibili.com/206214
func reservoir(shape []string) int {
	n, m := len(shape), len(shape[0])
	// 每个格子分成四个区域（上下左右），标上序号，方便用并查集连通
	// 假设左右下还有一圈格子，直接连到超级汇点 0
	u := make([][]int, n+1)
	d := make([][]int, n+1)
	l := make([][]int, n+1)
	r := make([][]int, n+1)
	for i := range u {
		u[i] = make([]int, m+2)
		d[i] = make([]int, m+2)
		l[i] = make([]int, m+2)
		r[i] = make([]int, m+2)
	}
	c := 1
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ { // 假设格子的列号从 1 开始，这样方便表示左右边界
			u[i][j] = c; c++
			d[i][j] = c; c++
			l[i][j] = c; c++
			r[i][j] = c; c++
		}
	}

	// 并查集模板
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) { fa[find(x)] = find(y) }

	ok := make([]bool, c) // 能否容纳水
	// 倒着判断每一行，寻找可能有水的区域
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= m; j++ {
			merge(r[i][j], l[i][j+1]) // 连通左右
		}
		for j := 1; j <= m; j++ {
			merge(d[i][j], u[i+1][j]) // 连通下
			// 根据格子的类型连接格子内部四个区域
			switch shape[i][j-1] {
			case '.':
				merge(l[i][j], u[i][j])
				merge(l[i][j], d[i][j])
				merge(l[i][j], r[i][j])
			case 'l':
				merge(l[i][j], d[i][j])
				merge(r[i][j], u[i][j])
			default:
				merge(l[i][j], u[i][j])
				merge(r[i][j], d[i][j])
			}
		}
		for j := 1; j <= m; j++ {
			// 在没有连接第 i-1 行的情况下，无法到达左右下边界 => 能容纳水
			ok[l[i][j]] = find(l[i][j]) != find(0)
			ok[r[i][j]] = find(r[i][j]) != find(0)
			ok[u[i][j]] = find(u[i][j]) != find(0)
			ok[d[i][j]] = find(d[i][j]) != find(0)
		}
	}

	// 第一行连上超级汇点，方便后面统一判断是否在闭合区域里面了
	for j := 1; j <= m; j++ {
		merge(u[0][j], 0)
	}

	ans := 0
	for i, b := range ok {
		if b && find(i) == find(0) {
			ans++
		}
	}
	return ans / 2
}
