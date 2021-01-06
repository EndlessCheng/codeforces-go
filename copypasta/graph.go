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

TIPS: 使用一个 fa 数组（初始化为 -1）记录搜索树中的节点的父节点，这样对每个节点都有一条到根的路径（根的 fa 为 -1）
NOTE: 独立集相关问题，可以从染色的角度考虑
NOTE: 度数大于 √M 的点不超过 2√M 个

https://oeis.org/A031878 Maximal number of edges in Hamiltonian path in complete graph on n nodes
a(n) = C(n, 2)        n%2==0
a(n) = C(n, 2)-n/2+1  n%2==1

CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=graphs

环与独立集 https://codeforces.com/problemset/problem/1364/D
匹配与独立集 https://codeforces.com/problemset/problem/1198/C

构造 https://codeforces.com/problemset/problem/41/E
转换 https://codeforces.com/problemset/problem/788/B
第k小路径 https://codeforces.com/problemset/problem/1196/F

todo《挑战》例题+练习题
2.5 节 - 最短路 & 最小生成树
3255 https://www.luogu.com.cn/problem/P2865 次短路
3723 http://poj.org/problem?id=3723 建模+MST
3169 https://www.luogu.com.cn/problem/P4878 差分约束
2139 http://poj.org/problem?id=2139 Floyd
3259 https://www.luogu.com.cn/problem/P2850 多源 SPFA（建议读原文，洛谷翻译不完整）
3268 https://www.luogu.com.cn/problem/P1821 反图 Dij
https://onlinejudge.u-aizu.ac.jp/problems/2249 Dij 的过程中更新花费，注意距离相等时取花费最小值
https://onlinejudge.u-aizu.ac.jp/problems/2200 todo
1258 https://www.luogu.com.cn/problem/P1546 Prim
2377 http://poj.org/problem?id=2377 最大生成树
https://onlinejudge.u-aizu.ac.jp/problems/2224 为了让原图无环，需要去除不在最大生成树上的边
2395 https://www.luogu.com.cn/problem/P1547 最小生成树的最长边：Kruskal 中最后一条加入 MST 中的边的长度
3.5 节 - 二分图
3041
3057
1274
2112
1486
1466
3692
2724
2226
AOJ 2251
3.5节 - 网络流
最大流
3281
3469
3713
2987
2914
3155
最小费用流
2135
2175
3686
3680
3068
2195
3422
AOJ 2266
AOJ 2230
4.3 节 - SCC & 2SAT
2186
3683
3180
1236
3678
2723
2749
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

// EXTRA: 先染色，再递归 https://codeforces.com/problemset/problem/1470/D
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
		// https://codeforces.com/problemset/problem/1217/D 给一个有向图着色，使得没有一个环只有一个颜色，求使用的颜色数量的最小值
		// https://codeforces.com/problemset/problem/698/B
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
		// 无向图: DFS 找长度至少为 k 的环
		// 注：如果只有一个环（基环树），见 treeWithCycle
		// 模板题 https://codeforces.com/problemset/problem/263/D
		// https://codeforces.com/problemset/problem/1325/F
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
		// 构建深度数组/最短路
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
		// 全源最短路
		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := range dist[i] {
				dist[i][j] = -1
			}
			dist[i][i] = 0
			q := []int{i}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				for _, w := range g[v] {
					if dist[i][w] == -1 {
						dist[i][w] = dist[i][v] + 1
						q = append(q, w)
					}
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
// 半欧拉图：具有欧拉路径而无欧拉回路的图
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

	// 倒序变正序
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
// 模板题 LC1192 https://leetcode-cn.com/problems/critical-connections-in-a-network/
//       https://codeforces.com/problemset/problem/1000/E
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

type vdPair struct {
	v   int
	dis int64
}
type vdHeap []vdPair

func (h vdHeap) Len() int              { return len(h) }
func (h vdHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h vdHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *vdHeap) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *vdHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *vdHeap) push(v vdPair)        { heap.Push(h, v) }
func (h *vdHeap) pop() vdPair          { return heap.Pop(h).(vdPair) }

