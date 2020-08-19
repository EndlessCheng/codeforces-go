package copypasta

import (
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

/*
Graph Theory Playlist https://www.youtube.com/playlist?list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P
图论的小技巧以及扩展 https://www.luogu.com.cn/blog/chengni5673/tu-lun-di-xiao-ji-qiao-yi-ji-kuo-zhan

CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=graphs
构造 https://codeforces.com/problemset/problem/41/E

TIPS: 使用一个 fa 数组（初始化为 -1）记录搜索树中的节点的父节点，这样对每个节点都有一条到根的路径（根的 fa 为 -1）
NOTE: 独立集相关问题，可以从染色的角度考虑
NOTE: 度数大于 √M 的点不超过 2√M 个

环与独立集 https://codeforces.com/problemset/problem/1364/D
匹配与独立集 https://codeforces.com/problemset/problem/1198/C
*/

// namespace
type graph struct{}

// len(g[v]) 表示结点 v 在无向图上的度/有向图上的出度
// 对于树来说叶结点有 len(g[v]) == 1
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

func (*graph) dfs(n, st int, g [][]int) {
	vis := make([]bool, n)
	var f func(int)
	f = func(v int) {
		vis[v] = true
		// ...
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	f(st)
	// 有向图/森林
	for i, b := range vis {
		if !b {
			f(i)
		}
	}

	{
		// 有向图的环/回边检测/012染色
		//《算法导论》p.353 边的分类
		// vis[v] == 0：该顶点未被访问
		// vis[v] == 1：该顶点已经被访问，其子树未遍历完
		// vis[v] == 2：该顶点已经被访问，其子树已遍历完
		vis := make([]int8, n)
		var f func(int)
		f = func(v int) {
			vis[v] = 1
			for _, w := range g[v] {
				if t := vis[w]; t == 0 { // 树边
					f(w)
				} else if t == 1 { // 后向边，说明有环

				} else { // 前向边或横向边

				}
			}
			vis[v] = 2
		}
		for i, t := range vis {
			if t == 0 {
				f(i)
			}
		}
	}

	{
		// 无向图的树边和非树边（含重边）
		vis := make([]int8, n)
		var f func(int)
		f = func(v int) {
			vis[v] = 1
			for _, w := range g[v] {
				if t := vis[w]; t == 0 { // 树边
					f(w)
				} else if t == 2 || w == v { // 非树边（没有自环的话可以去掉 w == v）

				}
			}
			vis[v] = 2
		}
		for i, t := range vis {
			if t == 0 {
				f(i)
			}
		}
	}

	{
		// 无向图: DFS 找长度至少为 k 的环
		// https://codeforces.com/problemset/problem/263/D
		var k, end, st int
		fa := make([]int, n)
		dep := make([]int, n)
		var f func(v, p, d int) bool
		f = func(v, p, d int) bool {
			fa[v] = p
			dep[v] = d
			for _, w := range g[v] {
				if dep[w] == 0 {
					if f(w, v, d+1) {
						return true
					}
				} else if d-dep[w] >= k {
					end, st = v, w
					return true
				}
			}
			return false
		}
		f(st, -1, 1)

		cycle := []interface{}{st + 1} // for print
		for v := end; v != st; v = fa[v] {
			cycle = append(cycle, v+1)
		}
	}
}

// DFS 应用：求连通分量以及每个点所属的连通分量 (Connected Component, CC)
// ccIDs 的值从 1 开始
func (*graph) calcCC(n int, g [][]int) (comps [][]int, ccIDs []int) {
	ccIDs = make([]int, n)
	idCnt := 0 // 也可以去掉，用 len(comps)+1 代替
	var comp []int
	var f func(int)
	f = func(v int) {
		ccIDs[v] = idCnt
		comp = append(comp, v)
		for _, w := range g[v] {
			if ccIDs[w] == 0 {
				f(w)
			}
		}
	}
	for i, id := range ccIDs {
		if id == 0 {
			idCnt++
			comp = []int{}
			f(i)
			comps = append(comps, comp)
		}
	}
	return
}

func (*graph) bfs(n, st int, g [][]int) {
	vis := make([]bool, n)
	vis[st] = true
	q := []int{st}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// do v...
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				q = append(q, w)
			}
		}
	}

	{
		// 构建深度数组
		dep := make([]int, n)
		for i := range dep {
			dep[i] = -1
		}
		dep[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			// do(v, dep[v]) ...

			for _, w := range g[v] {
				if dep[w] == -1 {
					dep[w] = dep[v] + 1
					q = append(q, w)
				}
			}
		}
	}

	{
		// BFS with rollback
		vis = make([]bool, n)
		vis[st] = true
		vs := []int{st}
		type pair struct{ v, fa int }
		q := []pair{{st, -1}}
	outer:
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			v, fa := p.v, p.fa
			for _, w := range g[v] {
				if !vis[w] {
					vis[w] = true
					q = append(q, pair{w, v})
					vs = append(vs, w)
				} else if w != fa {
					// ... (兼容自环和重边)
					break outer // 提前退出的情况
				}
			}
		}
		for _, v := range vs {
			vis[v] = false
		}
	}

	{
		// BFS 012 染色
		// 0 不在队列，未访问
		// 1 在队列，未访问
		// 2 不在队列，已访问
		// 相关题目 https://codeforces.com/contest/1385/problem/E
		vis := make([]int8, n)
		vis[st] = 1
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			// do v...
			for _, w := range g[v] {
				if vis[w] == 0 {
					vis[w] = 1
					q = append(q, w)
				}
			}
			vis[v] = 2
		}
	}
}

