package copypasta

// 效率是线段树的 3~10 倍（由数据决定）
func binaryIndexedTree() {
	lowbit := func(n int) int { return n & -n }

	n := int(1e5)
	tree := make([]int, n+1)
	add := func(idx int, val int) {
		for ; idx <= n; idx += lowbit(idx) {
			tree[idx] += val
		}
	}
	addRange := func(l, r int, val int) { // [l,r]
		add(l, val)
		add(r+1, -val)
		// now value at i is a[i] + sum(i)
	}
	sum := func(idx int) (res int) {
		for ; idx > 0; idx &= idx - 1 {
			res += tree[idx]
		}
		return
	}
	query := func(l, r int) int { // [l,r]
		return sum(r) - sum(l-1)
	}

	_ = []interface{}{add, addRange, query}
}
