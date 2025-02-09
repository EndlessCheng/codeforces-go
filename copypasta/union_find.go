package copypasta

import (
	. "fmt"
	"io"
	"slices"
)

/* 并查集
只有路径压缩的并查集复杂度是 O(nlogn) 的，这也是大多数情况下的实现方案
只有启发式合并（按深度合并）的并查集的复杂度也是 O(nlogn) 的，适用于可持久化的场景

只有路径压缩的并查集，可以构造一棵二项树（binomial tree）
结合图片讲解 https://upload.wikimedia.org/wikipedia/commons/c/cf/Binomial_Trees.svg (https://en.wikipedia.org/wiki/Binomial_heap)
每次把二项树的根节点连到一个新的孤立点上，然后对最深的点调用 find
这样可以得到一棵几乎一样的树（区别仅仅是根节点多了一个儿子）
所以，只要重复上述过程，就可以让每次 find 都是 O(logn) 级别的了

具体的时间复杂度证明见《算法导论》
https://zhuanlan.zhihu.com/p/553192435

随机合并下的时间复杂度 https://www.cis.upenn.edu/~sanjeev/papers/soda14_disjoint_set_union.pdf
*/

// 普通并查集
// 可视化 https://visualgo.net/zh/ufds
// https://oi-wiki.org/ds/dsu/
// https://cp-algorithms.com/data_structures/disjoint_set_union.html
// 并查集时间复杂度证明 https://oi-wiki.org/ds/dsu-complexity/
// RMQ 标准算法和线性树上并查集 https://ljt12138.blog.uoj.ac/blog/4874
//
// 另见 graph.go 中的 MST
//
// 模板题 LC547 https://leetcode.cn/problems/number-of-provinces/
// LC684 https://leetcode.cn/problems/redundant-connection/
// LC1267 https://leetcode.cn/problems/count-servers-that-communicate/
// 并查集+置换 LC2459 https://leetcode.cn/problems/sort-array-by-moving-items-to-empty-space/
// https://www.luogu.com.cn/problem/P1111
// https://www.luogu.com.cn/problem/P1967 经典题
// https://www.luogu.com.cn/problem/P3367
// https://atcoder.jp/contests/arc097/tasks/arc097_b
// 基础题 https://codeforces.com/problemset/problem/1167/C 1400
//       https://codeforces.com/problemset/problem/2060/E 1500
//       https://codeforces.com/problemset/problem/1411/C 1700
// https://codeforces.com/problemset/problem/755/C 1300
// LC1562 https://leetcode.cn/problems/find-latest-group-of-size-m/
// 转换 https://atcoder.jp/contests/abc304/tasks/abc304_e
// 转换 https://atcoder.jp/contests/abc238/tasks/abc238_e
// merge 后 from 还有用 https://atcoder.jp/contests/abc279/tasks/abc279_f
// 处理图上的环 https://codeforces.com/contest/1726/problem/D
// https://codeforces.com/problemset/problem/1851/G 2000 离线
// 对偶图 LC3235 https://leetcode.cn/problems/check-if-the-rectangle-corner-is-reachable/
// - 思考：如果允许圆心在矩形外呢？见我题解下面的讨论
// https://atcoder.jp/contests/abc383/tasks/abc383_e
// https://atcoder.jp/contests/abc392/tasks/abc392_e
//
// 质因子并查集 GCD>1 并查集
// 预处理质因子（见 math.go 中的 primeDivisorsAll）
// 枚举 a[i] 的质因子 p，用 pre[p] 记录质因子上一次出现的下标（初始化成 -1），然后 merge(i, pre[p]) if pre[p] > 0
// - [2709. 最大公约数遍历](https://leetcode.cn/problems/greatest-common-divisor-traversal/) 2172
// - [1627. 带阈值的图连通性](https://leetcode.cn/problems/graph-connectivity-with-threshold/) 2221
// - [952. 按公因数计算最大组件大小](https://leetcode.cn/problems/largest-component-size-by-common-factor/) 2272
//       下标距离有上界 https://codeforces.com/contest/1978/problem/F
// - [1998. 数组的最大公因数排序](https://leetcode.cn/problems/gcd-sort-of-an-array/) 2429
//
// 数组标记/区间合并相关
// - 经典模型是一维区间覆盖染色，通过倒序+并查集解决
// - 顺带补充下二维的情况（非并查集）：LC2718 https://leetcode.cn/problems/sum-of-matrix-after-queries/
// - [1851. 包含每个查询的最小区间](https://leetcode.cn/problems/minimum-interval-to-include-each-query/)
// - [2382. 删除操作后的最大子段和](https://leetcode.cn/problems/maximum-segment-sum-after-removals/)
// - [2334. 元素值大于变化阈值的子数组](https://leetcode.cn/problems/subarray-with-elements-greater-than-varying-threshold/)
// - [2612. 最少翻转操作数](https://leetcode.cn/problems/minimum-reverse-operations/)
// https://codeforces.com/problemset/problem/1041/C 1600
// https://codeforces.com/problemset/problem/827/A 1700
// https://codeforces.com/problemset/problem/1157/E 1700
// https://codeforces.com/problemset/problem/724/D 1900
// https://codeforces.com/problemset/problem/2018/D 2200
// https://www.codechef.com/problems/REMSUBARR
//
// 树+点权/边权的顺序
// LC2421 https://leetcode.cn/problems/number-of-good-paths/
// 贡献法 https://codeforces.com/problemset/problem/915/F
// 贡献法 https://atcoder.jp/contests/abc214/tasks/abc214_d
//
// LC2503 https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/
// 接水问题 https://codeforces.com/problemset/problem/371/D
// LC407 三维接雨水 https://leetcode.cn/problems/trapping-rain-water-ii/
// - https://www.luogu.com.cn/problem/P5930 
// 使某些点不在环上需要删除的最少边数 https://ac.nowcoder.com/acm/contest/7780/C
// todo https://codeforces.com/problemset/problem/292/D
// 任意合并+区间合并 https://codeforces.com/problemset/problem/566/D
// 动态加点 https://codeforces.com/contest/1494/problem/D
// 思维转换 https://nanti.jisuanke.com/t/43488
//         https://codeforces.com/problemset/problem/1012/B
//         https://codeforces.com/problemset/problem/1466/F
// https://codeforces.com/problemset/problem/455/C 2100
// 前缀和 后缀和 https://codeforces.com/problemset/problem/292/D
// 维护树或基环树 https://codeforces.com/problemset/problem/859/E
// 求矩阵的 rank 矩阵 https://codeforces.com/problemset/problem/650/C LC1632 https://leetcode.cn/problems/rank-transform-of-a-matrix/submissions/
// 分组排序套路 LC1998 https://leetcode.cn/problems/gcd-sort-of-an-array/
// 套题 https://blog.csdn.net/weixin_43914593/article/details/104108049 算法竞赛专题解析（3）：并查集
// 转换 https://codeforces.com/problemset/problem/1253/D
// 离散 + 四方向 Kick Start 2019 Round C Wiggle Walk https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aac#analysis
// 能力守恒+离线 https://codeforces.com/contest/1851/problem/G
// 技巧：去掉无用数据
// - https://codeforces.com/problemset/problem/1157/E
// - https://codeforces.com/problemset/problem/1791/F
// todo https://codeforces.com/contest/884/problem/E
// https://codeforces.com/problemset/problem/1416/D 2600 DSU 重构树

