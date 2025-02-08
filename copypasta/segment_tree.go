package copypasta

import "math/bits"

// 线段树讲解 by 灵茶山艾府（13:30 开始）https://www.bilibili.com/video/BV15D4y1G7ms

// 可视化 https://visualgo.net/zh/segmenttree

// 推荐阅读《算法竞赛进阶指南》0x43 和 0x48 节
// https://oi-wiki.org/ds/seg/
// https://cp-algorithms.com/data_structures/segment_tree.html
// [Monoid 幺半群] Generalizing Segment Trees https://sharmaeklavya2.github.io/blog/generalizing-segment-trees.html
// 总结得比较详细 https://www.acwing.com/blog/content/1684/
// 线段树进阶 Part 1 https://www.luogu.com.cn/blog/AlexWei/Segment-Tree-Part-1
// https://codeforces.com/blog/entry/18051
// https://codeforces.com/blog/entry/89313
// https://codeforces.com/blog/entry/15890
// todo 高效线段树 crazySegmentTree https://codeforces.com/blog/entry/89399
// https://en.algorithmica.org/hpc/data-structures/segment-trees/
// 像使用 STL 一样使用线段树 https://zhuanlan.zhihu.com/p/459679512 https://zhuanlan.zhihu.com/p/459880950
// 数组两倍空间线段树 https://www.cnblogs.com/chy-2003/p/11815396.html
// 线段树诡异题目收录 https://zhuanlan.zhihu.com/p/124181375
// Limitの线段树题单 https://www.luogu.com.cn/training/1124
// todo [题单] 线段树的进阶用法 https://www.luogu.com.cn/training/221#problems

// todo Offline Range MEX queries in O(log n) https://codeforces.com/blog/entry/117688

// 注：对于指针写法，必要时禁止 GC，能加速不少
// func init() { debug.SetGCPercent(-1) }

// 模板（单点修改、区间查询）https://www.luogu.com.cn/problem/P2068

/*
如果一个题目可以用分治解决，那么这个题目的带修改版本可以用线段树解决

带修最长连续相同子串 LC2213 https://leetcode.cn/problems/longest-substring-of-one-repeating-character/
带修最大子段和 https://www.luogu.com.cn/problem/P4513
- 代码 https://www.luogu.com.cn/record/50262292
- https://codeforces.com/edu/course/2/lesson/4/2/practice/contest/273278/problem/A
带修最大子段和+按位或 https://www.luogu.com.cn/problem/P7492 https://www.luogu.com.cn/contest/42328
带修打家劫舍 https://www.luogu.com.cn/problem/P3097
- LC https://leetcode.cn/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/
*/

// 势能线段树：区间开方、区间取模、区间 GCD 一个数，都是可以暴力更新的
// 关于线段树上的一些进阶操作 https://www.luogu.com/article/aentaeud
// 区间开方见 CF920F
// 区间取模见 CF438D
// 区间 GCD 一个数见 https://www.luogu.com.cn/problem/P9989 https://www.cnblogs.com/Athanasy/p/17940070
// https://www.luogu.com.cn/problem/P10516
// 另见吉老师线段树 Segment Tree Beats https://www.luogu.com.cn/problem/P6242 【模板】线段树 3（区间最值操作、区间历史最值）
// 另见 Kinetic Tournament 树 (KTT) https://www.luogu.com.cn/problem/P5693
// https://www.luogu.com.cn/problem/P10587

