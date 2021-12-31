package copypasta

import "sort"

/* 树状数组（Fenwick Tree），二叉索引树（Binary Index Tree, BIT）
https://en.wikipedia.org/wiki/Fenwick_tree
树状数组的基本用途是维护序列的前缀和
tree[i] = a[i-lowbit(i)+1] + ... + a[i]
可以看做是一个删去了右儿子的线段树

推荐阅读《算法竞赛进阶指南》0x42 节
https://oi-wiki.org/ds/bit/
todo 树状数组延申应用 https://www.luogu.com.cn/blog/kingxbz/shu-zhuang-shuo-zu-zong-ru-men-dao-ru-fen
 浅谈树状数组的优化及扩展 https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
 浅谈树状数组套权值树 https://www.luogu.com.cn/blog/bfqaq/qian-tan-shu-zhuang-shuo-zu-quan-zhi-shu
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FenwickTree.java.html

模板题 https://www.luogu.com.cn/problem/P3374
逆序对 https://codeforces.com/edu/course/2/lesson/4/3/practice/contest/274545/problem/A https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
经典技巧: 元素值和下标双变量的题目，转换成元素排序后对下标的操作（元素大小相等时下标大的在前）
    https://codeforces.com/problemset/problem/629/D
静态区间种类 - 离线做法
    https://www.luogu.com.cn/problem/P1972
    https://atcoder.jp/contests/abc174/tasks/abc174_f
    https://codeforces.com/problemset/problem/246/E
题目推荐 https://cp-algorithms.com/data_structures/fenwick.html#toc-tgt-12
EXTRA: 树状数组的性质能使其支持动态 [1,r] 范围上的最值更新查询等操作 https://codeforces.com/problemset/problem/629/D
好题 https://www.luogu.com.cn/problem/P2345 https://www.luogu.com.cn/problem/P5094
多变量统计 https://codeforces.com/problemset/problem/1194/E
三元逆序对 https://codeforces.com/problemset/problem/61/E
整除对统计 https://codeforces.com/problemset/problem/301/D
区间统计技巧 https://codeforces.com/problemset/problem/369/E
区间包含计数 https://codeforces.com/problemset/problem/652/D
todo https://codeforces.com/problemset/problem/961/E（不止一种做法）
 https://codeforces.com/gym/101649 I 题
 http://poj.org/problem?id=2155
 http://poj.org/problem?id=2886
*/
func fenwickTree(n int) {
	tree := make([]int, n+1) // int64
	add := func(i int, val int) {
		for ; i < len(tree); i += i & -i {
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

	// 求权值树状数组第 k 小的数（k > 0）
	// 这里 tree[i] 表示 i 的个数
	// 返回最小的 x 满足 ∑i=[1..x] tree[i] >= k
	// 思路类似倍增的查询，不断寻找 ∑<k 的数，最后 +1 就是答案
	// https://oi-wiki.org/ds/fenwick/#tricks
	// https://codeforces.com/blog/entry/61364
	// https://codeforces.com/problemset/problem/1404/C
	// todo https://codeforces.com/contest/992/problem/E
	// 二分 https://www.luogu.com.cn/problem/P4137
	// - 代码见下面的 rangeMex
	kth := func(k int) (res int) {
		const log = 17 // bits.Len(uint(n))
		for b := 1 << (log - 1); b > 0; b >>= 1 {
			if next := res | b; next < len(tree) && k > tree[next] {
				k -= tree[next]
				res = next
			}
		}
		return res + 1
	}

	// 常数优化：O(n) 建树
	// https://oi-wiki.org/ds/fenwick/#tricks
	init := func(a []int) { // len(tree) = len(a) + 1
		for i := 1; i < len(tree); i++ {
			tree[i] += a[i-1]
			if j := i + i&-i; j < len(tree) {
				tree[j] += tree[i]
			}
		}
	}

	// 常数优化（不推荐。实测只快了几毫秒）
	// https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
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
	// 如果 a 范围较大则需要离散化（但这样还不如直接用归并排序）
	// 归并做法见 misc.go 中的 mergeCount
	// 扩展 https://codeforces.com/problemset/problem/362/C
	// 环形最小逆序对 https://www.luogu.com.cn/problem/solution/P2995
	// 扩展：某些位置上的数待定时的逆序对的期望值 https://codeforces.com/problemset/problem/1096/F
	// https://codeforces.com/problemset/problem/1585/D
	cntInversions := func(a []int) int64 {
		n := len(a)
		tree := make([]int, n+1)
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
		invCnt := int64(0)
		for i, v := range a {
			// 由于 i 从 0 开始算，这里先 sum 后 add
			invCnt += int64(i - sum(v))
			add(v)
		}
		return invCnt
	}

	_ = []interface{}{add, sum, query, addRange, kth, init, cntInversions}
}

// 给一个数组 a 和一些询问 qs，对每个询问计算 mex(a[l..r])
// a[i]>=0, 1<=l<=r<=n
// 遍历数组 a，记录 a[i] 最后一次出现的位置 lastPos 以及上一个 a[i] 的位置 prevPos
// 建立一个权值树状数组，维护 lastPos[v] 的前缀最小值
// 树状数组维护前缀最小值的条件是每次修改只能往小改，那么从后往前做就好了
// 将询问离线：按照右端点排序（或分组），计算 mex。原理见代码中 query 的注释
// https://www.luogu.com.cn/problem/P4137
// LC2003/周赛258D https://leetcode-cn.com/problems/smallest-missing-genetic-value-in-each-subtree/
// - 需要将 a 转换成 DFS 序且从 0 开始，同时最终答案需要 +1
func rangeMex(a []int, qs []struct{ l, r, i int }, min func(int, int) int) []int {
	const mx int = 1e5 + 2
	// 权值树状数组
	// 这里 tree[v] = min{pos[v-lowbit(v)+1], ..., pos[v]}
	tree := [mx]int{}
	for i := range tree {
		tree[i] = 1e9
	}
	// 由于树状数组的下标需要为正，将所有 v 偏移 +1
	update := func(v, pos int) {
		for v++; v < mx; v += v & -v {
			tree[v] = min(tree[v], pos)
		}
	}
	// 根据上面的定义，对于第一个满足 if 条件的 next，有 min{pos[1], ..., pos[next]} >= l，即 mex >= next（这里的 1~next 是偏移 +1 后的）
	// 后面满足 if 的以此类推
	query := func(l int) (res int) {
		const log = 17 // bits.Len(uint(mx))
		for b := 1 << (log - 1); b > 0; b >>= 1 {
			if next := res | b; next < mx && tree[next] >= l {
				res = next
			}
		}
		return
	}

	n, m := len(a), len(qs)
	prevPos := make([]int, n)
	lastPos := make([]int, mx)
	for i, v := range a {
		prevPos[i] = lastPos[v]
		lastPos[v] = i + 1
	}
	for v, pos := range lastPos {
		update(v, pos)
	}

	ans := make([]int, m)
	sort.Slice(qs, func(i, j int) bool { return qs[i].r > qs[j].r })
	for i, qi := n-1, 0; i >= 0; i-- {
		for ; qi < m && qs[qi].r == i+1; qi++ {
			ans[qs[qi].i] = query(qs[qi].l)
		}
		update(a[i], prevPos[i])
	}
	return ans
}

// 结构体写法
type fenwick struct {
	tree []int64
}

func newFenwickTree(n int) fenwick {
	return fenwick{make([]int64, n+1)}
}

// 位置 i 增加 val
// 1<=i<=n
func (f fenwick) add(i int, val int64) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i] += val
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

// 求区间和 [l,r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int64 {
	return f.sum(r) - f.sum(l-1)
}
