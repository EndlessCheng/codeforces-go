package copypasta

import (
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

/*
注：这里的代码偏向于数据结构，其余的树上统计类算法见 dp.go 中的树形 DP 部分

从特殊到一般：先思考一条链的情况，然后逐渐增加分支来思考一般的情况

NOTE: 对于有根树的题，可以考虑加上 g[0] = append(g[0], -1) 来简化代码
NOTE: 由于树上任意两点间的路径等价于两条点到根的路径的对称差，处理一些树上异或的问题可以往这个方向思考
NOTE: 注意特判只有一条边的情况，此时两个叶结点对应同一条边
NOTE: 一些树上点对问题，可以从「每条边所能产生的贡献」来思考 https://codeforces.com/problemset/problem/700/B

树上统计（从下往上）典型题 https://codeforces.com/problemset/problem/766/E
不错的构造 https://codeforces.com/problemset/problem/260/D
分类讨论的好题 https://codeforces.com/problemset/problem/765/E
*/

// namespace
type tree struct{}

// DFS: 树上两点路径
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

// 两个基本信息：节点深度和子树大小
// 节点深度：
// - 深度与祖先：v 是 w 的祖先，当且仅当 dep[v]+dist(v,w)=dep[w]
// - 与 DFS 序结合，可以表达子树在某个深度上的一段信息（见 tree.inOutTimestamp）
// - 直径 中心（见 tree.diameter）
// 子树大小：
// - 与 DFS 序结合，把子树转化成区间（见 tree.subtreeSize）
// - 重心 点分治（见 tree.findCentroid 等）
// - 重链剖分（见 tree.heavyLightDecomposition）
// - 用于计算每条边对所有路径产生的贡献 https://codeforces.com/problemset/problem/1401/D
//
// 离线好题 https://codeforces.com/problemset/problem/570/D
// 这题的在线写法是把相同深度的 dfn 放入同一组（同组内的 dfn 是有序的），对于一颗子树的某个深度，在该组中必对应着连续的一段 dfn，二分即可找到
func (*tree) depthSize(n, root int, g [][]int, max func(int, int) int, v int) {
	dep := make([]int, n)
	size := make([]int, n)
	maxDep := make([]int, n) // EXTRA: 子树最大深度
	var build func(v, fa, d int) int
	build = func(v, fa, d int) int {
		dep[v] = d
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += build(w, v, d+1)
				maxDep[v] = max(maxDep[v], maxDep[w]+1)
			}
		}
		size[v] = sz
		return sz
	}
	build(root, -1, 0)

	// EXTRA: 一种贪心策略是，将 g[v] 按照 maxDep 从大到小排序
	// https://codeforces.com/contest/1510/submission/111986751
	sort.Slice(g[v], func(i, j int) bool { return maxDep[g[v][i]] > maxDep[g[v][j]] })
}

// 树上每个子树的信息：子树大小，DFS 序（从 1 开始）
// 这样的话 [o.dfn, o.dfn+o.size-1] 就表示一颗子树，方便用树状数组/线段树维护
// 模板题 https://ac.nowcoder.com/acm/contest/6383/B
// 例题 https://codeforces.com/problemset/problem/383/C
//     https://codeforces.com/problemset/problem/877/E
func (*tree) subtreeSize(n, root int, g [][]int) {
	type node struct{ dfn, size int }
	nodes := make([]node, n)
	dfn := 0
	var build func(v, fa int) int
	build = func(v, fa int) int {
		dfn++
		nodes[v].dfn = dfn
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += build(w, v)
			}
		}
		nodes[v].size = sz
		return sz
	}
	build(root, -1)

	isFa := func(fa, v int) bool { return nodes[fa].dfn < nodes[v].dfn && nodes[v].dfn < nodes[fa].dfn+nodes[fa].size }

	{
		dfnToNodeID := make([]int, n+1)
		for i, o := range nodes {
			dfnToNodeID[o.dfn] = i
		}
	}

	{
		var v int
		var update, query func(int, int)
		var queryOne func(int)

		// 注意 o.dfn 从 1 开始
		o := nodes[v]
		update(o.dfn, o.dfn+o.size-1) // 更新子树
		query(o.dfn, o.dfn+o.size-1)  // 查询子树
		queryOne(nodes[v].dfn)        // 查询单个节点
	}

	_ = isFa
}

// 每个节点的入出时间戳
// 应用：可以 O(1) 判断 fa 是否为 v 的祖先节点（是否在根到 v 的路径上）
// 例题 https://codeforces.com/problemset/problem/1328/E
// 好题（需要充分利用入出时间戳的性质）https://codeforces.com/problemset/problem/1528/C
// 给定一颗 n 个点的完全 k 叉树的先序遍历，还原这棵树 https://ac.nowcoder.com/acm/contest/9247/B
//    先用 BFS 建树，然后 DFS 跑建好的树
//    也可以不用 BFS，根据完全 k 叉树的性质直接建图：（点的范围从 0 到 n-1）
//    for w := 1; w < n; w++ {
//	      v := (w - 1) / k
//        g[v] = append(g[v], w)
//    }
// 其他：见 mo.go 中的树上莫队部分
func (*tree) inOutTimestamp(n, root int, g [][]int) {
	timeIn := make([]int, n)
	timeOut := make([]int, n)
	at := make([]int, n+1)
	clock := 0
	var f func(v, fa int)
	f = func(v, fa int) {
		clock++
		timeIn[v] = clock
		at[clock] = v
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
		timeOut[v] = clock
	}
	f(root, -1)
	isPa := func(pa, v int) bool { return timeIn[pa] < timeIn[v] && timeIn[v] <= timeOut[pa] }

	{
		// 与深度时间戳结合，二分求某个子树在某个深度的节点范围
		// https://codeforces.com/problemset/problem/208/E 加强版 https://www.luogu.com.cn/problem/P5384（需要差分）
		// https://codeforces.com/problemset/problem/246/E
		// https://codeforces.com/problemset/problem/570/D
		// https://codeforces.com/problemset/problem/1076/E
		type info struct{ tin, tout, dep int }
		is := make([]info, n)
		depT := make([][]int, n)
		t := 0
		var f func(v, fa, d int)
		f = func(v, fa, d int) {
			t++
			is[v].tin = t
			is[v].dep = d
			depT[d] = append(depT[d], t)
			for _, w := range g[v] {
				if w != fa {
					f(w, v, d+1)
				}
			}
			is[v].tout = t
		}
		f(root, -1, 0)

		// 深度 d 上的这一排节点与子树 v 求交集，返回对应的深度 d 的节点区间 [l,r)
		query := func(v, d int) (int, int) {
			nf := is[v]
			//d += nf.dep // 如果 d 是从 v 开始算的话还要加上节点在整棵树的深度
			l := sort.SearchInts(depT[d], nf.tin)
			r := sort.SearchInts(depT[d], nf.tout+1)
			return l, r
		}
		_ = query
	}

	_ = isPa
}

