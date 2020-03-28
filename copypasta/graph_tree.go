package copypasta

import (
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

/*
NOTE: 对于有根树的题，可以考虑加上 g[0] = append(g[0], -1) 来简化代码
NOTE: 由于树上任意两点间的路径等价于两条点到根的路径的对称差，处理一些树上异或的问题可以往这个方向思考
NOTE: 注意特判只有一条边的情况，此时两个叶结点对应同一条边
*/

// namespace
type tree struct{}

// 节点深度
func (*tree) depth(n, root int, g [][]int) []int {
	dep := make([]int, n)
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		dep[v] = d
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 0)
	return dep
}

// 树上两点路径
func (*tree) path(st, end int, g [][]int) (path []int) {
	var f func(v, fa int) bool
	f = func(v, fa int) bool {
		if v == st {
			path = append(path, v)
			return true
		}
		for _, w := range g[v] {
			if w != fa {
				if f(w, v) {
					path = append(path, v)
					return true
				}
			}
		}
		return false
	}
	// 反向寻找
	f(end, -1)
	return
}

// 树上每个子树的信息：子树大小，DFS 序（从 1 开始）
// 这样的话 [o.dfn, o.dfn+o.size-1] 就表示一颗子树，方便用线段树维护
func (*tree) subtreeSize(n, root int, g [][]int) {
	type node struct{ size, dfn int }
	nodes := make([]node, n)
	dfn := 0
	var buildNode func(int, int) int
	buildNode = func(v, p int) int {
		dfn++
		nodes[v] = node{1, dfn}
		o := &nodes[v]
		for _, w := range g[v] {
			if w != p {
				o.size += buildNode(w, v)
			}
		}
		return o.size
	}
	buildNode(root, -1)
}

// 每个节点的入出时间戳
// 预处理后可以 O(1) 判断 fa 是否为 v 的祖先节点
// 例题 https://codeforces.ml/contest/1328/problem/E
func (*tree) inOutTimestamp(n, root int, g [][]int) {
	timeIn := make([]int, n)
	timeOut := make([]int, n)
	clock := 0
	var f func(v, fa int)
	f = func(v, fa int) {
		clock++
		timeIn[v] = clock
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
		clock++
		timeOut[v] = clock
	}
	f(0, -1)
	isFa := func(fa, v int) bool { return timeIn[fa] <= timeIn[v] && timeOut[v] <= timeOut[fa] }

	_ = isFa
}

// 树的直径/最长链
// 性质：直径的中点到所有叶子的距离和最小
// 也可以用树形 DP，计算每个根往下的最长链和次长链从而得到答案（维护最大时记录是从哪个节点取到的，维护次大时跳过该节点）
func (*tree) diameter(st int, g [][]int) (dv, dw int) {
	var u, maxD int
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD = d
			u = v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	maxD = -1
	f(st, -1, 0)
	dv = u
	maxD = -1
	f(dv, -1, 0)
	dw = u
	return
}

// 树的重心
// 以重心为根时，最大子树结点数最少
// 性质：
// - 一棵树最多有两个重心，且相邻
// - 树中所有点到某个点的距离和中，到重心的距离和是最小的；如果有两个重心，那么距离和一样
// - 把两棵树通过一条边相连得到一棵新的树，新重心在旧重心的路径上
// - 在一棵树上添加或删除一个叶结点后，重心保持不变或移动至相邻的结点上
// 常用作点分治中的一个划分步骤
// https://oi-wiki.org/graph/tree-centroid/
func (*tree) findCentroid(n int, g [][]int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minMaxSubSize := int(1e9)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		size := 1
		maxSubSize := 0
		for _, w := range g[v] {
			if w != fa {
				sz := f(w, v)
				size += sz
				maxSubSize = max(maxSubSize, sz)
			}
		}
		maxSubSize = max(maxSubSize, n-size) // 向上的子树大小
		if maxSubSize < minMaxSubSize {
			minMaxSubSize = maxSubSize
			ans = v
		}
		return size
	}
	return
}