// 轻量级模板
// 采用非递归写法，效率更好
func _(n int) {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	_ = find
}

type UnionFind struct {
	Fa     []int
	Groups int // 连通分量个数
}

func NewUnionFind(n int) UnionFind {
	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	return UnionFind{fa, n}
}

// 非递归版本
func (u UnionFind) Find(x int) int {
	root := x
	for u.Fa[root] != root {
		root = u.Fa[root]
	}
	for u.Fa[x] != root {
		u.Fa[x], x = root, u.Fa[x]
	}
	return root
}

// 递归版本
func (u UnionFind) FindR(x int) int {
	if u.Fa[x] != x {
		u.Fa[x] = u.FindR(u.Fa[x])
	}
	return u.Fa[x]
}

// newRoot = -1 表示未发生合并
func (u *UnionFind) Merge(from, to int) (newRoot int) {
	x, y := u.Find(from), u.Find(to)
	if x == y {
		return -1
	}
	u.Fa[x] = y
	u.Groups--
	return y
}

func (u UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

// 以下代码可用于比赛时复制

func _(n int) {
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
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	// 总是合并到代表元更大的树上
	mergeBig := func(from, to int) int {
		ff, ft := find(from), find(to)
		if ff > ft {
			ff, ft = ft, ff
		}
		fa[ff] = ft
		return ft
	}

	{
		// 哈希表版本离散化版本
		// LC947 https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column/
		// https://codeforces.com/problemset/problem/506/D 2400
		// class 版本见 https://codeforces.com/problemset/submission/506/247878263
		fa := map[int]int{}
		groups := 0
		var find func(int) int
		find = func(x int) int {
			fx, ok := fa[x]
			if !ok {
				fa[x] = x
				fx = x
				groups++
			}
			if fx != x {
				fa[x] = find(fx)
				return fa[x]
			}
			return x
		}

		// merge，并返回新的 root
		merge := func(from, to int) int {
			x, y := find(from), find(to)
			if x == y {
				return -2e9
			}
			fa[x] = y
			groups--
			return y
		}

		_ = merge
	}

	// 区间并查集 / 涂色并查集 / 刷墙并查集
	// LC1851 https://leetcode.cn/problems/minimum-interval-to-include-each-query/ 2286
	// - 变形：额外传入数组 nums，计算包含元素值为 queries[i] 的最短的 intervals[i]
	// LC2158 https://leetcode.cn/problems/amount-of-new-area-painted-each-day/
	// https://codeforces.com/problemset/problem/371/D 1800 经典题
	// https://codeforces.com/problemset/problem/2020/D 1800
	// https://codeforces.com/problemset/problem/724/D 1900
	mergeRangeTo := func(l, r, to int) { 
		// 常用：to = r + 1，这时建议用左闭右开表示区间
		// 或者 to = find(r)
		// l = max(l, 1)
		// r = min(r, n)
		for i := find(l); i <= r; i = find(i + 1) { // initFa 需要开 n+1 空间（或者 n+2，如果下标从 1 开始）
			fa[i] = to
		}
	}

	//rangeFullMerged := func() bool { return find(0) == n }

	// 连通分量个数
	// countRoots > 1 表示整个图不连通
	countRoots := func(st int) (cnt int) {
		for i := st; i < len(fa); i++ {
			if find(i) == i {
				cnt++
			}
		}
		return
	}

	// 所有代表元
	getRoots := func() (roots []int) {
		for i := range fa {
			if find(i) == i {
				roots = append(roots, i)
			}
		}
		return
	}

	// 连通分量
	getComps := func() (comps map[int][]int) {
		comps = map[int][]int{}
		for i := range fa {
			f := find(i)
			comps[f] = append(comps[f], i)
		}
		return
	}

	// 同一连通块内的数字从小到大排序
	// https://codeforces.com/contest/1971/problem/G
	sortCC := func(a []int) {
		n := len(a)
		comps := make([][]int, n)
		for i := 0; i < n; i++ {
			rt := find(i)
			comps[rt] = append(comps[rt], i)
		}
		for _, cc := range comps {
			if cc == nil {
				continue
			}
			// 注意 cc 已经是有序的
			// 收集连通块内的元素，排序，然后重新填回去
			b := make([]int, len(cc))
			for ci, i := range cc {
				b[ci] = a[i]
			}
			slices.Sort(b)
			for ci, i := range cc {
				a[i] = b[ci]
			}
		}
	}

	{
		// 按秩合并
		rank := make([]int, n)
		merge := func(x, y int) {
			x, y = find(x), find(y)
			if x == y {
				return
			}
			if rank[x] > rank[y] {
				x, y = y, x
			}
			fa[x] = y
			if rank[x] == rank[y] {
				rank[y]++
			}
		}
		_ = merge
	}

	_ = []interface{}{merge, same, mergeBig, mergeRangeTo, getRoots, countRoots, getComps, sortCC}
}

// 用并查集实现有序集合的删除、查找前驱和查找后继
// 无法构造一棵二项树，单次操作均摊复杂度是 O(1) 的
// LC https://leetcode.cn/problems/block-placement-queries/
type delUf struct {
	left  []int
	right []int
}

func newDelUf(n int) delUf {
	// 把 0 和 n+1 当作哨兵
	// 如果有删除 0 的情况，想清楚有没有 corner case
	left := make([]int, n+2)
	right := make([]int, n+2)
	for i := range left {
		left[i] = i
		right[i] = i
	}
	return delUf{left, right}
}

func (f delUf) _find(fa []int, x int) int {
	if fa[x] != x {
		fa[x] = f._find(fa, fa[x])
	}
	return fa[x]
}

// 删除 x
func (f delUf) delete(x int) {
	if f._find(f.left, x) != x { // x 已经被删除
		return
	}
	f.left[x] = x - 1
	f.right[x] = x + 1
}

// 查找前驱：返回严格小于 x 的最大元素
func (f delUf) prev(x int) int {
	if x <= 0 {
		panic("x must be positive")
	}
	return f._find(f.left, x-1)
}

// 查找后继：返回严格大于 x 的最小元素
func (f delUf) next(x int) int {
	return f._find(f.right, x+1)
}

//

// 二维并查集
type ufPoint struct{ x, y int } // int32
type uf2d map[ufPoint]ufPoint

func (u uf2d) find(x ufPoint) ufPoint {
	if f, ok := u[x]; ok && f != x {
		u[x] = u.find(f)
		return u[x]
	}
	return x
}
func (u uf2d) merge(from, to ufPoint) { u[u.find(from)] = u.find(to) }

// Kick Start 2019C - Wiggle Walk https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aac
func moveRobot(start ufPoint, command string) ufPoint {
	p := start
	w, n, e, s := uf2d{}, uf2d{}, uf2d{}, uf2d{}
	for _, c := range command {
		// 注意这里是矩阵
		w.merge(p, ufPoint{p.x, p.y - 1})
		n.merge(p, ufPoint{p.x - 1, p.y})
		e.merge(p, ufPoint{p.x, p.y + 1})
		s.merge(p, ufPoint{p.x + 1, p.y})
		switch c {
		case 'W':
			p = w.find(p)
		case 'N':
			p = n.find(p)
		case 'E':
			p = e.find(p)
		default:
			p = s.find(p)
		}
	}
	return p
}

// 点权并查集
// 维护的可以是集合的大小、最值、XOR、GCD 等
// https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/B
// https://codeforces.com/problemset/problem/1609/D
// LC1562 https://leetcode.cn/problems/find-latest-group-of-size-m/
// 哈希表写法 https://leetcode.cn/problems/groups-of-strings/
// https://atcoder.jp/contests/arc107/tasks/arc107_c
func _(n int) {
	groups := n
	maxSize := 1
	fa := make([]int, n) // n+1
	sz := make([]int, len(fa))
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	// newRoot = -1 表示未发生合并
	merge := func(from, to int) (newRoot int) {
		from, to = find(from), find(to)
		if from == to {
			return -1
		}
		fa[from] = to
		sz[to] += sz[from]
		//sz[from] = 0 // 有些题目需要保证总 sz 和不变（如 CF1609D）
		if sz[to] > maxSize {
			maxSize = sz[to]
		}
		groups--
		return to
	}
	same := func(x, y int) bool { return find(x) == find(y) }
	size := func(x int) int { return sz[find(x)] }

	_ = []interface{}{merge, same, size}
}

// 边权并查集（种类并查集）
// 核心在于：
//    2 ------ 4
//   /        /
//  1 ------ 3
// 如果知道 1->2 的距离和 3->4 的距离，现在告诉你 1->3 的距离
// 由于 1->3->4 和 1->2->4 的距离相等（相当于从 1 到 4 有两条路径）
// 那么就可以推出 2->4 的距离为 (1->3) + (3->4) - (1->2)
//
// https://www.bilibili.com/video/av68342657?p=2
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-11
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-12
// https://oi-wiki.org/ds/dsu/#_9
//
// 模板题 https://codeforces.com/problemset/problem/1850/H 1700
//       https://codeforces.com/problemset/problem/1074/D 2400? 1700!
//       https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/D
// 种类并查集：狼人和平民 https://codeforces.com/problemset/problem/1594/D 1700
// 种类并查集：同义词反义词 https://codeforces.com/problemset/problem/766/D 2000
// 种类并查集：食物链 https://www.luogu.com.cn/problem/P2024
// 种类并查集：不能构成二分图的第一条边 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/J
// 种类并查集 + 维护集合大小 https://codeforces.com/problemset/problem/1290/C 2400
// todo https://codeforces.com/contest/1615/problem/D
//      https://codeforces.com/contest/1713/problem/E
// 边权：https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/C
// 边权：LC399 除法求值 https://leetcode.cn/problems/evaluate-division/
//      LC2307 https://leetcode.cn/problems/check-for-contradictions-in-equations/ 也可以 DFS
// https://codeforces.com/problemset/problem/1788/F 2500
func _(n int) {
	// 注：kinds 为 2 时可以用异或来代替加减法
	const kinds = 3
	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	dis := make([]int, len(fa)) // dis[i] 表示 i 到其所在集合根节点（代表元）的距离
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			ffx := find(fa[x])
			dis[x] += dis[fa[x]] //
			fa[x] = ffx
		}
		return fa[x]
	}
	same := func(x, y int) bool { return find(x) == find(y) }
	delta := func(x, y int) int {
		if !same(x, y) { // 如果题目保证 same，则可以去掉
			return -1
		}
		find(x)
		find(y)
		return ((dis[x]-dis[y])%kinds + kinds) % kinds
	}
	merge := func(from, to int, d int) bool { // 返回是否与已知条件矛盾
		if fFrom, fTo := find(from), find(to); fFrom != fTo {
			dis[fFrom] = d + dis[to] - dis[from]
			fa[fFrom] = fTo
			return true
		}
		return delta(from, to) == d
	}

	// 统计每个集合中各个类型的个数
	cnt := make([][kinds]int, len(fa))
	for i := range fa {
		cnt[find(i)][dis[i]%kinds]++
	}

	// 离散化版本
	faMap, disMap := map[int]int{}, map[int]int{}
	find = func(x int) int {
		if fx, ok := faMap[x]; ok && fx != x {
			ffx := find(fx)
			disMap[x] += disMap[fx]
			faMap[x] = ffx
			return ffx
		}
		return x
	}

	_ = []interface{}{merge, same, delta}
}

