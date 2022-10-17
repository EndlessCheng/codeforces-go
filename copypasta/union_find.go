package copypasta

import (
	. "fmt"
	"io"
)

/* 并查集
只有路径压缩的并查集复杂度是 O(nlogn) 的，这也是大多数情况下的实现方案
只有启发式合并（按深度合并）的并查集的复杂度也是 O(nlogn) 的，适用于可持久化的场景

具体的时间复杂度证明见《算法导论》
https://zhuanlan.zhihu.com/p/553192435

随机合并下的时间复杂度 https://www.cis.upenn.edu/~sanjeev/papers/soda14_disjoint_set_union.pdf
*/

// 普通并查集
// 可视化 https://visualgo.net/zh/ufds
// https://oi-wiki.org/ds/dsu/
// https://cp-algorithms.com/data_structures/disjoint_set_union.html
// 并查集时间复杂度证明 https://oi-wiki.org/ds/dsu-complexity/
//
// 模板题 https://www.luogu.com.cn/problem/P3367
// 接水问题 https://codeforces.com/problemset/problem/371/D
// 三维接雨水 https://www.luogu.com.cn/problem/P5930 LC407 https://leetcode-cn.com/problems/trapping-rain-water-ii/
// 使某些点不在环上需要删除的最少边数 https://ac.nowcoder.com/acm/contest/7780/C
// todo https://codeforces.com/problemset/problem/292/D
// 任意合并+区间合并 https://codeforces.com/problemset/problem/566/D
// 动态加点 https://codeforces.com/contest/1494/problem/D
// 思维转换 https://nanti.jisuanke.com/t/43488
//         https://codeforces.com/problemset/problem/1012/B
//         https://codeforces.com/problemset/problem/1466/F
// 前缀和 后缀和 https://codeforces.com/problemset/problem/292/D
// 维护树或基环树 https://codeforces.com/problemset/problem/859/E
// 求矩阵的 rank 矩阵 https://codeforces.com/problemset/problem/650/C LC1632/周赛212D https://leetcode-cn.com/problems/rank-transform-of-a-matrix/submissions/
// 分组排序套路 LC1998/周赛257D https://leetcode-cn.com/problems/gcd-sort-of-an-array/
// 套题 https://blog.csdn.net/weixin_43914593/article/details/104108049 算法竞赛专题解析（3）：并查集
// [1700] 转换 https://codeforces.com/problemset/problem/1253/D
// 离散 + 四方向 https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aac#analysis
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

func (u UnionFind) Find(x int) int {
	if u.Fa[x] != x {
		u.Fa[x] = u.Find(u.Fa[x])
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
		// 离散化版本
		// https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column/
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

	mergeRangeTo := func(l, r, to int) { // 常用：to=r+1，这时建议用左闭右开表示区间
		//if l < 0 {
		//	l = 0
		//}
		//if r > n {
		//	r = n
		//}
		for i := find(l); i <= r; i = find(i + 1) { // initFa 需要开 n+1 空间
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

	_ = []interface{}{merge, same, mergeBig, mergeRangeTo, getRoots, countRoots, getComps}
}

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
		case 'W': p = w.find(p)
		case 'N': p = n.find(p)
		case 'E': p = e.find(p)
		default:  p = s.find(p)
		}
	}
	return p
}

// 并查集 - 维护点权
// 维护的可以是集合的大小、最值、XOR、GCD 等
// https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/B
// https://codeforces.com/problemset/problem/1609/D
// LC 周赛 203D https://leetcode-cn.com/contest/weekly-contest-203/problems/find-latest-group-of-size-m/
// 哈希表写法 https://leetcode-cn.com/problems/groups-of-strings/
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

// 并查集 - 维护边权（种类）
// 简单易懂的讲解：https://www.bilibili.com/video/av68342657?p=2
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-11
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-12
// https://oi-wiki.org/ds/dsu/#_9
// 模板题 https://codeforces.com/problemset/problem/1074/D https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/D
// 种类并查集：同义词反义词 https://codeforces.com/problemset/problem/766/D
// 种类并查集：狼人和平民 https://codeforces.com/problemset/problem/1594/D
// 种类并查集：食物链 https://www.luogu.com.cn/problem/P2024
// 种类并查集：不能构成二分图的第一条边 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/J
// 种类并查集 + 维护集合大小 https://codeforces.com/problemset/problem/1290/C
// todo https://codeforces.com/contest/1615/problem/D
//      https://codeforces.com/contest/1713/problem/E
// 边权：https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/C
// 边权：LC399 除法求值 https://leetcode-cn.com/problems/evaluate-division/
func _(n int) {
	// 注：kinds 为 2 时可以用异或来代替加减法
	const kinds = 2
	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	dis := make([]int, len(fa)) // dis[i] 表示 i 到其所在集合根节点（代表元）的距离
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			ffx := find(fa[x])
			dis[x] += dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	// 调用前需要保证 same(x, y) 为 true
	delta := func(x, y int) int {
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
	same := func(x, y int) bool { return find(x) == find(y) }

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
// todo https://codeforces.com/contest/891/problem/C
//  https://codeforces.com/contest/1681/problem/F
// todo Dynamic connectivity contest https://codeforces.com/gym/100551
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