// BFS 应用：求无向无权图最小环长度
// 好题 https://codeforces.com/problemset/problem/1325/E
func (*graph) shortestCycleBFS(n int, g [][]int) int {
	const inf int = 1e9
	ans := inf
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	type pair struct{ v, fa int }
	for st := range g {
		vs := []int{st}
		dist[st] = 0
		q := []pair{{st, -1}}
	outer:
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			v, fa := p.v, p.fa
			for _, w := range g[v] {
				if dist[w] == -1 {
					dist[w] = dist[v] + 1
					q = append(q, pair{w, v})
					vs = append(vs, w)
				} else if w != fa {
					if l := dist[w] + dist[v] + 1; l < ans {
						ans = l
					}
					break outer
				}
			}
		}
		for _, v := range vs {
			dist[v] = -1
		}
	}
	return ans
}

// 欧拉回路（欧拉图）：连通且每个点的度数为偶数；对于有向图需要入度和出度相同
// 欧拉路径：连通且恰好有两个点的度数为奇数；对于有向图需要这起点的出度比入度大一，终点的入度比出度大一
// 逐步插入回路法（Hierholzer 算法） https://oi-wiki.org/graph/euler/
// todo 混合图欧拉回路
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/EulerianCycle.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/EulerianPath.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/DirectedEulerianCycle.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/DirectedEulerianPath.java.html
// https://algs4.cs.princeton.edu/42digraph/DirectedEulerianCycle.java.html
// NOTE: 递归前对边排序可保证输出的是字典序最小的路径
// 模板题（输出顶点）
//		无向图 https://www.luogu.com.cn/problem/P2731
//      无向图 https://www.luogu.com.cn/problem/P1341
//      有向图 LC332 https://leetcode-cn.com/problems/reconstruct-itinerary/solution/javadfsjie-fa-by-pwrliang/
func (*graph) eulerianPath(n, m int) []int {
	type neighbor struct{ to, eid int }
	g := make([][]neighbor, n)
	// read g ...

	// 排序，保证字典序最小
	for _, vs := range g {
		sort.Slice(vs, func(i, j int) bool { return vs[i].to < vs[j].to })
	}

	var st int
	oddDegCnt := 0
	for i := len(g) - 1; i >= 0; i-- { // 倒着遍历保证起点的字典序最小
		//if len(g[i]) == 0 {
		//	continue
		//}
		if deg := len(g[i]); deg > 0 {
			if deg&1 == 1 {
				st = i
				oddDegCnt++
			} else if oddDegCnt == 0 {
				st = i
			}
		}
	}
	if oddDegCnt > 2 {
		return nil
	}

	// NOTE: 若没有奇度数，则返回的是欧拉回路
	path := make([]int, 0, m+1)
	{
		// 无向图
		vis := make([]bool, m)
		var f func(int)
		f = func(v int) {
			for len(g[v]) > 0 {
				e := g[v][0]
				g[v] = g[v][1:]
				if i := e.eid; !vis[i] {
					vis[i] = true
					f(e.to)
					// NOTE: 输出边的话移在这里 append i
				}
			}
			path = append(path, v)
		}
		f(st)
	}

	{
		// 有向图
		var f func(int)
		f = func(v int) {
			for len(g[v]) > 0 {
				e := g[v][0]
				g[v] = g[v][1:]
				f(e.to)
				// NOTE: 输出边的话移在这里 append e.eid
			}
			path = append(path, v)
		}
		f(st)
	}

	// 倒序输出 path
	for i, j := 0, len(path)-1; i < j; i++ {
		path[i], path[j] = path[j], path[i]
		j--
	}

	return path
}

/* Topic - DFS 树
讲解+套题 https://codeforces.com/blog/entry/68138
好题：https://codeforces.com/problemset/problem/1325/F
*/

// 割点（割顶） cut vertices / articulation points
// https://codeforces.com/blog/entry/68138
// https://oi-wiki.org/graph/cut/#_1
// low(v): 在不经过 v 父亲的前提下能到达的最小的时间戳
// 模板题 https://www.luogu.com.cn/problem/P3388
func (*graph) findCutVertices(n int, g [][]int) (isCut []bool) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	isCut = make([]bool, n)
	dfn := make([]int, n) // 值从 1 开始
	dfsClock := 0
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
				if lowW >= dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 的祖先（可以连到 v 上，这也算割顶）
					isCut[v] = true
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] { // 找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 { // 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割顶
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

	cuts := []int{}
	for v, is := range isCut {
		if is {
			cuts = append(cuts, v) // v+1
		}
	}

	return
}

