package copypasta

import "math/bits"

/* 稀疏表 Sparse Table
st[i][j] 对应的区间是 [i, i+2^j)
https://oi-wiki.org/ds/sparse-table/
https://codeforces.com/blog/entry/66643
Tarjan RMQ https://codeforces.com/blog/entry/48994
一些 RMQ 的性能对比 https://codeforces.com/blog/entry/78931
一个 RMQ 问题的快速算法，以及区间众数 https://zhuanlan.zhihu.com/p/79423299
将 LCA、RMQ、LA 优化至理论最优复杂度 https://www.luogu.com.cn/blog/ICANTAKIOI/yi-shang-shou-ke-ji-jiang-lcarmqla-you-hua-zhi-zui-you-fu-za-du
RMQ 标准算法和线性树上并查集 https://ljt12138.blog.uoj.ac/blog/4874
随机 RMQ https://www.luogu.com.cn/problem/P3793

模板题 https://www.luogu.com.cn/problem/P3865
https://codeforces.com/contest/1709/problem/D
二分/三指针 https://codeforces.com/problemset/problem/689/D
变长/种类 https://www.jisuanke.com/contest/11346/challenges
GCD https://codeforces.com/contest/1548/problem/B
题目推荐 https://cp-algorithms.com/data_structures/sparse-table.html#toc-tgt-5
*/

type ST [][]int

func NewST(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, v := range a {
		st[i] = make([]int, sz)
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = st.Op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// 查询区间 [l,r)，注意 l 和 r 是从 0 开始算的
func (st ST) Query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return st.Op(st[l][k], st[r-1<<k][k])
}

// min, max, gcd, ...
func (ST) Op(int, int) (_ int) { return }

//

// 下标版本，查询返回的是区间最值的下标
// https://codeforces.com/problemset/problem/675/E
// - 此题另一种做法是单调栈二分，见 https://www.luogu.com.cn/problem/solution/CF675E
type stPair struct{ v, i int }
type ST2 [][]stPair

func NewST2(a []int) ST2 {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST2, n)
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
func (st ST2) Query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	a, b := st[l][k], st[r-1<<k][k]
	if a.v <= b.v { // 最小值，相等时下标取左侧
		return a.i
	}
	return b.i
}
