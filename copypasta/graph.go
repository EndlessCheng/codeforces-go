package copypasta

import (
	"container/heap"
	. "fmt"
	"io"
	"math"
	"math/bits"
	"sort"
)

/*
Graph Theory Playlist https://www.youtube.com/playlist?list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P
图论的小技巧以及扩展 https://www.luogu.com.cn/blog/chengni5673/tu-lun-di-xiao-ji-qiao-yi-ji-kuo-zhan

边权转点权：在 v-w 之间加一个点，这个点的点权就是原来的边权（原图的点的点权视作 0）
点权转边权：将一个点拆分成两个点，用一条边连起来，新边的边权就是该点的点权（原图的边的边权视作 0）
其它情况：也可以用 min/max 等价转换 http://codeforces.com/problemset/problem/915/F

TIPS: 使用一个 fa 数组（初始化为 -1）记录搜索树中的节点的父节点，这样对每个节点都有一条到根的路径（根的 fa 为 -1）
NOTE: 独立集相关问题，可以从染色的角度考虑
NOTE: 度数大于 √M 的点不超过 2√M 个
      相关题目 & 无向图定向 https://leetcode-cn.com/problems/minimum-degree-of-a-connected-trio-in-a-graph/solution/gei-wu-xiang-tu-ding-xiang-by-lucifer100-c72d/

https://oeis.org/A031878 Maximal number of edges in Hamiltonian path in complete graph on n nodes
a(n) = C(n, 2)        n%2==0
a(n) = C(n, 2)-n/2+1  n%2==1

环与独立集 https://codeforces.com/problemset/problem/1364/D
匹配与独立集 https://codeforces.com/problemset/problem/1198/C

建图 https://codeforces.com/problemset/problem/1635/E
归纳 https://codeforces.com/problemset/problem/412/D
构造 https://codeforces.com/problemset/problem/41/E
转换 https://codeforces.com/problemset/problem/788/B
转换 https://codeforces.com/problemset/problem/788/C
加边 https://codeforces.com/problemset/problem/723/E
第k小路径 https://codeforces.com/problemset/problem/1196/F
给一无向图，从中删除恰好一条边，求可以让图变成二分图的所有边的下标 https://codeforces.com/problemset/problem/19/E
倒水问题 https://www.luogu.com.cn/problem/P1432
顶点有限制的生成树 https://codeforces.com/problemset/problem/723/F

Trémaux tree https://en.wikipedia.org/wiki/Tr%C3%A9maux_tree
DFS 树与 BFS 树 https://atcoder.jp/contests/abc251/tasks/abc251_f
证明 https://atcoder.jp/contests/abc251/editorial/3987

奇妙 BFS https://codeforces.com/problemset/problem/1651/D

竞赛图
竞赛图的一些性质 https://www.cnblogs.com/acha/p/9042984.html
- SCC 的拓扑序是唯一的
- 拓扑序上，不同 SCC 的点的入度，越靠前的严格越小
https://codeforces.com/problemset/problem/1498/E
https://codeforces.com/problemset/problem/1514/E
todo 竞赛图与三元环 https://codeforces.com/problemset/problem/117/C

定义连通性
https://codeforces.com/problemset/problem/1689/E

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

// https://atcoder.jp/contests/arc111/tasks/arc111_b
// EXTRA: 先染色，再递归 https://codeforces.com/problemset/problem/1470/D
// 无向图后向边定向 https://codeforces.com/problemset/problem/1519/E
// https://codeforces.com/problemset/problem/1176/E
// 与 MST 结合 https://codeforces.com/problemset/problem/1707/C
func (*graph) dfs(n, st int, g [][]int) {
	vis := make([]bool, n)
	var cntV, cntE int
	var f func(int)
	f = func(v int) {
		vis[v] = true
		cntV++
		cntE += len(g[v])
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i, b := range vis {
		if !b { // && len(g[i]) > 0
			cntV, cntE = 0, 0
			f(i)      // 注意自环和重边
			cntE /= 2 // 无向图
			if cntV-1 == cntE {
				// 树
			} else {
				// 有环
			}
		}
	}

	{
		// 奇偶标记法
		// https://codeforces.com/problemset/problem/936/B
		vis := make([][2]bool, n)
		var f func(int, int8)
		f = func(v int, step int8) {
			vis[v][step] = true
			// ...
			for _, w := range g[v] {
				if !vis[w][step^1] {
					f(w, step^1)
				}
			}
		}
		f(st, 0)
	}

	{
		// 欧拉序列
		eulerPath := []int{}
		vis := make([]bool, n)
		var f func(int)
		f = func(v int) {
			eulerPath = append(eulerPath, v)
			vis[v] = true
			for _, w := range g[v] {
				if !vis[w] {
					f(w)
					eulerPath = append(eulerPath, v)
				}
			}
		}
		f(st)
	}

	{
		// 有向图的环/回边检测/012染色
		//《算法导论》p.353 边的分类
		// vis[v] == 0：该顶点未被访问
		// vis[v] == 1：该顶点已经被访问，其子树未遍历完
		// vis[v] == 2：该顶点已经被访问，其子树已遍历完
		// LC802 https://leetcode-cn.com/problems/find-eventual-safe-states/
		// http://codeforces.com/problemset/problem/25/D
		// https://codeforces.com/problemset/problem/698/B
		// https://codeforces.com/problemset/problem/936/B
		// https://codeforces.com/problemset/problem/1217/D 给一个有向图着色，使得没有一个环只有一个颜色，求使用的颜色数量的最小值
		// https://codeforces.com/problemset/problem/1547/G
		color := make([]int8, n)
		var f func(int)
		f = func(v int) {
			color[v] = 1
			for _, w := range g[v] {
				if c := color[w]; c == 0 { // 未访问过，即 DFS 树上的树边【树枝边】
					f(w)
				} else if c == 1 { // 后向边，说明有环

				} else { // 前向边或横向边，说明有多条路径可以到 w

				}
			}
			color[v] = 2
		}
		for i, c := range color {
			if c == 0 {
				f(i)
			}
		}
	}

	{
		// 无向图分类：无环/自环/一般环
		// https://codeforces.com/contest/1770/problem/D
		c := 0 // 默认：无环
		var f func(int, int)
		f = func(v, fa int) {
			vis[v] = true
			for _, w := range g[v] {
				if w != fa {
					if w == v {
						// 自环
						c = 1
					} else if vis[w] { // 返祖边或者横向边（v 连向不在子树 v 上的点 w）
						// 一般环
						c = 2
					} else { // 树枝边
						f(w, v)
					}
				}
			}
		}
		_ = c
		f(0, -1)
	}

	{
		// 无向图: DFS 找长度至少为 k 的环
		// 注：如果只有一个环（基环树），见 pseudotree
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

	// 其它找环题目
	// https://codeforces.com/contest/1817/problem/B
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

// BFS
// 基础题 https://leetcode.cn/problems/keys-and-rooms/
// 建模 https://codeforces.com/problemset/problem/1272/E
// 锻炼分类讨论能力 https://codeforces.com/contest/1790/problem/G
// 带撤销的 BFS https://codeforces.com/problemset/problem/1721/D
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

// 字典序最小最短路
// 入门经典第二版 p.173
// 理想路径（NEERC10）https://codeforces.com/gym/101309 I 题
// 从终点倒着 BFS 求最短路，然后从起点开始一层一层向终点走，每一步都选颜色最小的，并记录最小颜色对应的所有节点，供下一层遍历
// 如果求的是字典序最小的顶点，每一步需选择符合 dis[w] == dis[v]-1 的下标最小的顶点
// LC499 https://leetcode.cn/problems/the-maze-iii/
func (*graph) lexicographicallySmallestShortestPath(g [][]struct{ to, color int }, st, end int) []int {
	const inf int = 1e9
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = inf
	}
	dis[end] = 0
	q := []int{end}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			if w := e.to; dis[v]+1 < dis[w] {
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
	}

	if dis[st] == inf {
		return nil
	}

	colorPath := []int{}
	check := []int{st}
	inC := make([]bool, len(g))
	inC[st] = true
	for loop := dis[st]; loop > 0; loop-- {
		minC := inf
		tmp := check
		check = nil
		for _, v := range tmp {
			for _, e := range g[v] {
				if w, c := e.to, e.color; dis[w] == dis[v]-1 {
					if c < minC {
						for _, w := range check {
							inC[w] = false
						}
						minC, check, inC[w] = c, []int{w}, true
					} else if c == minC && !inC[w] {
						check = append(check, w)
						inC[w] = true
					}
				}
			}
		}
		colorPath = append(colorPath, minC)
	}
	return colorPath
}

// BFS 应用：求无向无权图最小环长度
// 好题 https://codeforces.com/problemset/problem/1325/E
// LC2608 https://leetcode.cn/problems/shortest-cycle-in-a-graph/
/* 注意不能提前推出（哪怕是遍历完一个找到环的点的所有邻居）
0 3
0 5
3 4
4 5
1 9
1 11
9 10
11 10
2 6
2 8
6 7
8 7
0 1
0 2
1 2
*/
func (*graph) shortestCycleBFS(n int, g [][]int, min func(int, int) int) int {
	const inf int = 1e9
	ans := inf
	dis := make([]int, n)
	for st := range g {
		for i := range dis {
			dis[i] = -1
		}
		dis[st] = 0
		type pair struct{ v, fa int }
		q := []pair{{st, -1}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			v, fa := p.v, p.fa
			for _, w := range g[v] {
				if dis[w] == -1 {
					dis[w] = dis[v] + 1
					q = append(q, pair{w, v})
				} else if w != fa {
					ans = min(ans, dis[w]+dis[v]+1)
				}
			}
		}
	}
	return ans
}

// 欧拉图（欧拉回路）   半欧拉图（欧拉路径）
// 半欧拉图：具有欧拉路径而无欧拉回路的图
// 判别法如下 https://oi-wiki.org/graph/euler/#_3
// 无向图-欧拉回路：连通且没有奇度数点
// 无向图-欧拉路径：连通且恰有 0 或 2 个奇度数点（若有则选择其中一奇度数点为起点）
// 有向图-欧拉回路：SCC 只有一个且每个点的入度和出度相同
// 有向图-欧拉路径：1. 对应的无向图是连通的；2. 若每个点的入度和出度相同则起点任意；否则起点的出度比入度多一，终点的入度比出度多一，且其余点的入度和出度相同
//
// 逐步插入回路法（Hierholzer 算法）https://oi-wiki.org/graph/euler/
// todo 混合图欧拉回路
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/EulerianCycle.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/EulerianPath.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/DirectedEulerianCycle.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/DirectedEulerianPath.java.html
// https://algs4.cs.princeton.edu/42digraph/DirectedEulerianCycle.java.html
// NOTE: 递归前对边排序可保证输出的是字典序最小的路径
// 模板题（输出顶点）
// - 无向图 https://www.luogu.com.cn/problem/P2731 https://www.luogu.com.cn/problem/P1341
// - 有向图 https://www.luogu.com.cn/problem/P7771 LC332 https://leetcode-cn.com/problems/reconstruct-itinerary/solution/javadfsjie-fa-by-pwrliang/
// 模板题（输出边）
// - 有向图 LC2097 https://leetcode-cn.com/problems/valid-arrangement-of-pairs/
// 构造 https://ac.nowcoder.com/acm/contest/4010/H
// 构造 https://codeforces.com/problemset/problem/1511/D
// 虚点 https://codeforces.com/problemset/problem/723/E
// 转换 https://codeforces.com/problemset/problem/1361/C
// https://codeforces.com/problemset/problem/1186/F
func (*graph) eulerianPathOnUndirectedGraph(n, m int) []int {
	// 无向图
	type neighbor struct{ to, eid int }
	g := make([][]neighbor, n)
	// read g ...

	// 排序，保证字典序最小
	for _, es := range g {
		sort.Slice(es, func(i, j int) bool { return es[i].to < es[j].to })
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
	path := make([]int, 0, len(g)) // m
	vis := make([]bool, m)
	var f func(int)
	f = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			i := e.eid
			if vis[i] {
				continue
			}
			vis[i] = true
			w := e.to
			f(w)
			// 输出边的写法，注意是倒序
			// path = append(path, i)
		}
		// 输出点的写法，最后需要反转 path
		path = append(path, v)
	}
	f(st) // for i := range g { f(i) }

	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}

	return path
}

func (*graph) eulerianPathOnDirectedGraph(n, m int) []int {
	// 有向图
	type neighbor struct{ to, eid int }
	g := make([][]neighbor, n)
	inDeg := make([]int, n) // 统计入度
	// read g ...

	// 排序，保证字典序最小
	for _, es := range g {
		sort.Slice(es, func(i, j int) bool { return es[i].to < es[j].to })
	}

	st := -1
	end := -1
	for i, es := range g {
		if len(es) == inDeg[i]+1 { // 出度比入度大一，为起点
			if st >= 0 {
				return nil // 无欧拉路径
			}
			st = i
			//break // 如果保证有欧拉路径就直接 break
		}
		if len(es)+1 == inDeg[i] { // 入度比出度大一，为终点
			if end >= 0 {
				return nil // 无欧拉路径
			}
			end = i
			//break
		}
	}
	if st < 0 {
		st = 0 // 任选一起点（比如字典序最小），此时返回的是欧拉回路
	}

	path := make([]int, 0, m+1)
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

	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
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
// LC928 https://leetcode-cn.com/problems/minimize-malware-spread-ii/
func (*graph) findCutVertices(n int, g [][]int, min func(int, int) int) (isCut []bool) {
	isCut = make([]bool, n)
	dfn := make([]int, n) // 值从 1 开始
	dfsClock := 0
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
			} else if w != fa { // 找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 { // 特判：只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割顶
			isCut[v] = false
		}
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
// 与最短路结合 https://codeforces.com/problemset/problem/567/E
// https://codeforces.com/problemset/problem/118/E
// todo 构造 https://codeforces.com/problemset/problem/550/D
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
	isBridge = make([]bool, len(edges))
	dfn := make([]int, len(g)) // 值从 1 开始
	dfsClock := 0
	var f func(int, int) int
	f = func(v, fid int) int { // 使用 fid 而不是 fa，可以兼容重边的情况
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		for _, e := range g[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := f(w, e.eid)
				if lowW > dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 或 v 的祖先，所以 v-w 必定是桥
					isBridge[e.eid] = true
				}
				lowV = min(lowV, lowW)
			} else if e.eid != fid { // 找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			f(v, -1)
		}
	}

	// EXTRA: 所有桥边的下标
	bridgeEIDs := []int{}
	for eid, b := range isBridge {
		if b {
			bridgeEIDs = append(bridgeEIDs, eid)
		}
	}

	return
}