// 可持久化并查集
// 需要关 GC：func init() { debug.SetGCPercent(-1) }
// 模板题 https://www.luogu.com.cn/problem/P3402
type pufNode struct {
	lo, ro  *pufNode
	l, r    int
	fa, dep int
}

// t := make([]*pufNode, 1, maxVersion+1)
// t[0] = buildPUF(1, n)
func buildPUF(l, r int) *pufNode {
	o := &pufNode{l: l, r: r}
	if l == r {
		o.fa = l
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPUF(l, m)
	o.ro = buildPUF(m+1, r)
	return o
}

func (o *pufNode) _find(x int) *pufNode {
	if o.l == o.r {
		return o
	}
	if m := o.lo.r; x <= m {
		return o.lo._find(x)
	}
	return o.ro._find(x)
}

func (o *pufNode) find(x int) *pufNode {
	f := o._find(x)
	if f.fa == x {
		return f
	}
	return o.find(f.fa)
}

// 注意为了拷贝一份 pufNode，这里的接收器不是指针
func (o pufNode) setFa(from, to int) *pufNode {
	if o.l == o.r {
		o.fa = to
		return &o
	}
	if m := o.lo.r; from <= m {
		o.lo = o.lo.setFa(from, to)
	} else {
		o.ro = o.ro.setFa(from, to)
	}
	return &o
}

func (o *pufNode) addDep(x int) {
	if o.l == o.r {
		o.dep++
		return
	}
	if m := o.lo.r; x <= m {
		o.lo.addDep(x)
	} else {
		o.ro.addDep(x)
	}
}

// 启发式合并：把深度小的合并到深度大的。若二者深度一样，则合并后的深度加一
func (o *pufNode) merge(x, y int) *pufNode {
	from, to := o.find(x), o.find(y)
	if from.fa == to.fa {
		return o
	}
	if from.dep > to.dep {
		from, to = to, from
	}
	p := o.setFa(from.fa, to.fa)
	if from.dep == to.dep {
		p.addDep(to.fa)
	}
	return p
}

// 可撤销并查集 回滚并查集
// 应用：动态图连通性·离线（求 CC 个数或判断 v 和 w 是否连通）
// https://en.wikipedia.org/wiki/Dynamic_connectivity
// https://codeforces.com/gym/100551/problem/A
// https://codeforces.com/edu/course/2/lesson/7/3/practice/contest/289392/problem/C
// https://loj.ac/p/121
// https://atcoder.jp/contests/abc302/tasks/abc302_h （基于 https://atcoder.jp/contests/arc111/tasks/arc111_b）
// https://www.luogu.com.cn/problem/P5631
// https://codeforces.com/contest/891/problem/C
// https://codeforces.com/problemset/problem/1217/F
// https://codeforces.com/contest/1681/problem/F
// Dynamic connectivity contest https://codeforces.com/gym/100551
func dynamicConnectivity(in io.Reader, n, q int) (ans []int) {
	if q == 0 {
		return
	}

	type edge struct{ v, w int }
	type query struct {
		e edge
		t int
		// 记 i 为 query 在 qs 中的下标
		// t > i 表示 i 时刻的加边操作对应的删边时间为 t
		// t < i 表示 i 时刻的删边操作对应的加边时间为 t
	}
	qs := make([]query, q)
	addTime := map[edge]int{}
	for i := range qs {
		var op string
		if Fscan(in, &op); op[0] == '?' {
			qs[i].t = -1 // 表示要输出答案
			continue
		}
		var v, w int
		Fscan(in, &v, &w)
		if v > w {
			v, w = w, v
		}
		e := edge{v, w}
		qs[i].e = e
		if op[0] == '+' {
			addTime[e] = i
		} else {
			addT := addTime[e]
			delete(addTime, e)
			qs[i].t = addT
			qs[addT].t = i
		}
	}
	for e, t := range addTime {
		qs[t].t = len(qs)
		qs = append(qs, query{e, t})
	}

	cc := n
	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	undo := []int{}
	find := func(x int) int {
		for x != fa[x] {
			x = fa[x]
		}
		return x
	}
	merge := func(e edge) {
		if x, y := find(e.v), find(e.w); x != y {
			if sz[x] > sz[y] {
				x, y = y, x
			}
			fa[x] = y
			sz[y] += sz[x]
			undo = append(undo, x)
			cc--
		}
	}
	rollbackTo := func(k int) {
		for len(undo) > k {
			x := undo[len(undo)-1]
			undo = undo[:len(undo)-1]
			sz[fa[x]] -= sz[x]
			fa[x] = x
			cc++
		}
	}
	var f func(l, r int)
	f = func(l, r int) {
		if l+1 == r {
			if qs[l].t < 0 {
				ans = append(ans, cc)
				//ans = append(ans, find(qs[l].e.v) == find(qs[l].e.w))
			}
			return
		}
		mid := (l + r) >> 1
		initSize := len(undo)
		// 遍历 [l,mid) 之前，连接所有在 l 时刻前连接的，且不在 mid 时刻前删除的边
		for _, q := range qs[mid:r] {
			if q.t < l {
				merge(q.e)
			}
		}
		f(l, mid)
		rollbackTo(initSize)
		// 遍历 [mid,r) 之前，连接所有在 mid 时刻前连接的，且不在 r 时刻前删除的边
		for _, q := range qs[l:mid] {
			if q.t >= r {
				merge(q.e)
			}
		}
		f(mid, r)
		rollbackTo(initSize)
	}
	f(0, len(qs))

	return
}
