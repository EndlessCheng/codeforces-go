package copypasta

import (
	"maps"
	"math"
	"slices"
	"sort"
)

/* 树状数组（Fenwick Tree），二叉索引树（Binary Index Tree, BIT）
https://en.wikipedia.org/wiki/Fenwick_tree
【推荐阅读】带你发明树状数组！https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/

可视化 https://visualgo.net/zh/fenwicktree

todo 树状数组延申应用 https://www.luogu.com.cn/blog/kingxbz/shu-zhuang-shuo-zu-zong-ru-men-dao-ru-fen
 浅谈树状数组的优化及扩展 https://www.luogu.com.cn/blog/countercurrent-time/qian-tan-shu-zhuang-shuo-zu-you-hua
 浅谈树状数组套权值树 https://www.luogu.com.cn/blog/bfqaq/qian-tan-shu-zhuang-shuo-zu-quan-zhi-shu
https://oi-wiki.org/ds/bit/
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FenwickTree.java.html

模板题
LC307 https://leetcode.cn/problems/range-sum-query-mutable/
https://www.luogu.com.cn/problem/P3374

关于逆序对，见下面的 cntInversions

https://codeforces.com/problemset/problem/1234/D 1600
https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
https://atcoder.jp/contests/arc075/tasks/arc075_c
静态区间种类 - 离线做法
    https://www.luogu.com.cn/problem/P1972
    https://atcoder.jp/contests/abc174/tasks/abc174_f
    https://codeforces.com/problemset/problem/246/E
置换 LC2179 https://leetcode.cn/problems/count-good-triplets-in-an-array/
- 同样的置换思想 LC1713 https://leetcode.cn/problems/minimum-operations-to-make-a-subsequence/
题目推荐 https://cp-algorithms.com/data_structures/fenwick.html#toc-tgt-12
树状数组的性质能使其支持动态 [1,x] 或 [x,n] 范围上的最值更新查询等操作
    https://codeforces.com/problemset/problem/629/D
    https://codeforces.com/problemset/problem/1635/F
好题 https://www.luogu.com.cn/problem/P2345 https://www.luogu.com.cn/problem/P5094
多变量统计 https://codeforces.com/problemset/problem/1194/E
         T4 https://www.nowcoder.com/discuss/1022136
最多交换 k 次相邻字母后，得到的最小字典序
- LC1505 https://leetcode.cn/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/
LC2921 https://leetcode.cn/problems/maximum-profitable-triplets-with-increasing-prices-ii/

https://codeforces.com/problemset/problem/1915/F 1500
https://codeforces.com/problemset/problem/1234/D 1600
https://atcoder.jp/contests/abc157/tasks/abc157_e 1443~CF1600
https://codeforces.com/problemset/problem/627/B 1700 模板题
https://codeforces.com/problemset/problem/652/D 1800 区间包含计数
https://codeforces.com/problemset/problem/597/C 1900 长为 k 的上升子序列个数
https://codeforces.com/problemset/problem/961/E 1900（不止一种做法）
https://codeforces.com/problemset/problem/2042/D 1900
https://codeforces.com/problemset/problem/629/D 2000
https://codeforces.com/problemset/problem/1891/F 2000 离线 树 回溯
https://codeforces.com/problemset/problem/703/D 2100 区间元素去重后的异或和
- 联系 https://www.luogu.com.cn/problem/P1972
https://codeforces.com/problemset/problem/1660/F2 2100 建模
https://codeforces.com/problemset/problem/301/D 2200 整除对统计
https://codeforces.com/problemset/problem/369/E 2200 区间统计技巧
https://codeforces.com/problemset/problem/762/E 2200 离散化
- https://codeforces.com/problemset/problem/1045/G 2200 同 762E
https://codeforces.com/problemset/problem/1194/E 2200 多变量统计
https://codeforces.com/problemset/problem/1167/F 2300
https://codeforces.com/problemset/problem/1967/C 2300
https://codeforces.com/problemset/problem/12/D 2400 三维偏序
https://codeforces.com/problemset/problem/246/E 2400
https://codeforces.com/problemset/problem/1334/F 2500
https://codeforces.com/problemset/problem/1635/F 2800

https://atcoder.jp/contests/abc392/tasks/abc392_f 也可以用 Splay 树
https://atcoder.jp/contests/abc256/tasks/abc256_f 多重前缀和
https://www.lanqiao.cn/problems/5131/learning/?contest_id=144
贡献 https://www.lanqiao.cn/problems/12467/learning/?contest_id=167
https://codeforces.com/gym/101649 I 题
http://poj.org/problem?id=2155
http://poj.org/problem?id=2886
*/