// 无向图的双连通分量 Biconnected Components (BCC)          也叫重连通图
// v-BCC：任意割点都是至少两个不同 v-BCC 的公共点              广义圆方树
// 每个 v-BCC 的点数就是一个极大环，所有即
// https://oi-wiki.org/graph/bcc/
// https://www.csie.ntu.edu.tw/~hsinmu/courses/_media/dsa_13spring/horowitz_306_311_biconnected.pdf
// 好题 https://codeforces.com/problemset/problem/962/F
// https://leetcode-cn.com/problems/s5kipK/
// 结合树链剖分 https://codeforces.com/problemset/problem/487/E
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
func (G *graph) findVertexBCC(g [][]int, min func(int, int) int) (comps [][]int, bccIDs []int) {
	bccIDs = make([]int, len(g)) // ID 从 1 开始编号
	idCnt := 0
	isCut := make([]bool, len(g))

	dfn := make([]int, len(g)) // 值从 1 开始
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
	cutIDs := make([]int, len(g))
	for i, is := range isCut {
		if is {
			idCnt++ // 接在 BCC 之后给割点编号
			cutIDs[i] = idCnt
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
// 缩点后形成一棵 bridge tree
// 模板题 https://codeforces.com/problemset/problem/1000/E
// 较为综合的一道题 http://codeforces.com/problemset/problem/732/F
func (G *graph) findEdgeBCC(in io.Reader, n, m int) (comps [][]int, bccIDs []int) {
	type neighbor struct{ to, eid int }
	type edge struct{ v, w int }
	g := make([][]neighbor, n)
	edges := make([]edge, m)

	// *copy* 包含读图
	isBridge := G.findBridges(in, n, m)

	// 求原图中每个点的 bccID
	bccIDs = make([]int, len(g))
	idCnt := 0
	var comp []int
	var f2 func(int)
	f2 = func(v int) {
		bccIDs[v] = idCnt
		comp = append(comp, v)
		for _, e := range g[v] {
			if w := e.to; !isBridge[e.eid] && bccIDs[w] == 0 {
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
	for eid, b := range isBridge {
		if b {
			e := edges[eid]
			v, w := bccIDs[e.v]-1, bccIDs[e.w]-1
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
		}
	}

	return
}

type dijkstraPair struct {
	v   int
	dis int64
}
type dijkstraHeap []dijkstraPair

func (h dijkstraHeap) Len() int              { return len(h) }
func (h dijkstraHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h dijkstraHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *dijkstraHeap) Push(v interface{})   { *h = append(*h, v.(dijkstraPair)) }
func (h *dijkstraHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *dijkstraHeap) push(v dijkstraPair)  { heap.Push(h, v) }
func (h *dijkstraHeap) pop() dijkstraPair    { return heap.Pop(h).(dijkstraPair) }

// 单源最短路 Dijkstra
// 适用于稀疏图 O(mlogm)
// 根据《算法(第4版)》，这里实现的是 lazy 版本的 Dijkstra，复杂度为 O(mlogm)；若在插入堆时元素已在堆中，改成更新元素而不是插入元素可使复杂度降为 O(mlogn)
// st 也可以是一个点集，这相当于同时对多个点跑最短路
// 视频讲解（第四题）https://www.bilibili.com/video/BV1wj411G7sH/
// 可视化 https://visualgo.net/zh/sssp
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 最短路问题笔记 https://www.luogu.com.cn/blog/SCN/zui-duan-lu-wen-ti-bi-ji
//
// 模板题 https://www.luogu.com.cn/problem/P3371 https://www.luogu.com.cn/problem/P4779
//       https://codeforces.com/problemset/problem/20/C
//       LC743 https://leetcode-cn.com/problems/network-delay-time/
// 最短路个数 https://www.luogu.com.cn/problem/P1608
// 通过最短路找到可以删除的边 https://codeforces.com/problemset/problem/449/B
// 稠密图 https://atcoder.jp/contests/arc064/tasks/arc064_c
// 【理解本质】https://atcoder.jp/contests/abc271/tasks/abc271_e
// 建模 https://www.luogu.com.cn/problem/P4644
// 建模 LC864 https://leetcode-cn.com/problems/shortest-path-to-get-all-keys/
// 建模【好题】https://codeforces.com/contest/1528/problem/D
// 建模+转换+多源最短路 https://codeforces.com/problemset/problem/1753/D
// 还能再走多远？LC882 https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph/
// 转换 LC2577 https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/
// 转换 https://atcoder.jp/contests/abc237/tasks/abc237_e
// 双关键字+记录路径编号 https://codeforces.com/problemset/problem/507/E
// 关键边、伪关键边（与割边结合）https://codeforces.com/problemset/problem/567/E
// 基于 max LC1631 https://leetcode-cn.com/problems/path-with-minimum-effort/
// 题目推荐 https://cp-algorithms.com/graph/dijkstra.html#toc-tgt-5
// 线段树建图优化 https://codeforces.com/problemset/problem/786/B
// 涉及到相邻两条边的最短路 https://codeforces.com/contest/1486/problem/E
// todo 与扩欧结合 https://www.acwing.com/problem/content/3418/
// 跑两遍最短路，第二次修正边权来改变最短路 https://codeforces.com/problemset/problem/715/B
// 分层图最短路
//    空间压缩 https://codeforces.com/problemset/problem/1442/C
//    转换 https://codeforces.com/problemset/problem/1473/E
// todo 动态最短路 https://codeforces.com/problemset/problem/1163/F
//
// 最短路径树
// todo https://xyzl.blog.luogu.org/Shortest-Path-Tree-SPT
// 最短路树上跑拓扑排序 https://codeforces.com/contest/1076/problem/D
// todo https://codeforces.com/problemset/problem/1005/F
// todo MST https://codeforces.com/problemset/problem/545/E
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
	h := dijkstraHeap{{st, 0}}
	for len(h) > 0 {
		p := h.pop()
		v := p.v
		if vis[v] { // p.dis > dist[v]
			continue
		}
		vis[v] = true
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				fa[w] = v
				h.push(dijkstraPair{w, dist[w]})
			}
		}
	}

	//if dist[end] == inf {
	//	return -1
	//}

	// EXTRA: path from end to start
	// 记录边的编号 https://codeforces.com/problemset/problem/507/E
	path := []int{}
	end := n - 1 //
	for x := end; x != -1; x = fa[x] {
		path = append(path, x)
	}

	// EXTRA: 构建有向最短路树（除根节点外的点的入度均为 1）
	// https://www.acwing.com/problem/content/3631/
	{
		g2 := make([][]neighbor, n)
		vis := make([]bool, n)
		for v, es := range g {
			for _, e := range es {
				w := e.to
				if vis[w] { // 已经有一条指向 w 的边了。若去掉该判断则构建的是 DAG
					continue
				}
				if dist[v]+e.wt == dist[w] {
					g2[v] = append(g2[v], e)
					vis[w] = true
				}
			}
		}
	}

	// EXTRA: 在最短路 DAG 上跑拓扑（如最短路计数）
	// https://www.luogu.com.cn/problem/P1144 https://www.luogu.com.cn/problem/P1608 LC1976 https://leetcode-cn.com/problems/number-of-ways-to-arrive-at-destination/
	// 也可以把转移写在求最短路的代码中，见 https://www.luogu.com.cn/record/56683589
	// 紧急情况 https://www.acwing.com/problem/content/1477/
	// 条条大路通罗马 https://www.acwing.com/problem/content/1579/
	{
		deg := make([]int, n)
		for v, es := range g {
			for _, e := range es {
				if dist[v]+e.wt == dist[e.to] {
					deg[e.to]++
				}
			}
		}

		dp := make([]int, n)
		dp[st] = 0 //
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; dist[v]+e.wt == dist[w] {
					// do dp[v] -> dp[w] ...

					if deg[w]--; deg[w] == 0 {
						q = append(q, w)
					}
				}
			}
		}
	}

	// EXTRA: 次短路
	// 模板题 https://www.luogu.com.cn/problem/P2865
	// LC2045 https://leetcode-cn.com/problems/second-minimum-time-to-reach-destination/
	// 次短路计数 https://www.acwing.com/problem/content/385/ https://codeforces.com/contest/1650/problem/G
	// 长度不超过最短路长度+K 的路径个数 [NOIP2017 提高组] 逛公园 https://www.luogu.com.cn/problem/P3953
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
		h := dijkstraHeap{{st, 0}}
		for len(h) > 0 {
			p := h.pop()
			v, d := p.v, p.dis
			if dist2[v] < d { // 注意是 dist2
				continue
			}
			for _, e := range g[v] {
				w := e.to
				newD := d + e.wt
				if newD < dist[w] {
					h.push(dijkstraPair{w, newD})
					dist[w], newD = newD, dist[w]
				}
				if dist[w] < newD && newD < dist2[w] {
					h.push(dijkstraPair{w, newD})
					dist2[w] = newD
				}
			}
		}
		_ = dist2[end]
	}

	return
}

// 另一种 Dijkstra 写法
// 适用于稠密图 O(n^2)
// 建模 https://codeforces.com/contest/1528/problem/D
func (*graph) shortestPathDijkstra2(g [][]int64, st int) []int64 {
	n := len(g)

	const inf int64 = 1e18 // 1e9+1
	dis := make([]int64, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	vis := make([]bool, n)
	for {
		v := -1
		for w, b := range vis {
			if !b && (v < 0 || dis[w] < dis[v]) {
				v = w
			}
		}
		if v < 0 {
			return dis
		}
		vis[v] = true
		for w, wt := range g[v] {
			if newD := dis[v] + wt; newD < dis[w] {
				dis[w] = newD
			}
		}
	}
}

// 0-1 最短路 / 0-1 BFS
// https://oi-wiki.org/graph/bfs/#bfs_3
// https://codeforces.com/blog/entry/22276
// EXTRA: 1-2 最短路 https://codeforces.com/blog/entry/90917
// 例题: https://codeforces.com/problemset/problem/173/B
// 网格图 https://codeforces.com/problemset/problem/590/C
// 建图技巧 https://codeforces.com/problemset/problem/821/D
// 建图技巧 https://codeforces.com/problemset/problem/1340/C
// 哪里有 1 https://atcoder.jp/contests/abc213/tasks/abc213_e
//         https://atcoder.jp/contests/abc176/tasks/abc176_d
// https://codeforces.com/problemset/problem/877/D（也可以 BFS）
// https://codeforces.com/problemset/problem/1063/B
// https://codeforces.com/problemset/problem/1442/C
func (*graph) bfs01(g [][]struct{ to, wt int }, st int) []int {
	const inf int = 1e9
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	ql, qr := []int{st}, []int{}
	for len(ql) > 0 || len(qr) > 0 {
		var v int
		if len(ql) > 0 {
			ql, v = ql[:len(ql)-1], ql[len(ql)-1]
		} else {
			v, qr = qr[0], qr[1:]
		}
		for _, e := range g[v] {
			w, d := e.to, e.wt
			if newD := dis[v] + d; newD < dis[w] {
				dis[w] = newD
				if d == 0 {
					ql = append(ql, w)
				} else {
					qr = append(qr, w)
				}
			}
		}
	}
	return dis
}

// 单源最短路 SPFA O(nm)   队列优化的 Bellman-Ford
// 对于构建一个让 SPFA 跑到最坏情况的（网格）图，见 main/testutil/rand.go 中的 GraphHackSPFA
// 可视化 https://visualgo.net/zh/sssp
// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
//
// 模板题 https://www.luogu.com.cn/problem/P3385
// 用到 Bellman-Ford 的思想 LC1928 https://leetcode.cn/problems/minimum-cost-to-reach-destination-in-time/
// 边权递增最短路 http://acm.hdu.edu.cn/showproblem.php?pid=4479
//
// EXTRA: 差分约束系统
// 若有 Xi-Xj<=Ck，则连一条有向边 j->i，边权为 Ck
// 然后再添加一个 0 号节点，向其他节点连一条边权为 0 的有向边，表示 Xi-X0<=0
// 这样，在无负环时会得到一组非正数解
// 模板题 https://www.luogu.com.cn/problem/P4878
// 每个区间至少选 ci 个 https://www.luogu.com.cn/problem/P1250 https://www.luogu.com.cn/problem/SP116 http://poj.org/problem?id=1201
// - todo 加强版 https://leetcode-cn.com/problems/t3fKg1/
// - todo 打印方案 https://atcoder.jp/contests/abc216/tasks/abc216_g
func (*graph) shortestPathSPFA(in io.Reader, n, m, st int) (dist []int64) { // 有负环时返回 nil
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
				// 找到一个从 st 出发可达的负环
				if relaxedCnt[w] >= n {
					return nil
				}
				if !inQ[w] {
					inQ[w] = true
					q = append(q, w)
				}
			}
		}
	}

	// EXTRA: 只是找负环的话，初始时将所有点入队即可
	// 注意不能只从一个点出发找负环，因为可能这个点无法到达负环
	// https://www.luogu.com.cn/problem/P2868

	// EXTRA: 打印负环
	// https://cp-algorithms.com/graph/finding-negative-cycle-in-graph.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BellmanFordSP.java.html

	return
}