// 树上最小路径覆盖，要求路径之间不相交，即每个顶点恰好被覆盖一次（路径长度可以为 0，即一个点）
// 贪心做法是对每个点及其子树，尽量让这个点能够拼接两个子树的路径，从而形成一条路径（把这个点当作「拐点」）
// DP 做法是定义 dp[i][0/1] 表示以 i 为根的子树的最小路径覆盖数，0 即 i 不与父节点构成路径（当作拐点），1 即 i 与父节点构成路径（不当作拐点）
// - DP 写法见 https://blog.csdn.net/clove_unique/article/details/52724434
// https://en.wikipedia.org/wiki/Path_cover
// https://codeforces.com/problemset/problem/618/D
func (*tree) minPathCover(g [][]int) int {
	// 考虑最小路径覆盖中，在树上的边数
	// 显然每多一条路径覆盖，在树上的边数就少一条
	// 因此有：最小路径覆盖数=n-在树上的最大边数
	edgeNum := 0
	var f func(int, int) bool
	f = func(v, fa int) bool {
		c := 0
		for _, w := range g[v] {
			if w != fa && f(w, v) {
				c++
			}
		}
		if c < 2 {
			edgeNum += c
			return true // 可以与父节点合并
		}
		edgeNum += 2
		return false // 不与父节点合并
	}
	f(0, -1)
	return len(g) - edgeNum
}

// 树的直径/最长链（DP 求法另见 dp.go 中的 diameter）
// 返回树的某条直径的两端点以及直径长度（最长链长度）
// 树的中心：树的直径的中点。直径长度为偶数时有一个，为奇数时有两个
//    如果给每条边加一个中点，那么可以保证树的中心为一个
// 性质：
//    直径的中点到所有叶子的距离和最小
//    对于两棵树，记第一棵树直径两端点为 u 和 v，第二棵树直径两端点为 x 和 y。若用一条边连接两棵树，则新树存在某条直径，其两端点一定是 u,v,x,y 中的两个点
// 随机树的直径 https://zhuanlan.zhihu.com/p/398621082
// 树的直径与重心（含动态维护） https://www.luogu.com.cn/blog/Loveti/problem-tree
// https://leetcode-cn.com/contest/biweekly-contest-12/problems/tree-diameter/
// EXTRA: 森林的情况 https://codeforces.com/problemset/problem/455/C
// 转换的好题 https://codeforces.com/problemset/problem/734/E
// 转换成求部分直径 https://codeforces.com/problemset/problem/1617/E https://oeis.org/A072339
// 必须边 https://www.luogu.com.cn/problem/P3304 https://www.acwing.com/problem/content/description/391/
func (*tree) diameter(st int, g [][]int) (int, int, int) {
	maxD, u := -1, 0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1) // d+e.wt
			}
		}
	}
	f(st, -1, 0)
	dv := u
	maxD = -1
	f(u, -1, 0)
	dw := u

	// EXTRA: 获取所有直径端点
	// 任意直径端点都可以从其中的一条直径的端点之一出发获得
	// 找出直径 dv-dw 后，对两个端点各自跑一次 DFS，即可找出全部直径端点
	// 注：这种写法加上上面，总共用了四次 DFS，可以仅需两次，见 https://codeforces.com/contest/592/submission/119472149
	// 下标最小的直径端点 https://codeforces.com/problemset/problem/592/D
	// 树上非严格次长距离 https://ac.nowcoder.com/acm/contest/9557/C（另一种做法见下面的 secondDiameter）
	isEnd := make([]bool, len(g))
	var findEnds func(v, fa, d int)
	findEnds = func(v, fa, d int) {
		if d == maxD {
			isEnd[v] = true
			return
		}
		for _, w := range g[v] {
			if w != fa {
				findEnds(w, v, d+1)
			}
		}
	}
	findEnds(dv, -1, 0)
	findEnds(dw, -1, 0)
	ends := []int{}
	for v, e := range isEnd {
		if e {
			ends = append(ends, v)
		}
	}

	// EXTRA: 获取所有在直径上的点
	// https://ac.nowcoder.com/acm/contest/9753/C
	onDiameter := make([]bool, len(g))
	var findVerticesOnDiameter func(v, fa, d int) bool
	findVerticesOnDiameter = func(v, fa, d int) bool {
		if d == maxD {
			onDiameter[v] = true
			return true
		}
		for _, w := range g[v] {
			if w != fa {
				if findVerticesOnDiameter(w, v, d+1) {
					onDiameter[v] = true
				}
			}
		}
		return onDiameter[v]
	}
	findVerticesOnDiameter(dv, -1, 0)
	findVerticesOnDiameter(dw, -1, 0)
	dvs := []int{}
	for v, on := range onDiameter {
		if on {
			dvs = append(dvs, v)
		}
	}

	// EXTRA: 获取直径上的所有节点 path
	// path[len(path)/2] 即为树的中心（之一）
	path := []int{}
	var findDiameterPath func(v, fa int) bool
	findDiameterPath = func(v, fa int) bool {
		if v == dw {
			path = append(path, v)
			return true
		}
		for _, w := range g[v] {
			if w != fa && findDiameterPath(w, v) {
				path = append(path, v)
				return true
			}
		}
		return false
	}
	findDiameterPath(dv, -1)

	// EXTRA: 求出无根树上每个点的最远点及距离（紫书 p.282 思考题）
	// 从任意直径的两个端点出发跑 DFS，取最大值
	// 相关题目 https://codeforces.com/problemset/problem/337/D
	farthest := make([]struct{ v, d int }, len(g))
	for i := range farthest {
		farthest[i].d = -1
	}
	var cur int
	var findFarthest func(v, fa, d int)
	findFarthest = func(v, fa, d int) {
		if d > farthest[v].d {
			farthest[v].d = d
			farthest[v].v = cur
		}
		for _, w := range g[v] {
			if w != fa {
				findFarthest(w, v, d+1)
			}
		}
	}
	cur = dv
	findFarthest(dv, -1, 0)
	cur = dw
	findFarthest(dw, -1, 0)

	return dv, dw, maxD
}

