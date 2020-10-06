package copypasta

import (
	. "fmt"
	"io"
)

/* 套题
https://blog.csdn.net/weixin_43914593/article/details/104108049 算法竞赛专题解析（3）：并查集
*/

// 普通并查集
// https://oi-wiki.org/ds/dsu/
// https://cp-algorithms.com/data_structures/disjoint_set_union.html
// 并查集时间复杂度证明 https://oi-wiki.org/ds/dsu-complexity/
// 模板题 https://www.luogu.com.cn/problem/P3367
// 思维转换题! https://nanti.jisuanke.com/t/43488
// 使某些点不在环上需要删除的最少边数 https://ac.nowcoder.com/acm/contest/7780/C
// https://codeforces.com/problemset/problem/292/D
// 任意合并+区间合并 https://codeforces.com/problemset/problem/566/D
func unionFind(n int) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	initFa(n + 1) //
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

	// 离散化版本
	faMap := map[int]int{}
	find = func(x int) int {
		if fx, ok := faMap[x]; ok && fx != x {
			faMap[x] = find(fx)
			return faMap[x]
		}
		return x
	}

	// merge，并返回是否有新的 merge
	sameMerge := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		fa[x] = y
		return true
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
		rank := make([]int, n)
		merge := func(x, y int) {
			x = find(x)
			y = find(y)
			if x == y {
				return
			}
			if rank[x] < rank[y] {
				fa[x] = y
			} else {
				fa[y] = x
				if rank[x] == rank[y] {
					rank[x]++
				}
			}
		}
		_ = merge
	}

	_ = []interface{}{
		initFa, merge, same,
		mergeBig, sameMerge, mergeRangeTo, getRoots, countRoots, getComps,
	}
}

// 并查集 - 维护点权
// 维护的可以是集合的大小、最值、XOR、GCD 等
// https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/B
// LC 周赛 203D https://leetcode-cn.com/contest/weekly-contest-203/problems/find-latest-group-of-size-m/
func unionFindVertexWeight(n int) {
	var fa, size []int
	initFa := func(n int) {
		fa = make([]int, n)
		size = make([]int, n)
		for i := range fa {
			fa[i] = i
			size[i] = 1
		}
	}
	initFa(n + 1) //
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from != to {
			size[to] += size[from]
			fa[from] = to
		}
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	_ = []interface{}{initFa, merge, same}
}

// 并查集 - 维护边权（种类）
// 简单易懂的讲解：https://www.bilibili.com/video/av68342657?p=2
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-11
// https://cp-algorithms.com/data_structures/disjoint_set_union.html#toc-tgt-12
// https://oi-wiki.org/ds/dsu/#_9
// 模板题 https://codeforces.com/problemset/problem/1074/D https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/D
// 种类并查集：同义词反义词 https://codeforces.com/problemset/problem/766/D
// 种类并查集：食物链 https://www.luogu.com.cn/problem/P2024
// 种类并查集：不能构成二分图的第一条边 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/J
// 边权：https://codeforces.com/edu/course/2/lesson/7/1/practice/contest/289390/problem/C
func unionFindEdgeWeight(n int) {
	const kinds = 2
	var fa, dis []int // dis 表示点到其所在集合根节点（代表元）的距离
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		dis = make([]int, n)
	}
	initFa(n + 1) //
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			ffx := find(fa[x])
			dis[x] += dis[fa[x]]
			fa[x] = ffx
		}
		return fa[x]
	}
	// 调用前需要：1. 调用 find(x); find(y)  2. 保证 same(x, y) 为 true
	delta := func(x, y int) int { return ((dis[x]-dis[y])%kinds + kinds) % kinds }
	merge := func(from, to int, d int) bool { // 返回是否与已知条件矛盾
		if fFrom, fTo := find(from), find(to); fFrom != fTo {
			dis[fFrom] = d + dis[to] - dis[from]
			fa[fFrom] = fTo
			return true
		}
		return delta(from, to) == d
	}
	same := func(x, y int) bool { return find(x) == find(y) }

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

	_ = []interface{}{initFa, merge, same, delta}
}

// 并查集组（一般用于涉及到位运算的题目）
// 也可以写成后面的 struct 形式
func multiUnionFind(n, m int) {
	fas := make([][]int, m)
	for i := range fas {
		fas[i] = make([]int, n) // n+1
		for j := range fas[i] {
			fas[i][j] = j
		}
	}
	var find func([]int, int) int
	find = func(fa []int, x int) int {
		if fa[x] != x {
			fa[x] = find(fa, fa[x])
		}
		return fa[x]
	}
	merge := func(fa []int, from, to int) { fa[find(fa, from)] = find(fa, to) }
	same := func(fa []int, x, y int) bool { return find(fa, x) == find(fa, y) }
	mergeRange := func(fa []int, l, r int) { // 左闭右开
		for i := find(fa, l); i < r; i = find(fa, i+1) {
			fa[i] = r // merge 到 r 上
		}
	}

	_ = []interface{}{merge, same, mergeRange}
}

type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}
func newUnionFinds(m, n int) []uf {
	us := make([]uf, m)
	for i := range us {
		us[i] = newUnionFind(n)
	}
	return us
}
func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}
func (u uf) merge(from, to int) (isNewMerge bool) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u.fa[x] = y
	return true
}
func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }
func (u uf) countRoots(st int) (cnt int) { // st = 0 or 1 ...
	for i := st; i < len(u.fa); i++ {
		if u.find(i) == i {
			cnt++
		}
	}
	return
}

// 可持久化并查集
// todo 模板题 https://www.luogu.com.cn/problem/P3402

// 动态图连通性（求 CC 个数或判断 v 和 w 是否连通）
// https://en.wikipedia.org/wiki/Dynamic_connectivity
// https://codeforces.com/gym/100551/problem/A https://codeforces.com/edu/course/2/lesson/7/3/practice/contest/289392/problem/C
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
				ans = append(ans, cc) // Fprintln(out, cc)
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
