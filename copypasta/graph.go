package copypasta

import (
	"container/heap"
	"fmt"
	"io"
	"math"
	"math/bits"
	"slices"
	"sort"
)

/*
本页面的力扣题目已整理至【题单】图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）
https://leetcode.cn/circle/discuss/01LUak/

有关【网格图】的题目，见 search.go

建图
https://codeforces.com/problemset/problem/2027/C 1500
https://codeforces.com/problemset/problem/765/D 1700
https://codeforces.com/problemset/problem/1635/E 2200

题单 https://www.luogu.com.cn/training/81272#problems
Graph Theory Playlist https://www.youtube.com/playlist?list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P
图论的小技巧以及扩展 https://www.luogu.com.cn/blog/chengni5673/tu-lun-di-xiao-ji-qiao-yi-ji-kuo-zhan

边权转点权：在 v-w 之间加一个点，这个点的点权就是原来的边权（原图的点的点权视作 0）
点权转边权：将一个点拆分成两个点，用一条边连起来，新边的边权就是该点的点权（原图的边的边权视作 0）
其它情况：也可以用 min/max 等价转换 https://codeforces.com/problemset/problem/915/F

数学归纳法
https://codeforces.com/problemset/problem/1515/F 2600

平面图最小割转最短路
https://www.luogu.com.cn/problem/P4001
https://www.luogu.com.cn/problem/P7916
https://codeforces.com/contest/1749/problem/E 2400
https://codeforces.com/gym/104821/problem/E

TIPS: 使用一个 fa 数组（初始化为 -1）记录搜索树中的节点的父节点，这样对每个节点都有一条到根的路径（根的 fa 为 -1）
NOTE: 独立集相关问题，可以从染色的角度考虑
NOTE: 度数大于 √M 的点不超过 2√M 个
      相关题目 & 无向图定向 https://leetcode.cn/problems/minimum-degree-of-a-connected-trio-in-a-graph/solution/gei-wu-xiang-tu-ding-xiang-by-lucifer100-c72d/

https://oeis.org/A031878 Maximal number of edges in Hamiltonian path in complete graph on n nodes
a(n) = C(n, 2)        n%2==0
a(n) = C(n, 2)-n/2+1  n%2==1

https://codeforces.com/problemset/problem/1364/D 环与独立集
https://codeforces.com/problemset/problem/1198/C 匹配与独立集

https://codeforces.com/problemset/problem/41/E 1900 构造
https://codeforces.com/problemset/problem/412/D 2000 归纳
https://codeforces.com/problemset/problem/788/B 2100 转换
https://codeforces.com/problemset/problem/723/E 2200 加边
https://codeforces.com/problemset/problem/1196/F 2200 第k小路径
https://codeforces.com/problemset/problem/788/C 2300 转换
https://codeforces.com/problemset/problem/632/F 2400 转换
给一无向图，从中删除恰好一条边，求可以让图变成二分图的所有边的下标 https://codeforces.com/problemset/problem/19/E 2900
倒水问题 https://www.luogu.com.cn/problem/P1432
顶点有限制的生成树 https://codeforces.com/problemset/problem/723/F
辅助证明 https://codeforces.com/contest/1839/problem/E
https://ac.nowcoder.com/acm/contest/68572/H
https://www.luogu.com.cn/problem/P10247
- https://www.luogu.com.cn/contest/157761

集合哈希 set hashing https://codeforces.com/problemset/problem/154/C

Trémaux tree https://en.wikipedia.org/wiki/Tr%C3%A9maux_tree
DFS 树与 BFS 树 https://atcoder.jp/contests/abc251/tasks/abc251_f
证明 https://atcoder.jp/contests/abc251/editorial/3987

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

// 建图：邻接表写法
// g[v] 表示 v 的邻居
func (*graph) readGraph(in io.Reader, n, m int) {
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w, wt int
		fmt.Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}
}

// 建图：链表写法（链式前向星）
// 节点 v 的邻居形成一条链表（用数组下标代替指针）
// 添加 v 的一个邻居 w 时，用链表的【头插法】把 w 插入到链表的头节点之前
func (*graph) readGraphList(in io.Reader, n, m int) {
	type node struct{ to, next int }
	head := make([]int, n) // head[i] 表示 i 的邻居链表的头节点
	for i := range head {
		head[i] = -1 // -1 表示 nil
	}
	nodes := make([]node, m) // 无向图是 m*2
	for i := 0; i < m; i++ {
		var v, w int
		fmt.Fscan(in, &v, &w)
		v--
		w--
		nodes[i] = node{w, head[v]} // 头插法
		head[v] = i
	}

	// 遍历 v 的所有邻居 w
	// 和遍历链表是一样的
	var v int
	for cur := head[v]; cur != -1; cur = nodes[cur].next {
		w := nodes[cur].to
		_ = w // do(w) ...
	}
}

// XOR Linked Tree 
// https://codeforces.com/blog/entry/135239

/* 图上的 DFS
https://atcoder.jp/contests/abc126/tasks/abc126_e
https://atcoder.jp/contests/agc012/tasks/agc012_b 思维扩展
https://codeforces.com/problemset/problem/744/A 1500
https://codeforces.com/problemset/problem/659/E 1600
https://codeforces.com/problemset/problem/1176/E 1700
https://codeforces.com/contest/1927/problem/F 1900 找包含指定边的环
https://codeforces.com/contest/1547/problem/G 2100 对每个点 i，判断 1 到 i 有多少条路径（0/1/>=2/∞）
https://codeforces.com/problemset/problem/1470/D 2200 先染色，再递归
https://codeforces.com/problemset/problem/1707/C 2400 与 MST 结合
https://codeforces.com/problemset/problem/1519/E 2700 无向图后向边定向
https://atcoder.jp/contests/arc111/tasks/arc111_b
*/
func (*graph) dfs(g [][]int, st int) {
	// 代码来自 https://codeforces.com/problemset/problem/659/E
	vis := make([]bool, len(g))
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

	// 返回一个以 start 为起点和终点的简单环
	// ！必须保证 start 在一个环中
	// 例如 cycle=[1,2,3] 表示一个 1->2->3->1 的环
	// https://codeforces.com/contest/1927/problem/F 先用 tarjan 去掉所有割边
	cycleAt := func(start int) []int {
		vis := make([]bool, len(g))
		cycle := []int{}
		var dfs func(int, int) bool
		dfs = func(v, fa int) bool {
			vis[v] = true
			cycle = append(cycle, v) // v+1
			for _, w := range g[v] {
				if w != fa && (w == start || !vis[w] && dfs(w, v)) {
					return true
				}
			}
			cycle = cycle[:len(cycle)-1]
			return false
		}
		dfs(start, -1)
		return cycle
	}
	_ = cycleAt

	{
		// 奇偶标记法
		// https://codeforces.com/problemset/problem/936/B
		vis := make([][2]bool, len(g))
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
		vis := make([]bool, len(g))
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
		// LC802 https://leetcode.cn/problems/find-eventual-safe-states/
		// https://codeforces.com/problemset/problem/25/D
		// https://codeforces.com/problemset/problem/698/B
		// https://codeforces.com/problemset/problem/936/B
		// https://codeforces.com/problemset/problem/1217/D 给一个有向图着色，使得没有一个环只有一个颜色，求使用的颜色数量的最小值
		// https://codeforces.com/problemset/problem/1547/G 2100
		// 与 AC 自动机结合 https://www.luogu.com.cn/problem/P2444
		color := make([]int8, len(g))
		var dfs func(int) bool
		dfs = func(v int) bool {
			color[v] = 1
			for _, w := range g[v] {
				c := color[w]
				if c == 0 { // 未访问过，即 DFS 树上的树边【树枝边】
					if dfs(w) {
						return true
					}
				} else if c == 1 { // 后向边，说明有环
					return true
				} // else: 前向边或横向边，说明有多条路径可以到 w
			}
			color[v] = 2
			return false
		}
		for i, c := range color {
			if c == 0 {
				dfs(i) // ...
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
				if w == fa {
					continue
				}
				if w == v {
					// 自环
					c = 1
				} else if vis[w] { // 返祖边 v-anc 或者 anc-v（反向返祖边），例如 1-2-3-1，其中 3-1 和 1-3 都是 vis 状态
					// 一般环
					c = 2
				} else { // 树枝边
					f(w, v)
				}
			}
		}
		_ = c
		f(0, -1)
	}

	{
		// 无向图找长度 >= k 的环
		// 注：如果只有一个环（基环树），见 pseudotree
		// 模板题 https://codeforces.com/problemset/problem/263/D
		// https://codeforces.com/problemset/problem/1325/F
		var k, end, begin int
		fa := make([]int, len(g))
		dep := make([]int, len(g))
		var f func(int, int, int) bool
		f = func(v, p, d int) bool {
			fa[v] = p
			dep[v] = d
			for _, w := range g[v] {
				if dep[w] == 0 {
					if f(w, v, d+1) {
						return true
					}
				} else if d-dep[w] >= k {
					end, begin = v, w
					return true
				}
			}
			return false
		}
		f(0, -1, 1)

		cycle := []any{begin + 1} // for print
		for v := end; v != begin; v = fa[v] {
			cycle = append(cycle, v+1)
		}
	}

	// 基环树找环见下面的基环树

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

/* 图上的 BFS
https://atcoder.jp/contests/abc132/tasks/abc132_e 分层图
https://codeforces.com/problemset/problem/689/B 1600 最短路模板题
https://codeforces.com/problemset/problem/601/A 1600 脑筋急转弯
https://codeforces.com/problemset/problem/1721/D 1800 带撤销的 BFS
https://codeforces.com/problemset/problem/1851/F 1800 带撤销的 BFS
https://codeforces.com/problemset/problem/1272/E 1900 建模
https://codeforces.com/problemset/problem/1790/G 2300 锻炼分类讨论能力
https://codeforces.com/problemset/problem/1874/B 2400
*/
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
		vis := make([]bool, n)
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

// 字典序最小最短路 · 其一
// 路径的点权序列的字典序最小（边权为 1）
// https://codeforces.com/problemset/problem/1209/F 2600
// https://ac.nowcoder.com/acm/contest/76652/B
func (*graph) lexicographicallySmallestShortestPath(g [][]struct{ to, color int }, st, end int) []int {
	dis := make([]int, len(g))
	from := make([]int, len(g)) // 可选（输出具体路径用）
	vis := make([]bool, len(g))
	vis[st] = true
	q := [][]int{{st}}
	for len(q) > 0 {
		vs := q[0]
		q = q[1:]
		type edge struct{ from, to int }
		nxt := map[int][]edge{}
		for _, v := range vs {
			for _, e := range g[v] {
				nxt[e.color] = append(nxt[e.color], edge{v, e.to})
			}
		}

		_keys := make([]int, 0, len(nxt))
		for k := range nxt {
			_keys = append(_keys, k)
		}
		slices.Sort(_keys)

		// 优先走 color 小的边
		for _, color := range _keys {
			ws := []int{}
			for _, e := range nxt[color] {
				w := e.to
				if !vis[w] {
					vis[w] = true
					from[w] = e.from // 记录每个节点的前驱
					dis[w] = dis[e.from] ^ color
					ws = append(ws, w)
				}
			}
			if len(ws) > 0 {
				q = append(q, ws)
			}
		}
	}

	{
		// EXTRA：输出从起点到终点的字典序最小最短路
		path := []int{}
		for v := end; v != st; v = from[v] {
			path = append(path, v)
		}
		path = append(path, st)
		slices.Reverse(path)
	}

	return dis
}

// 字典序最小最短路 · 其二
// 上面（其一）的做法需要排序，如何避免排序呢？
// 入门经典第二版 p.173：理想路径（NEERC10）https://codeforces.com/gym/101309 I 题
// - 从终点倒着 BFS 求最短路，然后从起点开始一层一层向终点走，每一步都选颜色最小的，并记录最小颜色对应的所有节点，供下一层遍历
// EXTRA: 如果要求路径节点编号的字典序最小，每一步需选择符合 dis[w] == dis[v]-1 的编号最小的顶点
// LC499 https://leetcode.cn/problems/the-maze-iii/
func (*graph) lexicographicallySmallestShortestPath2(g [][]struct{ to, color int }, st, end int) []int {
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
func (*graph) shortestCycleBFS(n int, g [][]int) int {
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
// 通过图中每条边恰好一次
// 半欧拉图：具有欧拉路径而无欧拉回路的图
// 判定方法：
// 无向图欧拉回路：连通且没有奇度数点（全为偶度数点）
// 无向图欧拉路径：连通且恰有 0 或 2 个奇度数点（若有 2 个，则选择其中一奇度数点为起点）
// 有向图欧拉回路：SCC 只有一个且每个点的入度和出度相同
// 有向图欧拉路径：1. 对应的无向图是连通的；2. 若每个点的入度和出度相同则起点任意；否则起点的出度比入度多一，终点的入度比出度多一，且其余点的入度和出度相同
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
// - 有向图 https://www.luogu.com.cn/problem/P7771 
//         LC332 双向边 https://leetcode.cn/problems/reconstruct-itinerary/
// 模板题（输出边）
// - 有向图 LC2097 https://leetcode.cn/problems/valid-arrangement-of-pairs/ 2651
// LC753 https://leetcode.cn/problems/cracking-the-safe/ 2274
// https://codeforces.com/problemset/problem/1511/D 1600 构造
// https://codeforces.com/problemset/problem/723/E 2200 虚点
// https://codeforces.com/problemset/problem/209/C 2400 添加边使得图存在欧拉回路
// https://codeforces.com/problemset/problem/1186/F 2400
// https://codeforces.com/problemset/problem/1361/C 2500 转换
// https://codeforces.com/problemset/problem/527/E 2600
// https://ac.nowcoder.com/acm/contest/4010/H 构造
//
// BEST 定理
// https://en.wikipedia.org/wiki/BEST_theorem
// https://cmwqf.github.io/2020/09/12/%E6%B5%85%E8%B0%88BEST%E5%AE%9A%E7%90%86/
// https://luckyglass.github.io/2020/20Nov24thArt1/
// https://www.luogu.com.cn/problem/P5807
// https://atcoder.jp/contests/abc336/tasks/abc336_g

// 无向图欧拉回路/欧拉路径（有向图见下面）
func (*graph) eulerianPathOnUndirectedGraph(m int, g [][]struct{ to, eid int }) []int {
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
	// 注意：如果 m~1e6，建议把 dfs 写在外面，防止 MLE
	path := make([]int, 0, len(g)) // m
	vis := make([]bool, m)
	var dfs func(int)
	dfs = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			i := e.eid
			if vis[i] {
				continue
			}
			vis[i] = true
			w := e.to
			dfs(w)
			// 输出边的写法，注意是倒序
			// path = append(path, i)   [2]int{v, w}
		}
		// 输出点的写法（最后需要反转 path）
		path = append(path, v)
	}
	dfs(st) // for i := range g { dfs(i) }

	slices.Reverse(path) // 如果输出的是点
	return path
}

// 无向【完全图】欧拉回路/欧拉路径
// 如果 n 是偶数，会去掉一些边
// https://codeforces.com/contest/1981/problem/D 2400
func (*graph) eulerianPathOnUndirectedCompleteGraph(n int) []int {
	g := make([]int, n)
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	if n%2 == 0 {
		// 去掉 n/2-1 条边，保证至多两个奇度数点
		// 如果要求欧拉回路的话，i 改成从 1 开始（去掉 n/2 条边，保证所有点度数都是偶数）
		for i := 2; i < n; i += 2 {
			vis[i-1][i] = true
			vis[i][i-1] = true
		}
	}

	// 不允许自环的话，加上这个 for 循环（注意自环不影响度数的奇偶性）
	for i, r := range vis {
		r[i] = true
	}

	// 注意：如果 n~1e3，会导致递归深度 ~1e6，建议把 dfs 写在外面，防止 MLE
	path := []int{}
	var dfs func(int)
	dfs = func(v int) {
		for ; g[v] < n; g[v]++ {
			w := g[v]
			if vis[v][w] {
				continue
			}
			vis[v][w] = true
			vis[w][v] = true
			dfs(w)
		}
		path = append(path, v) // 记录节点
	}
	dfs(0)
	slices.Reverse(path)
	return path
}

// 有向图欧拉回路/欧拉路径
func (*graph) eulerianPathOnDirectedGraph(m int, g [][]struct{ to, eid int }) []int {
	// 读图的时候，统计入度 ...
	inDeg := make([]int, len(g))

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

	// 注意：如果 m~1e6，建议把 dfs 写在外面，防止 MLE
	path := make([]int, 0, m+1)
	var dfs func(int)
	dfs = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			dfs(e.to)
			// NOTE: 输出边的话移在这里 append e.eid
		}
		path = append(path, v)
	}
	dfs(st)

	slices.Reverse(path)
	return path
}