const fenwickInitVal = 0 // -1e18

type fenwick []int

func newFenwickTree(n int) fenwick {
	t := make(fenwick, n+1)
	for i := range t {
		t[i] = fenwickInitVal
	}
	return t
}

func (fenwick) op(a, b int) int {
	return a + b // max(a, b)
}

// a[i] 增加 val
// 1<=i<=n
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = f.op(f[i], val)
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1<=i<=n
func (f fenwick) pre(i int) int {
	res := fenwickInitVal
	i = min(i, len(f)-1)
	for ; i > 0; i &= i - 1 {
		res = f.op(res, f[i])
	}
	return res
}

// 求区间和 a[l] + ... + a[r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

// 静态二维数点
// 对于每个询问，计算 [x1,x2] x [y1,y2] 中的点的个数
// 离线，拆分成两个更小的询问：[1,x2] x [y1,y2] 中的点的个数，减去 [1,x1-1] x [y1,y2] 中的点的个数
// 一边从小到大枚举 x，一边更新【值域树状数组】，一边回答离线后的询问
// https://www.luogu.com.cn/problem/P2163
// LC3382 https://leetcode.cn/problems/maximum-area-rectangle-with-point-constraints-ii/
// https://codeforces.com/problemset/problem/1899/G 1900
func areaPointCountOffline(points []struct{ x, y int }, queries []struct{ x1, x2, y1, y2 int }) []int {
	xMap := map[int][]int{} // 同一列的所有点的纵坐标
	yMap := map[int][]int{} // 同一行的所有点的横坐标（没有用到，可以简单改成 []int，下面排序去重得到 ys）
	for _, p := range points {
		x, y := p.x, p.y
		xMap[x] = append(xMap[x], y)
		yMap[y] = append(yMap[y], x)
	}

	// 离散化用
	xs := slices.Sorted(maps.Keys(xMap))
	ys := slices.Sorted(maps.Keys(yMap))

	// 离线询问
	type data struct{ qid, sign, y1, y2 int }
	qs := make([][]data, len(xs))
	for i, q := range queries {
		x1 := sort.SearchInts(xs, q.x1) // 离散化，下标从 0 开始
		x2 := sort.SearchInts(xs, q.x2+1) - 1
		if x1 > x2 {
			continue
		}
		y1 := sort.SearchInts(ys, q.y1)
		y2 := sort.SearchInts(ys, q.y2+1) - 1
		if y1 > y2 {
			continue
		}
		if x1 > 0 {
			qs[x1-1] = append(qs[x1-1], data{i, -1, y1, y2})
		}
		qs[x2] = append(qs[x2], data{i, 1, y1, y2})
	}

	// 回答询问
	ans := make([]int, len(queries))
	t := make(fenwick, len(ys)+1)
	for i, x := range xs { // 从小到大枚举 x
		// 把横坐标为 x 的所有点都加到树状数组中
		for _, y := range xMap[x] {
			t.update(sort.SearchInts(ys, y)+1, 1) // 离散化，并且下标从 1 开始
		}
		for _, q := range qs[i] {
			// 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
			ans[q.qid] += q.sign * t.query(q.y1+1, q.y2+1) // 下标从 1 开始
		}
	}
	return ans
}

//

// 差分版本
// 参考《算法竞赛进阶指南》《挑战程序设计竞赛》
// 利用差分数组，实现 O(log n) 的区间加、区间查询
// a[1] = diff[1]
// a[2] = diff[1] + diff[2]
// a[m] = diff[1] + ... + diff[m]
// 所以 a[1] + ... + a[m]
//   = ∑(m-i+1)*diff[i]
//   = (m+1)∑diff[i] - ∑i*diff[i]
// https://ac.nowcoder.com/acm/problem/50454
// https://codeforces.com/problemset/problem/383/C 2000
// https://codeforces.com/problemset/problem/916/E 2400
// todo 二维差分 上帝造题的七分钟 https://www.luogu.com.cn/problem/P4514
// todo 离线询问（按 x y 分组）https://codeforces.com/contest/1824/problem/D