// https://www.luogu.com.cn/problem/P4588 乘法 单点修改
// https://codeforces.com/problemset/problem/2050/F 1700 GCD
// https://codeforces.com/problemset/problem/914/D 1900 GCD 
// https://codeforces.com/problemset/problem/380/C 2000 区间最长括号子序列 
// https://codeforces.com/contest/474/problem/E 2000 最值及其下标 
// https://codeforces.com/problemset/problem/522/D 2000 区间最短线段长度 
// https://codeforces.com/problemset/problem/920/F 2000 开方（也可以用并查集做）
// - https://www.luogu.com.cn/problem/P4145
// - http://acm.hdu.edu.cn/showproblem.php?pid=4027
// 区间（绝对）众数及其次数（摩尔投票算法）https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
// - LC169 https://leetcode.cn/problems/majority-element/
// - LC1157 https://leetcode.cn/problems/online-majority-element-in-subarray/
// - https://www.luogu.com.cn/problem/P3567
// - https://www.luogu.com.cn/problem/P3765
// - https://codeforces.com/contest/1514/problem/D 2000
// https://codeforces.com/problemset/problem/703/D 2100 区间元素去重后的异或和 
// - 联系 https://www.luogu.com.cn/problem/P1972
// todo https://codeforces.com/problemset/problem/1567/E 2200 区间连续递增子数组个数
// https://codeforces.com/problemset/problem/1179/C 2200
// https://codeforces.com/problemset/problem/1906/F 2200 最大子数组和（非空） 离线
// https://codeforces.com/problemset/problem/438/D 2300 取模
// https://codeforces.com/problemset/problem/1093/G 2300 k 维曼哈顿（单点修改+区间最大值）
// https://codeforces.com/problemset/problem/1187/D 2400 转换的好题
// https://codeforces.com/problemset/problem/1401/F 2400 区间 swap & reverse
// - 联想 reverse bit 的递归思路
// https://codeforces.com/problemset/problem/1436/E 2400 所有子数组的 mex 的 mex
// 区间 mex https://www.luogu.com.cn/problem/P4137
// - 做法之一是离线+线段树二分 https://www.luogu.com.cn/blog/user7035/solution-p4137
// - 也可以用树状数组 https://www.luogu.com.cn/blog/Atalod/ti-xie-p4137-post
// - 反向构造题 https://www.luogu.com.cn/problem/P6852
// https://codeforces.com/problemset/problem/1609/E 2400 单点修改 + 不含子序列 abc
// https://codeforces.com/problemset/problem/1665/E 2500 区间最小的 31 个数
// https://codeforces.com/problemset/problem/2042/F 2600 两段最大子段和
// todo https://codeforces.com/problemset/problem/19/D 2800
// https://codeforces.com/problemset/problem/765/F 3100 最小差值
// - https://codeforces.com/problemset/problem/1793/F 2600 和 CF765F 是同一题
// - 不允许相等 https://www.luogu.com.cn/problem/P5926
// https://atcoder.jp/contests/abc285/tasks/abc285_f
//
// 题目推荐 https://cp-algorithms.com/data_structures/segment_tree.html#toc-tgt-12
// 力扣 https://leetcode.cn/tag/segment-tree/
// 另见 dp.go 的数据结构优化 DP
// 另见 dp.go 的动态 DP
//
// todo http://poj.org/problem?id=2991
// 变换成值域 http://poj.org/problem?id=3368
// - https://www.luogu.com.cn/problem/SP1684
// - https://www.luogu.com.cn/problem/UVA11235
// http://poj.org/problem?id=3470
// todo http://poj.org/problem?id=1201

// 线段树二分
// LC2286 https://leetcode.cn/problems/booking-concert-tickets-in-groups/
// LC2940 https://leetcode.cn/problems/find-building-where-alice-and-bob-can-meet/

// EXTRA: 权值线段树
// 讲解与习题 https://www.luogu.com.cn/blog/bfqaq/qian-tan-quan-zhi-xian-duan-shu
// 浅谈权值线段树到主席树 https://www.luogu.com.cn/blog/your-alpha1022/WeightSegmentTree-ChairmanTree
// 谈树状数组套权值树 https://www.luogu.com.cn/blog/bfqaq/qian-tan-shu-zhuang-shuo-zu-quan-zhi-shu
// https://codeforces.com/problemset/problem/1042/D 1800
// todo 区间只出现一次的数的最大值 https://codeforces.com/problemset/problem/69/E

// EXTRA: 树套树
// 代码见 fenwick_tree.go
// 三维偏序：树状数组套动态开点权值线段树 https://www.luogu.com.cn/problem/P3810
// 二逼平衡树 https://www.luogu.com.cn/problem/P3380
// - 两种 O((n+q)lognlogU) 做法：
// - 树状数组套动态开点权值线段树（AC）
// - 动态开点权值线段树套下标平衡树（TLE）https://www.luogu.com.cn/record/136191286
// todo 二维线段树 https://www.luogu.com.cn/problem/P3437
// - LC308 https://leetcode.cn/problems/range-sum-query-2d-mutable/
//  https://www.luogu.com.cn/problem/P4514
//  树套树 标记永久化 https://www.luogu.com.cn/blog/Hoshino-kaede/chao-leng-men-shuo-ju-jie-gou-er-wei-xian-duan-shu-yang-xie

// EXTRA: 线段树优化建图
// 每个位置对应着 O(logn) 个线段树上的节点，每个区间可以拆分成至多 O(logn) 个线段树上的区间
// 这个性质可以用于优化建图
// https://www.luogu.com.cn/blog/forever-captain/DS-optimize-graph
// https://codeforces.com/problemset/problem/786/B 2300
// https://codeforces.com/problemset/problem/1903/F 2500
// todo https://www.luogu.com.cn/problem/P6348

