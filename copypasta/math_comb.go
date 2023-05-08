package copypasta

import "math/bits"

/* 组合数学

概率与期望见 dp.go 中的「概率 DP」
部分计数问题见 dp.go 中的「计数 DP」

https://en.wikipedia.org/wiki/Combination
https://en.wikipedia.org/wiki/Enumerative_combinatorics
https://en.wikipedia.org/wiki/Binomial_theorem

一些组合问题
没有思路的话可以尝试：
- 打表 + OEIS
- 用 DP 推导，然后尝试优化
https://codeforces.com/problemset/problem/1391/C 1500
https://codeforces.com/problemset/problem/213/B
https://codeforces.com/problemset/problem/300/C
https://codeforces.com/problemset/problem/520/E
https://codeforces.com/problemset/problem/559/C
https://codeforces.com/problemset/problem/869/C
https://codeforces.com/problemset/problem/1204/E 推荐
https://codeforces.com/problemset/problem/1261/D2 推荐
https://codeforces.com/problemset/problem/1288/C
https://codeforces.com/problemset/problem/1342/E
https://codeforces.com/problemset/problem/1359/E
https://codeforces.com/problemset/problem/1761/D https://www.luogu.com.cn/blog/linyihdfj/solution-cf1761d https://www.cnblogs.com/linyihdfj/p/16893607.html
https://codeforces.com/problemset/problem/1763/D 推荐 分类讨论
https://atcoder.jp/contests/abc171/tasks/abc171_f 推荐 巧妙去重

放球问题
最基础的问题，把 n 个无区别的球放入 m 个有区别的盒子中，不允许空盒（n>=m）：
- 解答：考虑用 m-1 个隔板隔开这些球，这些球之间有 n-1 个位置可以放置隔板，所以方案数为 C(n-1,m-1)
变形：允许空盒
- 解答：假设多了 m 个球，往每个盒子中都放一个球，就可以转换成上面的情况 C(n+m-1,m-1)
https://baike.baidu.com/item/%E6%94%BE%E7%90%83%E9%97%AE%E9%A2%98
https://www.luogu.com.cn/blog/over-knee-socks/post-ball-box
https://www.cnblogs.com/Xing-Ling/p/11176939.html
https://blog.csdn.net/weixin_33759269/article/details/86017932
扩展例题 https://codeforces.com/problemset/problem/893/E
todo 十二重计数法 https://www.luogu.com.cn/problem/P5824

todo 组合数性质 | 二项式推论 https://oi-wiki.org/math/combination/#_13
todo NOI 一轮复习 IV：组合计数 https://www.luogu.com.cn/blog/ix-35/noi-yi-lun-fu-xi-iv-zu-ge-ji-shuo
一些常用组合恒等式的解释 https://www.zhihu.com/question/26094736 https://zhuanlan.zhihu.com/p/82241906
递推式 C(n-1, k-1) + C(n-1, k) = C(n, k)
上项求和 C(r, r) + C(r+1, r) + ... + C(n, r) = C(n+1, r+1)
上式亦为 C(n, 0) + C(n+1, 1) + ... + C(n+m, m) = C(n+m+1, m)
   https://atcoder.jp/contests/abc154/tasks/abc154_f
   https://codeforces.com/contest/1696/problem/E
范德蒙德恒等式 Vandermonde's identity https://en.wikipedia.org/wiki/Vandermonde%27s_identity
∑i=[0..k] C(n,i)*C(m,k-i) = C(n+m,k)   https://www.luogu.com.cn/problem/P7386
特别地：∑i=[0..m] C(n,i)*C(m,i) = ∑i=[0..m] C(n,i)*C(m,m-i) = C(n+m,m)   https://codeforces.com/problemset/problem/785/D
∑i>=n and k-i>=m C(i,n)*C(k-i,m) = C(k+1,n+m+1)   https://www.luogu.com.cn/blog/hanzhongtlx/ti-xie-0-1-trie
组合恒等式之万金油方法 https://zhuanlan.zhihu.com/p/25195967
∑i*C(n,i) = n*2^(n-1)
组合数奇偶性：n&m==m 时 C(n,m) 为奇数，否则为偶数
联立 (1+1)^n 和 (1+(-1))^n 的二项式展开，可得 ∑C(n,2k+1) = ∑C(n,2k) = 2^(n-1)
https://oeis.org/A000244 3^n 子集的子集个数

NOTE: 涉及到相邻的组合问题：可以考虑当前位置和左侧位置所满足的性质（例题 https://atcoder.jp/contests/abc167/tasks/abc167_e）

杨辉三角每行之积 https://oeis.org/A001142 a(n) = Product_{k=1..n} k^(2k-1-n)
A001142(n) 为奇数时的 n 是 2^k - 1 形式的

https://oeis.org/A002109 Hyperfactorials: Product_{k=1..n} k^k

隔三组合数 https://oeis.org/A024493 https://oeis.org/A024494 https://oeis.org/A024495 C(n,0) + C(n,3) + ... + C(n,3[n/3])
隔四组合数 https://oeis.org/A038503 https://oeis.org/A038504 https://oeis.org/A038505 https://oeis.org/A000749

Tetrahedral (or triangular pyramidal) numbers: a(n) = C(n+2,3) = n*(n+1)*(n+2)/6 https://oeis.org/A000292
a(n) = Sum_{1<=i<=j<=n} j-i
a(n) = sum of all the possible products p*q where (p,q) are ordered pairs and p + q = n + 1
a(n) = 长度为 n 的字符串的所有子串长度之和

隔板法 https://zh.wikipedia.org/wiki/%E9%9A%94%E6%9D%BF%E6%B3%95

todo 可重集排列组合 https://oi-wiki.org/math/combination/
todo https://codeforces.com/problemset/problem/451/E
不相邻的排列 https://oi-wiki.org/math/combination/#_10
错排 https://oeis.org/A000166 subfactorial numbers  a[n]=(n-1)*(a[n-1]+a[n-2])  https://zh.wikipedia.org/wiki/%E9%94%99%E6%8E%92%E9%97%AE%E9%A2%98
	https://oeis.org/A082491 n! * A000166(n)   a(n+2) = (n+2)*(n+1)*(a(n+1)+(n+1)*a(n))
	https://oeis.org/A000255 错排的比较对象的范围是 [1,n+1]  a(n) = n*a(n-1) + (n-1)*a(n-2), a(0) = a(1) = 1
	https://oeis.org/A000153 错排的比较对象的范围是 [1,n+2]  a(n) = n*a(n-1) + (n-2)*a(n-2), a(0) = 0, a(1) = 1
	https://oeis.org/A000261 错排的比较对象的范围是 [1,n+3]  a(n) = n*a(n-1) + (n-3)*a(n-2), a(1) = 0, a(2) = 1
	https://oeis.org/A001909 错排的比较对象的范围是 [1,n+4]  a(n) = n*a(n-1) + (n-4)*a(n-2), a(2) = 0, a(3) = 1
		https://atcoder.jp/contests/abc172/tasks/abc172_e
    https://oeis.org/A127548 和两个排列都不同的错排数（这两个排列也互为错排）
圆排列 https://zh.wikipedia.org/wiki/%E5%9C%86%E6%8E%92%E5%88%97
    Q(n,n) = (n-1)!

https://oeis.org/A000522 Total number of arrangements of a set with n elements: a(n) = Sum_{k=0..n} n!/k!    Total number of permutations of all subsets of an n-set
                          a(n) = n*a(n-1) + 1, a(0) = 1
                               = floor(e * n!)
https://oeis.org/A007526 A000522(n)-1 去掉空集
https://oeis.org/A030297 a(n) = sum_{k=0...n} (n! / k!) * k^2     For n>=2, a(n) = floor(2*e*n! - n - 2)
https://oeis.org/A019461 Add 1, multiply by 1, add 2, multiply by 2, etc.; start with 0

二阶递推数列通项 https://zhuanlan.zhihu.com/p/75096951
凯莱公式 Cayley’s formula: the number of trees on n labeled vertices is n^(n-2).
普吕弗序列 Prüfer sequence: 由树唯一地产生的序列
约瑟夫问题 Josephus Problem https://cp-algorithms.com/others/josephus_problem.html https://en.wikipedia.org/wiki/Josephus_problem
Stern-Brocot 树与 Farey 序列 https://oi-wiki.org/misc/stern-brocot/ https://cp-algorithms.com/others/stern_brocot_tree_farey_sequences.html

* 生成函数/母函数 *
https://en.wikipedia.org/wiki/Generating_function
https://oi-wiki.org/math/gen-func/intro/
todo 一些常见数列的生成函数推导 https://www.luogu.com.cn/blog/nederland/girl-friend

整数分拆 https://oeis.org/A000041 https://en.wikipedia.org/wiki/Partition_(number_theory)
五边形数与整数拆分问题
    https://en.wikipedia.org/wiki/Pentagonal_number_theorem
    https://en.wikipedia.org/wiki/Fermat_polygonal_number_theorem
    https://studyingfather.com/archives/3000
    https://blog.csdn.net/visit_world/article/details/52734860
    相关题目 https://www.luogu.com.cn/problem/P6189
https://oeis.org/A104513 The number of consecutive integers > 1 beginning with A104512(n), the sum of which equals n, or 0 if impossible.
						a(n)=0 iff n=2^k
https://oeis.org/A069283 将 n 分拆成至少两个连续整数的方法数 = n 的奇因子数 - 1
						见上面的 oddDivisorsNum 函数
https://oeis.org/A018819 Binary partition function: number of partitions of n into powers of 2
	相关题目 https://www.luogu.com.cn/problem/P6065 http://poj.org/problem?id=2229
https://oeis.org/A000740 将 n 分拆成若干互质整数的方法数 a(n) = sum_{d|n} mu(n/d)*2^(d-1)

华林问题 Waring's problem
https://en.wikipedia.org/wiki/Waring%27s_problem
https://oeis.org/A002804 (Presumed) solution to Waring's problem: g(k) = 2^k + floor((3/2)^k) - 2
g(k) is the smallest number s such that every natural number is the sum of at most s k-th powers of natural numbers
k=2 https://oeis.org/A002828 Least number of squares that add up to n
	https://oeis.org/A000415 Numbers that are the sum of 2 but no fewer nonzero squares
	https://oeis.org/A000419 Numbers that are the sum of 3 but no fewer nonzero squares
	https://oeis.org/A004215 Numbers that are the sum of 4 but no fewer nonzero squares
	四平方和定理 Lagrange's four-square theorem https://en.wikipedia.org/wiki/Lagrange%27s_four-square_theorem
	https://oeis.org/A006431 Numbers that have a unique partition into a sum of four non-negative squares
k=3 https://oeis.org/A002376 Least number of positive cubes needed to represent n
k=4 https://oeis.org/A002377 Least number of 4th powers needed to represent n

	贪心分拆
	Number of positive k-th powers needed to sum to n using the greedy algorithm
	k=2 https://oeis.org/A053610 Records https://oeis.org/A006892
	k=3 https://oeis.org/A055401 Records https://oeis.org/A055402
		相关题目 https://codeforces.com/problemset/problem/679/B

质数分拆
https://oeis.org/A000607 Number of partitions of n into prime parts
	https://www.luogu.com.cn/problem/P1832
https://oeis.org/A061358 Number of ways of writing n=p+q with p, q primes and p>=q
https://oeis.org/A067187 Numbers that can be expressed as the sum of two primes in exactly one way
https://oeis.org/A068307 number of partitions of n into a sum of three primes
https://oeis.org/A071335 Number of partitions of n into a sum of at most three primes
https://oeis.org/A023022 Number of partitions of n into two relatively prime parts
	a(n) = phi(n)/2 for n >= 3

https://oeis.org/A000404 Numbers that are the sum of 2 nonzero squares
https://oeis.org/A003325 Numbers that are the sum of 2 positive cubes

https://oeis.org/A000081 Number of unlabeled rooted trees with n nodes (or connected functions with a fixed point)

Maximum product of two integers whose sum is n https://oeis.org/A002620
Quarter-squares: floor(n/2)*ceiling(n/2). Equivalently, floor(n^2/4)
https://oeis.org/A024206 = A002620(n+1)-1 = floor((n-1)(n+3)/4)

	Maximal product of three numbers with sum n: a(n) = max(r*s*t), n = r+s+t https://oeis.org/A006501
	a(n) = floor(n/3)*floor((n+1)/3)*floor((n+2)/3)
	Expansion of (1+x^2) / ( (1-x)^2 * (1-x^3)^2 )

	Maximal product of four nonnegative integers whose sum is n https://oeis.org/A008233
	a(n) = floor(n/4)*floor((n+1)/4)*floor((n+2)/4)*floor((n+3)/4)

	...

	相关题目 https://codeforces.com/problemset/problem/1368/B

没有相邻元素差值为 1 的排列个数
https://oeis.org/A002464 Hertzsprung's problem: ways to arrange n non-attacking kings on an n X n board, with 1 in each row and column
Also number of permutations of length n without rising or falling successions
if n = 0 or 1 then a(n) = 1
if n = 2 or 3 then a(n) = 0
otherwise a(n) = (n+1)*a(n-1) - (n-2)*a(n-2) - (n-5)*a(n-3) + (n-3)*a(n-4)
https://oeis.org/A129535 补集
https://oeis.org/A086852 恰有一个相邻元素差值为 1 的排列个数
https://oeis.org/A086853 恰有两个相邻元素差值为 1 的排列个数

记 A = [1,2,...,n]，A 的全排列中与 A 的最大差值为 n^2/2 https://oeis.org/A007590
Maximum sum of displacements of elements in a permutation of (1..n)
For example, with n = 9, permutation (5,6,7,8,9,1,2,3,4) has displacements (4,4,4,4,4,5,5,5,5) with maximal sum = 40

https://oeis.org/A176127 The number of permutations of {1,2,...,n,1,2,...,n} with the property that there are k numbers between the two k's in the set for k=1,...,n
相关题目：《程序员的算法趣题》Q53 同数包夹

n married couples are seated in a row so that every wife is to the left of her husband
若不考虑顺序，则所有排列的个数为 (2n)!
考虑顺序可以发现，对于每一对夫妻来说，妻子在丈夫左侧的情况和在右侧的情况相同且不同对夫妻之间是独立的
因此每有一对夫妻，符合条件的排列个数就减半
所以结果为 a(n) = (2n)!/2^n https://oeis.org/A000680
或者见这道题目的背景 LC1359 https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/

NxN 大小的对称置换矩阵的个数 https://oeis.org/A000085
这里的对称指仅关于主对角线对称
a[i] = (a[i-1] + (i-1)*a[i-2]) % mod
The number of n X n symmetric permutation matrices
Number of self-inverse permutations on n letters, also known as involutions; number of standard Young tableaux with n cells
Proof of the recurrence relation a(n) = a(n-1) + (n-1)*a(n-2):
	number of involutions of [n] containing n as a fixed point is a(n-1);
	number of involutions of [n] containing n in some cycle (j, n),
	where 1 <= j <= n-1, is (n-1) times the number of involutions of [n] containing the cycle (n-1 n) = (n-1)*a(n-2)
相关题目 https://ac.nowcoder.com/acm/contest/5389/C

The number of 3 X n matrices of integers for which the upper-left hand corner is a 1,
the rows and columns are weakly increasing, and two adjacent entries differ by at most 1
a(n+2) = 5*a(n+1) - 2*a(n), with a(0) = 1, a(1) = 4
https://oeis.org/A052913
相关题目 LC1411 https://leetcode-cn.com/problems/number-of-ways-to-paint-n-x-3-grid/

男厕问题 / 电话问题 https://oeis.org/A185456
Assume that the first person to use a bank of payphones selects one at the end,
and all subsequent users select the phone which puts them farthest from the current phone users.
U(n) is the smallest number of phones such that n may be used without any two adjacent phones being used
https://www.zhihu.com/question/278361000/answer/1004606685

https://oeis.org/A089934 Table T(n,k) of the number of n X k matrices on {0,1} without adjacent 0's in any row or column
https://oeis.org/A006506 上面这个 table 的对角线
    Number of configurations of non-attacking princes on an n X n board, where a "prince" attacks the four adjacent (non-diagonal) squares
    Also number of independent vertex sets in an n X n grid

https://oeis.org/A001224 The number of inequivalent ways to pack a 2 X n rectangle with dominoes
    If F(n) is the n-th Fibonacci number, then a(2n) = (F(2n+1) + F(n+2))/2 and a(2n+1) = (F(2n+2) + F(n+1))/2
    https://oeis.org/A060312

https://oeis.org/A001187 n 个节点的无向连通图的个数  Number of connected labeled graphs with n nodes
相关题目：https://www.acwing.com/problem/content/309/

韦德伯恩-埃瑟林顿数
https://oeis.org/A001190 Wedderburn-Etherington numbers: unlabeled binary rooted trees (every node has out-degree 0 or 2) with n endpoints (and 2n-1 nodes in all)
https://en.wikipedia.org/wiki/Wedderburn%E2%80%93Etherington_number
https://mathworld.wolfram.com/WeaklyBinaryTree.html 给出了如下公式：
    a(2n-1) = a(1)a(2n-2) + a(2)a(2n-3) + ... + a(n-1)a(n)
    a(2n)   = a(1)a(2n-1) + a(2)a(2n-2) + ... + a(n-1)a(n+1) + a(n)(a(n)+1)/2
https://oeis.org/A000598 Number of rooted ternary trees with n nodes
https://oeis.org/A268172 Binary-ternary Wedderburn-Etherington numbers
相关题目：《程序员的算法趣题》Q30 用插线板制作章鱼脚状线路

https://oeis.org/A003991 Multiplication table read by antidiagonals = (n-k+1)*k
https://oeis.org/A059036 = A003991(n, k) - 1

一些二进制的计数问题见 bits.go

CF 上的一些组合计数问题 http://blog.miskcoo.com/2015/06/codeforces-combinatorics-and-probabilities-problem

置换群、Burnside 引理与 Pólya 定理          Polya 计数
https://en.wikipedia.org/wiki/P%C3%B3lya_enumeration_theorem
https://oi-wiki.org/math/permutation-group/
todo https://atcoder.jp/contests/abc198/tasks/abc198_f
     https://oeis.org/A054473 Number of ways of numbering the faces of a cube with nonnegative integers so that the sum of the 6 numbers is n

找出 50% 作弊者 GCJ2021 QR https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d1155
讨论 https://codeforces.com/blog/entry/84822
*/