// 桥（割边）
// https://oi-wiki.org/graph/cut/#_4
// https://algs4.cs.princeton.edu/41graph/Bridge.java.html
// 模板题 https://codeforces.com/problemset/problem/1000/E
// 题目推荐 https://cp-algorithms.com/graph/bridge-searching.html#toc-tgt-2
// 与 MST 结合 https://codeforces.com/problemset/problem/160/D
func (*graph) findBridges(in io.Reader, n, m int) (isBridge []bool) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type neighbor struct{ to, eid int }
	type edge struct{ v, w int }

	g := make([][]neighbor, n)
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], neighbor{w, i})
		g[w] = append(g[w], neighbor{v, i})
		edges[i] = edge{v, w}
	}
	isBridge = make([]bool, m)
	dfn := make([]int, n) // 值从 1 开始
	dfsClock := 0
	tmpID := make([]int, n) // EXTRA
	for i := range tmpID {
		tmpID[i] = -1
	}
	//low := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		for _, e := range g[v] {
			w := e.to
			if dfn[w] == 0 {
				lowW := f(w, v)
				if lowW > dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 或 v 的祖先，所以 v-w 必定是桥
					isBridge[e.eid] = true
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] { // 找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}

		{
			// EXTRA: 割边上有重边时，不能算作割边 https://codeforces.com/problemset/problem/160/D
			for _, e := range g[v] {
				if tmpID[e.to] == -1 {
					tmpID[e.to] = e.eid
				} else {
					isBridge[tmpID[e.to]] = false
					isBridge[e.eid] = false
				}
			}
			for _, e := range g[v] {
				tmpID[e.to] = -1
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

	bridgeEIDs := []int{}
	for eid, is := range isBridge {
		if is {
			bridgeEIDs = append(bridgeEIDs, eid)
		}
	}

	return
}

// 无向图的双连通分量 Biconnected Components (BCC)
// v-BCC：任意割点都是至少两个不同 v-BCC 的公共点
// https://oi-wiki.org/graph/bcc/
// https://www.csie.ntu.edu.tw/~hsinmu/courses/_media/dsa_13spring/horowitz_306_311_biconnected.pdf
// 好题 https://codeforces.com/problemset/problem/962/F
/*
使用 https://csacademy.com/app/graph_editor/ 显示下面的样例
基础样例 - 一个割点两个简单环
1 2
2 3
3 4
4 1
3 5
5 6
6 7
7 3
基础样例 - 两个割点三个简单环（注意那条含有两个割点的边）
7 3
7 4
1 2
2 3
3 1
3 4
4 5
5 6
6 4
*/
func (G *graph) findVertexBCC(n int, g [][]int) (comps [][]int, bccIDs []int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	bccIDs = make([]int, n) // ID 从 1 开始编号
	idCnt := 0
	isCut := make([]bool, n)

	dfn := make([]int, n) // 值从 1 开始
	dfsClock := 0
	type edge struct{ v, w int } // eid
	stack := []edge{}
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			e := edge{v, w} // ne.eid
			if dfn[w] == 0 {
				stack = append(stack, e)
				childCnt++
				lowW := f(w, v)
				if lowW >= dfn[v] {
					isCut[v] = true
					idCnt++
					comp := []int{}
					//eids := []int{}
					for {
						e, stack = stack[len(stack)-1], stack[:len(stack)-1]
						if bccIDs[e.v] != idCnt {
							bccIDs[e.v] = idCnt
							comp = append(comp, e.v)
						}
						if bccIDs[e.w] != idCnt {
							bccIDs[e.w] = idCnt
							comp = append(comp, e.w)
						}
						//eids = append(eids, e.eid)
						if e.v == v && e.w == w {
							break
						}
					}
					// 点数和边数相同，说明该 v-BCC 是一个简单环，且环上所有的边只属于一个简单环
					//if len(comp) == len(eids) {
					//	for _, eid := range eids {
					//		onSimpleCycle[eid] = true
					//	}
					//}
					comps = append(comps, comp)
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] {
				stack = append(stack, e)
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 {
			isCut[v] = false
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			if len(g[v]) == 0 { // 零度，即孤立点（isolated vertex）
				idCnt++
				bccIDs[v] = idCnt
				comps = append(comps, []int{v})
				continue
			}
			f(v, -1)
		}
	}

	// EXTRA: 缩点
	// BCC 和割点作为新图中的节点，并在每个割点与包含它的所有 BCC 之间连边
	vid := idCnt
	cutIDs := make([]int, n) // 接在 BCC 之后给割点编号
	for i, is := range isCut {
		if is {
			vid++
			cutIDs[i] = vid
		}
	}
	for v, cp := range comps {
		v++
		for _, w := range cp {
			if w = cutIDs[w]; w > 0 {
				// add(v,w); add(w,v) ...
			}
		}
	}

	return
}

// e-BCC：删除无向图中所有的割边后，剩下的每一个 CC 都是 e-BCC
// 缩点后形成一颗 bridge tree
// 模板题 https://codeforces.com/problemset/problem/1000/E
func (G *graph) findEdgeBCC(in io.Reader, n, m int) (comps [][]int, bccIDs []int) {
	type neighbor struct{ to, eid int }
	type edge struct{ v, w int }

	g := make([][]neighbor, n)
	edges := make([]edge, m)
	// read g ...

	isBridge := G.findBridges(in, n, m)

	// 求原图中每个点的 bccID
	bccIDs = make([]int, n)
	idCnt := 0
	var comp []int
	var f2 func(int)
	f2 = func(v int) {
		bccIDs[v] = idCnt
		comp = append(comp, v)
		for _, e := range g[v] {
			if w := e.to; bccIDs[w] == 0 && !isBridge[e.eid] {
				f2(w)
			}
		}
	}
	for i, id := range bccIDs {
		if id == 0 {
			idCnt++
			comp = []int{}
			f2(i)
			comps = append(comps, comp)
		}
	}

	// EXTRA: 缩点，复杂度 O(M)
	// 遍历 edges，若两端点的 bccIDs 不同（割边）则建边
	g2 := make([][]int, idCnt)
	for _, e := range edges {
		if v, w := bccIDs[e.v]-1, bccIDs[e.w]-1; v != w {
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
		}
	}

	// 也可以遍历 isBridge，割边两端点 bccIDs 一定不同
	for eid, is := range isBridge {
		if is {
			e := edges[eid]
			v, w := bccIDs[e.v]-1, bccIDs[e.w]-1
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
		}
	}

	return
}

// 任意两点最短路
// 传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// https://oi-wiki.org/graph/shortest-path/#floyd
// 题目推荐 https://cp-algorithms.com/graph/all-pair-shortest-path-floyd-warshall.html#toc-tgt-5
// https://codeforces.com/problemset/problem/1204/C
// 好题 https://codeforces.com/problemset/problem/295/B
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

// 最小环
// 传入邻接矩阵 weights
// weights[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/min-circle/#floyd
// NOTE: 无权图的情况见 shortestCycleBFS
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

type hPair struct {
	v   int
	dis int64
}
type pairHeap []hPair

func (h pairHeap) Len() int              { return len(h) }
func (h pairHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis } // > 权值最大
func (h pairHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(v interface{})   { *h = append(*h, v.(hPair)) }
func (h *pairHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *pairHeap) push(v hPair)         { heap.Push(h, v) }
func (h *pairHeap) pop() hPair           { return heap.Pop(h).(hPair) }

// 单源最短路 Dijkstra
// 适用于稀疏图 O(mlogn)
// start 也可以是一个点集，这相当于同时对多个点跑最短路
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 模板题 https://www.luogu.com.cn/problem/P3371 https://www.luogu.com.cn/problem/P4779
//       https://codeforces.com/problemset/problem/20/C
// 稠密图 https://atcoder.jp/contests/arc064/tasks/arc064_c
// 题目推荐 https://cp-algorithms.com/graph/dijkstra.html#toc-tgt-5
// todo 与线段树结合跑单源最短路 https://codeforces.com/problemset/problem/786/B
func (*graph) shortestPathDijkstra(in io.Reader, n, m, st int) (dist []int64) {
	type neighbor struct {
		to int
		wt int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		v, w, wt := 0, 0, int64(0)
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	const inf int64 = 1e18 // 1e9+1
	dist = make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	vis := make([]bool, n) // 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
	fa := make([]int, n)
	for i := range fa {
		fa[i] = -1
	}

	q := pairHeap{{st, 0}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		// 注：不要 vis 相关的代码也可以，对运行速度几乎没有影响
		if vis[v] { // dist[v] < p.dis
			continue
		}
		vis[v] = true
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] { // > 权值最大
				dist[w] = newD
				fa[w] = v
				q.push(hPair{w, newD})
			}
		}
	}

	//if dist[end] == inf {
	//	return -1
	//}

	// EXTRA: path from end to start
	var end = n - 1
	path := []int{}
	for x := end; x != -1; x = fa[x] {
		path = append(path, x)
	}

	// EXTRA: dist 也可以理解成「时刻」
	// 这种情况下，dist[v] 表示从 start 出发到 v 的耗时

	// EXTRA: 对于相邻的两点，记边为 e，若有：
	// abs(dist[v], dist[w]) == e.wt  =>  e 在最短路上（不带绝对值的话就有先后关系）
	// abs(dist[v], dist[w])  < e.wt  =>  e 不在最短路上（可以想象成一条鼓起的线）
	// 这里的最短路可以有多条

	// EXTRA: 次短路
	{
		const inf int64 = 1e18 // 1e9+1
		dist := make([]int64, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		dist2 := make([]int64, n)
		for i := range dist2 {
			dist2[i] = inf
		}
		vis := make([]bool, n)

		q := pairHeap{{st, 0}}
		for len(q) > 0 {
			v := q.pop().v
			if vis[v] {
				continue
			}
			vis[v] = true
			for _, e := range g[v] {
				w := e.to
				newD := dist[v] + e.wt
				if newD < dist[w] {
					q.push(hPair{w, newD})
					dist[w], newD = newD, dist[w]
				}
				if newD > dist[w] && newD < dist2[w] {
					dist2[w] = newD
					q.push(hPair{w, newD})
				}
			}
		}
	}

	return
}

// 01 最短路
// https://oi-wiki.org/graph/bfs/#bfs_3
// https://codeforces.com/blog/entry/22276
// 例题: https://codeforces.com/problemset/problem/173/B
func (*graph) bfs01(in io.Reader, n, m, st int) []int {
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
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
			w, d := e.to, e.wt
			if newD := dist[v] + d; newD < dist[w] {
				dist[w] = newD
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

// 单源最短路 SPFA O(nm)   Bellman-Ford
// 有负环时返回 nil
// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
// EXTRA: 差分约束，若有 Xi-Xj<=Ck，则连一条有向边 j->i，边权为 Ck
//        然后再添加一个 0 号节点，向其他节点连一条边权为 0 的有向边，表示 Xi-X0<=0
//        这样，在无负环时会得到一组非正数解
//        模板题 https://www.luogu.com.cn/problem/P3385
func (*graph) shortestPathSPFA(in io.Reader, n, m, st int) (dist []int64) {
	type neighbor struct {
		to int
		wt int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		v, w, wt := 0, 0, int64(0)
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	const inf int64 = 1e18 // 1e9+1
	dist = make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	onQ := make([]bool, n)
	onQ[st] = true
	relaxedCnt := make([]int, n)
	q := []int{st}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		onQ[v] = false
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				relaxedCnt[w] = relaxedCnt[v] + 1
				if relaxedCnt[w] >= n {
					return nil
				} // found negative cycle
				if !onQ[w] {
					q = append(q, w)
					onQ[w] = true
				}
			}
		}
	}

	// EXTRA: 打印负环
	// https://cp-algorithms.com/graph/finding-negative-cycle-in-graph.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BellmanFordSP.java.html

	return
}

// EXTRA: 同余最短路
// todo 跳楼机 https://www.luogu.com.cn/problem/P3403

// 最小生成树 Kruskal
// 适用于稀疏图 O(mlogm)，或者边已经按权值排序的情况
// https://oi-wiki.org/graph/mst/#kruskal
// 模板题 https://www.luogu.com.cn/problem/P3366 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/E
// 题目推荐 https://cp-algorithms.com/graph/mst_kruskal.html#toc-tgt-5
// 关键边、伪关键边（与割边结合）https://codeforces.com/problemset/problem/160/D
func (*graph) mstKruskal(in io.Reader, n, m int) int64 {
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
		v, w int
		wt   int // int64
		eid  int // 某些题目需要
	}
	edges := make([]edge, m)
	for i := range edges {
		v, w, wt := 0, 0, 0
		Fscan(in, &v, &w, &wt)
		v--
		w--
		edges[i] = edge{v, w, wt, i}
	}

	// 边权范围小的话也可以用桶排
	sort.Slice(edges, func(i, j int) bool { return edges[i].wt < edges[j].wt })
	initFa(n)
	sum := int64(0)
	cntE := 0
	for _, e := range edges {
		if fv, fw := find(e.v), find(e.w); fv != fw {
			sum += int64(e.wt)
			fa[fv] = fw
			cntE++
		}
	}
	// 图不连通
	if cntE < n-1 {
		return -1
	}
	return sum
}

// 最小生成树 Prim
// 适用于稠密图 O(n^2)，传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/mst/#prim
// 题目：https://codeforces.com/contest/1245/problem/D
func (*graph) mstPrim(dist [][]int) (sum int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const inf int = 1.1e9
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

// Boruvka's algorithm
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BoruvkaMST.java.html

// 次小生成树
// Second best Minimum Spanning Tree
// Using Kruskal and Lowest Common Ancestor
// todo https://cp-algorithms.com/graph/second_best_mst.html
func (*graph) secondMST(n, m int) (sum int64) {
	panic("TODO")
}

// Kruskal 重构树
// https://oi-wiki.org/graph/mst/#kruskal_1

// 最小生成树计数 Kirchhoff's theorem

// 最小树形图 - 朱刘算法

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
	colors := make([]int8, n) // 0 表示未访问该节点
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			if colors[w] == c || colors[w] == 0 && !f(w, 3-c) {
				return false
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

// 寻找二分图中的奇环（无奇环是该图为二分图的充要条件 https://www.zhihu.com/question/52978925）
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/Bipartite.java.html
func (*graph) bipartiteFindOddLengthCycle(n int, g [][]int) (cycle []int) {
	colors := make([]int8, n) // 0 表示未访问该节点
	fa := make([]int, n)
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			if colors[w] == c {
				cycle = append(cycle, w)
				for x := v; x != w; x = fa[x] {
					cycle = append(cycle, x)
				}
				//cycle = append(cycle, w)
				return false
			}
			if colors[w] == 0 {
				fa[w] = v
				if !f(w, 3-c) {
					return false
				}
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 && !f(i, 1) {
			return
		}
	}
	return
}

/* 匹配 独立集 覆盖（边覆盖 顶点覆盖） 支配集
https://en.wikipedia.org/wiki/Matching_(graph_theory)
https://en.wikipedia.org/wiki/Independent_set_(graph_theory)
https://en.wikipedia.org/wiki/Edge_cover
https://en.wikipedia.org/wiki/Vertex_cover
https://en.wikipedia.org/wiki/Dominating_set
https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/

最大匹配+最小边覆盖=n （图中无孤立点）
最大独立集+最小顶点覆盖=n
对于二分图，最小顶点覆盖=最大匹配，最大独立集=n-最大匹配

不是 n-匹配就是 n-独立集 https://codeforces.com/problemset/problem/1198/C
*/

// 二分图最大匹配 - 匈牙利算法/增广路算法 O(nm)    Hungarian algorithm
// https://www.renfei.org/blog/bipartite-matching.html 推荐
// https://www.geeksforgeeks.org/maximum-bipartite-matching/
// https://oi-wiki.org/graph/bi-graph/#_9
// https://zhuanlan.zhihu.com/p/62981901
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BipartiteMatching.java.html
// 模板题 https://www.luogu.com.cn/problem/P3386
// 有向图最小路径覆盖：
//   起初把每个点都视作一条路径，这样共有 n 条不相交路径
//   拆点图：每个点拆成出点和入点，原图的 v->w 相当于拆点图的 出点v->入点w，这说明拆点图一定是二分图
//   在拆点图里找到一条匹配边就相当于把两条路径合成了一条路径，也就相当于路径数减少了 1
//   所以找到了多少匹配边，路径数就减少了多少
//   所以有最小路径覆盖=原图的结点数-拆点图的最大匹配数
// todo 树上最小路径覆盖 http://codeforces.com/problemset/problem/618/D
func (*graph) maxMatchingHungarian(n int, g [][]int) (match []int, cnt int) {
	match = make([]int, n)
	for i := range match {
		match[i] = -1
	}
	var used []bool
	var f func(int) bool
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

// 匈牙利算法的另一种写法，适用左右两侧节点有明确区分的情况，要求 g 中存储的是左侧到右侧的单向边
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
// 算法第四版 https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/HopcroftKarp.java.html
// http://pepcy.cf/icpc-templates/003Graph/hk.html
// https://en.wikipedia.org/wiki/Hopcroft%E2%80%93Karp_algorithm
func (*graph) maxMatchingHopcroftKarp(n int, g [][]int) (match []int, cnt int) {
	// TODO
	return
}

// 带权二分图最大匹配 - Kuhn–Munkres 算法 O(n^3)   KM
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/AssignmentProblem.java.html
// https://zhuanlan.zhihu.com/p/62981901
// https://www.luogu.com.cn/problem/solution/P6577
// 模板题 https://www.luogu.com.cn/problem/P3967 https://www.luogu.com.cn/problem/P6577
func (*graph) maxMatchingKuhnMunkres(n int, g [][]int) (match []int, cnt int) {
	// TODO
	return
}

// 一般图最大匹配（带花树 Edmonds's blossom algorithm）
// https://en.wikipedia.org/wiki/Blossom_algorithm
// TODO https://www.cnblogs.com/cjyyb/p/8719368.html 带花树算法学习笔记
// 模板题 https://www.luogu.com.cn/problem/P6113
//       https://www.luogu.com.cn/problem/P4258

// EXTRA: 完美匹配 Perfect Match
// 完美匹配同时也是一个原图的最小边数的边覆盖

// EXTRA: Min Cost Perfect Matching (MCPM)
// https://courses.engr.illinois.edu/cs598csc/sp2010/Lectures/Lecture11.pdf

// 有向图的拓扑排序 Kahn's algorithm
// 可以用来判断有向图是否有环、求 DAG 上的 DP 等
// https://oi-wiki.org/graph/topo/
// https://cp-algorithms.com/graph/topological-sort.html
// DAG DP https://ac.nowcoder.com/acm/contest/6384/C https://www.luogu.com.cn/problem/P3387
// 好题 https://codeforces.com/problemset/problem/915/D
// 混合图拓扑排序 https://codeforces.com/contest/1385/problem/E
// EXTRA: todo 拓扑排序是否唯一：算法第四版 p.387 4.2.25
// LC 套题 https://leetcode-cn.com/tag/topological-sort/
func (*graph) topSort(in io.Reader, n, m int) (orders []int, isDAG bool) {
	g := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < m; i++ {
		v, w := 0, 0
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w) // 注意不能有自环
		deg[w]++
	}

	//fa := make([]int, n)
	//for i := range fa {
	//	fa[i] = -1
	//}
	//levels := make([]int, n)

	orders = make([]int, 0, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
			//levels[i] = 1
			// NOTE: 对于点权型 DAG DP，这里记得对起点初始化
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		orders = append(orders, v)
		for _, w := range g[v] {
			// do v-w ...

			deg[w]--
			if deg[w] == 0 {
				//fa[w] = v
				//levels[w] = levels[v] + 1
				q = append(q, w)
			}
		}
	}

	isDAG = len(orders) == n

	{
		fa := make([]int, n)

		// EXTRA: path from end to start
		var end = n - 1
		path := make([]int, 0, n)
		for v := end; v != -1; v = fa[v] {
			path = append(path, v)
		}
	}

	return
}

// 强连通分量分解 Strongly Connected Component (SCC)
// sccIDs[v] 表示点 v 所属的 SCC 的拓扑序
// https://oi-wiki.org/graph/scc/#kosaraju
// https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KosarajuSharirSCC.java.html
// 模板题 https://www.luogu.com.cn/problem/P2341
func (*graph) sccKosaraju(in io.Reader, n, m int, g [][]int) (comps [][]int, sccIDs []int) {
	type edge struct{ v, w int }
	edges := make([]edge, m)
	g = make([][]int, n)
	rg := make([][]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
		edges[i] = edge{v, w}
	}

	vs := make([]int, 0, n)
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
		vs = append(vs, v) // 后序。保证遍历出来的结果的逆后序生成的 SCC 一定是拓扑序
	}
	for v := range g {
		if !vis[v] {
			dfs(v)
		}
	}

	vis = make([]bool, n)
	var comp []int
	var rdfs func(int)
	rdfs = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range rg[v] {
			if !vis[w] {
				rdfs(w)
			}
		}
	}
	// 逆后序遍历，就可以像无向图那样求出 SCC
	comps = [][]int{}
o:
	for i := len(vs) - 1; i >= 0; i-- {
		if v := vs[i]; !vis[v] {
			comp = []int{}
			rdfs(v)
			// EXTRA: len(comp) >= 3 说明有环，注意环的个数可能不止一个
			comps = append(comps, comp)
			// EXTRA: 判断缩点后是否出度为 0
			for _, u := range comp {
				for _, w := range g[u] {
					if !vis[w] {
						continue o
					}
				}
			}
			// 出度为 0
		}
	}

	// comps 的结果就是拓扑序
	sccIDs = make([]int, n)
	for i, cp := range comps {
		for _, v := range cp {
			sccIDs[v] = i
		}
	}

	// EXTRA: 缩点: 将边 v-w 转换成 sccIDs[v]-sccIDs[w]
	// 缩点后得到了一张 DAG，点的编号范围为 [0,len(comps)-1]
	// 注意这样可能会产生重边，不能有重边时可以对每个点排序去重
	// 模板题 点权 https://www.luogu.com.cn/problem/P3387
	// 		 边权 https://codeforces.com/contest/894/problem/E
	for _, e := range edges {
		if v, w := sccIDs[e.v], sccIDs[e.w]; v != w {

		} else {
			// EXTRA: 汇合同一个 SCC 的权值 ...

		}
	}

	// EXTRA: 求有多少个点能被其他所有点访问到 https://www.luogu.com.cn/problem/P2341
	lastComp := comps[len(comps)-1]
	numCanBeVisitedFromAll := len(lastComp)
	vis = make([]bool, n)
	rdfs(lastComp[0])
	for _, use := range vis {
		if !use {
			numCanBeVisitedFromAll = 0
		}
	}
	_ = numCanBeVisitedFromAll

	return
}

// TODO: SCC Tarjan
// 常数比 Kosaraju 小
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TarjanSCC.java.html
// https://stackoverflow.com/questions/32750511/does-tarjans-scc-algorithm-give-a-topological-sort-of-the-scc

// Gabow's algorithm
// 常数比 Kosaraju 大
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GabowSCC.java.html

// 2-Satisfiability (2-SAT)
// https://oi-wiki.org/graph/2-sat/
// https://cp-algorithms.com/graph/2SAT.html
// https://zhuanlan.zhihu.com/p/50211772
// todo https://www.luogu.com.cn/blog/85514/post-2-sat-xue-xi-bi-ji
// 讲解+套题 https://codeforces.com/blog/entry/16205
// 2-SAT 总结 by kuangbin https://www.cnblogs.com/kuangbin/archive/2012/10/05/2712429.html
// 一般 ¬x 用 x+n 表示
// NOTE: 单独的条件 x为a 可以用 (x为a)∨(x为a) 来表示
// NOTE: 一些建边的转换：
//       A,B 至少存在一个 (A|B)    ¬A⇒B, ¬B⇒A
//       A,B 不能同时存在 (¬A|¬B)  A⇒¬B, B⇒¬A
//       A,B 必须且只一个 (A^B)    A⇒¬B, B⇒¬A, ¬A⇒B, ¬B⇒A
//       A,B 同时或都不在 (¬(A^B)) A⇒B, B⇒A, ¬A⇒¬B, ¬B⇒¬A
//       A 必须存在       (A)     ¬A⇒A
//       A 不能存在       (¬A)     A⇒¬A
// 下面的代码基于模板题 https://www.luogu.com.cn/problem/P4782
// 读入 m 条数据，每条数据表示 (x为a)∨(y为b)，a b 为 0 或 1
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

// 基环树
// 对于内向基环树，由于每个点的出度均为一，可以用 []int 来表示图
// https://www.luogu.com.cn/blog/user52918/qian-tan-ji-huan-shu
// https://codeforces.com/problemset/problem/1027/D
// https://codeforces.com/problemset/problem/1335/F
func (*graph) treeWithCycle(n int, g []int, rg [][]int) {
	inDeg := make([]int, n)
	// 计算入度 ...

	vis := make([]int8, n)
	// 拓扑排序，之后 vis[v]==0 的点必定在基环上
	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		vis[v] = 1
		w := g[v]
		inDeg[w]--
		if inDeg[w] == 0 {
			q = append(q, w)
		}
	}

	for i, b := range vis {
		if b > 0 {
			continue
		}
		cycle := []int{}
		for v := i; vis[v] == 0; v = g[v] {
			vis[v] = 2 // 标记为基环上的点
			cycle = append(cycle, v)
		}

		sz := len(cycle)
		for j, cv := range cycle {
			for _, root := range rg[cv] {
				if vis[root] == 2 {
					continue
				}
				// EXTRA: 从非基环的根部出发，遍历反图 rg ...

			}

			// EXTRA: 从基环上的位置 j 倒着走 k 步到达的点
			var k int
			_ = cycle[((j-k)%sz+sz)%sz]
		}
	}
}

// 圆方树
// todo https://www.luogu.com.cn/blog/PinkRabbit/Introduction-to-Round-Square-Tree

/* 网络流
上下界网络流 https://oi-wiki.org/graph/flow/bound/
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FordFulkerson.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FlowNetwork.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FlowEdge.java.html
全局最小割 https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GlobalMincut.java.html
EXTRA: Disjoint paths
       Edge-disjoint paths: It turns out that the maximum number of edge-disjoint paths equals the maximum flow of the graph,
                            assuming that the capacity of each edge is one.
       Node-disjoint paths: 拆点法
EXTRA: todo 路径覆盖 Path cover + 打印
todo Competitive Programmer’s Handbook Ch.20
     线性规划与网络流24题-最小路径覆盖问题 https://byvoid.com/zhs/blog/lpf24-3/
*/

// 最大流 Dinic's algorithm O(n^2 * m)
// Ford–Fulkerson algorithm (FFA) 的改进版本
// https://oi-wiki.org/graph/flow/max-flow/#dinic
// https://cp-algorithms.com/graph/dinic.html
// 模板题 https://www.luogu.com.cn/problem/P3376
// 建模题 https://atcoder.jp/contests/arc085/tasks/arc085_c
//       https://codeforces.com/problemset/problem/1360/G (满流时有解)
//       https://codeforces.com/problemset/problem/546/E
func (*graph) maxFlowDinic(in io.Reader, n, m, st, end int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// st--; end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for i := 0; i < m; i++ {
		var v, w, cap int
		Fscan(in, &v, &w, &cap)
		v--
		w--
		addEdge(v, w, cap)
	}

	// 计算从源点 st 出发的距离标号
	level := make([]int, n)
	calcLevel := func() bool {
		for i := range level {
			level[i] = -1
		}
		level[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && level[w] < 0 {
					level[w] = level[v] + 1
					q = append(q, w)
				}
			}
		}
		return level[end] >= 0
	}
	// 寻找增广路
	var iter []int // 当前弧，在其之前的边已经没有用了，避免对没有用的边进行多次检查
	var dfs func(int, int) int
	dfs = func(v int, minF int) int {
		if v == end {
			return minF
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && level[w] > level[v] {
				if f := dfs(w, min(minF, e.cap)); f > 0 {
					e.cap -= f
					g[w][e.rid].cap += f
					return f
				}
			}
		}
		return 0
	}

	maxFlow := 0        // int64
	const inf int = 1e9 // 1e18
	for calcLevel() {
		iter = make([]int, n)
		for {
			if f := dfs(st, inf); f > 0 {
				maxFlow += f
			} else {
				break
			}
		}
	}

	// EXTRA: 打印流的分配方案，输出反向边上的 cap

	return maxFlow
}