/* Topic - DFS 树
讲解+套题 https://codeforces.com/blog/entry/68138
好题：https://codeforces.com/problemset/problem/1325/F
*/

// 割点 cut vertices / articulation points
// https://codeforces.com/blog/entry/68138
// https://oi-wiki.org/graph/cut/#_1
// low(v): 在不经过 v 父亲的前提下能到达的最小的时间戳
// 模板题 https://www.luogu.com.cn/problem/P3388
func (*graph) findCutVertices(n int, g [][]int) (isCut []bool) {
	isCut = make([]bool, n)
	dfn := make([]int, n) // DFS 到结点 v 的时间（从 1 开始）
	// low[v] 定义为以下两种情况的最小值
	// 1. dfn[v]
	// 2. subtree(v) 的返祖边所指向的节点的 dfn，也就是经过恰好一条不在 DFS 树上的边，能够到达 subtree(v) 的节点的 dfn
	dfsClock := 0
	var tarjan func(int, int) int
	tarjan = func(v, fa int) int { // 无需考虑重边
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			if dfn[w] == 0 {
				childCnt++
				lowW := tarjan(w, v)
				lowV = min(lowV, lowW)
				if lowW >= dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 的祖先（可以连到 v 上，这也算割点）
					isCut[v] = true
				}
			} else if w != fa { // （w!=fa 可以省略，但为了保证某些题目没有重复统计所以保留）   找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 { // 特判：在 DFS 树上只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割点
			isCut[v] = false
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(v, -1)
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
// 模板题 LC1192 https://leetcode.cn/problems/critical-connections-in-a-network/
//       https://codeforces.com/problemset/problem/1000/E
// https://codeforces.com/contest/1927/problem/F
// 题目推荐 https://cp-algorithms.com/graph/bridge-searching.html#toc-tgt-2
// 与 MST 结合 https://codeforces.com/problemset/problem/160/D
// 与最短路结合 https://codeforces.com/problemset/problem/567/E
// https://codeforces.com/problemset/problem/118/E
// todo 构造 https://codeforces.com/problemset/problem/550/D
func (*graph) findBridges(n int, edges [][]int) (isBridge []bool) {
	type neighbor struct{ to, eid int }
	g := make([][]neighbor, n)
	for i, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], neighbor{w, i})
		g[w] = append(g[w], neighbor{v, i})
	}

	isBridge = make([]bool, len(edges))
	dfn := make([]int, len(g)) // 值从 1 开始
	dfsClock := 0
	var tarjan func(int, int) int
	tarjan = func(v, fid int) int { // 使用 fid 而不是 fa，可以兼容重边的情况
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		for _, e := range g[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := tarjan(w, e.eid)
				lowV = min(lowV, lowW)
				if lowW > dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 或 v 的祖先，所以 v-w 必定是桥
					isBridge[e.eid] = true
				}
			} else if e.eid != fid { // 找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(v, -1)
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

/*
无向图的双连通分量 Biconnected Components (BCC) / 重连通图
点双连通分量 v-BCC：任意割点都是至少两个不同点双的公共点
点双也叫 Block
缩点后形成一棵 block-cut tree / BC-tree https://en.wikipedia.org/wiki/Biconnected_component#Block-cut_tree
每条树边恰好在一个点双内
每个点双的点数就是一个极大环的点数
https://oi-wiki.org/graph/bcc/
https://www.csie.ntu.edu.tw/~hsinmu/courses/_media/dsa_13spring/horowitz_306_311_biconnected.pdf

模板题 https://www.luogu.com.cn/problem/P8435
好题 https://codeforces.com/problemset/problem/962/F
LCP54 https://leetcode.cn/problems/s5kipK/
todo 结合树链剖分 https://codeforces.com/problemset/problem/487/E

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
func (G *graph) findVertexBCC(g [][]int) (comps [][]int, bccIDs []int) {
	bccIDs = make([]int, len(g)) // ID 从 1 开始编号
	idCnt := 0
	isCut := make([]bool, len(g)) // 缩点用

	dfn := make([]int, len(g)) // 值从 1 开始
	dfsClock := 0
	type edge struct{ v, w int } // eid
	st := []edge{}               // 存边是为了解决一些特殊题目（基本写法存点就行）
	var tarjan func(int, int) int
	tarjan = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			e := edge{v, w} // ne.eid
			if dfn[w] == 0 {
				st = append(st, e)
				childCnt++
				lowW := tarjan(w, v)
				lowV = min(lowV, lowW)
				if lowW >= dfn[v] {
					isCut[v] = true
					idCnt++
					comp := []int{}
					//eids := []int{}
					for {
						e, st = st[len(st)-1], st[:len(st)-1]
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
					// 仙人掌点双：点数和边数相同，说明该 v-BCC 是一个简单环，且环上所有的边只属于一个简单环
					//if len(comp) == len(eids) {
					//	for _, eid := range eids {
					//		onSimpleCycle[eid] = true
					//	}
					//}
					comps = append(comps, comp)
				}
			} else if w != fa && dfn[w] < dfn[v] {
				st = append(st, e) // 简单写法中，可以省略
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
			tarjan(v, -1)
		}
	}

	// EXTRA: 缩点
	// BCC 和割点作为新图中的节点，并在每个割点与包含它的所有 BCC 之间连边
	cutIDs := make([]int, len(g))
	for i, ok := range isCut {
		if ok {
			idCnt++ // 接在 BCC 之后给割点编号
			cutIDs[i] = idCnt
		}
	}
	for v, cp := range comps {
		v++
		for _, w := range cp {
			if w = cutIDs[w]; w > 0 {
				// add(v, w)
				// add(w, v)
				// ...
			}
		}
	}

	return
}

// 边双连通分量 e-BCC：删除无向图中所有的割边后，剩下的每一个 CC 都是 e-BCC
// 缩点后形成一棵 bridge tree
// 模板题 https://www.luogu.com.cn/problem/P8436
// 模板题 https://codeforces.com/problemset/problem/1000/E
// 较为综合的一道题 https://codeforces.com/problemset/problem/732/F
// https://codeforces.com/gym/104270/problem/B DSU
// - https://www.luogu.com.cn/problem/P9886
func (G *graph) findEdgeBCC(n int, edges [][]int) (comps [][]int, bccIDs []int) {
	isBridge := G.findBridges(n, edges)
	type neighbor struct{ to, eid int }
	var g [][]neighbor

	// 求原图中每个点的 bccID
	bccIDs = make([]int, len(g))
	idCnt := 0
	var comp []int
	var initBid func(int)
	initBid = func(v int) {
		bccIDs[v] = idCnt
		comp = append(comp, v)
		for _, e := range g[v] {
			if w := e.to; !isBridge[e.eid] && bccIDs[w] == 0 {
				initBid(w)
			}
		}
	}
	for i, id := range bccIDs {
		if id == 0 {
			idCnt++
			comp = []int{}
			initBid(i)
			comps = append(comps, comp)
		}
	}

	// EXTRA: 缩点，复杂度 O(M)
	// 遍历 edges，若两端点的 bccIDs 不同（割边）则建边
	g2 := make([][]int, idCnt)
	for _, e := range edges {
		v, w := e[0], e[1]
		v, w = bccIDs[v]-1, bccIDs[w]-1
		if v != w {
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
		}
	}

	// 也可以遍历 isBridge，割边两端点 bccIDs 一定不同
	for i, b := range isBridge {
		if b {
			e := edges[i]
			v, w := e[0], e[1]
			v, w = bccIDs[v]-1, bccIDs[w]-1
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
		}
	}

	return
}

// 圆方树 Round-square tree
// 圆点：原图中的点
// 方点：每个点双的虚拟节点
// 动机：直接点双缩点会丢失信息，应当如何把点双和原图上的点联系起来？
// 目标：能够快速求出图中任意两点间的最短路
//
// 狭义圆方树（仙人掌图） Cactus / Cactus tree
// 一条边至多在一个简单回路中 / 任意两个简单回路至多有一个公共点
// 注意：只有两个点的点双无需建方点
// 图片请看 https://en.wikipedia.org/wiki/Cactus_graph
// A connected graph in which any two simple cycles have at most one vertex in common
// Equivalently, it is a connected graph in which every edge belongs to at most one simple cycle
// 如果图没有偶环，则不可能有两个奇环共用一条边（因为这样会构成一个偶环），因此没有两个环共用一条边，图一定为仙人掌（注意：反过来，一个仙人掌图是可能有偶环的）
//
// 广义圆方树 Round-square tree
// 一条边可以在多个简单回路中
// 注意：只有两个点的点双也要建方点
// https://oi-wiki.org/graph/block-forest/
// Block graph：每个点双都是完全图 https://en.wikipedia.org/wiki/Block_graph
//
// 静态仙人掌 https://www.luogu.com.cn/problem/P5236
// todo https://www.luogu.com.cn/problem/P4129
//  https://www.luogu.com.cn/problem/P4244
//  https://www.luogu.com.cn/problem/P3687
// todo 动态仙人掌 https://www.luogu.com.cn/problem/P5237
// todo 题单 圆方树——处理仙人掌的利器 https://immortalco.blog.uoj.ac/blog/1955
//
// todo 题单 https://www.cnblogs.com/alex-wei/p/high_level_graph_theory.html
//  https://www.luogu.com.cn/problem/P4630
//  https://www.luogu.com.cn/problem/P4606
//  https://www.luogu.com.cn/problem/P5058
//  https://www.luogu.com.cn/problem/P4334
//  https://www.luogu.com.cn/problem/P3225
//  https://www.luogu.com.cn/problem/P8456
//  https://codeforces.com/problemset/problem/1763/F 3000
//  https://codeforces.com/problemset/problem/487/E 3200
//  https://www.luogu.com.cn/problem/UVA1464
// https://atcoder.jp/contests/abc318/tasks/abc318_g 2260
func (*graph) roundSquareTree(n int, edges [][]int, abs func(int) int) {
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	g2 := make([][]nb, len(g)*2)
	ringWtSum := make([]int, len(g)*2)
	bccID := len(g)
	dis := make([]int, len(g))

	dfn := make([]int, len(g))
	dfsClock := 0
	st := []int{}
	aWt := make([]int, len(g)) // 如果栈中存边的话可以合并到栈中
	var tarjan func(int, int) int
	tarjan = func(v, fa int) int {
		st = append(st, v)
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		for _, e := range g[v] {
			w := e.to
			if dfn[w] == 0 {
				aWt[w] = e.wt
				dis[w] = dis[v] + e.wt
				lowW := tarjan(w, v)
				lowV = min(lowV, lowW)
				if lowW >= dfn[v] { // v 是割点
					// v 和方点 bccID 连边，边权为 0
					g2[v] = append(g2[v], nb{bccID, 0})
					g2[bccID] = append(g2[bccID], nb{v, 0})
					ancestor := st[len(st)-1]
					ringWtSum[bccID] = aWt[ancestor] + dis[ancestor] - dis[v] // 这个点双的环长
					// 遍历这个点双中的点（除了割点 v）
					for {
						x := st[len(st)-1]
						st = st[:len(st)-1]
						d := dis[x] - dis[v]
						d = min(d, ringWtSum[bccID]-d)
						// 点 x 和方点 bccID 连边，边权为 x 到 v 的最短距离 d
						g2[x] = append(g2[x], nb{bccID, d})
						g2[bccID] = append(g2[bccID], nb{x, d})
						if x == w {
							break
						}
					}
					bccID++
				}
			} else if w != fa && dfn[w] < dfn[v] { // 返祖边
				aWt[v] = e.wt
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	tarjan(0, -1)

	const mx = 20 // bits.Len(最大节点数*2)
	pa := make([][mx]int, bccID)
	dep := make([]int, bccID)
	dis2 := make([]int, bccID)
	var buildPa func(int, int)
	buildPa = func(v, fa int) {
		pa[v][0] = fa
		for _, e := range g2[v] {
			if w := e.to; w != fa {
				dis2[w] = dis2[v] + e.wt
				dep[w] = dep[v] + 1
				buildPa(w, v)
			}
		}
	}
	buildPa(0, -1)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	getLca := func(v, w int) (int, int) {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for k := uint(dep[w] - dep[v]); k > 0; k &= k - 1 {
			w = pa[w][bits.TrailingZeros(k)]
		}
		if w == v { // 一个点是另一个点的祖先节点（必然是圆点）
			return v, -1
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		if pa[v][0] < len(g) { // 圆点
			return pa[v][0], -1
		}
		// 此时 pa[v][0] 是方点
		return v, w // 方点下面的两个圆点
	}
	// 返回仙人掌上两点间的最短距离
	getDis := func(v, w int) int {
		x, y := getLca(v, w)
		if y < 0 { // x = lca
			return dis2[v] + dis2[w] - dis2[x]*2
		}
		// 看图理解 https://cdn.luogu.com.cn/upload/image_hosting/k0rtviky.png
		dxy := abs(dis[x] - dis[y]) // x 和 y 在同一个点双，应当使用原图的 dis 信息
		dxy = min(dxy, ringWtSum[pa[x][0]]-dxy)
		return dis2[v] + dis2[w] - dis2[x] - dis2[y] + dxy
	}

	// EX: 判断从 src 到 dst 的简单路径上是否包含点 target
	// 只要路径上有个方点与 y 相连即可
	// https://atcoder.jp/contests/abc318/tasks/abc318_g
	pathContains := func(src, dst, target int) (ok bool) {
		var g2 [][]int
		var dfs func(int, int) bool
		dfs = func(v, fa int) bool {
			if v == dst {
				return true
			}
			for _, w := range g2[v] {
				if w != fa && dfs(w, v) {
					if !ok && w >= n {
						for _, x := range g2[w] {
							if x == target {
								ok = true
								break
							}
						}
					}
					return true
				}
			}
			return false
		}
		dfs(src, -1)
		return
	}

	_ = []any{getDis, pathContains}
}

//

type dijkstraPair struct{ v, dis int }
type dijkstraHeap []dijkstraPair

func (h dijkstraHeap) Len() int             { return len(h) }
func (h dijkstraHeap) Less(i, j int) bool   { return h[i].dis < h[j].dis }
func (h dijkstraHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *dijkstraHeap) Push(v any)          { *h = append(*h, v.(dijkstraPair)) }
func (h *dijkstraHeap) Pop() (v any)        { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *dijkstraHeap) push(v dijkstraPair) { heap.Push(h, v) }
func (h *dijkstraHeap) pop() dijkstraPair   { return heap.Pop(h).(dijkstraPair) }

// 单源最短路 Dijkstra
// 适用于稀疏图 O(mlogm)
// 这里实现的是懒更新（lazy）版本的 Dijkstra，复杂度为 O(mlogm)
// 若在插入堆时元素已在堆中，改成直接更新元素，而不是插入元素，可使复杂度降为 O(mlogn)
//
// st 也可以是一个点集，这相当于同时对多个点跑最短路
// 视频讲解（第四题）https://www.bilibili.com/video/BV1wj411G7sH/
// 可视化 https://visualgo.net/zh/sssp
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 最短路问题笔记 https://www.luogu.com.cn/blog/SCN/zui-duan-lu-wen-ti-bi-ji
// Universal Optimality https://mp.weixin.qq.com/s/V0MlmsoZdcRIZcxsRgAzuA
//
// 模板题 https://www.luogu.com.cn/problem/P3371 https://www.luogu.com.cn/problem/P4779
//       https://codeforces.com/problemset/problem/20/C 1900
// 最短路个数 https://www.luogu.com.cn/problem/P1144
//           https://www.luogu.com.cn/problem/P1608
// 【理解本质】https://atcoder.jp/contests/abc271/tasks/abc271_e
// https://atcoder.jp/contests/abc191/tasks/abc191_e
// https://atcoder.jp/contests/abc192/tasks/abc192_e
// https://codeforces.com/problemset/problem/229/B 1700 二分
// https://codeforces.com/problemset/problem/2023/B 1700 建模
// https://codeforces.com/problemset/problem/1725/M 1800 分层图最短路
// https://codeforces.com/problemset/problem/1915/G 1800 双关键字
// https://codeforces.com/problemset/problem/1842/D 1900 转换
// https://codeforces.com/problemset/problem/449/B 2000 【理解本质】通过最短路找到可以删除的边
// https://codeforces.com/problemset/problem/545/E 2000 MST
// - https://atcoder.jp/contests/arc090/tasks/arc090_c
// https://codeforces.com/problemset/problem/938/D 2000 超级源点
// - LC2473 https://leetcode.cn/problems/minimum-cost-to-buy-apples/
// https://atcoder.jp/contests/abc204/tasks/abc204_e 1710=CF2004
// https://codeforces.com/problemset/problem/507/E 2100 双关键字+记录路径编号 
// https://codeforces.com/problemset/problem/1005/F 2100
// https://codeforces.com/problemset/problem/567/E 2200 * 关键边、伪关键边（与割边结合）
// https://codeforces.com/problemset/problem/715/B 2300 https://leetcode.cn/problems/modify-graph-edge-weights/
// https://codeforces.com/problemset/problem/1693/C 2300 转换
// https://codeforces.com/problemset/problem/1753/D 2400 建模+转换+多源最短路 
// https://atcoder.jp/contests/abc245/tasks/abc245_g 2270=CF2428
// https://codeforces.com/problemset/problem/1528/D 2500 建模【好题】
// https://atcoder.jp/contests/arc064/tasks/arc064_c * 稠密图
// https://atcoder.jp/contests/abc237/tasks/abc237_e 转换
// LC864 https://leetcode.cn/problems/shortest-path-to-get-all-keys/ 2259 建模 
// https://www.luogu.com.cn/problem/P4644 建模 
// [SDOI2010]大陆争霸 https://www.luogu.com.cn/problem/P2446
// [AHOI2014/JSOI2014]骑士游戏 https://www.luogu.com.cn/problem/P4042
// 题目推荐 https://cp-algorithms.com/graph/dijkstra.html#toc-tgt-5
// 线段树优化建图 https://codeforces.com/problemset/problem/786/B 2300
// - todo [SNOI2017] 炸弹 https://www.luogu.com.cn/problem/P5025
// 涉及到相邻两条边的最短路 https://codeforces.com/contest/1486/problem/E 2200
// todo 与扩欧结合 https://www.acwing.com/problem/content/3418/
// 分层图最短路
// - 空间压缩 https://codeforces.com/problemset/problem/1442/C 2400
// - 转换 https://codeforces.com/problemset/problem/1473/E 2400
// todo 动态最短路 https://codeforces.com/problemset/problem/1163/F 3000
// todo 与类似并查集的东西结合 https://ac.nowcoder.com/acm/problem/19789
//
// 最短路树
// todo https://xyzl.blog.luogu.org/Shortest-Path-Tree-SPT
// 最短路树上跑拓扑排序 LC1786 https://leetcode.cn/problems/number-of-restricted-paths-from-first-to-last-node/ 2079
// 最短路树上跑拓扑排序 https://codeforces.com/contest/1076/problem/D 1800
func (*graph) shortestPathDijkstra(n, st int, edges [][]int) (dis []int) {
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	const inf int = 1e18
	dis = make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0            // 如果不写这一行，那么计算出的 dis[st] 是包含 st 的最小环长度
	from := make([]int, n) // 见下面「从 st 到 end 的路径」
	for i := range from {
		from[i] = -1
	}
	h := dijkstraHeap{{st, 0}}
	for len(h) > 0 {
		p := h.pop()
		v := p.v
		// 下面循环中的 newD < dist[w] 可能会把重复的节点 w 入堆
		// 也就是说，堆中可能会包含多个相同节点，且这些相同节点的 dist 值【互不相同】
		// 那么这个节点第二次及其后面出堆的时候，由于 dist[v] 已经更新成最短路了，其邻居的最短路也已经被更新过了，所以直接跳过
		if p.dis > dis[v] {
			continue
		}
		// 第 k 个执行到这一行的最短路 dist[v] 就是所有 dist 中第 k 小的最短路
		// 用这一思想解决 https://ac.nowcoder.com/acm/contest/65157/C
		for _, e := range g[v] {
			w := e.to
			newD := p.dis + e.wt
			if newD < dis[w] {
				dis[w] = newD
				from[w] = v
				h.push(dijkstraPair{w, newD})
			}
		}
	}

	// 判断是否可以到达终点
	//if dist[end] == inf { return -1 }

	// EXTRA: 从 st 到 end 的任意一条路径
	// 记录边的编号 https://codeforces.com/problemset/problem/507/E
	getAnyPath := func(end int) (path []int) {
		for x := end; x != -1; x = from[x] {
			path = append(path, x)
		}
		slices.Reverse(path)
		return
	}
	_ = getAnyPath

	// EXTRA: 哪些边在从 st 到 end 的最短路上？
	//
	// 第一种方法：
	// 从 end 出发，算一次最短路 dist2，那么满足
	//     dist[v] + wt + dist2[w] == dist[end] || dist[w] + wt + dist2[v] == dist[end]
	// 的边就在最短路上
	//
	// 第二种方法：
	// 从 end 出发 DFS/BFS，满足 dist[v] + wt == dist[w] 的边就在最短路上
	// 详见 https://leetcode.cn/problems/find-edges-in-shortest-paths/solution/dijkstra-zui-duan-lu-dfsbfs-zhao-bian-py-yf48/
	{
		vis := make([]bool, len(g))
		var rdfs func(int)
		rdfs = func(w int) {
			vis[w] = true
			for _, e := range g[w] {
				v := e.to
				if dis[v]+e.wt != dis[w] {
					continue
				}
				// 标记在最短路上的边（建图的时候把边的下标也存进去）
				//onPath[e.i] = true
				if !vis[v] {
					rdfs(v)
				}
			}
		}
		var end int
		rdfs(end)
	}

	// EXTRA: 构建有向最短路树（除根节点外的点的入度均为 1）
	// https://atcoder.jp/contests/abc252/tasks/abc252_e
	{
		g2 := make([][]neighbor, n)
		vis := make([]bool, n)
		for v, es := range g {
			for _, e := range es {
				w := e.to
				if vis[w] { // 已经有一条指向 w 的边了。若去掉该判断则构建的是 DAG
					continue
				}
				if dis[v]+e.wt == dis[w] {
					g2[v] = append(g2[v], e)
					vis[w] = true
				}
			}
		}
	}

	// EXTRA: 在最短路 DAG 上跑拓扑（如最短路计数）
	// LC1976 https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/
	// 最短路个数 https://www.luogu.com.cn/problem/P1144 https://www.luogu.com.cn/problem/P1608 
	// 也可以把转移写在求最短路的代码中，见 https://www.luogu.com.cn/record/56683589
	// 紧急情况 https://www.acwing.com/problem/content/1477/
	// 条条大路通罗马 https://www.acwing.com/problem/content/1579/
	{
		deg := make([]int, n)
		for v, es := range g {
			for _, e := range es {
				if dis[v]+e.wt == dis[e.to] {
					deg[e.to]++
				}
			}
		}

		f := make([]int, n)
		f[st] = 0 //
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; dis[v]+e.wt == dis[w] {
					// do f[v] -> f[w] ...

					if deg[w]--; deg[w] == 0 {
						q = append(q, w)
					}
				}
			}
		}
	}

	// EXTRA: 次短路·其一
	// 模板题 https://www.luogu.com.cn/problem/P2865
	// LC2045 https://leetcode.cn/problems/second-minimum-time-to-reach-destination/ 2202
	// 次短路计数 https://codeforces.com/contest/1650/problem/G http://poj.org/problem?id=3463 https://www.acwing.com/problem/content/385/
	// 长度不超过最短路长度+K 的路径个数 [NOIP2017 提高组] 逛公园 https://www.luogu.com.cn/problem/P3953
	{
		const inf int = 1e18
		dist := make([]int, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		dist2 := make([]int, n)
		for i := range dist2 {
			dist2[i] = inf
		}
		h := dijkstraHeap{{st, 0}}
		for len(h) > 0 {
			p := h.pop()
			v, d := p.v, p.dis
			if d > dist2[v] { // 注意是 dist2
				continue
			}
			for _, e := range g[v] {
				w := e.to
				newD := d + e.wt
				if newD < dist[w] {
					dist2[w] = dist[w]
					dist[w] = newD
					h.push(dijkstraPair{w, newD})
				} else if dis[w] < newD && newD < dist2[w] { // 求严格次短路需要写 dis[w] < newD
					dist2[w] = newD
					h.push(dijkstraPair{w, newD})
				}
			}
		}
	}

	// EXTRA: 次短路·其二
	// 另一种次短路写法，适用性更广
	// https://atcoder.jp/contests/abc245/tasks/abc245_g 2270=CF2428
	{
		const inf int = 1e18
		type pair struct{ d, d2 int }
		dis := make([]pair, n)
		for i := range dis {
			dis[i].d = inf
			dis[i].d2 = inf
		}
		h := dijkstraHeap{{st, 0}}
		for len(h) > 0 {
			p := h.pop()
			v, d := p.v, p.dis
			if dis[v].d == inf {
				dis[v].d = d
			} else if dis[v].d2 == inf && d > dis[v].d { // 如果不要求严格次短路，可以去掉 d > dis[v].d
				dis[v].d2 = d
			} else {
				continue
			}
			for _, e := range g[v] {
				h.push(dijkstraPair{e.to, d + e.wt})
			}
		}
	}

	return
}

// 另一种 Dijkstra 写法
// 适用于稠密图 O(n^2)
// LC2662 https://leetcode.cn/problems/minimum-cost-of-a-path-with-special-roads/ 2154
// 建模 https://codeforces.com/contest/1528/problem/D
func (*graph) shortestPathDijkstra2(g [][]int, st int) []int {
	n := len(g)
	const inf int = 1e18
	dis := make([]int, n+1) // dis[n] 当作哨兵
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	vis := make([]bool, n)
	for {
		v := n
		for w, b := range vis {
			if !b && dis[w] < dis[v] {
				v = w
			}
		}
		if dis[v] == inf {
			return dis
		}
		vis[v] = true
		for w, wt := range g[v] {
			dis[w] = min(dis[w], dis[v]+wt)
		}
	}
}

// 0-1 最短路 / 0-1 BFS
// https://oi-wiki.org/graph/bfs/#%E5%8F%8C%E7%AB%AF%E9%98%9F%E5%88%97-bfs
// https://codeforces.com/blog/entry/22276
// EXTRA: 1-2 最短路 https://codeforces.com/blog/entry/90917
//
// - [1824. 最少侧跳次数](https://leetcode.cn/problems/minimum-sideway-jumps/) 1778
// - [1368. 使网格图至少有一条有效路径的最小代价](https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/) 2069
// - [2290. 到达角落需要移除障碍物的最小数目](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/) 2138
// https://codeforces.com/problemset/problem/173/B 1800
// https://codeforces.com/problemset/problem/1063/B 1800
// https://codeforces.com/problemset/problem/1941/G 2000 也可以 BFS（需要合适的建图方式）
// https://codeforces.com/problemset/problem/877/D 2100 也可以 BFS
// https://codeforces.com/problemset/problem/590/C 2200 网格图
// https://codeforces.com/problemset/problem/821/D 2200 建图技巧
// https://codeforces.com/problemset/problem/1340/C 2400 建图技巧
// https://codeforces.com/problemset/problem/1442/C 2400
// https://codeforces.com/problemset/problem/1749/E 2400 建图技巧
// 建图技巧 https://www.lanqiao.cn/problems/6281/learning/?contest_id=146
// 哪里有 1 https://atcoder.jp/contests/abc213/tasks/abc213_e
//         https://atcoder.jp/contests/abc176/tasks/abc176_d
func (*graph) bfs01(g [][]struct{ to, wt int }, st int) []int {
	const inf int = 1e18
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	type vd struct{ v, d int }
	ql, qr := []vd{{st, dis[st]}}, []vd{}
	for len(ql) > 0 || len(qr) > 0 {
		var p vd
		if len(ql) > 0 {
			ql, p = ql[:len(ql)-1], ql[len(ql)-1]
		} else {
			p, qr = qr[0], qr[1:]
		}
		v := p.v
		if p.d > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w, wt := e.to, e.wt
			newD := p.d + wt
			if newD < dis[w] {
				dis[w] = newD
				if wt == 0 {
					ql = append(ql, vd{w, newD})
				} else {
					qr = append(qr, vd{w, newD})
				}
			}
		}
	}
	return dis
}

// 单源最短路 SPFA O(nm)   队列优化的 Bellman-Ford
// 对于构建一个让 SPFA（或者纯队列写法）跑到最坏情况的（网格）图，见 main/testutil/rand.go 中的 GraphHackSPFA
// 可视化 https://visualgo.net/zh/sssp
// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
//
// 模板题 https://www.luogu.com.cn/problem/P3385
// 用到 Bellman-Ford 的思想
// - LC787 https://leetcode.cn/problems/cheapest-flights-within-k-stops/
// - LC1928 https://leetcode.cn/problems/minimum-cost-to-reach-destination-in-time/
// 边权递增最短路 http://acm.hdu.edu.cn/showproblem.php?pid=4479
//
// EXTRA: 差分约束系统
// 若有 Xi-Xj<=Ck，则连一条有向边 j->i，边权为 Ck。这样求最短路可以保证 Xi-Xj<=Ck
// 然后再添加一个 0 号节点，向其他节点连一条边权为 0 的有向边，表示 Xi-X0<=0
// 这样，在无负环时会得到一组非正数解（先求一组解，然后再全部加上 Δ）
// 模板题 https://www.luogu.com.cn/problem/P4878
// 第 i 个区间至少选 limit[i] 个 LC2589 https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/
// - 【思路讲解】O(nlogn) 贪心+单调栈二分 https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/solution/tan-xin-pythonjavacgo-by-endlesscheng-w3k3/
// - 加强版 LCP32 https://leetcode.cn/problems/t3fKg1/
// - https://www.luogu.com.cn/problem/P1250
// - https://www.luogu.com.cn/problem/SP116
// - http://poj.org/problem?id=1201
// - todo 打印方案 https://atcoder.jp/contests/abc216/tasks/abc216_g
func (*graph) shortestPathSPFA(n, st int, edges [][]int) (dis []int) { // 有负环时返回 nil
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	const inf int = 1e18 // 1e9+1
	dis = make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	q := []int{st}
	inQ := make([]bool, n)
	inQ[st] = true
	relaxedCnt := make([]int, n) // 检测负环用
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		inQ[v] = false
		for _, e := range g[v] {
			w := e.to
			newD := dis[v] + e.wt
			if newD < dis[w] {
				dis[w] = newD
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
	// 注意循环完了 inQ[i] 都是 false

	// EXTRA: 只是找负环的话，初始时将所有点入队即可
	// 注意不能只从一个点出发找负环，因为可能这个点无法到达负环
	// https://www.luogu.com.cn/problem/P2868

	// EXTRA: 打印负环
	// https://cp-algorithms.com/graph/finding-negative-cycle-in-graph.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BellmanFordSP.java.html

	return
}

// 任意两点最短路 Floyd-Warshall  O(n^3)  本质是求 Min-plus matrix multiplication
// 【图解】带你发明 Floyd 算法！https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// https://en.wikipedia.org/wiki/Min-plus_matrix_multiplication
// https://cp-algorithms.com/graph/all-pair-shortest-path-floyd-warshall.html#toc-tgt-5
//
// Solving the all pairs shortest path problem after minor update of a large dense graph
// https://arxiv.org/abs/2412.15122
//
// 模板题 https://www.luogu.com.cn/problem/B3647
// 传递闭包 https://www.luogu.com.cn/problem/B3611
// 传递闭包+矩阵快速幂 https://codeforces.com/contest/691/problem/E 1900
// https://codeforces.com/problemset/problem/1204/C
// LC1334 https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/ 1855
// LC1462 https://leetcode.cn/problems/course-schedule-iv/ 1693
// LC2976 https://leetcode.cn/problems/minimum-cost-to-convert-string-i/
// - 原题 https://codeforces.com/problemset/problem/33/B 1800
// LC2959 https://leetcode.cn/problems/number-of-possible-sets-of-closing-branches/
// 动态加点 https://codeforces.com/problemset/problem/295/B
// 动态加边 LC2642 https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/ 1811
// - https://codeforces.com/problemset/problem/25/C
// https://atcoder.jp/contests/abc143/tasks/abc143_e 最少加油次数 跑两次 Floyd
// DP LC2977 https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/
// https://atcoder.jp/contests/abc243/tasks/abc243_e 寻找不影响最短路的边
// 传递闭包 UVa247 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=4&page=show_problem&problem=183
// 注：求传递闭包时，若 i-k 不连通，则最内层循环无需运行
// 任意两点最大边权最小路径 UVa10048 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=12&page=show_problem&problem=989
// 利用最短路简化分类讨论 https://codeforces.com/gym/105139/problem/L
func (*graph) shortestPathFloydWarshall(n int, edges [][]int) [][]int {
	// g[k][i][j] 表示「经过若干个编号不超过 k 的中间节点」时，从 i 到 j 的最短路长度，其中第一维可以压缩掉
	// 为什么可以把第一维度去掉？g[i][k] 和 g[k][j] 不会被覆盖掉吗？
	// 见算法导论第三版练习 25.2-4（网络上有习题解答）

	// 初始化，保证 g[i][i] = 0
	// 注：如果没有动态加边的需求，其实 g[i][i] 是否初始化成 0 无所谓
	// 但还是需要注意，代码中是否有判断 g[i][i] 的值的逻辑
	// 【技巧】用 Floyd 计算完 g[i][j] 之后，如果要判断所有 g[i][j]，可以利用 g[i][j] = g[j][i] 的性质，只判断 j < i 的情况，从而减少一半的计算量

	const inf = math.MaxInt / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i {
				g[i][j] = inf
			}
		}
	}
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v][w] = min(g[v][w], wt)
		g[w][v] = min(g[w][v], wt)
	}
	for k := range g {
		for i := range g {
			if g[i][k] == inf {
				continue
			}
			for j := range g {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}

	// 如果出现 g[i][i] < 0 则说明有负环

	// 动态加边
	// https://codeforces.com/problemset/problem/25/C
	// https://atcoder.jp/contests/abc375/tasks/abc375_f
	// LC2642 https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/
	// LC2959 https://leetcode.cn/problems/number-of-possible-sets-of-closing-branches/ 结合状压 DP
	for i := range g {
		// 注意 from=i 或者 to=j 时，下面的 g[i][from] 和 g[to][j] 都需要 g[i][i] 这样的值
		// 所以初始化成 0 方便计算
		g[i][i] = 0
	}
	addEdge := func(from, to, wt int) {
		// 无法让任何最短路变短
		if wt >= g[from][to] {
			return
		}
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][from]+wt+g[to][j]) // g[i][to]+wt+g[from][j] 如果是双向边的话额外 min 这种情况
			}
		}
	}
	_ = addEdge

	return g
}

// 位压缩版 Floyd
// 时间复杂度 O(n^3/w), w=bits.UintSize，一般是 64
// LC2101 https://leetcode.cn/problems/detonate-the-maximum-bombs/ 1880
// https://atcoder.jp/contests/abc287/tasks/abc287_h
func (*graph) floydWarshallBitset(n int, edges [][]int) []int {
	f := make([]Bitset, n) // f[i] 表示从 i 出发可以到达的节点集合
	for i := range f {
		f[i] = NewBitset(n)
		f[i].Set(i) // i 可以到它自己
	}
	for _, e := range edges {
		v, w := e[0], e[1] // 有向边 v->w
		f[v].Set(w)
	}
	for k := range f {
		for i := range f {
			if f[i].Has(k) { // i 可以到 k
				f[i].Or(f[k]) // k 能到的点，i 也可以到           i->j 现在可以 i->k->j
			}
		}
	}
	reach := make([]int, n) // reach[i] 表示从 i 出发可以到达的节点数（注意读题，一般都要包括自己）
	for i, s := range f {
		reach[i] = s.OnesCount()
	}
	return reach
}

// 子集 Floyd
// LC2959 https://leetcode.cn/problems/number-of-possible-sets-of-closing-branches/
func (*graph) floydWarshallSubset(n int, edges [][]int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2
		}
	}
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	f := make([][][]int, 1<<n)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
			for k := range f[i][j] {
				f[i][j][k] = math.MaxInt / 2
			}
		}
	}

	f[0] = g
	// check f[0][i][j] ...

	for s := uint(1); s < 1<<n; s++ {
		k := bits.TrailingZeros(s) // 中转节点（选任一 s 中的元素）
		t := s ^ 1<<k
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[s][i][j] = min(f[t][i][j], f[t][i][k]+f[t][k][j])
				if j < i {
					// 因为 f[s][i][j] = f[s][j][i]，只需要检查一半
					// check f[s][i][j] ...
				}
			}
		}
	}
}

// 最小环
//
// 另见 Dijkstra 删边：
//（无向图）对于每条边 v-w，计算从 v 出发，在不经过边 v-w 时，到 w 的最短路，再加上 v-w 的边权
//（有向图）v->w，计算从 w 到 v 的最短路，再加上 v->w 的边权
//
// 传入邻接矩阵 weights
// weights[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/min-circle/#floyd
// NOTE: 无权图的情况见 shortestCycleBFS
func (*graph) shortestCycleFloydWarshall(weights [][]int) int {
	const inf int = 1e18
	n := len(weights)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = slices.Clone(weights[i])
	}
	ans := inf
	for k := range dist {
		for i := range k {
			for j := range i {
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
func (G *graph) shortestPathJohnson(n int, edges [][]int) [][]int {
	const inf int = 1e18
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n+1)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2] // 输入的节点编号从 1 开始	
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}

	// 建虚拟节点 0 并且往其他的点都连一条边权为 0 的边
	const src = 0
	for v := 1; v <= n; v++ {
		g[src] = append(g[src], neighbor{v, 0})
		g[v] = append(g[v], neighbor{src, 0})
	}

	// 跑 SPFA（省略建图，注意点数为 n+1）
	h := G.shortestPathSPFA(n+1, 0, edges)
	if h == nil {
		return nil
	}

	// 求新的边权
	for v := 1; v <= n; v++ {
		for i, e := range g[v] {
			g[v][i].wt += h[v] - h[e.to]
		}
	}

	// 以每个点为源点跑一遍 Dijkstra（省略建图，注意点数为 n+1）
	dist := make([][]int, n+1)
	for st := 1; st <= n; st++ {
		dist[st] = G.shortestPathDijkstra(n+1, st, edges)
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
func (*graph) shortestPathMod(a []int, limit int) (ans int) {
	slices.Sort(a) // 常数优化
	dis := make([]int, a[0])
	for i := range dis {
		dis[i] = math.MaxInt
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
			if newD := dis[v] + ai; newD < dis[w] {
				dis[w] = newD
				h.push(dijkstraPair{w, newD})
			}
		}
	}
	// dis[i] 表示通过累加除了 a[0] 以外的数，可以得到的最小的数，且满足 dis[i] ≡ i (mod a[0])
	// 如果 dis[i] != inf 说明可以得到 dis[i] + k*a[0]
	// 如果 dis[target%a[0]] <= target 说明可以得到 target
	for _, d := range dis {
		if d <= limit {
			ans += (limit-d)/a[0] + 1
		}
	}
	return
}

// k 短路
// A* 算法   astar
// 可持久化可并堆优化
// https://en.wikipedia.org/wiki/K_shortest_path_routing
// https://www.youtube.com/watch?v=CgW0HPHqFE8
// todo https://oi-wiki.org/graph/kth-path/
// todo 模板题 https://www.luogu.com.cn/problem/P2483

// 最小斯坦纳树 minimum Steiner tree problem
// 输入：一个带权无向图，k 个关键节点
// 返回：连通 k 个关键节点的最小代价
// https://en.wikipedia.org/wiki/Steiner_tree_problem
// https://oi-wiki.org/graph/steiner-tree/
// 模板题 https://www.luogu.com.cn/problem/P6192
// todo 输出方案 WC08 游览计划 https://www.luogu.com.cn/problem/P4294
func (*graph) minimumSteinerTree(n int, edges [][]int, points []int) int {
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	const inf int = 1e9
	k := len(points)
	f := make([][]int, 1<<k)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	for i, v := range points {
		f[1<<i][v] = 0
	}

	inQ := make([]bool, n)
	for s := 1; s < len(f); s++ {
		fs := f[s]
		for sub := s & (s - 1); sub > s^sub; sub = (sub - 1) & s {
			for i := 0; i < n; i++ {
				fs[i] = min(fs[i], f[sub][i]+f[s^sub][i]) // 合并最短路
			}
		}
		// SPFA
		q := []int{}
		for i, fv := range fs {
			if fv < inf {
				q = append(q, i)
				inQ[i] = true
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for _, e := range g[v] {
				w := e.to
				if newD := fs[v] + e.wt; newD < fs[w] {
					fs[w] = newD
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
	}
	ans := slices.Min(f[len(f)-1])
	return ans
}

// 最小生成树 MST Kruskal
// 适用于稀疏图 O(mlogm)，或者边已经按权值排序的情况
// 性质：
// - 对于不同的 MST，同一边权的边的个数都是相同的（应用见后面的最小生成树计数）
// - 对于任意正确加边方案，加完小于某权值的边后，图的连通性是一样的
// - 任意点对间的路径，MST 做到了最小化最大边权（反证法）
// https://oi-wiki.org/graph/mst/#kruskal
// https://cp-algorithms.com/graph/mst_kruskal.html
// 边权 [0,1] 的随机完全图的 MST 权值和是 ζ(3) = 1.202…
// - https://en.wikipedia.org/wiki/Random_minimum_spanning_tree
// - https://www.sciencedirect.com/science/article/pii/0166218X85900587
//
// TIPS: 混合点权边权的问题，可以创建一个超级源点，把每个点 i 和超级源点相连，边权为点 i 的点权。这样就转换成了 MST 问题。
// LC1168 https://leetcode.cn/problems/optimize-water-distribution-in-a-village/
//
// 模板题 https://www.luogu.com.cn/problem/P3366 
//       https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/E
//       https://atcoder.jp/contests/abc218/tasks/abc218_e
// 关键边、伪关键边（与割边结合）https://codeforces.com/problemset/problem/160/D 2300
// - LC1489 https://leetcode.cn/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
// https://atcoder.jp/contests/abc210/tasks/abc210_e 需要一些数论知识
// https://atcoder.jp/contests/abc270/tasks/abc270_f 枚举
// https://codeforces.com/problemset/problem/1468/J 1800 分类讨论
// https://codeforces.com/problemset/problem/1095/F 1900 边权为 a[i]+a[j] 的混合 MST
// - 完全图找个最小的 a[i] 和其余点连边
// https://atcoder.jp/contests/arc076/tasks/arc076_b 1615=CF1932 需要一点注意力
// - https://www.luogu.com.cn/problem/P8074
// https://codeforces.com/problemset/problem/1108/F 2100 MST 的唯一性
// https://codeforces.com/problemset/problem/733/F 2200 与 LCA 结合
// https://codeforces.com/problemset/problem/891/C 2300 判断给定的边是否均在同一棵 MST 中
// https://codeforces.com/problemset/problem/1408/E 2400 二分图无环
// https://www.luogu.com.cn/problem/P1547 最小生成树的最长边：Kruskal 中最后一条加入 MST 中的边的长度
// https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/ EXTRA: 与树链剖分结合可以在线查询两点间路径最大边权的最小值
// https://codeforces.com/problemset/problem/1149/D todo 只有两种边权的图的 MST 的性质 + 所有 MST 中的单源最短路的最小值 
// https://ac.nowcoder.com/acm/contest/1056/A 算法竞赛进阶指南 走廊泼水节
// - https://codeforces.com/problemset/problem/1857/G 2000
// https://codeforces.com/problemset/problem/1707/C 2400 与 DFS 搜索树结合
// https://codeforces.com/problemset/problem/632/F 2400 转换
// https://codeforces.com/problemset/problem/1981/E 2600
// https://atcoder.jp/contests/abc282/tasks/abc282_e 无环即生成树
// https://atcoder.jp/contests/abc355/tasks/abc355_f
// https://atcoder.jp/contests/typical90/tasks/typical90_ai 子树 MST 必须包含特殊点
func (*graph) mstKruskal(n int, edges [][]int) int {
	// 边权范围小的话也可以用桶排
	slices.SortFunc(edges, func(a, b []int) int { return a[2] - b[2] })

	fa := make([]int, n) // n+1
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

	sum := 0
	cntE := 0
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		//v, w, wt := e.v, e.w, e.wt
		fv, fw := find(v), find(w)
		if fv != fw {
			fa[fv] = fw
			sum += wt
			cntE++
		}
	}

	// 图不连通
	if cntE < n-1 {
		return -1
	}
	return sum
}

// 最小生成树 MST Prim
// 适用于稠密图 O(n^2)，传入邻接矩阵 dis
// dis[v][w] == inf 表示没有 v-w 边
// 有些题目需要在连通分量上求 MST，这时就需要用到 root
// 可视化 https://visualgo.net/zh/mst
// https://oi-wiki.org/graph/mst/#prim
// 模板题 https://www.luogu.com.cn/problem/P1546
// LC1584 https://leetcode.cn/problems/min-cost-to-connect-all-points/
// https://codeforces.com/problemset/problem/1245/D 1900 超级源点 打印方案 
// https://codeforces.com/contest/1508/problem/C
// https://codeforces.com/problemset/problem/632/F
// todo https://codeforces.com/problemset/problem/959/E
func (*graph) mstPrim(dis [][]int, root int) (mstSum int, edges [][2]int) {
	edges = make([][2]int, 0, len(dis)-1)

	// 注意：dis 需要保证 dis[i][i] = inf，从而避免自环的影响

	const inf int = 2e9
	// minD[i].d 表示当前 MST 到点 i 的最小距离，对应的边为 minD[i].v-i
	minD := make([]struct{ v, d int }, len(dis))
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
		mstSum += minD[v].d
		if v != root {
			edges = append(edges, [2]int{minD[v].v, v})
		}

		// 更新 minD
		for w, d := range dis[v] {
			// 注：若 mstPrim 结束后 minD 无其他用途，!inMST[w] 的判断可以去掉
			if !inMST[w] && d < minD[w].d {
				minD[w].d = d
				minD[w].v = v
			}
		}
	}
}

// 处理多个连通块的 Prim 算法
// 返回每个连通块的 MST 边权和 mstSum，以及每个 MST 的边 edges
// 注意：dis 需要保证 dis[i][i] = inf，从而避免自环的影响
func (*graph) mstPrimMultiComp(dis [][]int) (totalMST int, mstSum []int, edges [][][2]int) {
	const inf int = 2e9

	nodes := []int{}
	vis := make([]bool, len(dis))
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		nodes = append(nodes, v)
		for w, d := range dis[v] {
			if d < inf && !vis[w] {
				dfs(w)
			}
		}
	}

	// minD[i].d 表示当前 MST 到点 i 的最小距离，对应的边为 minD[i].v-i
	minD := make([]struct{ v, d int }, len(dis))
	for i := range minD {
		minD[i].d = inf
	}
	inMST := make([]bool, len(dis)) // 初始时所有点都不在 MST 中
	for root, b := range vis {
		if b {
			continue
		}
		nodes = []int{}
		dfs(root)

		sum := 0
		es := [][2]int{}
		minD[root].d = 0
		for {
			// 根据切分定理，求不在当前 MST 的点到当前 MST 的最小距离，即 minD[v].d
			v := -1
			for _, w := range nodes {
				if !inMST[w] && (v < 0 || minD[w].d < minD[v].d) {
					v = w
				}
			}
			if v < 0 { // 已求出 MST
				break
			}

			// 加入 MST
			inMST[v] = true
			sum += minD[v].d
			if v != root {
				es = append(es, [2]int{minD[v].v, v})
			}

			// 更新 minD
			for _, w := range nodes {
				if !inMST[w] && dis[v][w] < minD[w].d {
					minD[w].d = dis[v][w]
					minD[w].v = v
				}
			}
		}
		totalMST += sum
		mstSum = append(mstSum, sum)
		edges = append(edges, es)
	}
	return
}

// Boruvka's algorithm
// 用于求解边权互不相同的无向图的最小生成森林
// 见 trie01.go 中的 xorMST
// https://en.wikipedia.org/wiki/Bor%C5%AFvka%27s_algorithm
// https://oi-wiki.org/graph/mst/#boruvka
// https://www.geeksforgeeks.org/boruvkas-algorithm-greedy-algo-9/
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BoruvkaMST.java.html
// https://codeforces.com/problemset/problem/888/G
// todo https://codeforces.com/problemset/problem/1550/F https://www.luogu.com.cn/blog/ETHANK/boruvka-xiao-ji

// 单点度数（单度）限制最小生成树   O(n^2)
// 点 root 的度数不超过 lim
// 不超过 http://poj.org/problem?id=1639 https://codeforces.com/gym/100227 A https://www.acwing.com/problem/content/349/
// todo EXTRA: 恰好的情况（需要用 WQS 二分）https://codeforces.com/problemset/problem/125/E
func (*graph) limitDegreeMST(dis [][]int, root, lim int) int {
	const inf int = 2e9

	n := len(dis)
	mstSum := 0
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
func (*graph) strictlySecondMST(n int, edges []struct{ v, w, wt int }) int {
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

	mstSum := 0
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
// LC1584 https://leetcode.cn/problems/min-cost-to-connect-all-points/
// 做法见官方题解 https://leetcode.cn/problems/min-cost-to-connect-all-points/solution/lian-jie-suo-you-dian-de-zui-xiao-fei-yo-kcx7/
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
		slices.SortFunc(ps, func(a, b pair) int { return a.v - b.v })
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
			mst += e.dis
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
	slices.SortFunc(edges, func(a, b [3]int) int { return a[2] - b[2] })
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
func (*graph) msaEdmonds(n, root int, edges [][3]int) (ans int) {
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
			ans += wt
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
// 具体请看我的题解：https://www.luogu.com.cn/blog/endlesscheng/solution-cf1242b
// https://codeforces.com/blog/entry/93652
//
// https://codeforces.com/problemset/problem/1242/B 1900
// https://codeforces.com/problemset/problem/190/E  2100
// https://codeforces.com/problemset/problem/920/E  2100
// https://codeforces.com/problemset/problem/653/E  2400 单点度数限制
// https://codeforces.com/problemset/problem/1508/C 2500 MST
func (*graph) inverseGraphComponents(g [][]int) [][]int {
	// 原图度数最小的点即为反图度数最大的点
	maxV := 0
	for v, ws := range g {
		if len(ws) < len(g[maxV]) {
			maxV = v
		}
	}

	// 特殊情况：原图是完全图，反图无边
	if len(g[maxV]) == len(g)-1 {
		components := make([][]int, len(g))
		for i := range components {
			components[i] = []int{i} // i+1
		}
		return components
	}

	ccNum := len(g)
	fa := make([]int, len(g))
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
	merge := func(from, to int) {
		f, t := find(from), find(to)
		if f == t {
			return
		}
		ccNum--
		fa[f] = t
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	has := make([]int, len(g))
	ts := 0
	mergeInv := func(v int) {
		ts++
		has[v] = ts
		for _, w := range g[v] {
			has[w] = ts
		}
		for i := range g {
			if has[i] != ts {
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

	tmp := map[int][]int{}
	for i := range fa {
		rootV := find(i)
		tmp[rootV] = append(tmp[rootV], i) // i+1
	}
	components := make([][]int, 0, len(tmp))
	for _, nodes := range tmp {
		components = append(components, nodes)
	}
	return components
}

// 二分图判定+染色      二分图染色
// 注：二分图也叫偶图
// https://en.wikipedia.org/wiki/Bipartite_graph
// https://oi-wiki.org/graph/bi-graph/#_3
// https://cp-algorithms.com/graph/bipartite-check.html
// 辅助证明 https://codeforces.com/contest/1839/problem/E
//
// 模板题 LC886 https://leetcode.cn/problems/possible-bipartition/
// https://codeforces.com/problemset/problem/862/B 1300 考察定义 树至多加多少条边仍然是二分图
// https://codeforces.com/problemset/problem/1093/D 1700
// https://codeforces.com/problemset/problem/1354/E 2100 与分组背包结合
// https://codeforces.com/problemset/problem/1537/F 2200
// https://codeforces.com/problemset/problem/553/C 2200 染色的技巧
// https://codeforces.com/problemset/problem/662/B 2200 染色的技巧
// https://codeforces.com/problemset/problem/85/E 2600 转换 
// https://codeforces.com/problemset/problem/547/D 2600 转换 
// todo https://codeforces.com/problemset/problem/741/C 2600 转换 
// https://www.luogu.com.cn/problem/P6185
// 与背包结合（NEERC01，紫书例题 9-19，UVa 1627）https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4502
func (*graph) isBipartite(g [][]int) bool {
	colors := make([]int8, len(g)) // 0 表示未访问该节点
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			// 如果要分组，传入 3^c，后续可以当成下标
			// 如果要根据颜色来确定正负号，传入 -c
			if colors[w] == c || colors[w] == 0 && !f(w, 3^c) {
				return false
			}
		}
		return true
	}
	// 可能有多个连通块
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
// 注：更推荐用 Dinic，可以达到 O(m√n) 的复杂度（而且这个复杂度一般不会跑满）
// - 超级源点 S 连左部，右部连超级汇点 T，所有边的容量均为 1，那么 S 到 T 的最大流即为二分图的最大匹配
// - 时间复杂度证明 https://www.cnblogs.com/Itst/p/12556871.html
// - 时间复杂度证明 https://oi-wiki.org/graph/flow/max-flow/#%E7%89%B9%E6%AE%8A%E6%83%85%E5%BD%A2%E4%B8%8B%E7%9A%84%E6%97%B6%E9%97%B4%E5%A4%8D%E6%9D%82%E5%BA%A6%E5%88%86%E6%9E%90
// - 代码 https://www.luogu.com.cn/record/123020820
// 可视化【推荐】 https://visualgo.net/zh/matching
// - 选择左下的图示 -> CP4 3.11a*，然后选择增广路 -> 标准
// https://www.renfei.org/blog/bipartite-matching.html 推荐
// https://oi-wiki.org/topic/graph-matching/bigraph-match/
// https://zhuanlan.zhihu.com/p/62981901
// https://en.wikipedia.org/wiki/Hall%27s_marriage_theorem
// https://www.geeksforgeeks.org/maximum-bipartite-matching/
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BipartiteMatching.java.html
// 有关随机贪心(匹配)预处理的 hack https://bzoj.blog.uoj.ac/blog/2427
//
// 模板题 https://www.luogu.com.cn/problem/P3386 https://www.luogu.com.cn/problem/B3605
// LC1349 https://leetcode.cn/problems/maximum-students-taking-exam/
// LCP04 https://leetcode.cn/problems/broken-board-dominoes/
// LC2123 https://leetcode.cn/problems/minimum-operations-to-remove-adjacent-ones-in-matrix/
// 正则二分图匹配 https://loj.ac/p/180
// - 正则就是所有点的度数都一样
//【网络流 24 题】飞行员配对方案 https://loj.ac/p/6000 https://www.luogu.com.cn/problem/P2756
//【网络流 24 题】骑士共存（这题 Dinic 更快）https://loj.ac/p/6226 https://www.luogu.com.cn/problem/P3355
// https://codeforces.com/problemset/problem/1139/E 2400
// todo https://codeforces.com/contest/1404/problem/E
//      https://codeforces.com/problemset/problem/1783/F 题解 https://www.luogu.com.cn/blog/DaiRuiChen007/CF1783F

// 匈牙利算法 · 写法一
func (*graph) maxBipartiteMatchingHungarian(g [][]int) (match []int, cnt int) {
	match = make([]int, len(g))
	for i := range match {
		match[i] = -1
	}
	vis := make([]int, len(g))
	ts := 0
	var dfs func(int) bool
	dfs = func(v int) bool {
		vis[v] = ts
		for _, w := range g[v] {
			mw := match[w]
			if mw == -1 || vis[mw] != ts && dfs(mw) {
				match[w] = v
				match[v] = w
				return true
			}
		}
		return false
	}
	for root := range g {
		if match[root] != -1 {
			continue
		}
		ts++
		if dfs(root) {
			cnt++ // +=2
		}
	}
	return
}

// 匈牙利算法 · 写法二
// 适用左右两侧节点有明确区分的情况，要求 g 中存储的是左侧到右侧的单向边
// 常见于棋盘放置 1x2 骨牌，或者一些排列型约束的题目
// 找 m 个完美匹配 https://atcoder.jp/contests/abc317/tasks/abc317_g
// https://codeforces.com/problemset/problem/1139/E 2400
func (*graph) maxBipartiteMatchingHungarianLR(nl, nr int, g [][]int) (matchL []int, cnt int) {
	// matchL[leftNode] = rightNode
	// matchR[rightNode] = leftNode
	matchL = make([]int, nl) // 可选
	for i := range matchL {
		matchL[i] = -1
	}
	matchR := make([]int, nr)
	for i := range matchR {
		matchR[i] = -1
	}
	vis := make([]int, len(matchL))
	ts := 0
	var dfs func(int) bool
	dfs = func(v int) bool {
		vis[v] = ts
		for _, w := range g[v] {
			lv := matchR[w]
			if lv == -1 || vis[lv] != ts && dfs(lv) {
				matchR[w] = v
				matchL[v] = w
				return true
			}
		}
		return false
	}
	for root := range g {
		ts++
		if dfs(root) {
			cnt++ // +=2
		}
	}
	return
}

// 二分图最大匹配 Hopcroft–Karp 算法 O(m√n)
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
// LC1595 https://leetcode.cn/problems/minimum-cost-to-connect-two-groups-of-points/solution/kai-kai-yan-jie-zhuan-huan-cheng-zui-da-dai-quan-p/
// LC1879 https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/
// LC1947 https://leetcode.cn/problems/maximum-compatibility-score-sum/
// LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/
// LC2403 https://leetcode.cn/problems/minimum-time-to-kill-all-monsters/
// todo GCJ21 Round2D https://codingcompetitions.withgoogle.com/codejam/round/0000000000435915/00000000007dc2de
func (*graph) maxWeightedBipartiteMatchingKuhnMunkresSlow(wt [][]int) (match []int, sum int) {
	const inf int = 1e18
	// NOTE: wt 中不存在的边应初始化为 -inf

	match = make([]int, len(wt)) // 右部点匹配了哪一个左部点
	for i := range match {
		match[i] = -1
	}
	// 初始化顶标
	la := make([]int, len(wt))
	for i, r := range wt {
		la[i] = r[0]
		for _, w := range r[1:] {
			if w > la[i] {
				la[i] = w
			}
		}
	}
	lb := make([]int, len(wt))
	slack := make([]int, len(wt))
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
func (*graph) maxWeightedBipartiteMatchingKuhnMunkres(wt [][]int) (match []int, sum int) {
	const inf int = 1e18
	// NOTE: wt 中不存在的边应初始化为 -inf

	match = make([]int, len(wt)) // 右部点匹配了哪一个左部点
	la := make([]int, len(wt))
	for i, row := range wt {
		la[i] = -inf
		for _, v := range row {
			if v > la[i] {
				la[i] = v
			}
		}
	}
	lb := make([]int, len(wt))
	slack := make([]int, len(wt))
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
//
// 拓扑序关键节点：对于【优先 pop 小 id 的拓扑排序】和【优先 pop 大 id 的拓扑排序】，如果一个点在这两个拓扑序中的下标都是一样的，那么这个点就是关键节点
//
// 模板题 https://www.luogu.com.cn/problem/B3644
// 树上拓扑+记录变成叶子的时间 LC2603 https://leetcode.cn/problems/collect-coins-in-a-tree/
// DAG DP LC2050 https://leetcode.cn/problems/parallel-courses-iii/
//        LC1857 https://leetcode.cn/problems/largest-color-value-in-a-directed-graph/
//        https://ac.nowcoder.com/acm/contest/6384/C
//        https://www.luogu.com.cn/problem/P3387
//        https://codeforces.com/problemset/problem/721/C 1800
// 删点 LC310 https://leetcode.cn/problems/minimum-height-trees/
// - https://codeforces.com/problemset/problem/1881/F
// https://codeforces.com/problemset/problem/1572/A 1800
// 混合图拓扑排序+定向 https://codeforces.com/problemset/problem/1385/E 2000
//                  https://codeforces.com/problemset/problem/1100/E 2200
// 构造 https://codeforces.com/problemset/problem/269/C 2100
// 与「环形」结合 https://codeforces.com/contest/1863/problem/E  2100
// 好题 https://codeforces.com/problemset/problem/915/D 2200
// 利用拓扑思想 https://codeforces.com/problemset/problem/1037/E 2200
// 与堆结合 https://codeforces.com/problemset/problem/825/E 2300
//         https://codeforces.com/problemset/problem/1283/F 2200
// 缩点后的拓扑序 https://codeforces.com/contest/1463/problem/E 2400
// 关键点 次关键点 https://codeforces.com/contest/1062/problem/F 2900
// 拓扑序是否唯一：任意时刻队列中不能有超过一个元素
// - LC444 https://leetcode.cn/problems/sequence-reconstruction/
// 检查一个序列是否为拓扑序，可以仿造拓扑排序的算法，从前往后检查节点的入度是否为 0，然后减少相邻节点的入度，直到找到一个入度不为 0 的点或者遍历到末尾
func (*graph) topoSort(n int, edges [][]int) []int {
	g := make([][]int, n)
	deg := make([]int, len(g))
	for _, e := range edges {
		v, w := e[0], e[1]
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
// 支持自环和重边
// https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
// https://oi-wiki.org/graph/scc/#kosaraju
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KosarajuSharirSCC.java.html
// 模板题 https://www.luogu.com.cn/problem/P1726
//       https://atcoder.jp/contests/practice2/tasks/practice2_g
// 缩点后与 DAG DP 结合 https://codeforces.com/contest/1900/problem/E
// 计数 https://codeforces.com/problemset/problem/427/C
// https://www.luogu.com.cn/problem/P2341
// 建图转换 https://codeforces.com/problemset/problem/1239/D
// 与高斯消元结合 https://www.luogu.com.cn/problem/P6030
// https://codeforces.com/problemset/problem/999/E 2000
func (*graph) sccKosaraju(n int, edges [][]int) ([][]int, []int) {
	g := make([][]int, n)
	rg := make([][]int, len(g))
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
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

	allScc := [][]int{}
	clear(vis)
	var scc []int
	var rdfs func(int)
	rdfs = func(v int) {
		vis[v] = true
		scc = append(scc, v)
		for _, w := range rg[v] {
			if !vis[w] {
				rdfs(w)
			}
		}
	}
outer:
	for i := len(g) - 1; i >= 0; i-- { // 逆后序遍历，就可以像无向图那样求出 SCC
		v := vs[i]
		if vis[v] {
			continue
		}
		scc = []int{}
		rdfs(v)
		allScc = append(allScc, scc)

		// EXTRA: 直接在这里判断缩点后是否出度为 0（无需实际缩点）
		{
			for _, u := range scc {
				for _, w := range g[u] {
					if !vis[w] { // 出度不为 0
						continue outer
					}
				}
			}
			// 出度为 0
		}
	}

	// 记录每个点所属 SCC 的下标，用于缩点和查询
	sid := make([]int, len(g))
	for i, cc := range allScc {
		// 还可以汇合同一个 SCC 的权值等 ...
		for _, v := range cc {
			sid[v] = i
		}
	}

	// 缩点: 将边 v-w 转换成 sid[v]-sid[w]
	// 缩点后得到了一张 DAG，点的编号范围为 [0,len(scc)-1]
	// 注意这样可能会产生重边，不能有重边时可以用 map 或对每个点排序去重
	// 模板题 点权 https://www.luogu.com.cn/problem/P3387
	// 		 边权 https://codeforces.com/contest/894/problem/E
	// 检测路径是否可达/唯一/无穷 https://codeforces.com/problemset/problem/1547/G
	ns := len(allScc)
	g2 := make([][]int, ns)
	deg := make([]int, ns) // 求拓扑序 DP
	for v, ws := range g {
		v = sid[v]
		for _, w := range ws {
			w = sid[w]
			if v != w {
				g2[v] = append(g2[v], w)
				deg[w]++
			}
		}
	}

	// EXTRA: 求有多少个点能被其他所有点访问到
	// https://www.luogu.com.cn/problem/P2341
	// 能到其它所有点 https://codeforces.com/problemset/problem/1777/E
	// - Tarjan 写法 https://codeforces.com/problemset/submission/1777/204187501
	numCanBeVisitedFromAll := func() int {
		clear(vis)
		lastComp := allScc[len(allScc)-1]
		rdfs(lastComp[0]) // 在反图上遍历
		for _, b := range vis {
			// 原图不是连通的
			if !b {
				return 0
			}
		}
		return len(lastComp)
	}
	_ = numCanBeVisitedFromAll

	return allScc, sid
}

// SCC Tarjan
// 支持自环和重边
// 常数比 Kosaraju 略小（在 AtCoder 上的测试显示，5e5 的数据下比 Kosaraju 快了约 100ms）
// CF 上差不多 Kosaraju 436ms https://codeforces.com/problemset/submission/1900/234527101
//            Tarjan   451ms https://codeforces.com/problemset/submission/1900/234526944
// - 估计是因为递归栈的开销导致的（64 位递归栈比 32 位的更耗内存，更容易 cache miss）
// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
// https://oi-wiki.org/graph/scc/#tarjan
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TarjanSCC.java.html
// https://stackoverflow.com/questions/32750511/does-tarjans-scc-algorithm-give-a-topological-sort-of-the-scc
// 与最小割结合 https://www.luogu.com.cn/problem/P4126
func (*graph) sccTarjan(g [][]int) ([][]int, []int) {
	// low(v) 定义为以下两种情况的最小值
	// 1. dfn[v]
	// 2. subtree(v) 的返祖边所指向的节点的 dfn，也就是经过恰好一条不在 DFS 树上的边，能够到达 subtree(v) 的节点的 dfn
	allScc := [][]int{}
	dfn := make([]int, len(g)) // 值从 1 开始
	dfsClock := 0
	st := []int{}
	var tarjan func(int) int
	tarjan = func(v int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		st = append(st, v)
		for _, w := range g[v] {
			if dfn[w] == 0 {
				lowW := tarjan(w)
				lowV = min(lowV, lowW)
			} else {
				// 如果 0 < dfn[w] != inf，说明 w 在 st 中，那么找到 v 的到其祖先节点的边 v-w，用 dfn[w] 来更新 lowV
				// 实际上，不需要判断 dfn[w] == inf 的情况，直接取 min 即可
				lowV = min(lowV, dfn[w])
			}
		}
		if dfn[v] == lowV { // 无法从 v 出发，到达 v 的任一祖先
			scc := []int{}
			for {
				w := st[len(st)-1]
				st = st[:len(st)-1]
				// 避免搜索树上的另一棵子树上的点 v，通过横向边，把 dfn[w] 错误地更新到 lowV（注意 dfn[w] 都小于后面新遍历到的点的 dfn 值）
				dfn[w] = math.MaxInt // 相当于 inStack[w] = false
				scc = append(scc, w) // 如果只要计算 sid，可以在这里写 sid[w] = sccCnt
				if w == v {
					break
				}
			}
			allScc = append(allScc, scc) // sccCnt++
		}
		return lowV
	}
	for i, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(i)
		}
	}

	// 由于每个强连通分量都是在它的所有后继强连通分量被求出之后求得的
	// 上面得到的 allScc 是拓扑序的逆序
	slices.Reverse(allScc)

	// 缩点
	sid := make([]int, len(g))
	for i, scc := range allScc {
		// 还可以汇合同一个 SCC 的权值等 ...
		for _, v := range scc {
			sid[v] = i
		}
	}

	g2 := make([][]int, len(allScc))
	deg := make([]int, len(allScc))
	for v, ws := range g {
		v = sid[v]
		for _, w := range ws {
			w = sid[w]
			if v != w {
				g2[v] = append(g2[v], w)
				deg[w]++
			}
		}
	}

	return allScc, sid
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
//       A 为真          (A)     ¬A⇒A     注：假可以推出一切   或者说：真不能推出假，假可以推出真
//       A 为假          (¬A)    A⇒¬A     同上
//       A 为真 B 就为真          A⇒B, ¬B⇒¬A
//       A 为假 B 就为假          ¬A⇒¬B, B⇒A
//       A,B 至少存在一个 (A|B)    ¬A⇒B, ¬B⇒A 意思是一个为假的时候，另一个一定为真 https://www.luogu.com.cn/problem/P4782
//       A,B 不能同时存在 (¬A|¬B)  A⇒¬B, B⇒¬A 就是上面的式子替换了一下（一个为真，另一个一定为假）
//       A,B 必须且只一个 (A^B)    A⇒¬B, B⇒¬A, ¬A⇒B, ¬B⇒A
//       A,B 同时或都不在 (¬(A^B)) A⇒B, B⇒A, ¬A⇒¬B, ¬B⇒¬A
// 模板题 https://www.luogu.com.cn/problem/P4782
// 建边练习 https://codeforces.com/contest/468/problem/B
// 建边练习 https://codeforces.com/contest/1971/problem/H
// 定义 Ai 表示「选 Xi」，这样若两个旗子 i j 满足 |Xi-Xj|<D 时，就相当于 Ai Aj 至少一个为假。其他情况类似 
// - https://atcoder.jp/contests/practice2/tasks/practice2_h
// https://codeforces.com/problemset/problem/1657/F 2600
func (G *graph) sat2(n int) []bool {
	// g 分为左右两部，左部 [0,n-1] 的点表示 x 为假的状态，右部 [n,2*n-1] 的点表示 x 为真的状态
	// 我们要做的，就是把状态连起来
	// 比如根据题意，若 x 为真，可以推出 y 为假
	// 就是添加一条 x+n 到 y 的有向边（以及逆否命题的边，见下）
	// 建好图后，求 SCC
	// 如果发现 x 和 x+n 在同一个 SCC 中，说明 x 为真和 x 为假同时成立，矛盾！
	// 如果不矛盾，我们还可以知道 x 是真还是假（见下）
	g := make([][]int, n*2)

	// 通过题意挖掘出如下性质（a b 为 0/1 表示假/真）
	// x=a ⇒ y=b
	// 注意这同时意味着其逆否命题也成立
	// y=b^1 ⇒ x=a^1
	// 如果原命题及其逆否命题成对出现，那么只需要对原命题连边
	// https://codeforces.com/contest/1971/problem/H
	addEdge := func(x, a, y, b int) {
		x--
		y--
		v, w := x+a*n, y+b*n
		g[v] = append(g[v], w)
		v, w = y+(b^1)*n, x+(a^1)*n
		g[v] = append(g[v], w)
	}

	// 通过题意挖掘出如下性质（a b 为 0/1 表示假/真）
	// x=a 和 y=b 两个条件至少一个成立
	// 用编程语言来说就是 x==a || y==b 成立
	// 比如 a=b=1，那么 x 为假可以推出 y 为真，y 为假可以推出 x 为真
	// 即 addEdge(x,a^1,y,b)
	// 推论：
	// 令 a=0，即 addEdge(x,1,x,0)，表示 x 恒为假
	// 令 a=1，即 addEdge(x,0,x,1)，表示 x 恒为真
	// https://www.luogu.com.cn/problem/P4782
	// https://codeforces.com/contest/1971/problem/H addClause(abs(a),a>0,abs(b),b>0)
	addClause := func(x int, a bool, y int, b bool) {
		va := 0
		if a {
			va = 1
		}
		vb := 0
		if b {
			vb = 1
		}
		addEdge(x, va^1, y, vb)
	}

	// read input & addEdge ...

	_, sid := G.sccTarjan(g)
	ans := make([]bool, n)
	for i, id := range sid[:n] {
		// x 和 ¬x 处于同一个 SCC 时无解（因为 x ⇔ ¬x）
		if id == sid[i+n] {
			return nil
		}
		// 由于假可以推出全部，所以：
		// sid[x] < sid[¬x] ⇔ (¬x ⇒ x) ⇔ x 为真
		// sid[x] > sid[¬x] ⇔ (x ⇒ ¬x) ⇔ x 为假
		ans[i] = id < sid[i+n]
	}

	_ = addClause
	return ans
}

// 基环树（环套树），英文名叫 pseudotree，基环树森林叫 pseudoforest
// https://en.wikipedia.org/wiki/Pseudoforest
// 对于内向基环树，由于每个点的出度均为一，可以用 []int 来表示图
// 基环树的讲解可以看一下我的这篇题解 https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/
// 如果题目涉及到基环树上的路径，用倍增会更简单
// todo https://www.luogu.com.cn/blog/user52918/qian-tan-ji-huan-shu
// todo 题单 https://www.luogu.com.cn/blog/ShadderLeave/ji-huan-shu-bi-ji
//
// LC684 并查集更简单 https://leetcode.cn/problems/redundant-connection/
// LC2876 每个点能访问到的点的个数 https://leetcode.cn/problems/count-visited-nodes-in-a-directed-graph/
// LC2127 https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/
// LC2359 单源最短路 https://leetcode.cn/problems/find-closest-node-to-given-two-nodes/ 这题样例没有给环的例子
// todo LC2836 https://leetcode.cn/problems/maximize-value-of-function-in-a-ball-passing-game
// - 加强版 https://codeforces.com/problemset/problem/702/E 2100
// - 更简单的做法是倍增
// 统计每个环的长度 LC2360 https://leetcode.cn/problems/longest-cycle-in-a-graph/
// - https://codeforces.com/problemset/problem/1833/E 1600
// - https://codeforces.com/problemset/problem/711/D 1900
// https://codeforces.com/problemset/problem/1670/C 1400 无向图
// https://codeforces.com/problemset/problem/131/D  1600 【无向图 建反图】
// https://codeforces.com/problemset/problem/1530/D 1600 把基环树拆分成若干条链，然后把这些链合并成一个大环 
// https://codeforces.com/problemset/problem/1873/H 1700 复杂分类讨论
// https://codeforces.com/problemset/problem/1027/D 1700
// https://codeforces.com/problemset/problem/1770/D 2000
// https://codeforces.com/problemset/problem/1335/F 2200
// https://codeforces.com/problemset/problem/1907/G 2200
// https://codeforces.com/problemset/problem/1200/F 2300 拆点
// https://codeforces.com/problemset/problem/1270/G 2700 构造 建图
// https://atcoder.jp/contests/abc266/tasks/abc266_f
// https://ac.nowcoder.com/acm/contest/9977/c 删除一条边使得直径最长
// [IOI2008] 岛屿 https://www.luogu.com.cn/problem/P4381
// todo [NOI2013] 快餐店 https://www.luogu.com.cn/problem/P1399
func (*graph) pseudotree(g []int) { // g 为内向基环树（森林）
	deg := make([]int, len(g)) // g 上每个节点的入度
	for _, w := range g {
		deg[w]++
	}

	// 拓扑排序，剪掉 g 上的所有树枝
	// 拓扑排序后 deg 值为 1 的点必定在基环上，为 0 的点必定在树枝上
	// 注：拓扑排序时还可以做 DP，比如给树枝上的每个点标记反向深度
	rg := make([][]int, len(g)) // g 的反图（外向基环树）
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			//f[i] = 1
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w := g[v]
		// 顺便建反图（在这里建反图可以避免加入基环上的边）
		rg[w] = append(rg[w], v)
		//f[v] += a[v]
		//f[w] = max(f[w], f[v])
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	// 在反图上遍历入环点+树枝
	// 由于考虑了入环点，所以可以遍历到基环树上的所有点
	var rdfs func(int, int)
	rdfs = func(v, depth int) {
		// ...
		for _, w := range rg[v] {
			rdfs(w, depth+1)
		}
	}

	// 注意可能有多棵基环树
	for i0, d := range deg {
		if d == 0 {
			continue
		}

		// 收集基环上的点
		ring := []int{i0}
		for v := g[i0]; v != i0; v = g[v] {
			deg[v] = 0
			ring = append(ring, v)
		}

		// 遍历基环
		// 特别注意基环大小小于 3 的特殊情况
		for _, v := range ring {
			rdfs(v, 0)
		}
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

		branchQueue := []int{}
		for i, d := range deg {
			if d == 1 {
				branchQueue = append(branchQueue, i)
			}
		}
		for len(branchQueue) > 0 {
			v := branchQueue[0]
			branchQueue = branchQueue[1:]
			for _, w := range g[v] {
				if deg[w]--; deg[w] == 1 {
					branchQueue = append(branchQueue, w)
				}
			}
		}

		var curRoot int
		_ = curRoot
		// v 到 curRoot 的距离为 d
		var dfsBranch func(v, fa, d int) int
		dfsBranch = func(v, fa, d int) int {
			size := 1
			for _, w := range g[v] {
				if w != fa && deg[w] < 2 {
					sz := dfsBranch(w, v, d+1)
					// do sz ...
					size += sz
				}
			}
			return size
		}
		size := 0
		for root, d := range deg {
			if d > 1 { // root 在基环上
				curRoot = root
				sz := dfsBranch(root, -1, 0)
				// do sz ...
				size += sz
			}
		}

		{
			// 遍历基环（一般是求长度）
			vis := make([]bool, len(g))
			var f func(int) int
			f = func(v int) int {
				vis[v] = true
				size := 1
				for _, w := range g[v] {
					if !vis[w] && deg[w] > 1 {
						size += f(w)
					}
				}
				return size
			}
			for root, b := range vis {
				if !b && deg[root] > 1 {
					sz := f(root)
					// do sz ...
					_ = sz
				}
			}

			// 效率更高的非递归写法
			// 见 https://leetcode.cn/problems/longest-cycle-in-a-graph/solution/nei-xiang-ji-huan-shu-zhao-huan-li-yong-pmqmr/
			var to []int // 有向图无向图都可以
			time := make([]int, len(to))
			clock := 1
			for x, t := range time {
				if t > 0 { // 已访问
					continue
				}
				for t0 := clock; x >= 0; x = to[x] {
					if time[x] > 0 {
						if time[x] >= t0 { // 如果 < t0 意味着当前走的是树枝
							// 注：记录每个时间对应的节点，那么从 time[x] 到 clock-1 的点按顺序都在环上
							sz := clock - time[x]
							// do sz ...
							_ = sz
						}
						break
					}
					time[x] = clock
					clock++
				}
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

// 网格图上的基环树
// 输入一个由上下左右箭头组成的网格图
// a[i][j] 表示第 i 行第 j 列的方向，0123 分别表示上下左右
// 例如这张图：https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/29/grid2.png
// - 注意图中有一个大小为 6 的基环
// 问：从网格图中的哪个坐标出发，可以访问到的不同坐标数最多？（不含出界坐标）
// 返回三个数：坐标个数、起点行坐标、起点纵坐标（从 0 开始）
// 如果有多个符合要求的起点坐标，返回最小字典序（取行坐标小的起点，行坐标相等的起点取纵坐标小的）
func (*graph) pseudotreeGrid(a [][]int) []int {
	n, m := len(a), len(a[0])
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	deg := make([][]int, n)
	for i := range deg {
		deg[i] = make([]int, m)
	}

	g := make([][]pair, n)
	for i, row := range a {
		g[i] = make([]pair, m)
		for j, v := range row {
			d := dir4[v]
			x, y := i+d.x, j+d.y
			if x < 0 || x >= n || y < 0 || y >= m {
				x, y = i, j // 出界指向自己，从而保证网格图是一棵内向基环树
			}
			g[i][j] = pair{x, y}
			deg[x][y]++
		}
	}

	rg := make([][][]pair, n)
	for i := range rg {
		rg[i] = make([][]pair, m)
	}
	q := []pair{}
	for i, row := range deg {
		for j, d := range row {
			if d == 0 {
				q = append(q, pair{i, j})
			}
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		to := g[from.x][from.y]
		if to.x < 0 {
			continue
		}
		x, y := to.x, to.y
		rg[x][y] = append(rg[x][y], from)
		if deg[x][y]--; deg[x][y] == 0 {
			q = append(q, to)
		}
	}

	var ans, ansI, ansJ int
	var rdfs func(int, int, int)
	rdfs = func(x, y, depth int) {
		if depth > ans || depth == ans && (x < ansI || x == ansI && y < ansJ) {
			ans, ansI, ansJ = depth, x, y
		}
		for _, from := range rg[x][y] {
			rdfs(from.x, from.y, depth+1)
		}
	}
	for i0, row := range deg {
		for j0, d := range row {
			if d == 0 {
				continue
			}
			ring := []pair{{i0, j0}}
			for p := g[i0][j0]; p.x != i0 || p.y != j0; p = g[p.x][p.y] {
				deg[p.x][p.y] = 0
				ring = append(ring, p)
			}
			for _, p := range ring {
				rdfs(p.x, p.y, len(ring))
			}
		}
	}
	return []int{ans, ansI, ansJ}
}

// yy 的东西：
// n 点 n+1 条边（双环树）
// 分类讨论，多加的一条边在哪：
// - 这条边的两个点在同一个树枝上：再形成一个基环（再细分，可以讨论 LCA 是否在基环上）
// -- 真桥、割点
// - 不在同一个树枝上：两个环，共享一条环上的路径
// - 树枝连基环：同上
// - 基环上两点相连
// -- 伪桥
// 性质：
// - 直径？
// - 任意两点最短路？
// - MST？
// - 最大独立集
// - 最小点覆盖
// 变形：
// - 这条边加在基环树的哪里，从而满足<?>性质/最优化<?>

// 圆方树的内容在上面

/* 网络流·总结·题单 ################################################################################

最大流等于最小割的证明 https://seineo.github.io/%E5%9B%BE%E8%AE%BA%EF%BC%9A%E6%9C%80%E5%A4%A7%E6%B5%81%E6%9C%80%E5%B0%8F%E5%89%B2%E8%AF%A6%E8%A7%A3.html

网络流建模方式总结
最小割问题秒杀三板斧 https://www.bilibili.com/video/BV1jt4y1t7pd/
https://atcoder.jp/contests/abc326/tasks/abc326_g
套三板斧：
先把所有收益都拿到，然后减去放弃的成就和花费
这样问题变成计算最小代价
第一板斧：
skill 属于 S 集合表示不选择这个技能，属于 T 集合表示选择这个技能
achieve 属于 T 集合表示不放弃这个成就，属于 S 集合表示放弃这个成就
第二板斧：
skill 的每个等级依赖自己的上一个等级，也就是如果该等级选择而上一等级未选择，不合法；此时上一等级属于 S 而下一等级属于 T，因此可以建上一等级到下一等级的边，流量 inf
achieve 属于 T 合而依赖的 skill 属于 S 集合不合法，因此可以建对应的 skill 到 achieve 的边，流量 inf
第三板斧：
S 到 skill 建边，代价为选择这个技能（等级）的代价
achieve 到 T 建边，流量为其收益

todo
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

Max-Flow，简称 MF

可视化 https://visualgo.net/zh/maxflow
选择左下的图示 - CS4234 MF Demo 或者 CP4 8.15*，然后选择左下的 Dinic - 前进

https://en.wikipedia.org/wiki/Maximum_flow

二分图最大匹配
超级源点连左部，右部连超级汇点，所有边的容量均为 1，最大流即为最大匹配
模板题 https://www.luogu.com.cn/problem/P3386
代码 https://www.luogu.com.cn/record/123020820
https://codeforces.com/problemset/problem/489/B 1200

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

如何输出最小割（只需要求一个解）
1. 求最大流
2. 从源点出发在残余网络上 DFS，标记所有能够到达的点
3. 遍历原边集 edges，若其中一端有标记，另一端没有标记，则这条边为最小割上的边

技巧 1：用容量为 inf 的边来防止割断
技巧 2：给边权加上很大的数来约束删除次数
https://codeforces.com/problemset/problem/700/C 2600

建模·转换
https://www.luogu.com.cn/problem/P1345
https://www.acwing.com/problem/content/2282/
平均边权最小 https://www.acwing.com/problem/content/2281/
点连通度 SEERC04 F https://codeforces.com/gym/101461 http://poj.org/problem?id=1966 UVa1660 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4535
   https://en.wikipedia.org/wiki/Connectivity_(graph_theory)
   https://en.wikipedia.org/wiki/Menger%27s_theorem
LCP38/21春·战队赛F https://leetcode.cn/problems/7rLGCR/
todo https://atcoder.jp/contests/arc085/tasks/arc085_c

todo 最小割必经边？

最大权闭合图 Maximum Weight Closure of a Graph
https://en.wikipedia.org/wiki/Closure_problem
源点向所有正权点连边，容量为相应点权
所有负权点向汇点连边，容量为相应点权的相反数
原图边的容量为 inf（从而保证不会在最小割中）
最后用正权点总和减去源点到汇点的最小割即为答案
以「最大获利」这题（边权和减点权和最大子图）来解释，割掉源点到正权点的边，意味着放弃对应用户的收益；割掉负权点到汇点的边，意味着建立对应基站
NOI06 最大获利 https://www.luogu.com.cn/problem/P4174
- https://codeforces.com/problemset/problem/1082/G 2400
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
todo https://atcoder.jp/contests/abc285/tasks/abc285_g

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
从源点到集合 A 中各点连边，容量为 1，费用为 0
从集合 B 中各点到汇点连边，容量为 1，费用为 0
集合 A 和 B 两两连边，容量为 inf（或者题目指定），费用为 F(Ai,Bj)（题目指定）
这样跑 MCMF 得到的结果是匹配全部 A（或 B）的最小花费
代表题目 https://codeforces.com/problemset/problem/237/E 2000 | 代码 https://codeforces.com/problemset/submission/237/241222973
https://codeforces.com/problemset/problem/1437/C 1800
LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/ 2392
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

/* 网络流建模技巧/转换技巧 ################################################################################
todo 整合到其它 blocks
todo 重新看一下挑战

标准建模（指派问题）：
	http://poj.org/problem?id=2175
	http://poj.org/problem?id=3686
边容量减少：
	若 flow<=cap' 则最大流不变；若 flow>cap' 需要将多出的流退回去 todo
    最小割+退流 https://www.luogu.com.cn/problem/P3308
流量任意：
	todo
容量为负数：
	todo
费用为负数：
	todo 挑战:228
求最小割划分成两个集合：
	Dual Core CPU http://poj.org/problem?id=3469
无重复边的往返最短路：
	http://poj.org/problem?id=2135
	转换成流量为 2 的最小费用流
点边转换
   将点拆为入点和出点（v 和 v+n），即可把点的属性变成边的属性，从而方便应用最大流、最小割等算法
   将边的中间加一个节点，把边的属性体现在中间的点上
上下界费用流
	对每条边新增一条边 e'
	e.cap-=minCap
	e'.cap=minCap
	e'.cost=e.cost-M // 一个足够大的常数
	跑完 MCMF 后加上 M*∑minCap

Disjoint paths
Edge-disjoint paths: It turns out that the maximum number of edge-disjoint paths equals the maximum flow of the graph, assuming that the capacity of each edge is one.
Node-disjoint paths: 拆点法

路径覆盖问题 Path cover + 打印
todo https://zhuanlan.zhihu.com/p/125759333
todo Competitive Programmer’s Handbook Ch.20
todo 线性规划与网络流 24 题 - 最小路径覆盖问题 https://byvoid.com/zhs/blog/lpf24-3/

给一 DAG，求它的最大反链大小、一组最大反链构造以及所有最大反链的并
https://yhx-12243.github.io/OI-transit/records/lydsy1143%3Blg4298.html
*/

// 最大流 Dinic's algorithm O(n^2 * m)  二分图上为 O(m√n)
// https://en.wikipedia.org/wiki/Dinic%27s_algorithm
// https://oi-wiki.org/graph/flow/max-flow/#dinic
// https://cp-algorithms.com/graph/dinic.html
// [Tutorial] My way of understanding Dinitz's ("Dinic's") algorithm https://codeforces.com/blog/entry/104960
// https://www.bilibili.com/video/BV1j64y1R7yK/
// 时间复杂度证明 https://www.zhihu.com/question/34374412
// - 一次增广令至少一条边饱和（cap = 0），所以有至多 O(m) 条增广路，每条增广路的长度为 O(n)，所以多路增广的时间是 O(nm)
// - 由于每次 BFS 都会使 d[end] 变大，至多变大 O(n) 次，所以总的时间复杂度为 O(n) * O(nm) = O(n^2 * m)
// 关于二分图上的时间复杂度，见 https://www.cnblogs.com/Itst/p/12556871.html
//
// 模板题 https://www.luogu.com.cn/problem/P3376
//       https://www.luogu.com.cn/problem/P2740
func (*graph) maxFlowDinic(n, st, end int, edges [][]int, a, b []int) int {
	type neighbor struct{ to, rid, cap, eid int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, eid int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, eid})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -1}) // 无向图上 0 换成 cap
	}
	for i, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap, i)
	}

	{
		// 最大匹配的建图
		st := len(a) + len(b)
		end := st + 1
		type neighbor struct{ to, rid, cap int }
		g := make([][]neighbor, end+1)
		addEdge := func(from, to, cap int) {
			g[from] = append(g[from], neighbor{to, len(g[to]), cap})
			g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
		}
		// 超级源点连左部，右部连超级汇点，所有边的容量均为 1，最大流即为最大匹配
		for i, v := range a {
			addEdge(st, i, 1) // 如果题目允许一对多，比如一对二，把 1 改成 2
			for j, w := range b {
				if v+w < 100 { // 和题目有关，满足该约束即可匹配 a[i] 和 b[j]
					addEdge(i, j+len(a), 1)
				}
			}
		}
		for j := range b {
			addEdge(j+len(a), end, 1) // 如果题目允许多对一，比如二对一，把 1 改成 2
		}
	}

	d := make([]int, len(g))
	bfs := func() bool {
		clear(d) // d[i] = 0 表示没有访问过
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
	// 当前弧，在其之前的边已经没有用了，避免对没有用的边进行多次检查
	iter := make([]int, len(g))
	// 寻找增广路（多路增广）
	var dfs func(int, int) int
	dfs = func(v, totalFlow int) (curFlow int) {
		if v == end {
			return totalFlow
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && d[w] > d[v] {
				f := dfs(w, min(totalFlow-curFlow, e.cap))
				if f == 0 {
					continue
				}
				e.cap -= f
				g[w][e.rid].cap += f
				curFlow += f
				if curFlow == totalFlow {
					break
				}
			}
		}
		return
	}
	maxFlow := 0
	for bfs() {
		clear(iter)
		maxFlow += dfs(st, math.MaxInt)
	}

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
	ans := make([]int, len(edges))
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
func (*graph) maxFlowISAP(n, st, end int, edges [][]int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for _, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap)
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
	const inf int = 1e18
	maxFlow := 0
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

func (*graph) maxFlowHLPP(n, st, end int, edges [][]int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap int } // rid 为反向边在邻接表中的下标
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
	}
	for _, e := range edges {
		v, w, edgeCap := e[0], e[1], e[2]
		addEdge(v, w, edgeCap)
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
//
// 模板题 https://www.luogu.com.cn/problem/P3381
// https://codeforces.com/problemset/problem/237/E 2000
//
// 常见建模方式（下面代码按照这种建模写的）
// 建模的时候，一般可以理解成在一个矩阵 a 上，每行每列至多选一个数，问所选数字之和的最小值
// 创建一个（完全）二分图，左部为行，右部为列
// - 行 -> 列，容量为 inf，费用为 grid[i][j]
// - 超级源点 S -> 行，容量为 1，费用为 0
// - 列 -> 超级汇点 T，容量为 1，费用为 0
// 如果求最大值可以把元素值（费用）取反，最后答案再取反
// 如果每行可以选多个，可以修改从超级源点到左部（行）的容量
// 如果每列可以选多个，可以修改从右部（列）到超级汇点的容量
// 如果要限制至多选 k 个元素，可以在超级源点前面再加一个节点，连到超级源点，容量为 k，费用为 0（相当于超级源点的流出量至多为 k）。如果满流，则表示恰好选了 k 个元素
//
// 完全二分图 + 一对一 LC3376 https://leetcode.cn/problems/minimum-time-to-break-locks-i/
// 完全二分图 + 一对多 LC2850 https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/
// 完全二分图 + 至多选 k=3 个数 LC3257 https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/
// 多对一 LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/
// 二分图 + 稀疏矩阵 LC3276 https://leetcode.cn/problems/select-cells-in-grid-with-maximum-score/
// https://codeforces.com/problemset/problem/1107/F 2600
func (*graph) minCostFlowSPFA(a [][]int) (int, int) {
	n := len(a)
	m := len(a[0])
	S := n + m
	T := S + 1

	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int } // 如果输入的是 edges，可以额外记录边的下标
	g := make([][]neighbor, T+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i, row := range a {
		for j, v := range row {
			addEdge(i, n+j, math.MaxInt, v) // 如果求最大，改成 -v
		}
		addEdge(S, i, 1, 0) // 如果是一对多，改 cap
		// 特别地，如果这一行的所有 v 都相同，可以把 S->i 的 cost 改成 v，i->n+j 的 cost 改成 0
	}
	for j := range a[0] {
		addEdge(n+j, T, 1, 0) // 如果是多对一，改 cap
	}
	//addEdge(T+1, S, k, 0) // 如果要限制至多选 k 个元素（g 的大小改成 T+2）

	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[T] < math.MaxInt
	}

	maxFlow := 0 // 可选
	minCost := 0
	for spfa() {
		// 沿 st-end 的最短路尽量增广
		// 特别地，如果建图时所有边的容量都设为 1，那么 minF 必然为 1，下面第一个 for 循环可以省略
		minF := math.MaxInt
		for v := T; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := T; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		maxFlow += minF
		minCost += dis[T] * minF
	}
	return maxFlow, minCost
}