// 点分治
// 例：求树上距离不超过 upperDis 的点对数 http://poj.org/problem?id=1741
// TODO: 需要重新整理
func (*tree) numPairsWithDistanceLimit(in io.Reader, n int, upperDis int64) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type neighbor struct {
		to     int
		weight int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < n-1; i++ {
		var v, w int
		var weight int64
		Fscan(in, &v, &w, &weight)
		v--
		w--
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}
	usedCentroid := make([]bool, n)

	subtreeSize := make([]int, n)
	var calcSubtreeSize func(v, fa int) int
	calcSubtreeSize = func(v, fa int) int {
		sz := 1
		for _, e := range g[v] {
			if w := e.to; w != fa && !usedCentroid[w] {
				sz += calcSubtreeSize(w, v)
			}
		}
		subtreeSize[v] = sz
		return sz
	}

	var findCentroid func(v, fa, compSize int) (int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct int) {
		minSize = int(1e9)
		maxSubSize := 0
		sizeV := 1 // 除去了 usedCentroid 子树的剩余大小
		for _, e := range g[v] {
			if w := e.to; w != fa && !usedCentroid[w] {
				if minSizeW, ctW := findCentroid(w, v, compSize); minSizeW < minSize {
					minSize = minSizeW
					ct = ctW
				}
				maxSubSize = max(maxSubSize, subtreeSize[w])
				sizeV += subtreeSize[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-sizeV)
		if maxSubSize < minSize {
			minSize = maxSubSize
			ct = v
		}
		return
	}

	var disToCentroid []int64
	var calcDisToCentroid func(v, fa int, d int64)
	calcDisToCentroid = func(v, fa int, d int64) {
		disToCentroid = append(disToCentroid, d)
		for _, e := range g[v] {
			if w := e.to; w != fa && !usedCentroid[w] {
				calcDisToCentroid(w, v, d+e.weight)
			}
		}
	}

	countPairs := func(ds []int64) int64 {
		cnt := int64(0)
		//sort.Ints(ds)
		sort.Slice(ds, func(i, j int) bool { return ds[i] < ds[j] })
		j := len(ds)
		for i, di := range ds {
			for ; j > 0 && di+ds[j-1] > upperDis; j-- {
			}
			cnt += int64(j)
			if j > i {
				cnt--
			}
		}
		return cnt >> 1
	}

	var f func(int) int64
	f = func(v int) (ans int64) {
		calcSubtreeSize(v, -1)
		_, ct := findCentroid(v, -1, subtreeSize[v])
		usedCentroid[ct] = true
		// 统计按 ct 分割后的子树中的点对数
		for _, e := range g[ct] {
			if !usedCentroid[e.to] {
				ans += f(e.to)
			}
		}
		// 统计经过 ct 的点对数
		// 0 是方便统计包含 ct 的部分
		ds := []int64{0}
		for _, e := range g[ct] {
			if !usedCentroid[e.to] {
				disToCentroid = []int64{}
				calcDisToCentroid(e.to, ct, e.weight)
				ans -= countPairs(disToCentroid)
				ds = append(ds, disToCentroid...)
			}
		}
		ans += countPairs(ds)
		usedCentroid[ct] = false
		return
	}
	return f(0)
}

// 最近公共祖先 - 其一 - 基于倍增和二分搜索
// 适用于查询量和节点数等同的情形
// NOTE: 多个点的 LCA 等于 dfn_min 和 dfn_max 的 LCA
// https://oi-wiki.org/graph/lca/#_5
// TODO 模板题 https://www.luogu.com.cn/problem/P3379
// 题目推荐 https://cp-algorithms.com/graph/lca.html#toc-tgt-2
// TODO log 优化
func (*tree) lcaBinarySearch(n, root int, g [][]int) {
	const mx = 18 // floor(log2(最大树节点))+1
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var dfs func(v, p, d int)
	dfs = func(v, p, d int) {
		pa[v][0] = p
		dep[v] = d
		for _, w := range g[v] {
			if w != p {
				dfs(w, v, d+1)
			}
		}
	}
	dfs(root, -1, 0)
	for k := 0; k+1 < mx; k++ {
		for v := range pa {
			if p := pa[v][k]; p != -1 {
				pa[v][k+1] = pa[p][k]
			} else {
				pa[v][k+1] = -1
			}
		}
	}
	_lca := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for k := 0; k < mx; k++ {
			if (dep[w]-dep[v])>>k&1 == 1 {
				w = pa[w][k]
			}
		}
		if v == w {
			return v
		}
		for k := mx - 1; k >= 0; k-- {
			if pa[v][k] != pa[w][k] {
				v, w = pa[v][k], pa[w][k]
			}
		}
		return pa[v][0]
	}
	_d := func(v, w int) int { return dep[v] + dep[w] - dep[_lca(v, w)]<<1 }

	// EXTRA: 其他树上二分
	var dist []int // 预处理略
	// 二分搜索 dist(x,v) <= d 的离根最近的 x
	search := func(v int, d int) int {
		dv := dist[v]
		for i := mx - 1; i >= 0; i-- {
			if p := pa[v][i]; p != -1 && dv-dist[p] <= d {
				v = p
			}
		}
		return v
	}

	_ = []interface{}{_d, search}
}

// 最近公共祖先 - 其二 - 基于 RMQ
// 由于预处理 ST 表是基于一个长度为 2n 的序列，所以常数上是比倍增算法要大的。内存占用也比倍增要大一倍左右（这点可忽略）
// 优点是查询的复杂度低，适用于查询量大的情形
// https://oi-wiki.org/graph/lca/#rmq
func (*tree) lcaRMQ(n, root int, g [][]int) {
	vs := make([]int, 0, 2*n-1)  // 欧拉序列
	pos := make([]int, n)        // pos[v] 表示 v 在 vs 中第一次出现的位置编号
	dep := make([]int, 0, 2*n-1) // 深度序列，和欧拉序列一一对应
	dis := make([]int, n)        // dis[v] 表示 v 到 root 的距离
	var dfs func(v, p, d int)    // 若有边权需额外传参 dis
	dfs = func(v, p, d int) {
		pos[v] = len(vs)
		vs = append(vs, v)
		dep = append(dep, d)
		dis[v] = d
		for _, w := range g[v] {
			if w != p {
				dfs(w, v, d+1) // 若有边权则额外传入 dis+e.weight
				vs = append(vs, v)
				dep = append(dep, d)
			}
		}
	}
	dfs(root, -1, 0)
	type pair struct{ v, i int }
	const mx = 18 // floor(log2(最大树节点))+1
	var st [][mx]pair
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]pair, n)
		for i, v := range a {
			st[i][0] = pair{v, i}
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				if a, b := st[i][j-1], st[i+1<<(j-1)][j-1]; a.v < b.v {
					st[i][j] = a
				} else {
					st[i][j] = b
				}
			}
		}
	}
	stInit(dep)
	stQuery := func(l, r int) int { // [l,r) 注意 l r 是从 0 开始算的
		k := bits.Len(uint(r-l)) - 1
		a, b := st[l][k], st[r-1<<k][k]
		if a.v < b.v {
			return a.i
		}
		return b.i
	}
	// 注意下标的换算，打印 LCA 的话要 +1
	_lca := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return vs[stQuery(pv, pw+1)]
	}
	_d := func(v, w int) int { return dis[v] + dis[w] - dis[_lca(v, w)]<<1 }

	_ = _d
}