// 非严格次长直径
// https://ac.nowcoder.com/acm/contest/9557/C
func (*tree) secondDiameter(st int, g [][]int) int {
	n := len(g)
	maxD, u, cntD := -1, 0, make([]int, n)
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		cntD[d]++
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	maxD = -1
	f(st, -1, 0)

	// 从直径的两端点 p q 出发求深度列表，最大的两个为直径 p-q 和 q-p，倒数第三个为非严格次长直径
	// 这里用基数排序
	maxD, cntD = -1, make([]int, n)
	f(u, -1, 0)
	f(u, -1, 0)
	if cntD[maxD] > 2 {
		return maxD
	}
	return maxD - 1
}

// 树的重心
// 性质：
//    以重心为根时，最大子树结点数最少，且所有子树的大小都不超过 节点数/2
//        反之，若存在一颗子树其大小超过 节点数/2，则重心在该子树中
//    一棵树最多有两个重心，且相邻
//    拥有奇数个节点的树只有一个重心
//    树中所有点到某个点的距离和中，到重心的距离和是最小的；如果有两个重心，那么距离和一样
//    把两棵树通过一条边相连得到一棵新的树，新重心在旧重心的路径上
//    在一棵树上添加或删除一个叶结点后，重心保持不变或移动至相邻的结点上（换句话说，重心个数可能一个变两个，两个变一个，或者说至多移动半条边）
//    树的重心一定在根节点的重链上 https://www.luogu.com.cn/problem/P5666
//    树的重心一定在它重儿子的重心到根节点的路径上 https://www.luogu.com.cn/problem/P5666
// 常用作点分治中的一个划分步骤
// https://oi-wiki.org/graph/tree-centroid/
// 树的直径与重心（含动态维护） https://www.luogu.com.cn/blog/Loveti/problem-tree
// 树重心的性质及动态维护 https://www.cnblogs.com/qlky/p/5781081.html
// 求两个重心 https://codeforces.com/problemset/problem/1406/C
// Edge replacement 后哪些点可以是重心 https://codeforces.com/problemset/problem/708/C
func (*tree) findCentroid(n, st int, g [][]int) (ct int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minMaxSubSize := int(1e9)
	var findCt func(v, fa int) int
	findCt = func(v, fa int) int {
		size := 1
		maxSubSize := 0
		for _, w := range g[v] {
			if w != fa {
				sz := findCt(w, v)
				size += sz
				maxSubSize = max(maxSubSize, sz)
			}
		}
		maxSubSize = max(maxSubSize, n-size) // 向上的子树大小
		if maxSubSize < minMaxSubSize {
			minMaxSubSize = maxSubSize
			ct = v
		}
		return size
	}
	findCt(st, -1)
	return
}

// 点分治 - 重心分解（CD, Centroid Decomposition）
// https://oi-wiki.org/graph/tree-divide/
// todo https://zhuanlan.zhihu.com/p/359209926
// https://codeforces.com/blog/entry/81661
// todo 点分治略解 https://www.luogu.com.cn/blog/user9012/dian-fen-zhi-lve-xie
// todo 模板题 https://www.luogu.com.cn/problem/P3806
// 模板题 https://codeforces.com/problemset/problem/321/C
// todo poj1741 poj2114 uva12161 spoj QTREE5
// 好题 https://codeforces.com/contest/1174/problem/F https://codeforces.com/contest/1174/submission/82371930
// todo ∑∑min(av,aw)*dis(v,w) https://ac.nowcoder.com/acm/contest/11171/D
func (*tree) centroidDecomposition(n, root int, g [][]int) {
	type node struct{ dep, fa int }
	nodes := make([]node, n)
	var build func(v, fa, d int)
	build = func(v, fa, d int) {
		nodes[v] = node{d, fa}
		for _, w := range g[v] {
			if w != fa {
				build(w, v, d+1)
			}
		}
	}
	build(root, -1, 0)

	usedCentroid := make([]bool, n)
	size := make([]int, n)
	var calcSize func(v, fa int) int
	calcSize = func(v, fa int) int {
		sz := 1
		for _, w := range g[v] {
			if w != fa && !usedCentroid[w] {
				sz += calcSize(w, v)
			}
		}
		size[v] = sz
		return sz
	}
	var compSize int
	var findCentroid func(v, fa int) int
	findCentroid = func(v, fa int) int {
		for _, w := range g[v] {
			if w != fa && !usedCentroid[w] && size[w] > compSize>>1 {
				return findCentroid(w, v)
			}
		}
		return v
	}

	var f func(v int)
	f = func(v int) {
		calcSize(v, -1)
		compSize = size[v]
		ct := findCentroid(v, -1)
		usedCentroid[ct] = true
		//defer func() { usedCentroid[ct] = false }()

		// do ct...

		for _, w := range g[ct] {
			if !usedCentroid[w] {
				f(w)
			}
		}
	}
	f(root)
}

