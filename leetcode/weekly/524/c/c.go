package main

// https://space.bilibili.com/206214
var mp = [...]int{'W': 0, 'R': 1, 'B': 2}

func findColoredRoot2(edges [][]int, colors string) int {
	n := len(colors)
	cnts := make([][3]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		cnts[x][mp[colors[y]]]++
		cnts[y][mp[colors[x]]]++
	}

	ans := -1
	for x, cnt := range cnts {
		sz := cnt[0] + cnt[1] + cnt[2]
		colorX := mp[colors[x]]
		if sz > 3 || cnt[colorX] > 0 || cnt[(colorX+2)%3] > 1 {
			return -1
		}

		if sz < 3 && cnt[(colorX+2)%3] == 0 {
			ans = x
		}
	}

	return ans
}

func findColoredRoot1(edges [][]int, colors string) int {
	n := len(colors)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := -1
	for x, to := range g {
		// 一个点至多三个邻居（父节点和左右儿子）
		if len(to) > 3 {
			return -1
		}

		cnt := [3]int{}
		for _, y := range to {
			cnt[mp[colors[y]]]++
		}

		colorX := mp[colors[x]]
		// 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
		if cnt[colorX] > 0 || cnt[(colorX+2)%3] > 1 {
			return -1
		}

		// 判断 x 能否是根
		if len(to) < 3 && cnt[(colorX+2)%3] == 0 {
			ans = x
		}
	}

	return ans
}

// 返回 0 到 n 的异或和
func xorN(n int) int {
	switch n % 4 {
	case 0: return n
	case 1: return 1
	case 2: return n + 1
	default: return 0
	}
}

var childColor = [...]byte{'R': 'B', 'B': 'W', 'W': 'R'}

func findColoredRoot(edges [][]int, colors string) int {
	n := len(colors)
	inDeg := make([]int8, n)
	outDeg := make([]int8, n)
	root := xorN(n - 1)

	for _, e := range edges {
		x, y := e[0], e[1]
		cx, cy := colors[x], colors[y]
		if cx == cy {
			return -1
		}

		// 确定谁是谁的父节点
		if childColor[cx] != cy {
			x, y = y, x
		}

		// x 是 y 的父节点
		outDeg[x]++
		inDeg[y]++
		if outDeg[x] > 2 || inDeg[y] > 1 {
			return -1
		}

		root ^= y
	}

	return root
}
