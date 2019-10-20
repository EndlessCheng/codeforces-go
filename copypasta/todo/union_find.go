package copypasta

func unionFind() {
	n := int(1e5)
	// NOTE: 离散化时，可以改用 map[int]int
	fa := make([]int, n+1)
	//size := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		//size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	//merge := func(from, to int) {
	//	from, to := find(from), find(to)
	//	if from != to {
	//		fa[from] = to
	//		size[to] += size[from]
	//	}
	//}
	same := func(x, y int) bool { return find(x) == find(y) }

	_ = []interface{}{merge, same}
}

func multiUnionFind() {
	n, m := int(1e5), 2
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

func unionFindWithWeights() {
	//n := int(1e5)
	//type pair struct{ i, sum int }
	//fa := make([]pair, n)
	//for i := range fa {
	//	fa[i] = pair{i, 1} // or read sum from stdin
	//}
	//var find func(int) pair
	//find = func(x int) pair {
	//	if pfa := fa[x]; pfa.i != x {
	//		fa[x] = find(pfa.i)
	//		fa[x].sum += pfa.sum
	//	}
	//	return fa[x]
	//}
	//merge := func(from, to int, val int) bool {
	//	pf, pt := find(from), find(to)
	//	if pf.i == pt.i {
	//		return false
	//	} else {
	//		fa[pf] = o2
	//		// *custom*
	//	}
	//}
	//query := func(i, j int) (sum int, ok bool) {
	//	pi, pj := find(i), find(j)
	//	if
	//}
	//
	//_ = []interface{}{merge}
}
