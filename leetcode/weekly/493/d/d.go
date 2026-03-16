package main

// https://space.bilibili.com/206214
func maxActivated(points [][]int) int {
	// 哈希表并查集
	fa := map[int]int{}
	var find func(int) int
	find = func(x int) int {
		fx, ok := fa[x]
		if !ok {
			return x
		}
		if fx != x {
			fa[x] = find(fx)
			return fa[x]
		}
		return x
	}

	const offset int = 3e9
	for _, p := range points {
		fa[find(p[0])] = find(p[1] + offset)
	}

	size := map[int]int{}
	for _, p := range points {
		size[find(p[0])]++ // 统计连通块的大小
	}

	mx1, mx2 := 0, 0
	for _, sz := range size {
		if sz > mx1 {
			mx2 = mx1
			mx1 = sz
		} else if sz > mx2 {
			mx2 = sz
		}
	}
	return mx1 + mx2 + 1
}
