package main

import "math/bits"

// https://space.bilibili.com/206214
type sparseTable2D[T any] struct {
	// st[k1][k2][n][m] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
	st [][][][]T
	op func(T, T) T
}

// 时间复杂度 O(n * m * log n * log m)
func newSparseTable2D[T any](a [][]T, op func(T, T) T) sparseTable2D[T] {
	n, m := len(a), len(a[0])
	wn, wm := bits.Len(uint(n)), bits.Len(uint(m))

	// st[k1][k2][n][m] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
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
	r1 = max(r1, 0)
	c1 = max(c1, 0)
	r2 = min(r2, len(s.st[0][0]))
	c2 = min(c2, len(s.st[0][0][0]))
	k1 := bits.Len(uint(r2-r1)) - 1
	k2 := bits.Len(uint(c2-c1)) - 1
	// 视作四个子矩阵的并集
	return s.op(
		s.op(s.st[k1][k2][r1][c1], s.st[k1][k2][r2-1<<k1][c1]),
		s.op(s.st[k1][k2][r1][c2-1<<k2], s.st[k1][k2][r2-1<<k1][c2-1<<k2]),
	)
}

func countLocalMaximums1(matrix [][]int) (ans int) {
	st := newSparseTable2D(matrix, func(a, b int) int { return max(a, b) })

	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && max(st.query(i-x, j-x+1, i+x+1, j+x), st.query(i-x+1, j-x, i+x, j+x+1)) <= x {
				ans++
			}
		}
	}
	return
}

//

// 一维 ST 表（泛型版本）
type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

func newSparseTable[T any](a []T, op func(T, T) T) sparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	st[0] = a
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len8(uint8(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
type seg []sparseTable[int]

func (t seg) build(a [][]int, node, l, r int) {
	if l == r { // 叶子
		t[node] = newSparseTable(a[l], func(a, b int) int { return max(a, b) })
		return
	}

	m := (l + r) / 2
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树

	merged := make([]int, len(a[0]))
	for i := range merged {
		merged[i] = max(t[node*2].st[0][i], t[node*2+1].st[0][i]) // 行号 [l, r] 中的第 i 列的最大值
	}
	t[node] = newSparseTable(merged, func(a, b int) int { return max(a, b) })
}

// 行号闭区间 [r1, r2]，列号左闭右开 [c1, c2)
func (t seg) query(node, l, r, r1, r2, c1, c2 int) int {
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

func countLocalMaximums(matrix [][]int) (ans int) {
	n, m := len(matrix), len(matrix[0])
	// 线段树每个节点 [l, r] 保存的是，当上下边界固定为 l 和 r 时，把每一列的最大值视作一个 int，这 m 个数的一维 ST 表
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(matrix, 1, 0, n-1)

	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && max(t.query(1, 0, n-1, max(i-x, 0), min(i+x, n-1), max(j-x+1, 0), min(j+x, m)),
				t.query(1, 0, n-1, max(i-x+1, 0), min(i+x-1, n-1), max(j-x, 0), min(j+x+1, m))) <= x {
				ans++
			}
		}
	}
	return
}
