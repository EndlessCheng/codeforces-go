package copypasta

import (
	"maps"
	"math"
	"math/bits"
	"reflect"
	"slices"
	"sort"
	"unsafe"
)

/*
注：这里的代码偏向于数据结构，其余的树上统计类算法见 dp.go 中的树形 DP 部分

从特殊到一般：先思考一条链的情况，然后逐渐增加分支来思考一般的情况

NOTE: 对于有根树的题，可以考虑加上 g[0] = append(g[0], -1) 来简化代码
NOTE: 由于树上任意两点间的路径等价于两条点到根的路径的对称差，处理一些树上异或的问题可以往这个方向思考
NOTE: 注意特判整棵树只有一条边的情况，此时两个叶结点对应同一条边
NOTE: 一些树上点对问题，可以从「每条边所能产生的贡献」来思考 https://codeforces.com/problemset/problem/700/B
NOTE: 节点数小于 √n 的同层节点对不超过 n√n，节点数大于 √n 的层的数量小于 √n 个 https://codeforces.com/problemset/problem/1806/E
NOTE: 树上两点的关系：v 和 w 相等【特判】、v 是 w 的祖先、w 是 v 的祖先、其它（v 和 w 在两棵不同子树中）https://codeforces.com/problemset/problem/1778/E
NOTE: 记录从 x 到根的路径上的每个点到 x 的距离，就可以从 y 走到根的路径上，找到到 x 的距离，从而求出 y 到 x 的距离 https://codeforces.com/problemset/problem/1790/F

随机树有期望 n/2 个叶子节点
On the number of leaves in a random recursive tree https://projecteuclid.org/journals/brazilian-journal-of-probability-and-statistics/volume-29/issue-4/On-the-number-of-leaves-in-a-random-recursive-tree/10.1214/14-BJPS252.pdf

简单 DFS
- [2368. 受限条件下可到达节点的数目](https://leetcode.cn/problems/reachable-nodes-with-restrictions/) 1477
- [3004. 相同颜色的最大子树](https://leetcode.cn/problems/maximum-subtree-of-the-same-color/)（会员题）
https://codeforces.com/problemset/problem/1675/D 1300 树分成尽量少的链
https://codeforces.com/problemset/problem/580/C 1500
https://codeforces.com/problemset/problem/34/D 1600

巧妙 DFS
https://atcoder.jp/contests/abc163/tasks/abc163_f 2470=CF2579 树上路径计数

BFS
https://codeforces.com/problemset/problem/2018/C 1700

利用递归栈快速标记祖先节点 https://codeforces.com/problemset/problem/1774/E
树上统计（从下往上）典型题 https://codeforces.com/problemset/problem/766/E
不错的构造 https://codeforces.com/problemset/problem/260/D
分类讨论的好题 https://codeforces.com/problemset/problem/765/E

树上贪心
https://codeforces.com/problemset/problem/1029/E 2100

树上路径异或
LC2791 https://leetcode.cn/problems/count-paths-that-can-form-a-palindrome-in-a-tree/
http://poj.org/problem?id=3764
https://www.luogu.com.cn/problem/UVA13277 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=878&page=show_problem&problem=5201

树上移动 move on tree
https://codeforces.com/problemset/problem/1774/E

https://en.wikipedia.org/wiki/Tree_traversal#Pre-order,_NLR
前序中序构造二叉树 + 判定是否合法 https://atcoder.jp/contests/abc255/tasks/abc255_f

树的最小表示：复杂度分析
https://leetcode.cn/problems/special-binary-string/solutions/1731760/on-log-n-by-hqztrue-nrmw/

其它
https://codeforces.com/problemset/problem/1491/E 2400
*/

// namespace
type tree struct{}

// 树哈希 Hashing root trees
// O(nlogn)
// https://codeforces.com/blog/entry/113465?#comment-1010870
// 判断是否为对称树（可以调整儿子顺序）https://codeforces.com/problemset/problem/1800/G
func (*tree) hash(g [][]int, root int) {
	tid := map[string]int{}
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		ids := make([]int, 0, len(g[v]))
		for _, w := range g[v] {
			if w != fa {
				ids = append(ids, dfs(w, v))
			}
		}
		slices.Sort(ids)
		// do ids...

		_ids := append(ids[:0:0], ids...) // 如果后面用不到 ids 可以去掉
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&_ids))
		sh.Len *= bits.UintSize / 8
		s := *(*string)(unsafe.Pointer(sh))
		id, ok := tid[s]
		if !ok {
			id = len(tid) // 从 0 开始
			tid[s] = id
		}
		return id
	}
	dfs(root, -1)
}

// 树同构
// AHU 算法
// https://oi-wiki.org/graph/tree-ahu/
// https://wwwmayr.in.tum.de/konferenzen/Jass08/courses/1/smal/Smal_Paper.pdf
// https://logic.pdmi.ras.ru/~smal/files/smal_jass08_slides.pdf
// todo https://www.zhihu.com/question/55484468/answer/991551284
// todo hashing 的一些正确姿势 https://zhuanlan.zhihu.com/p/104346215
//
// https://www.luogu.com.cn/problem/P5043
// 与换根 DP 结合：
// - https://codeforces.com/contest/763/problem/D
// - https://codeforces.com/problemset/problem/1794/E 参考代码 https://codeforces.com/contest/1794/submission/196015876
// https://open.kattis.com/problems/twochartsbecomeone

// https://codeforces.com/contest/342/problem/E
func (*tree) bfsMultiSources(g [][]int, starts []int) {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = 1e9
	}
	type pair struct{ v, fa int }
	q := []pair{}
	for _, v := range starts {
		q = append(q, pair{v, -1})
	}
	bfs := func(q []pair) {
		for _, p := range q {
			dis[p.v] = 0
		}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			v := p.v
			for _, w := range g[v] {
				if w != p.fa && dis[v]+1 < dis[w] {
					dis[w] = dis[v] + 1
					q = append(q, pair{w, v})
				}
			}
		}
	}
	bfs(q)
}

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

// 预处理从 v 到 w 走一步的节点 move1[v][w]
// 定义 v 到 v 走一步的节点为 v
// https://codeforces.com/problemset/problem/1771/D
func (*tree) move1(g [][]int) [][]int {
	move1 := make([][]int, len(g))
	for i := range move1 {
		move1[i] = make([]int, len(g))
	}
	for rt := range move1 {
		var build func(int, int)
		build = func(v, fa int) {
			move1[v][rt] = fa
			for _, w := range g[v] {
				if w != fa {
					build(w, v)
				}
			}
		}
		build(rt, rt)
	}
	return move1
}

// 两个基本信息：节点深度和子树大小
// 节点深度：
// - 深度与祖先：v 是 w 的祖先，当且仅当 dep[v]+dist(v,w)=dep[w]
// - 与 DFS 序结合，可以表达子树在某个深度上的一段信息（见 tree.inOutTimestamp）
// - 直径 中心（见 tree.diameter）
// 子树大小：
// - 与 DFS 序结合，把子树转化成区间（见 tree.dfnOrder）
// - 重心 点分治（见 tree.findCentroid 等）
// - 重链剖分（见 tree.heavyLightDecomposition）
// - 用于计算每条边对所有路径产生的贡献 https://codeforces.com/problemset/problem/1401/D
//
// 离线好题 https://codeforces.com/problemset/problem/570/D
// 这题的在线写法是把相同深度的 dfn 放入同一组（同组内的 dfn 是有序的），对于一棵子树的某个深度，在该组中必对应着连续的一段 dfn，二分即可找到
func (*tree) depthSize(n, root int, g [][]int, v int) {
	dep := make([]int, n)
	size := make([]int, n)
	maxDep := make([]int, n) // EXTRA: 子树最大深度
	var build func(int, int, int) int
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
	slices.SortFunc(g[v], func(a, b int) int { return maxDep[b] - maxDep[a] })
}

