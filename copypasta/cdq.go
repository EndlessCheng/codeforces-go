package copypasta

import "slices"

/*
CDQ 分治
https://oi-wiki.org/misc/cdq-divide/
https://www.luogu.com.cn/blog/ljc20020730/cdq-fen-zhi-xue-xi-bi-ji
https://codeforces.com/blog/entry/150934
https://blog.nowcoder.net/n/f44d4aada5a24f619442dd6ddffa7320
https://zhuanlan.zhihu.com/p/332996578

题单 https://www.luogu.com.cn/training/9435
https://www.luogu.com.cn/problem/P3810 模板题 三维偏序
https://www.luogu.com.cn/problem/P3157 动态逆序对 https://www.luogu.com.cn/problem/UVA11990
https://www.luogu.com.cn/problem/P4390 带修二维数点
https://www.luogu.com.cn/problem/P4169 拆开绝对值，四种情况分别计算
https://www.luogu.com.cn/problem/P4093 CDQ 优化 DP
https://www.luogu.com.cn/problem/P2487 CDQ 优化 DP
https://codeforces.com/problemset/problem/762/E  2200 做到复杂度与 k 无关
https://codeforces.com/problemset/problem/1045/G 2200 同 CF762E
https://codeforces.com/problemset/problem/1093/E 2400
https://codeforces.com/problemset/problem/848/C  2600 转化成带修二维数点
https://codeforces.com/problemset/problem/1175/G 3000 斜率优化
https://acm.hdu.edu.cn/showproblem.php?pid=5126 四维偏序

值域分治
LC4055 https://leetcode.cn/problems/count-shadow-pairs-ii/
https://www.luogu.com.cn/problem/P13508

*/

// 动态逆序对
// https://www.luogu.com.cn/problem/P3157
func dynamicInversion(nums, del []int) (res []int) {
	n := len(nums)
	m := len(del)

	type pair struct{ i, t, inv int }
	a := make([]pair, n+1)
	f := make(fenwick, n+1)
	inv := 0 // 这里算出初始逆序对个数
	for i, v := range nums {
		inv += f.query(v+1, n)
		f.update(v, 1)
		a[v].i = i
	}
	for t, v := range del {
		a[v].t = t + 1
	}
	for i, p := range a {
		if p.t == 0 {
			a[i].t = m + 1 // 没被删除的数，删除时间视作 m+1
		}
	}

	f = make(fenwick, m+2)
	var solve func(int, int)
	solve = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		solve(l, mid)
		solve(mid, r)

		// v < w
		// p[v] > p[w]
		// 删 w：t[v] > t[w]
		// 删 v：t[w] > t[v]

		v := l
		for w := mid; w < r; w++ {
			for ; v < mid && a[v].i > a[w].i; v++ {
				f.update(a[v].t, 1)
			}
			a[w].inv += f.query(a[w].t+1, m+1)
		}
		for v--; v >= l; v-- {
			f.update(a[v].t, -1)
		}

		w := r - 1
		for v := mid - 1; v >= l; v-- {
			for ; w >= mid && a[w].i < a[v].i; w-- {
				f.update(a[w].t, 1)
			}
			a[v].inv += f.query(a[v].t+1, m+1)
		}
		for w++; w < r; w++ {
			f.update(a[w].t, -1)
		}

		slices.SortFunc(a[l:r], func(a, b pair) int { return b.i - a.i })
	}
	solve(1, n+1)

	// 这里的排序也可以改用数组记录 https://www.luogu.com.cn/record/297905725
	slices.SortFunc(a[1:], func(a, b pair) int { return a.t - b.t })
	for _, p := range a[1:] {
		if p.t > m {
			break
		}
		res = append(res, inv)
		inv -= p.inv
	}
	return
}

/* 整体二分 Parallel Binary Search
https://oi-wiki.org/misc/parallel-binsearch/
https://codeforces.com/blog/entry/45578
todo 整体二分解决静态区间第 k 小的优化 https://www.luogu.com/article/gbzqyzwn
模板题 https://www.luogu.com.cn/problem/P3527
https://www.luogu.com.cn/problem/P2617
https://atcoder.jp/contests/agc002/tasks/agc002_d
https://www.hackerrank.com/contests/hourrank-23/challenges/selective-additions/problem
https://www.codechef.com/problems/MCO16504
*/
func parallelBinarySearch(n int, qs []struct{ l, r, v int }) []int {
	// 读入询问时可以处理成左闭右开的形式
	ans := make([]int, n)
	tar := make([]int, n)
	for i := range tar {
		tar[i] = i
	}
	var f func([]int, int, int)
	f = func(tar []int, ql, qr int) {
		if len(tar) == 0 {
			return
		}
		if ql+1 == qr {
			for _, c := range tar {
				ans[c] = ql // qr
			}
			return
		}
		qm := (ql + qr) / 2
		for _, q := range qs[ql:qm] {
			_ = q
			// apply(q)

		}

		// 根据此刻查询的结果将 tar 分成左右两部分
		var left, right []int
		for _, who := range tar {
			_ = who

		}

		for _, q := range qs[ql:qm] {
			_ = q
			// rollback(q)

		}
		f(left, ql, qm)
		f(right, qm, qr)
	}
	f(tar, 0, len(qs)+1) // 这样可以将无法满足要求的 ans[i] 赋值为 len(qs)
	return ans
}