// 单源最短路 Dijkstra
// 适用于稀疏图 O(mlogm)
// 根据《算法(第4版)》，这里实现的是 lazy 版本的 Dijkstra，复杂度为 O(mlogm)；若在插入堆时元素已在堆中，改成更新元素而不是插入元素可使复杂度降为 O(mlogn)
// st 也可以是一个点集，这相当于同时对多个点跑最短路
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 模板题 https://www.luogu.com.cn/problem/P3371 https://www.luogu.com.cn/problem/P4779
//       https://codeforces.com/problemset/problem/20/C
// 最短路个数 https://www.luogu.com.cn/problem/P1608
// 通过最短路找到可以删除的边 https://codeforces.com/problemset/problem/449/B
// 稠密图 https://atcoder.jp/contests/arc064/tasks/arc064_c
// 建模题 https://www.luogu.com.cn/problem/P4644
// 最短路树上跑拓扑排序 https://codeforces.com/contest/1076/problem/D
// 基于 max LC1631 https://leetcode-cn.com/problems/path-with-minimum-effort/
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
	// 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
	vis := make([]bool, n)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = -1
	}
	q := vdHeap{{st, 0}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if vis[v] { // dist[v] < p.dis
			continue
		}
		vis[v] = true
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				fa[w] = v
				q.push(vdPair{w, newD})
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

	// EXTRA: 对于相邻的两点，记边为 e，若有：
	// abs(dist[v], dist[w]) == e.wt  =>  e 在最短路上（不带绝对值的话就有先后关系）
	// abs(dist[v], dist[w])  < e.wt  =>  e 不在最短路上（可以想象成一条鼓起的线）
	// 这里的最短路可以有多条

	// EXTRA: 次短路
	// 模板题 https://www.luogu.com.cn/problem/P2865
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
		q := vdHeap{{st, 0}}
		for len(q) > 0 {
			p := q.pop()
			v, d := p.v, p.dis
			if dist2[v] < d { // 注意是 dist2
				continue
			}
			for _, e := range g[v] {
				w := e.to
				newD := d + e.wt
				if newD < dist[w] {
					q.push(vdPair{w, newD})
					dist[w], newD = newD, dist[w]
				}
				if dist[w] < newD && newD < dist2[w] {
					dist2[w] = newD
					q.push(vdPair{w, newD})
				}
			}
		}
		_ = dist2[end]
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
// todo DFS 写法
// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
// EXTRA: 差分约束，若有 Xi-Xj<=Ck，则连一条有向边 j->i，边权为 Ck
//        然后再添加一个 0 号节点，向其他节点连一条边权为 0 的有向边，表示 Xi-X0<=0
//        这样，在无负环时会得到一组非正数解
//        模板题 https://www.luogu.com.cn/problem/P3385
//        模板题 https://www.luogu.com.cn/problem/P4878 todo 需要复习
//        https://www.luogu.com.cn/problem/SP116
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
	q := []int{st}
	inQ := make([]bool, n)
	inQ[st] = true
	relaxedCnt := make([]int, n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		inQ[v] = false
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				relaxedCnt[w] = relaxedCnt[v] + 1
				if relaxedCnt[w] >= n {
					return nil
				} // 找到一个从 st 出发可达的负环
				if !inQ[w] {
					inQ[w] = true
					q = append(q, w)
				}
			}
		}
	}

	// EXTRA: 只是找负环的话可以初始所有点入队
	// https://www.luogu.com.cn/problem/P2868

	// EXTRA: 打印负环
	// https://cp-algorithms.com/graph/finding-negative-cycle-in-graph.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BellmanFordSP.java.html

	return
}

// 任意两点最短路 Floyd-Warshall O(n^3)    本质是求 Min-plus matrix multiplication
// 传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// https://en.wikipedia.org/wiki/Min-plus_matrix_multiplication
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
		var v, w, wt int
		Fscan(in, &v, &m, &wt)
		dist[v][w] = wt
		dist[w][v] = wt
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
		dist[i] = append([]int64(nil), weights[i]...)
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

