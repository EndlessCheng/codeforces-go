package copypasta

import "math/bits"

/* 稀疏表 Sparse Table
st[i][j] 对应的区间是 [i, i+2^j)
https://oi-wiki.org/ds/sparse-table/
https://codeforces.com/blog/entry/66643
一些 RMQ 的性能对比 https://codeforces.com/blog/entry/78931
一个 RMQ 问题的快速算法，以及区间众数 https://zhuanlan.zhihu.com/p/79423299
将 LCA、RMQ、LA 优化至理论最优复杂度 https://www.luogu.com.cn/blog/ICANTAKIOI/yi-shang-shou-ke-ji-jiang-lcarmqla-you-hua-zhi-zui-you-fu-za-du

模板题 https://www.luogu.com.cn/problem/P3865
变长/种类 https://www.jisuanke.com/contest/11346/challenges
GCD https://codeforces.com/contest/1548/problem/B
题目推荐 https://cp-algorithms.com/data_structures/sparse-table.html#toc-tgt-5
*/
func sparseTableCollections() {
	core := func(int, int) (_ int) { return } // min, max, gcd, ...

	// 17: n<131072, 18: n<262144, 19: n<524288, 20: n<1048576
	// 动态写法 mx = bits.Len(uint(n))
	const mx = 17
	var st [][mx]int
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]int, n)
		for i, v := range a {
			st[i][0] = v
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				st[i][j] = core(st[i][j-1], st[i+1<<(j-1)][j-1])
			}
		}
	}
	// [l,r) 注意 l r 是从 0 开始算的
	stQuery := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return core(st[l][k], st[r-1<<k][k]) }

	_, _ = stInit, stQuery

	// 下标版本，查询返回的是区间最值的下标
	{
		type pair struct{ v, i int }
		const mx = 17
		var st [][mx]pair
		stInit := func(a []int) {
			n := len(a)
			st = make([][mx]pair, n)
			for i, v := range a {
				st[i][0] = pair{v, i}
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
		}
		// [l,r) 注意 l r 是从 0 开始算的
		stQuery := func(l, r int) int {
			k := bits.Len(uint(r-l)) - 1
			a, b := st[l][k], st[r-1<<k][k]
			if a.v <= b.v { // 最小值，相等时下标取左侧
				return a.i
			}
			return b.i
		}
		_, _ = stInit, stQuery
	}
}
