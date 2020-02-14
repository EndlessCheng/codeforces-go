package copypasta

import (
	. "container/heap"
	. "fmt"
	"io"
	"sort"
)

// namespace
type graph struct{}

// deg[v] == len(g[v])
func (*graph) readGraph(in io.Reader, n, m int) [][]int {
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	return g
}

// 链式前向星
// https://oi-wiki.org//graph/save/#_14
func (*graph) readGraphList(in io.Reader, n, m int) {
	type edge struct{ to, prev int }
	edgeID := make([]int, n)
	for i := range edgeID {
		edgeID[i] = -1
	}
	edges := make([]edge, m) // 2*m
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		edges[i] = edge{w, edgeID[v]}
		edgeID[v] = i
	}

	// loop all edges start at v
	var v int
	for i := edgeID[v]; i != -1; {
		e := edges[i]
		w := e.to
		_ = w // do(w)
		i = e.prev
	}
}

func (*graph) readWeightedGraph(in io.Reader, n, m int) {
	type neighbor struct{ to, weight int }
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w, weight int
		Fscan(in, &v, &w, &weight)
		v--
		w--
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}
}

func (*graph) simpleSearch(n, st int, g [][]int) {
	// DFS
	vis := make([]bool, n)
	var f func(int)
	f = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	f(st)

	// BFS
	vis = make([]bool, n)
	vis[st] = true
	q := []int{st}
	for len(q) > 0 {
		var v int
		v, q = q[0], q[1:]
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				q = append(q, w)
			}
		}
	}
}

func (*graph) bfsWithDepth(n, st int, g [][]int, do func(v, dep int)) {
	vis := make([]bool, n)
	vis[st] = true
	type pair struct{ v, dep int }
	queue := []pair{{st, 0}}
	for len(queue) > 0 {
		var p pair
		p, queue = queue[0], queue[1:]
		do(p.v, p.dep)
		for _, w := range g[p.v] {
			if !vis[w] {
				vis[w] = true
				queue = append(queue, pair{w, p.dep + 1})
			}
		}
	}
}

func (*graph) depthArray(n, st int, g [][]int) []int {
	depths := make([]int, n)
	vis := make([]bool, n)
	vis[st] = true
	queue := []int{st}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				depths[w] = depths[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return depths
}

// 标记所属连通块
func (*graph) markComponentIDs(n int, g [][]int) []int {
	id := make([]int, n)
	for i := range id {
		id[i] = -1
	}
	idCnt := 0
	var f func(int)
	f = func(v int) {
		id[v] = idCnt
		for _, w := range g[v] {
			if id[w] == 0 {
				f(w)
			}
		}
	}
	for i, idi := range id {
		if idi == 0 {
			f(i)
			idCnt++
		}
	}
	return id
}

// 求出所有连通块
func (*graph) getAllComponents(n int, g [][]int) [][]int {
	comps := [][]int{}
	vis := make([]bool, n)
	var comp []int
	var f func(int)
	f = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i, vi := range vis {
		if !vi {
			comp = []int{}
			f(i)
			comps = append(comps, comp)
		}
	}
	return comps
}

// 割点 cut vertices / articulation points
// https://codeforces.com/blog/entry/68138
// https://oi-wiki.org/graph/cut/#_1
// low(v): 在不经过 v 父亲的前提下能到达的最小的时间戳
func (*graph) findCutVertices(n int, g [][]int) (isCut []bool) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	isCut = make([]bool, n)
	dfsClock := 0
	dfn := make([]int, n) // 值从 1 开始
	//low := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			if dfn[w] == 0 {
				childCnt++
				lowW := f(w, v)
				if lowW >= dfn[v] { // 该子树没有反向边能连回 v 的祖先
					isCut[v] = true
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] { // 找到反向边，用来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 { // 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数
			isCut[v] = false
		}
		//low[v] = lowV
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			f(v, -1)
		}
	}
	vs := make([]int, 0, n)
	for v, is := range isCut {
		if is {
			vs = append(vs, v) // v+1
		}
	}
	return
}