// DFS 序（从 0 开始）
// 另见后面的 inOutTimestamp
// 讲解：https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solution/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/
//
// LC3327 https://leetcode.cn/problems/check-if-dfs-strings-are-palindromes/
// https://codeforces.com/problemset/problem/1899/G 1900
// https://codeforces.com/problemset/problem/877/E 2000
// https://codeforces.com/problemset/problem/383/C 2000
// https://codeforces.com/problemset/problem/620/E 2100
// https://codeforces.com/problemset/problem/916/E 2400
// https://codeforces.com/problemset/problem/1110/F 2600
// https://codeforces.com/problemset/problem/163/E 2800 结合 AC 自动机 
// https://ac.nowcoder.com/acm/contest/6383/B
func (*tree) dfnOrder(root int, g [][]int, a []int) {
	n := len(a)
	newOrder := make([]int, n)
	// 闭区间 [l,r]   0 <= l <= r <= n-1
	nodes := make([]struct{ l, r int }, n)
	dfn := 0
	var buildDFN func(int, int)
	buildDFN = func(v, fa int) {
		nodes[v].l = dfn
		newOrder[dfn] = a[v] // 写这里是先序遍历，如果题目要求后序遍历，把这一行和下一行的 dfn++ 移到 for 循环后面
		dfn++
		for _, w := range g[v] {
			if w != fa {
				buildDFN(w, v)
			}
		}
		nodes[v].r = dfn - 1
	}
	buildDFN(root, -1)

	// 返回 [f 是 v 的祖先节点]
	// f == v 的情况请单独处理
	// LC2322 https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/ 2392
	// 判断给定点集是否都在一条路径上 https://codeforces.com/contest/1702/problem/G2 2000
	// https://codeforces.com/problemset/problem/1527/D 2400
	isAncestor := func(f, v int) bool { return nodes[f].l < nodes[v].l && nodes[v].l <= nodes[f].r }

	{
		dfnToNodeID := make([]int, n)
		for v, o := range nodes {
			dfnToNodeID[o.l] = v
		}
	}

	{
		// 如何使用？一般配合树状数组/线段树等数据结构
		var v int
		var update, query func(int, int)
		var queryOne func(int)

		o := nodes[v]
		update(o.l, o.r) // 更新子树（闭区间）
		query(o.l, o.r)  // 查询子树（闭区间）
		queryOne(o.l)    // 查询单个节点
	}

	{
		// 如果递归消耗太多内存，可以改为手动模拟栈
		// 下面的代码是有向树，不需要传入 fa
		// https://codeforces.com/contest/163/submission/233981400
		root := 0
		nodes := make([]struct{ l, r int }, n) // 左闭右开
		type stackInfo struct{ v, i int }
		st := []stackInfo{{root, 0}}
		nodes[root].l = 1
		dfn := 0
		for len(st) > 0 {
			p := st[len(st)-1]
			v, i := p.v, p.i
			if i < len(g[v]) {
				dfn++
				w := g[v][i]
				nodes[w].l = dfn
				st[len(st)-1].i++
				st = append(st, stackInfo{w, 0})
			} else {
				nodes[v].r = dfn + 1
				st = st[:len(st)-1]
			}
		}
	}

	_ = isAncestor
}

// 每个节点的入出时间戳
// 应用：可以 O(1) 判断 fa 是否为 v 的祖先节点（是否在根到 v 的路径上）
// 视频讲解 https://www.bilibili.com/video/BV1pW4y1r7xs/
// 文字讲解 https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solution/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/
// LC2322 https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/ 2392
// https://codeforces.com/problemset/problem/1328/E 1900 例题
// https://codeforces.com/problemset/problem/1528/C 2300 好题（需要充分利用入出时间戳的性质）
// https://codeforces.com/problemset/problem/2002/D2 2300
// https://codeforces.com/problemset/problem/1416/D 2600
// 给定一棵 n 个点的完全 k 叉树的先序遍历，还原这棵树 https://ac.nowcoder.com/acm/contest/9247/B
//    先用 BFS 建树，然后 DFS 跑建好的树
//    也可以不用 BFS，根据完全 k 叉树的性质直接建图：（点的范围从 0 到 n-1）
//    for w := 1; w < n; w++ {
//	      v := (w - 1) / k
//        g[v] = append(g[v], w)
//    }
// 其他：见 mo.go 中的树上莫队部分
func (*tree) inOutTimestamp(g [][]int, root int) {
	// DFS 序
	timeIn := make([]int, len(g))
	timeOut := make([]int, len(g))
	at := make([]int, len(g)+1)
	clock := 0 // 也可以改成从 -1 开始
	var build func(int, int)
	build = func(v, fa int) {
		clock++
		timeIn[v] = clock
		at[clock] = v
		for _, w := range g[v] {
			if w != fa {
				build(w, v)
			}
		}
		timeOut[v] = clock
	}
	build(root, -1) // 注意森林的情况

	// 返回 [f 是 v 的祖先节点]
	// f == v 的情况请单独处理
	// https://codeforces.com/problemset/problem/916/E 2400
	isAncestor := func(f, v int) bool { return timeIn[f] < timeIn[v] && timeIn[v] <= timeOut[f] }
	isAncestor2 := func(v, w int) bool { return isAncestor(v, w) || isAncestor(w, v) }

	{
		// 与深度时间戳结合，二分求某个子树在某个深度的节点范围
		// https://www.lanqiao.cn/problems/5892/learning/?contest_id=145
		// https://codeforces.com/problemset/problem/1076/E 1900
		// https://codeforces.com/problemset/problem/208/E 2100 
		// - https://www.luogu.com.cn/problem/P5384
		// - https://atcoder.jp/contests/abc202/tasks/abc202_e
		// https://codeforces.com/problemset/problem/570/D 2200
		// https://codeforces.com/problemset/problem/246/E 2400 CF208E + 离线
		// https://www.luogu.com.cn/problem/P7768
		var a []int // 点权   read... 
		type info struct{ tin, tout, dep int }
		nodes := make([]info, len(g))
		depTS := make([][]int, len(g))
		rowVal := make([][]int, len(g))
		ts := 0
		var f func(int, int, int)
		f = func(v, fa, d int) {
			ts++
			nodes[v].tin = ts
			nodes[v].dep = d
			depTS[d] = append(depTS[d], ts)
			rowVal[d] = append(rowVal[d], a[v])
			for _, w := range g[v] {
				if w != fa {
					f(w, v, d+1)
				}
			}
			nodes[v].tout = ts
		}
		f(root, -1, 0)

		// 返回子树 v 中的绝对深度为 d 的这一排节点在 depTS[d] 中的下标范围 [l,r)
		// 结合 RMQ 可以求出这一排节点的最大点权
		// d += nodes[v].dep // 相对深度：如果 d 是从 v 开始算的话，要加上节点在整棵树的深度
		query := func(v, d int) (int, int) {
			nf := nodes[v]
			l := sort.SearchInts(depTS[d], nf.tin)
			r := sort.SearchInts(depTS[d], nf.tout+1)
			return l, r
		}
		_ = query
	}

	_, _ = isAncestor, isAncestor2
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

// 树的直径/最长链       （树形 DP 做法另见 dp.go 中的 diameter）
// 返回树的某条直径的两端点以及直径长度（最长链长度）
// 树的中心：树的直径的中点。直径长度为偶数时有一个，为奇数时有两个
// - 如果给每条边加一个中点，那么可以保证树的中心为一个
// 性质：
// - 直径的中点到所有叶子的距离和最小
// - 对于两棵树，记第一棵树直径两端点为 u 和 v，第二棵树直径两端点为 x 和 y。若用一条边连接两棵树，则新树存在某条直径，其两端点一定是 u,v,x,y 中的两个点
//
// 为什么不能用类似找直径的做法求**图**的直径呢？比如做两次 BFS
// 反例：
// 1 2
// 1 3
// 2 4
// 3 5
// 2 5
// 从 1 出发跑 BFS，是可能会把 5 当成最远点的，然后从 5 出发跑 BFS 就算错了，正确的应该是从 4 出发跑 BFS
//
// 随机树的直径 https://zhuanlan.zhihu.com/p/398621082
// 树的直径与重心（含动态维护） https://www.luogu.com.cn/blog/Loveti/problem-tree
//
// LC1245 https://leetcode.cn/problems/tree-diameter/
// https://codeforces.com/problemset/problem/1404/B 1900
// https://codeforces.com/problemset/problem/455/C 2100 两棵树连边，连边之后直径最小
// - 简化版 LC3203 https://leetcode.cn/problems/find-minimum-diameter-after-merging-two-trees/
// https://codeforces.com/problemset/problem/734/E 2100 转换的好题 
// https://codeforces.com/problemset/problem/1000/E 2100 e-BCC
// https://codeforces.com/problemset/problem/379/F 2400
// https://codeforces.com/problemset/problem/911/F 2400 贪心
// https://codeforces.com/problemset/problem/1819/C 2400
// https://codeforces.com/problemset/problem/1617/E 2700 转换成求部分直径 
// - https://oeis.org/A072339
// https://www.luogu.com.cn/problem/P3304 必须边 
// https://www.luogu.com.cn/problem/T238762?contestId=65460 求树中任意一个与 x 距离为 k 的点 
// https://www.lanqiao.cn/problems/5890/learning/?contest_id=145
func (*tree) diameter(st int, g [][]int) (int, int, int) {
	maxD, u := -1, 0
	var findMaxDepth func(int, int, int)
	findMaxDepth = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				findMaxDepth(w, v, d+1) // d+e.wt
			}
		}
	}
	findMaxDepth(st, -1, 0)
	dv := u
	maxD = -1
	findMaxDepth(u, -1, 0)
	dw := u

	// EXTRA: 获取所有直径端点
	// 任意直径端点都可以从其中的一条直径的端点之一出发获得
	// 找出直径 dv-dw 后，对两个端点各自跑一次 DFS，即可找出全部直径端点
	// 注：这种写法加上上面，总共用了四次 DFS，可以仅需两次，见 https://codeforces.com/contest/592/submission/119472149
	// 下标最小的直径端点 https://codeforces.com/problemset/problem/592/D
	// 树上非严格次长距离 https://ac.nowcoder.com/acm/contest/9557/C（另一种做法见下面的 secondDiameter）
	isEnd := make([]bool, len(g))
	var findAllEnds func(v, fa, d int)
	findAllEnds = func(v, fa, d int) {
		if d == maxD {
			isEnd[v] = true
			return
		}
		for _, w := range g[v] {
			if w != fa {
				findAllEnds(w, v, d+1)
			}
		}
	}
	findAllEnds(dv, -1, 0)
	findAllEnds(dw, -1, 0)
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
	// https://codeforces.com/problemset/problem/1819/C
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
	// https://codeforces.com/problemset/problem/337/D 2000
	// https://codeforces.com/problemset/problem/911/F 2400
	// 每个点相距为 k 的点 https://atcoder.jp/contests/abc267/tasks/abc267_f
	// https://leetcode.cn/problems/find-the-last-marked-nodes-in-tree/
	farthest := make([]struct{ v, d int }, len(g))
	for i := range farthest {
		farthest[i].d = -1
	}
	var findFarthest func(int, int, int, int)
	findFarthest = func(v, fa, d, tar int) {
		if d > farthest[v].d {
			farthest[v].d = d
			farthest[v].v = tar
		}
		for _, w := range g[v] {
			if w != fa {
				findFarthest(w, v, d+1, tar)
			}
		}
	}
	findFarthest(dv, -1, 0, dv)
	findFarthest(dw, -1, 0, dw)

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
//    以重心为根时，最大子树结点数最少，且所有子树的大小都 < 节点数/2，或者说最大子树结点数 < 节点数/2
//        反之，若存在一棵子树其大小 ≥ 节点数/2，则重心在该子树中
//    一棵树最多有两个重心，且相邻
//    拥有奇数个节点的树只有一个重心
//    树中所有点到某个点的距离和中，到重心的距离和是最小的；如果有两个重心，那么距离和一样
//    把两棵树通过一条边相连得到一棵新的树，新重心在旧重心的路径上
//    在一棵树上添加或删除一个叶结点后，重心保持不变或移动至相邻的结点上（换句话说，重心个数可能一个变两个，两个变一个，或者说至多移动半条边）
//    树的重心一定在根节点的重链上 https://www.luogu.com.cn/problem/P5666
//    树的重心一定在它重儿子的重心到根节点的路径上 https://www.luogu.com.cn/problem/P5666
// 常用作点分治中的一个划分步骤
// https://oi-wiki.org/graph/tree-centroid/
// https://en.wikipedia.org/wiki/Tree_(graph_theory)#Properties
//    Every tree has a center consisting of one vertex or two adjacent vertices.
//    The center is the middle vertex or middle two vertices in every longest path.
//    Similarly, every n-vertex tree has a centroid consisting of one vertex or two adjacent vertices.
//    In the first case removal of the vertex splits the tree into subtrees of fewer than n/2 vertices.
//    In the second case, removal of the edge between the two centroidal vertices splits the tree into two subtrees of exactly n/2 vertices.
// 树的直径与重心（含动态维护） https://www.luogu.com.cn/blog/Loveti/problem-tree
// 树重心的性质及动态维护 https://www.cnblogs.com/qlky/p/5781081.html
// 求两个重心 https://codeforces.com/problemset/problem/1406/C
// 求每棵子树的重心 https://codeforces.com/problemset/problem/685/B
// Edge replacement 后哪些点可以是重心 https://codeforces.com/problemset/problem/708/C
// todo https://atcoder.jp/contests/abc362/tasks/abc362_f 重心性质
func (*tree) findCentroid(n, root int, g [][]int) (centroid int) {
	minOfMaxSubSize := math.MaxInt
	var findCt func(int, int) int
	findCt = func(v, fa int) int {
		size := 1
		maxSubSize := 0
		for _, w := range g[v] {
			if w != fa {
				sz := findCt(w, v)
				maxSubSize = max(maxSubSize, sz)
				size += sz
			}
		}
		maxSubSize = max(maxSubSize, n-size) // 向上的子树大小
		if maxSubSize < minOfMaxSubSize {
			minOfMaxSubSize = maxSubSize
			centroid = v
		}
		return size
	}
	findCt(root, -1)
	return
}

