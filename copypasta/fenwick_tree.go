package copypasta

// 效率是线段树的 3~10 倍（由数据决定）
func binaryIndexedTree(n int) {
	tree := make([]int, n+1)
	add := func(idx int, val int) {
		for ; idx <= n; idx += idx & -idx { // idx += lowbit(idx)
			tree[idx] += val
		}
	}
	addRange := func(l, r int, val int) { // [l,r]
		add(l, val)
		add(r+1, -val)
		// now value at i is a[i] + sum(i)
	}
	sum := func(idx int) (res int) {
		for ; idx > 0; idx &= idx - 1 { // idx -= lowbit(idx)
			res += tree[idx]
		}
		return
	}
	query := func(l, r int) int { // [l,r]
		return sum(r) - sum(l-1)
	}

	_ = []interface{}{add, addRange, query}
}

func multiBinaryIndexedTrees(n, m int) {
	trees := make([][]int, m)
	for i := range trees {
		trees[i] = make([]int, n+1)
	}
	add := func(tree []int, idx int, val int) {
		for ; idx <= n; idx += idx & -idx {
			tree[idx] += val
		}
	}
	sum := func(tree []int, idx int) (res int) {
		for ; idx > 0; idx &= idx - 1 {
			res += tree[idx]
		}
		return
	}
	query := func(tree []int, l, r int) int {
		return sum(tree, r) - sum(tree, l-1)
	}

	_ = []interface{}{add, query}
}