// 任意两点最短路 Johnson O(nmlogm)
// 若有负环返回 nil
// https://en.wikipedia.org/wiki/Johnson%27s_algorithm
// https://oi-wiki.org/graph/shortest-path/#johnson
// 模板题 https://www.luogu.com.cn/problem/P5905
func (G *graph) shortestPathJohnson(in io.Reader, n, m int) [][]int64 {
	const inf int64 = 1e18 // 1e9+1
	type neighbor struct {
		to int
		wt int64
	}
	g := make([][]neighbor, n+1)
	for i := 0; i < m; i++ {
		v, w, wt := 0, 0, int64(0)
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	// 建虚拟节点 0 并且往其他的点都连一条边权为 0 的边
	for v := 1; v <= n; v++ {
		g[0] = append(g[0], neighbor{v, 0})
		g[v] = append(g[v], neighbor{})
	}

	// 跑 SPFA，代码略（注意点数为 n+1）
	h := G.shortestPathSPFA(in, n+1, m, 0)
	if h == nil {
		return nil
	}

	// 求新的边权
	for v := 1; v <= n; v++ {
		for i, e := range g[v] {
			g[v][i].wt += h[v] - h[e.to]
		}
	}

	// 以每个点为源点跑一遍 Dijkstra，代码略（注意点数为 n+1）
	dist := make([][]int64, n+1)
	for st := 1; st <= n; st++ {
		dist[st] = G.shortestPathDijkstra(in, n+1, m, st)
		for end, d := range dist[st] {
			if d < inf {
				dist[st][end] -= h[st] - h[end]
			}
		}
	}
	return dist
}

// EXTRA: 同余最短路
// todo https://oi-wiki.org/graph/mod-shortest-path/
// todo 跳楼机 https://www.luogu.com.cn/problem/P3403

// k 短路
// A* 算法
// 可持久化可并堆优化
// https://en.wikipedia.org/wiki/K_shortest_path_routing
// todo https://oi-wiki.org/graph/kth-path/
// todo 模板题 https://www.luogu.com.cn/problem/P2483

// 最小斯坦纳树
// https://oi-wiki.org/graph/steiner-tree/
// todo 模板题 https://www.luogu.com.cn/problem/P6192

// 最小生成树 Kruskal
// 适用于稀疏图 O(mlogm)，或者边已经按权值排序的情况
// https://oi-wiki.org/graph/mst/#kruskal
// 模板题 https://www.luogu.com.cn/problem/P3366 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/E
// 题目推荐 https://cp-algorithms.com/graph/mst_kruskal.html#toc-tgt-5
// 关键边、伪关键边（与割边结合）https://codeforces.com/problemset/problem/160/D
// 最小生成树的最长边：Kruskal 中最后一条加入 MST 中的边的长度 https://www.luogu.com.cn/problem/P1547
// EXTRA: 与树链剖分结合可以在线查询两点间路径最大边权的最小值 https://leetcode-cn.com/contest/weekly-contest-220/problems/checking-existence-of-edge-length-limited-paths/
func (*graph) mstKruskal(in io.Reader, n, m int) int64 {
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
// 模板题 https://www.luogu.com.cn/problem/P1546
// 建模+打印方案 https://codeforces.com/contest/1245/problem/D
func (*graph) mstPrim(dist [][]int) int {
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
	minWeights[0] = 0 // 任选一点为起点
	used := make([]bool, n)
	ans := 0 // int64
	for {
		v := -1
		for i, u := range used {
			if !u && (v == -1 || minWeights[i] < minWeights[v]) {
				v = i
			}
		}
		if v == -1 {
			break
		}
		used[v] = true
		ans += minWeights[v]
		for w := range minWeights {
			minWeights[w] = min(minWeights[w], dist[v][w])
		}
	}
	return ans
}

// Boruvka's algorithm
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BoruvkaMST.java.html

// 次小生成树 Second best Minimum Spanning Tree
// Using Kruskal and Lowest Common Ancestor
// https://oi-wiki.org/graph/mst/#_9
// https://cp-algorithms.com/graph/second_best_mst.html
// todo 模板题 https://www.luogu.com.cn/problem/P4180
func (*graph) secondMST(n, m int) (sum int64) {
	return
}

// Kruskal 重构树
// https://oi-wiki.org/graph/mst/#kruskal_1

// 最小生成树计数 Kirchhoff's theorem

// 最小树形图 - 朱刘算法
// todo 模板题 https://www.luogu.com.cn/problem/P4716

// 反图的连通分量 O(n+m)
// https://www.luogu.com.cn/blog/endlesscheng/solution-cf1242b
// https://codeforces.com/contest/190/problem/E
// https://codeforces.com/contest/920/problem/E
// https://codeforces.com/contest/1243/problem/D
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
// https://en.wikipedia.org/wiki/Bipartite_graph
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

/* 匹配 带权匹配 独立集 边覆盖 顶点覆盖 支配集
https://en.wikipedia.org/wiki/Matching_(graph_theory)
https://en.wikipedia.org/wiki/Maximum_weight_matching
https://en.wikipedia.org/wiki/Independent_set_(graph_theory)
https://en.wikipedia.org/wiki/Edge_cover
https://en.wikipedia.org/wiki/Vertex_cover
https://en.wikipedia.org/wiki/Dominating_set
https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/

独立集+顶点覆盖 https://codeforces.com/problemset/problem/1470/D

最大匹配+最小边覆盖=n （图中无孤立点）
最大独立集+最小顶点覆盖=n https://www.geeksforgeeks.org/vertex-cover-problem-set-1-introduction-approximate-algorithm-2/
对于二分图，最小顶点覆盖=最大匹配，最大独立集=n-最大匹配

激光覆盖转换成最小顶点覆盖 http://poj.org/problem?id=3041
不是 n-匹配就是 n-独立集 https://codeforces.com/problemset/problem/1198/C

DAG 上的最小路径覆盖（挑战 p.272 / 进阶指南 p.436）：
    起初把每个点都视作一条路径，这样共有 n 条不相交路径
    拆点图：每个点拆成出点和入点，原图的 v->w 相当于拆点图的 出点v->入点w，这说明拆点图一定是二分图
    在拆点图里找到一条匹配边就相当于把两条路径合成了一条路径，也就相当于路径数减少了 1
    所以找到了多少匹配边，路径数就减少了多少
    所以有最小路径覆盖=原图的结点数-拆点图的最大匹配数
todo 树上最小路径覆盖 http://codeforces.com/problemset/problem/618/D
*/

// 二分图最大匹配 - 匈牙利算法/增广路算法 O(nm)    Hungarian algorithm
// 注：使用 Dinic 可以达到 O(m√n) 的复杂度
// https://www.renfei.org/blog/bipartite-matching.html 推荐
// https://oi-wiki.org/topic/graph-matching/bigraph-match/
// https://zhuanlan.zhihu.com/p/62981901
// https://www.geeksforgeeks.org/maximum-bipartite-matching/
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BipartiteMatching.java.html
// 模板题 https://www.luogu.com.cn/problem/P3386
func (*graph) maxBipartiteMatchingHungarian(n int, g [][]int) (match []int, cnt int) {
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
				cnt++ // +=2
			}
		}
	}
	return
}

// 匈牙利算法的另一种写法，适用左右两侧节点有明确区分的情况，要求 g 中存储的是左侧到右侧的单向边
func (*graph) maxBipartiteMatchingHungarianLR(nl, nr int, g [][]int) (matchL []int, cnt int) {
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
			cnt++ // +=2
		}
	}
	return
}

// 二分图最大匹配 -  Hopcroft–Karp 算法 O(m√n)
// https://en.wikipedia.org/wiki/Hopcroft%E2%80%93Karp_algorithm
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/HopcroftKarp.java.html
// todo http://pepcy.cf/icpc-templates/003Graph/hk.html
func (*graph) maxBipartiteMatchingHopcroftKarp(n int, g [][]int) (match []int, cnt int) {
	return
}