// 最高标号预流推进 (HLPP, High Level Preflow Push)
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm
// todo https://oi-wiki.org/graph/flow/max-flow/#hlpp
//      https://www.luogu.com.cn/blog/ONE-PIECE/jiu-ji-di-zui-tai-liu-suan-fa-isap-yu-hlpp
// 模板题 https://www.luogu.com.cn/problem/P4722

// 最小费用流
// 有两种实现：SPFA O(fnm) 和 Dijkstra O(fmlogn)
// 要求图中无负圈
// https://oi-wiki.org/graph/flow/min-cost/
// 模板题 https://www.luogu.com.cn/problem/P3381
// 性能对比（由于数据不强所以 SPFA 很快）：SPFA 1.05s   Dijkstra 1.91s
// todo 对比下 lrj 的建图方式
func (*graph) minCostFlowSPFA(in io.Reader, n, m, st, end, F int) int64 {
	// st--; end--

	type neighbor struct{ to, rid, cap, cost int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i := 0; i < m; i++ {
		var v, w, cap, cost int
		Fscan(in, &v, &w, &cap, &cost)
		v--
		w--
		addEdge(v, w, cap, cost)
	}

	const inf int64 = 1e18
	dist := make([]int64, n)
	type pair struct{ v, i int }
	fa := make([]pair, n)
	spfa := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		onQ := make([]bool, n)
		onQ[st] = true
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			onQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + int64(e.cost); newD < dist[w] {
					dist[w] = newD
					fa[w] = pair{v, i}
					if !onQ[w] {
						q = append(q, w)
						onQ[w] = true
					}
				}
			}
		}
		return dist[end] < inf
	}
	minCost := int64(0)
	for F > 0 && spfa() { // 若求最小费用最大流 MCMF，把 F>0 去掉
		// 沿 st-end 的最短路尽量增广
		minF := F // inf
		for v := end; v != st; {
			p := fa[v]
			if c := g[p.v][p.i].cap; c < minF {
				minF = c
			}
			v = p.v
		}
		F -= minF // maxFlow += minF
		minCost += dist[end] * int64(minF)
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
	}
	if F > 0 {
		return -1
	}
	return minCost
}