/* 一些细节
了解下列内容将有助于 Hack 代码

区间 [1,1] 对应的节点编号为 1<<bits.Len(uint(n-1))   1e5 => 2^17       这也说明 [1,1] 对应节点编号x2 的节点是不会下标越界的
区间 [n,n] 对应的节点编号为 1<<bits.Len(uint(n))-1   1e5 => 2^17-1
当 n≠2^k 时，在内存中区间 [n,n] 的下一个就是区间 [1,1]

最下面一排的节点个数
最下面一排的节点编号之和
https://www.luogu.com.cn/problem/P9689?contestId=133572

什么时候空间最浪费？（指一般线段树的实现方式）
当 n=2^k-1 时，此时只需要 2n 的空间

什么时候空间最不浪费？（指一般线段树的实现方式）
当 n=2^k+d (d<<2^k) 时
record table
i/n    n i range
1.0000 1 1 [1,1]
1.5000 2 3 [2,2]
1.6667 3 5 [2,2]
1.7500 4 7 [4,4]
1.8000 5 9 [2,2]
2.1667 6 13 [5,5]
2.5000 10 25 [7,7] **《算法竞赛进阶指南》用的例子
2.7222 18 49 [11,11]
2.8500 20 57 [17,17]
2.8529 34 97 [19,19]
3.1389 36 113 [29,29] ** 首个超过 3n 的例子
3.3088 68 225 [53,53]
3.3472 72 241 [65,65]
3.4015 132 449 [101,101]
3.5368 136 481 [121,121]
3.6402 264 961 [233,233]
3.6507 272 993 [257,257]
3.6942 520 1921 [457,457]
3.7595 528 1985 [497,497]
3.8163 1040 3969 [977,977]
3.8191 1056 4033 [1025,1025]
3.8454 2064 7937 [1937,1937]
3.8774 2080 8065 [2017,2017]
3.9072 4128 16129 [4001,4001]
3.9079 4160 16257 [4097,4097]
3.9223 8224 32257 [7969,7969]
3.9381 8256 32513 [8129,8129]
3.9534 16448 65025 [16193,16193]
3.9535 16512 65281 [16385,16385]
3.9610 32832 130049 [32321,32321]
3.9689 32896 130561 [32641,32641]
3.9766 65664 261121 [65153,65153]
3.9767 65792 261633 [65537,65537] ** 1e5 以内的最大值
3.9805 131200 522241 [130177,130177]
3.9844 131328 523265 [130817,130817]
3.9883 262400 1046529 [261377,261377]
3.9883 262656 1047553 [262145,262145]
3.9902 524544 2093057 [522497,522497]
3.9922 524800 2095105 [523777,523777]
3.9941 1049088 4190209 [1047041,1047041]
3.9941 1049600 4192257 [1048577,1048577]
*/

// 线段树有两个下标，一个是线段树底层数组的下标，一个是线段树维护的区间的下标。
// 底层数组的下标：一般是从 1 开始的，从 0 开始也可以，把左右子树下标改成 2*i+1 和 2*i+2 就行。下面的代码从 1 开始。
// 线段树维护的区间的下标：这个其实无所谓，从 0 从 1 开始都可以。下面的代码从 0 开始。

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中
// 如果想记录最值及其下标，可以把 val 的类型改成 pair
type seg []struct {
	l, r int
	val  int // info
}

// 合并两个节点上的数据：maintain 和 query 通用
// 要求操作满足区间可加性
// 例如 + * | & ^ min max gcd mulMatrix 摩尔投票 最大子段和 ...
func (seg) mergeInfo(a, b int) int {
	return max(a, b)
}

// 单点更新：见 update
func (t seg) set(o, val int) {
	t[o].val = t.mergeInfo(t[o].val, val)
}

// 下标从 0 开始
// 调用时 o=1, l=0, r=n-1
func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 调用时 o=1  0<=i<=n-1
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].val = t.mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lRes := t.query(o<<1, l, r)
	rRes := t.query(o<<1|1, l, r)
	return t.mergeInfo(lRes, rRes)
}

func (t seg) queryAll() int { return t[1].val }

// 线段树二分：返回 [l,r] 内第一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内第一个大于等于 target 的元素下标，需要线段树维护区间最大值
//     t.findFirst(1, l, r, func(nodeMax int) bool { return nodeMax >= target })
// 调用时 o=1
// https://leetcode.cn/problems/booking-concert-tickets-in-groups/
// - https://leetcode.cn/problems/booking-concert-tickets-in-groups/submissions/517574644/
// https://leetcode.cn/problems/find-building-where-alice-and-bob-can-meet/
// - https://leetcode.cn/problems/find-building-where-alice-and-bob-can-meet/submissions/517575667/
func (t seg) findFirst(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].val) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	idx := t.findFirst(o<<1, l, r, f)
	if idx < 0 {
		idx = t.findFirst(o<<1|1, l, r, f)
	}
	return idx
}