// 带权二分图最大匹配 - 任务分配问题/婚姻匹配问题 - KM(Kuhn–Munkres) 算法 O(n^4)  todo BFS 优化后的 O(n^3)
// https://en.wikipedia.org/wiki/Assignment_problem
// https://en.wikipedia.org/wiki/Hungarian_algorithm
// https://oi-wiki.org/topic/graph-matching/bigraph-weight-match/
// 直观理解 KM 算法 https://www.cnblogs.com/wenruo/p/5264235.html
// 讲解+题单 https://www.luogu.com.cn/blog/suxxsfe/xiong-ya-li-suan-fa
// https://resources.mpi-inf.mpg.de/departments/d1/teaching/ss12/AdvancedGraphAlgorithms/Slides06.pdf
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/AssignmentProblem.java.html
// 模板题 https://www.luogu.com.cn/problem/P6577
// https://www.luogu.com.cn/problem/P3967
// http://acm.hdu.edu.cn/showproblem.php?pid=2426
// EXTRA: 带权二分图最小边覆盖
// 转换成带权二分图最大匹配 https://cstheory.stackexchange.com/questions/14690/reducing-a-minimum-cost-edge-cover-problem-to-minimum-cost-weighted-bipartie-per
// https://leetcode-cn.com/problems/minimum-cost-to-connect-two-groups-of-points/solution/kai-kai-yan-jie-zhuan-huan-cheng-zui-da-dai-quan-p/
func (*graph) maxWeightedBipartiteMatchingKuhnMunkres(n int, wt [][]int) (match []int, sum int64) {
	const inf int = 2e9
	// NOTE: 若需要判断是否不存在完备匹配，wt 应初始化为 -inf，否则初始化为 0

	// 右部点匹配了哪一个左部点
	match = make([]int, n)
	for i := range match {
		match[i] = -1
	}
	// 顶标
	la := make([]int, n)
	for i, r := range wt {
		la[i] = r[0]
		for _, w := range r[1:] {
			if w > la[i] {
				la[i] = w
			}
		}
	}
	lb := make([]int, n)
	slack := make([]int, n)
	for i := 0; i < n; i++ {
		for {
			va := make([]bool, n)
			vb := make([]bool, n)
			for j := range slack {
				slack[j] = inf
			}
			var f func(int) bool
			f = func(v int) bool {
				va[v] = true
				for w, b := range vb {
					if !b {
						if d := la[v] + lb[w] - wt[v][w]; d == 0 {
							vb[w] = true
							if match[w] == -1 || f(match[w]) {
								match[w] = v
								return true
							}
						} else if d < slack[w] {
							slack[w] = d
						}
					}
				}
				return false
			}
			if f(i) {
				break
			}
			d := inf
			for j, b := range vb {
				if !b && slack[j] < d {
					d = slack[j]
				}
			}
			for j := 0; j < n; j++ {
				if va[j] {
					la[j] -= d
				}
				if vb[j] {
					lb[j] += d
				}
			}
		}
	}
	for w, v := range match {
		// 无解，或者不选
		if v == -1 {
			//continue
			return nil, 0
		}
		sum += int64(wt[v][w])
	}
	return
}

// 一般图（带权）最大匹配（带花树 Edmonds's blossom algorithm）    MWM
// https://en.wikipedia.org/wiki/Maximum_weight_matching
// https://en.wikipedia.org/wiki/Blossom_algorithm
// https://oi-wiki.org/topic/graph-matching/general-match/
// https://oi-wiki.org/topic/graph-matching/general-weight-match/
// TODO https://www.cnblogs.com/cjyyb/p/8719368.html 带花树算法学习笔记
// todo http://pepcy.cf/icpc-templates/003Graph/bls.html
// 模板题 https://www.luogu.com.cn/problem/P6113
//       https://www.luogu.com.cn/problem/P4258
//       https://www.luogu.com.cn/problem/P6699

// EXTRA: 完美匹配 Perfect Match     MWPM
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
// 构造 https://codeforces.com/problemset/problem/269/C
// 缩点后的拓扑序 https://codeforces.com/contest/1463/problem/E
// 拓扑序是否唯一：任意时刻队列中不能有超过一个元素
// 检查一个序列是否为拓扑序，可以仿造拓扑排序的算法，从前往后检查节点的入度是否为 0，然后减少相邻节点的入度，直到找到一个入度不为 0 的点或者遍历到末尾
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

			if deg[w]--; deg[w] == 0 {
				//fa[w] = v
				//levels[w] = levels[v] + 1
				q = append(q, w)
			}
		}
	}

	isDAG = len(orders) == n // 或所有 deg 均为 0

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
	// 注意这样可能会产生重边，不能有重边时可以用 map 或对每个点排序去重
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
// NOTE: 一些建边的转换：（¬x 用 x+n 表示）
//       A,B 至少存在一个 (A|B)    ¬A⇒B, ¬B⇒A 意思是一个为假的时候，另一个一定为真 https://www.luogu.com.cn/problem/P4782
//       A,B 不能同时存在 (¬A|¬B)  A⇒¬B, B⇒¬A 就是上面的式子替换了一下
//       A,B 必须且只一个 (A^B)    A⇒¬B, B⇒¬A, ¬A⇒B, ¬B⇒A
//       A,B 同时或都不在 (¬(A^B)) A⇒B, B⇒A, ¬A⇒¬B, ¬B⇒¬A
//       A 必须存在       (A)     ¬A⇒A
//       A 不能存在       (¬A)     A⇒¬A
// NOTE: 单独的条件 x为a 可以用 (x为a)∨(x为a) 来表示
// 下面的代码基于模板题 https://www.luogu.com.cn/problem/P4782
// todo 模板题 https://atcoder.jp/contests/practice2/tasks/practice2_h
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
		// x 和 ¬x 处于同一个 SCC 时无解，即 x ⇔ ¬x
		if id == sccIDs[i+n] {
			return nil
		}
		// sccIDs[x] > sccIDs[¬x] ⇔ (¬x ⇒ x) ⇔ x 为真
		// sccIDs[x] < sccIDs[¬x] ⇔ (x ⇒ ¬x) ⇔ x 为假
		ans[i] = id > sccIDs[i+n]
	}
	return ans
}