// 点分治：求树上距离不超过 upperDis 的点对数
// todo 待整理 https://www.luogu.com.cn/problem/P4178 http://poj.org/problem?id=1741
// todo 求树上距离等于 k 的点对数 https://codeforces.com/problemset/problem/161/D 可以参考洛谷的代码
func (*tree) numPairsWithDistanceLimit(in io.Reader, n, root int, upperDis int64) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type neighbor struct {
		to int
		wt int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < n-1; i++ {
		var v, w int
		var wt int64
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}
	usedCentroid := make([]bool, n)

	size := make([]int, n)
	var calcSize func(v, fa int) int
	calcSize = func(v, fa int) int {
		sz := 1
		for _, e := range g[v] {
			if w := e.to; w != fa && !usedCentroid[w] {
				sz += calcSize(w, v)
			}
		}
		size[v] = sz
		return sz
	}

	var compSize int
	var findCentroid func(v, fa int) (int, int)
	findCentroid = func(v, fa int) (minSize, ct int) {
		minSize = int(1e9)
		maxSubSize := 0
		sizeV := 1 // 除去了 usedCentroid 子树的剩余大小
		for _, e := range g[v] {
			if w := e.to; w != fa && !usedCentroid[w] {
				if minSizeW, ctW := findCentroid(w, v); minSizeW < minSize {
					minSize = minSizeW
					ct = ctW
				}
				maxSubSize = max(maxSubSize, size[w])
				sizeV += size[w]
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
				calcDisToCentroid(w, v, d+e.wt)
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

	var f func(v, fa int) int64
	f = func(v, fa int) (ans int64) {
		calcSize(v, fa)
		compSize = size[v]
		_, ct := findCentroid(v, fa)
		usedCentroid[ct] = true
		defer func() { usedCentroid[ct] = false }()

		// 统计按 ct 分割后的子树中的点对数
		for _, e := range g[ct] {
			if w := e.to; !usedCentroid[w] {
				ans += f(w, v)
			}
		}

		// 统计经过 ct 的点对数
		// 0 是方便统计包含 ct 的部分
		ds := []int64{0}
		for _, e := range g[ct] {
			if w := e.to; !usedCentroid[w] {
				disToCentroid = []int64{}
				calcDisToCentroid(w, ct, e.wt)
				ans -= countPairs(disToCentroid)
				ds = append(ds, disToCentroid...)
			}
		}
		ans += countPairs(ds)
		return
	}
	return f(root, -1)
}

// 点分树（动态点分治）
// todo https://oi-wiki.org/graph/dynamic-tree-divide/
// todo 模板题 https://www.luogu.com.cn/problem/P6329

// 最近公共祖先 · 其一 · 基于树上倍增和二分搜索
// O(nlogn) 预处理，O(logn) 查询
// 适用于查询量和节点数等同的情形
// NOTE: 多个点的 LCA 等于 dfn_min 和 dfn_max 的 LCA
// https://oi-wiki.org/graph/lca/#_5
// 模板题 https://www.luogu.com.cn/problem/P3379
// https://atcoder.jp/contests/arc060/tasks/arc060_c
// https://codeforces.com/problemset/problem/33/D
// 路径点权乘积 https://ac.nowcoder.com/acm/contest/6913/C
// 树上倍增应用（静态路径查询）：代码见下面的 EXTRA 部分
//    维护最大值（与 MST 结合）https://codeforces.com/problemset/problem/609/E
//       变体 https://codeforces.com/problemset/problem/733/F
//    维护最大值（与 MST 结合）LC1697 https://leetcode-cn.com/problems/checking-existence-of-edge-length-limited-paths/
//    维护最大值（与 MST 结合）LC1724（上面这题的在线版）https://leetcode-cn.com/problems/checking-existence-of-edge-length-limited-paths-ii/
//    维护最大值和严格次大值（严格次小 MST）：见 graph.go 中的 strictlySecondMST
//    维护前十大（点权）https://codeforces.com/problemset/problem/587/C
// 树上倍增-查询深度最小的未被标记的点 https://codeforces.com/problemset/problem/980/E
// 题目推荐 https://cp-algorithms.com/graph/lca.html#toc-tgt-2
// todo poj2763 poj1986 poj3728
// 其他：见 mo.go 中的树上莫队部分
func (*tree) lcaBinarySearch(n, root int, g [][]int) {
	const mx = 17 // bits.Len(最大节点数)
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var build func(v, p, d int)
	build = func(v, p, d int) {
		pa[v][0] = p
		dep[v] = d
		for _, w := range g[v] {
			if w != p {
				build(w, v, d+1)
			}
		}
	}
	build(root, -1, 0)
	// 倍增
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	// 从 v 开始向上跳 k 步，不存在返回 -1
	// O(1) 求法见长链剖分
	uptoKthPa := func(v, k int) int {
		for i := 0; i < mx && v != -1; i++ {
			if k>>i&1 > 0 {
				v = pa[v][i]
			}
		}
		return v
	}
	// 从 v 开始向上跳到指定深度 d，d<=dep[v]
	// https://en.wikipedia.org/wiki/Level_ancestor_problem
	// https://codeforces.com/problemset/problem/1535/E
	uptoDep := func(v, d int) int {
		for i := 0; i < mx; i++ {
			if (dep[v]-d)>>i&1 > 0 {
				v = pa[v][i]
				//if v == -1 { panic(-9) }
			}
		}
		return v
	}
	_lca := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}
	disVW := func(v, w int) int { return dep[v] + dep[w] - dep[_lca(v, w)]<<1 }

	// EXTRA: 输入 u 和 v，u 是 v 的祖先，返回 u 到 v 路径上的第二个节点
	down := func(u, v int) int {
		// assert dep[u] < dep[v]
		v = uptoDep(v, dep[u]+1)
		if pa[v][0] == u {
			return v
		}
		return -1
	}

	{
		// 加权树上二分
		var dep []int64 // 加权深度，dfs 预处理略
		// 从 v 开始向根移动至多 d 距离，返回最大移动次数，以及能移动到的离根最近的点
		// 变形 https://codeforces.com/problemset/problem/932/D
		uptoDep := func(v int, d int64) (int, int) {
			step := 0
			dv := dep[v]
			for i := mx - 1; i >= 0; i-- {
				if p := pa[v][i]; p != -1 && dv-dep[p] <= d {
					step |= 1 << i
					v = p
				}
			}
			return step, v
		}
		_ = uptoDep
	}

	{
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		// EXTRA: 倍增的时候维护其他属性，如边权最值等
		// 下面的代码来自 https://codeforces.com/problemset/problem/609/E
		// EXTRA: 额外维护最值边的下标，见 https://codeforces.com/contest/733/submission/120955685
		type nb struct{ to, wt int }
		g := make([][]nb, n)
		// read g ...
		const mx = 18
		type pair struct{ p, maxWt int }
		pa := make([][mx]pair, n)
		dep := make([]int, n)
		var build func(v, p, d int)
		build = func(v, p, d int) {
			pa[v][0].p = p
			dep[v] = d
			for _, e := range g[v] {
				if w := e.to; w != p {
					pa[w][0].maxWt = e.wt
					build(w, v, d+1)
				}
			}
		}
		build(0, -1, 0)

		for i := 0; i+1 < mx; i++ {
			for v := range pa {
				if p := pa[v][i]; p.p != -1 {
					pp := pa[p.p][i]
					pa[v][i+1] = pair{pp.p, max(p.maxWt, pp.maxWt)}
				} else {
					pa[v][i+1].p = -1
				}
			}
		}

		// 求 LCA(v,w) 的同时，顺带求出 v-w 上的边权最值
		_lca := func(v, w int) (lca, maxWt int) {
			if dep[v] > dep[w] {
				v, w = w, v
			}
			for i := 0; i < mx; i++ {
				if (dep[w]-dep[v])>>i&1 > 0 {
					p := pa[w][i]
					maxWt = max(maxWt, p.maxWt)
					w = p.p
				}
			}
			if w != v {
				for i := mx - 1; i >= 0; i-- {
					if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
						maxWt = max(maxWt, max(pv.maxWt, pw.maxWt))
						v, w = pv.p, pw.p
					}
				}
				maxWt = max(maxWt, max(pa[v][0].maxWt, pa[w][0].maxWt))
				v = pa[v][0].p
			}
			// 如果是点权的话这里加上 maxWt = max(maxWt, pa[v][0].maxWt)
			lca = v
			return
		}

		_ = _lca
	}

	_ = []interface{}{disVW, uptoKthPa, down}
}

// 最近公共祖先 · 其二 · 基于 RMQ
// O(nlogn) 预处理，O(1) 查询
// 由于预处理 ST 表是基于一个长度为 2n 的序列，所以常数上是比倍增算法要大的。内存占用也比倍增要大一倍左右（这点可忽略）
// 优点是查询的复杂度低，适用于查询量大的情形
// https://oi-wiki.org/graph/lca/#rmq
func (*tree) lcaRMQ(n, root int, g [][]int) {
	vs := make([]int, 0, 2*n-1)  // 欧拉序列
	pos := make([]int, n)        // pos[v] 表示 v 在 vs 中第一次出现的位置编号
	dep := make([]int, 0, 2*n-1) // 深度序列，和欧拉序列一一对应
	disRoot := make([]int, n)    // disRoot[v] 表示 v 到 root 的距离
	var build func(v, p, d int)  // 若有边权需额外传参 dis
	build = func(v, p, d int) {
		pos[v] = len(vs)
		vs = append(vs, v)
		dep = append(dep, d)
		disRoot[v] = d
		for _, w := range g[v] {
			if w != p {
				build(w, v, d+1) // d+e.wt
				vs = append(vs, v)
				dep = append(dep, d)
			}
		}
	}
	build(root, -1, 0)
	type pair struct{ v, i int }
	const mx = 17 // bits.Len(最大节点数)
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
	_d := func(v, w int) int { return disRoot[v] + disRoot[w] - disRoot[_lca(v, w)]<<1 }

	_ = _d
}

// 最近公共祖先 · 其三 · Tarjan 离线算法
// 时间和空间复杂度均为 O(n+q)
// 虽然用了并查集但是由于数据的特殊性，操作的均摊结果是 O(1) 的，见 https://core.ac.uk/download/pdf/82125836.pdf
// https://oi-wiki.org/graph/lca/#tarjan
// https://cp-algorithms.com/graph/lca_tarjan.html
// 扩展：Tarjan RMQ https://codeforces.com/blog/entry/48994
func (*tree) lcaTarjan(in io.Reader, n, q, root int) []int {
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		v, w := 0, 0
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	pa := make([]int, n)
	for i := range pa {
		pa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
	}

	lca := make([]int, q)
	dis := make([]int, q) // dis(q.v,q.w)
	type query struct{ w, i int }
	qs := make([][]query, n)
	for i := 0; i < q; i++ {
		v, w := 0, 0
		Fscan(in, &v, &w)
		v--
		w--
		if v != w {
			qs[v] = append(qs[v], query{w, i})
			qs[w] = append(qs[w], query{v, i})
		} else {
			// do v==w...
			lca[i] = v
			dis[i] = 0
		}
	}

	dep := make([]int, n)
	vis := make([]int8, n)
	var _f func(v, d int)
	_f = func(v, d int) {
		dep[v] = d
		vis[v] = 1
		for _, w := range g[v] {
			if vis[w] == 0 {
				_f(w, d+1)
				pa[w] = v
			}
		}
		for _, q := range qs[v] {
			if w := q.w; vis[w] == 2 {
				// do(v, w, lcaVW)...
				lcaVW := find(w)
				lca[q.i] = lcaVW
				dis[q.i] = dep[v] + dep[w] - dep[lcaVW]<<1
			}
		}
		vis[v] = 2
	}
	_f(root, 0)
	return lca
}

// LCA 应用：树上差分
// 操作为更新 v-w 路径上的点权或边权（初始为 0）
// 点权时 diff[lca] -= val
// 边权时 diff[lca] -= 2 * val（定义 diff 为点到父亲的差分值）
// https://www.luogu.com.cn/blog/RPdreamer/ci-fen-and-shu-shang-ci-fen
// todo https://loj.ac/d/1698
// 模板题（边权）https://codeforces.com/problemset/problem/191/C
func (*tree) differenceInTree(in io.Reader, n, root int, g [][]int) []int {
	_lca := func(v, w int) (_ int) { return }

	diff := make([]int, n)
	update := func(v, w int, val int) {
		lca := _lca(v, w)
		diff[v] += val
		diff[w] += val
		diff[lca] -= val // 点权
		//diff[lca] -= 2 * val // 边权
	}
	var q int
	Fscan(in, &q)
	for i := 0; i < q; i++ {
		var v, w, val int
		Fscan(in, &v, &w, &val)
		v--
		w--
		update(v, w, val)
	}

	// 自底向上求出每个点的点权/边权
	ans := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		sum := diff[v]
		for _, w := range g[v] {
			if w != fa {
				// 边权的话在这里记录 ans
				//s := f(w, v)
				//ans[e.eid] = s
				//sum += s
				sum += f(w, v)
			}
		}
		// 点权的话在这里记录 ans
		ans[v] = sum
		return sum
	}
	f(root, -1)

	return ans
}