// 线段树二分：返回 [l,r] 内最后一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内最后一个小于等于 target 的元素下标，需要线段树维护区间最小值
//     t.findLast(1, l, r, func(nodeMin int) bool { return nodeMin <= target })
// 调用时 o=1
func (t seg) findLast(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].val) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	idx := t.findLast(o<<1|1, l, r, f)
	if idx < 0 {
		idx = t.findLast(o<<1, l, r, f)
	}
	return idx
}

// a 的下标从 0 开始，线段树的区间下标也从 0 开始
func newSegmentTree(a []int) seg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

//

// 延迟标记（区间修改）
// 有关区间 flip 的 0-1 线段树，见 segment_tree01.go
//
// 【单个更新操作】
// + min/max https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/A
//           【推荐】https://codeforces.com/problemset/problem/1208/D 1900
//           https://codeforces.com/problemset/problem/1321/E 2000
//           https://codeforces.com/problemset/problem/52/C 2200
//           https://codeforces.com/problemset/problem/1110/F 2600
//           转换 https://codeforces.com/gym/294041/problem/E
//           todo 转换 https://atcoder.jp/contests/abc327/tasks/abc327_f
//           DP https://atcoder.jp/contests/dp/tasks/dp_w
// + ∑ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/D
//     https://www.luogu.com.cn/problem/P2068
//     https://www.luogu.com.cn/problem/P3372
// + + ∑ https://atcoder.jp/contests/abc357/tasks/abc357_f
// | & https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/C
// = min https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/E
// = ∑ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/F 
//     https://codeforces.com/problemset/problem/558/E 2300
// + 某个区间的不小于 x 的最小下标 https://codeforces.com/edu/course/2/lesson/5/3/practice/contest/280799/problem/C
// 单点查询的简化写法 https://codeforces.com/problemset/problem/292/E 1900
// - https://codeforces.com/contest/292/submission/173659179
// 不含任何长度 >= 2 的回文串 https://codeforces.com/contest/1881/problem/G 2000
// https://codeforces.com/problemset/problem/620/E   2100
// https://codeforces.com/problemset/problem/1295/E  2200
// https://codeforces.com/problemset/problem/1557/D  2200 max max 离散化 
// https://codeforces.com/problemset/problem/718/C   2300 矩阵乘法 ∑ 
// https://codeforces.com/problemset/problem/1797/E  2300 phi*
// https://codeforces.com/problemset/problem/145/E   2400
// https://codeforces.com/problemset/problem/1114/F  2400
// https://codeforces.com/problemset/problem/240/F   2600
// https://codeforces.com/problemset/problem/1439/C  2600 =max 求和的 O(log^2) 性质 
// https://codeforces.com/problemset/problem/1614/E  2600
// https://codeforces.com/problemset/problem/2009/G3 2700
// https://codeforces.com/problemset/problem/794/F   2800 数位修改 考察对懒标记的理解 
// todo https://codeforces.com/problemset/problem/1209/G2 3200
// https://atcoder.jp/contests/abc389/tasks/abc389_f 线段树二分
// 线段树二分与更新合并 LC2589 https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/
//                   LCP32 https://leetcode.cn/problems/t3fKg1/
// 维护平方和 LC2916 https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/
// todo https://www.luogu.com.cn/problem/P1471
//  斐波那契 https://codeforces.com/problemset/problem/446/C 2400
// 区间加等差数列（差分法）https://www.luogu.com.cn/problem/P1438 https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/B
//   多项式更新 Competitive Programmer’s Handbook Ch.28
// 区间加等比数列（q 恒定）
//   提出首项后，剩余的是固定的 1 + q^1 + q^2 + ... + q^(线段树区间长度-1)
//    可以预处理 1 + q^1 + q^2 + ... 的前缀和
//    线段树维护首项的和
//
// 【多个更新操作复合】
// = + max https://www.luogu.com.cn/problem/P1253
//         代码 https://www.luogu.com.cn/record/138265877
// * + ∑ https://www.luogu.com.cn/problem/P3373
//       LC1622 https://leetcode.cn/problems/fancy-sequence/
// = + ∑ https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/A
// * + ∑ai^k(k≤10) https://www.zhihu.com/question/564007656 B
// 线段树维护区间加、乘、赋值、平方和、立方和 http://acm.hdu.edu.cn/showproblem.php?pid=4578
// - https://www.cnblogs.com/dyhaohaoxuexi/p/14046275.html
// 注：区间赋值（=x）可以看成是先 *0 再 +x
// 三维向量 * + 交换 ∑^2 http://118.190.20.162/view.page?gpid=T119（需要注册 CCF）
//
// 吉老师线段树 吉司机线段树 Segment Tree Beats (Seg-beats)
// todo https://oi-wiki.org/ds/seg-beats/
//  https://codeforces.com/blog/entry/57319
//  区间最值操作与区间历史最值详解 https://www.luogu.com.cn/blog/Hakurei-Reimu/seg-beats
//  模板题 https://www.luogu.com.cn/problem/P6242
const todoInit = 0

