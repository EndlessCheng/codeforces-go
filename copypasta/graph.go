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

CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=graphs
构造 https://codeforces.com/problemset/problem/41/E

TIPS: 使用一个 fa 数组（初始化为 -1）记录搜索树中的节点的父节点，这样对每个节点都有一条到根的路径（根的 fa 为 -1）
NOTE: 独立集相关问题，可以从染色的角度考虑
NOTE: 度数大于 √M 的点不超过 2√M 个

一个错误的结论：有向图，按照深度奇偶性分类，子图一定不存在长度大于 1 的路径
from https://codeforces.com/problemset/problem/1368/E
input
4 5
3 5
1 3
1 5
2 3
3 4
2 4

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

// 0-1 矩阵，有向图
func (*graph) matrixToEdges(mat []string) [][2]int {
	edges := [][2]int{}
	for v, row := range mat {
		for w, weight := range row {
			if weight != '0' {
				edges = append(edges, [2]int{v, w})
			}
		}
	}
	return edges
}

func (*graph) simpleSearch(n, st int, g [][]int) {
	// DFS
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
	// 或者（有向图/森林）
	for i, b := range vis {
		if !b {
			f(i)
		}
	}

	{
		// 有向图的环/回边检测
		//《算法导论》p.353 边的分类
		// vis[v] == 0：该顶点未被访问
		// vis[v] == 1：该顶点已经被访问，其子树未遍历完
		// vis[v] == 2：该顶点已经被访问，其子树已遍历完
		vis := make([]int8, n)
		var f func(int)
		f = func(v int) {
			vis[v] = 1
			for _, w := range g[v] {
				if b := vis[w]; b == 0 {
					f(w) // 树边
				} else if b == 1 {
					// 后向边，说明有环
				} else {
					// 前向边或横向边
				}
			}
			vis[v] = 2
		}
		for i, b := range vis {
			if b == 0 {
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
		f(0, -1, 1)

		ans := []interface{}{st + 1} // for print
		for v := end; v != st; v = fa[v] {
			ans = append(ans, v+1)
		}
	}

	// BFS
	vis = make([]bool, n)
	vis[st] = true
	q := []int{st}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// ...
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				q = append(q, w)
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
}

func (*graph) depthArray(n, st int, g [][]int) []int {
	dep := make([]int, n)
	for i := range dep {
		dep[i] = -1
	}
	dep[st] = 0
	q := []int{st}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// do(v, dep[v])
		for _, w := range g[v] {
			if dep[w] == -1 {
				dep[w] = dep[v] + 1
				q = append(q, w)
			}
		}
	}
	return dep
}

// BFS 应用：求无向无权图最小环长度
// 好题 https://codeforces.ml/problemset/problem/1325/E
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
// 模板题 https://www.luogu.com.cn/problem/P2731
//       https://www.luogu.com.cn/problem/P1341
func (*graph) eulerianPath(n, m int) bool {
	type neighbor struct{ to, eid int }
	g := make([][]neighbor, n)
	// read g ...

	var st int
	oddDegCnt := 0
	for i := len(g) - 1; i >= 0; i-- { // 倒着遍历保证起点的字典序最小
		if deg := len(g[i]); deg > 0 {
			if deg&1 != 0 {
				st = i
				oddDegCnt++
			} else if oddDegCnt == 0 {
				st = i
			}
		}
	}

	if oddDegCnt > 2 {
		return false
	} // NOTE: 若没有奇度数，则寻找的是欧拉回路

	path := make([]int, 0, m)
	iter := make([]int, n)
	vis := make([]bool, m)
	var f func(int)
	f = func(v int) {
		for ; iter[v] < len(g[v]); {
			e := g[v][iter[v]]
			iter[v]++
			if id := e.eid; !vis[id] {
				vis[id] = true
				f(e.to)
				path = append(path, id)
			}
		}
	}
	f(st)

	// NOTE: 若输出的是顶点，可以在递归内部 path = append(path, e.to)，最后把 st 添加到末尾

	// 倒序输出 path...

	return true
}

/* Topic - DFS 树
讲解+套题 https://codeforces.ml/blog/entry/68138
好题：https://codeforces.ml/problemset/problem/1325/F
*/

// 割点（割顶） cut vertices / articulation points
// https://codeforces.ml/blog/entry/68138
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
// 模板题 https://codeforces.ml/problemset/problem/1000/E
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
// 好题 https://codeforces.ml/problemset/problem/962/F
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
// 模板题 https://codeforces.ml/problemset/problem/1000/E
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
// https://codeforces.ml/problemset/problem/1204/C
// 好题 https://codeforces.ml/problemset/problem/295/B
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
	dis int64
	v   int
}
type pairHeap []hPair