// [0] 维护 ∑diff[i]
// [1] 维护 ∑i*diff[i]
// 为了更好地利用缓存，写成 [][2] 而不是 [2][]
type fenwickDiff [][2]int

func newFenwickTreeDiff(n int) fenwickDiff {
	return make(fenwickDiff, n+1)
}

func (t fenwickDiff) _add(i, val int) {
	for iv := i * val; i < len(t); i += i & -i {
		t[i][0] += val
		t[i][1] += iv
	}
}

// a[l] 到 a[r] 增加 val
// 1<=l<=r<=n
func (t fenwickDiff) add(l, r, val int) {
	t._add(l, val)
	t._add(r+1, -val)
}

// 求前缀和 a[1] + ... + a[i]
// 1<=i<=n
func (t fenwickDiff) pre(i0 int) int {
	var s0, s1 int
	for i := i0; i > 0; i &= i - 1 {
		s0 += t[i][0]
		s1 += t[i][1]
	}
	return (i0+1)*s0 - s1
}

// 求区间和 a[l] + ... + a[r]
// 1<=l<=r<=n
func (t fenwickDiff) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

//

// 二维差分树状数组
// https://codeforces.com/problemset/problem/869/E 2400
type fenwickDiff2 [][]int

func newFenwickTreeDiff2(n, m int) fenwickDiff2 {
	t := make(fenwickDiff2, n+1)
	for i := range t {
		t[i] = make([]int, m+1)
	}
	return t
}

func (t fenwickDiff2) add(x, y, val int) {
	for i := x; i < len(t); i += i & -i {
		for j := y; j < len(t[i]); j += j & -j {
			t[i][j] += val
		}
	}
}

// 二维矩阵左上角 (x1,y1) 右下角 (x2,y2) 区域增加 val
// 下标从 1 开始
func (t fenwickDiff2) update(x1, y1, x2, y2, val int) {
	t.add(x1, y1, val)
	t.add(x1, y2+1, -val)
	t.add(x2+1, y1, -val)
	t.add(x2+1, y2+1, val)
}

// 获取二维矩阵 (x,y) 的值
// 下标从 1 开始
func (t fenwickDiff2) get(x, y int) (res int) {
	for i := x; i > 0; i &= i - 1 {
		for j := y; j > 0; j &= j - 1 {
			res += t[i][j]
		}
	}
	return
}

//

// 树套树：树状数组套动态开点权值线段树
// 三维偏序 https://www.luogu.com.cn/problem/P3810 https://www.luogu.com.cn/record/136178821
// 二逼平衡树 https://www.luogu.com.cn/problem/P3380 https://www.luogu.com.cn/record/136286395
// 树状数组在这里就是纯纯工具人，只用来拆分区间
// 注：如果 TLE 可以使用 func init() { debug.SetGCPercent(-1) } 加速
type fenwickWithSeg []*stNode // 见 segment_tree.go

func newFenwickTreeWithSeg(n, mx int) fenwickWithSeg {
	t := make(fenwickWithSeg, n+1)
	for i := range t {
		t[i] = newStRoot(0, mx) // 注意下界
	}
	return t
}

// 二维单点更新：位置 (i,j) 用 val 更新
func (f fenwickWithSeg) update(i, j, val int) {
	for ; i < len(f); i += i & -i {
		f[i].update(j, val)
	}
}

// 二维前缀和：累加所有 x <= i 且 y <= j 的值
func (f fenwickWithSeg) pre(i, j int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i].query(0, j) // 注意下界
	}
	return
}

// 返回 [l,r] 内 v 的排名，即小于 v 的元素个数 + 1
func (f fenwickWithSeg) rank(l, r, v int) int {
	return f.pre(r, v-1) - f.pre(l-1, v-1) + 1
}