// 桥
// https://oi-wiki.org/graph/cut/#_4
// 题目推荐 https://cp-algorithms.com/graph/bridge-searching.html#toc-tgt-2
func (*graph) findBridges(n, m int, g [][]int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type bridge struct{ v, w int }
	bridges := make([]bridge, 0, m)
	dfsClock := 0
	dfn := make([]int, n) // 值从 1 开始
	//low := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			if dfn[w] == 0 {
				childCnt++
				lowW := f(w, v)
				if lowW > dfn[v] { // w 子树的反向边只能连回其自身，所以 v-w 必定是桥
					bridges = append(bridges, bridge{v, w})
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] { // 找到反向边，用来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		//low[v] = lowV
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			f(v, -1)
		}
	}
	// bridges...
}

// 无向图的双连通分量 biconnected components
// https://oi-wiki.org/graph/bcc/
// https://www.csie.ntu.edu.tw/~hsinmu/courses/_media/dsa_13spring/horowitz_306_311_biconnected.pdf
func (*graph) findBCC() (comps [][]int, bccIDs []int) {
	// TODO: implement！！
	// 點雙連通、邊雙連通
	return
}

// 传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/shortest-path/#floyd
// 题目推荐 https://cp-algorithms.com/graph/all-pair-shortest-path-floyd-warshall.html#toc-tgt-5
func (*graph) shortestPathFloydWarshall(in io.Reader, n, m int) [][]int {
	min := func(a, b int) int { // int64
		if a < b {
			return a
		}
		return b
	}
	const inf int = 1e9
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}
	for i := 0; i < m; i++ {
		var v, w, dis int
		Fscan(in, &v, &m, &dis)
		dist[v][w] = dis
		dist[w][v] = dis
	}
	for k := range dist { // 阶段
		for i := range dist { // 状态
			for j := range dist { // 决策
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return dist
}

// 传入邻接矩阵 weights
// weights[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/min-circle/#floyd
func (*graph) shortestCycleFloydWarshall(weights [][]int64) int64 {
	min := func(a, b int64) int64 {
		if a <= b {
			return a
		}
		return b
	}
	const inf int64 = 1e18
	//const inf int = 1e8 // *NOTE*
	n := len(weights)
	dist := make([][]int64, n)
	for i := range dist {
		dist[i] = make([]int64, n)
		copy(dist[i], weights[i])
	}
	ans := inf
	for k := range dist { // 阶段
		for i := 0; i < k; i++ { // 状态
			for j := 0; j < i; j++ { // 决策
				ans = min(ans, dist[i][j]+weights[i][k]+weights[k][j])
			}
		}
		for i := range dist { // 状态
			for j := range dist { // 决策
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return ans
}

// 适用于稀疏图 O((|E|+|V|)⋅log|V|)
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 题目推荐 https://cp-algorithms.com/graph/dijkstra.html#toc-tgt-5
func (*graph) shortestPathDijkstra(in io.Reader, n, m, start int) (dist []int64, parents []int) {
	type neighbor struct {
		to     int
		weight int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w int
		var weight int64
		Fscan(in, &v, &w, &weight)
		v--
		w--
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	const inf int64 = 1e18 // 1e9+1
	dist = make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	parents = make([]int, n)
	for i := range parents {
		parents[i] = -1
	}

	h := &pairHeap{}
	Push(h, hPair{0, start})
	for h.Len() > 0 {
		p := Pop(h).(hPair)
		d, v := p.x, p.y
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.weight; newD < dist[w] {
				dist[w] = newD
				parents[w] = v
				Push(h, hPair{newD, w})
			}
		}
	}

	// path from end to start
	var end = n - 1
	path := make([]int, 0, n)
	for v := end; v != -1; v = parents[v] {
		path = append(path, v)
	}

	return
}

// https://oi-wiki.org/graph/bfs/#bfs_3
// https://codeforces.com/blog/entry/22276
// 例题: https://codeforces.com/problemset/problem/173/B
func (*graph) bfs01(in io.Reader, n, m, st int) []int {
	type neighbor struct{ to, weight int }
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w, weight int
		Fscan(in, &v, &w, &weight)
		v--
		w--
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	const inf int = 1e9
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	q := &deque{}
	q.pushL(st)
	for !q.empty() {
		v := q.popL()
		for _, e := range g[v] {
			w, d := e.to, e.weight
			if dist[v]+d < dist[w] {
				dist[w] = dist[v] + d
				if d == 0 {
					q.pushL(w)
				} else {
					q.pushR(w)
				}
			}
		}
	}
	return dist
}

// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
func (*graph) shortestPathBellmanFord(in io.Reader, n, m, start int) (dist []int64) {
	type neighbor struct {
		to     int
		weight int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w int
		var weight int64
		Fscan(in, &v, &w, &weight)
		v--
		w--
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	const inf int64 = 1e18 // 1e9+1
	dist = make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	onQ := make([]bool, n)
	onQ[start] = true
	relaxedCnt := make([]int, n)

	queue := []int{start}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		onQ[v] = false
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.weight; newD < dist[w] {
				dist[w] = newD
				relaxedCnt[w]++
				if relaxedCnt[w] > n {
					// found negative cycle
					return nil
				}
				// there is no reason to put a vertex in the queue if it is already in.
				if !onQ[w] {
					queue = append(queue, w)
					onQ[w] = true
				}
			}
		}
	}
	return
}

// https://cp-algorithms.com/graph/finding-negative-cycle-in-graph.html
func (*graph) findNegativeCycleBellmanFord() []int {
	// TODO return negative cycle
	return nil
}

// 适用于稀疏图 O(|E|⋅log|E|)
// https://oi-wiki.org/graph/mst/#kruskal
// 题目推荐 https://cp-algorithms.com/graph/mst_kruskal.html#toc-tgt-5
func (*graph) mstKruskal(in io.Reader, n, m int) (sum int64) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	type edge struct {
		v, w   int
		weight int
	}
	edges := make([]edge, m)
	for i := range edges {
		var v, w, weight int
		Fscan(in, &v, &w, &weight)
		v--
		w--
		edges[i] = edge{v, w, weight}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
	initFa(n)
	for _, e := range edges {
		if from, to := find(e.v), find(e.w); from != to {
			sum += int64(e.weight)
			fa[from] = to
		}
	}
	return
}

// 适用于稠密图 O(|V|^2)，传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/mst/#prim
func (*graph) mstPrim(dist [][]int) (sum int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const inf = 1.1e9
	n := len(dist)
	minWeights := make([]int, n)
	for i := range minWeights {
		minWeights[i] = inf
	}
	minWeights[0] = 0
	used := make([]bool, n)
	for {
		v := -1
		for i, use := range used {
			if !use && (v == -1 || minWeights[i] < minWeights[v]) {
				v = i
			}
		}
		if v == -1 {
			break
		}
		used[v] = true
		sum += minWeights[v]
		for w := range minWeights {
			minWeights[w] = min(minWeights[w], dist[v][w])
		}
	}
	return
}

// Second best Minimum Spanning Tree
// Using Kruskal and Lowest Common Ancestor
// https://cp-algorithms.com/graph/second_best_mst.html
func (*graph) secondMST(n, m int) (sum int64) {
	// TODO implement
	return
}

// 反图的连通分量 O(n+m)
// https://www.luogu.com.cn/blog/endlesscheng/solution-cf1242b
func (*graph) inverseGraphComponents(n int, g [][]int) (components [][]int) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	maxDegInv, maxDegInvV := 0, 0
	for v, edges := range g {
		if degInv := n - 1 - len(edges); degInv > maxDegInv {
			maxDegInv = degInv
			maxDegInvV = v
		}
	}
	if maxDegInv == 0 {
		components = make([][]int, n)
		for i := range components {
			components[i] = []int{i} // i+1
		}
		return
	}

	mergeInv := func(v int, edges []int) {
		vs := map[int]bool{v: true}
		for _, w := range edges {
			vs[w] = true
		}
		for i := 0; i < n; i++ {
			if !vs[i] {
				merge(i, v)
			}
		}
	}
	initFa(n)
	mergeInv(maxDegInvV, g[maxDegInvV])
	for v, edges := range g {
		if !same(v, maxDegInvV) {
			mergeInv(v, edges)
		}
	}

	componentsMap := map[int][]int{}
	for i := range fa {
		rootV := find(i)
		componentsMap[rootV] = append(componentsMap[rootV], i) // i+1
	}
	components = make([][]int, 0, len(componentsMap))
	for _, vs := range componentsMap {
		components = append(components, vs)
	}
	return
}

// 二分图判定
// https://oi-wiki.org/graph/bi-graph/#_3
// https://cp-algorithms.com/graph/bipartite-check.html
func (*graph) isBipartite(n int, g [][]int) bool {
	colors := make([]int8, n)
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			if colors[w] == c {
				return false
			}
			if colors[w] == 0 {
				if !f(w, 3-c) {
					return false
				}
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 && !f(i, 1) {
			return false
		}
	}
	return true
}

// 二分图最大匹配 - 匈牙利算法/增广路算法 O(nm)
// https://www.renfei.org/blog/bipartite-matching.html 推荐
// https://www.geeksforgeeks.org/maximum-bipartite-matching/
// https://oi-wiki.org/graph/bi-graph/#_9
// https://zhuanlan.zhihu.com/p/62981901
func (*graph) maxMatchingHungarian(n int, g [][]int) (match []int, cnt int) {
	match = make([]int, n)
	for i := range match {
		match[i] = -1
	}
	var used []bool
	var f func(v int) bool
	f = func(v int) bool {
		used[v] = true
		for _, w := range g[v] {
			if mw := match[w]; mw == -1 || !used[mw] && f(mw) {
				match[w] = v
				match[v] = w
				return true
			}
		}
		return false
	}
	for v := range g {
		if match[v] == -1 {
			used = make([]bool, n)
			if f(v) {
				cnt++
			}
		}
	}
	return
}

// 另一种写法，适用左右两侧节点有明确区分的情况，要求 g 中存储的是左侧到右侧的单向边
func (*graph) maxMatchingHungarianLR(nl, nr int, g [][]int) (matchL []int, cnt int) {
	matchL = make([]int, nl)
	matchR := make([]int, nr)
	for i := range matchL {
		matchL[i] = -1
	}
	for i := range matchR {
		matchR[i] = -1
	}
	var used []bool
	var f func(v int) bool
	f = func(v int) bool {
		used[v] = true
		for _, w := range g[v] {
			if lv := matchR[w]; lv == -1 || !used[lv] && f(lv) {
				matchR[w] = v
				matchL[v] = w
				return true
			}
		}
		return false
	}
	for v := range g {
		used = make([]bool, nl)
		if f(v) {
			cnt++
		}
	}
	return
}

// 二分图最大匹配 -  Hopcroft–Karp 算法 O(m√n)
// http://pepcy.cf/icpc-templates/003Graph/hk.html
// https://en.wikipedia.org/wiki/Hopcroft%E2%80%93Karp_algorithm
func (*graph) maxMatchingHopcroftKarp(n int, g [][]int) (match []int, cnt int) {
	// TODO
	return
}

// 带权二分图最大匹配 - Kuhn–Munkres 算法 O(n^3)
// https://zhuanlan.zhihu.com/p/62981901
func (*graph) maxMatchingKuhnMunkres(n int, g [][]int) (match []int, cnt int) {
	// TODO
	return
}

// 拓扑排序 Kahn's algorithm
// https://oi-wiki.org/graph/topo/
// https://cp-algorithms.com/graph/topological-sort.html
func (*graph) topSort(in io.Reader, n, m int) (orders []int, parents []int, levels []int, isDAG bool) {
	g := make([][]int, n)
	inDeg := make([]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w) // 注意不能有自环
		inDeg[w]++
	}

	// 拓扑排序生成的额外有用信息
	parents = make([]int, n)
	for i := range parents {
		parents[i] = -1
	}
	levels = make([]int, n)

	orders = make([]int, 0, n)
	q := []int{}
	for i, deg := range inDeg {
		if deg == 0 {
			q = append(q, i)
			levels[i] = 1
		}
	}
	for len(q) > 0 {
		var v int
		v, q = q[0], q[1:]
		orders = append(orders, v)
		for _, w := range g[v] {
			inDeg[w]--
			if inDeg[w] == 0 {
				q = append(q, w)
				levels[w] = levels[v] + 1
				parents[w] = v
			}
		}
	}

	isDAG = len(orders) == n

	// path from end to start
	var end = n - 1
	path := make([]int, 0, n)
	for v := end; v != -1; v = parents[v] {
		path = append(path, v)
	}

	return
}

// 强连通分量分解 Strongly connected component
// https://oi-wiki.org/graph/scc/#kosaraju
// https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
// sccIDs[v] 表示点 v 所属的 SCC 的拓扑序
func (*graph) sccKosaraju(in io.Reader, n, m int, g [][]int) (comps [][]int, sccIDs []int) {
	type edge struct{ v, w int }
	edges := make([]edge, m)
	g = make([][]int, n)
	rg := make([][]int, n)

	addEdge := func(v, w int) {
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		addEdge(v, w)
		edges[i] = edge{v, w}
	}

	used := make([]bool, n)
	vs := make([]int, 0, n)
	var dfs func(int)
	dfs = func(v int) {
		used[v] = true
		for _, w := range g[v] {
			if !used[w] {
				dfs(w)
			}
		}
		// 后序遍历
		vs = append(vs, v)
	}
	for v := range g {
		if !used[v] {
			dfs(v)
		}
	}

	used = make([]bool, n)
	var comp []int
	var rdfs func(int)
	rdfs = func(v int) {
		used[v] = true
		comp = append(comp, v)
		for _, w := range rg[v] {
			if !used[w] {
				rdfs(w)
			}
		}
	}
	comps = [][]int{}
	for i := len(vs) - 1; i >= 0; i-- {
		if v := vs[i]; !used[v] {
			comp = []int{}
			rdfs(v)
			comps = append(comps, comp)
		}
	}

	sccIDs = make([]int, n)
	for i, cp := range comps {
		for _, v := range cp {
			sccIDs[v] = i
		}
	}

	// EXTRA: 缩点: 将边 v-w 转换成 sccIDs[v]-sccIDs[w]
	// 缩点后的点的编号范围为 [0,len(comps)-1]
	// 注意这样会产生重边和自环
	for _, e := range edges {
		if v, w := sccIDs[e.v], sccIDs[e.w]; v != w {
			// custom
		}
	}

	// EXTRA: 求有多少个点能被其他所有点访问到
	lastComp := comps[len(comps)-1]
	numCanBeVisitedFromAll := len(lastComp)
	_ = numCanBeVisitedFromAll
	used = make([]bool, n)
	rdfs(lastComp[0])
	for _, use := range used {
		if !use {
			numCanBeVisitedFromAll = 0
		}
	}

	return
}

// TODO: SCC Tarjan

// https://oi-wiki.org/graph/2-sat/
// https://cp-algorithms.com/graph/2SAT.html
// https://zhuanlan.zhihu.com/p/50211772
// 下面的代码基于模板题 https://www.luogu.com.cn/problem/P4782
// 读入 m 条数据，每条数据表示 (x为a)∨(y为b)，a b 为 0 或 1
// ¬x 用 x+n 表示
// NOTE: 单独的条件 x为a 可以用 (x为a)∨(x为a) 来表示
// NOTE: 一些建边的转换：
//       A,B 至少存在一个 (A|B)    ¬A⇒B, ¬B⇒A
//       A,B 不能同时存在 (¬A|¬B)  A⇒¬B, B⇒¬A
//       A,B 必须且只一个 (A^B)    A⇒¬B, B⇒¬A, ¬A⇒B, ¬B⇒A
//       A,B 同时或都不在 (¬(A^B)) A⇒B, B⇒A, ¬A⇒¬B, ¬B⇒¬A
//       A 必须存在       (A)     ¬A⇒A
//       A 不能存在       (¬A)     A⇒¬A
func (G *graph) solve2SAT(in io.Reader, n, m int) []bool {
	g := make([][]int, 2*n)
	for i := 0; i < m; i++ {
		var x, a, y, b int
		Fscan(in, &x, &a, &y, &b)
		x--
		y--
		v, w := x+a&1*n, y+(b^1)*n
		g[v] = append(g[v], w)
		v, w = y+b&1*n, x+(a^1)*n
		g[v] = append(g[v], w)
	}
	_, sccIDs := G.sccKosaraju(in, 2*n, m, g) // *两倍空间*
	ans := make([]bool, n)
	for i, id := range sccIDs[:n] {
		if id == sccIDs[i+n] { // x ⇔ ¬x
			return nil
		}
		// sccIDs[x] > sccIDs[¬x] ⇔ (¬x ⇒ x) ⇔ x 为真
		// sccIDs[x] < sccIDs[¬x] ⇔ (x ⇒ ¬x) ⇔ x 为假
		ans[i] = id > sccIDs[i+n]
	}
	return ans
}