func pow(x int64, n int) (res int64) {
	x %= _mod
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % _mod
		}
		x = x * x % _mod
	}
	return
}

// 一种避免不小心把数组开小的写法
// https://codeforces.com/problemset/submission/1794/205053722
type comb struct{ _f, _invF []int64 }

func newComb(mx int) *comb {
	c := &comb{[]int64{1}, []int64{1}}
	c._init(mx)
	return c
}

func (c *comb) _init(mx int) {
	n := len(c._f)
	c._f = append(make([]int64, 0, mx+1), c._f...)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * int64(i) % _mod
	}
	c._invF = append(make([]int64, 0, mx+1), c._invF...)[:mx+1]
	c._invF[mx] = pow(c._f[mx], _mod-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * int64(i) % _mod
	}
}

func (c *comb) f(n int) int64 {
	if n >= len(c._f) {
		c._init(n * 2)
	}
	return c._f[n]
}

func (c *comb) invF(n int) int64 {
	if n >= len(c._f) {
		c._init(n * 2)
	}
	return c._invF[n]
}

func (c *comb) c(n, k int) int64 {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % _mod * c.invF(n-k) % _mod
}

// 容斥原理 (PIE, the principle of inclusion and exclusion)
// 参考《挑战程序设计竞赛》P296
// https://codeforces.com/blog/entry/64625
// https://ac.nowcoder.com/acm/contest/6219/C
//
// 多重集组合数 https://codeforces.com/problemset/problem/451/E
// https://codeforces.com/problemset/problem/1342/E
// 如何将问题转化成可以容斥的结构 https://codeforces.com/problemset/problem/1228/E
// 不重不漏 https://codeforces.com/problemset/problem/1007/B
// 与 SOS DP 结合 https://codeforces.com/problemset/problem/449/D
// 用因子容斥 https://codeforces.com/problemset/problem/900/D
func solveInclusionExclusion(a []int) (ans int64) {
	n := len(a)
	const mod int64 = 1e9 + 7 // 998244353
	for sub := uint(0); sub < 1<<n; sub++ {
		res := int64(0)
		for i, v := range a {
			if sub>>i&1 > 0 {
				// 视情况而定，有时候包含元素 i 表示考虑这种情况，有时候表示不考虑这种情况
				_ = v // do v...

			}
		}
		if bits.OnesCount(sub)&1 > 0 { // 某些题目是 == 0
			res = -res
		}
		ans += res // mod
	}
	ans = (ans%mod + mod) % mod
	return
}
