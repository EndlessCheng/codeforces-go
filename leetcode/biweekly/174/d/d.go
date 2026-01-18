package main

import "slices"

// https://space.bilibili.com/206214
func minimumFlips(n int, edges [][]int, start, target string) (ans []int) {
	type edge struct{ to, i int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, i})
		g[y] = append(g[y], edge{x, i})
	}

	// 返回是否需要翻转 x-fa 这条边
	var dfs func(int, int) bool
	dfs = func(x, fa int) bool {
		rev := start[x] != target[x] // x-fa 是否要翻转
		for _, e := range g[x] {
			y := e.to
			if y != fa && dfs(y, x) {
				ans = append(ans, e.i) // 需要翻转 y-x
				rev = !rev             // x 被翻转了
			}
		}
		return rev
	}

	if dfs(0, -1) { // 只剩下一个根节点需要翻转，无法操作
		return []int{-1}
	}
	slices.Sort(ans)
	return
}
