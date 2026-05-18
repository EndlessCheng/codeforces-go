package copypasta

import (
	"fmt"
	"math/bits"
)

/* 稀疏表 Sparse Table
st[i][j] 对应的区间是 [i, i+2^j)
https://oi-wiki.org/ds/sparse-table/
https://codeforces.com/blog/entry/66643
扩展：Tarjan RMQ https://codeforces.com/blog/entry/48994
一些 RMQ 的性能对比 https://codeforces.com/blog/entry/78931
一个 RMQ 问题的快速算法，以及区间众数 https://zhuanlan.zhihu.com/p/79423299
将 LCA、RMQ、LA 优化至理论最优复杂度 https://www.luogu.com.cn/blog/ICANTAKIOI/yi-shang-shou-ke-ji-jiang-lcarmqla-you-hua-zhi-zui-you-fu-za-du
RMQ 标准算法和线性树上并查集 https://ljt12138.blog.uoj.ac/blog/4874
随机 RMQ https://www.luogu.com.cn/problem/P3793
todo O(n)-O(1) lca/rmq, not method of 4 russians https://codeforces.com/blog/entry/125371
todo O(n)-O(1) RMQ https://atcoder.jp/contests/arc165/submissions/45673031

模板题 https://www.luogu.com.cn/problem/P3865
模板题 https://www.luogu.com.cn/problem/P2880
模板题 https://www.luogu.com.cn/problem/P1816
https://codeforces.com/problemset/problem/1709/D 1700
https://codeforces.com/problemset/problem/2050/F 1700 GCD
https://codeforces.com/problemset/problem/1548/B 1800 GCD
https://codeforces.com/problemset/problem/689/D 2100 二分/三指针
https://codeforces.com/problemset/problem/1107/G 2500
https://www.jisuanke.com/contest/11346/challenges 变长/种类
todo https://ac.nowcoder.com/acm/problem/240870 https://ac.nowcoder.com/acm/contest/view-submission?submissionId=53616019

题单 https://cp-algorithms.com/data_structures/sparse-table.html#toc-tgt-5
*/

type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// 时间复杂度 O(n * log n)
func newSparseTable[T any](a []T, op func(T, T) T) sparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], a)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

// [l,r) 左闭右开，下标从 0 开始
// 返回 op(nums[l:r])
// 时间复杂度 O(1)
func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

// 使用方法举例
func sparseTableExample() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	st := newSparseTable(nums, func(a, b int) int { return max(a, b) })
	fmt.Println(st.query(0, 5)) // 5
	fmt.Println(st.query(1, 1)) // 错误：必须保证 l < r
}

//

// 下标版本，查询返回的是区间最值的下标
// https://codeforces.com/problemset/problem/675/E
// - 此题另一种做法是单调栈二分，见 https://www.luogu.com.cn/problem/solution/CF675E
type stPair struct{ v, i int }
type sparseTableWithIndex [][]stPair

func newSparseTableWithIndex(a []int) sparseTableWithIndex {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(sparseTableWithIndex, n)
	for i, v := range a {
		st[i] = make([]stPair, sz)
		st[i][0] = stPair{v, i}
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			if a, b := st[i][j-1], st[i+1<<(j-1)][j-1]; a.v <= b.v { // 最小值，相等时下标取左侧
				st[i][j] = a
			} else {
				st[i][j] = b
			}
		}
	}
	return st
}

// 查询区间 [l,r)，注意 l 和 r 是从 0 开始算的
func (st sparseTableWithIndex) query(l, r int) int {
	k := bits.Len32(uint32(r-l)) - 1
	a, b := st[l][k], st[r-1<<k][k]
	if a.v <= b.v { // 最小值，相等时下标取左侧
		return a.i
	}
	return b.i
}

//

// 不相交 ST 表（Disjoint Sparse Table，DST）
// 可以用来查询区间矩阵乘法、最大子段和等，这种不能相交的运算
// 国内算法竞赛圈称其为「猫树」 https://oi-wiki.org/ds/cat-tree/
// 类似算法：双栈滑动窗口
// https://codeforces.com/edu/course/3/lesson/18/3
// https://codeforces.com/edu/course/3/lesson/18/3/practice/contest/619579/problem/A
type disjointSparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

func newDisjointSparseTable[T any](a []T, op func(T, T) T) disjointSparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], a)
	for k := 1; k < w; k++ {
		B := 1 << (k + 1)
		for m := 0; m < n; m += B {
			mid := min(m+1<<k, n)

			// 左半算后缀 op
			st[k][mid-1] = a[mid-1]
			for i := mid - 2; i >= m; i-- {
				st[k][i] = op(a[i], st[k][i+1])
			}

			// 右半算前缀 op
			if mid < n {
				end := min(m+B, n)
				st[k][mid] = a[mid]
				for i := mid + 1; i < end; i++ {
					st[k][i] = op(st[k][i-1], a[i])
				}
			}
		}
	}
	return disjointSparseTable[T]{st, op}
}

