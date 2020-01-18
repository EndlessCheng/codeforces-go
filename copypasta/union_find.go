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

	_ = []interface{}{initFa, merge, same, getRoots, countRoots}
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
