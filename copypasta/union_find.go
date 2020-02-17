package copypasta

// https://oi-wiki.org/ds/dsu/
// https://cp-algorithms.com/data_structures/disjoint_set_union.html

func unionFind() {
	// NOTE: 离散化时，可以改用 map[int]int
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n) // n+1
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

	// EXTRA
	mergeCheck := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return true
		}
		fa[x] = y
		return false
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
	getRoots := func() (roots []int) {
		for i := range fa {
			if find(i) == i {
				roots = append(roots, i)
			}
		}
		return
	}
	countRoots := func() (cnt int) { // 连通分量个数
		for i := range fa {
			if find(i) == i {
				cnt++
			}
		}
		return
	}
	getComps := func() (comps map[int][]int) { // 连通分量
		comps = map[int][]int{}
		for i := range fa {
			f := find(i)
			comps[f] = append(comps[f], i)
		}
		return
	}

	_ = []interface{}{
		initFa, merge, same,
		mergeCheck, mergeRangeTo, getRoots, countRoots, getComps,
	}
}

// https://oi-wiki.org/ds/dsu/#_9
// TODO: 一般化
// TODO: https://codeforces.com/contest/1074/problem/D
func unionFindWithMaintain() {
	var fa, size []int
	initFa := func(n int) {
		fa = make([]int, n)   // n+1
		size = make([]int, n) // n+1
		for i := range fa {
			fa[i] = i
			size[i] = 1
		}
	}
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
			fa[from] = to
			size[to] += size[from]
		}
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	_ = []interface{}{initFa, merge, same}
}

// 二维并查集
func multiUnionFind(n, m int) {
	fas := make([][]int, m)
	for i := range fas {
		fas[i] = make([]int, n+1)
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

	_ = []interface{}{merge, same}
}