// 树链剖分/重链剖分 (HLD, Heavy Light Decomposition）
// 性质：
//    树上每个结点都属于且仅属于一条重链
//    如果 v-w 是一条轻边，那么 size[w] < size[v]/2
//    从根结点到任意结点所经过的重链数为 O(logn)，轻边数为 O(logn)
//    与重心的关系（见 findCentroid）
// https://en.wikipedia.org/wiki/Heavy_path_decomposition
//    把每条重链当成一个节点，每条轻边作为边，我们可以得到一颗路径树。显然路径树的高度为 O(logn)
//    The paths of the decomposition may themselves be organized into a tree called the "path tree", "heavy path tree", or "compressed tree".
//    Each node of the path tree corresponds to a path of the heavy path decomposition.
//    If p is a path of the heavy path decomposition, then the parent of p in the path tree is the path containing the parent of the head of p.
//    The root of the path tree is the path containing the root of the original tree.
//    Alternatively, the path tree may be formed from the original tree by edge contraction of all the heavy edges.
// https://oi-wiki.org/graph/hld/
// https://cp-algorithms.com/graph/hld.html
// https://codeforces.com/blog/entry/81317
// 树链剖分详解 https://www.cnblogs.com/zwfymqz/p/8094500.html
// 树链剖分详解 https://www.luogu.com.cn/blog/communist/shu-lian-pou-fen-yang-xie
//
// 注：若没有修改操作，见 lcaBinarySearch（路径查询）以及 subtreeSize（子树查询）
//
// 模板题（点权）https://www.luogu.com.cn/problem/P3384
// 与最小生成树结合（边权）https://codeforces.com/problemset/problem/609/E
// 好题 https://codeforces.com/contest/1174/problem/F
// 归并树 https://codeforces.com/problemset/problem/587/C
// todo 子异和 https://www.luogu.com.cn/problem/P5127
// todo 完成题单 https://www.luogu.com.cn/training/1654
// TODO: 处理边权的情况
// todo NOI21 轻重边 https://www.luogu.com.cn/problem/P7735
func (*tree) heavyLightDecomposition(n, root int, g [][]int, vals []int64) { // vals 为点权
	// 深度，子树大小，重儿子，父节点，所处重链顶点（深度最小），DFS 序（作为线段树中的编号，从 1 开始）
	type node struct{ depth, size, hson, fa, top, dfn int }
	nodes := make([]node, n)
	//idv := make([]int, n+1) // idv[nodes[v].dfn] == v

	var build func(v, fa, dep int) int
	build = func(v, fa, dep int) int {
		size, hsz, hson := 1, 0, -1
		for _, w := range g[v] {
			if w != fa {
				sz := build(w, v, dep+1)
				size += sz
				if sz > hsz {
					hsz, hson = sz, w
				}
			}
		}
		nodes[v] = node{depth: dep, size: size, hson: hson, fa: fa}
		return size
	}
	build(root, -1, 0)

	{
		// EXTRA: 寻找以 st 为重链顶点的重链
		// hPath[-1] 即为重链末端节点
		getHP := func(st int) []int {
			hPath := []int{st}
			for o := nodes[st]; o.hson != -1; o = nodes[o.hson] {
				hPath = append(hPath, o.hson)
			}
			return hPath
		}

		_ = getHP
	}

	dfn := 0
	var markTop func(v, top int)
	markTop = func(v, top int) {
		o := &nodes[v]
		o.top = top
		dfn++
		o.dfn = dfn
		//idv[dfn] = v
		if o.hson != -1 {
			// 优先遍历重儿子，保证在同一条重链上的点的 DFS 序是连续的
			markTop(o.hson, top)
			for _, w := range g[v] {
				if w != o.fa && w != o.hson {
					markTop(w, w)
				}
			}
		}
	}
	markTop(root, root)

	// 按照 DFS 序对应的点权初始化线段树
	dfnVals := make([]int64, n)
	for i, v := range vals {
		dfnVals[nodes[i].dfn-1] = v
	}
	t := newLazySegmentTree(dfnVals)

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
	updatePath := func(v, w int, add int64) { doPath(v, w, func(l, r int) { t.update(1, l, r, add) }) }
	queryPath := func(v, w int) (sum int64) { doPath(v, w, func(l, r int) { sum += t.query(1, l, r) }); return } // % mod
	updateSubtree := func(v int, add int64) { o := nodes[v]; t.update(1, o.dfn, o.dfn+o.size-1, add) }
	querySubtree := func(v int) (sum int64) { o := nodes[v]; return t.query(1, o.dfn, o.dfn+o.size-1) }

	_ = []interface{}{updatePath, queryPath, updateSubtree, querySubtree}
}

