package copypasta

// 树状数组
// 效率是线段树的 3~10 倍（由数据决定）
// https://oi-wiki.org/ds/bit/
// 模板题 https://www.luogu.com.cn/problem/P3374
// 题目推荐 https://cp-algorithms.com/data_structures/fenwick.html#toc-tgt-12

func fenwickTree(n int) {
	tree := make([]int, n+1)
	add := func(i int, val int) {
		for ; i <= n; i += i & -i { // i += lowbit(i)
			tree[i] += val
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 { // i -= lowbit(i)
			res += tree[i]
		}
		return
	}
	query := func(l, r int) int { return sum(r) - sum(l-1) } // [l,r]

	// 差分树状数组，可用于区间更新
	// 模板题 https://www.luogu.com.cn/problem/P3368
	addRange := func(l, r int, val int) { // [l,r]
		add(l, val)
		add(r+1, -val)
		// now value at i is a[i] + sum(i)
	}

	_ = []interface{}{add, sum, query, addRange}
}

func multiFenwickTree(m, n int) {
	trees := make([][]int, m)
	for i := range trees {
		trees[i] = make([]int, n+1)
	}
	add := func(tree []int, i int, val int) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	sum := func(tree []int, i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	query := func(tree []int, l, r int) int { return sum(tree, r) - sum(tree, l-1) }

	_ = []interface{}{add, sum, query}
}
