package main

import (
	"fmt"
	"strings"
)

// https://space.bilibili.com/206214
func maxLen(n int, edges [][]int, label string) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, 1<<n)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}

	// 计算从 x 和 y 向两侧扩展，最多还能访问多少个节点（不算 x 和 y）
	var dfs func(int, int, int) int
	dfs = func(x, y, vis int) (res int) {
		p := &memo[x][y][vis]
		if *p >= 0 { // 之前计算过
			return *p
		}
		for _, v := range g[x] {
			if vis>>v&1 > 0 { // v 在路径中
				continue
			}
			for _, w := range g[y] {
				if vis>>w&1 == 0 && w != v && label[w] == label[v] {
					// 保证 v < w，减少状态个数
					r := dfs(min(v, w), max(v, w), vis|1<<v|1<<w)
					res = max(res, r+2)
				}
			}
		}
		*p = res // 记忆化
		return
	}

	for x, to := range g {
		// 奇回文串，x 作为回文中心
		ans = max(ans, dfs(x, x, 1<<x)+1)
		// 偶回文串，x 和 x 的邻居 y 作为回文中心
		for _, y := range to {
			// 保证 x < y，减少状态个数
			if x < y && label[x] == label[y] {
				ans = max(ans, dfs(x, y, 1<<x|1<<y)+2)
			}
		}
	}
	return
}

func main() {
	n := 14
	es := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			es = append(es, []int{i, j})
		}
	}
	fmt.Println(maxLen(n, es, strings.Repeat("a", n)))
}