// 基于原始对偶方法 (primal-dual method)
// https://blog.xehoth.cc/DurationPlan-Primal-Dual/
func (*graph) minCostFlowDijkstra(n, st, end, flowLimit int, edges [][]int) int {
	st--
	end--

	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, n)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for _, e := range edges {
		v, w, edgeCap, edgeCost := e[0], e[1], e[2], e[3]
		addEdge(v, w, edgeCap, edgeCost)
	}

	h := make([]int, len(g)) // 顶点的势
	dist := make([]int, len(g))
	type pair struct{ v, i int }
	fa := make([]pair, len(g))
	dijkstra := func() bool {
		const _inf int = 1e18
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
				if newD := dist[v] + e.cost + h[v] - h[w]; newD < dist[w] {
					dist[w] = newD
					fa[w] = pair{v, i}
					q.push(dijkstraPair{w, newD})
				}
			}
		}
		return dist[end] < _inf
	}
	minCost := 0
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
		flowLimit -= minF        // maxFlow += minF
		minCost += h[end] * minF // 注意这里是 h 不是 dist
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
// An independent set of EDGES is a set of edges of which no two have a vertex in common. It is usually called a matching. https://en.wikipedia.org/wiki/Independent_set_(graph_theory)#See_also
// https://en.wikipedia.org/wiki/Clique_problem
// 另见 Bron–Kerbosch 算法 https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
// Measure and Conquer: A Simple O(2^0.288n) Independent Set Algorithm http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.321.6920&rep=rep1&type=pdf
// todo 剪枝写法
// https://codeforces.com/problemset/problem/1105/E
// todo https://codeforces.com/problemset/problem/1767/E
func (*graph) maximalCliques(g []int) int {
	// 一种求最大团的做法，适用于点数不超过 50 的图
	// 传入的 g 为状压后的邻接矩阵
	// 定义 f(s) 为 s 的所有子集中最大团的大小
	// 则转移时要么不取 lb（lowbit），要么取 lb 并去掉不与 lb 相邻的点（包括 lb）
	// 将这一过程记忆化可大幅减少运行时间，理由如下：
	// 由于每次都会去掉 lb，所以至多 k=len(g)/2 次递归后会进入右半部分没有 1 的状态
	// 将这 k 次递归过程视作一棵二叉树，则其耗时为 O(2^k)
	// 之后记忆化占主导，耗时也为 O(2^k)
	// 主要注意的是，k 次递归的结果是否记忆化并不重要，因为这部分最多也只有 O(2^k) 个状态
	// 总的来说，记忆化将计算量由原来的「二叉树规模」变成了「meet in the middle 规模」
	memo := map[int]int{0: 0}
	var dfs func(int) int
	dfs = func(s int) int {
		if v, ok := memo[s]; ok {
			return v
		}
		memo[s] = max(dfs(s&(s-1)), 1+dfs(s&g[bits.TrailingZeros(uint(s))]))
		return memo[s]
	}
	ans := dfs(1<<len(g) - 1)
	return ans
}