// 点分治 重心分解（CD, Centroid Decomposition）
// 适合处理树上路径相关问题
// 考察完经过某个重心的所有路径，后面就无需再考察这个重心了，直接将其删除
// 每次以重心为根递归处理，这样做递归深度不会超过 O(logn)
// 或者说，每个点至多被 O(log n) 个在重心拐弯的路径覆盖
// https://oi-wiki.org/graph/tree-divide/
// https://zhuanlan.zhihu.com/p/359209926
// https://codeforces.com/blog/entry/81661
// https://www.luogu.com.cn/blog/user9012/dian-fen-zhi-lve-xie
// https://liu-cheng-ao.blog.uoj.ac/blog/2969
// todo 重心树 代码 https://www.luogu.com.cn/record/103317317
//
// 模板题 https://www.luogu.com.cn/problem/P4178
// - http://poj.org/problem?id=1741
// todo 无需去重的做法（染色法）https://www.luogu.com.cn/blog/1239004072Angel/solution-p4178
// 多个询问 https://www.luogu.com.cn/problem/P3806
// - http://poj.org/problem?id=2114
// https://www.luogu.com.cn/problem/P4149
// 也可以树形 DP https://codeforces.com/problemset/problem/161/D 1800
// https://codeforces.com/problemset/problem/321/C 2100
// https://codeforces.com/problemset/problem/914/E 2400
// 好题 https://codeforces.com/contest/1174/problem/F 2400
// - https://codeforces.com/contest/1174/submission/82371930
// todo https://codeforces.com/contest/776/problem/F 2800
//  https://www.luogu.com.cn/problem/P2664
//  https://www.luogu.com.cn/problem/SP2939
//  ∑∑min(a[i],a[j])*dis(i,j) https://ac.nowcoder.com/acm/contest/11171/D
//  UVa12161 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=3313
// 距离相关
// - https://atcoder.jp/contests/yahoo-procon2018-final/tasks/yahoo_procon2018_final_c
// - https://leetcode.cn/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/ 
// -- https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/solutions/6103520/centroid-decomposition-o-n-log-n-m-log-m-solution/
// - https://judge.yosupo.jp/problem/vertex_add_range_contour_sum_on_tree

