package copypasta

/* 树状数组
https://oi-wiki.org/ds/bit/
todo 浅谈树状数组的优化及扩展 https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
todo 浅谈树状数组套权值树 https://www.luogu.com.cn/blog/bfqaq/qian-tan-shu-zhuang-shuo-zu-quan-zhi-shu
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FenwickTree.java.html

模板题 https://www.luogu.com.cn/problem/P3374
逆序对 https://codeforces.com/edu/course/2/lesson/4/3/practice/contest/274545/problem/A https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
经典技巧: 元素值和下标双变量的题目，转换成元素排序后对下标的操作（元素大小相等时下标大的在前）
			https://codeforces.com/problemset/problem/629/D
静态区间种类 - 离线做法 https://www.luogu.com.cn/problem/P1972 https://atcoder.jp/contests/abc174/tasks/abc174_f
题目推荐 https://cp-algorithms.com/data_structures/fenwick.html#toc-tgt-12
EXTRA: 树状数组的性质能使其支持动态 [1,r] 范围上的最值更新查询等操作 https://codeforces.com/problemset/problem/629/D
好题 https://www.luogu.com.cn/problem/P2345 https://www.luogu.com.cn/problem/P5094
https://codeforces.com/gym/101649 I
todo http://poj.org/problem?id=2155
http://poj.org/problem?id=2886
*/
func fenwickTree(n int) {
	tree := make([]int, n+1)
	add := func(i int, val int) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	query := func(l, r int) int { return sum(r) - sum(l-1) } // [l,r]

	// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
	// r+1 即使超过 n 也没关系，因为不会用到
	// 模板题 https://www.luogu.com.cn/problem/P3368
	addRange := func(l, r int, val int) { add(l, val); add(r+1, -val) } // [l,r]

	// 常数优化：O(n) 建树 https://oi-wiki.org/ds/fenwick/#tricks
	init := func(a []int) {
		n := len(a)
		for i := 1; i <= n; i++ {
			tree[i] += a[i-1]
			if j := i + i&-i; j <= n {
				tree[j] += tree[i]
			}
		}
	}

	// 常数优化（不推荐。实测只快了几毫秒）https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
	query = func(l, r int) (s int) {
		if l > r {
			panic(9)
		}
		l--
		for ; r > l; r &= r - 1 {
			s += tree[r]
		}
		for ; l > r; l &= l - 1 {
			s -= tree[l]
		}
		return
	}

	// 求逆序对的方法之一
	cntInversions := func(a []int) int64 {
		n := len(a)
		tree := [1e5 + 1]int{} // 注：如果 a 范围较大则需要离散化
		add := func(i int) {
			for ; i <= n; i += i & -i {
				tree[i]++
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res += tree[i]
			}
			return
		}
		cnt := int64(0)
		for i, v := range a {
			cnt += int64(i - sum(v))
			add(v)
		}
		return cnt
	}

	_ = []interface{}{
		add, sum, query, addRange, init,
		cntInversions,
	}
}

// 也可以写成后面的 struct 形式
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

type fenwick struct {
	tree []int // int64
}

func newFenwickTree(n int) fenwick {
	return fenwick{make([]int, n+1)}
}
func (f fenwick) add(i int, val int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i] += val
	}
}
func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}
func (f fenwick) query(l, r int) (res int) { // [l,r]
	return f.sum(r) - f.sum(l-1)
}