// 返回 [l,r] 内的第 k 个数（设为 v），即有 k-1 个数小于 v
// 需要保证 1 <= k <= r-l+1
func (f fenwickWithSeg) kth(l, r, k int) int {
	var ar, al []*stNode // 如果禁止 GC，需要把这行放在外面，然后用下面注释的两行代码（复用内存）
	//ar = ar[:0]
	//al = al[:0]
	for ; r > 0; r &= r - 1 {
		ar = append(ar, f[r])
	}
	for l--; l > 0; l &= l - 1 {
		al = append(al, f[l])
	}
	left, right := 0, f[0].r
	for left < right {
		s := 0
		for _, o := range ar {
			s += o.lo.val
		}
		for _, o := range al {
			s -= o.lo.val
		}
		mid := (left + right) >> 1
		if s >= k {
			for i, o := range ar {
				ar[i] = o.lo
			}
			for i, o := range al {
				al[i] = o.lo
			}
			right = mid
		} else {
			k -= s
			for i, o := range ar {
				ar[i] = o.ro
			}
			for i, o := range al {
				al[i] = o.ro
			}
			left = mid + 1
		}
	}
	return left
}

// 返回 [l,r] 内严格小于 v 的最大的数
// 如果不存在，返回 math.MinInt
func (f fenwickWithSeg) prev(l, r, v int) int {
	rk := f.rank(l, r, v)
	if rk == 1 { // 没有比 v 小的数
		return math.MinInt
	}
	return f.kth(l, r, rk-1)
}

// 返回 [l,r] 内严格大于 v 的最小的数
// 如果不存在，返回 math.MaxInt
func (f fenwickWithSeg) next(l, r, v int) int {
	rk := f.rank(l, r, v+1)
	if rk > r-l+1 { // [l,r] 内的所有数都 <= v
		return math.MaxInt
	}
	return f.kth(l, r, rk)
}

//