// 点分治 · 其一
// 遍历右边的子树 B，去看左边的子树 A（路径是 A 中的点 - 重心 - B 中的点）
func (tree) centroidDecomposition(g [][]struct{ to, wt int }) int {
	deleted := make([]bool, len(g))
	// 注：其实只需要保存 ct 的邻居的 size，但这并不好维护
	size := make([]int, len(g))
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, e := range g[v] {
			w := e.to
			if w != fa && !deleted[w] {
				minSizeW, ctW, faCtW := findCentroid(w, v, compSize)
				if minSizeW < minSize {
					minSize, ct, faCt = minSizeW, ctW, faCtW
				}
				maxSubSize = max(maxSubSize, size[w])
				size[v] += size[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-size[v])
		if maxSubSize < minSize {
			minSize, ct, faCt = maxSubSize, v, fa
		}
		return
	}

	ans := 0
	tmp := make([]int, len(g))
	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)
		//if size[v] == 1 {
		//	return
		//}

		pathValSet := map[int]bool{0: true} // 0 表示重心的数据
		for _, e := range g[ct] {
			w := e.to
			if deleted[w] {
				continue
			}
			tmp := tmp[:0]
			var f func(int, int, int)
			f = func(v, fa, pathVal int) {
				// do path & pathValSet like 2sum

				tmp = append(tmp, pathVal)
				for _, e := range g[v] {
					if w := e.to; w != fa && !deleted[w] {
						f(w, v, pathVal+e.wt)
					}
				}
			}
			f(w, ct, e.wt)
			// 注意在递归结束后，才把 tmp 的数据加入 pathValSet
			// 否则会把在同一棵子树内的数据当作另一棵子树的数据
			for _, pathVal := range tmp {
				pathValSet[pathVal] = true
			}
		}

		// 删除重心
		deleted[ct] = true
		for _, e := range g[ct] {
			w := e.to
			if !deleted[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}
	}
	dfs(0, -1, len(g))
	return ans
}

// 点分治 · 其二
// 适用场景：对每个点，计算和这个点有关的【路径信息】，比如距离、路径上的节点满足什么性质等
// 1. 从重心出发，统计整个重心连通块的数据
// 2. 对于重心的子树 A，先去掉子树 A 在第 1 步中统计的数据
// 3. 然后递归子树 A，对于子树 A 中的每个点 x，计算子树 A 中的点经过 x 到不在子树 A 的点的【路径信息】
// 4. 递归结束，恢复子树 A 的数据，处理重心的下一棵子树 B，回到第 2 步
// 5. 去掉整个重心连通块的数据，处理下一个重心连通块
// 下面的代码以 https://codeforces.com/problemset/problem/914/E 2400 为例
// - 对于每个顶点，输出经过该顶点的回文路径的数量
// todo https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/solutions/6103520/centroid-decomposition-o-n-log-n-m-log-m-solution/
func (tree) centroidDecomposition2(g [][]int, s string) []int {
	deleted := make([]bool, len(g))
	size := make([]int, len(g))
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				minSizeW, ctW, faCtW := findCentroid(w, v, compSize)
				if minSizeW < minSize {
					minSize, ct, faCt = minSizeW, ctW, faCtW
				}
				maxSubSize = max(maxSubSize, size[w])
				size[v] += size[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-size[v])
		if maxSubSize < minSize {
			minSize, ct, faCt = maxSubSize, v, fa
		}
		return
	}

	ans := make([]int, len(g))
	for i := range ans {
		ans[i] = 1
	}

	// 更新从 ct 出发的路径信息
	// delta = 1 或 -1
	cnt := [1 << 20]int{}
	var updateCC func(int, int, int, int)
	updateCC = func(v, fa, delta, pathMask int) {
		//pathMask++
		pathMask ^= 1 << (s[v] - 'a')
		cnt[pathMask] += delta
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				updateCC(w, v, delta, pathMask)
			}
		}
	}

	// （用两数之和的思路思考）
	// 计算「经过 v 向上，在重心拐弯，到其它子树」的路径信息
	// pathMask 从 ct 的儿子开始
	// 每个 v 至多被 O(log n) 个在重心拐弯的路径覆盖
	var calc func(int, int, int) int
	calc = func(v, fa, pathMask int) int {
		pathMask ^= 1 << (s[v] - 'a')
		// 单独计算：从 v 出发的路径信息
		res := cnt[pathMask]
		for i := 1; i < len(cnt); i <<= 1 {
			res += cnt[pathMask^i]
		}
		// 把 v 下面的也加上，这样最终算出的是经过 v 的路径信息
		for _, w := range g[v] {
			if w != fa && !deleted[w] {
				res += calc(w, v, pathMask)
			}
		}
		ans[v] += res
		return res
	}

	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)

		updateCC(ct, -1, 1, 0)
		// 单独计算：从 ct 出发的路径信息
		res := cnt[0]
		for i := 1; i < len(cnt); i <<= 1 {
			res += cnt[i]
		}
		// 再加上经过 ct 的路径信息
		for _, w := range g[ct] {
			if deleted[w] {
				continue
			}
			// 排除 w 子树后再计算
			updateCC(w, ct, -1, 1<<(s[ct]-'a'))
			res += calc(w, ct, 0)
			updateCC(w, ct, 1, 1<<(s[ct]-'a'))
		}
		// v->w 和 w->v 算了两次，同时去掉 [v] 这一个点的路径信息（注意初始化的时候 ans[i] = 1）
		ans[ct] += res / 2

		// 去掉整个重心连通块的路径信息
		updateCC(ct, -1, -1, 0)
		// 删除重心
		deleted[ct] = true

		// 处理其它重心连通块
		for _, w := range g[ct] {
			if !deleted[w] {
				if w != faCt {
					dfs(w, ct, size[w])
				} else {
					dfs(w, ct, compSize-size[ct])
				}
			}
		}
	}
	dfs(0, -1, len(g))
	return ans
}

// 动态点分治 点分树
// 维护 x 到 ct 的信息和 ct 到其余点的信息，就可以快速地处理 x 到其余所有点的信息（对于一个 x 来说，它的 ct 有 O(logn) 个）
// 适用于不关心树的形态的问题，比如路径问题，联通块问题，寻找关键点问题等等
// 提示：结合【贡献法】
// https://oi-wiki.org/graph/dynamic-tree-divide/
// https://oi-wiki.org/graph/tree-divide/#%E7%82%B9%E5%88%86%E6%A0%91
//
// 无修改 https://www.luogu.com.cn/problem/P3241
// 单点修改 https://www.luogu.com.cn/problem/P6329
// todo 点分树+堆 https://www.luogu.com.cn/problem/P2056 https://www.luogu.com.cn/problem/SP2666
//  借助点分树移动答案 https://www.luogu.com.cn/problem/P3345
//  动态加点的点分树+平衡树 https://www.luogu.com.cn/problem/P3920 
//  - 除去动态加点就是点分树套路。加点时默认新点的点分父亲为原树父亲，当某点分子树不平衡度超过某个阈值，重新点分治即可。
//  边分树+虚树 https://www.luogu.com.cn/problem/P4220
//  边分树+虚树 https://www.luogu.com.cn/problem/P4565
//  思维 | 最大深度最小的点分树 https://www.luogu.com.cn/problem/P5912
//  所有路径的点权异或和 + 单点修改 https://codeforces.com/gym/527120/problem/B CCPC 2024 上海市赛
func (*tree) centroidDecompositionTree(g [][]struct{ to, wt int }, root int, a []int) {
	deleted := make([]bool, len(g))
	size := make([]int, len(g))
	var findCentroid func(int, int, int) (int, int, int)
	findCentroid = func(v, fa, compSize int) (minSize, ct, faCt int) {
		minSize = math.MaxInt
		maxSubSize := 0
		size[v] = 1
		for _, e := range g[v] {
			w := e.to
			if w != fa && !deleted[w] {
				if minSizeW, ctW, faCtW := findCentroid(w, v, compSize); minSizeW < minSize {
					minSize, ct, faCt = minSizeW, ctW, faCtW
				}
				maxSubSize = max(maxSubSize, size[w])
				size[v] += size[w]
			}
		}
		maxSubSize = max(maxSubSize, compSize-size[v])
		if maxSubSize < minSize {
			minSize, ct, faCt = maxSubSize, v, fa
		}
		return
	}

	// paCts[x] 存储着 x 到其 ct 的信息（x 在 O(logn) 个重心连通块中）
	// paCts[x][0] 是最大的重心连通块，paCts[x][-1] 是最小的重心连通块
	// 注意：这个顺序不能表示与 x 的远近关系（可能 x 就在 paCts[x][0] 旁边，但离 paCts[x][1] 比较远）
	type disInfo struct{ ct, sonI, ctDis int }
	paCts := make([][]disInfo, len(g))

	// sonInfo[ct][i] 存储着 ct 到其重心连通块内的 g[ct][i] 这棵子树的其余点的信息
	sonInfo := make([][]fenwick, len(g))

	// mergeSonInfo[ct] 存储着 ct 到其重心连通块内的其余点的信息（不含 ct 这个点）
	mergeSonInfo := make([]fenwick, len(g))

	var dfs func(int, int, int)
	dfs = func(v, fa, compSize int) {
		_, ct, faCt := findCentroid(v, fa, compSize)

		sonInfo[ct] = make([]fenwick, len(g[ct]))
		totA := make(fenwick, compSize+1)
		for idx, e := range g[ct] {
			w := e.to
			if deleted[w] {
				continue
			}
			var sizeW int
			if w != faCt {
				sizeW = size[w]
			} else {
				sizeW = compSize - size[ct]
			}
			sumA := make(fenwick, sizeW+1)
			var f func(int, int, int)
			f = func(v, fa, d int) {
				paCts[v] = append(paCts[v], disInfo{ct, idx, d})
				sumA[d] += a[v]
				totA[d] += a[v]
				for _, e := range g[v] {
					w := e.to
					if w != fa && !deleted[w] {
						f(w, v, d+e.wt) // d+1
					}
				}
			}
			f(w, ct, e.wt) // 1
			for i := 1; i <= sizeW; i++ {
				if j := i + i&-i; j <= sizeW {
					sumA[j] += sumA[i]
				}
			}
			sonInfo[ct][idx] = sumA
		}
		for i := 1; i <= compSize; i++ {
			if j := i + i&-i; j <= compSize {
				totA[j] += totA[i]
			}
		}
		mergeSonInfo[ct] = totA

		deleted[ct] = true
		for _, e := range g[ct] {
			w := e.to
			if deleted[w] {
				continue
			}
			if w != faCt {
				dfs(w, ct, size[w])
			} else {
				dfs(w, ct, compSize-size[ct])
			}
		}
	}
	dfs(root, -1, len(g))

	update := func(x, val int) {
		delta := val - a[x]
		a[x] = val
		for _, p := range paCts[x] {
			sonInfo[p.ct][p.sonI].update(p.ctDis, delta)
			mergeSonInfo[p.ct].update(p.ctDis, delta)
		}
	}

	// https://www.luogu.com.cn/problem/P6329
	query := func(cur, k int) int {
		// 单独统计 cur
		res := a[cur]
		// 统计 cur 作为重心时，到它所在重心连通块的其余点的信息
		res += mergeSonInfo[cur].pre(min(k, len(mergeSonInfo[cur])-1))
		for _, p := range paCts[cur] {
			if p.ctDis > k {
				continue
			}
			// 单独统计 cur 与 p.ct 的信息
			res += a[p.ct]
			// 统计 p.ct 到重心连通块内的其余点的信息（个数），注意不能包含 cur 所在子树（否则就重复统计了）
			res += mergeSonInfo[p.ct].pre(min(k-p.ctDis, len(mergeSonInfo[p.ct])-1)) -
				sonInfo[p.ct][p.sonI].pre(min(k-p.ctDis, len(sonInfo[p.ct][p.sonI])-1))
		}
		return res
	}

	_ = []any{update, query}
}