// 基环树（环套树），英文名叫 pseudotree，基环树森林叫 pseudoforest
// https://en.wikipedia.org/wiki/Pseudoforest
// 对于内向基环树，由于每个点的出度均为一，可以用 []int 来表示图
// todo https://www.luogu.com.cn/blog/user52918/qian-tan-ji-huan-shu
// https://codeforces.com/problemset/problem/1027/D
// https://codeforces.com/problemset/problem/1335/F
// todo [IOI2008] 岛屿 https://www.luogu.com.cn/problem/P4381
// todo [NOI2013] 快餐店 https://www.luogu.com.cn/problem/P1399
func (*graph) pseudotree(n int, g []int, rg [][]int, inDeg []int) {
	// 输入：g 为内向基环树，rg 为反图（外向基环树），inDeg 为 g 的入度，在读入时计算出

	// 拓扑排序，之后 inDeg 为 1 的点必定在基环上
	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w := g[v]
		if inDeg[w]--; inDeg[w] == 0 {
			q = append(q, w)
		}
	}

	// 在反图上遍历树枝
	var f func(int)
	f = func(v int) {
		for _, w := range rg[v] {
			if inDeg[w] == 0 {
				f(w)
			}
		}
	}
	// 遍历基环
	for root, d := range inDeg {
		if d > 0 {
			f(root)
		}
	}

	// EXTRA: 求基环
	var cycle []int
	for i, d := range inDeg {
		if d > 0 {
			for v := i; ; v = g[v] {
				cycle = append(cycle, v)
				if g[v] == i {
					break
				}
			}
			break
		}
	}

	{
		// EXTRA: 无向图的情况
		g := [][]int{}
		deg := []int{}
		// read ...

		q := []int{}
		for i, d := range deg {
			if d == 1 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, w := range g[v] {
				if deg[w]--; deg[w] == 1 {
					q = append(q, w)
				}
			}
		}

		// 遍历树枝
		var f func(v, fa int) int
		f = func(v, fa int) int {
			size := 1
			for _, w := range g[v] {
				if w != fa && deg[w] < 2 {
					sz := f(w, v)
					// do sz ...
					size += sz
				}
			}
			return size
		}
		// 遍历基环
		size := 0
		for root, d := range deg {
			if d > 1 {
				sz := f(root, -1)
				// do sz ...
				size += sz
			}
		}

		// EXTRA: 求基环
		var cycle []int
		for i, d := range deg {
			if d > 1 {
				pre, v := -1, i
				for {
					cycle = append(cycle, v)
					for _, w := range g[v] {
						if w != pre && deg[w] > 1 {
							pre, v = v, w
							break
						}
					}
					if v == i {
						break
					}
				}
				break
			}
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

Disjoint paths
Edge-disjoint paths: It turns out that the maximum number of edge-disjoint paths equals the maximum flow of the graph, assuming that the capacity of each edge is one.
Node-disjoint paths: 拆点法

路径覆盖 Path cover + 打印
todo Competitive Programmer’s Handbook Ch.20
todo 线性规划与网络流 24 题 - 最小路径覆盖问题 https://byvoid.com/zhs/blog/lpf24-3/

全局最小割 Stoer-Wagner 算法 O(nm+n^2logn)
https://en.wikipedia.org/wiki/Stoer%E2%80%93Wagner_algorithm
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GlobalMincut.java.html
模板题 https://www.luogu.com.cn/problem/P5632

最大权闭合子图
*/

/* 网络流建模技巧/转换技巧
标准建模（指派问题）：
	http://poj.org/problem?id=2175
	http://poj.org/problem?id=3686
顶点上有容量限制：
	可以将其拆成两个，边容量为顶点容量
有最小容量限制-最大流（挑战 p.214）：
	增加源点 S 和汇点 T，修改每条边 v-w 的 cap-=minCap，并 S->v 连边 cap=minCap，w->T 连边 cap=minCap
	检查最大流是否为 ΣminCap（不是满流无解）
	S->st 连边 cap=inf，end->T 连边 cap=inf
	最大流为 maxFlow-ΣminCap
有最小容量限制-最小费用最大流（挑战 p.227）：
	对每条边新增一条边 e'
	e.cap-=minCap
	e'.cap=minCap
	e'.cost=e.cost-M // 一个足够大的常数
	跑完 MCMF 后加上 M*ΣminCap
边容量增加：
	重新跑增广路
边容量减少：
	若 flow<=cap' 则最大流不变；若 flow>cap' 需要将多出的流退回去 todo
流量任意：
	todo
容量为负数：
	todo
费用为负数：
	todo 挑战 p.228
一对多的最大匹配：
	Dining https://www.luogu.com.cn/problem/P2891 http://poj.org/problem?id=3281
求最小割划分成两个集合：
	Dual Core CPU http://poj.org/problem?id=3469
无重复边的往返最短路：
	http://poj.org/problem?id=2135
	转换成流量为 2 的最小费用流

另见下面的 minCostFlowSPFA

最小割模型
主要参考胡伯涛《最小割模型在信息学竞赛中的应用》（PDF 在 misc 文件夹下）
https://blog.csdn.net/qq_35649707/article/details/77482691
最大权闭合图 Maximum Weight Closure of a Graph
最大密度子图 Maximum Density Subgraph
二分图最小点权覆盖集/最大点权独立集 Minimum Weight Vertex Covering Set (MinWVCS) and Maximum Weight Vertex Independent Set (MaxWVIS) in a Bipartite Graph
   建立一个源 s，向 X 部每个点连边；建立一个汇 t，从 Y 部每个点向汇 t 连边，把二分图中的边看成是有向的，
   则任意一条从 s 到 t 的路径，一定具有 s-v-w-t 的形式（v∈X, w∈Y）。
   割的性质是不存在一条从 s 到 t 的路径。故路径上的三条边 s-v, v-w, w-t 中至少有一条边在割中。
   若人为地令 v-w 不可能在最小割中，即令其容量为正无限，
   可将条件简化为 s-v, w-t 中至少有一条边在最小割中，这正好与点覆盖集限制条件的形式相符（边的两端点中至少一个在覆盖集内），
   而目标是最小化点权之和，这恰好也是最小割的优化目标。
   对于最大点权独立集，其等价于点权之和减去最小点权覆盖集。
   https://codeforces.com/contest/808/problem/F

点边转换
   将点拆为入点和出点（v 和 v+n），即可把点的属性变成边的属性，从而方便应用最大流、最小割等算法
   将边的中间加一个节点，把边的属性体现在中间的点上

todo 线性规划与网络流 24 题 解题报告 https://byvoid.com/zhs/blog/lpf24-solution/

todo 题单！https://www.zybuluo.com/xzyxzy/note/992041
  网络流从入门到入土 #1 https://www.luogu.com.cn/training/12097#problems
  网络流从入门到入土 #2 https://www.luogu.com.cn/training/12098#problems
  网络流从入门到入土 #3 https://www.luogu.com.cn/training/12099#problems
  网络流建模经典题 https://www.luogu.com.cn/training/1230#problems
  网络流经典题目 https://www.luogu.com.cn/training/3144#problems

CF Tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=flows
*/

// 最大流 Dinic's algorithm O(n^2 * m)  二分图上为 O(m√n)
// Ford–Fulkerson algorithm (FFA) 的改进版本
// https://oi-wiki.org/graph/flow/max-flow/#dinic
// https://cp-algorithms.com/graph/dinic.html
// 模板题 https://www.luogu.com.cn/problem/P3376
// 建模题 https://atcoder.jp/contests/arc085/tasks/arc085_c
//       https://codeforces.com/problemset/problem/1360/G (满流时有解)
//       https://codeforces.com/problemset/problem/546/E (满流时有解)
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
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0}) // 无向图上 0 改 cap
	}
	for i := 0; i < m; i++ {
		var v, w, cap int
		Fscan(in, &v, &w, &cap)
		v--
		w--
		addEdge(v, w, cap)
	}

	// 计算从源点 st 出发的距离
	d := make([]int, n)
	bfs := func() bool {
		for i := range d {
			d[i] = -1
		}
		d[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && d[w] < 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d[end] >= 0
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
			if w := e.to; e.cap > 0 && d[w] > d[v] {
				if f := dfs(w, min(minF, e.cap)); f > 0 {
					e.cap -= f
					g[w][e.rid].cap += f
					return f
				}
			}
		}
		return 0
	}

	const inf int = 1e9 // 1e18
	maxFlow := 0        // int64
	for bfs() {
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
	//w := e.to
	//ans[v][w] = g[w][e.rid].cap

	return maxFlow
}