func _(n int) {
	tree := make([]int, n+1)
	add := func(i, val int) {
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
	addRange := func(l, r, val int) { add(l, val); add(r+1, -val) } // [l,r]

	// 求权值树状数组第 k 小的数（k > 0）
	// 这里 tree[i] 表示 i 的个数
	// 返回最小的 x 满足 ∑i=[1..x] tree[i] >= k
	// 思路类似倍增的查询，不断寻找 ∑<k 的数，最后 +1 就是答案
	// https://oi-wiki.org/ds/fenwick/#tricks
	//
	// https://codeforces.com/blog/entry/61364
	// https://codeforces.com/problemset/problem/1404/C
	// todo https://codeforces.com/contest/992/problem/E
	// https://atcoder.jp/contests/abc287/tasks/abc287_g
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

	// 输出权值树状数组的 mex（这里的定义是第一个没出现的正数）
	// 注意不能有重复元素
	mex := func() (res int) {
		const log = 17 // bits.Len(uint(n))
		for b := 1 << (log - 1); b > 0; b >>= 1 {
			if next := res | b; tree[next] == next {
				res = next
			}
		}
		res++ // mex
		return
	}

	// 常数优化：O(n) 建树
	// https://oi-wiki.org/ds/fenwick/#tricks
	init := func(a []int) { // len(tree) = len(a) + 1
		for i, v := range a {
			i++
			tree[i] += v
			if j := i + i&-i; j < len(tree) {
				tree[j] += tree[i]
			}
		}
	}

	// a 的下标从 1 开始
	initFrom := func(a []int) {
		for i := 1; i < len(a); i++ {
			if j := i + i&-i; j < len(a) {
				a[j] += a[i]
			}
		}
		tree = a
	}

	// 另外一种写法（效率和算两次一样）
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

	{
		// 时间戳优化（通常用于多组数据+值域树状数组）https://oi-wiki.org/ds/fenwick/#%E6%97%B6%E9%97%B4%E6%88%B3%E4%BC%98%E5%8C%96
		// https://codeforces.com/problemset/submission/1801/205042964
		const mx int = 1e6
		tree := [mx + 1]int{} // 默认都是 0
		time := [mx + 1]int{}
		curCase := 1 // 从 1 开始
		upd := func(i int, v int) {
			for ; i <= mx; i += i & -i {
				if time[i] != curCase {
					time[i] = curCase
					tree[i] = 0 // reset
				}
				tree[i] += v
			}
		}
		pre := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				if time[i] == curCase {
					res += tree[i]
				} // 否则，相当于 res += 0
			}
			return
		}
		_, _ = upd, pre
	}

	// 求逆序对的方法之一
	// 如果 a 范围较大则需要离散化（但这样还不如直接用归并排序）
	// 归并做法见 misc.go 中的 mergeCount
	// LCR 170. 交易逆序对的总数 https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
	// 扩展 https://codeforces.com/problemset/problem/362/C 1900
	// 环形最小逆序对 https://www.luogu.com.cn/problem/solution/P2995
	// 扩展：某些位置上的数待定时的逆序对的期望值 https://codeforces.com/problemset/problem/1096/F
	// https://codeforces.com/problemset/problem/1585/D 1900
	// https://codeforces.com/edu/course/2/lesson/4/3/practice/contest/274545/problem/A
	// 逆序对的奇偶性 https://www.luogu.com.cn/blog/203623/sol-p3760-tjoi2017-yi-huo-hu
	// - https://ac.nowcoder.com/acm/contest/308/D
	// https://codeforces.com/problemset/problem/749/E 期望 贡献
	// https://atcoder.jp/contests/arc075/tasks/arc075_c
	// https://codeforces.com/problemset/problem/911/D
	// https://codeforces.com/contest/987/problem/E
	// https://atcoder.jp/contests/chokudai_S001/tasks/chokudai_S001_l
	// https://atcoder.jp/contests/abc296/tasks/abc296_f
	// https://atcoder.jp/contests/arc136/tasks/arc136_b
	// https://www.codechef.com/problems/DYNAINV?tab=statement
	// https://ac.nowcoder.com/acm/problem/20861
	// 1e9 范围逆序对 https://codeforces.com/problemset/problem/540/E
	// 三元逆序对 https://codeforces.com/problemset/problem/61/E
	// 互质逆序对 小白月赛 87G https://ac.nowcoder.com/acm/contest/73854/G
	cntInversions := func(a []int) int {
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
		invCnt := 0
		for i, v := range a {
			// 由于 i 从 0 开始算，这里先 sum 后 add
			invCnt += i - sum(v)
			add(v)
		}
		return invCnt
	}

	// 通过邻项交换，把数组 a 变成数组 b，需要的最小操作次数
	// 如果无法做到，返回 -1
	// https://atcoder.jp/contests/arc120/tasks/arc120_c
	// LC1850 https://leetcode.cn/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/ 2073
	minSwap := func(a, b []int) (res int) {
		tree := make([]int, len(a)+1)
		add := func(i int) {
			for i++; i < len(tree); i += i & -i {
				tree[i]++
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res += tree[i]
			}
			return
		}

		pos := map[int][]int{}
		for i, v := range a {
			pos[v] = append(pos[v], i)
		}
		for i, v := range b {
			p := pos[v]
			if len(p) == 0 {
				return -1
			}
			j := p[0]
			pos[v] = p[1:]
			res += i - sum(j)
			add(j)
		}
		return
	}

	_ = []interface{}{add, sum, query, addRange, kth, mex, init, initFrom, cntInversions, minSwap}
}

// 给一个数组 a 和一些询问 qs，对每个询问计算 mex(a[l..r])
// a[i]>=0, 1<=l<=r<=n
// 遍历数组 a，记录 a[i] 最后一次出现的位置 lastPos 以及上一个 a[i] 的位置 prevPos
// 建立一个权值树状数组，维护 lastPos[v] 的前缀最小值
// 树状数组维护前缀最小值的条件是每次修改只能往小改，那么从后往前做就好了
// 将询问离线：按照右端点排序（或分组），计算 mex。原理见代码中 query 的注释
// https://www.luogu.com.cn/problem/P4137
// LC2003 https://leetcode.cn/problems/smallest-missing-genetic-value-in-each-subtree/
// - 需要将 a 转换成 DFS 序且从 0 开始，同时最终答案需要 +1
func rangeMex(a []int, qs []struct{ l, r, i int }) []int {
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