// 边分治
// todo https://oi-wiki.org/graph/tree-divide/#%E8%BE%B9%E5%88%86%E6%B2%BB

// 最近公共祖先 · 其一 · 基于树上倍增
// 【模板讲解】树上倍增算法（以及最近公共祖先） 
// - 请看 https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/
// O(nlogn) 预处理，O(logn) 查询
// 适用于查询量和节点数等同的情形
// 适用于可以动态添加节点（挂叶子）的情形
// NOTE: 多个点的 LCA 等于 dfn_min 和 dfn_max 的 LCA
// https://oi-wiki.org/graph/lca/#_5
// 另见 mo.go 中的【树上莫队】
//
// 倍增 LC1483 https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/
// 模板题 https://www.luogu.com.cn/problem/P3379
// https://codeforces.com/problemset/problem/33/D 2000
// https://codeforces.com/problemset/problem/1304/E 2000
// https://codeforces.com/problemset/problem/1702/G2 2000 分类讨论
// - https://codeforces.com/problemset/problem/1527/D 2400 分类讨论（思路同 CF1702G2，但不需要倍增）
// https://atcoder.jp/contests/arc060/tasks/arc060_c 2154 倍增思想
// https://codeforces.com/problemset/problem/519/E 2100 到两点距离相同的点的数量
// https://codeforces.com/problemset/problem/1535/E 2200
// https://codeforces.com/problemset/problem/379/F 2400
// https://codeforces.com/problemset/problem/916/E 2400
// https://ac.nowcoder.com/acm/contest/6913/C 路径点权乘积 
//
// 维护元素和 LC2836 https://leetcode.cn/problems/maximize-value-of-function-in-a-ball-passing-game/
// 维护边权出现次数 LC2846 https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/
// 维护最大值（与 MST 结合）https://codeforces.com/problemset/problem/609/E 2100
//    变体 https://codeforces.com/problemset/problem/733/F 2200
// 维护最大值（与 MST 结合）LC1697 https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/
// 维护最大值（与 MST 结合）LC1724（上面这题的在线版）https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths-ii/
// 维护最大值和严格次大值（严格次小 MST）：见 graph.go 中的 strictlySecondMST
// 维护前十大（点权）https://codeforces.com/problemset/problem/587/C 2200
// 维护最大子段和 https://codeforces.com/contest/1843/problem/F2
// 维护从 x 往上有几个不同的 OR https://codeforces.com/contest/1878/problem/G
// 维护最大值 https://www.hackerearth.com/practice/algorithms/graphs/graph-representation/practice-problems/algorithm/optimal-connectivity-c6ae79ca/
// http://acm.hdu.edu.cn/showproblem.php?pid=7345
//
// 树上倍增-查询深度最小的未被标记的点 https://codeforces.com/problemset/problem/980/E 2200
// 题目推荐 https://cp-algorithms.com/graph/lca.html#toc-tgt-2
// todo poj2763 poj1986 poj3728
func (*tree) lcaBinaryLifting(root int, g [][]int) {
	const mx = 17 // bits.Len(最大节点数)
	pa := make([][mx]int, len(g))
	dep := make([]int, len(g)) // 根节点的深度为 0
	var buildPa func(int, int)
	buildPa = func(v, p int) {
		pa[v][0] = p
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				buildPa(w, v)
			}
		}
	}
	buildPa(root, -1) // pa[root][0] = -1
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	// 从 v 开始，向上跳到指定深度 d
	// https://en.wikipedia.org/wiki/Level_ancestor_problem
	// https://codeforces.com/problemset/problem/1535/E 2200
	uptoDep := func(v, d int) int {
		if d > dep[v] {
			panic(-1)
		}
		for k := uint(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
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
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}
	getDis := func(v, w int) int { return dep[v] + dep[w] - dep[getLCA(v, w)]*2 }

	// EXTRA: 输入 v 和 to，to 可能是 v 的子孙，返回从 v 到 to 路径上的第二个节点（v 的一个儿子）
	// 如果 to 不是 v 的子孙，返回 -1
	// https://codeforces.com/problemset/problem/1702/G2 2000
	// https://codeforces.com/problemset/problem/916/E 2400
	down1 := func(v, to int) int {
		if dep[to] <= dep[v] {
			return -1
		}
		to = uptoDep(to, dep[v]+1)
		if pa[to][0] == v {
			return to
		}
		return -1
	}

	// EXTRA: 从 v 出发，向 to 方向走一步
	// 输入需要保证 v != to
	move1 := func(v, to int) int {
		if v == to {
			panic(-1)
		}
		if getLCA(v, to) == v { // to 在 v 下面
			return uptoDep(to, dep[v]+1)
		}
		// lca 在 v 上面
		return pa[v][0]
	}

	// EXTRA: 从 v 开始，向上跳 k 步
	// 不存在则返回 -1
	// O(1) 求法见长链剖分
	uptoKthPa := func(v, k int) int {
		for ; k > 0 && v != -1; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(uint(k))]
		}
		return v
	}

	// EXTRA: 输入 v 和 w，返回 v 到 w 路径上的中点
	// 返回值是一个数组，因为可能有两个中点
	// 在有两个中点的情况下，保证返回值的第一个中点离 v 更近
	midPath := func(v, w int) []int {
		lca := getLCA(v, w)
		dv := dep[v] - dep[lca]
		dw := dep[w] - dep[lca]
		if dv == dw {
			return []int{lca}
		}
		if dv > dw {
			mid := uptoKthPa(v, (dv+dw)/2)
			if (dv+dw)%2 == 0 {
				return []int{mid}
			}
			return []int{mid, pa[mid][0]}
		} else {
			mid := uptoKthPa(w, (dv+dw)/2)
			if (dv+dw)%2 == 0 {
				return []int{mid}
			}
			return []int{pa[mid][0], mid} // pa[mid][0] 离 v 更近
		}
	}

	{
		// 加权树上二分
		var dep []int // 加权深度，dfs 预处理略
		// 从 v 开始向根移动至多 d 距离，返回最大移动次数，以及能移动到的离根最近的点
		// NOIP2012·提高 疫情控制 https://www.luogu.com.cn/problem/P1084
		// 变形 https://codeforces.com/problemset/problem/932/D
		// 点权写法 https://codeforces.com/problemset/problem/1059/E 2400
		// https://www.luogu.com.cn/problem/P7167
		uptoDep := func(v, d int) (int, int) {
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
		// EXTRA: 倍增的时候维护其他属性，如边权最值等
		// 下面的代码来自 https://codeforces.com/problemset/problem/609/E 2100
		// EXTRA: 额外维护最值边的下标，见 https://codeforces.com/contest/733/submission/120955685
		// 点权写法 https://codeforces.com/problemset/problem/1059/E 2400
		type nb struct{ to, wt int }
		var g [][]nb // read g ...

		const mx = 18
		type data int
		type pair struct {
			p     int
			maxWt data
		}
		pa := make([][mx]pair, len(g))
		dep := make([]int, len(g))
		var build func(v, p, d int)
		build = func(v, p, d int) {
			pa[v][0].p = p
			dep[v] = d
			for _, e := range g[v] {
				if w := e.to; w != p {
					pa[w][0].maxWt = data(e.wt)
					build(w, v, d+1)
				}
			}
		}
		build(0, -1, 0)

		merge := func(a, b data) data {
			return data(max(int(a), int(b)))
		}

		for i := 0; i+1 < mx; i++ {
			for v := range pa {
				if p := pa[v][i]; p.p != -1 {
					pp := pa[p.p][i]
					pa[v][i+1] = pair{pp.p, merge(p.maxWt, pp.maxWt)}
				} else {
					pa[v][i+1].p = -1
				}
			}
		}

		// 求 LCA(v,w) 的同时，顺带求出 v-w 上的边权最值
		getLCA := func(v, w int) (lca int, maxWt data) {
			//pathLen := dep[v] + dep[w]
			if dep[v] > dep[w] {
				v, w = w, v
			}
			for k := dep[w] - dep[v]; k > 0; k &= k - 1 {
				p := pa[w][bits.TrailingZeros(uint(k))]
				maxWt = merge(maxWt, p.maxWt)
				w = p.p
			}
			if w != v {
				for i := mx - 1; i >= 0; i-- {
					if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
						maxWt = merge(maxWt, merge(pv.maxWt, pw.maxWt))
						v, w = pv.p, pw.p
					}
				}
				maxWt = merge(maxWt, merge(pa[v][0].maxWt, pa[w][0].maxWt))
				v = pa[v][0].p
			}
			// 如果是点权的话这里加上 maxWt = merge(maxWt, pa[v][0].maxWt)
			lca = v
			//pathLen -= dep[lca] * 2
			return
		}

		_ = getLCA
	}

	_ = []interface{}{getDis, uptoKthPa, down1, move1, midPath}
}