type lazySeg []struct {
	l, r int
	sum  int // info
	todo int
}

func (lazySeg) mergeInfo(a, b int) int {
	return a + b // % mod
}

func (t lazySeg) do(o int, v int) {
	to := &t[o]

	// 更新 v 对整个区间的影响
	to.sum += v * (to.r - to.l + 1)

	// 更新 v 对左右儿子的影响
	to.todo += v
	// % mod
}

func (t lazySeg) spread(o int) {
	if v := t[o].todo; v != todoInit {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit
	}
}

// 下标从 0 开始
// 调用时 o=1, l=0, r=n-1
func (t lazySeg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].sum = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t lazySeg) update(o, l, r int, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t lazySeg) maintain(o int) {
	t[o].sum = t.mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t lazySeg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lRes := t.query(o<<1, l, r)
	rRes := t.query(o<<1|1, l, r)
	return t.mergeInfo(lRes, rRes)
}

func (t lazySeg) queryAll() int { return t[1].sum }

// 线段树二分：返回 [l,r] 内第一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内第一个大于等于 target 的元素下标，需要线段树维护区间最大值
//     t.findFirst(1, l, r, func(nodeMax int) bool { return nodeMax >= target })
// 调用时 o=1
func (t lazySeg) findFirst(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].sum) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findFirst(o<<1, l, r, f)
	if idx < 0 {
		idx = t.findFirst(o<<1|1, l, r, f)
	}
	return idx
}

// 线段树二分：返回 [l,r] 内最后一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内最后一个小于等于 target 的元素下标，需要线段树维护区间最小值
//     t.findLast(1, l, r, func(nodeMin int) bool { return nodeMin <= target })
// 调用时 o=1
func (t lazySeg) findLast(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].sum) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findLast(o<<1|1, l, r, f)
	if idx < 0 {
		idx = t.findLast(o<<1, l, r, f)
	}
	return idx
}

// a 的下标从 0 开始，线段树的区间下标也从 0 开始
func newLazySegmentTree(a []int) lazySeg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

// EXTRA: 适用于需要提取所有元素值的场景
func (t lazySeg) spreadAll(o int) {
	if t[o].l == t[o].r {
		return
	}
	t.spread(o)
	t.spreadAll(o << 1)
	t.spreadAll(o<<1 | 1)
}

//

// 动态开点线段树·其一·单点修改
// 注：如果 TLE 可以使用 func init() { debug.SetGCPercent(-1) } 加速
// LC327 https://leetcode.cn/problems/count-of-range-sum/
// LC2770 https://leetcode.cn/problems/maximum-number-of-jumps-to-reach-the-last-index/ 1533
// LC2736 https://leetcode.cn/problems/maximum-sum-queries/ 2533
// https://codeforces.com/problemset/problem/2000/H 2200
// - 相似题目 https://www.luogu.com.cn/problem/P2894
// todo https://codeforces.com/problemset/problem/1614/E 2600
// https://atcoder.jp/contests/abc351/tasks/abc351_f
// 树套树见 fenwick_tree.go
var emptyStNode = &stNode{val: stNodeDefaultVal}

func init() {
	emptyStNode.lo = emptyStNode
	emptyStNode.ro = emptyStNode
}

const stNodeDefaultVal = 0 // 如果求最大值并且有负数，改成 math.MinInt

type stNode struct {
	lo, ro *stNode
	l, r   int
	val    int
}

func (stNode) mergeInfo(a, b int) int {
	return max(a, b)
}

func (o *stNode) maintain() {
	// 注意这里没有判断 o.lo 和 o.ro 是空的，因为用的 emptyStNode
	o.val = o.mergeInfo(o.lo.val, o.ro.val)
}

func (o *stNode) update(i, val int) {
	if o.l == o.r {
		o.val = o.mergeInfo(o.val, val)
		return
	}
	m := (o.l + o.r) >> 1 // 使用右移而不是除法，这样保证负数也是下取整     或者 o.l + (o.r - o.l) / 2
	if i <= m {
		if o.lo == emptyStNode {
			o.lo = &stNode{lo: emptyStNode, ro: emptyStNode, l: o.l, r: m, val: stNodeDefaultVal}
		}
		o.lo.update(i, val)
	} else {
		if o.ro == emptyStNode {
			o.ro = &stNode{lo: emptyStNode, ro: emptyStNode, l: m + 1, r: o.r, val: stNodeDefaultVal}
		}
		o.ro.update(i, val)
	}
	o.maintain()
}

func (o *stNode) query(l, r int) int {
	if o == emptyStNode || l > o.r || r < o.l {
		return stNodeDefaultVal
	}
	if l <= o.l && o.r <= r {
		return o.val
	}
	return o.mergeInfo(o.lo.query(l, r), o.ro.query(l, r))
}