// todo 极大团计数

// 分团覆盖问题 Clique cover
// https://en.wikipedia.org/wiki/Clique_cover

// todo 图的同构

// 树的同构见 graph_tree.go

// 支配树 Lengauer−Tarjan 算法
// https://en.wikipedia.org/wiki/Dominator_(graph_theory)
// https://oi-wiki.org/graph/dominator-tree/
// todo https://www.luogu.com.cn/blog/Zenislt/qian-tan-zhi-pei-shu-lengauer-tarjan-algorithm
// todo 模板题 https://www.luogu.com.cn/problem/P5180
//  https://www.luogu.com.cn/problem/P2597

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
func (*graph) countCycle3(n int, edges [][2]int) (ans int) {
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		v--
		w--
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
func (*graph) countCycle4(n int, edges [][2]int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		v--
		w--
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
					ans += cnt[u]
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
				mn, mx := v, v
				for x := v; x != w; {
					x = fa[x]
					if x < mn {
						mn = x
					} else if x > mx {
						mx = x
					}
				}
				minMax[v] = [2]int{mn, mx}
				for x := v; x != w; {
					x = fa[x]
					minMax[x] = [2]int{mn, mx}
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
				mn, mx := w, w
				for i := len(s) - 1; s[i] != w; i-- {
					x := s[i]
					if x < mn {
						mn = x
					} else if x > mx {
						mx = x
					}
				}
				minMax[w] = [2]int{mn, mx}
				for i := len(s) - 1; s[i] != w; i-- {
					x := s[i]
					minMax[x] = [2]int{mn, mx}
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

反复迭代 a <- a + a%10

https://codeforces.com/problemset/problem/1714/E 1400
https://codeforces.com/problemset/problem/1848/D 2200

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