// 长链剖分
// 长链剖分和重链剖分一样，是把一棵树分成若干条不相交的链
// 但是，这里的重儿子不再是子树大小最大的，而是深度最大的子节点（长儿子）
// 根据这一定义可推出，从根结点到任意结点所经过的轻边数为 O(√n) (想象一颗长儿子不断递减的二叉树)
// https://oi-wiki.org/graph/hld/#_14
// https://www.luogu.com.cn/blog/Ynoi/zhang-lian-pou-fen-xue-xi-bi-ji
// https://www.cnblogs.com/cj-chd/p/10076199.html
// https://www.cnblogs.com/zhoushuyu/p/9468669.html
// 应用：树上 k 级祖先 https://www.luogu.com.cn/problem/P5903 https://codeforces.com/problemset/problem/208/E
// 长链剖分优化树形 DP：
//    若树形 DP 的转移只和节点深度有关，我们完全可以把一颗子树拍扁成一条垂直的链
//    那么在合并子树时，长儿子将会占据主导优势，即其余子树均往长儿子上合并，这会使每个节点至多被合并一次，从而得到 O(n) 的优秀复杂度
//    具体实现时还有一些技巧，见后面的有关 DP 优化的 EXTRA
// 子树深度众数 https://codeforces.com/problemset/problem/1009/F
// 距离等于 k 的点对数 https://codeforces.com/problemset/problem/161/D
func (*tree) heavyLightDecompositionByDepth(n, root int, g [][]int) {
	// 深度，子树最大深度，重儿子，父节点，所处长链顶点（深度最小）
	type node struct{ depth, maxDepth, hson, fa, top int }
	nodes := make([]node, n)

	var build func(v, fa, dep int) int
	build = func(v, fa, dep int) int {
		maxDep, hson := dep, -1
		for _, w := range g[v] {
			if w != fa {
				if mxD := build(w, v, dep+1); mxD > maxDep {
					maxDep, hson = mxD, w
				}
			}
		}
		nodes[v] = node{depth: dep, maxDepth: maxDep, hson: hson, fa: fa}
		return maxDep
	}
	build(root, -1, 0) // 为了方便，fa 指定为 -1

	var markTop func(v, top int)
	markTop = func(v, top int) {
		o := &nodes[v]
		o.top = top
		if o.hson != -1 {
			markTop(o.hson, top)
			for _, w := range g[v] {
				if w != o.fa && w != o.hson {
					markTop(w, w)
				}
			}
		}
	}
	markTop(root, root)

	{
		// EXTRA: 对每个长链顶点 o，向上向下各记录长度为该长链长度的节点列表（此代码可以整合进 markTop）
		up := make([][]int, n)
		down := make([][]int, n)
		for v, o := range nodes {
			if o.top != v {
				continue
			}
			l := o.maxDepth - o.depth + 1
			up[v] = make([]int, 0, l) // 可能填不满
			for fa := v; fa != -1 && len(up[v]) < l; fa = nodes[fa].fa {
				up[v] = append(up[v], fa)
			}
			down[v] = make([]int, 0, l)
			for hson := v; hson != -1; hson = nodes[hson].hson {
				down[v] = append(down[v], hson)
			}
		}
		// 结合 up down，可以在倍增预处理后，O(1) 求出树上 k 级祖先
		const mx = 17 // bits.Len(最大节点数)
		pa := make([][mx]int, n)
		for i, o := range nodes {
			pa[i][0] = o.fa
		}
		for i := 0; i+1 < mx; i++ {
			for v := range pa {
				if p := pa[v][i]; p != -1 {
					pa[v][i+1] = pa[p][i]
				} else {
					pa[v][i+1] = -1
				}
			}
		}
		uptoKthPa := func(v, k int) int {
			if k == 0 {
				return v
			}
			lk := bits.Len(uint(k)) - 1
			v = pa[v][lk]
			top := nodes[v].top
			k ^= 1 << lk
			k -= nodes[v].depth - nodes[top].depth
			if k >= 0 {
				return up[top][k]
			}
			return down[top][-k]
		}
		_ = uptoKthPa
	}

	{
		// EXTRA: 对一些 DP 的优化
		hson := make([]int, n)
		var build func(v, fa int) int
		build = func(v, fa int) int {
			maxDep, hs := 0, -1
			for _, w := range g[v] {
				if w != fa {
					if mxD := build(w, v); mxD > maxDep {
						maxDep, hs = mxD, w
					}
				}
			}
			hson[v] = hs
			return maxDep + 1
		}
		build(root, -1)

		// 不妨定义当前子树最大深度为子子树最大深度+1，也就是越往下深度越小
		// 这样带来的好处是在合并之后，我们可以将当前节点的信息直接添加在 DP 数组末尾
		// 下面的代码来自 https://codeforces.com/problemset/problem/1009/F
		ans := make([]int, n)
		var f func(v, fa int) ([]int, int)
		f = func(v, fa int) (cnt []int, maxI int) { // maxI 为 cnt 中最大最右元素的下标
			if hson[v] == -1 {
				return []int{1}, 0
			}
			cnt, maxI = f(hson[v], v)
			for _, w := range g[v] {
				if w != fa && w != hson[v] {
					subCnt, _ := f(w, v)
					// do...
					shift := len(cnt) - len(subCnt)
					for i, c := range subCnt {
						i += shift
						if cnt[i] += c; cnt[i] >= cnt[maxI] {
							maxI = i
						}
					}
				}
			}
			cnt = append(cnt, 1)
			if cnt[maxI] == 1 {
				maxI++ // 即 len(cnt)-1
			}
			ans[v] = len(cnt) - 1 - maxI // 转化成题目要求的定义
			return
		}
		f(root, -1)
	}
}