// [l,r] 闭区间，下标从 0 开始
func (s disjointSparseTable[T]) query(l, r int) T {
	if l > r {
		panic("入参不合法：l > r")
	}
	if l == r {
		return s.st[0][l] // % mod
	}
	k := bits.Len(uint(l^r)) - 1
	return s.op(s.st[k][l], s.st[k][r])
}

//

// 二维 ST 表
// https://blog.nowcoder.net/n/3eccd1386a8846398d3bee62b485309b
// https://codeforces.com/problemset/problem/1301/E 2500
// https://leetcode.cn/problems/largest-local-values-in-a-matrix-ii/
type sparseTable2D[T any] struct {
	// st[k1][k2][n][m] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
	st [][][][]T
	op func(T, T) T
}

// 时间复杂度 O(n * m * log n * log m)
func newSparseTable2D[T any](a [][]T, op func(T, T) T) sparseTable2D[T] {
	n, m := len(a), len(a[0])
	wn, wm := bits.Len(uint(n)), bits.Len(uint(m))

	st := make([][][][]T, wn)
	for k1 := range st {
		st[k1] = make([][][]T, wm)
		for k2 := range st[k1] {
			st[k1][k2] = make([][]T, n)
			for i := range st[k1][k2] {
				st[k1][k2][i] = make([]T, m)
			}
		}
	}

	// 初始值
	st[0][0] = a // 如果后面会修改 a，这里要 clone

	// 单独计算 k1 = 0
	for k2 := 1; k2 < wm; k2++ {
		for i := range n {
			for j := range m - 1<<k2 + 1 {
				st[0][k2][i][j] = op(st[0][k2-1][i][j], st[0][k2-1][i][j+1<<(k2-1)])
			}
		}
	}

	for k1 := 1; k1 < wn; k1++ {
		for k2 := range wm {
			for i := range n - 1<<k1 + 1 {
				for j := range m - 1<<k2 + 1 {
					st[k1][k2][i][j] = op(st[k1-1][k2][i][j], st[k1-1][k2][i+1<<(k1-1)][j])
				}
			}
		}
	}

	return sparseTable2D[T]{st, op}
}

// 返回子矩阵最大值
// 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)，下标从 0 开始
// 时间复杂度 O(1)
func (s sparseTable2D[T]) query(r1, c1, r2, c2 int) T {
	//r1 = max(r1, 0)
	//c1 = max(c1, 0)
	//r2 = min(r2, len(s.st[0][0]))
	//c2 = min(c2, len(s.st[0][0][0]))
	k1 := bits.Len(uint(r2-r1)) - 1
	k2 := bits.Len(uint(c2-c1)) - 1
	// 视作四个子矩阵的并集
	return s.op(
		s.op(s.st[k1][k2][r1][c1], s.st[k1][k2][r2-1<<k1][c1]),
		s.op(s.st[k1][k2][r1][c2-1<<k2], s.st[k1][k2][r2-1<<k1][c2-1<<k2]),
	)
}

// 使用方法举例
func sparseTable2DExample() {
	a := [][]int{
		{3, 1, 4},
		{1, 5, 9},
	}
	st := newSparseTable2D(a, func(a, b int) int { return max(a, b) })
	// 注意是左闭右开
	fmt.Println(st.query(0, 1, 2, 3)) // 9
}

// 

// 也可以用线段树套 ST 表实现，如果 N*M ~ Q，那么总体复杂度是 O(Q log Q)，比二维 ST 表少一个 log
type segST []sparseTable[int]

func newSegSt(a [][]int) segST {
	n := len(a)
	// 线段树每个节点 [l, r] 保存的是，当上下边界固定为 l 和 r 时，把每一列的最大值视作一个 int，这 m 个数的一维 ST 表
	t := make(segST, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func (t segST) build(a [][]int, node, l, r int) {
	if l == r { // 叶子
		t[node] = newSparseTable(a[l], func(a, b int) int { return max(a, b) })
		return
	}

	m := (l + r) / 2
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树

	merged := make([]int, len(a[0]))
	p, q := t[node*2].st[0], t[node*2+1].st[0]
	for i, x := range p {
		merged[i] = max(x, q[i]) // 行号 [l, r] 中的第 i 列的最大值
	}
	t[node] = newSparseTable(merged, func(a, b int) int { return max(a, b) })
}

// 行号闭区间 [r1, r2]，列号左闭右开 [c1, c2)，下标从 0 开始
// t.query(1, 0, n-1, r1, r2, c1, c2)
func (t segST) query(node, l, r, r1, r2, c1, c2 int) int {
	if r1 <= l && r <= r2 { // 当前子树完全在 [r1, r2] 内
		return t[node].query(c1, c2)
	}
	m := (l + r) / 2
	if r2 <= m { // [r1, r2] 在左子树
		return t.query(node*2, l, m, r1, r2, c1, c2)
	}
	if r1 > m { // [r1, r2] 在右子树
		return t.query(node*2+1, m+1, r, r1, r2, c1, c2)
	}
	return max(t.query(node*2, l, m, r1, r2, c1, c2), t.query(node*2+1, m+1, r, r1, r2, c1, c2))
}