// ISAP, Improved Shortest Augmenting Path O(n^2 * m)
// https://oi-wiki.org/graph/flow/max-flow/#isap
// https://www.renfei.org/blog/isap.html
// todo 目前来看性能比 Dinic 好一点点，需要找数据量更大的做对比
// todo 非二分图上 Dinic 据说更快？
func (*graph) maxFlowISAP(in io.Reader, n, m, st, end int) int {
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

	// 计算从汇点 end 出发的距离
	d := make([]int, n)
	for i := range d {
		d[i] = -1
	}
	d[end] = 0
	cd := make([]int, n+1) // 注意有 d[i] == n 的情况
	q := []int{end}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		cd[d[v]]++
		for _, e := range g[v] {
			if w := e.to; d[w] < 0 {
				d[w] = d[v] + 1
				q = append(q, w)
			}
		}
	}
	if d[st] < 0 {
		return -1
	}

	// 寻找增广路
	const inf int = 1e9 // 1e18
	maxFlow := 0        // int64
	iter := make([]int, n)
	type pair struct{ v, i int }
	fa := make([]pair, n)
o:
	for v := st; d[st] < n; {
		if v == end {
			minF := inf
			for v := end; v != st; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := end; v != st; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			v = st
		}
		for i := iter[v]; i < len(g[v]); i++ {
			e := g[v][i]
			if w := e.to; e.cap > 0 && d[w] < d[v] {
				fa[w] = pair{v, i}
				iter[v] = i
				v = w
				continue o
			}
		}
		if cd[d[v]] == 1 {
			break // gap 优化
		}
		cd[d[v]]--
		minD := n - 1
		for _, e := range g[v] {
			if e.cap > 0 && d[e.to] < minD {
				minD = d[e.to]
			}
		}
		d[v] = minD + 1
		cd[d[v]]++
		iter[v] = 0
		if v != st {
			v = fa[v].v
		}
	}
	return maxFlow
}