// 树上启发式合并 DSU on tree / small to large
// https://oi-wiki.org/graph/dsu-on-tree/
// NOTE: 合并的时候最好先循环计算一遍答案，再循环合并一遍，这样的习惯可避免产生把合并之后的数值算入答案中的 bug
// 讲解+套题 https://pzy.blog.luogu.org/dsu-on-tree-xue-xi-bi-ji
// 讲解+套题 https://codeforces.com/blog/entry/44351 补充 https://codeforces.com/blog/entry/67696
// 模板题 https://www.luogu.com.cn/problem/U41492
//       https://codeforces.com/problemset/problem/600/E
// 距离等于 k 的点对数 https://codeforces.com/problemset/problem/161/D
//            变形题 https://ac.nowcoder.com/acm/contest/4853/E 题解 https://ac.nowcoder.com/discuss/394080
// todo https://ac.nowcoder.com/acm/contest/4010/E
func (*tree) dsu(n, root int, g [][]int, vals []int) { // vals 为点权
	hson := make([]int, n)
	var build func(v, fa int) int
	build = func(v, fa int) int {
		sz, hsz, hs := 1, 0, -1
		for _, w := range g[v] {
			if w != fa {
				s := build(w, v)
				sz += s
				if s > hsz {
					hsz, hs = s, w
				}
			}
		}
		hson[v] = hs
		return sz
	}
	build(root, -1)

	// 例如：统计子树的点权种类数
	ans := make([]int, n) // int64
	var f func(v, fa int) map[int]bool
	f = func(v, fa int) map[int]bool {
		if hson[v] < 0 {
			ans[v] = 1
			return map[int]bool{vals[v]: true}
		}
		has := f(hson[v], v)
		merge := func(val int) {
			// do...
			has[val] = true
		}
		for _, w := range g[v] {
			if w != fa && w != hson[v] {
				mp := f(w, v)
				for val := range mp {
					merge(val)
				}
			}
		}
		merge(vals[v])
		ans[v] = len(has)
		return has
	}
	f(root, -1)
}

