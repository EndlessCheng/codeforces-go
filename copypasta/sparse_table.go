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

type sparseTable struct {
	st [][]int
	op func(int, int) int
}

// 时间复杂度 O(n * log n)
func newSparseTable(a []int, op func(int, int) int) sparseTable {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]int, w)
	for i := range st {
		st[i] = make([]int, n)
	}
	copy(st[0], a)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable{st, op}
}

// [l, r) 左闭右开，下标从 0 开始
// 返回 op(nums[l:r])
// 时间复杂度 O(1)
func (s sparseTable) query(l, r int) int {
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

// 二维 ST 表
// 其一 · 保证询问的是正方形区域
// https://blog.csdn.net/weixin_41162823/article/details/98471161
// https://www.luogu.com.cn/problem/P2216

// 其二 · 询问任意长宽
// https://blog.nowcoder.net/n/3eccd1386a8846398d3bee62b485309b
// https://codeforces.com/problemset/problem/1301/E 2500
