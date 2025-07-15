package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
func maxLen(n int, edges [][]int, label string) (ans int) {
	// 计算理论最大值
	cnt := [26]int{}
	for _, ch := range label {
		cnt[ch-'a']++
	}
	odd := 0
	for _, c := range cnt {
		odd += c % 2
	}
	theoreticalMax := n - max(odd-1, 0) // 奇数选一个放正中心，其余全弃

	if len(edges) == n*(n-1)/2 { // 完全图，可以达到理论最大值
		return theoreticalMax
	}

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
					// 保证 v < w，减少状态个数和计算量
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
		if ans == theoreticalMax {
			return
		}
		// 偶回文串，x 和 x 的邻居 y 作为回文中心
		for _, y := range to {
			// 保证 x < y，减少状态个数和计算量
			if x < y && label[x] == label[y] {
				ans = max(ans, dfs(x, y, 1<<x|1<<y)+2)
				if ans == theoreticalMax {
					return
				}
			}
		}
	}
	return
}

func maxLenGroupByLabel(n int, edges [][]int, label string) (ans int) {
	g := make([][]int, n)
	labelG := make([][26][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		labelG[x][label[y]-'a'] = append(labelG[x][label[y]-'a'], y)
		labelG[y][label[x]-'a'] = append(labelG[y][label[x]-'a'], x)
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
			for _, w := range labelG[y][label[v]-'a'] {
				if vis>>w&1 == 0 && w != v {
					// 保证 v < w，减少状态个数和计算量
					r := dfs(min(v, w), max(v, w), vis|1<<v|1<<w)
					res = max(res, r+2)
				}
			}
		}
		*p = res // 记忆化
		return
	}

	for x, to := range labelG {
		// 奇回文串，x 作为回文中心
		ans = max(ans, dfs(x, x, 1<<x)+1)
		if ans == n {
			return
		}
		// 偶回文串，x 和 x 的邻居 y 作为回文中心
		for _, y := range to[label[x]-'a'] {
			// 保证 x < y，减少状态个数和计算量
			if x < y {
				ans = max(ans, dfs(x, y, 1<<x|1<<y)+2)
				if ans == n {
					return
				}
			}
		}
	}
	return
}

func maxLenBFS(n int, edges [][]int, label string) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([][][]bool, n)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, 1<<n)
		}
	}
	type tuple struct{ x, y, vis int }
	q := []tuple{}
	// 奇回文串，x 作为回文中心
	for x := range n {
		vis[x][x][1<<x] = true
		q = append(q, tuple{x, x, 1 << x})
	}
	// 偶回文串，x 和 x 的邻居 y 作为回文中心
	for x, to := range g {
		for _, y := range to {
			// 保证 x < y，减少状态个数
			if x < y && label[x] == label[y] {
				vis[x][y][1<<x|1<<y] = true
				q = append(q, tuple{x, y, 1<<x | 1<<y})
			}
		}
	}
	var t tuple
	for len(q) > 0 {
		t = q[0]
		q = q[1:]
		for _, v := range g[t.x] {
			if t.vis>>v&1 > 0 { // v 在路径中
				continue
			}
			for _, w := range g[t.y] {
				if t.vis>>w&1 == 0 && w != v && label[w] == label[v] {
					// 保证 v < w，减少状态个数
					p := &vis[min(v, w)][max(v, w)][t.vis|1<<v|1<<w]
					if !*p {
						*p = true
						q = append(q, tuple{min(v, w), max(v, w), t.vis | 1<<v | 1<<w})
					}
				}
			}
		}
	}
	return bits.OnesCount(uint(t.vis))
}