// 树分块
// https://oi-wiki.org/ds/tree-decompose/
// https://ouuan.github.io/post/%E8%8E%AB%E9%98%9F%E5%B8%A6%E4%BF%AE%E8%8E%AB%E9%98%9F%E6%A0%91%E4%B8%8A%E8%8E%AB%E9%98%9F%E8%AF%A6%E8%A7%A3/#%E5%88%86%E5%9D%97%E6%96%B9%E5%BC%8F
// https://yfscfs.gitee.io/post/%E5%90%84%E7%A7%8D%E6%A0%91%E4%B8%8A%E5%88%86%E5%9D%97%E5%A7%BF%E5%8A%BF%E7%9A%84%E7%89%B9%E7%82%B9%E5%8F%8A%E5%B0%8F%E7%BB%93/
// https://www.cnblogs.com/hua-dong/p/8275227.html
// https://www.luogu.com.cn/blog/gxy001/shu-fen-kuai-xue-xi-bi-ji
// 模板题 王室联邦 https://www.luogu.com.cn/problem/P2325
func (*tree) limitSizeDecomposition(n, blockSize int, g [][]int) {
	// 树分块有多种方式，下面的代码基于《王室联邦》这题
	rt := []int{}         // 每个块的根节点
	bid := make([]int, n) // bid[i] 表示节点 i 所属的块的编号 rt[bid[i]] 即为 i 所属块的根节点
	s := []int{}
	var f func(v, fa int)
	f = func(v, fa int) {
		sz := len(s)
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
				if len(s)-sz >= blockSize {
					rt = append(rt, v)
					for len(s) > sz {
						bid[s[len(s)-1]] = len(rt) - 1
						s = s[:len(s)-1]
					}
				}
			}
		}
		s = append(s, v)
	}
	f(0, -1)
	if len(rt) == 0 {
		rt = []int{0}
	}
	for _, v := range s {
		bid[v] = len(rt) - 1
	}
}

// TODO: 虚树 Virtual Tree / Auxiliary Tree
// https://oi-wiki.org/graph/virtual-tree/
// https://www.luogu.com.cn/problem/P5891 https://class.luogu.com.cn/classroom/lgr66

// 普吕弗序列（Prufer 序列，Prüfer sequence）
// https://en.wikipedia.org/wiki/Pr%C3%BCfer_sequence
// https://oeis.org/A000272 Cayley's formula https://en.wikipedia.org/wiki/Cayley%27s_formula
// https://www.luogu.com.cn/problem/P6086
// todo 光之大陆 https://www.acwing.com/problem/content/2420/
func (tree) treeToPrufer(n int, pa []int) []int { // 传入的 pa 是以 n 为根时的每个节点的父节点
	deg := make([]int, n+1)
	for i := 1; i < n; i++ {
		deg[pa[i]]++
	}
	prufer := make([]int, n-2)
	for i, j := 0, 1; i < n-2; j++ {
		for deg[j] > 0 {
			j++
		}
		prufer[i] = pa[j]
		for i++; i < n-2; i++ {
			p := prufer[i-1]
			if deg[p]--; deg[p] > 0 || p > j {
				break
			}
			prufer[i] = pa[p]
		}
	}
	return prufer
}

func (tree) pruferToTree(n int, prufer []int) []int {
	deg := make([]int, n+1)
	for _, p := range prufer {
		deg[p]++
	}
	prufer = append(prufer, n) // 设置 prufer[n-2] = n 方便后续使用
	pa := make([]int, n+1)
	for i, j := 0, 1; i < n-1; j++ {
		for deg[j] > 0 {
			j++
		}
		pa[j] = prufer[i]
		for ; i < n-2; i++ {
			p := prufer[i]
			if deg[p]--; deg[p] > 0 || p > j {
				break
			}
			pa[p] = prufer[i+1]
		}
		i++
	}
	pa[n] = -1
	return pa
}