func (h pairHeap) Len() int              { return len(h) }
func (h pairHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h pairHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(v interface{})   { *h = append(*h, v.(hPair)) }
func (h *pairHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *pairHeap) push(v hPair)         { heap.Push(h, v) }
func (h *pairHeap) pop() hPair           { return heap.Pop(h).(hPair) }
func (h pairHeap) empty() bool           { return len(h) == 0 }

// 单源最短路 Dijkstra
// 适用于稀疏图 O((|E|+|V|)⋅log|V|)
// start 也可以是一个点集，这相当于同时对多个点跑最短路
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 模板题 https://www.luogu.com.cn/problem/P4779
//       https://codeforces.com/problemset/problem/20/C
// 题目推荐 https://cp-algorithms.com/graph/dijkstra.html#toc-tgt-5
// 与线段树结合跑单源最短路 https://codeforces.ml/problemset/problem/786/B
func (*graph) shortestPathDijkstra(in io.Reader, n, m, st int) (dist []int64) {
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
	dist[st] = 0
	fa := make([]int, n)
	for i := range fa {
		fa[i] = -1
	}

	h := pairHeap{{0, st}}
	for !h.empty() {
		p := h.pop()
		d, v := p.dis, p.v
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := d + e.weight; newD < dist[w] {
				dist[w] = newD
				fa[w] = v
				h.push(hPair{newD, w})
			}
		}
	}

	//if dist[n-1] == inf {
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
	// abs(dist[v], dist[w]) == e.weight => e 在最短路上（不带绝对值的话就有先后关系）
	// abs(dist[v], dist[w])  < e.weight => e 不在最短路上（可以想象成一条鼓起的线）
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

		h := pairHeap{{0, st}}
		for !h.empty() {
			p := h.pop()
			d, v := p.dis, p.v
			if dist2[v] < d {
				continue
			}
			for _, e := range g[v] {
				w := e.to
				newD := d + e.weight
				if newD < dist[w] {
					h.push(hPair{newD, w})
					dist[w], newD = newD, dist[w]
				}
				if newD > dist[w] && newD < dist2[w] {
					dist2[w] = newD
					h.push(hPair{newD, w})
				}
			}
		}
	}

	return
}

// 01 最短路
// https://oi-wiki.org/graph/bfs/#bfs_3
// https://codeforces.ml/blog/entry/22276
// 例题: https://codeforces.ml/problemset/problem/173/B
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

// 单源最短路 SPFA
// 有负环时返回 nil
// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
// EXTRA: 差分约束，若有 Xi-Xj<=Ck，则连一条有向边 j->i，边权为 Ck
//        然后再添加一个 0 号节点，向其他节点连一条边权为 0 的有向边，表示 Xi-X0<=0
//        这样，在无负环时会得到一组非正数解
//        模板题 https://www.luogu.com.cn/problem/P3385
func (*graph) shortestPathBellmanFord(in io.Reader, n, m, st int) (dist []int64) {
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
			if newD := dist[v] + e.weight; newD < dist[w] {
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
// 适用于稀疏图 O(|E|⋅log|E|)，或者边已经按权值排序的情况
// https://oi-wiki.org/graph/mst/#kruskal
// 模板题 https://www.luogu.com.cn/problem/P3366
// 题目推荐 https://cp-algorithms.com/graph/mst_kruskal.html#toc-tgt-5
// 关键边、伪关键边（与割边结合）https://codeforces.com/problemset/problem/160/D
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
		weight int // int64
		eid    int // 某些题目需要
	}
	edges := make([]edge, m)
	for i := range edges {
		var v, w, weight int
		Fscan(in, &v, &w, &weight)
		v--
		w--
		edges[i] = edge{v, w, weight, i}
	}

	// weight 范围小的话也可以用桶排
	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
	initFa(n)
	for _, e := range edges {
		if fv, fw := find(e.v), find(e.w); fv != fw {
			sum += int64(e.weight)
			fa[fv] = fw
		}
	}
	return
}