// LCA 应用：树上差分
// 操作为更新 v-w 路径上的点权或边权（初始为 0）
// 点权时 diff[lca] -= val
// 边权时 diff[lca] -= 2 * val（定义 diff 为点到父亲的差分值）
func (*tree) differenceOnTree(n, root int, g [][]int) {
	diff := make([]int, n)
	update := func(v, w int, val int) {
		var lca int // = _lca(v, w)
		diff[v] += val
		diff[w] += val
		diff[lca] -= val
	}

	// 自底向上求出每个点的点权
	ans := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		sum := diff[v]
		for _, w := range g[v] {
			if w != fa {
				sum += f(w, v)
			}
		}
		ans[v] = sum
		return sum
	}

	_ = update
}

// 树链剖分（重链剖分）
// https://en.wikipedia.org/wiki/Heavy_path_decomposition
// https://oi-wiki.org/graph/hld/
// https://cp-algorithms.com/graph/hld.html
// vals 为点权
// 模板题（点权）https://www.luogu.com.cn/problem/P3384
// TODO: 处理边权的情况
func (*tree) hld(n, root int, g [][]int, vals []int64) {
	// 重儿子，父节点，深度，子树大小，所处重链顶点（深度最小），DFS 序（作为线段树中的编号，从 1 开始）
	type node struct{ hson, fa, depth, size, top, dfn int }
	nodes := make([]node, n)
	//idv := make([]int, n+1) // idv[nodes[v].dfn] == v

	var build func(v, fa, d int) *node
	build = func(v, fa, d int) *node {
		nodes[v] = node{hson: -1, fa: fa, depth: d, size: 1}
		o := &nodes[v]
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			son := build(w, v, d+1)
			o.size += son.size
			if o.hson == -1 || son.size > nodes[o.hson].size {
				o.hson = w
			}
		}
		return o
	}
	build(root, -1, 0)

	dfn := 0
	var decomposition func(v, fa, top int)
	decomposition = func(v, fa, top int) {
		o := &nodes[v]
		o.top = top
		dfn++
		o.dfn = dfn
		//idv[dfn] = v
		if o.hson != -1 {
			// 优先遍历重儿子，保证在同一条重链上的点的 DFS 序是连续的
			decomposition(o.hson, v, top)
			for _, w := range g[v] {
				if w != fa && w != o.hson {
					decomposition(w, v, w)
				}
			}
		}
	}
	decomposition(root, -1, root)

	t := make(lazySegmentTree, 4*n)
	// 按照 DFS 序对应的点权初始化线段树
	dfnVals := make([]int64, n)
	for i, v := range vals {
		dfnVals[nodes[i].dfn-1] = v
	}
	t.init(dfnVals)

	doPath := func(v, w int, do func(l, r int)) {
		ov, ow := nodes[v], nodes[w]
		for ; ov.top != ow.top; ov, ow = nodes[v], nodes[w] {
			topv, topw := nodes[ov.top], nodes[ow.top]
			// v 所处的重链顶点必须比 w 的深
			if topv.depth < topw.depth {
				v, w = w, v
				ov, ow = ow, ov
				topv, topw = topw, topv
			}
			do(topv.dfn, ov.dfn)
			// TODO: 边权下，处理轻边的情况
			v = topv.fa
		}
		if ov.depth > ow.depth {
			//v, w = w, v
			ov, ow = ow, ov
		}
		do(ov.dfn, ow.dfn)
		// TODO: 边权下，处理轻边的情况
	}
	updatePath := func(v, w int, add int64) { doPath(v, w, func(l, r int) { t.update(l, r, add) }) }
	queryPath := func(v, w int) (sum int64) { doPath(v, w, func(l, r int) { sum += t.query(l, r) }); return } // % mod
	updateSubtree := func(v int, add int64) { o := nodes[v]; t.update(o.dfn, o.dfn+o.size-1, add) }
	querySubtree := func(v int) (sum int64) { o := nodes[v]; return t.query(o.dfn, o.dfn+o.size-1) }

	_ = []interface{}{updatePath, queryPath, updateSubtree, querySubtree}
}

// TODO: Morris Traversal
// https://www.cnblogs.com/anniekim/archive/2013/06/15/morristraversal.html

// TODO: link/cut tree
// https://en.wikipedia.org/wiki/Link/cut_tree
// https://oi-wiki.org/ds/lct/
// 模板题 https://www.luogu.com.cn/problem/P3690

// TODO: 树上启发式合并 DSU on tree
// https://oi-wiki.org/graph/dsu-on-tree/
// todo https://ac.nowcoder.com/discuss/394080 E 题

// TODO: 虚树
// https://oi-wiki.org/graph/virtual-tree/

// 基环树
// TODO

// 仙人掌图 Cactus graph
// A connected graph in which any two simple cycles have at most one vertex in common
// https://en.wikipedia.org/wiki/Cactus_graph
// TODO 圆方树
// 模板题 https://www.luogu.com.cn/problem/P5236
