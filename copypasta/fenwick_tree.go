package copypasta

// 树状数组
// 效率是线段树的 3~10 倍（由数据决定）
// https://oi-wiki.org/ds/bit/
// todo 浅谈树状数组的优化及扩展 https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
// todo 浅谈树状数组套权值树 https://www.luogu.com.cn/blog/bfqaq/qian-tan-shu-zhuang-shuo-zu-quan-zhi-shu
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FenwickTree.java.html
// 模板题 https://www.luogu.com.cn/problem/P3374
// 题目推荐 https://cp-algorithms.com/data_structures/fenwick.html#toc-tgt-12
// EXTRA: 树状数组的性质能使其支持动态 [1,r] 范围上的最值更新查询等操作 https://codeforces.ml/problemset/problem/629/D
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

	// 常数优化
	// 参考 https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
	query = func(l, r int) (s int) {
		l--
		for ; r > l; r &= r - 1 {
			s += tree[r]
		}
		for ; l > r; l &= l - 1 {
			s -= tree[l]
		}
		return
	}

	// 差分树状数组，可用于区间更新+单点查询
	// 单点查询 query(i) = a[i] + sum(i)
	// 模板题 https://www.luogu.com.cn/problem/P3368
	addRange := func(l, r int, val int) { add(l, val); add(r+1, -val) } // [l,r]

	_ = []interface{}{add, sum, query, addRange}
}

// NOTE: 也可以写成 struct 的形式
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