// 最高标号预流推进 (HLPP, High Level Preflow Push)   O(n^2 * √m)   复杂度上界相比 Dinic/ISAP 比较紧
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm#Highest_label_selection_rule
// https://oi-wiki.org/graph/flow/max-flow/#hlpp
// 模板题 https://www.luogu.com.cn/problem/P4722
// todo deque 优化 + 全局重贴标签等 https://www.luogu.com.cn/problem/solution/P4722
type dh struct {
	sort.IntSlice
	d []int
}

func (h dh) Less(i, j int) bool  { return h.d[h.IntSlice[i]] > h.d[h.IntSlice[j]] } // 处于堆中的节点的 d 值不会改变，所以可以直接比较
func (h *dh) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *dh) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *dh) push(v int)         { heap.Push(h, v) }
func (h *dh) pop() int           { return heap.Pop(h).(int) }

func (*graph) maxFlowHLPP(in io.Reader, n, m, st, end int) int {
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

	// 计算从汇点 end 出发的距离
	d := make([]int, n)
	for i := range d {
		d[i] = -1
	}
	d[end] = 0
	cd := make([]int, 2*n)
	_q := []int{end}
	for len(_q) > 0 {
		v := _q[0]
		_q = _q[1:]
		cd[d[v]]++
		for _, e := range g[v] {
			if w := e.to; d[w] < 0 {
				d[w] = d[v] + 1
				_q = append(_q, w)
			}
		}
	}
	if d[st] < 0 {
		return -1
	}
	d[st] = n

	exFlow := make([]int, n)
	q := dh{d: d}
	inQ := make([]bool, n)
	push := func(v, f int, e *neighbor) {
		w := e.to
		e.cap -= f
		g[w][e.rid].cap += f
		exFlow[v] -= f
		exFlow[w] += f
		if w != st && w != end && !inQ[w] {
			q.push(w)
			inQ[w] = true
		}
	}
	// 将源点的所有边都满流地推送出去
	for i := range g[st] {
		if e := &g[st][i]; e.cap > 0 {
			push(st, e.cap, e)
		}
	}
	for len(q.IntSlice) > 0 {
		v := q.pop()
		inQ[v] = false
	o:
		for {
			for i := range g[v] {
				if e := &g[v][i]; e.cap > 0 && d[e.to] < d[v] {
					push(v, min(e.cap, exFlow[v]), e)
					if exFlow[v] == 0 {
						break o
					}
				}
			}
			dv := d[v]
			cd[dv]--
			if cd[dv] == 0 { // gap 优化
				for i, h := range d {
					if i != st && i != end && dv < h && h <= n {
						d[i] = n + 1 // 超过 n，从而尽快将流量推回 st
					}
				}
			}
			// relabel
			minD := int(1e9)
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && d[w] < minD {
					minD = d[w]
				}
			}
			d[v] = minD + 1
			cd[d[v]]++
		}
	}
	return exFlow[end]
}

// 最小费用流 最小费用最大流 MCMF
// 有两种实现：SPFA O(fnm) 和 Dijkstra O(fmlogn)
// 要求图中无负圈
// https://oi-wiki.org/graph/flow/min-cost/
// 性能对比（由于数据不强所以 SPFA 很快）：SPFA 1.05s(max 365ms)   Dijkstra 1.91s(max 688ms)
// 模板题 https://www.luogu.com.cn/problem/P3381
// 常用的建模思路是
// - 从源点 (0) 连 len(A) 条容量为 1，费用为 0 的边到集合 A，
// - 从集合 B 连 len(B) 条容量为 1，费用为 0 的边到汇点 (len(A)+len(B)+1)
// - 集合 A B 之间连边，容量为 inf，费用为 f(Ai,Bi)，f 根据题意
// - 这样跑 MCMF 得到的结果是匹配全部 A 或者 B 的最小花费
// 例如 https://codeforces.com/problemset/problem/1437/C
// 流通问题 circulation problem
// 最小费用流通问题 minimum-cost-circulation problem
//     https://en.wikipedia.org/wiki/Circulation_problem
//     The circulation problem and its variants are a generalisation of network flow problems,
//     with the added constraint of a lower bound on edge flows,
//     and with flow conservation also being required for the source and sink (i.e. there are no special nodes).
//    《算法导论》思考题 29-5
//     https://codeforces.com/blog/entry/85186?#comment-728533
func (*graph) minCostFlowSPFA(in io.Reader, n, m, st, end, flowLimit int) (int, int64) {
	// st--; end--

	const inf int64 = 1e18
	type neighbor struct{ to, rid, cap, cost int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost}) // 无向图上 0 改 cap
	}
	for i := 0; i < m; i++ {
		var v, w, cap, cost int
		Fscan(in, &v, &w, &cap, &cost)
		v--
		w--
		addEdge(v, w, cap, cost)
	}

	// n = len(A)+len(B)+2  or  end+1
	dist := make([]int64, n)
	type pair struct{ v, i int }
	fa := make([]pair, n)
	spfa := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		inQ := make([]bool, n)
		inQ[st] = true
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + int64(e.cost); newD < dist[w] {
					dist[w] = newD
					fa[w] = pair{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[end] < inf
	}
	maxFlow, minCost := 0, int64(0)
	for maxFlow < flowLimit && spfa() { // 若求 MCMF，把 maxFlow < flowLimit 去掉
		// 沿 st-end 的最短路尽量增广
		minF := flowLimit // inf
		for v := end; v != st; {
			p := fa[v]
			if c := g[p.v][p.i].cap; c < minF {
				minF = c
			}
			v = p.v
		}
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		maxFlow += minF
		minCost += dist[end] * int64(minF)
	}
	if maxFlow < flowLimit {
		return -1, -1
	}
	return maxFlow, minCost
}

