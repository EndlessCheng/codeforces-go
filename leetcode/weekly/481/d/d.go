package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func interactionCosts1(n int, edges [][]int, group []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mx := slices.Max(group)
	total := make([]int, mx+1)
	for _, x := range group {
		total[x]++
	}

	var dfs func(int, int) []int
	dfs = func(x, fa int) []int {
		cntX := make([]int, mx+1)
		cntX[group[x]] = 1
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			cntY := dfs(y, x)
			for i, c := range cntY {
				ans += int64(c) * int64(total[i]-c)
				cntX[i] += c
			}
		}
		return cntX
	}
	dfs(0, -1)
	return
}

func interactionCosts(n int, edges [][]int, group []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mx := slices.Max(group)
	total := make([]int, mx+1)
	for _, x := range group {
		total[x]++
	}

	for target, tot := range total {
		if tot == 0 {
			continue
		}
		var dfs func(int, int) int
		dfs = func(x, fa int) (cntX int) {
			if group[x] == target {
				cntX = 1
			}
			for _, y := range g[x] {
				if y == fa {
					continue
				}
				cntY := dfs(y, x)
				ans += int64(cntY) * int64(tot-cntY)
				cntX += cntY
			}
			return
		}
		dfs(0, -1)
	}
	return
}

func interactionCosts3(n int, edges [][]int, group []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dfn := make([]int, n)
	ts := 0
	pa := make([][17]int, n)
	dep := make([]int, n)
	var build func(int, int)
	build = func(v, p int) {
		dfn[v] = ts
		ts++
		pa[v][0] = p
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				build(w, v)
			}
		}
	}
	build(0, -1)
	mx := bits.Len(uint(n))
	for i := range mx - 1 {
		for v := range pa {
			p := pa[v][i]
			if p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros32(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			pv, pw := pa[v][i], pa[w][i]
			if pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	nodesMap := map[int][]int{}
	for i, x := range group {
		nodesMap[x] = append(nodesMap[x], i)
	}

	vt := make([][]int, n)   // 虚树
	isNode := make([]int, n) // 用来区分是关键节点还是 LCA
	for i := range isNode {
		isNode[i] = -1
	}
	addVtEdge := func(v, w int) {
		vt[v] = append(vt[v], w) // 往虚树上添加一条有向边
	}
	const root = 0
	st := []int{root} // 用根节点作为栈底哨兵

	for val, nodes := range nodesMap {
		// 对于相同点权的这一组关键节点 nodes，构建虚树
		slices.SortFunc(nodes, func(a, b int) int { return dfn[a] - dfn[b] })
		vt[root] = vt[root][:0] // 重置虚树
		st = st[:1]
		for _, v := range nodes {
			isNode[v] = val
			if v == root {
				continue
			}
			vt[v] = vt[v][:0]
			lca := getLCA(st[len(st)-1], v) // 路径的拐点（LCA）也加到虚树中
			// 回溯，加边
			for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
				addVtEdge(st[len(st)-2], st[len(st)-1])
				st = st[:len(st)-1]
			}
			if lca != st[len(st)-1] { // lca 不在栈中（首次遇到）
				vt[lca] = vt[lca][:0]
				addVtEdge(lca, st[len(st)-1])
				st[len(st)-1] = lca // 加到栈中
			}
			st = append(st, v)
		}
		// 最后的回溯，加边
		for i := 1; i < len(st); i++ {
			addVtEdge(st[i-1], st[i])
		}

		var dfs func(int) int
		dfs = func(v int) (size int) {
			// 如果 isNode[v] != t，那么 v 只是关键节点之间路径上的「拐点」
			if isNode[v] == val {
				size = 1
			}
			for _, w := range vt[v] {
				sz := dfs(w)
				wt := dep[w] - dep[v] // 虚树边权
				// 贡献法
				ans += int64(wt) * int64(sz) * int64(len(nodes)-sz)
				size += sz
			}
			return
		}

		rt := root
		if isNode[rt] != val && len(vt[rt]) == 1 {
			// 注意 root 只是一个哨兵，不一定在虚树上，得从真正的根节点开始
			rt = vt[rt][0]
		}
		dfs(rt)
	}

	return
}