// 任意两点最短路 Floyd-Warshall  O(n^3)  本质是求 Min-plus matrix multiplication
// 传入邻接矩阵 dis
// dis[v][w] == inf 表示没有 v-w 边
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// https://en.wikipedia.org/wiki/Min-plus_matrix_multiplication
// https://oi-wiki.org/graph/shortest-path/#floyd
// https://zhuanlan.zhihu.com/p/623757829
// 题目推荐 https://cp-algorithms.com/graph/all-pair-shortest-path-floyd-warshall.html#toc-tgt-5
// https://codeforces.com/problemset/problem/1204/C
// LC1334 https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
// 动态加点 https://codeforces.com/problemset/problem/295/B
// 动态加边 https://codeforces.com/problemset/problem/25/C LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
// todo https://atcoder.jp/contests/abc243/tasks/abc243_e
// 传递闭包 UVa247 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=4&page=show_problem&problem=183
// 注：求传递闭包时，若 i-k 不连通，则最内层循环无需运行
// 任意两点最大边权最小路径 UVa10048 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=12&page=show_problem&problem=989
func (*graph) shortestPathFloydWarshall(dis [][]int, min func(int, int) int) [][]int {
	// dis[k][i][j] 表示「经过若干个编号不超过 k 的中间节点」时，从 i 到 j 的最短路长度，其中第一维可以压缩掉
	// 为什么可以把第一维度去掉？dis[i][k] 和 dis[k][j] 不会被覆盖掉吗？
	// 见算法导论第三版练习 25.2-4（网络上有习题解答）
	for k := range dis { // 中间节点的最大编号
		for i := range dis {
			for j := range dis {
				// 决策（k 不是中间节点，k 是中间节点）
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	// 如果出现 dis[i][i] < 0 则说明有负环

	// 动态加边
	// https://codeforces.com/problemset/problem/25/C
	// LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
	for i := range dis {
		// 注意 from=i 或者 to=j 时，下面的 dis[i][from] 和 dis[to][j] 都需要 dis[i][i] 这样的值
		// 所以初始化成 0 方便计算
		dis[i][i] = 0
	}
	addEdge := func(from, to int, wt int) { // wt int64
		// 无法让任何最短路变短
		if wt >= dis[from][to] {
			return
		}
		for i := range dis {
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][from]+wt+dis[to][j])
			}
		}
	}
	_ = addEdge

	return dis
}

// 位压缩版 O(n^3/w)
// LC2101 https://leetcode-cn.com/problems/detonate-the-maximum-bombs/
// https://atcoder.jp/contests/abc287/tasks/abc287_h
func (*graph) floydWarshallBitset(in io.Reader, n, m int) []int {
	vs := make([]Bitset, n) // vs[i] 表示从 i 出发可以到达的节点
	for i := range vs {
		vs[i] = NewBitset(n)
		vs[i].Set(i) // i 可以到达 i
	}
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &m)
		v--
		w--
		vs[v].Set(w) // 有向边 v->w
	}
	for k := range vs {
		for i := range vs {
			if vs[i].Has(k) {
				vs[i].UnionFrom(vs[k]) // i->j 现在可以 i->k->j
			}
		}
	}
	reach := make([]int, n) // reach[i] 表示从 i 出发可以到达的节点数（注意读题，一般都要包括自己）
	for i, bs := range vs {
		reach[i] = bs.OnesCount()
	}
	return reach
}

