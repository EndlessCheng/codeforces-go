package main

// https://space.bilibili.com/206214
var mp = [...]int{'W': 0, 'R': 1, 'B': 2}

func findColoredRoot(edges [][]int, colors string) int {
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
			if ans >= 0 {
				return -1
			}
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

		// x 可以是根
		if len(to) < 3 && cnt[(colorX+2)%3] == 0 {
			// 不能有多个点能是根，那样当一个点是根的时候，另一个点不满足要求
			if ans >= 0 {
				return -1
			}
			ans = x
		}
	}

	return ans
}