// 最近公共祖先 · 其二 · 基于 RMQ
// O(nlogn) 预处理，O(1) 查询
// 由于预处理 ST 表是基于一个长度为 2n 的序列，所以常数上是比倍增算法要大的。内存占用也比倍增要大一倍左右（这点可忽略）
// 优点是查询的复杂度低，适用于查询量大的情形
// https://oi-wiki.org/graph/lca/#rmq
// todo DFS 序求 LCA（常数更小） https://www.cnblogs.com/alex-wei/p/DFN_LCA.html
// https://codeforces.com/problemset/problem/342/E
// 注：如果只有路径修改+查询，可以用欧拉序列 + 树状数组/线段树，见《挑战》p.332 http://poj.org/problem?id=2763
func (*tree) lcaRMQ(root int, g [][]int) {
	vs := make([]int, 0, 2*len(g)-1)  // 欧拉序列
	pos := make([]int, len(g))        // pos[v] 表示 v 在 vs 中第一次出现的位置编号
	dep := make([]int, 0, 2*len(g)-1) // 深度序列，和欧拉序列一一对应
	disRoot := make([]int, len(g))    // disRoot[v] 表示 v 到 root 的距离
	var build func(v, p, d int)       // 若有边权需额外传参 dis
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

	type stPair struct{ v, i int }
	const mx = 17 + 1 // bits.Len(len(dep))
	var st [][mx]stPair
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]stPair, n)
		for i, v := range a {
			st[i][0] = stPair{v, i}
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
	getLCA := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return vs[stQuery(pv, pw+1)]
	}
	getDis := func(v, w int) int { return disRoot[v] + disRoot[w] - disRoot[getLCA(v, w)]*2 }

	_ = getDis
}

// 最近公共祖先 · 其三 · Tarjan 离线算法
// 时间复杂度 O(n+qα)
// 原论文 https://dl.acm.org/doi/pdf/10.1145/800061.808753
// https://core.ac.uk/download/pdf/82125836.pdf
// https://oi-wiki.org/graph/lca/#tarjan
// https://cp-algorithms.com/graph/lca_tarjan.html
// todo 线性做法 https://ljt12138.blog.uoj.ac/blog/4874
// 类似思路可以用在求 RMQ 上 https://codeforces.com/blog/entry/48994
//
// LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
func (*tree) lcaTarjan(root int, edges, queries [][]int) []int {
	n := len(edges) - 1
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	nq := len(queries)
	lca := make([]int, nq)
	dis := make([]int, nq) // dis(q.v,q.w)
	type query struct{ w, i int }
	qs := make([][]query, len(g))
	for i, q := range queries {
		v, w := q[0], q[1]
		// 第一种写法：保证在 v=w 时恰好只更新一个（结合下面的 if w := q.w; w == v || ... 理解）
		qs[v] = append(qs[v], query{w, i})
		if v != w {
			qs[w] = append(qs[w], query{v, i})
		}

		// 第二种写法：单独处理 v==w 的情况
		//if v != w {
		//	qs[v] = append(qs[v], query{w, i})
		//	qs[w] = append(qs[w], query{v, i})
		//} else {
		//	// do v==w...
		//	lca[i] = v
		//	dis[i] = 0
		//}
	}

	_fa := make([]int, len(g))
	for i := range _fa {
		_fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if _fa[x] != x {
			_fa[x] = find(_fa[x])
		}
		return _fa[x]
	}

	dep := make([]int, len(g))
	// 为什么不用 bool 数组？
	// 对于下面代码中的 do(v, w, lcaVW)
	// 如果 v 是 w 的祖先节点，那么 w 递归结束后会触发一次，v 递归结束后又会触发一次
	// 如果 do 中有增量更新，这样就错了
	// 而三色标记法可以保证只会触发一次
	color := make([]int8, len(g))
	var tarjan func(int, int)
	tarjan = func(v, d int) {
		dep[v] = d
		color[v] = 1
		for _, w := range g[v] {
			if color[w] == 0 {
				tarjan(w, d+1)
				_fa[w] = v // 相当于把 w 的子树节点全部 merge 到 v
			}
		}
		for _, q := range qs[v] {
			w := q.w
			// color[w] == 2 意味着 y 所在子树已经遍历完
			// 也就意味着 w 已经 merge 到它和 v 的 LCA 上了
			if w == v || color[w] == 2 {
				lcaVW := find(w)
				lca[q.i] = lcaVW
				dis[q.i] = dep[v] + dep[w] - dep[lcaVW]*2
				// do(v, w, lcaVW)...
			}
		}
		color[v] = 2
	}
	tarjan(root, 0)
	return lca
}