// 基于原始对偶方法 (primal-dual method)
// https://blog.xehoth.cc/DurationPlan-Primal-Dual/
func (*graph) minCostFlowDijkstra(in io.Reader, n, m, st, end, flowLimit int) int64 {
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
		q := vdHeap{{st, 0}}
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
					q.push(vdPair{w, newD})
				}
			}
		}
		return dist[end] < inf
	}
	minCost := int64(0)
	for flowLimit > 0 && dijkstra() {
		for i, d := range dist {
			h[i] += d
		}
		minF := flowLimit // inf
		for v := end; v != st; {
			p := fa[v]
			if c := g[p.v][p.i].cap; c < minF {
				minF = c
			}
			v = p.v
		}
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		flowLimit -= minF               // maxFlow += minF
		minCost += h[end] * int64(minF) // 注意这里是 h 不是 dist
	}
	if flowLimit > 0 {
		return -1
	}
	return minCost
}

// todo 基于 Capacity Scaling 的弱多项式复杂度最小费用流算法 https://ouuan.github.io/post/%E5%9F%BA%E4%BA%8E-capacity-scaling-%E7%9A%84%E5%BC%B1%E5%A4%9A%E9%A1%B9%E5%BC%8F%E5%A4%8D%E6%9D%82%E5%BA%A6%E6%9C%80%E5%B0%8F%E8%B4%B9%E7%94%A8%E6%B5%81%E7%AE%97%E6%B3%95/

// ZKW 费用流
// https://artofproblemsolving.com/community/c1368h1020435

//

// 支配树
// todo 模板题 https://www.luogu.com.cn/problem/P5180

// 弦图：任意长度大于 3 的环都有一个弦（连接环中不相邻两点的边）的图称为弦图
// 单纯点 完美消除序列
// 最大势算法 Maximum Cardinality Search (MCS) http://www.ii.uib.no/~pinar/MCS-M.pdf
// https://oi-wiki.org/graph/chord/
// https://www.luogu.com.cn/blog/hsfzLZH1/chord-graph

// 寻找一个子图，要么其是一个 k-团，要么其每个顶点都至少有 k 个邻居（度不小于 k）
// https://codeforces.com/contest/1439/problem/B
func (*graph) findPseudoClique(g []map[int]bool, k int) []int {
	n := len(g)
	deg := make([]int, n)
	for i, vs := range g {
		deg[i] = len(vs)
	}

	left := n
	del := make([]bool, n)
	q := []int{}
	for i, d := range deg {
		if d < k {
			del[i] = true
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		left--
		if deg[v] == k-1 {
			// 检查 v 和它的小伙伴们(邻居)能否组成 k-团
			clique := []int{}
			for w := range g[v] {
				if deg[w] >= k-1 {
					clique = append(clique, w)
				}
			}
			if len(clique) == k-1 {
				for i, v := range clique {
					for _, w := range clique[:i] {
						if !g[v][w] {
							goto deleteV
						}
					}
				}
				clique = append(clique, v)
				return clique
			}
		}
	deleteV:
		deg[v] = 0
		for w := range g[v] {
			if deg[w]--; deg[w] < k && !del[w] {
				del[w] = true
				q = append(q, w)
			}
		}
	}

	if left == 0 {
		return nil
	}

	pseudoClique := make([]int, 0, left)
	for i, d := range del {
		if !d {
			pseudoClique = append(pseudoClique, i)
		}
	}
	return pseudoClique
}

// todo 最大团

// todo 极大团计数

// todo 图的同构

// todo 树的同构
// AHU 算法
// https://oi-wiki.org/graph/tree-ahu/

//

// 支配树
// todo 模板题 https://www.luogu.com.cn/problem/P5180

// 弦图：任意长度大于 3 的环都有一个弦（连接环中不相邻两点的边）的图称为弦图
// 单纯点 完美消除序列
// 最大势算法 Maximum Cardinality Search (MCS) http://www.ii.uib.no/~pinar/MCS-M.pdf
// https://oi-wiki.org/graph/chord/
// https://www.luogu.com.cn/blog/hsfzLZH1/chord-graph

// todo 最大团

// todo 极大团计数

// todo 图的同构

// todo 树的同构
// AHU 算法
// https://oi-wiki.org/graph/tree-ahu/
