package main

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) (ans bool) {
	fa := make([]int, n)
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
	countRoots := func() (cnt int) {
		for i := range fa {
			if find(i) == i {
				cnt++
			}
		}
		return
	}

	for v, w := range leftChild {
		if w != -1 {
			if same(v, w) {
				return
			}
			merge(v, w)
		}
		if w = rightChild[v]; w != -1 {
			if same(v, w) {
				return
			}
			merge(v, w)
		}
	}
	return countRoots() == 1
}
