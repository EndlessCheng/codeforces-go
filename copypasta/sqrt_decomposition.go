package copypasta

import (
	"math"
	"slices"
)

/* 根号算法 根号分治 Sqrt Decomposition
一种技巧：组合两种算法从而降低复杂度 O(n^2) -> O(n√n)
常用于图论或者某些数组统计类题目
参考 Competitive Programmer’s Handbook Ch.27
王悦同《根号算法——不只是分块》
todo 暴力美学——浅谈根号分治 https://www.luogu.com.cn/blog/Amateur-threshold/pu-li-mei-xue-qian-tan-gen-hao-fen-zhi

题目花样很多，下面举个例子
有 n 个对象，每个对象有一个「关于其他对象的统计量」ci（一个数、一个集合的元素个数，等等）
为方便起见，假设 ∑ci 的数量级和 n 一样，下面用 n 表示 ∑ci
当 ci > √n 时，这样的对象不超过 √n 个，暴力枚举这些对象之间的关系（或者，该对象与其他所有对象的关系），时间复杂度为 O(n) 或 O(n√n)。此乃算法一
当 ci ≤ √n 时，这样的对象有 O(n) 个，由于统计量 ci 很小，暴力枚举当前对象的统计量，时间复杂度为 O(n√n)。此乃算法二
这样，以 √n 为界，我们将所有对象划分成了两组，并用两个不同的算法处理
这两种算法是看待同一个问题的两种不同方式，通过恰当地组合（平衡）这两个算法，复杂度由 O(n^2) 降至 O(n√n)
例子是 https://codeforces.com/problemset/problem/1806/E
注意：**枚举时要做到不重不漏**

- [1714. 数组中特殊等间距元素的和](https://leetcode.cn/problems/sum-of-special-evenly-spaced-elements-in-array/)（会员题）
可以从这题上手 https://www.luogu.com.cn/problem/P3396
- 同 https://codeforces.com/contest/103/problem/D 2100
https://www.luogu.com.cn/problem/T279521?contestId=65460
- https://www.luogu.com.cn/blog/cyffff/solution-JRKSJ-Eltaw
LCP16 https://leetcode.cn/problems/you-le-yuan-de-you-lan-ji-hua/
https://codeforces.com/problemset/problem/1921/F 1900
https://codeforces.com/problemset/problem/797/E 2000
https://codeforces.com/problemset/problem/1207/F 2100
https://codeforces.com/problemset/problem/1806/E 2200
https://codeforces.com/problemset/problem/1968/G2 2200 也可以直接记忆化
https://codeforces.com/problemset/problem/425/D 2300
https://codeforces.com/problemset/problem/677/D 2300
https://codeforces.com/problemset/problem/1468/M 2300 或四元环
https://codeforces.com/problemset/problem/342/E 2400
https://codeforces.com/problemset/problem/506/D 2400
https://codeforces.com/problemset/problem/786/C 2400 见下面的 floorDivide
https://codeforces.com/problemset/problem/1619/H 2400 大步+小步，有点分段打表的味道
https://codeforces.com/problemset/problem/1270/F 2600
https://codeforces.com/problemset/problem/1039/D 2800
https://codeforces.com/problemset/problem/1039/E 3400
https://atcoder.jp/contests/abc293/tasks/abc293_f 四次方根
https://atcoder.jp/contests/abc365/tasks/abc365_g
https://leetcode.cn/problems/maximum-number-of-matching-indices-after-right-shifts/
- 频率高 (>t) 的字符拿出来做卷积，频率低的字符枚举所有对
- 这样前一部分复杂度是 O(n/t*n log n)，后一部分是 O(n*t)，取 t=sqrt(n log n) 得复杂度 O(n^1.5 log^0.5 n)
https://leetcode.com/discuss/interview-question/3517350/
- Given a list of pairs {L,R} & an array.
  Find out the total number of pairs (i, j) where
  (arr[i], arr[j]) should be equal to exactly one of the given Q pairs. And i < j.

自动 O(n√n)
https://codeforces.com/problemset/problem/1790/F 2100

随机+分块
https://codeforces.com/contest/1840/problem/G2 2500
*/

// TIPS: n 的整数分拆中，不同数字的个数至多有 O(√n) 种

/* 分散层叠算法 Fractional Cascading
https://en.wikipedia.org/wiki/Fractional_cascading
https://www.zhihu.com/question/33776070/answer/59405602
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
区间 rank 的 N 种解法 https://www.luogu.com.cn/blog/Peterprpr/HunterNoHorse
浅谈归约矩乘 https://www.luogu.com.cn/blog/Ynoi/qian-tan-gui-yue-ju-cheng
为什么有些题只能用根号科技？https://www.luogu.com.cn/blog/include-c/sqrt-tech

题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
好题 https://codeforces.com/problemset/problem/91/E
todo 动态逆序对 https://www.luogu.com.cn/problem/P3157 https://www.luogu.com.cn/problem/UVA11990
https://cp-algorithms.com/sequences/rmq.html
todo https://www.luogu.com.cn/problem/P3396
 https://codeforces.com/problemset/problem/1207/F
 https://codeforces.com/contest/455/problem/D
https://codeforces.com/problemset/problem/13/E 2700
https://www.luogu.com.cn/problem/P3203 [HN10] 弹飞绵羊
https://www.nowcoder.com/discuss/353159150542725120 K 题 https://ac.nowcoder.com/acm/contest/view-submission?submissionId=50861318
*/
func _(a []int) {
	// 下标从 0 开始
	type block struct {
		l, r           int // [l,r]
		origin, sorted []int
		//lazyAdd int
	}
	n := len(a)
	blockSize := int(math.Sqrt(float64(n))) // 建议设为常量，避免超时 【提示】大一点可能更快一些
	//blockSize := int(math.Sqrt(float64(n) * math.Log2(float64(n+1))))
	blocks := make([]block, (n-1)/blockSize+1)
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
		b.sorted = slices.Clone(b.origin)
		slices.Sort(b.sorted)
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

	_ = sqrtOp
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