// LCA 应用：树上差分
// 入门（单点更新）LC2445 https://leetcode.cn/problems/number-of-nodes-with-value-one/
// 操作为更新 v-w 路径上的点权或边权（初始为 0）
// 点权时 diff[lca] -= val 且 diff[father[lca]] -= val
//    把 x-lca-y 看成 x-lca'-lca-y，这里 lca' 是 lca 的儿子，
//    那么 x-lca' 就对应着 diff[x] += val 且 diff[lca] -= val
//    lca-y 就对应着 diff[y] += val 且 diff[father[lca]] -= val
// 边权时 diff[lca] -= 2 * val（定义 diff 为点到父亲的差分值）
// https://www.luogu.com.cn/blog/RPdreamer/ci-fen-and-shu-shang-ci-fen
// https://zhuanlan.zhihu.com/p/61299306
// todo https://loj.ac/d/1698
// 模板题（点权）https://www.luogu.com.cn/problem/P3128 LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
// 模板题（边权）https://codeforces.com/problemset/problem/191/C
// todo https://www.luogu.com.cn/problem/P2680
// https://codeforces.com/problemset/problem/1707/C
func (*tree) differenceInTree(n, root int, g, queries [][]int) []int {
	var pa [][]int
	var getLCA func(int, int) int

	diff := make([]int, n)
	update := func(v, w int, val int) {
		diff[v] += val
		diff[w] += val
		lca := getLCA(v, w)
		diff[lca] -= val // 点权
		if f := pa[lca][0]; f >= 0 {
			diff[f] -= val // 点权
		}
		//diff[lca] -= 2 * val // 边权
	}
	for _, q := range queries {
		v, w, val := q[0], q[1], q[2]
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

// LCA+DFN：虚树 Virtual Tree / Auxiliary Tree
// https://oi-wiki.org/graph/virtual-tree/ 栈相比两次排序，效率更高
// 【点评】除了 DFS 遍历虚树，也可以把虚树上的所有点整合到数组 vtNodes 中，并计算出每个点的父节点 vtPa
//        然后按照 DFN 从小到大排序 vtNodes，倒着遍历 vtNodes 数组，就可以直接自底向上算了
//        但这样做还要再排序，并使用几个额外的数组，所以相比之下直接 DFS 更好
//
// 题单 https://www.luogu.com.cn/training/3682#problems
// 入门 https://codeforces.com/problemset/problem/613/D 2800
// 换根 DP https://codeforces.com/problemset/problem/1320/E 3000
// https://www.luogu.com.cn/problem/P4103 [HE14] 大工程（点对距离和，最短路径，最长路径/直径）
// - https://atcoder.jp/contests/abc359/tasks/abc359_g 点对距离和
// https://www.luogu.com.cn/problem/P3233 [HN14] 世界树
// https://www.luogu.com.cn/problem/P2495 [SD11] 消耗战
// https://www.luogu.com.cn/problem/P5891
// https://www.luogu.com.cn/problem/P7409
func (*tree) virtualTree(g [][]int) {
	var dep []int
	var getLCA func(int, int) int

	dfn := make([]int, len(g))
	ts := 0
	_ = ts
	// 向上查找<lcaBinaryLifting>
	// 在 buildPa 开头添加：
	// dfn[v] = ts; ts++

	vt := make([][]int, len(g))
	// vt := make([][]edge, len(g))
	inNodes := make([]int, len(g))
	for i := range inNodes {
		inNodes[i] = -1
	}
	addVtEdge := func(v, w int) {
		vt[v] = append(vt[v], w)
		// wt := dep[w] - dep[v]
		// vt[v] = append(vt[v], edge{w, wt})
		// 也可以在 DFS 的时候算出边权
	}
	const root = 0
	st := []int{root} // 用根节点作为栈底哨兵
	// nodes 为询问的「关键节点」
	do := func(nodes []int, qid int) {
		slices.SortFunc(nodes, func(a, b int) int { return dfn[a] - dfn[b] })
		vt[root] = vt[root][:0]
		st = st[:1]
		for _, v := range nodes {
			inNodes[v] = qid // 时间戳
			if v == root {
				continue
			}
			// ... 某些题目需要判断 v 和 pa[v][0] 是否都在 nodes 中
			vt[v] = vt[v][:0]
			lca := getLCA(st[len(st)-1], v)
			// 回溯
			for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
				addVtEdge(st[len(st)-2], st[len(st)-1])
				st = st[:len(st)-1]
			}
			if lca != st[len(st)-1] { // lca 不在栈中（首次遇到）
				vt[lca] = vt[lca][:0]
				addVtEdge(lca, st[len(st)-1])
				st[len(st)-1] = lca // 加到栈中
				// ... 标记 lca 是虚点
			}
			st = append(st, v)
		}
		// 最后的回溯
		for i := 1; i < len(st); i++ {
			addVtEdge(st[i-1], st[i])
		}

		// 计算 nodes 两两间的：
		// 路径长度和
		// 最短路径
		// 最长路径（直径）
		const inf int = 1e18
		sumWt, gMinL, gMaxL := 0, inf, 0
		var f func(int) (int, int, int)
		f = func(v int) (size, minL, maxL int) {
			// 如果 inNodes[v] != qid，那么 v 只是关键节点之间路径上的「拐点」
			// 在处理虚树 DP 时，可能需要额外考虑 v 不在 nodes 中的情况
			imp := inNodes[v] == qid
			if imp {
				size = 1
			}
			minL = inf
			for _, w := range vt[v] {
				sz, mn, mx := f(w)
				wt := dep[w] - dep[v]
				sumWt += wt * sz * (len(nodes) - sz) // 贡献法
				size += sz

				mn += wt
				if imp {
					gMinL = min(gMinL, mn)
				} else {
					gMinL = min(gMinL, minL+mn)
					minL = min(minL, mn)
				}

				mx += wt
				gMaxL = max(gMaxL, maxL+mx)
				maxL = max(maxL, mx)
			}
			if minL == inf {
				minL = 0
			}
			return
		}
		rt := root
		if inNodes[rt] != qid && len(vt[rt]) == 1 {
			// 注意 root 只是一个哨兵，得从真正的根节点开始
			rt = vt[rt][0]
		}
		f(rt)

		// 收尾工作：reset 数组
		//for _, v := range nodes {
		//	idx[v] = 0
		//}
	}

	_ = do
}

// 树链剖分/重链剖分 (HLD, Heavy Light Decomposition）
// 性质：
//    树上每个结点都属于且仅属于一条重链
//    如果 v-w 是一条轻边，那么 size[w] < size[v]/2
//    从根结点到任意结点所经过的重链数为 O(logn)，轻边数为 O(logn)
//    与重心的关系（见 findCentroid）
// https://en.wikipedia.org/wiki/Heavy_path_decomposition
//    把每条重链当成一个节点，每条轻边作为边，我们可以得到一棵路径树。显然路径树的高度为 O(logn)
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
// O(log n) 查询 https://codeforces.com/blog/entry/127896
//
// 注：若没有修改操作，更简单的做法见 lcaBinaryLifting（路径查询）以及 dfnOrder（子树查询）
// 注：如果只有路径修改+查询，可以用欧拉序列 + 树状数组/线段树，见《挑战》p.332
//
// 模板题（点权）https://www.luogu.com.cn/problem/P3384
//            https://codeforces.com/problemset/problem/343/D 2100
// 模板题（边权）https://atcoder.jp/contests/abc294/tasks/abc294_g
// - 也可以转换成子树所有点的 dis 都增加了 delta，用欧拉序+差分树状数组维护
// 与最小生成树结合（边权）https://codeforces.com/problemset/problem/609/E
// 好题 https://codeforces.com/contest/1174/problem/F
// 归并树 https://codeforces.com/problemset/problem/587/C
// todo 题单 https://www.luogu.com.cn/training/1654
// todo https://www.luogu.com.cn/problem/P5127
// TODO: 处理边权的情况
//   https://www.luogu.com.cn/problem/P1505
//   https://www.luogu.com.cn/problem/P4315
//   https://www.luogu.com.cn/problem/P4114
// todo NOI21 轻重边 https://www.luogu.com.cn/problem/P7735
//  https://www.luogu.com.cn/problem/P4211
// 结合广义圆方树 https://codeforces.com/problemset/problem/487/E
func (*tree) heavyLightDecomposition(n, root int, g [][]int, vals []int) { // vals 为点权
	// 深度，子树大小，重儿子，父节点，所处重链顶点（深度最小），DFS 序（作为线段树中的编号，从 1 开始）
	type node struct{ depth, size, hson, fa, top, dfn int }
	nodes := make([]node, n)
	//idv := make([]int, n+1) // idv[nodes[v].dfn] == v
	var build func(int, int, int) int
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

	dfn := 0
	var markTop func(int, int)
	markTop = func(v, top int) {
		o := &nodes[v]
		o.top = top
		o.dfn = dfn
		//idv[dfn] = v
		dfn++
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
	dfnVals := make([]int, n)
	for i, v := range vals {
		dfnVals[nodes[i].dfn] = v
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
	updatePath := func(v, w, add int) { doPath(v, w, func(l, r int) { t.update(1, l, r, add) }) }
	queryPath := func(v, w int) (sum int) { doPath(v, w, func(l, r int) { sum += t.query(1, l, r) }); return } // % mod
	updateSubtree := func(v, add int) { o := nodes[v]; t.update(1, o.dfn, o.dfn+o.size-1, add) }
	querySubtree := func(v int) (sum int) { o := nodes[v]; return t.query(1, o.dfn, o.dfn+o.size-1) }

	// EXTRA: 寻找以 st 为重链顶点的重链
	// hPath[-1] 即为重链末端节点
	getHP := func(st int) []int {
		hPath := []int{st}
		for o := nodes[st]; o.hson != -1; o = nodes[o.hson] {
			hPath = append(hPath, o.hson)
		}
		return hPath
	}

	_ = []any{updatePath, queryPath, updateSubtree, querySubtree, getHP}
}

// 长链剖分
// 长链剖分和重链剖分一样，是把一棵树分成若干条不相交的链
// 但是，这里的重儿子不再是子树大小最大的，而是深度最大的子节点（长儿子）
// 根据这一定义可推出，从根结点到任意结点所经过的轻边数为 O(√n) (想象一棵长儿子不断递减的二叉树)
// https://oi-wiki.org/graph/hld/#_14
// https://www.luogu.com.cn/blog/Ynoi/zhang-lian-pou-fen-xue-xi-bi-ji
// https://www.cnblogs.com/cj-chd/p/10076199.html
// https://www.cnblogs.com/zhoushuyu/p/9468669.html
// 应用：树上 k 级祖先 https://www.luogu.com.cn/problem/P5903 https://codeforces.com/problemset/problem/208/E
// 长链剖分优化树形 DP：
//    若树形 DP 的转移只和节点深度有关，我们完全可以把一棵子树拍扁成一条垂直的链
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
// 每个节点合并到另外一个集合中的次数不超过 O(nlogn)，因为一旦合并到另一个集合，这个节点所在的集合大小至少倍增
// https://oi-wiki.org/graph/dsu-on-tree/
// NOTE: 合并的时候最好先循环计算一遍答案，再循环合并一遍，这样的习惯可避免产生把合并之后的数值算入答案中的 bug
// 讲解+套题 https://pzy.blog.luogu.org/dsu-on-tree-xue-xi-bi-ji
// 讲解+套题 https://codeforces.com/blog/entry/44351 补充 https://codeforces.com/blog/entry/67696
// todo 套题 https://blog.csdn.net/m0_49959202/article/details/114925708
// 模板题 https://www.luogu.com.cn/problem/U41492
//       https://codeforces.com/problemset/problem/600/E https://www.acwing.com/problem/content/3191/
// todo HNOI09 梦幻布丁 https://www.luogu.com.cn/problem/P3201 https://www.acwing.com/problem/content/2156/
// LC2003 所有子树 mex https://leetcode.cn/problems/smallest-missing-genetic-value-in-each-subtree/
// 距离等于 k 的点对数 https://codeforces.com/problemset/problem/161/D
//            变形题 https://ac.nowcoder.com/acm/contest/4853/E 题解 https://ac.nowcoder.com/discuss/394080
// https://ac.nowcoder.com/acm/contest/4010/E
// https://atcoder.jp/contests/abc183/tasks/abc183_f
// https://codeforces.com/contest/1455/problem/G
// https://codeforces.com/contest/570/problem/D
// https://codeforces.com/contest/246/problem/E
// https://codeforces.com/contest/208/problem/E
// https://codeforces.com/contest/1009/problem/F
// https://codeforces.com/contest/375/problem/D
// https://codeforces.com/contest/741/problem/D
// https://codeforces.com/problemset/problem/1805/E
// https://codeforces.com/contest/1824/problem/C

// 写法一：按 map 的大小合并
// 路径点权异或 https://codeforces.com/problemset/problem/1709/E
func (*tree) smallToLarge(root int, g [][]int, vals []int) { // vals 为点权
	var f func(int, int, int) map[int]bool
	f = func(v, fa, xor int) map[int]bool {
		xor ^= vals[v]
		m := map[int]bool{xor: true}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			subM := f(w, v, xor)
			if len(subM) > len(m) {
				m, subM = subM, m
			}
			// check subM ...
			maps.Copy(m, subM) // m <- subM
		}
		return m
	}
	f(root, -1, 0)
}

// 写法二：轻重儿子合并 · map 版本
// 根节点到树上任意节点的轻边数不超过 O(logn) 条
// 某些题目可以不用 map，而是像莫队那样添加和撤销，这样只用数组就行，例如 https://codeforces.com/problemset/problem/375/D
func (*tree) dsuMap(root int, g [][]int, vals []int) { // vals 为点权
	hson := make([]int, len(g))
	var build func(int, int) int
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
		hson[v] = hs // 叶子的重儿子是 -1
		return sz
	}
	build(root, -1)

	// 例如：统计子树的点权种类数
	ans := make([]int, len(g))
	var f func(int, int) map[int]bool
	f = func(v, fa int) map[int]bool {
		if hson[v] < 0 { // 叶结点
			ans[v] = 1
			return map[int]bool{vals[v]: true}
		}
		has := f(hson[v], v) // 先算重儿子
		merge := func(val int) {
			// do...
			has[val] = true
		}
		for _, w := range g[v] {
			if w != fa && w != hson[v] { // 其余儿子合并到重儿子的结果中
				subM := f(w, v)
				for val := range subM {
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

// 写法三：轻重儿子合并 · 数组版本
// 根节点到树上任意节点的轻边数不超过 O(logn) 条
// 不用 map，而是像莫队那样添加和撤销，这样只需要数组，常数更小
// 例如 https://codeforces.com/problemset/problem/375/D
func (*tree) dsuArr(root int, g [][]int, vals []int) { // vals 为点权
	dfn := 0
	nodes := make([]struct{ l, r, hson int }, len(g)) // [l,r)
	nodeVals := make([]int, 0, len(g))
	var build func(int, int) int
	build = func(v, fa int) int {
		nodes[v].l = dfn
		dfn++
		nodeVals = append(nodeVals, vals[v])
		size, hsz, hson := 1, 0, -1
		for _, w := range g[v] {
			if w != fa {
				sz := build(w, v)
				size += sz
				if sz > hsz {
					hsz, hson = sz, w
				}
			}
		}
		nodes[v].r = nodes[v].l + size
		nodes[v].hson = hson
		return size
	}
	build(root, -1)

	//（离线）统计子树中的出现次数 >= k 的点权个数（1 <= k <= n）
	cnt := [1e5 + 1]int{}
	cc := make([]int, len(nodeVals)+1)
	var f func(int, int)
	f = func(v, fa int) {
		hson := nodes[v].hson
		for _, w := range g[v] {
			if w == fa || w == hson {
				continue
			}
			f(w, v)
			// 恢复现场，这样下一棵子树不会受到影响
			for _, x := range nodeVals[nodes[w].l:nodes[w].r] {
				cc[cnt[x]]--
				cnt[x]--
			}
		}
		if hson >= 0 {
			f(hson, v)
			// 此时重儿子的数据已经添加
		}

		// 添加根节点的数据
		cnt[vals[v]]++
		cc[cnt[vals[v]]]++

		for _, w := range g[v] {
			if w == fa || w == hson {
				continue
			}
			// 添加非重儿子的数据
			for _, x := range nodeVals[nodes[w].l:nodes[w].r] {
				cnt[x]++
				cc[cnt[x]]++
			}
		}

		// 子树 v 的每个节点的数据都已添加，回答询问
		// 此时 cc[k] 就是子树 v 中的出现次数 >= k 的点权个数
		// ...
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

// 普吕弗序列（Prufer 序列，Prüfer sequence）
// https://en.wikipedia.org/wiki/Pr%C3%BCfer_sequence
// https://oeis.org/A000272 Cayley's formula https://en.wikipedia.org/wiki/Cayley%27s_formula
// https://www.luogu.com.cn/problem/P6086
// todo 光之大陆 https://www.acwing.com/problem/content/2420/
// todo https://codeforces.com/problemset/problem/156/D
func (*tree) treeToPrufer(n int, pa []int) []int { // 传入的 pa 是以 n 为根时的每个节点的父节点
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

func (*tree) pruferToTree(n int, prufer []int) []int {
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