// 基于原始对偶方法 (primal-dual method)
// https://blog.xehoth.cc/DurationPlan-Primal-Dual/
func (*graph) minCostFlowDijkstra(in io.Reader, n, m, st, end, F int) int64 {
	// st--; end--

	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i := 0; i < m; i++ {
		var v, w, cap, cost int
		Fscan(in, &v, &w, &cap, &cost)
		v--
		w--
		addEdge(v, w, cap, cost)
	}

	const inf int64 = 1e18
	h := make([]int64, n) // 顶点的势
	dist := make([]int64, n)
	type pair struct{ v, i int }
	fa := make([]pair, n)
	dijkstra := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		q := pairHeap{{st, 0}}
		for len(q) > 0 {
			p := q.pop()
			v := p.v
			if dist[v] < p.dis {
				continue
			}
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + int64(e.cost) + h[v] - h[w]; newD < dist[w] {
					dist[w] = newD
					fa[w] = pair{v, i}
					q.push(hPair{w, newD})
				}
			}
		}
		return dist[end] < inf
	}
	minCost := int64(0)
	for F > 0 && dijkstra() {
		for i, d := range dist {
			h[i] += d
		}
		minF := F // inf
		for v := end; v != st; {
			p := fa[v]
			if c := g[p.v][p.i].cap; c < minF {
				minF = c
			}
			v = p.v
		}
		F -= minF // maxFlow += minF
		minCost += h[end] * int64(minF) // 注意这里是 h 不是 dist
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
	}
	if F > 0 {
		return -1
	}
	return minCost
}

// todo 基于 Capacity Scaling 的弱多项式复杂度最小费用流算法 https://ouuan.github.io/post/%E5%9F%BA%E4%BA%8E-capacity-scaling-%E7%9A%84%E5%BC%B1%E5%A4%9A%E9%A1%B9%E5%BC%8F%E5%A4%8D%E6%9D%82%E5%BA%A6%E6%9C%80%E5%B0%8F%E8%B4%B9%E7%94%A8%E6%B5%81%E7%AE%97%E6%B3%95/

// ZKW 费用流
// https://artofproblemsolving.com/community/c1368h1020435