// 最小生成树 Prim
// 适用于稠密图 O(|V|^2)，传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/mst/#prim
// 题目：https://codeforces.ml/contest/1245/problem/D
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

// 二分图最大匹配 - 匈牙利算法/增广路算法 O(nm)
// 最大匹配+最小边覆盖=n
// 最大独立集+最小顶点覆盖=n    最大独立集=n-最大匹配
// 最大匹配=最小顶点覆盖
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
// todo 树上最小路径覆盖 http://codeforces.ml/problemset/problem/618/D
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

// 带权二分图最大匹配 - Kuhn–Munkres 算法 O(n^3)
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/AssignmentProblem.java.html
// https://zhuanlan.zhihu.com/p/62981901
func (*graph) maxMatchingKuhnMunkres(n int, g [][]int) (match []int, cnt int) {
	// TODO
	return
}

// 一般图最大匹配（带花树）
// 模板题 https://www.luogu.com.cn/problem/P6113
// TODO

// 有向图的拓扑排序 Kahn's algorithm
// 可以用来判断有向图是否有环、求 DAG 上的 DP 等
// https://oi-wiki.org/graph/topo/
// https://cp-algorithms.com/graph/topological-sort.html
// 模板题 https://www.luogu.com.cn/problem/P3387
// EXTRA: todo 拓扑排序是否唯一：算法第四版 p.387 4.2.25
// LC 套题 https://leetcode-cn.com/tag/topological-sort/
func (*graph) topSort(in io.Reader, n, m int) (orders []int, isDAG bool) {
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

	//fa := make([]int, n)
	//for i := range fa {
	//	fa[i] = -1
	//}
	//levels := make([]int, n)

	orders = make([]int, 0, n)
	q := []int{}
	for i, deg := range inDeg {
		if deg == 0 {
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
			inDeg[w]--
			if inDeg[w] == 0 {
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
		for x := end; x != -1; x = fa[x] {
			path = append(path, x)
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
	used := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		used[v] = true
		for _, w := range g[v] {
			if !used[w] {
				dfs(w)
			}
		}
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
	// 按照 DFS 逆后序就可以像无向图那样求出 CC
	comps = [][]int{}
	for i := len(vs) - 1; i >= 0; i-- {
		if v := vs[i]; !used[v] {
			comp = []int{}
			rdfs(v)
			// EXTRA: len(comp) >= 3 说明有环，注意环的个数可能不止一个
			comps = append(comps, comp)
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
	// 模板题 https://www.luogu.com.cn/problem/P3387
	for _, e := range edges {
		if v, w := sccIDs[e.v], sccIDs[e.w]; v != w {
			// custom
		}
	}

	// EXTRA: 求有多少个点能被其他所有点访问到 https://www.luogu.com.cn/problem/P2341
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
// 常数比 Kosaraju 小
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TarjanSCC.java.html

// Gabow's algorithm
// 常数比 Kosaraju 大
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GabowSCC.java.html

// 2-Satisfiability (2-SAT)
// https://oi-wiki.org/graph/2-sat/
// https://cp-algorithms.com/graph/2SAT.html
// https://zhuanlan.zhihu.com/p/50211772
// 讲解+套题 https://codeforces.ml/blog/entry/16205
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
// https://www.luogu.com.cn/blog/user52918/qian-tan-ji-huan-shu
func (*graph) treeWithCycle(n int, g [][]int) {
	// EXTRA: 内向基环树找环
	inDeg := make([]int, n) // 计算入度 ...
	visCnt := make([]int8, n)
	var f func(v int)
	f = func(v int) {
		if visCnt[v] == 2 {
			return
		}
		visCnt[v]++
		w := g[v][0] // 所有点的出度均为一
		f(w)
		if visCnt[v] == 1 && visCnt[w] == 2 { // w 为树枝和环的交叉点
			// EXTRA: 从 v 开始遍历反图可以遍历树枝上的所有点
			//        遍历时赋值 visCnt[x] = 1 和 inDeg[x] = -1

			// EXTRA: 环上的反图遍历需要额外判断反图中的点是否满足 visCnt[x] == 2
		}
	}
	for i, d := range inDeg {
		if d == 0 {
			f(i)
		}
	}
	for i, c := range visCnt {
		if c == 0 { // 基环树无树枝
			f(i)
		}
	}
	cycleSize := 0
	for _, c := range visCnt {
		if c == 2 {
			cycleSize++
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
EXTRA: 路径覆盖 Path cover
todo Competitive Programmer’s Handbook Ch.20
     线性规划与网络流24题-最小路径覆盖问题 https://byvoid.com/zhs/blog/lpf24-3/
*/

// 最大流 Dinic's algorithm O(n2m)
// Ford–Fulkerson algorithm (FFA) 的改进版本
// https://oi-wiki.org/graph/flow/max-flow/#dinic
// https://cp-algorithms.com/graph/dinic.html
// 模板题 https://www.luogu.com.cn/problem/P3376
// 建模题 https://atcoder.jp/contests/arc085/tasks/arc085_c
func (*graph) maxFlowDinic(in io.Reader, numV, numE, st, end int) (maxFlow int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// st-- end--

	type edge struct {
		to, rev int // 反向边
		cap     int
	}
	edges := make([][]edge, numV)
	addEdge := func(from, to int, cap int) {
		edges[from] = append(edges[from], edge{to, len(edges[to]), cap})
		edges[to] = append(edges[to], edge{from, len(edges[from]) - 1, 0})
	}
	for i := 0; i < numE; i++ {
		var v, w int
		var cap int
		Fscan(in, &v, &w, &cap)
		v--
		w--
		addEdge(v, w, cap)
	}

	// 计算从源点 st 出发的距离标号
	level := make([]int, numV)
	calcLevel := func() bool {
		for i := range level {
			level[i] = -1
		}
		level[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range edges[v] {
				if w := e.to; e.cap > 0 && level[w] < 0 {
					level[w] = level[v] + 1
					q = append(q, w)
				}
			}
		}
		return level[end] >= 0
	}
	// 寻找增广路
	var iter []int // 当前弧，在其之前的边已经没有用了
	var dfs func(int, int) int
	dfs = func(v int, mf int) int {
		if v == end {
			return mf
		}
		for i := iter[v]; i < len(edges[v]); i++ {
			e := &edges[v][i]
			if w := e.to; e.cap > 0 && level[w] > level[v] {
				if f := dfs(w, min(mf, e.cap)); f > 0 {
					e.cap -= f
					edges[w][e.rev].cap += f
					return f
				}
			}
			iter[v]++ // 当前弧优化（避免对没有用的边进行多次检查）
		}
		return 0
	}
	const inf int = 1e9 // 1e18
	for calcLevel() {
		iter = make([]int, numV)
		for {
			if f := dfs(st, inf); f > 0 {
				maxFlow += f
			} else {
				break
			}
		}
	}

	return
}

// 预流推进 HLPP 算法（High Level Preflow Push）
// todo https://oi-wiki.org/graph/flow/max-flow/#hlpp
//      https://www.luogu.com.cn/blog/ONE-PIECE/jiu-ji-di-zui-tai-liu-suan-fa-isap-yu-hlpp
// 模板题 https://www.luogu.com.cn/problem/P4722

// 最小费用最大流
// todo https://oi-wiki.org/graph/flow/min-cost/
// 模板题 https://www.luogu.com.cn/problem/P3381
