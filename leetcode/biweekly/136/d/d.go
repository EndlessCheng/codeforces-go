package main

// https://space.bilibili.com/206214
func timeTaken(edges [][]int) []int {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	nodes := make([]struct{ maxD, maxD2, y int }, len(g))
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		p := &nodes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			maxD := dfs(y, x) + 2 - y%2 // 从 x 出发，往 y 方向的最大深度
			if maxD > p.maxD {
				p.maxD2 = p.maxD
				p.maxD = maxD
				p.y = y
			} else if maxD > p.maxD2 {
				p.maxD2 = maxD
			}
		}
		return p.maxD
	}
	dfs(0, -1)

	ans := make([]int, len(g))
	var reroot func(int, int, int)
	reroot = func(x, fa, fromUp int) {
		p := nodes[x]
		ans[x] = max(fromUp, p.maxD)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			w := 2 - x%2
			if y == p.y { // 对于 y 来说，上面要选次大的
				reroot(y, x, max(fromUp, p.maxD2)+w)
			} else { // 对于 y 来说，上面要选最大的
				reroot(y, x, max(fromUp, p.maxD)+w)
			}
		}
	}
	reroot(0, -1, 0)
	return ans
}
