package copypasta

func unionFindCollections() {
	n := int(1e5)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }

	_ = []interface{}{merge}

	// TODO: 带权并查集
}