// 最小环
// 传入邻接矩阵 weights
// weights[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/min-circle/#floyd
// NOTE: 无权图的情况见 shortestCycleBFS
func (*graph) shortestCycleFloydWarshall(weights [][]int64, min func(int64, int64) int64) int64 {
	const inf int64 = 1e18
	//const inf int = 1e8 // *NOTE*
	n := len(weights)
	dist := make([][]int64, n)
	for i := range dist {
		dist[i] = append([]int64(nil), weights[i]...)
	}
	ans := inf
	for k := range dist {
		for i := 0; i < k; i++ {
			for j := 0; j < i; j++ {
				ans = min(ans, dist[i][j]+weights[i][k]+weights[k][j])
			}
		}
		for i := range dist {
			for j := range dist {
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
// 返回使 ∑a[i]*x[i] = b 有非负数解的 b 的个数，其中 0<=b<=limit
// a[i] >= 0
// https://oi-wiki.org/graph/mod-shortest-path/
// 跳楼机 https://www.luogu.com.cn/problem/P3403
// https://www.luogu.com.cn/problem/P2371
// https://codeforces.com/problemset/problem/986/F
func (*graph) shortestPathMod(a []int, limit int64) (ans int64) {
	sort.Ints(a) // 常数优化
	dis := make([]int64, a[0])
	for i := range dis {
		dis[i] = math.MaxInt64
	}
	dis[0] = 0
	h := dijkstraHeap{{}}
	for len(h) > 0 {
		top := h.pop()
		v := top.v
		if top.dis > dis[v] {
			continue
		}
		for _, ai := range a[1:] {
			w := (v + ai) % a[0]
			if newD := dis[v] + int64(ai); newD < dis[w] {
				dis[w] = newD
				h.push(dijkstraPair{w, newD})
			}
		}
	}
	for _, d := range dis {
		if d <= limit {
			ans += (limit-d)/int64(a[0]) + 1
		}
	}
	return
}

// k 短路
// A* 算法
// 可持久化可并堆优化
// https://en.wikipedia.org/wiki/K_shortest_path_routing
// todo https://oi-wiki.org/graph/kth-path/
// todo 模板题 https://www.luogu.com.cn/problem/P2483

// 最小斯坦纳树
// https://oi-wiki.org/graph/steiner-tree/
// todo 模板题 https://www.luogu.com.cn/problem/P6192
//  WC08 游览计划 https://www.luogu.com.cn/problem/P4294

// 最小生成树 Kruskal
// 适用于稀疏图 O(mlogm)，或者边已经按权值排序的情况
// 性质：
// - 对于不同的 MST，同一边权的边的个数都是相同的（应用见后面的最小生成树计数）
// - 对于任意正确加边方案，加完小于某权值的边后，图的连通性是一样的
// https://oi-wiki.org/graph/mst/#kruskal
// 边权 [0,1] 的随机完全图的 MST 权值和是 ζ(3) = 1.202…	https://en.wikipedia.org/wiki/Random_minimum_spanning_tree https://www.sciencedirect.com/science/article/pii/0166218X85900587
//
// 模板题 https://www.luogu.com.cn/problem/P3366 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/E
// 题目推荐 https://cp-algorithms.com/graph/mst_kruskal.html#toc-tgt-5
// 需要一些数论知识 https://atcoder.jp/contests/abc210/tasks/abc210_e
// 枚举 https://atcoder.jp/contests/abc270/tasks/abc270_f
// 关键边、伪关键边（与割边结合）https://codeforces.com/problemset/problem/160/D LC1489 https://leetcode.cn/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
// 判断给定的边是否均在同一棵 MST 中 https://codeforces.com/problemset/problem/891/C
// 二分图无环 https://codeforces.com/problemset/problem/1408/E
// 与 LCA 结合 https://codeforces.com/problemset/problem/733/F
// 最小生成树的最长边：Kruskal 中最后一条加入 MST 中的边的长度 https://www.luogu.com.cn/problem/P1547
// EXTRA: 与树链剖分结合可以在线查询两点间路径最大边权的最小值 https://leetcode-cn.com/problems/checking-existence-of-edge-length-limited-paths/
// 边权为 a[i]+a[j] 的混合 MST https://codeforces.com/problemset/problem/1095/F
// - 完全图找个最小的 a[i] 和其余点连边
// todo 只有两种边权的图的 MST 的性质 + 所有 MST 中的单源最短路的最小值 http://codeforces.com/problemset/problem/1149/D
// 与 DFS 搜索树结合 https://codeforces.com/problemset/problem/1707/C
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
			fa[fv] = fw
			sum += int64(e.wt)
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
// 适用于稠密图 O(n^2)，传入邻接矩阵 dis
// dis[v][w] == inf 表示没有 v-w 边
// 有些题目需要在连通分量上求 MST，这时就需要用到 root
// 可视化 https://visualgo.net/zh/mst
// https://oi-wiki.org/graph/mst/#prim
// 模板题 https://www.luogu.com.cn/problem/P1546
// 建模+打印方案 https://codeforces.com/problemset/problem/1245/D
// https://codeforces.com/contest/1508/problem/C
func (*graph) mstPrim(dis [][]int, root int) (mst int, edges [][2]int) {
	edges = make([][2]int, 0, len(dis)-1)

	// 注意：dis 需要保证 dis[i][i] = inf，从而避免自环的影响

	const inf int = 2e9
	minD := make([]struct{ v, d int }, len(dis)) // minD[i].d 表示当前 MST 到点 i 的最小距离，对应的边为 minD[i].v-i
	for i := range minD {
		minD[i].d = inf
	}
	minD[root].d = 0
	inMST := make([]bool, len(dis)) // 初始时所有点都不在 MST 中
	for {
		// 根据切分定理，求不在当前 MST 的点到当前 MST 的最小距离，即 minD[v].d
		v := -1
		for w, in := range inMST {
			if !in && (v < 0 || minD[w].d < minD[v].d) {
				v = w
			}
		}
		if v < 0 { // 已求出 MST
			return
		}

		// 加入 MST
		inMST[v] = true
		mst += minD[v].d // int64
		if v != root {
			edges = append(edges, [2]int{minD[v].v, v})
		}

		// 更新 minD
		for w, d := range dis[v] {
			if !inMST[w] && d < minD[w].d { // 注：若 mstPrim 结束后 minD 无其他用途，!inMST[w] 的判断可以去掉
				minD[w].d = d
				minD[w].v = v
			}
		}
	}
}

// Boruvka's algorithm
// 用于求解边权互不相同的无向图的最小生成森林
// https://en.wikipedia.org/wiki/Bor%C5%AFvka%27s_algorithm
// https://oi-wiki.org/graph/mst/#boruvka
// https://www.geeksforgeeks.org/boruvkas-algorithm-greedy-algo-9/
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BoruvkaMST.java.html
// todo http://codeforces.com/problemset/problem/888/G
//  https://codeforces.com/problemset/problem/1550/F https://www.luogu.com.cn/blog/ETHANK/boruvka-xiao-ji
func (*graph) boruvkaMST(n, m int) (sum int64) {
	return
}

// 单点度数（单度）限制最小生成树   O(n^2)
// 点 root 的度数不超过 lim
// 不超过 http://poj.org/problem?id=1639 https://codeforces.com/gym/100227 A https://www.acwing.com/problem/content/349/
// EXTRA: 恰好的情况（需要用 WQS 二分）https://codeforces.com/problemset/problem/125/E
func (*graph) limitDegreeMST(dis [][]int, root, lim int) int {
	const inf int = 2e9

	n := len(dis)
	mstSum := 0 // int64
	rootDeg := 0
	mst := make([][]int, n)
	for i := range mst {
		mst[i] = make([]int, n)
		for j := range mst[i] {
			mst[i][j] = inf
		}
	}

	inComp := make([]bool, n)
	inComp[root] = true // 下面求出去掉 root 后的所有连通块
	minD := make([]struct{ v, d int }, n)
	for i := range minD {
		minD[i].d = inf
	}
	inMST := make([]bool, n)
	for st, inC := range inComp {
		if !inC {
			comp := []int{}
			var dfs func(int)
			dfs = func(v int) {
				comp = append(comp, v)
				inComp[v] = true
				for w, d := range dis[v] {
					if d < inf && !inComp[w] {
						dfs(w)
					}
				}
			}
			dfs(st)

			// 求该连通块的 MST
			minD[st].d = 0
			for {
				v := -1
				for _, w := range comp {
					if !inMST[w] && (v < 0 || minD[w].d < minD[v].d) {
						v = w
					}
				}
				if v < 0 {
					break
				}
				mstSum += minD[v].d
				inMST[v] = true
				for _, w := range comp {
					if !inMST[w] && dis[v][w] < minD[w].d {
						minD[w].d = dis[v][w]
						minD[w].v = v
					}
				}
			}

			// 连通块内部的最小生成树，连边
			closest := st
			for _, w := range comp {
				if w == st {
					continue
				}
				mst[minD[w].v][w] = minD[w].d
				mst[w][minD[w].v] = minD[w].d
				if dis[root][w] < dis[root][closest] {
					closest = w // 顺带求出该连通块到 root 最近的点
				}
			}

			// 每个连通块通过 closest 跟 root 连边
			d := dis[root][closest]
			mstSum += d
			mst[root][closest] = d
			mst[closest][root] = d
			rootDeg++
		}
	}

	type maxEdge struct{ v, w, wt int }
	dp := make([]maxEdge, n) // dp[i] 表示从 root 到 i 这条路径上的边权最大的边
	var f func(int, int)
	f = func(v, fa int) {
		for w, wt := range mst[v] {
			if wt < inf && w != fa {
				if wt < dp[v].wt {
					dp[w] = dp[v]
				} else {
					dp[w] = maxEdge{v, w, wt}
				}
				f(w, v)
			}
		}
	}
	f(root, -1)

	for rootDeg < lim {
		maxDec, maxV := 0, 0
		for v, d := range dis[root] {
			// 枚举从 root 出发的非树边，看看能否做到减小 mstSum
			if d < inf && mst[root][v] == inf {
				if dec := mst[dp[v].v][dp[v].w] - d; dec > maxDec {
					maxDec, maxV = dec, v
				}
			}
		}
		if maxDec == 0 { // 无法再减小 mstSum
			break
		}

		// 减小 mstSum：删去树边 dp[maxV].v-dp[maxV].w，加入非树边 root-maxV
		mstSum -= maxDec
		mst[dp[maxV].v][dp[maxV].w] = inf
		mst[dp[maxV].w][dp[maxV].v] = inf
		mst[root][maxV] = dis[root][maxV]
		mst[maxV][root] = dis[root][maxV]
		rootDeg++

		// 重新计算以 maxV 为根的一条路径的 dp 状态
		dp[maxV] = maxEdge{root, maxV, dis[root][maxV]}
		f(maxV, root) // 注意这里 fa 填的 root，所以只会更新受到上面删边和加边影响的子树
	}
	return mstSum
}

// 严格次小生成树 Second best Minimum Spanning Tree
// https://oi-wiki.org/graph/mst/#_13
// https://cp-algorithms.com/graph/second_best_mst.html
// 模板题（严格）https://www.luogu.com.cn/problem/P4180 https://www.acwing.com/problem/content/358/
// 注：非严格次小生成树
//     做法更加简单，维护路径最大值即可，见 https://oi-wiki.org/graph/mst/#_10
func (*graph) strictlySecondMST(n int, edges []struct{ v, w, wt int }, min, max func(int, int) int) int {
	sort.Slice(edges, func(i, j int) bool { return edges[i].wt < edges[j].wt })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	mstSum := 0 // int64
	inMST := make([]bool, len(edges))
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i, e := range edges {
		v, w, wt := e.v, e.w, e.wt
		if fv, fw := find(v), find(w); fv != fw {
			fa[fv] = fw
			mstSum += wt
			inMST[i] = true
			g[v] = append(g[v], nb{w, wt}) // MST
			g[w] = append(g[w], nb{v, wt})
		}
	}

	const mx = 17
	type pair struct{ p, fi, se int }
	pa := make([][mx]pair, n)
	dep := make([]int, n)
	var build func(v, p, d int)
	build = func(v, p, d int) {
		pa[v][0].p = p
		dep[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].fi = e.wt
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	merge := func(xFi, xSe, yFi, ySe int) (int, int) {
		fi, se := max(xFi, yFi), 0
		if xFi == yFi {
			se = max(xSe, ySe)
		} else if xFi > yFi {
			se = max(xSe, yFi)
		} else {
			se = max(xFi, ySe)
		}
		return fi, se
	}
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				fi, se := merge(p.fi, p.se, pp.fi, pp.se)
				pa[v][i+1] = pair{pp.p, fi, se}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}

	// 返回路径最大边权和严格次大边权
	queryPath := func(v, w int) (fi, se int) {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (dep[w]-dep[v])>>i&1 > 0 {
				p := pa[w][i]
				fi, se = merge(fi, se, p.fi, p.se)
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					fi, se = merge(fi, se, pv.fi, pv.se)
					fi, se = merge(fi, se, pw.fi, pw.se)
					v, w = pv.p, pw.p
				}
			}
			fi, se = merge(fi, se, pa[v][0].fi, pa[v][0].se)
			fi, se = merge(fi, se, pa[w][0].fi, pa[w][0].se)
		}
		return
	}

	const inf int = 1e9 // 1e18
	delta := inf
	for i, e := range edges {
		v, w, wt := e.v, e.w, e.wt
		if inMST[i] || v == w { // 注意跳过自环
			continue
		}
		fi, se := queryPath(v, w)
		if wt > fi {
			delta = min(delta, wt-fi) // 替换从而得到更大的 MST，取最小的替换差值
		} else if se > 0 { // 此时必然有 wt == fi
			delta = min(delta, wt-se) // wt = fi > se，同样可以替换
		}
	}
	if delta == inf {
		return -1
	}
	mstSum += delta

	return mstSum
}

// 曼哈顿距离最小生成树 O(nlogn)
// LC1584 https://leetcode-cn.com/problems/min-cost-to-connect-all-points/
// 做法见官方题解 https://leetcode-cn.com/problems/min-cost-to-connect-all-points/solution/lian-jie-suo-you-dian-de-zui-xiao-fei-yo-kcx7/
func (*graph) manhattanMST(points []struct{ x, y, i int }, abs func(int) int) (mst int) {
	n := len(points)
	// 读入时把 points 加上下标

	type edge struct{ v, w, dis int }
	edges := []edge{}

	build := func() {
		sort.Slice(points, func(i, j int) bool { a, b := points[i], points[j]; return a.x < b.x || a.x == b.x && a.y < b.y })

		// 离散化 y-x
		type pair struct{ v, i int }
		ps := make([]pair, n)
		for i, p := range points {
			ps[i] = pair{p.y - p.x, i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })
		kth := make([]int, n)
		k := 1
		kth[ps[0].i] = k
		for i := 1; i < n; i++ {
			if ps[i].v != ps[i-1].v {
				k++
			}
			kth[ps[i].i] = k
		}

		const inf int = 2e9
		tree := make([]int, k+1)
		idRec := make([]int, k+1)
		for i := range tree {
			tree[i] = inf
			idRec[i] = -1
		}
		update := func(pos, val, id int) {
			for ; pos > 0; pos &= pos - 1 {
				if val < tree[pos] {
					tree[pos] = id
					idRec[pos] = id
				}
			}
		}
		query := func(pos int) int {
			minVal, minID := inf, -1
			for ; pos < len(tree); pos += pos & -pos {
				if tree[pos] < minVal {
					minVal = tree[pos]
					minID = idRec[pos]
				}
			}
			return minID
		}

		for i := n - 1; i >= 0; i-- {
			p := points[i]
			pos := kth[i]
			if j := query(pos); j != -1 {
				q := points[j]
				dis := abs(p.x-q.x) + abs(p.y-q.y)
				edges = append(edges, edge{p.i, q.i, dis})
			}
			update(pos, p.x+p.y, i)
		}
	}
	build()
	for i := range points {
		points[i].x, points[i].y = points[i].y, points[i].x
	}
	build()
	for i := range points {
		points[i].x = -points[i].x
	}
	build()
	for i := range points {
		points[i].x, points[i].y = points[i].y, points[i].x
	}
	build()

	sort.Slice(edges, func(i, j int) bool { return edges[i].dis < edges[j].dis })

	uf := NewUnionFind(n)
	left := n - 1
	for _, e := range edges {
		if uf.Merge(e.v, e.w) >= 0 {
			mst += e.dis // int64
			left--
			if left == 0 {
				break
			}
		}
	}
	return
}

// Kruskal 重构树
// https://oi-wiki.org/graph/mst/#kruskal_1

// 最小生成树计数 Kirchhoff's theorem
// https://www.luogu.com.cn/problem/P4208
// todo

// 最小差值生成树
// edges 中的节点编号从 0 开始，且无自环
// https://www.luogu.com.cn/problem/P4234
// https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/F
// Japan07，紫书例题 11-2，UVa1395 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=446&page=show_problem&problem=4141
func (*graph) minDiffMST(n int, edges [][3]int) int {
	m := len(edges)
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
	nodes := make([]*lctNode, n+m)
	fa := make([]int, n+m)
	for i := range nodes {
		nodes[i] = &lctNode{id: i}
		nodes[i].mi = nodes[i]
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		fa[x] = y
		return true
	}
	cc := 0

	const inf int = 2e9 + 1
	ans := inf
	del := make([]bool, m)
	l := 0
	for i, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		x, y, mid := nodes[v], nodes[w], nodes[n+i]
		if merge(v, w) {
			mid.link(x)
			mid.link(y)
			cc++
		} else {
			x.split(y)
			mi := y.mi
			del[mi.id-n] = true
			mi.cut(nodes[edges[mi.id-n][0]])
			mi.cut(nodes[edges[mi.id-n][1]])
			mid.link(x)
			mid.link(y)
		}
		for del[l] {
			l++
		}
		if cc == n-1 && wt-edges[l][2] < ans {
			ans = wt - edges[l][2]
		}
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

// 最小树形图 (MSA, Minimum weight Spanning Arborescence)   有向图上的最小生成树 (DMST)
// O(nm) 朱刘算法（Edmonds 算法）
// https://en.wikipedia.org/wiki/Edmonds%27_algorithm
// https://oi-wiki.org/graph/dmst/
// todo 另外还有 Tarjan 的 O(m+nlogn) 算法
//  https://oi-wiki.org/graph/dmst/#tarjan-dmst
//
// 模板题 https://www.luogu.com.cn/problem/P4716
// todo https://codeforces.com/problemset/problem/240/E
func (*graph) msaEdmonds(n, root int, edges [][3]int) (ans int64) {
	const inf int = 2e9
	minW := make([]int, n)
	fa := make([]int, n)
	id := make([]int, n)
	rt := make([]int, n)
	for {
		for i := range minW {
			minW[i] = inf
		}
		for _, e := range edges {
			if v, w, wt := e[0], e[1], e[2]; wt < minW[w] {
				minW[w] = wt
				fa[w] = v
			}
		}
		for i, wt := range minW {
			if i != root && wt == inf {
				return -1
			}
		}
		cid := 0
		for i := range id {
			id[i] = -1
			rt[i] = -1
		}
		for i, wt := range minW {
			if i == root {
				continue
			}
			ans += int64(wt)
			v := i
			for ; v != root && id[v] < 0 && rt[v] != i; v = fa[v] {
				rt[v] = i
			}
			if v != root && id[v] < 0 { // rt[v] == i，有环
				id[v] = cid
				for x := fa[v]; x != v; x = fa[x] {
					id[x] = cid
				}
				cid++
			}
		}
		if cid == 0 {
			return
		}
		for i, v := range id {
			if v < 0 {
				id[i] = cid
				cid++
			}
		}
		// 缩点
		tmp := edges
		edges = nil
		for _, e := range tmp {
			if v, w := id[e[0]], id[e[1]]; v != w {
				edges = append(edges, [3]int{v, w, e[2] - minW[e[1]]})
			}
		}
		root = id[root]
		minW = minW[:cid]
		id = id[:cid]
	}
}

// 反图的连通分量 O(n+m)
// 这个算法是我独自想出来的 :)
// https://www.luogu.com.cn/blog/endlesscheng/solution-cf1242b
// https://codeforces.com/blog/entry/93652
//
// https://codeforces.com/contest/190/problem/E
// https://codeforces.com/contest/920/problem/E
// https://codeforces.com/contest/1242/problem/B
// MST https://codeforces.com/contest/1508/problem/C
func (*graph) inverseGraphComponents(n int, g [][]int) [][]int {
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

	// 原图度数最小的点即为反图度数最大的点
	maxV := 0
	for v, ws := range g {
		if len(ws) < len(g[maxV]) {
			maxV = v
		}
	}
	if len(g[maxV]) == n-1 { // 特殊情况：反图无边
		components := make([][]int, n)
		for i := range components {
			components[i] = []int{i} // i+1
		}
		return components
	}

	initFa(n)
	mergeInv := func(v int) {
		has := map[int]bool{v: true}
		for _, w := range g[v] {
			has[w] = true
		}
		for i := range g {
			if !has[i] {
				merge(i, v)
			}
		}
	}
	mergeInv(maxV)
	// 暴力合并其余不在 maxV 所处连通分量的点
	for v := range g {
		if !same(v, maxV) {
			mergeInv(v)
		}
	}

	componentsMap := map[int][]int{} // make([][]int, n)
	for i := range fa {
		rootV := find(i)
		componentsMap[rootV] = append(componentsMap[rootV], i) // i+1
	}
	components := make([][]int, 0, len(componentsMap))
	for _, vs := range componentsMap {
		components = append(components, vs)
	}
	return components
}

// 二分图判定+染色
// 注：二分图也叫偶图
// https://en.wikipedia.org/wiki/Bipartite_graph
// https://oi-wiki.org/graph/bi-graph/#_3
// https://cp-algorithms.com/graph/bipartite-check.html
//
// 模板题 LC886 https://leetcode.cn/problems/possible-bipartition/
// https://codeforces.com/problemset/problem/1093/D
// https://www.luogu.com.cn/problem/P6185
// https://codeforces.com/problemset/problem/1537/F
// 转换 https://codeforces.com/problemset/problem/85/E
// 染色的技巧 https://codeforces.com/problemset/problem/553/C
//          https://codeforces.com/problemset/problem/662/B
// 树至多加多少条边仍然是二分图 https://codeforces.com/problemset/problem/862/B
// 与背包结合（NEERC01，紫书例题 9-19，UVa 1627）https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4502
// 与分组背包结合 https://codeforces.com/problemset/problem/1354/E
func (*graph) isBipartite(g [][]int) bool {
	colors := make([]int8, len(g)) // 0 表示未访问该节点
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			// 如果要分组，用 3^c，便于填入下标；如果要根据染色来 +/-，用 -c
			if colors[w] == c || colors[w] == 0 && !f(w, 3^c) {
				return false
			}
		}
		return true
	}
	//f(0, 1) // 只有一个 CC
	for i, c := range colors {
		if c != 0 {
			continue
		}
		if !f(i, 1) {
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

/* 匹配 完美匹配 带权匹配 独立集 边覆盖 顶点覆盖 路径覆盖 支配集
https://en.wikipedia.org/wiki/Matching_(graph_theory)
https://en.wikipedia.org/wiki/Perfect_matching 完美匹配 iff 唯一匹配
https://en.wikipedia.org/wiki/Maximum_weight_matching
https://en.wikipedia.org/wiki/Independent_set_(graph_theory)
https://en.wikipedia.org/wiki/Maximal_independent_set
https://en.wikipedia.org/wiki/Edge_cover
https://en.wikipedia.org/wiki/Vertex_cover
https://en.wikipedia.org/wiki/Path_cover
https://en.wikipedia.org/wiki/Dominating_set
https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/

最大独立集见后面的「最大团」

独立集+顶点覆盖 https://codeforces.com/problemset/problem/1470/D

最大匹配+最小边覆盖=n （图中无孤立点）
最大独立集+最小顶点覆盖=n https://www.geeksforgeeks.org/vertex-cover-problem-set-1-introduction-approximate-algorithm-2/
    最大独立集与最小顶点覆盖互为对方关于 V 的补集（V 是图的顶点集合）
Kőnig's theorem https://en.wikipedia.org/wiki/K%C5%91nig%27s_theorem_(graph_theory)
    对于二分图，最小顶点覆盖=最大匹配
    因此有：最大独立集=n-最大匹配

激光覆盖转换成最小顶点覆盖 http://poj.org/problem?id=3041
不是 n-匹配就是 n-独立集 https://codeforces.com/problemset/problem/1198/C

DAG 上的最小路径覆盖，要求路径之间不相交，即每个顶点恰好被覆盖一次（路径长度可以为 0，即一个点）
具体定义见进阶指南 p.436
另见挑战 p.272
起初把每个点都视作一条路径，这样共有 n 条不相交路径
拆点图：每个点拆成出点和入点，原图的 v->w 相当于拆点图的 出点v->入点w，这说明拆点图一定是二分图
在拆点图里找到一条匹配边就相当于把两条路径合成了一条路径，也就相当于路径数减少了 1
所以找到了多少匹配边，路径数就减少了多少
所以有最小路径覆盖=原图的结点数-拆点图的最大匹配数

允许路径相交的做法见进阶指南 p.437

树上的最小路径覆盖见 graph_tree.go
*/

// 二分图最大匹配 - 匈牙利算法/增广路算法 O(nm)    Hungarian algorithm
// 注：使用 Dinic 可以达到 O(m√n) 的复杂度
// 【推荐】可视化 https://visualgo.net/zh/matching
//        选择「图示 - CP4 3.11a*」，然后选择「增广路 - 标准」
// https://www.renfei.org/blog/bipartite-matching.html 推荐
// https://oi-wiki.org/topic/graph-matching/bigraph-match/
// https://zhuanlan.zhihu.com/p/62981901
// https://en.wikipedia.org/wiki/Hall%27s_marriage_theorem
// https://www.geeksforgeeks.org/maximum-bipartite-matching/
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BipartiteMatching.java.html
// 有关随机贪心(匹配)预处理的 hack https://bzoj.blog.uoj.ac/blog/2427
//
// 模板题 https://www.luogu.com.cn/problem/P3386
// LCP04 https://leetcode-cn.com/problems/broken-board-dominoes/
// LC2123 https://leetcode.cn/problems/minimum-operations-to-remove-adjacent-ones-in-matrix/
//【网络流 24 题】飞行员配对方案 https://loj.ac/p/6000 https://www.luogu.com.cn/problem/P2756
//【网络流 24 题】骑士共存（这题 Dinic 更快）https://loj.ac/p/6226 https://www.luogu.com.cn/problem/P3355
// todo https://codeforces.com/contest/1404/problem/E
//      https://codeforces.com/problemset/problem/1783/F 题解 https://www.luogu.com.cn/blog/DaiRuiChen007/CF1783F
func (*graph) maxBipartiteMatchingHungarian(g [][]int) (match []int, cnt int) {
	match = make([]int, len(g))
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
			used = make([]bool, len(g))
			if f(v) {
				cnt++ // +=2
			}
		}
	}
	return
}

// 匈牙利算法的另一种写法，适用左右两侧节点有明确区分的情况，要求 g 中存储的是左侧到右侧的单向边
func (*graph) maxBipartiteMatchingHungarianLR(nl, nr int, g [][]int) (matchL []int, cnt int) {
	// matchL[leftNode] = rightNode
	// matchR[rightNode] = leftNode
	matchL = make([]int, nl)
	matchR := make([]int, nr)
	for i := range matchL {
		matchL[i] = -1
	}
	for i := range matchR {
		matchR[i] = -1
	}
	var used []bool
	var f func(int) bool
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
		used = make([]bool, nl) // 每次 DFS 之前都要重置 used
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
// todo https://codeforces.com/problemset/problem/1630/F
func (*graph) maxBipartiteMatchingHopcroftKarp(n int, g [][]int) (match []int, cnt int) {
	panic("todo")
}

// 带权二分图最大完美匹配 - 任务分配问题/婚姻匹配问题 - KM (Kuhn–Munkres) 算法
// 注意：下面的代码是 O(n^4) 的, O(n^3) 的在后面
// https://en.wikipedia.org/wiki/Assignment_problem
// https://en.wikipedia.org/wiki/Hungarian_algorithm
// https://oi-wiki.org/topic/graph-matching/bigraph-weight-match/
// 直观理解 KM 算法 https://www.cnblogs.com/wenruo/p/5264235.html
// 讲解+题单 https://www.luogu.com.cn/blog/suxxsfe/xiong-ya-li-suan-fa
// https://resources.mpi-inf.mpg.de/departments/d1/teaching/ss12/AdvancedGraphAlgorithms/Slides06.pdf
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/AssignmentProblem.java.html
// 模板题 https://www.luogu.com.cn/problem/P6577
// https://www.luogu.com.cn/problem/P3967
// 【网络流 24 题】分配问题 https://loj.ac/p/6012 https://www.luogu.com.cn/problem/P4014
// http://acm.hdu.edu.cn/showproblem.php?pid=2426
// EXTRA: 带权二分图最小边覆盖
// 转换成带权二分图最大匹配 https://cstheory.stackexchange.com/questions/14690/reducing-a-minimum-cost-edge-cover-problem-to-minimum-cost-weighted-bipartie-per
// LC1066 https://leetcode.cn/problems/campus-bikes-ii/
// LC1595 https://leetcode-cn.com/problems/minimum-cost-to-connect-two-groups-of-points/solution/kai-kai-yan-jie-zhuan-huan-cheng-zui-da-dai-quan-p/
// LC1879 https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/
// LC1947 https://leetcode-cn.com/problems/maximum-compatibility-score-sum/
// LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/
// LC2403 https://leetcode.cn/problems/minimum-time-to-kill-all-monsters/
// todo GCJ21 Round2D https://codingcompetitions.withgoogle.com/codejam/round/0000000000435915/00000000007dc2de
func (*graph) maxWeightedBipartiteMatchingKuhnMunkresSlow(wt [][]int64) (match []int, sum int64) {
	const inf int64 = 1e18
	// NOTE: wt 中不存在的边应初始化为 -inf

	match = make([]int, len(wt)) // 右部点匹配了哪一个左部点
	for i := range match {
		match[i] = -1
	}
	// 初始化顶标
	la := make([]int64, len(wt))
	for i, r := range wt {
		la[i] = r[0]
		for _, w := range r[1:] {
			if w > la[i] {
				la[i] = w
			}
		}
	}
	lb := make([]int64, len(wt))
	slack := make([]int64, len(wt))
	for i := 0; i < len(wt); i++ {
		for { // 循环直到 DFS 找到一个匹配
			va := make([]bool, len(wt)) // 访问标记：是否在交错树中
			vb := make([]bool, len(wt))
			for j := range slack {
				slack[j] = inf
			}
			var f func(int) bool
			f = func(v int) bool {
				va[v] = true
				for w, b := range vb {
					if !b {
						if delta := la[v] + lb[w] - wt[v][w]; delta == 0 { // 相等子图
							vb[w] = true
							if match[w] == -1 || f(match[w]) {
								match[w] = v
								return true
							}
						} else if delta < slack[w] {
							slack[w] = delta
						}
					}
				}
				return false
			}
			if f(i) {
				break
			}
			// 更新顶标
			delta := inf
			for j, b := range vb {
				if !b && slack[j] < delta {
					delta = slack[j]
				}
			}
			for j := 0; j < len(wt); j++ {
				if va[j] {
					la[j] -= delta
				}
				if vb[j] {
					lb[j] += delta
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
		sum += wt[v][w]
	}
	return
}

// O(n^3)
// 下标需要从 1 开始
func (*graph) maxWeightedBipartiteMatchingKuhnMunkres(wt [][]int64) (match []int, sum int64) {
	const inf int64 = 1e18
	// NOTE: wt 中不存在的边应初始化为 -inf

	match = make([]int, len(wt)) // 右部点匹配了哪一个左部点
	la := make([]int64, len(wt))
	for i, row := range wt {
		la[i] = -inf
		for _, v := range row {
			if v > la[i] {
				la[i] = v
			}
		}
	}
	lb := make([]int64, len(wt))
	slack := make([]int64, len(wt))
	for i := 1; i < len(wt); i++ {
		vb := make([]bool, len(wt))
		for j := 1; j < len(wt); j++ {
			slack[j] = inf
		}
		last := make([]int, len(wt)) // 右部点在交错树中的上一个右部点，用于倒推得到交错路
		y := 0
		match[0] = i // 一开始假设有一条 i-0 的匹配
		for {
			vb[y] = true
			x, nextY := match[y], 0
			delta := inf
			for j := 1; j < len(wt); j++ {
				if !vb[j] {
					if d := la[x] + lb[j] - wt[x][j]; d < slack[j] {
						slack[j] = d
						last[j] = y
					}
					if slack[j] < delta {
						delta = slack[j]
						nextY = j
					}
				}
			}
			// 当 delta=0 时，相当于沿着相等子图向下搜索一层
			// 当 delta>0 时，相当于直接回到最小边（新加入相等子图的边）处开始搜索
			if delta > 0 {
				for j := 0; j < len(wt); j++ {
					if vb[j] {
						la[match[j]] -= delta
						lb[j] += delta
					} else {
						slack[j] -= delta
					}
				}
			}
			y = nextY
			if match[y] == 0 {
				break
			}
		}
		// 倒推更新增广路
		for ; y > 0; y = last[y] {
			match[y] = match[last[y]]
		}
	}
	for w := 1; w < len(wt); w++ {
		v := match[w]
		// 无解，或者不选
		if v == 0 {
			//continue
			return nil, 0
		}
		sum += wt[v][w]
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
// 模板题 https://www.luogu.com.cn/problem/P6113 https://uoj.ac/problem/79
// https://www.luogu.com.cn/problem/P4258
// https://www.luogu.com.cn/problem/P6699

// EXTRA: 完美匹配 Perfect Match     MWPM
// 完美匹配同时也是一个原图的最小边数的边覆盖

// EXTRA: Min Cost Perfect Matching (MCPM)
// https://courses.engr.illinois.edu/cs598csc/sp2010/Lectures/Lecture11.pdf

// 有向图的拓扑排序 Kahn's algorithm
// 可以用来判断有向图是否有环、求 DAG 上的 DP 等
// https://oi-wiki.org/graph/topo/
// https://cp-algorithms.com/graph/topological-sort.html
// 树上拓扑+记录变成叶子的时间 LC2603 https://leetcode.cn/problems/collect-coins-in-a-tree/
// DAG DP LC2050 https://leetcode-cn.com/problems/parallel-courses-iii/
//        LC1857 https://leetcode-cn.com/problems/largest-color-value-in-a-directed-graph/
//        https://ac.nowcoder.com/acm/contest/6384/C
//        https://www.luogu.com.cn/problem/P3387
//        https://codeforces.com/problemset/problem/721/C
// 删点 LC310 https://leetcode.cn/problems/minimum-height-trees/
// 好题 https://codeforces.com/problemset/problem/915/D
// 关键点 次关键点 https://codeforces.com/contest/1062/problem/F
// 混合图拓扑排序+定向 https://codeforces.com/problemset/problem/1385/E
//                  https://codeforces.com/problemset/problem/1100/E
// 构造 https://codeforces.com/problemset/problem/269/C
// 缩点后的拓扑序 https://codeforces.com/contest/1463/problem/E
// 与堆结合 https://codeforces.com/problemset/problem/825/E
//         https://codeforces.com/problemset/problem/1283/F
// 利用拓扑思想 https://codeforces.com/problemset/problem/1037/E
// https://codeforces.com/problemset/problem/1572/A
// 拓扑序是否唯一：任意时刻队列中不能有超过一个元素
// - LC444 https://leetcode.cn/problems/sequence-reconstruction/
// 检查一个序列是否为拓扑序，可以仿造拓扑排序的算法，从前往后检查节点的入度是否为 0，然后减少相邻节点的入度，直到找到一个入度不为 0 的点或者遍历到末尾
func (*graph) topSort(in io.Reader, n, m int) []int {
	g := make([][]int, n)
	deg := make([]int, len(g))
	for i := 0; i < m; i++ {
		v, w := 0, 0
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w) // 注意特判自环的情况
		deg[w]++
	}

	//fa := make([]int, len(g))
	//for i := range fa {
	//	fa[i] = -1
	//}
	//levels := make([]int, len(g))

	q := make([]int, 0, len(g))
	orders := q
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
			//levels[i] = 1
			// NOTE: 若起点有特殊性，可以在这里初始化 dp
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// update dp[v]...

		for _, w := range g[v] {
			// dp[w] = max(dp[w], dp[v])

			if deg[w]--; deg[w] == 0 {
				//fa[w] = v
				//levels[w] = levels[v] + 1
				q = append(q, w)
			}
		}
	}

	// NOTE: 若 cap(q) 大于 0 则说明图中有环
	orders = orders[:len(g)-cap(q)]

	// NOTE: 若要重复求拓扑排序记得拷贝一份 deg

	{
		fa := make([]int, len(g))

		// EXTRA: path from end to start
		var end = len(g) - 1
		path := make([]int, 0, len(g))
		for v := end; v != -1; v = fa[v] {
			path = append(path, v)
		}
	}

	return orders
}

// 强连通分量分解 Strongly Connected Component (SCC)
// https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
// https://oi-wiki.org/graph/scc/#kosaraju
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KosarajuSharirSCC.java.html
// 模板题 https://www.luogu.com.cn/problem/P1726
//       https://atcoder.jp/contests/practice2/tasks/practice2_g
// https://www.luogu.com.cn/problem/P2341
// 建图转换 https://codeforces.com/problemset/problem/1239/D
// 与高斯消元结合 https://www.luogu.com.cn/problem/P6030
func (*graph) sccKosaraju(n, m int) ([][]int, []int) {
	type edge struct{ v, w int }
	edges := make([]edge, 0, m) // 缩点用
	g := make([][]int, n)
	rg := make([][]int, len(g))
	addEdge := func(v, w int) {
		//v--
		//w--
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
		edges = append(edges, edge{v, w})
	}

	// 生成 DFS 后序（用于跑逆后序遍历，这样生成的 SCC 一定是拓扑序）
	vs := make([]int, 0, len(g))
	vis := make([]bool, len(g))
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
		vs = append(vs, v)
	}
	for i, b := range vis {
		if !b {
			dfs(i)
		}
	}

	vis = make([]bool, len(g))
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
	scc := [][]int{}
outer:
	for i := len(g) - 1; i >= 0; i-- { // 逆后序遍历，就可以像无向图那样求出 SCC
		if v := vs[i]; !vis[v] {
			comp = []int{}
			rdfs(v)
			scc = append(scc, comp)

			// EXTRA: 无需实际缩点，判断缩点后是否出度为 0
			{
				for _, u := range comp {
					for _, w := range g[u] {
						if !vis[w] { // 出度不为 0
							continue outer
						}
					}
				}
				// 出度为 0
			}
		}
	}

	// 记录每个点所属 SCC 的下标，用于缩点和查询
	sid := make([]int, len(g))
	for i, cc := range scc {
		for _, v := range cc {
			sid[v] = i
		}
	}

	// EXTRA: 缩点: 将边 v-w 转换成 sid[v]-sid[w]
	// 缩点后得到了一张 DAG，点的编号范围为 [0,len(scc)-1]
	// 注意这样可能会产生重边，不能有重边时可以用 map 或对每个点排序去重
	// 模板题 点权 https://www.luogu.com.cn/problem/P3387
	// 		 边权 https://codeforces.com/contest/894/problem/E
	// 检测路径是否可达/唯一/无穷 https://codeforces.com/problemset/problem/1547/G
	ns := len(scc)
	g2 := make([][]int, ns)
	deg := make([]int, ns)
	for _, e := range edges {
		if v, w := sid[e.v], sid[e.w]; v != w {
			g2[v] = append(g2[v], w)
			deg[w]++
		} else {
			// 这里可以记录自环（指 len(scc) == 1 但是有自环）、汇合同一个 SCC 的权值等 ...

		}
	}

	// EXTRA: 求有多少个点能被其他所有点访问到
	// https://www.luogu.com.cn/problem/P2341
	// 能到其它所有点 https://codeforces.com/problemset/problem/1777/E
	// - Tarjan 写法 https://codeforces.com/problemset/submission/1777/204187501
	numCanBeVisitedFromAll := func() int {
		lastComp := scc[len(scc)-1]
		vis = make([]bool, len(g))
		rdfs(lastComp[0]) // 在反图上遍历
		for _, b := range vis {
			// 原图不是连通的
			if !b {
				return 0
			}
		}
		return len(lastComp)
	}
	_, _ = addEdge, numCanBeVisitedFromAll
	return scc, sid
}

// SCC Tarjan
// 常数比 Kosaraju 略小（在 AtCoder 上的测试显示，5e5 的数据下比 Kosaraju 快了约 100ms）
// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
// https://oi-wiki.org/graph/scc/#tarjan
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TarjanSCC.java.html
// https://stackoverflow.com/questions/32750511/does-tarjans-scc-algorithm-give-a-topological-sort-of-the-scc
// 与最小割结合 https://www.luogu.com.cn/problem/P4126
func (*graph) sccTarjan(g [][]int, min func(int, int) int) ([][]int, []int) {
	scc := [][]int{}
	dfn := make([]int, len(g)) // 值从 1 开始
	dfsClock := 0
	stk := []int{} // 注意这不是递归栈，单纯地将搜索到的节点组成一个先进先出的数据结构
	inStk := make([]bool, len(g))
	var f func(int) int
	f = func(v int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		stk = append(stk, v)
		inStk[v] = true
		for _, w := range g[v] {
			if dfn[w] == 0 {
				lowW := f(w)
				lowV = min(lowV, lowW)
			} else if inStk[w] { // 找到 v 的到其祖先节点的边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		if dfn[v] == lowV { // 回不去了，再也回不去了，祖先（节点）已成历史
			comp := []int{}
			for {
				w := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				// 避免搜索树上的另一棵子树上的点 v，通过横向边，把 dfn[w] 错误地更新到 lowV（注意 dfn[w] 都小于后面新遍历到的点的 dfn 值）
				inStk[w] = false
				comp = append(comp, w)
				if w == v {
					break
				}
			}
			scc = append(scc, comp)
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			f(v)
		}
	}

	// 由于每个强连通分量都是在它的所有后继强连通分量被求出之后求得的
	// 上面得到的 scc 是拓扑序的逆序
	for i, n := 0, len(scc); i < n/2; i++ {
		scc[i], scc[n-1-i] = scc[n-1-i], scc[i]
	}

	sid := make([]int, len(g))
	for i, cp := range scc {
		for _, v := range cp {
			sid[v] = i
		}
	}

	return scc, sid
}

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
// NOTE: 一些建边的转换：
//       A 为真          (A)     ¬A⇒A     注：A ⇔ A∨A ⇔ ¬A⇒A∧¬A⇒A ⇔ ¬A⇒A
//       A 为假          (¬A)    A⇒¬A
//       A 为真 B 就为真          A⇒B, ¬B⇒¬A
//       A 为假 B 就为假          ¬A⇒¬B, B⇒A
//       A,B 至少存在一个 (A|B)    ¬A⇒B, ¬B⇒A 意思是一个为假的时候，另一个一定为真 https://www.luogu.com.cn/problem/P4782
//       A,B 不能同时存在 (¬A|¬B)  A⇒¬B, B⇒¬A 就是上面的式子替换了一下（一个为真，另一个一定为假）
//       A,B 必须且只一个 (A^B)    A⇒¬B, B⇒¬A, ¬A⇒B, ¬B⇒A
//       A,B 同时或都不在 (¬(A^B)) A⇒B, B⇒A, ¬A⇒¬B, ¬B⇒¬A
// NOTE: 单独的条件 x为a 可以用 (x为a)∨(x为a) 来表示
// 模板题 https://www.luogu.com.cn/problem/P4782
// 建边练习【模板代码】 https://codeforces.com/contest/468/problem/B
// 定义 Ai 表示「选 Xi」，这样若两个旗子 i j 满足 |Xi-Xj|<D 时，就相当于 Ai Aj 至少一个为假。其他情况类似 https://atcoder.jp/contests/practice2/tasks/practice2_h
func (G *graph) solve2SAT(n, m int) []bool {
	// 分为左右两部，左边 [0,n) 范围的点表示 x 为真，右边 [n,2*n) 范围的点表示 x 为假（¬x 用 x+n 表示）
	// 例如，当 x y 均为真时，就连一条 a
	g := make([][]int, n*2)
	rg := make([][]int, len(g))

	// x=a 和 y=b 两个条件至少满足一个（a b 为 0/1 表示 假/真）
	// 见 https://www.luogu.com.cn/problem/P4782
	addEdge := func(x, a, y, b int) {
		//x--
		//y--
		v, w := x+(a^1)*n, y+b*n // ¬A⇒B
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
		v, w = y+(b^1)*n, x+a*n // ¬B⇒A
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	// 读图，求 sid ...
	_, sid := G.sccKosaraju(n*2, m)
	ans := make([]bool, n)
	for i, id := range sid[:n] {
		// x 和 ¬x 处于同一个 SCC 时无解（因为 x ⇔ ¬x）
		if id == sid[i+n] {
			return nil
		}
		// sid[x] < sid[¬x] ⇔ (¬x ⇒ x) ⇔ x 为真
		// sid[x] > sid[¬x] ⇔ (x ⇒ ¬x) ⇔ x 为假
		ans[i] = id < sid[i+n]
	}

	_ = addEdge
	return ans
}

// 基环树（环套树），英文名叫 pseudotree，基环树森林叫 pseudoforest
// https://en.wikipedia.org/wiki/Pseudoforest
// 对于内向基环树，由于每个点的出度均为一，可以用 []int 来表示图
// 基环树的讲解可以看一下我的这篇题解 https://leetcode-cn.com/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/
// todo https://www.luogu.com.cn/blog/user52918/qian-tan-ji-huan-shu
// todo 题单 https://www.luogu.com.cn/blog/ShadderLeave/ji-huan-shu-bi-ji
//
// LC2127 https://leetcode-cn.com/problems/maximum-employees-to-be-invited-to-a-meeting/
// LC2359 单源最短路 https://leetcode.cn/problems/find-closest-node-to-given-two-nodes/
// LC2360 最长环 https://leetcode.cn/problems/longest-cycle-in-a-graph/
// 计数 https://codeforces.com/problemset/problem/711/D
// https://codeforces.com/problemset/problem/1027/D
// https://codeforces.com/problemset/problem/1335/F
// 拆点 https://codeforces.com/problemset/problem/1200/F
// https://codeforces.com/contest/1770/problem/D
// https://atcoder.jp/contests/abc266/tasks/abc266_f
// 删除一条边使得直径最长 https://ac.nowcoder.com/acm/contest/9977/c
// todo [IOI2008] 岛屿 https://www.luogu.com.cn/problem/P4381
// todo [NOI2013] 快餐店 https://www.luogu.com.cn/problem/P1399
func (*graph) pseudotree(g []int) { // g 为内向基环树（森林）
	rg := make([][]int, len(g)) // g 的反图（外向基环树）
	deg := make([]int, len(g))  // g 上每个节点的入度
	for v, w := range g {
		rg[w] = append(rg[w], v)
		deg[w]++
	}

	// 拓扑排序，剪掉 g 上的所有树枝
	// 拓扑排序后 deg 值为 1 的点必定在基环上，为 0 的点必定在树枝上
	// 注：拓扑排序时还可以做 DP，比如给树枝上的每个点标记反向深度
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		//dp[v]++
		w := g[v] // v 只有一条出边
		//dp[w] = max(dp[w], dp[v])
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	// 在反图上遍历树枝
	var rdfs func(int)
	rdfs = func(v int) {
		// ...

		for _, w := range rg[v] {
			if deg[w] == 0 { // 树枝上的点在拓扑排序后，入度均为 0

				rdfs(w)

			}
		}
	}

	for i, d := range deg {
		if d <= 0 {
			continue
		}
		// 遍历基环上的点（拓扑排序后入度大于 0）
		ring := []int{}
		for v := i; ; v = g[v] {
			deg[v] = -1 // 将基环上的点的入度标记为 -1，避免重复访问
			ring = append(ring, v)
			if g[v] == i {
				break
			}
		}
		// do ring ...
		// 特别注意基环大小小于 3 的特殊情况

	}

	{
		// EXTRA: 无向图的情况
		// 去掉树枝后，剩余点的度数均超过 1
		var n int
		g := make([][]int, n)
		deg := make([]int, len(g))
		//for i := 0; i < len(g); i++ {
		//  var v, w int
		//	Fscan(in, &v, &w)
		//	v--
		//	w--
		//	g[v] = append(g[v], w)
		//	g[w] = append(g[w], v)
		//	deg[v]++
		//	deg[w]++
		//}

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

		{
			// EXTRA：给每个树枝编号
			ids := make([]int, n)
			id := 0
			var f func(int, int)
			f = func(v, fa int) {
				ids[v] = id
				for _, w := range g[v] {
					if w != fa && deg[w] < 2 {
						f(w, v)
					}
				}
			}
			for i, d := range deg {
				if d > 1 {
					id++
					f(i, -1)
				}
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

/* 网络流·总结·题单 ################################################################################

最大流等于最小割的证明 https://seineo.github.io/%E5%9B%BE%E8%AE%BA%EF%BC%9A%E6%9C%80%E5%A4%A7%E6%B5%81%E6%9C%80%E5%B0%8F%E5%89%B2%E8%AF%A6%E8%A7%A3.html

todo 网络流建模方式总结
 https://www.cnblogs.com/victorique/p/8560656.html
 https://blog.bill.moe/network-flow-models/
 NOI 一轮复习 I：二分图网络流 https://www.luogu.com.cn/blog/ix-35/noi-yi-lun-fu-xi-i-er-fen-tu-wang-lao-liu
 2016 国家集训队论文《网络流的一些建模方法》姜志豪 https://github.com/enkerewpo/OI-Public-Library/blob/master/IOI%E4%B8%AD%E5%9B%BD%E5%9B%BD%E5%AE%B6%E5%80%99%E9%80%89%E9%98%9F%E8%AE%BA%E6%96%87/%E5%9B%BD%E5%AE%B6%E9%9B%86%E8%AE%AD%E9%98%9F2016%E8%AE%BA%E6%96%87%E9%9B%86.pdf

todo 网络流 24 题 https://loj.ac/p?tagIds=30 https://www.luogu.com.cn/problem/list?tag=332
 线性规划与网络流 24 题 解题报告 https://byvoid.com/zhs/blog/lpf24-solution/

todo 题单 https://www.zybuluo.com/xzyxzy/note/992041
 网络流从入门到入土 #1 https://www.luogu.com.cn/training/12097#problems
 网络流从入门到入土 #2 https://www.luogu.com.cn/training/12098#problems
 网络流从入门到入土 #3 https://www.luogu.com.cn/training/12099#problems
 网络流建模经典题 https://www.luogu.com.cn/training/1230#problems
 网络流经典题目 https://www.luogu.com.cn/training/3144#problems

Max-Flow in almost linear time https://codeforces.com/blog/entry/100510

CF Tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=flows
*/

/* 最大流·建模·转换 ################################################################################

可视化 https://visualgo.net/zh/maxflow

https://en.wikipedia.org/wiki/Maximum_flow

建模·转换
将点拆为入点和出点（v 和 v+n），即可把点上的约束变成边上的约束
https://www.luogu.com.cn/problem/P2891 http://poj.org/problem?id=3281
【网络流 24 题】最长不降子序列 https://loj.ac/p/6005 https://www.luogu.com.cn/problem/P2766
    注意这题用到了操纵超级源点的技巧：容量限制与解除容量限制
NWERC07 B https://codeforces.com/gym/100723 http://poj.org/problem?id=3498 UVa12125 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=243&page=show_problem&problem=3277
网格模型 https://codeforces.com/problemset/problem/1360/G
https://codeforces.com/problemset/problem/546/E
转换 https://www.acwing.com/problem/content/2239/ http://poj.org/problem?id=1149
转换 https://codeforces.com/problemset/problem/653/D
todo 转换 https://atcoder.jp/contests/arc085/tasks/arc085_c

顶点上有容量
将顶点拆成两个（入顶点 x 和出顶点 y），入点向 x 连边，y 向出点连边，x 向 y 连边，容量为顶点的容量

无向图
视作两条容量均为 cap 的有向边（具体实现见下面代码中 addEdge 的注释）

多源汇最大流
建立超级源点 S 和超级汇点 T，S 向所有源点连边，所有汇点向 T 连边，每条边的容量为 inf 或对应源汇的容量限制
https://www.acwing.com/problem/content/2236/

只能经过这条边一次 ⇔ 容量为 1
http://poj.org/problem?id=2455 https://www.acwing.com/problem/content/2279/

上下界可行流·总结
https://oi-wiki.org/graph/flow/bound/
https://www.acwing.com/solution/content/17067/
https://zhuanlan.zhihu.com/p/324507636
todo 题单 https://www.luogu.com.cn/training/8462

无源汇上下界可行流（循环流）
假设存在一个流量守恒的解 f，通过将每条边的流量减去 low，得到一个新图的流，但其不一定满足流量守恒
对于每个顶点 v，记 d(v) = ∑lowIn(v) - ∑lowOut(v)
- 若 d(v) > 0，说明流入减去的更多，则需将 v 的流入量增加 d(v)，这可以通过新建超级源点 S，并增加 S->v，容量为 d(v) 的边做到
- 若 d(v) < 0，说明流出减去的更多，则需将 v 的流出量增加 d(v)，这可以通过新建超级汇点 T，并增加 v->T，容量为 -d(v) 的边做到
跑从 S 到 T 的最大流，若满流（即最大流等于从 S 出发的容量之和），则说明可以让新图的流量守恒，从而说明原图存在可行流 f，其每条边的流量为 low 加上新图中每条边的流量；若不满流则无解
模板题 https://loj.ac/p/115 https://www.acwing.com/problem/content/2190/

有源汇上下界可行流
从汇点向源点连一条容量为 inf 的边，即转换成了无源汇上下界可行流

有源汇上下界最大流
1. 跑一遍有源汇上下界可行流，若有解，记此时源点到汇点的流量为 f1（通过汇点向源点的反向边的流量得到）
2. 删去汇点到源点的边（或将其容量置为 0，具体实现时可以将汇点->源点边最后加入，或者使用指针记录该边及其反向边）
3. 在残余网络上继续增广，记额外的最大流为 f2，那么答案即为 f1+f2
模板题 https://loj.ac/p/116 https://www.luogu.com.cn/problem/P5192

有源汇上下界最小流
将上面第 3 步改成退流，即减去残余网络上从汇点到源点的最大流
模板题 https://loj.ac/p/117 https://www.luogu.com.cn/problem/P4843

分层图
注意：可以在原图的基础上添加边/增加容量，然后继续寻找增广路增广
【网络流 24 题】星际转移 https://loj.ac/p/6015 https://www.luogu.com.cn/problem/P2754

关键边
关键边 v-w 需满足，在跑完最大流后：
1. 这条边的流量等于其容量
2. 在残余网络上，从源点可以到达 v，从 w 可以到达汇点（即从汇点顺着反向边可以到达 w）
http://poj.org/problem?id=3204 https://www.acwing.com/problem/content/2238/
具体实现见下面代码中的 EXTRA
*/

/* 最小割·建模·转换 ################################################################################

https://en.wikipedia.org/wiki/Max-flow_min-cut_theorem
最小割模型汇总 https://blog.csdn.net/qq_35649707/article/details/77482691
下面的 topic 参考胡伯涛《最小割模型在信息学竞赛中的应用》（PDF 见 https://github.com/EndlessCheng/cp-pdf）

求出最大流后，从源点出发在残余网络上 DFS，标记所有能够到达的点。遍历原边集 edges，若其中一端有标记，另一端没有标记，则这条边为最小割上的边

常用技巧：用容量为 inf 的边来防止割断

建模·转换
https://www.acwing.com/problem/content/2282/
平均边权最小 https://www.acwing.com/problem/content/2281/
点连通度 SEERC04 F https://codeforces.com/gym/101461 http://poj.org/problem?id=1966 UVa1660 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4535
   https://en.wikipedia.org/wiki/Connectivity_(graph_theory)
   https://en.wikipedia.org/wiki/Menger%27s_theorem
LCP38/21春·战队赛F https://leetcode-cn.com/problems/7rLGCR/

最大权闭合图 Maximum Weight Closure of a Graph
https://en.wikipedia.org/wiki/Closure_problem
源点向所有正权点连边，容量为相应点权
所有负权点向汇点连边，容量为相应点权的相反数
原图边的容量为 inf（从而保证不会在最小割中）
最后用正权点总和减去源点到汇点的最小割即为答案
以「最大获利」这题来解释，割掉源点到正权点的边，意味着放弃对应用户的收益；割掉负权点到汇点的边，意味着建立对应基站
NOI06 最大获利 https://www.luogu.com.cn/problem/P4174
【网络流 24 题】太空飞行计划 https://loj.ac/p/6001 https://www.luogu.com.cn/problem/P2762

最大密度子图 Maximum Density Subgraph
https://en.wikipedia.org/wiki/Dense_subgraph
参考 https://www.luogu.com.cn/problem/solution/UVA1389
二分上下界：最小密度为 1/n，最大密度为 m
二分精度：任意两个密度不同的子图，其密度差 >= 1/n^2
todo NEERC06 H https://codeforces.com/gym/100287 https://codeforces.com/gym/100532 http://poj.org/problem?id=3155 UVa1389 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=446&page=show_problem&problem=4135

二分图最小点权覆盖集 Minimum Weight Vertex Covering Set (MinWVCS) in a Bipartite Graph
二分图最大点权独立集 Maximum Weight Vertex Independent Set (MaxWVIS) in a Bipartite Graph
建立一个源 s，向 X 部每个点连边；建立一个汇 t，从 Y 部每个点向汇 t 连边，把二分图中的边看成是有向的，
则任意一条从 s 到 t 的路径，一定具有 s->v->w->t 的形式（v∈X, w∈Y）。
割的性质是不存在一条从 s 到 t 的路径。故路径上的三条边 s-v, v-w, w-t 中至少有一条边在割中。
若人为地令 v-w 不可能在最小割中，即令其容量为正无限，
可将条件简化为 s-v, w-t 中至少有一条边在最小割中，这正好与点覆盖集限制条件的形式相符（边的两端点中至少一个在覆盖集内），
而目标是最小化点权之和，这恰好也是最小割的优化目标。
对于最大点权独立集，其等价于点权之和减去最小点权覆盖集。
【网络流 24 题】骑士共存 https://loj.ac/p/6226 https://www.luogu.com.cn/problem/P3355
todo https://codeforces.com/contest/808/problem/F
NEERC03 D https://codeforces.com/gym/100725 https://codeforces.com/gym/101651 http://poj.org/problem?id=2125
黑白染色转化成二分图 https://www.acwing.com/problem/content/2328/

最小割的可行边和必须边（所有最小割集的并集和交集）
跑最大流，然后求整个残余网络的 SCC，则有：
- 可行边：两端不在一个 SCC 内，即不存在另一条从 v 到 w 的路径
- 必须边：一端在 S 的 SCC 内,另一端在 T 的 SCC 内
AHOI09 https://www.luogu.com.cn/problem/P4126
*/

/* 费用流·建模·转换 ################################################################################

https://en.wikipedia.org/wiki/Minimum-cost_flow_problem MCFP
https://en.wikipedia.org/wiki/Assignment_problem
https://en.wikipedia.org/wiki/Network_simplex_algorithm

NOTE: 对于修改容量的情况，由于 EK 是基于最短路的贪心算法，不能像最大流那样直接在残余网络上继续跑，必须重新建图重新跑 EK
todo https://codeforces.com/problemset/problem/362/E

建模·转换
从源点连容量为 1 费用为 0 的边到集合 A 中各点
从集合 B 中各点连容量为 1 费用为 0 的边到汇点
集合 A 和 B 之间连边，容量为 inf，费用为 F(Ai,Bj)，F 根据题意
这样跑 MCMF 得到的结果是匹配全部 A（或 B）的最小花费
LC2172 https://leetcode-cn.com/problems/maximum-and-sum-of-array/
https://codeforces.com/problemset/problem/1437/C
【网络流 24 题】运输问题 https://loj.ac/p/6011 https://www.luogu.com.cn/problem/P4015
【网络流 24 题】数字梯形 https://loj.ac/p/6010 https://www.luogu.com.cn/problem/P4013
【网络流 24 题】深海机器人 https://loj.ac/p/6224 https://www.luogu.com.cn/problem/P4012
k 取方格数 https://www.luogu.com.cn/problem/P2045 http://poj.org/problem?id=3422
    关键技巧：拆点时，从入点向出点连两条边，第一条边容量为 1，费用为点权，第二条边容量为 k-1，费用为 0
    这表示第一次经过该点时，可以把数取走，之后再经过时就不再计算
【网络流 24 题】餐巾计划 https://loj.ac/p/6008 https://www.luogu.com.cn/problem/P1251

最大费用
将每条边的费用反向，答案即为 -MCMF

无源汇上下界最小费用可行流
建图和上面的「无源汇上下界可行流」一样
NOI08 志愿者招募 https://www.luogu.com.cn/problem/P3980（也可以用线性规划做）
- 由于没有上界，建图的时候可以不用减去下界
- 把每天的人数要求看成是边的流量下界（从 i 天向 i+1 天连边）
- 由于要满足流量守恒，对于每个人 i，需要从结束日期向开始日期连边，容量为 inf，费用为 ci。这相当于每个人在流网络的一单位的流量流过了一个环
- 代码实现 https://www.luogu.com.cn/record/56398769
AHOI14/JSOI14 支线剧情 https://www.luogu.com.cn/problem/P4043
-「看完所有剧情」可以转换成每条边的流量下界为 1，容量为 inf，费用为过剧情花费的时间
-「开始新的游戏」可以转换成每个点向点 1 连边，容量为 inf，费用为 0
- 代码实现 https://www.luogu.com.cn/record/56402617

流通问题 circulation problem
最小费用流通问题 minimum-cost-circulation problem
https://en.wikipedia.org/wiki/Circulation_problem
The circulation problem and its variants are a generalisation of network flow problems,
with the added constraint of a lower bound on edge flows,
and with flow conservation also being required for the source and sink (i.e. there are no special nodes).
《算法导论》思考题 29-5
todo https://codeforces.com/contest/1455/problem/E
 https://codeforces.com/blog/entry/85186?#comment-728533
*/

// 最大流 Dinic's algorithm O(n^2 * m)  二分图上为 O(m√n)
// 如果容量是浮点数，下面代码中 > 0 的判断要改成 > eps
// https://en.wikipedia.org/wiki/Dinic%27s_algorithm
// https://oi-wiki.org/graph/flow/max-flow/#dinic
// https://cp-algorithms.com/graph/dinic.html
// 模板题 https://www.luogu.com.cn/problem/P3376 https://www.luogu.com.cn/problem/P2740
func (*graph) maxFlowDinic(in io.Reader, n, m, st, end int, min func(int, int) int) int {
	const inf int = 1e9 // 1e18
	st--
	end--

	type neighbor struct{ to, rid, cap, eid int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, eid int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, eid})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -1}) // 无向图上 0 换成 cap
	}
	for i := 0; i < m; i++ {
		var v, w, cp int
		Fscan(in, &v, &w, &cp)
		v--
		w--
		addEdge(v, w, cp, i)
	}

	var d []int // 从源点 st 出发的距离
	bfs := func() bool {
		d = make([]int, len(g))
		d[st] = 1
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && d[w] == 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d[end] > 0
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
	dinic := func() (maxFlow int) { // int64
		for bfs() {
			iter = make([]int, len(g))
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
	maxFlow := dinic()

	// EXTRA: 容量复原（不存原始容量的写法）
	for _, es := range g {
		for i, e := range es {
			if e.eid >= 0 { // 正向边
				es[i].cap += g[e.to][e.rid].cap
				g[e.to][e.rid].cap = 0
			}
		}
	}

	// EXTRA: 求流的分配方案（即反向边上的 cap）
	// https://loj.ac/p/115 https://www.acwing.com/problem/content/2190/
	ans := make([]int, m)
	for _, es := range g { // v
		for _, e := range es {
			w, i := e.to, e.eid
			if i >= 0 { // 正向边
				ans[i] = g[w][e.rid].cap
			}
		}
	}

	// EXTRA: 求关键边（扩容后可以增加最大流的边）的数量
	// 关键边 v-w 需满足，在跑完最大流后：
	// 1. 这条边的流量等于其容量
	// 2. 在残余网络上，从源点可以到达 v，从 w 可以到达汇点（即从汇点顺着反向边可以到达 w）
	// http://poj.org/problem?id=3204 https://www.acwing.com/problem/content/2238/
	{
		// 在残余网络上跑 DFS，看看哪些点能从源点和汇点访问到（从汇点出发的要判断反向边的流量）
		vis1 := make([]bool, len(g))
		var dfs1 func(int)
		dfs1 = func(v int) {
			vis1[v] = true
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && !vis1[w] {
					dfs1(w)
				}
			}
		}
		dfs1(st)

		vis2 := make([]bool, len(g))
		var dfs2 func(int)
		dfs2 = func(v int) {
			vis2[v] = true
			for _, e := range g[v] {
				if w := e.to; !vis2[w] && g[w][e.rid].cap > 0 {
					dfs2(w)
				}
			}
		}
		dfs2(end)

		ans := 0
		for v, es := range g {
			if !vis1[v] {
				continue
			}
			for _, e := range es {
				// 原图的边，流量为 0（说明该边满流），且边的两端点能分别从源汇访问到
				if e.eid >= 0 && e.cap == 0 && vis2[e.to] {
					ans++
				}
			}
		}
	}

	return maxFlow
}

// ISAP, Improved Shortest Augmenting Path O(n^2 * m)
// https://oi-wiki.org/graph/flow/max-flow/#isap
// https://www.renfei.org/blog/isap.html
// 测试了一下性能和 Dinic 差不多
func (*graph) maxFlowISAP(in io.Reader, n, m, st, end int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for i := 0; i < m; i++ {
		var v, w, cp int
		Fscan(in, &v, &w, &cp)
		v--
		w--
		addEdge(v, w, cp)
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

// 最高标号预流推进 (HLPP, High Level Preflow Push)   O(n^2 * √m)
// 注：虽然在复杂度上比增广路方法进步很多，但是预流推进算法复杂度的上界是比较紧的，因此有时差距并不会很大
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm
// https://en.wikipedia.org/wiki/Push%E2%80%93relabel_maximum_flow_algorithm#Highest_label_selection_rule
// https://oi-wiki.org/graph/flow/max-flow/#hlpp
// https://www.luogu.com.cn/blog/ONE-PIECE/jiu-ji-di-zui-tai-liu-suan-fa-isap-yu-hlpp
// 模板题 https://loj.ac/p/127 https://www.luogu.com.cn/problem/P4722
// todo deque 优化 + 全局重贴标签等 https://www.luogu.com.cn/problem/solution/P4722
type hlppHeap struct {
	sort.IntSlice
	d []int
}

func (h hlppHeap) Less(i, j int) bool { return h.d[h.IntSlice[i]] > h.d[h.IntSlice[j]] } // 处于堆中的节点的 d 值不会改变，所以可以直接比较
func (h *hlppHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hlppHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hlppHeap) push(v int)        { heap.Push(h, v) }
func (h *hlppHeap) pop() int          { return heap.Pop(h).(int) }

func (*graph) maxFlowHLPP(in io.Reader, n, m, st, end int, min func(int, int) int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for i := 0; i < m; i++ {
		var v, w, cp int
		Fscan(in, &v, &w, &cp)
		v--
		w--
		addEdge(v, w, cp)
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
	q := hlppHeap{d: d}
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

// 无向图全局最小割
// Stoer-Wagner 算法 O(nm+n^2logn)
// https://en.wikipedia.org/wiki/Stoer%E2%80%93Wagner_algorithm
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GlobalMincut.java.html
// todo 模板题 https://www.luogu.com.cn/problem/P5632 http://poj.org/problem?id=2914
func (*graph) minimumCutStoerWagner(dist [][]int) int {
	panic("todo")
}

// 最小费用流 MCFP
// 最小费用最大流 MCMF（即满流时的费用）
// 将 Edmonds-Karp 中的 BFS 改成 SPFA O(fnm) 或 Dijkstra O(fmlogn)
// 要求初始网络中无负权圈
// 性能对比（洛谷 P3381，由于数据不强所以 SPFA 很快）：SPFA 1.05s(max 365ms)   Dijkstra 1.91s(max 688ms)
// https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm
// https://oi-wiki.org/graph/flow/min-cost/
// https://cp-algorithms.com/graph/min_cost_flow.html
// 最小费用流的不完全算法博物馆 https://www.luogu.com.cn/blog/Atalod/zui-xiao-fei-yong-liu-di-fou-wan-quan-suan-fa-bo-wu-guan
// 模板题 https://www.luogu.com.cn/problem/P3381
func (*graph) minCostFlowSPFA(in io.Reader, n, m, st, end int) (int, int64) {
	const inf int = 1e9 // 1e18
	st--
	end--

	type neighbor struct{ to, rid, cap, cost, eid int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost, eid int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost, eid})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost, -1}) // 无向图上 0 换成 cap
	}
	for i := 0; i < m; i++ {
		var v, w, cp, cost int
		Fscan(in, &v, &w, &cp, &cost)
		v--
		w--
		addEdge(v, w, cp, cost, i)
	}

	dist := make([]int64, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	spfa := func() bool {
		const _inf int64 = 1e18
		for i := range dist {
			dist[i] = _inf
		}
		dist[st] = 0
		inQ := make([]bool, len(g))
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
					fa[w] = vi{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[end] < _inf
	}
	ek := func() (maxFlow int, minCost int64) {
		for spfa() {
			// 沿 st-end 的最短路尽量增广
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
			minCost += dist[end] * int64(minF)
		}
		return
	}
	return ek()
}

// 基于原始对偶方法 (primal-dual method)
// https://blog.xehoth.cc/DurationPlan-Primal-Dual/
func (*graph) minCostFlowDijkstra(in io.Reader, n, m, st, end, flowLimit int) int64 {
	st--
	end--

	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i := 0; i < m; i++ {
		var v, w, cp, cost int
		Fscan(in, &v, &w, &cp, &cost)
		v--
		w--
		addEdge(v, w, cp, cost)
	}

	h := make([]int64, len(g)) // 顶点的势
	dist := make([]int64, len(g))
	type pair struct{ v, i int }
	fa := make([]pair, len(g))
	dijkstra := func() bool {
		const _inf int64 = 1e18
		for i := range dist {
			dist[i] = _inf
		}
		dist[st] = 0
		q := dijkstraHeap{{st, 0}}
		for len(q) > 0 {
			p := q.pop()
			v := p.v
			if p.dis > dist[v] {
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
					q.push(dijkstraPair{w, newD})
				}
			}
		}
		return dist[end] < _inf
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

// 团 clique
// 图 G 中的团和图 G 的补图中的独立集是一一对应的
// 图 G 中的独立集和图 G 的补图中的团是一一对应的（等价说法）
// https://en.wikipedia.org/wiki/Clique_(graph_theory)

// 最大独立集 maximal independent set (MIS)
// 等价于在补图上找最大团 maximal cliques (MC)
// https://en.wikipedia.org/wiki/Clique_problem
// 另见 Bron–Kerbosch 算法 https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
// Measure and Conquer: A Simple O(2^0.288n) Independent Set Algorithm http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.321.6920&rep=rep1&type=pdf
// todo 剪枝写法
// https://codeforces.com/problemset/problem/1105/E
func (*graph) maximalCliques(g []int64, max func(int, int) int) int {
	// 一种求最大团的做法，适用于点数不超过 50 的图
	// 传入的 g 为状压后的邻接矩阵
	// 定义 f(s) 为 s 的所有子集中最大团的大小
	// 则转移时要么不取 lb（low bit），要么取 lb 并去掉不与 lb 相邻的点（包括 lb）
	// 将这一过程记忆化可大幅减少运行时间，理由如下：
	// 由于每次都会去掉 lb，所以至多 k=len(g)/2 次递归后会进入右半部分没有 1 的状态
	// 将这 k 次递归过程视作一棵二叉树，则其耗时为 O(2^k)
	// 之后记忆化占主导，耗时也为 O(2^k)
	// 主要注意的是，k 次递归的结果是否记忆化并不重要，因为这部分最多也只有 O(2^k) 个状态
	// 总的来说，记忆化将计算量由原来的「二叉树规模」变成了「meet in the middle 规模」
	dp := map[int64]int{0: 0}
	var f func(int64) int
	f = func(s int64) int {
		if v, has := dp[s]; has {
			return v
		}
		dp[s] = max(f(s&(s-1)), 1+f(s&g[bits.TrailingZeros64(uint64(s))]))
		return dp[s]
	}
	ans := f(1<<len(g) - 1)
	return ans
}

// todo 极大团计数

// 分团覆盖问题 Clique cover
// https://en.wikipedia.org/wiki/Clique_cover

// todo 图的同构

// 树的同构见 graph_tree.go

// 支配树
// https://en.wikipedia.org/wiki/Dominator_(graph_theory)
// todo https://www.luogu.com.cn/blog/Zenislt/qian-tan-zhi-pei-shu-lengauer-tarjan-algorithm
// todo 模板题 https://www.luogu.com.cn/problem/P5180

// 弦图：任意长度大于 3 的环都有一个弦（连接环中不相邻两点的边）的图称为弦图
// 单纯点 完美消除序列
// 最大势算法 Maximum Cardinality Search (MCS) http://www.ii.uib.no/~pinar/MCS-M.pdf
// https://oi-wiki.org/graph/chord/
// https://www.luogu.com.cn/blog/hsfzLZH1/chord-graph

// 三元环计数/判定 O(m√m)
// 三元环一定由 v-w、w-u 和 v-u 组成，暴力枚举 v-w-u，然后判断 u 是否为 v 的出点
// 因此枚举的复杂度为 ∑in[w]*out[w]
// 直接枚举肯定不行，不妨给边定向：度数大的指向度数小的，或者度数相同时编号小的指向编号大的（全序，从而得到一个 DAG）
// 考察枚举中的 w：若 deg[w]>√m，由于度数超过 √m 的点的个数至多为 O(√m)，根据定向规则，w 的入度也至多为 O(√m)；
//               若 deg[w]<=√m，由于 w 的邻边数量就是 deg[w]，所以入度也至多为 O(√m)
// 因此，这种定向方式保证了枚举中的 w 的入度至多为 O(√m)，从而总体复杂度为 ∑in[w]*out[w] < √m*∑out[w] < √m*m
// 注：若定向时从度数小的指向度数大的，同样地，w 的出度至多为 O(√m)，从而复杂度亦为 O(m√m)
// 讲解+题目 https://blog.nowcoder.net/n/9bfaeb850d6f4df6b257ffcf8e5889dd
// https://www.luogu.com.cn/blog/KingSann/fou-chang-yong-di-hei-ke-ji-san-yuan-huan-post
// https://www.luogu.com.cn/blog/i207M/san-yuan-huan-ji-shuo-xue-xi-bi-ji
// https://www.cnblogs.com/Khada-Jhin/p/10143074.html
// https://cdn.luogu.com.cn/upload/image_hosting/4ty1215p.png
// http://acm.hdu.edu.cn/showproblem.php?pid=6184
func (*graph) countCycle3(n int, edges [][2]int) (ans int64) {
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1] // -1
		deg[v]++
		deg[w]++
	}
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		if deg[v] > deg[w] || deg[v] == deg[w] && v > w {
			v, w = w, v
		}
		g[v] = append(g[v], w)
	}

	vis := make([]int, n)
	for i := range vis {
		vis[i] = -1
	}
	for v, ws := range g {
		for _, w := range ws {
			vis[w] = v // 这样写可以不用 reset vis
		}
		for _, w := range ws {
			for _, u := range g[w] {
				if vis[u] == v {
					ans++
				}
			}
		}
	}
	return
}

// 四元环计数/判定 O(m√m)
// 同三元环那样定向，将四元环拆成左右各两条边，这两条边为一条无向边和一条有向边，然后统计无向+有向的对数
// 具体参见 https://blog.csdn.net/weixin_43466755/article/details/112985722
// https://cdn.luogu.com.cn/upload/image_hosting/4ty1215p.png
// https://www.luogu.com.cn/blog/i207M/san-yuan-huan-ji-shuo-xue-xi-bi-ji
// 转换成判定 https://codeforces.com/problemset/problem/1468/M
func (*graph) countCycle4(n int, edges [][2]int) (ans int64) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1] // -1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	less := func(v, w int) bool { return deg[v] < deg[w] || deg[v] == deg[w] && v < w }

	g2 := make([][]int, n)
	for v, ws := range g {
		for _, w := range ws {
			if less(v, w) {
				g2[v] = append(g2[v], w)
			}
		}
	}

	cnt := make([]int, n)
	for v, ws := range g {
		for _, w := range ws {
			for _, u := range g2[w] {
				if less(v, u) {
					ans += int64(cnt[u])
					cnt[u]++
				}
			}
		}
		for _, w := range ws {
			for _, u := range g2[w] {
				if less(v, u) {
					cnt[u] = 0
				}
			}
		}
	}
	return
}

// 五元环
// https://blog.csdn.net/weixin_30563319/article/details/96009073
// https://nanti.jisuanke.com/t/A1644

// 拟阵 Matroid
// https://en.wikipedia.org/wiki/Matroid
// todo 拟阵与最优化问题 https://www.luogu.com.cn/blog/cpp/ni-zhen-yu-zui-you-hua-wen-ti

// 仙人掌图 Cactus graph
// https://en.wikipedia.org/wiki/Cactus_graph
// A connected graph in which any two simple cycles have at most one vertex in common
// Equivalently, it is a connected graph in which every edge belongs to at most one simple cycle
// 如果图没有偶环，则不可能有两个奇环共用一条边（因为这样会构成一个偶环），因此没有两个环共用一条边，图一定为仙人掌（注意：反过来，一个仙人掌图是可能有偶环的）
// 圆方树 https://oi-wiki.org/graph/block-forest/
// todo 模板题 静态仙人掌 https://www.luogu.com.cn/problem/P5236

// 仙人掌所有顶点所处环的最小顶点和最大顶点
// -1 表示不在环上
// https://codeforces.com/problemset/problem/901/C
// 与 FWT 结合 https://codeforces.com/problemset/problem/1218/D
func (*graph) cactusDFS(g [][]int, n int) [][2]int {
	minMax := make([][2]int, n)
	for i := range minMax {
		minMax[i] = [2]int{-1, -1}
	}
	vis := make([]int8, n)
	fa := make([]int, n)
	var f func(int)
	f = func(v int) {
		vis[v] = 1
		for _, w := range g[v] {
			if vis[w] == 0 {
				fa[w] = v
				f(w)
			} else if w != fa[v] && vis[w] == 1 {
				mi, mx := v, v
				for x := v; x != w; {
					x = fa[x]
					if x < mi {
						mi = x
					} else if x > mx {
						mx = x
					}
				}
				minMax[v] = [2]int{mi, mx}
				for x := v; x != w; {
					x = fa[x]
					minMax[x] = [2]int{mi, mx}
				}
			}
		}
		vis[v] = 2
	}
	// 连通图的话直接写 f(0)
	for i, b := range vis {
		if b == 0 {
			f(i)
		}
	}
	return minMax
}

// 另一种写法，基于栈
func (*graph) cactusDFS2(g [][]int, n int) [][2]int {
	minMax := make([][2]int, n)
	for i := range minMax {
		minMax[i] = [2]int{-1, -1}
	}
	vis := make([]int8, n)
	s := []int{}
	var f func(v, fa int)
	f = func(v, fa int) {
		vis[v] = 1
		s = append(s, v)
		for _, w := range g[v] {
			if vis[w] == 0 {
				f(w, v)
			} else if w != fa && vis[w] == 1 {
				mi, mx := w, w
				for i := len(s) - 1; s[i] != w; i-- {
					x := s[i]
					if x < mi {
						mi = x
					} else if x > mx {
						mx = x
					}
				}
				minMax[w] = [2]int{mi, mx}
				for i := len(s) - 1; s[i] != w; i-- {
					x := s[i]
					minMax[x] = [2]int{mi, mx}
				}
			}
		}
		s = s[:len(s)-1]
		vis[v] = 2
	}
	// 连通图的话直接写 f(0, -1)
	for i, b := range vis {
		if b == 0 {
			f(i, -1)
		}
	}
	return minMax
}

/*
misc

https://codeforces.com/problemset/problem/1714/E
0 0
5 0

1 2
2 4
3 6
4 8
6 12
7 14
8 16
9 18

11 12
12 14
13 16
14 18
16 2
17 4
18 6
19 8

变形：还有 k 次 +1 操作
*/
