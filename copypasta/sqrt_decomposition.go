package copypasta

import (
	"math"
	"sort"
)

/* 根号分治 Sqrt Decomposition
一种技巧：组合两种算法从而降低复杂度 O(n^2) -> O(n√n)
参考 Competitive Programmer’s Handbook Ch.27
王悦同《根号算法——不只是分块》

题目花样很多，下面举个例子
有 n 个对象，每个对象有一个「关于其他对象的统计量」ci（一个数、一个集合的元素个数，等等）
为方便起见，假设 ∑ci 的数量级和 n 一样，下面用 n 表示 ∑ci
当 ci > √n 时，这样的对象不超过 √n 个，暴力枚举这些对象之间的关系（或者，该对象与其他所有对象的关系），时间复杂度为 O(n) 或 O(n√n)。此乃算法一
当 ci ≤ √n 时，这样的对象有 O(n) 个，由于统计量 ci 很小，暴力枚举当前对象的统计量，时间复杂度为 O(n√n)。此乃算法二
这样，以 √n 为界，我们将所有对象划分成了两组，并用两个不同的算法处理
这两种算法是看待同一个问题的两种不同方式，通过恰当地组合（平衡）这两个算法，复杂度由 O(n^2) 降至 O(n√n)
例子是 https://codeforces.com/problemset/problem/1806/E
注意：**枚举时要做到不重不漏**

可以从这题上手 https://www.luogu.com.cn/problem/P3396 同 https://codeforces.com/contest/103/problem/D
https://www.luogu.com.cn/problem/T279521?contestId=65460 https://www.luogu.com.cn/blog/cyffff/solution-JRKSJ-Eltaw
https://codeforces.com/problemset/problem/342/E 2400
https://codeforces.com/problemset/problem/425/D
https://codeforces.com/problemset/problem/677/D
https://codeforces.com/problemset/problem/786/C 2400
https://codeforces.com/problemset/problem/797/E
https://codeforces.com/problemset/problem/1207/F
https://codeforces.com/problemset/problem/1468/M 或四元环
LCP16 https://leetcode-cn.com/problems/you-le-yuan-de-you-lan-ji-hua/
https://codeforces.com/problemset/problem/1039/D
https://codeforces.com/problemset/problem/1039/E
大步+小步，有点分段打表的味道 https://codeforces.com/problemset/problem/1619/H
https://codeforces.com/problemset/problem/1806/E
见下面的 floorDivide https://codeforces.com/problemset/problem/786/C

自动 O(n√n)
https://codeforces.com/problemset/problem/1790/F
*/

// TIPS: n 的整数分拆中，不同数字的个数至多有 O(√n) 种

/* 分散层叠算法 Fractional Cascading
https://en.wikipedia.org/wiki/Fractional_cascading
https://www.luogu.com.cn/blog/DPair2005/fen-san-ceng-die-suan-fa-xue-xi-bi-ji
https://www.luogu.com.cn/problem/P6466
*/

/*
分块数据结构
https://oi-wiki.org/ds/decompose/
https://oi-wiki.org/ds/block-array/
Unrolled linked list https://en.wikipedia.org/wiki/Unrolled_linked_list
【推荐】https://www.luogu.com.cn/blog/220037/Sqrt1
浅谈基础根号算法——分块 https://www.luogu.com.cn/blog/deco/qian-tan-ji-chu-gen-hao-suan-fa-fen-kuai
todo https://www.csie.ntu.edu.tw/~sprout/algo2018/ppt_pdf/root_methods.pdf
【都是我对数据结构的爱啊】区间 rank 的 N 种解法，你都会了吗 https://www.luogu.com.cn/blog/Peterprpr/HunterNoHorse

题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
好题 https://codeforces.com/problemset/problem/91/E
todo 动态逆序对 https://www.luogu.com.cn/problem/P3157 https://www.luogu.com.cn/problem/UVA11990
https://cp-algorithms.com/sequences/rmq.html
todo https://www.luogu.com.cn/problem/P3396
 https://codeforces.com/problemset/problem/1207/F
 https://codeforces.com/contest/455/problem/D
*/
func _(min, max func(int, int) int) {
	type block struct {
		l, r           int // [l,r]
		origin, sorted []int
		//lazyAdd int
	}
	var blocks []block
	sqrtInit := func(a []int) {
		n := len(a)
		blockSize := int(math.Sqrt(float64(n)))
		//blockSize := int(math.Sqrt(float64(n) * math.Log2(float64(n+1))))
		blockNum := (n-1)/blockSize + 1
		blocks = make([]block, blockNum)
		for i, v := range a {
			j := i / blockSize
			if i%blockSize == 0 {
				blocks[j] = block{l: i, origin: make([]int, 0, blockSize)}
			}
			blocks[j].origin = append(blocks[j].origin, v)
		}
		for i := range blocks {
			b := &blocks[i]
			b.r = b.l + len(b.origin) - 1
			b.sorted = append([]int(nil), b.origin...)
			sort.Ints(b.sorted)
		}
	}
	sqrtOp := func(l, r int, v int) { // [l,r], starts at 0
		for i := range blocks {
			b := &blocks[i]
			if b.r < l {
				continue
			}
			if b.l > r {
				break
			}
			if l <= b.l && b.r <= r {
				// do op on full block
			} else {
				// do op on part block
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl - b.l; j <= br-b.l; j++ {
					// do b.origin[j]...
				}
			}
		}
	}

	_ = []interface{}{sqrtInit, sqrtOp}
}

// 如果 f(i) 的计算结果近似 n/i
// 可以对 [1,n] 值域分治，如果区间内的结果都相同，则不再分治
// 时间复杂度类似整除分块 O(f(n)√n)
// https://codeforces.com/problemset/problem/786/C
func floorDivide(n int, f func(int) int) []int {
	ans := make([]int, n+1)
	var solve func(int, int)
	solve = func(l, r int) {
		if l > r {
			return
		}
		resL, resR := f(l), f(r)
		if resL == resR {
			for i := l; i <= r; i++ {
				ans[i] = resL
			}
			return
		}
		ans[l] = resL
		ans[r] = resR
		mid := (l + r) / 2
		solve(l+1, mid)
		solve(mid+1, r-1)
	}
	solve(1, n)
	return ans[1:]
}