// 询问范围的最小值、最大值
// 0 1e9
// -2e9 2e9
func newStRoot(l, r int) *stNode {
	return &stNode{lo: emptyStNode, ro: emptyStNode, l: l, r: r, val: stNodeDefaultVal}
}

// 动态开点线段树·其二·延迟标记（区间修改）
// https://codeforces.com/problemset/problem/915/E（注：此题有多种解法）
// https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/F https://www.luogu.com.cn/problem/P5848
//（内存受限）https://codeforces.com/problemset/problem/1557/D
var emptyLazyNode = &lazyNode{sum: lazyNodeDefaultVal, todo: lazyNodeDefaultTodoVal}

func init() {
	emptyLazyNode.lo = emptyLazyNode
	emptyLazyNode.ro = emptyLazyNode
}

const lazyNodeDefaultVal = 0
const lazyNodeDefaultTodoVal = 0

type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int // info
	todo   int
}

func (lazyNode) mergeInfo(a, b int) int {
	return a + b // max(a, b)
}

func (o *lazyNode) maintain() {
	o.sum = o.mergeInfo(o.lo.sum, o.ro.sum)
}

func (o *lazyNode) do(add int) {
	o.todo += add                  // % mod
	o.sum += (o.r - o.l + 1) * add // % mod
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == emptyLazyNode {
		o.lo = &lazyNode{lo: emptyLazyNode, ro: emptyLazyNode, l: o.l, r: m, sum: lazyNodeDefaultVal, todo: lazyNodeDefaultTodoVal}
	}
	if o.ro == emptyLazyNode {
		o.ro = &lazyNode{lo: emptyLazyNode, ro: emptyLazyNode, l: m + 1, r: o.r, sum: lazyNodeDefaultVal, todo: lazyNodeDefaultTodoVal}
	}
	if v := o.todo; v != lazyNodeDefaultTodoVal {
		o.lo.do(v)
		o.ro.do(v)
		o.todo = lazyNodeDefaultTodoVal
	}
}

func (o *lazyNode) update(l, r int, add int) {
	if l <= o.l && o.r <= r {
		o.do(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) int {
	if o == emptyLazyNode || l > o.r || r < o.l {
		return lazyNodeDefaultVal
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	lRes := o.lo.query(l, r)
	rRes := o.ro.query(l, r)
	return o.mergeInfo(lRes, rRes)
}

// 询问范围的最小值、最大值
// 0 1e9
// -2e9 2e9
func newLazyRoot(l, r int) *lazyNode {
	return &lazyNode{lo: emptyLazyNode, ro: emptyLazyNode, l: l, r: r, sum: lazyNodeDefaultVal, todo: lazyNodeDefaultTodoVal}
}

// EXTRA: 线段树合并
// https://www.luogu.com.cn/problem/P5494
// todo 一些题目 https://www.luogu.com.cn/blog/styx-ferryman/xian-duan-shu-ge-bing-zong-ru-men-dao-fang-qi
//   https://codeforces.com/blog/entry/83969
//   https://www.luogu.com.cn/problem/P4556
//   https://www.luogu.com.cn/problem/P5298
//   https://codeforces.com/problemset/problem/600/E
// rt = rt.merge(rt2)
func (o *stNode) merge(b *stNode) *stNode {
	if o == nil {
		return b
	}
	if b == nil {
		return o
	}
	if o.l == o.r {
		// 按照所需合并，如加法
		o.val += b.val
		return o
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	o.maintain()
	return o
}

// EXTRA: 线段树分裂
// 将区间 [l,r] 从 o 中分离到 b 上
// https://www.luogu.com.cn/blog/cyffff/talk-about-segument-trees-split
// https://www.luogu.com.cn/problem/P5494
// rt, rt2 := rt.split(nil, l, r)
func (o *stNode) split(b *stNode, l, r int) (*stNode, *stNode) {
	if o == emptyStNode || l > o.r || r < o.l {
		return o, emptyStNode
	}
	if l <= o.l && o.r <= r {
		return emptyStNode, o
	}
	if b == emptyStNode {
		b = &stNode{lo: emptyStNode, ro: emptyStNode, l: o.l, r: o.r, val: stNodeDefaultVal}
	}
	o.lo, b.lo = o.lo.split(b.lo, l, r)
	o.ro, b.ro = o.ro.split(b.ro, l, r)
	o.maintain()
	b.maintain()
	return o, b
}

// 权值线段树求第 k 小
// 调用前需保证 1 <= k <= root.val
func (o *stNode) kth(k int) int {
	if o.l == o.r {
		return o.l
	}
	cntL := o.lo.val
	if k <= cntL {
		return o.lo.kth(k)
	}
	return o.ro.kth(k - cntL)
}

//

// 线段树分治
// todo https://www.luogu.com.cn/problem/P5787
//  https://codeforces.com/problemset/problem/1140/F 2600

// 可持久化线段树（又称函数式线段树、主席树） Persistent Segment Tree
// https://oi-wiki.org/ds/persistent-seg/
// 静态+动态 https://www.acwing.com/blog/content/487/
// https://zhuanlan.zhihu.com/p/250565583
// https://blog.csdn.net/weixin_43914593/article/details/108861279
//
// 数组写法 https://codeforces.com/problemset/submission/840/254783792 
//        https://codeforces.com/problemset/submission/323/250523407
// 数组大小 n * (bits.Len(n-1) + 3)
//
// 另见 union_find.go 中的「可持久化并查集」
//
// 模板题 https://www.luogu.com.cn/problem/P3919
//       https://www.luogu.com.cn/problem/P3834
//       https://ac.nowcoder.com/acm/contest/7613/C
// https://atcoder.jp/contests/abc253/tasks/abc253_f 区间更新单点查询 
// https://codeforces.com/problemset/problem/1262/D2 1800 *在线做法 
// https://codeforces.com/problemset/problem/484/E 2500 二分，转换成找最长的已填入数字的区间，做法类似最大子段和 
// https://codeforces.com/problemset/problem/840/D 2500
// todo 种类数 https://codeforces.com/problemset/problem/620/E
//  https://codeforces.com/problemset/problem/786/C
//  差分 https://codeforces.com/problemset/problem/813/E
//  https://codeforces.com/problemset/problem/837/G
//  https://codeforces.com/problemset/problem/893/F 2300 与 DFS序+深度 结合
//  https://codeforces.com/problemset/problem/961/E（不止一种做法）
type pstNode struct {
	lo, ro *pstNode
	l, r   int // 注：如果 MLE 请换成传参的写法，或者使用数组版本
	sum    int
}

func (pstNode) mergeInfo(a, b int) int {
	return a + b // 根据题目修改
}

// t := make([]*pstNode, 1, maxVersion+1)
// t[0] = buildPST(a, 0, len(a)-1)
// 或者 t := []*pstNode{buildPST(a, 0, len(a)-1)}
func buildPST(a []int, l, r int) *pstNode {
	o := &pstNode{l: l, r: r}
	if l == r {
		o.sum = a[l]
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPST(a, l, m)
	o.ro = buildPST(a, m+1, r)
	o.maintain()
	return o
}

// 一般写法是更新到当前版本，然后把返回的新版本加在 t 的末尾，即
// t = append(t, t[len(t)-1].modify(i, add))
// t[i] = t[i-1].modify(i, add)
// 注意为了拷贝一份 pstNode，这里的接收器不是指针
func (o pstNode) modify(i int, add int) *pstNode {
	if o.l == o.r {
		o.sum += add
		return &o
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		o.lo = o.lo.modify(i, add)
	} else {
		o.ro = o.ro.modify(i, add)
	}
	o.maintain()
	return &o
}

func (o *pstNode) maintain() {
	o.sum = o.mergeInfo(o.lo.sum, o.ro.sum)
}

func (o *pstNode) queryRange(l, r int) int {
	if l <= o.l && o.r <= r {
		return o.sum
	}
	m := (o.l + o.r) >> 1
	if r <= m {
		return o.lo.queryRange(l, r)
	}
	if m < l {
		return o.ro.queryRange(l, r)
	}
	vl := o.lo.queryRange(l, r)
	vr := o.ro.queryRange(l, r)
	return o.mergeInfo(vl, vr)
}

// 区间更新（只能配合单点查询）
// 调用前需要拷贝一份 root 节点
// 需要保证 add 非负
// https://atcoder.jp/contests/abc253/tasks/abc253_f
func (o *pstNode) update(l, r int, add int) {
	if l <= o.l && o.r <= r {
		o.sum += add
		return
	}
	lo := *o.lo
	o.lo = &lo
	ro := *o.ro
	o.ro = &ro
	if add := o.sum; add > 0 {
		o.lo.sum += add
		o.ro.sum += add
		o.sum = 0
	}
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if r > m {
		o.ro.update(l, r, add)
	}
}

// 单点查询，配合上面的区间更新使用
func (o *pstNode) querySingle(i int) int {
	if o.l == o.r {
		return o.sum
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		return o.sum + o.lo.querySingle(i)
	}
	return o.sum + o.ro.querySingle(i)
}

// 主席树相当于对数组的每个前缀建立一棵线段树
// 离散化时，求 kth 需要将相同元素也视作不同的
// 附：Wavelet Trees https://codeforces.com/blog/entry/52854 https://ideone.com/Tkters
// 题目见上

// EXTRA: 查询区间 [l,r] 中第 k 小（k 从 1 开始）的数
// 初始 t[0] = buildPST(1, len(a))
//     t[i+1] = t[i].update(kth[i], 1)   kth[i] 为 a[i] 离散化后的值（从 1 开始）
// 查询 t[r].kth(t[l-1], k)               类似前缀和 [l,r] 1<=l<=r<=n
// https://www.luogu.com.cn/problem/P2617
// https://codeforces.com/problemset/problem/840/D
func (o *pstNode) kth(old *pstNode, k int) int {
	if o.l == o.r {
		return o.l
	}
	cntL := o.lo.sum - old.lo.sum
	if k <= cntL {
		return o.lo.kth(old.lo, k)
	}
	return o.ro.kth(old.ro, k-cntL)
}

// EXTRA: 查询区间 [l,r] 中 val 的出现次数
// t[r].query(t[l-1], val)
// https://codeforces.com/problemset/problem/840/D
func (o *pstNode) query(old *pstNode, val int) int {
	if o.l == o.r {
		return o.sum - old.sum
	}
	m := (o.l + o.r) >> 1
	if val <= m {
		return o.lo.query(old.lo, val)
	}
	return o.ro.query(old.ro, val)
}

// todo EXTRA: rank
//  二分答案？

// 在线二维数点
// 查询区间 [l,r] 中在 [low,high] 范围内的元素个数
// low 和 high 为离散化后的值（从 1 开始）
// 离线二维数点见 fenwick_tree.go 中的 areaPointCountOffline
// https://codeforces.com/problemset/problem/323/C 2400
// https://codeforces.com/problemset/problem/538/F 2200
// http://acm.hdu.edu.cn/showproblem.php?pid=4417
// 调用方式同上 t[r].countRange(t[l-1], low, high)
func (o *pstNode) countRange(old *pstNode, low, high int) int {
	if high < o.l || o.r < low {
		return 0
	}
	if low <= o.l && o.r <= high {
		return o.sum - old.sum
	}
	m := (o.l + o.r) >> 1
	if high <= m {
		return o.lo.countRange(old.lo, low, high)
	}
	if m < low {
		return o.ro.countRange(old.ro, low, high)
	}
	return o.lo.countRange(old.lo, low, high) + o.ro.countRange(old.ro, low, high)
}

// EXTRA: countDiff 区间不同元素个数
// 做法是维护左侧相同元素的位置
// todo 参考 https://www.luogu.com.cn/problem/solution/SP3267
// https://www.luogu.com.cn/problem/P1972
// https://www.luogu.com.cn/problem/SP3267

// EXTRA: 区间绝对众数及其出现次数（注意返回的众数是离散化后的值）
// 没找到时返回的 mode = -1
// 题目链接见最上
// 调用方式 t[r].countMode(t[l-1], k)，取 k = (r-l+2)/2，即区间长度一半的上取整
// 这种 k 值使得递归时左右儿子至多调用其中一个
func (o *pstNode) countMode(old *pstNode, k int) (mode, count int) {
	if o.l == o.r {
		return o.l, o.sum - old.sum
	}
	if o.lo.sum-old.lo.sum >= k {
		return o.lo.countMode(old.lo, k)
	}
	if o.ro.sum-old.ro.sum >= k {
		return o.ro.countMode(old.ro, k)
	}
	return -1, 0
}

//

// 动态开点主席树
// 如果 TLE 可以禁用垃圾回收 func init() { debug.SetGCPercent(-1) } 能快一倍
// 模板题 https://atcoder.jp/contests/abc339/tasks/abc339_g
// - 使用方法见代码 https://atcoder.jp/contests/abc339/submissions/50126709
// https://codeforces.com/problemset/problem/1771/F 代码 https://codeforces.com/contest/1771/submission/245657179

//

// EXTRA: 子序列自动机 Sequence Automation
// https://www.luogu.com.cn/problem/P5826

// EXTRA: 标记永久化
// todo 一本通 p239

//

// 李超线段树
// 用来维护在平面直角坐标系上的线段关系
// todo https://oi-wiki.org/ds/li-chao-tree/
//  https://zhuanlan.zhihu.com/p/64946571
//  https://www.luogu.com.cn/blog/fzber0103/Li-Chao-Tree
//  模板题 https://www.luogu.com.cn/problem/P4097
//  https://codeforces.com/contest/1303/problem/G

//

// Wavelet Trees
// todo https://www.geeksforgeeks.org/wavelet-trees-introduction/
// https://codeforces.com/blog/entry/52854
// https://www.youtube.com/watch?v=4aSv9PcecDw

// MISC: 线段树最下面一排的节点编号之和
// https://www.luogu.com.cn/problem/P9689?contestId=133572
