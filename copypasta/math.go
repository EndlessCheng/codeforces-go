package copypasta

import (
	"math"
	"math/big"
	"math/bits"
	"math/rand"
	"sort"
)

/* 数论 组合数学

鸽巢原理 抽屉原理
https://en.wikipedia.org/wiki/Pigeonhole_principle
http://codeforces.com/problemset/problem/1178/E

アルゴリズムと数学 演習問題集 https://atcoder.jp/contests/math-and-algorithm

一些不等式及其证明 https://www.luogu.com.cn/blog/chinesepikaync/oi-zhong-kuai-yong-dao-di-yi-suo-fou-deng-shi-ji-ji-zheng-ming

https://en.wikipedia.org/wiki/List_of_recreational_number_theory_topics
https://euler.stephan-brumme.com/toolbox/

NOTE: a%-b == a%b
NOTE: 对于整数来说有
       ax≤b  =>  x≤⌊b/a⌋       ax<b  =>  x<⌈b/a⌉
       ax>b  =>  x>⌊b/a⌋       ax≥b  =>  x≥⌈b/a⌉
NOTE: ⌊⌊x/n⌋/m⌋ = ⌊x/(n*m)⌋
NOTE: ⌈⌈x/n⌉/m⌉ = ⌈x/(n*m)⌉

https://oeis.org/A257212           Least d>0 such that floor(n/d) - floor(n/(d+1)) <= 1
https://oeis.org/A257213 mex(n/i); Least d>0 such that floor(n/d) = floor(n/(d+1))
另见数论分块

AP: Sn = n*(2*a1+(n-1)*d)/2
GP: Sn = a1*(q^n-1)/(q-1), q!=1
       = a1*n, q==1
∑i*q^(i-1) = n*q^n - (q^n-1)/(q-1)

若干无穷级数之和的公式 https://mathwords.net/mugenwa
∑^∞ r^i = 1/(1-r)
∑^∞ i*r^i = r/(1-r)^2

等幂和 Faulhaber's formula
https://zh.wikipedia.org/wiki/%E7%AD%89%E5%B9%82%E6%B1%82%E5%92%8C#%E4%B8%80%E8%88%AC%E6%95%B0%E5%88%97%E7%9A%84%E7%AD%89%E5%B9%82%E5%92%8C
1^2 + ... + n^2 = n*(n+1)*(2*n+1)/6
1^3 + ... + n^3 = [n(n+1)/2]^2

处理绝对值·曼哈顿距离转切比雪夫距离
每个点 (x,y) 改成 (x+y,x-y)
|x1-x2|+|y1-y2| 就可以用 max(|x1'-x2'|,|y1'-y2'|) 来计算了
https://codeforces.com/problemset/problem/1689/D
LC1131 https://leetcode.cn/problems/maximum-of-absolute-value-expression/

处理绝对值·分类讨论
https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/solution/bu-hui-hua-jian-qing-kan-zhe-pythonjavac-c2s6/

勾股数 https://oeis.org/A008846
斜边 https://oeis.org/A004613 Numbers that are divisible only by primes congruent to 1 mod 4
https://en.wikipedia.org/wiki/Pythagorean_triple https://zh.wikipedia.org/wiki/%E5%8B%BE%E8%82%A1%E6%95%B0

https://oeis.org/A000079 2^n
虽然是个很普通的序列，但也能出现在一些意想不到的地方
例如，在该页面搜索 permutation 可以找到一些有趣的计数问题
a(n) is the number of permutations on [n+1] such that every initial segment is an interval of integers.（每个前缀都对应一段连续的整数）
Example: a(3) counts 1234, 2134, 2314, 2341, 3214, 3241, 3421, 4321.
The map "p -> ascents of p" is a bijection from these permutations to subsets of [n].
An ascent of a permutation p is a position i such that p(i) < p(i+1).
The permutations shown map to 123, 23, 13, 12, 3, 2, 1 and the empty set respectively.
相关题目 https://codeforces.com/problemset/problem/1515/E

https://oeis.org/A001787 n*2^(n-1) = ∑i*C(n,i)   number of ones in binary numbers 1 to 111...1 (n bits)
https://oeis.org/A000337 ∑i*2^(i-1) = (n-1)*2^n+1
https://oeis.org/A036799 ∑i*2^i = (n-1)*2^(n+1)+2 = A000337(n)*2

https://oeis.org/A027992 a(n) = 2^n*(3n-1)+2 = The total sum of squares of parts in all compositions of n （做 https://codeforces.com/problemset/problem/235/B 时找到的序列）
https://oeis.org/A271638 a(n) = (13*n-36)*2^(n-1)+6*n+18 = 	The total sum of the cubes of all parts of all compositions of n

https://oeis.org/A014217 a(n) = floor(phi^n), where phi = (1+sqrt(5))/2 = 1.618...
a(n) = a(n-1) + 2*a(n-2) - a(n-3) - a(n-4)
证明 https://www.luogu.com.cn/discuss/show/318570

https://en.wikipedia.org/wiki/Faulhaber%27s_formula
https://oeis.org/A000330 平方和 = n*(n+1)*(2*n+1)/6
https://oeis.org/A000537 立方和 = (n*(n+1)/2)^2

https://oeis.org/A061168 ∑floor(log2(i)) = ∑(bits.Len(i)-1)

∑∑|ai-aj|
= 2*∑(i*ai-preSum(i-1)), i=[0,n-1], a 需要排序
https://www.luogu.com.cn/blog/DPair2005/solution-cf340c
https://codeforces.com/problemset/problem/340/C

https://oeis.org/A005326 Number of permutations p of (1,2,3,...,n) such that k and p(k) are relatively prime for all k in (1,2,3,...,n)
https://oeis.org/A009679 Number of partitions of {1, ..., 2n} into coprime pairs

https://oeis.org/A333885 Number of triples (i,j,k) with 1 <= i < j < k <= n such that i divides j divides k https://ac.nowcoder.com/acm/contest/7613/A

https://oeis.org/A000295 Eulerian numbers: Sum_{k=0..n} (n-k)*2^k = 2^n - n - 1
	Number of permutations of {1,2,...,n} with exactly one descent
	Number of partitions of an n-set having exactly one block of size > 1
	a(n-1) is the number of subsets of {1..n} in which the largest element of the set exceeds by at least 2 the next largest element
		For example, for n = 5, a(4) = 11 and the 11 sets are {1,3}, {1,4}, {1,5}, {2,4}, {2,5}, {3,5}, {1,2,4}, {1,2,5}, {1,3,5}, {2,3,5}, {1,2,3,5}
	a(n-1) is also the number of subsets of {1..n} in which the second smallest element of the set exceeds by at least 2 the smallest element
		For example, for n = 5, a(4) = 11 and the 11 sets are {1,3}, {1,4}, {1,5}, {2,4}, {2,5}, {3,5}, {1,3,4}, {1,3,5}, {1,4,5}, {2,4,5}, {1,3,4,5}

https://oeis.org/A064413 EKG sequence (or ECG sequence)
a(1) = 1; a(2) = 2; for n > 2, a(n) = smallest number not already used which shares a factor with a(n-1)

https://oeis.org/A002326 least m > 0 such that 2n+1 divides 2^m-1
LC1806 https://leetcode-cn.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation/

https://oeis.org/A003136 Loeschian number: numbers of the form x^2 + xy + y^2
https://en.wikipedia.org/wiki/Loeschian_number
https://www.bilibili.com/video/BV1or4y1A76q

数的韧性 https://en.wikipedia.org/wiki/Persistence_of_a_number 乘法: https://oeis.org/A003001 加法: https://oeis.org/A006050

Smallest number h such that n*h is a repunit (111...1), or 0 if no such h exists
https://oeis.org/A190301 111...1
https://oeis.org/A216485 222...2
相关题目 https://atcoder.jp/contests/abc174/tasks/abc174_c  快速算法见 https://img.atcoder.jp/abc174/editorial.pdf

	Least k such that the decimal representation of k*n contains only 1's and 0's
	https://oeis.org/A079339
	0's and d's (2~9): A096681-A096688

	a(n) is the least value of k such that k*n uses only digits 1 and 2. a(n) = -1 if no such multiple exists
	https://oeis.org/A216482

	a(n) is the smallest positive number such that the decimal digits of n*a(n) are all 0, 1 or 2
	https://oeis.org/A181061

Berlekamp–Massey algorithm
https://en.wikipedia.org/wiki/Berlekamp%E2%80%93Massey_algorithm
https://oi-wiki.org/math/berlekamp-massey/

椭圆曲线加密算法 https://ac.nowcoder.com/acm/contest/6916/C

Gaussian integer https://en.wikipedia.org/wiki/Gaussian_integer
Eisenstein integer https://en.wikipedia.org/wiki/Eisenstein_integer
Eisenstein prime https://en.wikipedia.org/wiki/Eisenstein_prime

https://oeis.org/A054710 Number of powers of 10 mod n https://codeforces.com/problemset/problem/1070/A

https://oeis.org/A050295 Number of strongly triple-free subsets of {1, 2, ..., n}
    https://leetcode.cn/circle/discuss/QH0XWr/

https://oeis.org/A005245 The (Mahler-Popken) complexity of n: minimal number of 1's required to build n using + and *
	3 log_3 n <= a(n) <= 3 log_2 n

https://oeis.org/A001108 a(n)-th triangular number is a square: a(n+1) = 6*a(n) - a(n-1) + 2, with a(0) = 0, a(1) = 1
https://oeis.org/A001109 a(n)^2 is a triangular number: a(n) = 6*a(n-1) - a(n-2) with a(0)=0, a(1)=1
https://oeis.org/A001110 Square triangular numbers: numbers that are both triangular and square

https://oeis.org/A034836 Number of ways to write n as n = x*y*z with 1 <= x <= y <= z
https://oeis.org/A331072 A034836 前缀和 O(n^(2/3))
	https://atcoder.jp/contests/abc227/tasks/abc227_c

https://oeis.org/A244478 a(0)=2, a(1)=0, a(2)=2; thereafter a(n) = a(n-1-a(n-1))+a(n-2-a(n-2)) unless a(n-1) <= n-1 or a(n-2) <= n-2 in which case the sequence terminates
https://oeis.org/A244479
LC1140 https://leetcode.cn/problems/stone-game-ii/ 需要记忆化的 M 的上界

Collatz conjecture (3n+1) https://en.wikipedia.org/wiki/Collatz_conjecture
https://oeis.org/A006577 Number of halving and tripling steps to reach 1 in '3x+1' problem, or -1 if 1 is never reached
https://oeis.org/A008884 3x+1 sequence starting at 27
LC1387 https://leetcode.cn/problems/sort-integers-by-the-power-value/

挑战 2.6 节练习题
2429 分解 LCM/GCD = a*b 且 gcd(a,b)=1 且 a+b 最小
1930 https://www.luogu.com.cn/problem/UVA10555 https://www.luogu.com.cn/problem/SP1166 floatToRat
3126 https://www.luogu.com.cn/problem/UVA12101 https://www.luogu.com.cn/problem/SP1841 BFS
3421 质因数幂次和 可重排列
3292 https://www.luogu.com.cn/problem/UVA11105 在 Z={4k+1} 上筛素数
3641 https://www.luogu.com.cn/problem/UVA11287 Carmichael Numbers https://oeis.org/A002997 https://en.wikipedia.org/wiki/Carmichael_number
4.1 节练习题（模运算的世界）
1150 https://www.luogu.com.cn/problem/UVA10212
1284
2115
3708
2720
GCJ Japan 2011 Final B

CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=number+theory
CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=combinatorics

*/

func _(abs func(int64) int64, max func(int64, int64) int64) {
	const mod = 1_000_000_007 // 998244353
	pow := func(x, n, p int64) (res int64) {
		x %= p
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % p
			}
			x = x * x % p
		}
		return
	}

	/* GCD LCM 相关
	https://mathworld.wolfram.com/EuclideanAlgorithm.html
	https://en.wikipedia.org/wiki/Euclidean_algorithm
	https://stackoverflow.com/questions/3980416/time-complexity-of-euclids-algorithm

	https://oeis.org/A051010 Triangle T(m,n) giving of number of steps in the Euclidean algorithm for gcd(m,n) with 0<=m<n
	https://oeis.org/A034883 Maximum length of Euclidean algorithm starting with n and any nonnegative i<n
	https://oeis.org/A049826 GCD(n,i) 的迭代次数之和，O(nlogn)

	Tighter time complexity for GCD https://codeforces.com/blog/entry/63771
	Runtime of finding the GCD of an array https://codeforces.com/blog/entry/92720

	GCD 套路：枚举倍数（调和级数复杂度）
	GCD(x,x+y) = GCD(x,y) https://codeforces.com/problemset/problem/1110/C
	GCD 与质因子 https://codeforces.com/problemset/problem/264/B
	数组中最小的 LCM(ai,aj) https://codeforces.com/problemset/problem/1154/G
	分拆与 LCM  https://ac.nowcoder.com/acm/contest/5961/D https://ac.nowcoder.com/discuss/439005
	TIPS: 一般 LCM 的题目都需要用 LCM=x*y/GCD 转换成研究 GCD 的性质
	todo https://atcoder.jp/contests/abc162/tasks/abc162_e
	     https://atcoder.jp/contests/abc206/tasks/abc206_e

	todo 基于值域预处理的快速 GCD https://www.luogu.com.cn/problem/P5435

	GCD = 1 的子序列个数 https://codeforces.com/problemset/problem/803/F https://ac.nowcoder.com/acm/problem/112055
	见后面的 mu

	a 中任意两数互质 <=> 每个质数至多整除一个 a[i]
	https://codeforces.com/contest/1770/problem/C

	todo https://codeforces.com/contest/1462/problem/D 的 O(nlogn) 解法

	Frobenius problem / Coin problem / Chicken McNugget Theorem
	两种硬币面额为 a 和 b，互质，数量无限，所不能凑出的数值的最大值为 a*b-a-b
	https://artofproblemsolving.com/wiki/index.php/Chicken_McNugget_Theorem
	https://en.wikipedia.org/wiki/Coin_problem
	https://www.luogu.com.cn/problem/P3951
	https://codeforces.com/contest/1526/problem/B

	*/

	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	// 例题 https://nanti.jisuanke.com/t/A1633
	gcdPrefix := func(a []int64) []int64 {
		n := len(a)
		gp := make([]int64, n+1)
		for i, v := range a {
			gp[i+1] = gcd(gp[i], v)
		}
		return gp
	}
	gcdSuffix := func(a []int64) []int64 {
		n := len(a)
		gs := make([]int64, n+1)
		for i := n - 1; i >= 0; i-- {
			gs[i] = gcd(gs[i+1], a[i])
		}
		return gs
	}

	lcm := func(a, b int64) int64 { return a / gcd(a, b) * b }

	// 前 n 个数的 LCM https://oeis.org/A003418 a(n) = lcm(1,...,n) ~ exp(n)
	// 相关题目 https://atcoder.jp/contests/arc110/tasks/arc110_a
	//         https://codeforces.com/problemset/problem/1485/D
	//         https://codeforces.com/problemset/problem/1542/C
	//         https://codeforces.com/problemset/problem/1603/A
	// a(n)/a(n-1) = https://oeis.org/A014963
	//     前缀和 https://oeis.org/A072107 https://ac.nowcoder.com/acm/contest/7607/A
	// LCM(2, 4, 6, ..., 2n) https://oeis.org/A051426
	// Mangoldt Function https://mathworld.wolfram.com/MangoldtFunction.html
	// a(n) 的因子个数 d(lcm(1,...,n)) https://oeis.org/A056793
	//     这同时也是 1~n 的子集的 LCM 的种类数
	// 另一种通分：「排水系统」的另一种解法 https://zxshetzy.blog.luogu.org/ling-yi-zhong-tong-fen
	// https://oeis.org/A000793	Landau's function g(n): largest order of permutation of n elements
	//                          Equivalently, largest LCM of partitions of n
	lcms := []int64{
		0, 1, 2, 6, 12, 60, 60, 420, 840, 2520, 2520, // 10
		27720, 27720, 360360, 360360, 360360, 720720, 12252240, 12252240, 232792560, 232792560, // 20
		232792560, 232792560, // 22 (int32)
		5354228880, 5354228880, 26771144400, 26771144400, 80313433200, 80313433200, 2329089562800, 2329089562800, // 30
		72201776446800, 144403552893600, 144403552893600, 144403552893600, 144403552893600, 144403552893600, 5342931457063200, 5342931457063200, 5342931457063200, 5342931457063200, // 40
		219060189739591200, 219060189739591200, // 9419588158802421600,
	}

	// GCD 性质统计相关
	// NOTE: 对于一任意非负序列，前 i 个数的 GCD 是非增序列，且至多有 O(logMax) 个不同值
	//       应用：https://codeforces.com/problemset/problem/1210/C
	// #{(a,b) | 1<=a<=b<=n, gcd(a,b)=1}   https://oeis.org/A002088
	//     = ∑phi(i)
	// #{(a,b) | 1<=a,b<=n, gcd(a,b)=1}   https://oeis.org/A018805
	//     = 2*(∑phi(i))-1
	//     = 2*A002088(n)-1
	// #{(a,b) | 1<=a,b<=n, gcd(a,b) is prime}  todo https://www.luogu.com.cn/problem/P2568
	// #{(a,b,c) | 1<=a,b,c<=n, gcd(a,b,c)=1}   https://oeis.org/A071778
	//     = ∑mu(i)*floor(n/i)^3
	// #{(a,b,c,d) | 1<=a,b,c,d<=n, gcd(a,b,c,d)=1}   https://oeis.org/A082540
	//     = ∑mu(i)*floor(n/i)^4

	// GCD 求和相关
	// ∑gcd(n,i) = ∑{d|n}d*phi(n/d)          https://oeis.org/A018804 https://www.luogu.com.cn/problem/P2303
	//     更简化的公式见小粉兔博客 https://www.cnblogs.com/PinkRabbit/p/8278728.html
	// ∑n/gcd(n,i) = ∑{d|n}d*phi(d)          https://oeis.org/A057660
	// ∑∑gcd(i,j) = ∑phi(i)*(floor(n/i))^2   https://oeis.org/A018806   https://www.luogu.com.cn/problem/P2398
	// ∑∑gcd(i,j) j<=i   = (1/2)∑phi(i)*floor(n/i)*(floor(n/i)+1)    https://oeis.org/A272718
	// ∑∑gcd(i,j) j<i    = (A018806(n) - n*(n+1)/2) / 2    https://oeis.org/A178881
	//     https://www.luogu.com.cn/problem/P1390
	//     训练指南例题 2-9，UVa11426 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=26&page=show_problem&problem=2421
	// ∑∑∑gcd(i,j,k) = ∑phi(i)*(floor(n/i))^3   https://ac.nowcoder.com/acm/contest/7608/B

	// LCM 性质统计相关
	// https://oeis.org/A048691 #{(a,b) | lcm(a,b)=n}，等价于 #{(x,y) | x|n, y|n, gcd(x,y)=1}
	//     = d(n^2)
	//     = (2*e1+1)(2*e2+1)...(2*ek+1), 其中 ei 是 n 的质因子分解中第 i 个质数的幂次
	// https://oeis.org/A018892 #{(a,b) | a<=b, lcm(a,b)=n}，等价于 #{(x,y) | x|n, y|n, x<=y, gcd(x,y)=1}
	//     = (d(n^2)+1)/2
	//     = ((2*e1+1)(2*e2+1)...(2*ek+1) + 1) / 2, 其中 ei 是 n 的质因子分解中第 i 个质数的幂次
	//     Number of ways to write 1/n as a sum of exactly 2 unit fractions
	//     Number of divisors of n^2 less than or equal to n
	//     训练指南 2.10 习题，UVa10892 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1833
	// https://oeis.org/A182082 A018892 的前缀和
	//     https://projecteuler.net/problem=379
	//     https://zhenweiliu.gitee.io/blog/2019/08/05/Project-Euler-Problem-379-Least-common-multiple-count/

	// LCM 求和相关
	// ∑lcm(n,i) = n*(1+∑{d|n}d*phi(d))/2 = n*(1+A057660(n))/2   https://oeis.org/A051193
	// ∑lcm(n,i)/n = A051193(n)/n = (1+∑{d|n}d*phi(d))/2 = (1+A057660(n))/2   https://oeis.org/A057661
	// ∑∑lcm(i,j)   https://oeis.org/A064951

	// 统计数组的所有子区间的 GCD 的不同个数
	// 代码和题目见 bits.go 中的 bitOpTrick

	// 统计数组的所有子序列的 GCD 的不同个数，复杂度 O(Clog^2C)
	// LC1819 https://leetcode-cn.com/problems/number-of-different-subsequences-gcds/
	// 我的题解 https://leetcode.cn/problems/number-of-different-subsequences-gcds/solution/ji-bai-100mei-ju-gcdxun-huan-you-hua-pyt-get7/
	countDifferentSubsequenceGCDs := func(a []int, gcd func(int, int) int) (ans int) {
		const mx int = 4e5 //
		has := [mx + 1]bool{}
		for _, v := range a {
			has[v] = true
		}
		for i := 1; i <= mx; i++ {
			g := 0
			for j := i; j <= mx && g != i; j += i { // 枚举 i 的倍数 j
				if has[j] { // 如果 j 在 nums 中
					g = gcd(g, j) // 更新最大公约数
				}
			}
			if g == i { // 找到一个答案
				ans++
			}
		}
		return
	}

	// 最简分数
	// https://codeforces.com/problemset/problem/1468/F
	type frac struct{ num, den int64 }

	// 如果有负数需要对 g 取绝对值
	makeFrac := func(a, b int64) frac { g := gcd(a, b); return frac{a / g, b / g} }

	// 比较两个（最简化后的）frac
	// 不使用高精度、浮点数等
	// 核心思路是将 a b 写成连分数形式，逐个比较
	// 复杂度 O(log)
	lessFrac := func(a, b frac) bool {
		// 如果保证 a b 均为正数，for 前面的这些 if 可以去掉
		if a == b {
			return false
		}
		if a.num == 0 {
			return b.num > 0
		}
		if b.num == 0 {
			return a.num < 0
		}
		if a.num > 0 != (b.num > 0) {
			return a.num < b.num
		}
		if a.num < 0 { // b.num < 0
			a, b = frac{-b.num, b.den}, frac{-a.num, a.den}
		}
		for {
			if a.den == 0 {
				return false
			}
			if b.den == 0 {
				return true
			}
			da, db := a.num/a.den, b.num/b.den
			if da != db {
				return da < db
			}
			a, b = frac{b.den, b.num - db*b.den}, frac{a.den, a.num - da*a.den}
		}
	}

	// 类欧几里得算法
	// ∑⌊(ai+b)/m⌋, i in [0,n-1]
	// https://oi-wiki.org/math/euclidean/
	// todo https://www.luogu.com.cn/blog/AlanWalkerWilson/Akin-Euclidean-algorithm-Basis
	//      https://www.luogu.com.cn/blog/Shuchong/qian-tan-lei-ou-ji-li-dei-suan-fa
	//      万能欧几里得算法 https://www.luogu.com.cn/blog/ILikeDuck/mo-neng-ou-ji-li-dei-suan-fa
	//
	// 模板题 https://atcoder.jp/contests/practice2/tasks/practice2_c
	//       https://www.luogu.com.cn/problem/P5170
	//       https://loj.ac/p/138
	// todo https://codeforces.com/problemset/problem/1182/F
	//  https://codeforces.com/problemset/problem/1098/E
	floorSum := func(n, m, a, b int64) (res int64) {
		if a < 0 {
			a2 := a%m + m
			res -= n * (n - 1) / 2 * ((a2 - a) / m)
			a = a2
		}
		if b < 0 {
			b2 := b%m + m
			res -= n * ((b2 - b) / m)
			b = b2
		}
		for {
			if a >= m {
				res += n * (n - 1) / 2 * (a / m)
				a %= m
			}
			if b >= m {
				res += n * (b / m)
				b %= m
			}
			yMax := a*n + b
			if yMax < m {
				break
			}
			n = yMax / m
			b = yMax % m
			m, a = a, m
		}
		return
	}

	sqCheck := func(a int64) bool { r := int64(math.Round(math.Sqrt(float64(a)))); return r*r == a }
	cubeCheck := func(a int64) bool { r := int64(math.Round(math.Cbrt(float64(a)))); return r*r*r == a }
	// 平方数开平方
	sqrt := func(a int64) int64 {
		r := int64(math.Round(math.Sqrt(float64(a))))
		if r*r == a {
			return r
		}
		return -1
	}
	// 立方数开立方
	cbrt := func(a int64) int64 {
		r := int64(math.Round(math.Cbrt(float64(a))))
		if r*r*r == a {
			return r
		}
		return -1
	}

	// 返回差分表的最后一个数
	// return the bottom entry in the difference table
	// 另一种做法是用公式 ∑(-1)^i * C(n,i) * a_i, i=0..n-1
	bottomDiff := func(a []int) int {
		for ; len(a) > 1; a = a[:len(a)-1] {
			for i := 0; i+1 < len(a); i++ {
				a[i] = a[i+1] - a[i]
			}
		}
		return a[0]
	}

	/* 质数 质因子分解 */

	// n/2^k https://oeis.org/A000265
	// A000265 的前缀和 https://oeis.org/A135013
	// a(n) = Sum_{k>=1} (round(n/2^k))^2

	// 质数表 https://oeis.org/A000040
	// primes[i]%10 https://oeis.org/A007652
	// 10-primes[i]%10 https://oeis.org/A072003
	// p-1 https://oeis.org/A006093
	// p+1 https://oeis.org/A008864
	// p^2+p+1 https://oeis.org/A060800 = sigma(p^2)
	// prime index prime https://oeis.org/A006450
	primes := []int{ // 预处理 mask 的见下
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199,
		211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293,
		307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397,
		401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499,
		503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599,
		601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691,
		701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797,
		809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887,
		907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997, /* #=168 */
		1009, 1013, 1019, 1021, 1031, 1033, 1039, 1049, 1051, 1061, 1063, 1069, 1087, 1091, 1093, 1097,
		1103, 1109, 1117, 1123, 1129, 1151, 1153, 1163, 1171, 1181, 1187, 1193,
		1201, 1213, 1217, 1223, 1229, 1231, 1237, 1249, 1259, 1277, 1279, 1283, 1289, 1291, 1297,
		1301, 1303, 1307, 1319, 1321, 1327, 1361, 1367, 1373, 1381, 1399,
		1409, 1423, 1427, 1429, 1433, 1439, 1447, 1451, 1453, 1459, 1471, 1481, 1483, 1487, 1489, 1493, 1499,
		1511, 1523, 1531, 1543, 1549, 1553, 1559, 1567, 1571, 1579, 1583, 1597,
		1601, 1607, 1609, 1613, 1619, 1621, 1627, 1637, 1657, 1663, 1667, 1669, 1693, 1697, 1699,
		1709, 1721, 1723, 1733, 1741, 1747, 1753, 1759, 1777, 1783, 1787, 1789,
		1801, 1811, 1823, 1831, 1847, 1861, 1867, 1871, 1873, 1877, 1879, 1889,
		1901, 1907, 1913, 1931, 1933, 1949, 1951, 1973, 1979, 1987, 1993, 1997, 1999, /* #=303 */
	}

	{
		// 小范围质数状压
		// Squarefree numbers https://oeis.org/A005117
		const mx = 30
		primeMask := [mx + 1]int{}
		for i := 2; i <= mx; i++ {
			for j, p := range primes {
				if i%p == 0 {
					//if i%(p*p) == 0 { // 有平方因子
					//	primeMask[i] = -1
					//	break
					//}
					primeMask[i] |= 1 << j // 把 j 加到集合中
				}
			}
		}
	}

	// 第 10^k 个素数
	// https://oeis.org/A006988
	// 补充：第 1e5, 2e5, 3e5, ..., 1e6 个素数
	// 1299709, 2750159, 4256233, 5800079, 7368787, 8960453, 10570841, 12195257, 13834103, 15485863
	primes10k := []int64{2, 29, 541, 7919, 104729, 1299709, 15485863, /* 1e6 */ 179424673, 2038074743, 22801763489, 252097800623, 2760727302517, 29996224275833, 323780508946331, 3475385758524527, 37124508045065437, 394906913903735329, 4185296581467695669}

	// map{小于 10^n 的素数个数: 小于 10^n 的最大素数} https://oeis.org/A006880 https://oeis.org/A003618   10^n-a(n): https://oeis.org/A033874
	primes10 := map[int]int64{
		4:         7,
		25:        97,
		168:       997, // 1e3
		1229:      9973,
		9592:      99991,
		78498:     999983, // 1e6
		664579:    9999991,
		5761455:   99999989,
		50847534:  999999937, // 1e9
		455052511: 9999999967,
	}

	// 大于 10^n 的最小素数 https://oeis.org/A090226 https://oeis.org/A003617    a(n)-10^n: https://oeis.org/A033873
	primes10_ := []int64{
		2,
		11,
		101,
		1009, // 1e3
		10007,
		100003,
		1000003, // 1e6
		10000019,
		100000007,
		1000000007, //1e9
		10000000019,
	}

	/* 质数性质统计相关

	Counting primes
	https://en.wikipedia.org/wiki/Meissel%E2%80%93Lehmer_algorithm
	https://oi-wiki.org/math/meissel-lehmer/
	https://www.zhihu.com/question/29580448
	O(n^(2/3)log^(1/3)(n)) https://codeforces.com/blog/entry/91632

	质数的幂次组成的集合 {p^k} https://oeis.org/A000961
	补集 https://oeis.org/A024619
	Exponential of Mangoldt function https://oeis.org/A014963

	质数前缀和 https://oeis.org/A007504
	a(n) ~ n^2 * log(n) / 2
	a(n)^2 - a(n-1)^2 = A034960(n)
	EXTRA: divide odd numbers into groups with prime(n) elements and add together https://oeis.org/A034960
		仍然是质数的前缀和 https://oeis.org/A013918 对应的前缀和下标 https://oeis.org/A013916

	质数前缀积 prime(n)# https://oeis.org/A002110
	the least number with n distinct prime factors
	2, 6, 30, 210, 2310, 30030, 510510, 9699690, 223092870, /9/
	6469693230, 200560490130, 7420738134810, 304250263527210, 13082761331670030, 614889782588491410

	质数间隙 https://en.wikipedia.org/wiki/Prime_gap https://oeis.org/A001223
	Positions of records https://oeis.org/A002386 https://oeis.org/A005669
	Values of records https://oeis.org/A005250
	Gap 均值 https://oeis.org/A286888 a(n)= floor((prime(n) - 2)/(n - 1))
	相关题目 https://www.luogu.com.cn/problem/P6104 https://class.luogu.com.cn/classroom/lgr69
	Kick Start 2021 Round B Consecutive Primes https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b/000000000077a8e6
	Numbers whose distance to the closest prime number is a prime number https://oeis.org/A160666

	孪生素数 https://en.wikipedia.org/wiki/Twin_prime https://oeis.org/A001359 https://oeis.org/A006512 https://oeis.org/A077800
	https://oeis.org/A113274 Record gaps between twin primes
		Upper bound: gaps between twin primes are smaller than 0.76*(log p)^3, where p is the prime at the end of the gap.
	https://oeis.org/A113275 Lesser of twin primes for which the gap before the following twin primes is a record

	Prime k-tuple https://en.wikipedia.org/wiki/Prime_k-tuple
	Prime constellations / diameter https://en.wikipedia.org/wiki/Prime_k-tuple#Prime_constellations https://oeis.org/A008407
	Cousin prime https://en.wikipedia.org/wiki/Cousin_prime https://oeis.org/A023200
	Sexy prime https://en.wikipedia.org/wiki/Sexy_prime https://oeis.org/A023201
	Prime triplet https://en.wikipedia.org/wiki/Prime_triplet https://oeis.org/A098420
	Primes in arithmetic progression https://en.wikipedia.org/wiki/Primes_in_arithmetic_progression

	First Hardy–Littlewood conjecture https://en.wikipedia.org/wiki/First_Hardy%E2%80%93Littlewood_conjecture
	Second Hardy–Littlewood conjecture https://en.wikipedia.org/wiki/Second_Hardy%E2%80%93Littlewood_conjecture 哈代-李特尔伍德第二猜想

	https://oeis.org/A007918 Least prime >= n (version 1 of the "next prime" function)
	https://oeis.org/A007920 Smallest number k such that n + k is prime

	任意质数之差 https://oeis.org/A030173
	非任意质数之差 https://oeis.org/A007921

	质数的逆二项变换 Inverse binomial transform of primes https://oeis.org/A007442

	合数前缀和 https://oeis.org/A053767

	合数前缀积 Compositorial number https://oeis.org/A036691

	不与质数相邻的合数 https://oeis.org/A079364

	半素数 https://oeis.org/A001358 也叫双素数/二次殆素数 Semiprimes (or biprimes): products of two primes
	https://en.wikipedia.org/wiki/Semiprime
	https://en.wikipedia.org/wiki/Almost_prime
	非平方半素数 https://oeis.org/A006881 Squarefree semiprimes: Numbers that are the product of two distinct primes.

	绝对素数 https://oeis.org/A003459 各位数字可以任意交换位置，其结果仍为素数
	https://en.wikipedia.org/wiki/Permutable_prime

	哥德巴赫猜想：大于 2 的偶数，都可表示成两个素数之和。
	偶数分拆的最小质数 Goldbach’s conjecture https://oeis.org/A020481
	Conjecture: a(n) ~ O(√n)
	https://en.wikipedia.org/wiki/Goldbach%27s_conjecture
		Positions of records https://oeis.org/A025018
		Values of records https://oeis.org/A025019
		1e9 内最大的为 a(721013438) = 1789
		2e9 内最大的为 a(1847133842) = 1861
	https://codeforces.com/problemset/problem/735/D
	将 1~n 这 n 个数分成若干组，使每组数之和为质数 https://codeforces.com/problemset/problem/45/G
		这题需要用到 a(n) ~ O(√n)

	勒让德猜想 - 在两个相邻平方数之间，至少有一个质数 Legendre’s conjecture
	https://en.wikipedia.org/wiki/Legendre%27s_conjecture
	Number of primes between n^2 and (n+1)^2 https://oeis.org/A014085
	Number of primes between n^3 and (n+1)^3 https://oeis.org/A060199

	伯特兰-切比雪夫定理 - n ~ 2n 之间至少有一个质数 Bertrand's postulate
	https://en.wikipedia.org/wiki/Bertrand%27s_postulate
	Number of primes between n and 2n (inclusive) https://oeis.org/A035250
	Number of primes between n and 2n exclusive https://oeis.org/A060715
	n ~ 1.5n https://codeforces.com/contest/1178/problem/D

	Least k such that H(k) > n, where H(k) is the harmonic number ∑{i=1..k} 1/i
	https://oeis.org/A002387
	https://oeis.org/A004080

		a(n) = smallest prime p such that ∑{primes q = 2, ..., p} 1/q exceeds n
		5, 277, 5_195_977, 1801241230056600523
		https://oeis.org/A016088 pi
		https://oeis.org/A046024 i

	a(n) = largest m such that the harmonic number H(m)= ∑{i=1..m} 1/i is < n
	https://oeis.org/A115515

		a(n) = largest prime p such that ∑{primes q = 2, ..., p} 1/q does not exceed n
		3, 271, 5_195_969, 1801241230056600467
		https://oeis.org/A223037

	Exponent of highest power of 2 dividing n, a.k.a. the binary carry sequence, the ruler sequence, or the 2-adic valuation of n
	a(n) = 0 if n is odd, otherwise 1 + a(n/2)
	https://oeis.org/A007814

	https://oeis.org/A000043 Mersenne exponents: primes p such that 2^p - 1 is prime. Then 2^p - 1 is called a Mersenne prime

	*/

	// 判断一个数是否为质数
	isPrime := func(n int64) bool {
		for i := int64(2); i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return n >= 2
	}
	// https://www.luogu.com.cn/problem/U82118
	isPrime = func(n int64) bool { return big.NewInt(n).ProbablyPrime(0) }

	// 判断质数+求最大质因子
	// 先用 Pollard-Rho 算法求出一个因子，然后递归求最大质因子
	// https://zhuanlan.zhihu.com/p/267884783
	// https://www.luogu.com.cn/problem/P4718
	pollardRho := func(n int64) int64 {
		if n == 4 {
			return 2
		}
		if isPrime(n) {
			return n
		}
		mul := func(a, b int64) (res int64) {
			for ; b > 0; b >>= 1 {
				if b&1 == 1 {
					res = (res + a) % n
				}
				a = (a + a) % n
			}
			return
		}
		for {
			c := 1 + rand.Int63n(n-1)
			f := func(x int64) int64 { return (mul(x, x) + c) % n }
			for t, r := f(0), f(f(0)); t != r; t, r = f(t), f(f(r)) {
				if d := gcd(abs(t-r), n); d > 1 {
					return d
				}
			}
		}
	}
	{
		cacheGPF := map[int64]int64{}
		var gpf func(int64) int64
		gpf = func(x int64) (res int64) {
			if cacheGPF[x] > 0 {
				return cacheGPF[x]
			}
			defer func() { cacheGPF[x] = res }()
			p := pollardRho(x)
			if p == x {
				return p
			}
			return max(gpf(p), gpf(x/p))
		}
	}

	// 预处理: [2,mx] 范围内的质数
	// 埃筛 埃氏筛 埃拉托斯特尼筛法 Sieve of Eratosthenes
	// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	// https://oeis.org/A055399 Number of stages of sieve of Eratosthenes needed to identify n as prime or composite
	// https://oeis.org/A230773 Minimum number of steps in an alternate definition of the Sieve of Eratosthenes needed to identify n as prime or composite
	// 质数个数 π(n) https://oeis.org/A000720
	//         π(10^n) https://oeis.org/A006880
	//         4, 25, 168, 1229, 9592, 78498, 664579, 5761455, 50847534, /* 1e9 */
	//         455052511, 4118054813, 37607912018, 346065536839, 3204941750802, 29844570422669, 279238341033925, 2623557157654233, 24739954287740860, 234057667276344607,
	// 思想应用 https://codeforces.com/contest/1646/problem/E
	sieve := func() {
		const mx int = 1e6
		primes := []int{}
		pid := [mx + 1]int{-1, -1}
		for i := 2; i <= mx; i++ {
			if pid[i] == 0 {
				primes = append(primes, i)
				pid[i] = len(primes)
				for j := i * i; j <= mx; j += i {
					pid[j] = -1
				}
			}
		}

		// EXTRA: pi(n), the number of primes <= n https://oeis.org/A000720
		pi := [mx + 1]int{}
		for i := 2; i <= mx; i++ {
			pi[i] = pi[i-1]
			if pid[i] > 0 {
				pi[i]++
			}
		}
	}

	// 线筛 线性筛 欧拉筛
	// 每个合数都从其 LPF 标记到（在遍历到 i = 合数/LPF 的时候，标记这些合数）
	// 参考 https://oi-wiki.org/math/sieve/ 以及进阶指南 p.136-137
	// mx = 3e7 时比埃氏筛大约快 100ms https://codeforces.com/problemset/submission/986/206447142
	//                              https://codeforces.com/problemset/submission/986/206445786
	// https://www.luogu.com.cn/problem/P3383
	sieveEuler := func() {
		const mx int = 1e7
		primes := []int{}
		pid := [mx + 1]int{-1, -1}
		for i := 2; i <= mx; i++ {
			if pid[i] == 0 {
				pid[i] = len(primes) + 1
				primes = append(primes, i)
			}
			for _, p := range primes {
				if p*i > mx {
					break
				}
				pid[p*i] = -1
				if i%p == 0 { // 后面的「质数*i」标记出的合数，其 LPF 不是该质数，应及时退出，从而避免重复标记
					break
				}
			}
		}
	}

	// 一般线性筛的模板
	// 记 f(n) 为积性函数
	// 其满足 1. f(p) = p ...
	//       2. f(p^(k+1)) = f(p^k) ... p
	//       3. f(x*p) = f(x) ... p (p 不是 x 的因子)
	// 一个典型的例子见下面 σ(n) 的线性筛求法
	// https://codeforces.com/contest/1512/problem/G
	// https://codeforces.com/gym/103107/problem/F
	sieveEulerTemplate := func() []int {
		const mx int = 1e7
		f := make([]int, mx+1)
		f[1] = 1 //
		vis := make([]bool, mx+1)
		primes := []int{}
		for i := 2; i <= mx; i++ {
			if !vis[i] {
				// 1: p
				f[i] = i
				primes = append(primes, i)
			}
			for _, p := range primes {
				v := p * i
				if v > mx {
					break
				}
				vis[v] = true
				if i%p == 0 {
					// 2: p^(k+1) <- p^k
					f[v] = f[i] * p
					break
				}
				// 3: x*p <- x 且 x 的质因子是没有 p 的
				f[v] = f[i] * p
			}
		}
		return f
	}

	// 区间筛法
	// 预处理 [2,√R] 的所有质数，去筛 [L,R] 之间的质数

	// todo 多组数据下的记忆化质因数分解 https://codeforces.com/contest/1512/submission/112590495

	// 质因数分解（完整版）prime factorization
	// 返回分解出的质数及其指数
	// 预处理 [2,√MX] 的素数可以加速这一过程
	// https://mathworld.wolfram.com/PrimeFactorization.html
	// todo 更高效的算法 - Pollard's Rho
	// n 的质因数分解中 2 的幂次 https://oeis.org/A007814
	// n 的质因数分解中非 2 的幂次之和 https://oeis.org/A087436
	type factor struct {
		p  int64
		e  int
		pe int64 // p^e
	}
	factorize := func(x int64) (factors []factor) {
		for i := int64(2); i*i <= x; i++ {
			if x%i > 0 {
				continue
			}
			e := 1
			pe := i
			for x /= i; x%i == 0; x /= i {
				e++
				pe *= i
			}
			factors = append(factors, factor{i, e, pe})
		}
		if x > 1 {
			factors = append(factors, factor{x, 1, x})
		}
		return
	}

	// 质因数分解（质数及其幂次）prime factorization
	// LC2507 https://leetcode.cn/problems/smallest-value-after-replacing-with-sum-of-prime-factors/
	// LC2584 https://leetcode.cn/problems/split-the-array-to-make-coprime-products/
	primeDivisors := func(x int64) (primes []int64) {
		for i := int64(2); i*i <= x; i++ {
			if x%i > 0 {
				continue
			}
			//e := 1
			for x /= i; x%i == 0; x /= i {
				//e++
			}
			primes = append(primes, i)
		}
		if x > 1 {
			//e := 1
			primes = append(primes, x)
		}
		return
	}

	// 质因数分解（加速：跳过偶数）prime factorization
	// 在 1e15 下比上面快大概 150ms
	// https://codeforces.com/contest/1334/submission/143919621
	// https://codeforces.com/contest/1334/submission/143919683
	primeDivisors2 := func(x int64) (primes []int64) {
		if x&1 == 0 {
			primes = append(primes, 2)
			x /= x & -x // 去掉所有的因子 2
		}
		for i := int64(3); i*i <= x; i += 2 {
			if x%i > 0 {
				continue
			}
			for x /= i; x%i == 0; x /= i {
			}
			primes = append(primes, i)
		}
		if x > 1 {
			primes = append(primes, x)
		}
		return
	}

	// 阶乘的质因数分解中 p 的幂次
	// https://cp-algorithms.com/algebra/factorial-divisors.html
	// https://codeforces.com/problemset/problem/633/B
	// https://codeforces.com/problemset/problem/1114/C
	// https://oeis.org/A027868 p=5 时为 n! 尾零的个数
	// https://oeis.org/A191610 Possible number of trailing zeros in n!
	// https://oeis.org/A000966 n! never ends in this many 0's
	//    The simplest way to obtain this sequence is by constructing a power series
	//       A(x) = Sum_{k >= 1} x^a(k) whose exponents give the terms of the sequence.
	//    Define e(n) = (5^n-1)/4, f(n) = (1-x^(e(n)-1))/(1-x^e(n-1)), t(n) = x^(e(n)-6).
	//    相关题目 LC793 https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function/
	//       数学解法 https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function/solution/shu-xue-tui-dao-by-jriver/
	powerOfFactorialPrimeDivisor := func(n, p int64) (k int64) {
		for n > 0 {
			n /= p
			k += n
		}
		return
	}

	// 预处理: [2,mx] 的质因数分解的系数和 bigomega(n) or Omega(n) https://oeis.org/A001222
	// https://en.wikipedia.org/wiki/Prime_omega_function
	// a(n) depends only on prime signature of n (cf. https://oeis.org/A025487)
	// So a(24) = a(375) since 24 = 2^3 * 3 and 375 = 3 * 5^3 both have prime signature (3, 1)
	//
	// 		Omega(n) - omega(n) https://oeis.org/A046660
	//
	// 另一种写法 https://math.stackexchange.com/questions/1955105/corectness-of-prime-factorization-over-a-range
	// 性质：Omega(nm)=Omega(n)+Omega(m)
	// 前缀和 https://oeis.org/A022559 = Omega(n!) ~ O(nloglogn)
	// EXTRA: https://oeis.org/A005361 Product of exponents of prime factorization of n
	//        https://oeis.org/A135291 Product of exponents of prime factorization of n!
	primeExponentsCountAll := func() {
		const mx int = 1e6
		Omega := [mx + 1]int{} // int8
		primes := []int{}
		for i := 2; i <= mx; i++ {
			if Omega[i] == 0 {
				Omega[i] = 1
				primes = append(primes, i)
			}
			for _, p := range primes {
				if p*i > mx {
					break
				}
				Omega[p*i] = Omega[i] + 1
			}
		}

		// EXTRA: 前缀和，即 Omega(n!) https://oeis.org/A022559
		for i := 3; i <= mx; i++ {
			Omega[i] += Omega[i-1]
		}
	}

	// 单个数的 Omega
	// https://codeforces.com/contest/1538/problem/D
	primeExponentsCount := func(x int) (c int) {
		for i := 2; i*i <= x; i++ {
			for ; x%i == 0; x /= i {
				c++
			}
		}
		if x > 1 {
			c++
		}
		return
	}

	/* 因子/因数/约数

	n 的因子个数 d(n) = Π(ei+1), ei 为第 i 个质数的系数 https://oeis.org/A000005 d(n) 也写作 τ(n) tau(n)
		Positions of records (高合成数，反素数) https://oeis.org/A002182
		Values of records https://oeis.org/A002183
		相关题目：范围内的最多约数个数 https://www.luogu.com.cn/problem/P1221 加强版 https://ac.nowcoder.com/acm/contest/82/A

		max(d(i)), i=1..10^n https://oeis.org/A066150
			方便估计复杂度 - 近似为开立方
			4, 12, 32, 64, 128, /5位数/
	        240, 448, 768, 1344, /9位数/
			2304, 4032, 6720, 10752, 17280, 26880, 41472, 64512, 103680, 161280, /19位数/

			上面这些数对应的最小的 n https://oeis.org/A066151
			6, 60, 840, 7560, 83160,
			720720, 8648640, 73513440, 735134400,
			6983776800, 97772875200, 963761198400, 9316358251200, 97821761637600, 866421317361600, 8086598962041600, 74801040398884800, 897612484786617600

		d(n) 前缀和 = ∑{k=1..n} floor(n/k) https://oeis.org/A006218
	               = 见后文「数论分块/除法分块」

		n+d(n) https://oeis.org/A062249
		n-d(n) https://oeis.org/A049820   count https://oeis.org/A060990   前缀和 https://oeis.org/A161664
		n*d(n) https://oeis.org/A038040   前缀和 https://oeis.org/A143127 = Sum_{i=1..floor(√n)}i*(i+floor(n/i))*(floor(n/i)+1-i) - 平方和(floor(√n))
												https://atcoder.jp/contests/abc172/tasks/abc172_d
		n*n*d(n) https://oeis.org/A034714   前缀和 https://oeis.org/A319085
		d(n)|n https://oeis.org/A033950 refactorable numbers / tau numbers
			n/d(n) https://oeis.org/A036762
		n%d(n) https://oeis.org/A054008
		a(1)=1, a(n+1)=a(n)+d(a(n)) https://oeis.org/A064491
		Smallest k such that d(k) = n https://oeis.org/A005179
			a(p) = 2^(p-1) for primes p
			相关题目 https://codeforces.com/problemset/problem/27/E https://www.luogu.com.cn/problem/P1128
			质数的情况 https://oeis.org/A061286
	    Number of divisors of n^2 less than n https://oeis.org/A063647 Also number of ways to write 1/n as a difference of exactly 2 unit fractions
	        a(n) = (d(n^2)-1)/2

	n 的因子之和 σ(n) = Π(pi^(ei+1)-1)/(pi-1) https://oeis.org/A000203 todo 相关题目 https://www.luogu.com.cn/problem/P1593
	    线性筛求法见后面
	    max(σ(i)), i=1..10^n https://oeis.org/A066410
	         对应的 n https://oeis.org/A066424
	    Smallest k such that sigma(k) = n https://oeis.org/A051444
		σ(n) 前缀和 = ∑{k=1..n} k*floor(n/k) https://oeis.org/A024916
		https://oeis.org/A001157 sigma_2(n): sum of squares of divisors of n
		https://oeis.org/A064602 sigma_2 前缀和 = Sum_{i=1..n} i^2 * floor(n/i)
		真因子之和 https://oeis.org/A001065 σ(n)-n
		完全数/完美数/完备数 https://oeis.org/A000396 Perfect numbers (σ(n) = 2n)
			https://en.wikipedia.org/wiki/Perfect_number
			https://en.wikipedia.org/wiki/Euclid%E2%80%93Euler_theorem
			LC507 https://leetcode-cn.com/problems/perfect-number/
		过剩数/丰数/盈数 https://oeis.org/A005101 Abundant numbers (σ(n) > 2n)
			https://en.wikipedia.org/wiki/Abundant_number
		亏数/缺数/不足数 https://oeis.org/A005100 Deficient numbers (σ(n) < 2n)
			https://en.wikipedia.org/wiki/Deficient_number
			https://ac.nowcoder.com/acm/contest/10322/A O(nlogn) 可以先预处理因子

	n 的因子之积 μ(n) = n^(d(n)/2.0) https://oeis.org/A007955
	because we can form d(n)/2 pairs from the factors, each with product n
		若 n 是完全平方数，则 ei+1 全为奇数，此时可以计算 [n^(1/2)]^d(n)
		否则，ei+1 中必有偶数，将其除二
		相关题目 https://codeforces.com/problemset/problem/615/D

	n 的因子的差分表的最后一个数 https://oeis.org/A187202 https://oeis.org/A187203
	NOTE: a(2^k) = 1

		正数 https://oeis.org/A193671
		零   https://oeis.org/A187204
		负数 https://oeis.org/A193672

	d(d(...d(n))) 迭代至 2 所需要的迭代次数
	0,0,1,0,2,0,2,1,2,0,3,0,2,2,1,0,3,0,3,2,2,0,3,1,2,2,3

	高合成数/反质数 Highly Composite Numbers https://oeis.org/A002182
	https://oi-wiki.org/math/prime/#_7
	性质：一个高合成数一定是由另一个高合成数乘一个质数得到
	见进阶指南 p.140-141
	Number of divisors of n-th highly composite number https://oeis.org/A002183
	Number of highly composite numbers not divisible by n https://oeis.org/A199337
	求出不超过 n 的最大的反质数 https://www.luogu.com.cn/problem/P1463

	Largest divisor of n having the form 2^i*5^j https://oeis.org/A132741
	a(n) = A006519(n)*A060904(n) = 2^A007814(n)*5^A112765(n)

	Squarefree numbers https://oeis.org/A005117 (介绍了一种筛法)
	Numbers that are not divisible by a square greater than 1
	https://en.wikipedia.org/wiki/Square-free_integer
	Lim_{n->infinity} a(n)/n = Pi^2/6，即密度为 6/(Pi^2) ≈ 0.608

		Numbers that are not squarefree https://oeis.org/A013929
		Numbers that are divisible by a square greater than 1

	https://oeis.org/A000188 square root of largest square dividing n
	https://oeis.org/A120486 Partial sums of A000188  a(n) = Sum_{k=1..n} phi(k)*floor(n/k^2)

	a(n) = Min {m>n | m has same prime factors as n ignoring multiplicity} https://oeis.org/A065642
		Numbers such that a(n)/n is not an integer are listed in https://oeis.org/A284342
	*/

	// 枚举一个数的全部因子
	divisors := func(n int64) (ds []int64) {
		for d := int64(1); d*d <= n; d++ {
			if n%d == 0 {
				ds = append(ds, d)
				if d*d < n {
					ds = append(ds, n/d)
				}
			}
		}
		//sort.Slice(ds, func(i, j int) bool { return ds[i] < ds[j] })
		return
	}

	// 不需要排序的写法
	divisors = func(n int64) (ds []int64) {
		ds2 := []int64{}
		for d := int64(1); d*d <= n; d++ {
			if n%d == 0 {
				ds = append(ds, d)
				if d*d < n {
					ds2 = append(ds2, n/d)
				}
			}
		}
		for i := len(ds2) - 1; i >= 0; i-- {
			ds = append(ds, ds2[i])
		}
		return
	}

	// 无需额外空间的写法
	// https://leetcode.cn/problems/smallest-integer-divisible-by-k/solution/san-chong-suan-fa-you-hua-pythonjavacgo-tk4cj/
	divisorsO1Space := func(n int64) {
		// 从小到大枚举不超过 sqrt(n) 的因子
		i := int64(1)
		for ; i*i <= n; i++ {
			if n%i == 0 {
				// do i ...
			}
		}
		// 从小到大枚举大于 sqrt(n) 的因子
		i--
		if i*i == n {
			i-- // 避免重复统计
		}
		for ; i > 0; i-- {
			if n%i == 0 {
				// do m/i ...
			}
		}
	}

	// Number of odd divisors of n https://oeis.org/A001227
	// a(n) = d(2*n) - d(n)
	// 亦为整数 n 分拆成若干连续整数的方法数
	// Number of partitions of n into consecutive positive integers including the trivial partition of length 1
	// e.g. 9 = 2+3+4 or 4+5 or 9 so a(9)=3
	// 相关题目 LC829 https://leetcode.cn/problems/consecutive-numbers-sum/
	// Kick Start 2021 Round C Alien Generator https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44/00000000007ec1cb
	oddDivisorsNum := func(n int) (ans int) {
		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				if i&1 == 1 {
					ans++
				}
				if i*i < n && n/i&1 == 1 {
					ans++
				}
			}
		}
		return
	}

	// 因子的中位数（偶数个因子时取小的那个）
	// Lower central (median) divisor of n https://oeis.org/A060775
	// EXTRA: Largest divisor of n <= sqrt(n) https://oeis.org/A033676
	maxSqrtDivisor := func(n int) int {
		for d := int(math.Sqrt(float64(n))); ; d-- {
			if n%d == 0 {
				return d
			}
		}
	}

	// 预处理: [1,mx] 范围内数的所有因子
	// 复杂度 O(nlogn)
	// NOTE: 1~n 的因子个数总和大约为 nlogn
	// NOTE: divisors[x] 为奇数 => x 为完全平方数 https://oeis.org/A000290
	// NOTE: halfDivisors(x) 为 ≤√x 的因数集合 https://oeis.org/A161906
	divisorsAll := func() {
		const mx int = 1e6
		divisors := [mx + 1][]int{}
		for i := 1; i <= mx; i++ {
			for j := i; j <= mx; j += i {
				divisors[j] = append(divisors[j], i)
			}
		}

		{
			// https://oeis.org/A038548 Number of divisors of n that are at most sqrt(n)
			// https://oeis.org/A094820 Partial sums of A038548
			// 更细致的优化：d 与 x/d 奇偶性相同 https://codeforces.com/contest/1081/problem/E
			divisors := [mx + 1][]int{}
			for i := 1; i*i <= mx; i++ {
				for j := i * i; j <= mx; j += i {
					divisors[j] = append(divisors[j], i)
				}
			}
		}

		{
			// 统计因子个数 d(n)
			// NOTE: 复杂度可以做到线性 https://codeforces.com/contest/920/submission/76859782
			// https://oeis.org/A055507 卷积 Sum_{k=1..n} d(k)*d(n+1-k)
			// https://atcoder.jp/contests/abc292/tasks/abc292_c
			const mx int = 1e6
			d := [mx + 1]int{}
			for i := 1; i <= mx; i++ {
				for j := i; j <= mx; j += i {
					d[j]++
				}
			}
		}

		{
			// 去掉 1 作为因子
			const mx = 1e6
			divisors := [mx + 1][]int{1: {1}} // 仅保留 1 的因子 1
			for i := 2; i <= mx; i++ {
				for j := i; j <= mx; j += i {
					divisors[j] = append(divisors[j], i)
				}
			}
		}

		{
			// 线性筛求 n 的因子之和 σ(n)
			// https://codeforces.com/contest/1512/problem/G
			const mx int = 1e7
			d := make([]int, mx+1)
			d[1] = 1
			s := make([]int, mx+1)
			primes := []int{}
			for i := 2; i <= mx; i++ {
				if d[i] == 0 {
					s[i] = 1 + i
					d[i] = s[i]
					primes = append(primes, i)
				}
				for _, p := range primes {
					if p*i > mx {
						break
					}
					if i%p == 0 {
						s[p*i] = s[i]*p + 1
						d[p*i] = d[i] / s[i] * s[p*i]
						break
					}
					s[p*i] = 1 + p
					d[p*i] = d[i] * s[p*i]
				}
			}
		}

		isSquareNumber := func(x int) bool { return len(divisors[x])&1 == 1 }
		halfDivisors := func(x int) []int { d := divisors[x]; return d[:(len(d)-1)/2+1] }

		_, _ = isSquareNumber, halfDivisors
	}

	// 预处理 Squarefree numbers
	// https://oeis.org/A005117
	// https://oeis.org/wiki/Squarefree_numbers
	// 密度（见后面 calcMu 的注释）为 6/Pi^2 ≈ 0.6079
	initSquarefreeNumbers := func() []int {
		const mx int = 1e6
		free := make([]bool, mx+1)
		for i := 1; i <= mx; i++ {
			free[i] = true
		}
		for i := 2; i*i <= mx; i++ {
			for j := 1; i*i*j <= mx; j++ {
				free[i*i*j] = false
			}
		}
		// 注意特判 1
		sf := []int{}
		for i, f := range free {
			if f {
				sf = append(sf, i)
			}
		}
		return sf
	}

	// LPF(n): least prime dividing n (when n > 1); a(1) = 1 https://oeis.org/A020639
	// 有时候数据范围比较大，用 primeFactorsAll 预处理会 MLE，这时候就要用 LPF 了（同样是预处理但是内存占用低）
	// 先预处理出 LPF，然后对要处理的数 v 不断地除 LPF(v) 直到等于 1
	// 		LPF 前缀和 https://oeis.org/A046669 https://oeis.org/A088821 前缀积 https://oeis.org/A072486
	//      - a(n) ~ n^2/(2 log n)
	//		n+LPF(n) https://oeis.org/A061228 the smallest number greater than n which is not coprime to n
	// 		n-LPF(n) https://oeis.org/A046666
	// 			迭代至 0 的次数 https://oeis.org/A175126 相关题目 https://codeforces.com/contest/1076/problem/B
	//		n*LPF(n) https://oeis.org/A285109
	// 		n/LPF(n) https://oeis.org/A032742 即 n 的最大因子 = Max{gcd(n,j); j=n+1..2n-1}
	//
	//		只考虑奇质数 https://oeis.org/A078701
	//
	// GPF(n): greatest prime dividing n, for n >= 2; a(1)=1 https://oeis.org/A006530
	//		GPF(p-1) https://oeis.org/A023503
	//		GPF(p+1) https://oeis.org/A023509
	// 		GPF 前缀和 https://oeis.org/A046670 前缀积 https://oeis.org/A104350
	//		n+GPF(n) https://oeis.org/A070229 the next m>n such that GPF(n)|m
	// 		n-GPF(n) https://oeis.org/A076563
	// 			迭代至 0 的次数 https://oeis.org/A309892
	// 		n*GPF(n) https://oeis.org/A253560
	// 		n/GPF(n) https://oeis.org/A052126
	//      a(1)=1, a(n+1)=a(n)+GPF(a(n)) https://oeis.org/A076271
	//
	// 		n/LPF(n)*GPF(n) https://oeis.org/A130064
	// 		n/GPF(n)*LPF(n) https://oeis.org/A130065
	//
	// https://codeforces.com/problemset/problem/385/C
	// https://codeforces.com/gym/103107/problem/F (另一种做法是欧拉筛）
	lpfAll := func() {
		const mx int = 1e6
		lpf := [mx + 1]int{1: 1}
		for i := 2; i <= mx; i++ {
			if lpf[i] == 0 {
				for j := i; j <= mx; j += i {
					if lpf[j] == 0 { // 去掉这个判断就变成求 GPF，也可以用来（从大到小地）分解质因数
						lpf[j] = i
					}
				}
			}
		}

		{
			// 也可以用欧拉筛求，实际测试下来耗时和上面差不多
			lpf := [mx + 1]int{1: 1}
			primes := []int{} // 可以提前确定空间
			for i := 2; i <= mx; i++ {
				if lpf[i] == 0 {
					lpf[i] = i
					primes = append(primes, i)
				}
				for _, p := range primes {
					if p*i > mx {
						break
					}
					lpf[p*i] = p
					if i%p == 0 {
						break
					}
				}
			}
		}

		// EXTRA: 分解 x
		factorize := func(x int) {
			for x > 1 {
				p := lpf[x]
				e := 1
				for x /= p; lpf[x] == p; x /= p {
					e++
				}
				// do(p, e) ...

			}
		}

		// 求 x 的所有因子
		// https://codeforces.com/problemset/problem/1614/D2
		// 简单的质因子分解 https://codeforces.com/problemset/problem/762/A
		//     在因子个数比较多时，效率比试除法高
		_ds := [1024]int{1} // 复用，避免反复扩容和 GC
		divisors := func(x int) []int {
			ds := _ds[:1]
			for x > 1 {
				p := lpf[x]
				e := 1
				for x /= p; lpf[x] == p; x /= p {
					e++
				}
				d := ds
				for powP := p; e > 0; e-- {
					for _, d := range d {
						ds = append(ds, d*powP)
					}
					powP *= p
				}
			}
			return ds // append([]int(nil), ds...)
		}

		// https://oeis.org/A046951 the number of squares dividing n
		// https://codeforces.com/contest/1822/problem/G2
		squareDivisors := func(x int) []int { //
			ds := _ds[:1]
			for _, p := range primes { // 预处理 U^(1/3) 内的质数
				p2 := p * p
				if p2 > x {
					break
				}
				if x%p2 == 0 {
					e := 1
					for x /= p2; x%p2 == 0; x /= p2 {
						e++
					}
					d := ds
					for pp := p; e > 0; e-- {
						for _, d := range d {
							ds = append(ds, d*pp)
						}
						pp *= p
					}
				}
				if x%p == 0 {
					x /= p
				}
			}
			if x > 1 {
				rt := int(math.Sqrt(float64(x)))
				if rt*rt == x {
					for _, d := range ds {
						ds = append(ds, d*rt)
					}
				}
			}
			// 上面结束后，得到的是 x 的所有平方因子的【平方根】
			for i := range ds {
				ds[i] *= ds[i]
			}
			return ds
		}

		// EXTRA: https://oeis.org/A007913 Squarefree part of n (also called core(n))
		// a(n) is the smallest positive number m such that n/m is a square
		// https://oeis.org/A013928 Number of (positive) squarefree numbers < n
		// https://oeis.org/A055204 core(n!)
		//     log a(n) ~ n log 2
		//     Square root of largest square dividing n! https://oeis.org/A055772
		// https://oeis.org/A008833 n/core(n)   Largest square dividing n
		// https://oeis.org/A055071 n!/core(n!) Largest square dividing n!
		// https://codeforces.com/contest/1470/problem/B
		// https://codeforces.com/contest/1497/problem/E2
		core := func(x int) int {
			c := 1
			for x > 1 {
				p := lpf[x]
				e := 1
				for x /= p; lpf[x] == p; x /= p {
					e ^= 1
				}
				if e > 0 {
					c *= p
				}
			}
			return c
		}

		coreAll := func() {
			symDiff := func(a, b []int) []int { // 对称差
				i, n := 0, len(a)
				j, m := 0, len(b)
				res := make([]int, 0, n+m)
				for {
					if i == n {
						return append(res, b[j:]...)
					}
					if j == m {
						return append(res, a[i:]...)
					}
					if a[i] < b[j] {
						res = append(res, a[i])
						i++
					} else if a[i] > b[j] {
						res = append(res, b[j])
						j++
					} else {
						i++
						j++
					}
				}
			}

			const mx int = 1e6
			core := [mx + 1][]int{}
			np := [mx + 1]bool{}
			primes := []int{}
			for i := 2; i <= mx; i++ {
				if !np[i] {
					core[i] = []int{i} // len(primes)
					primes = append(primes, i)
				}
				for _, p := range primes {
					if p*i > mx {
						break
					}
					np[p*i] = true
					core[p*i] = symDiff(core[i], core[p])
					if i%p == 0 {
						break
					}
				}
			}

			// EXTRA: 配合 bitset 可以求最长乘积为平方数的子数组
			// 也可以用 xor hashing（附题单）https://codeforces.com/blog/entry/85900
			maxLenSquare := func(a []int) (ans int) {
				const w = bits.UintSize
				mul := [9592/w + 1]uint{} // 9592 是 mx=1e5 下的质数个数
				pos := map[[9592/w + 1]uint]int{mul: -1}
				for i, v := range a {
					for _, pi := range core[v] {
						mul[pi/w] ^= 1 << (pi % w)
					}
					if j, ok := pos[mul]; !ok {
						pos[mul] = i
					} else if i-j > ans {
						ans = i - j
					}
				}
				return ans
			}

			_ = maxLenSquare
		}

		// EXTRA: https://oeis.org/A007947 Largest squarefree number dividing n: the squarefree kernel of n, rad(n), radical of n
		// https://oeis.org/A034386 rad(n!) Primorial numbers (second definition): n# = product of primes <= n
		//                                  = rad(LCM(1,...,n))
		//                                  = LCM(core(1), core(2), core(3), ..., core(n))
		// https://oeis.org/A003557 n/rad(n)  n divided by largest squarefree divisor of n
		// https://oeis.org/A049614 n!/rad(n!)
		rad := func(x int) int {
			r := 1
			for x > 1 {
				p := lpf[x]
				r *= p
				for x /= p; lpf[x] == p; x /= p {
				}
			}
			return r
		}

		// EXTRA: https://oeis.org/A008475 质因数分解中，各个 p^e 项之和

		// EXTRA: https://oeis.org/A001414 Integer log of n: sum of primes dividing n (with repetition)
		// 质因子分解，因子之和
		// https://oeis.org/A029908 不动点
		sopfr := func(x int) (s int) {
			for x > 1 {
				for p := lpf[x]; lpf[x] == p; x /= p {
					s += p
				}
			}
			return
		}

		// EXTRA: https://oeis.org/A008472 Sum of the distinct primes dividing n
		sopf := func(x int) (s int) {
			for x > 1 {
				p := lpf[x]
				s += p
				for x /= p; lpf[x] == p; x /= p {
				}
			}
			return
		}

		_ = []interface{}{factorize, divisors, squareDivisors, core, coreAll, rad, sopfr, sopf}
	}

	// 预处理质因子
	// 例如 pf[12] = [2,3]
	// for i>=2, pf[i][0] == i means i is prime
	primeFactorsAll := func() {
		const mx int = 1e6
		pf := [mx + 1][]int{}
		for i := 2; i <= mx; i++ {
			if pf[i] == nil {
				for j := i; j <= mx; j += i {
					pf[j] = append(pf[j], i)
				}
			}
		}
	}

	// 预处理: [2,mx] 的不同的质因子个数 omega(n) https://oeis.org/A001221
	// https://en.wikipedia.org/wiki/Prime_omega_function
	// 莫比乌斯反演 https://oeis.org/A062799 = Sum_{d|n} omega(d)
	// https://oeis.org/A007875 Number of ways of writing n as p*q, with p <= q, gcd(p, q) = 1
	//                          a(n) = 2^(omega(n)-1)
	//                          相关题目 https://www.luogu.com.cn/problem/T192681?contestId=38351 https://www.luogu.com.cn/blog/LonecharmRiver/Elevator
	// max omega(<10^n)
	// 2,3,4,5,6,7,8,8,9             1~9
	// 10,10,11,12,12,13,13,14,15    10~18
	distinctPrimesCountAll := func() {
		const mx int = 1e6
		omega := make([]int, mx+1) // int8
		for i := 2; i <= mx; i++ {
			if omega[i] == 0 {
				for j := i; j <= mx; j += i {
					omega[j]++
				}
			}
		}

		{
			// 线性筛
			omega := make([]int, mx+1) // int8
			primes := []int{}
			for i := 2; i <= mx; i++ {
				if omega[i] == 0 {
					omega[i] = 1
					primes = append(primes, i)
				}
				for _, p := range primes {
					if p*i > mx {
						break
					}
					if i%p == 0 {
						omega[p*i] = omega[i]
						break
					}
					omega[p*i] = omega[i] + 1
				}
			}
		}

		// EXTRA: 前缀和，即 omega(n!) https://oeis.org/A013939
		for i := 3; i <= mx; i++ {
			omega[i] += omega[i-1]
		}

		// EXTRA:
		// https://oeis.org/A034444 Number of unitary divisors of n (d such that d divides n, gcd(d, n/d) = 1)
		// a(n) = 1<<omega[n]
		// 另一种说法是，已知 LCM(x,y) 和 GCD(x,y)，求 (x,y) 的数量
		// 由于 (x/GCD) * (y/GCD) = LCM/GCD，且 x/GCD 和 y/GCD 互质
		// 所以相当于是在求 a(LCM/GCD) = 1<<omega[LCM/GCD]
		// 相关题目 https://codeforces.com/problemset/problem/1499/D

		// EXTRA:
		// https://oeis.org/A007875 Number of ways of writing n as p*q, with p <= q, gcd(p, q) = 1
		// a(n) = 1<<(omega[n]-1)
	}

	// 欧拉函数（互质的数的个数）Euler totient function https://oeis.org/A000010
	// https://en.wikipedia.org/wiki/Euler%27s_totient_function
	// 下界 https://en.wikipedia.org/wiki/Euler%27s_totient_function#Growth_rate
	// 比较松的下界 φ(n) >= √(n/2)
	// https://oi-wiki.org/math/euler/
	// 前缀和见下面的「phi 求和相关」
	// φ(φ...(n)) 收敛到 1 的迭代次数是 log 级别的：奇数减一，偶数减半 https://oeis.org/A003434
	//      https://codeforces.com/problemset/problem/1797/E
	// φ(n!) https://oeis.org/A048855
	//      If n is prime, then a(n) = a(n-1)*(n-1)
	//      If n is composite, then a(n) = a(n-1)*n
	//      紫书例题 10-26，UVa11440 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=26&page=show_problem&problem=2435
	// GCD(n, φ(n)) https://oeis.org/A009195
	// n+φ(n) https://oeis.org/A121048
	// n-φ(n) https://oeis.org/A051953 called Cototient
	// n*φ(n) https://oeis.org/A002618 = φ(n^2)
	// φ(n)|n https://oeis.org/A007694 iff n = 1 or n = 2^w * 3^u for w > 0 and u >= 0
	// 		n/φ(n) https://oeis.org/A049237
	//      the quotients can take only 3 distinct values:
	//			n/φ(n) = 1 iff n = 1
	//			n/φ(n) = 2 iff n = 2^w, w >= 1
	//			n/φ(n) = 3 iff n = 2^w * 3^u, w >= 1, u >= 1
	// n%φ(n) https://oeis.org/A068494
	// https://oeis.org/A074693 a(1)=1; a(n+1)=a(n)+φ(a(n))
	// https://oeis.org/A345965 a(1)=1; a(n)=φ(n)+a(n/lpf(n))
	// - 相关题目 https://codeforces.com/problemset/problem/772/C
	// Least number k such that phi(k) = n https://oeis.org/A002181    Inverse of Euler totient function
	// Number of values of k such that phi(k) = n https://oeis.org/A058277
	// φ集合 https://oeis.org/A002202
	// φ补集 https://oeis.org/A007617
	// https://oeis.org/A023900 Dirichlet inverse of Euler totient function
	// https://oeis.org/A070194 小于 n 且与 n 互质的数的最大间隔
	// https://oeis.org/A023896 小于 n 且与 n 互质的数之和 a(n) = phi(n^2)/2 = n*phi(n)/2
	// https://oeis.org/A053818 小于 n 且与 n 互质的数的平方之和 If n = p_1^e_1 * ... *p_r^e_r then a(n) = n^2*phi(n)/3 + (-1)^r*p_1*..._p_r*phi(n)/6 = n^2*A000010(n)/3 + n*A023900(n)/6, n>1
	// https://oeis.org/A053819 小于 n 且与 n 互质的数的立方之和 a(n) = n^2/4*(n*A000010(n) + A023900(n)), n>1
	// https://oeis.org/A069213 第 n 个与 n 互质的数

	// phi 求和相关
	// ∑φ(i) https://oeis.org/A002088 #{(x,y): 1<=x<=y<=n, gcd(x,y)=1}
	// a(n) = (3*n^2)/(Pi^2) + O(nlogn)，近似看成 n^2 / 3
	//      = (1/2)*Sum_{k>=1} mu(k)*floor(n/k)*floor(1+n/k)
	//     相关题目 https://codeforces.com/problemset/problem/1009/D
	// 1, 2, 4, 6, 10, 12, 18, 22, 28, 32, 42, 46, 58, 64, 72, 80, 96, 102
	// ∑φ(i)-1 https://oeis.org/A015614 #{(x,y): 1<=x<y<=n, gcd(x,y)=1}
	// 0, 1, 3, 5, 9, 11, 17, 21, 27, 31, 41, 45, 57, 63, 71, 79, 95, 101

	// Number of numbers "unrelated to n" https://oeis.org/A045763
	// m < n such that m is neither a divisor of n nor relatively prime to n
	// a(n) = n + 1 - d(n) - phi(n); where d(n) is the number of divisors of n

	// Unitary totient (or unitary phi) function uphi(n) https://oeis.org/A047994

	// 计算单个数 n 的欧拉函数（互质的数的个数）Euler totient function
	calcPhi := func(n int) int {
		ans := n
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				ans = ans / i * (i - 1)
				for n /= i; n%i == 0; n /= i {
				}
			}
		}
		if n > 1 {
			ans = ans / n * (n - 1)
		}
		return ans
	}

	// 预处理 [1,mx] 欧拉函数
	initPhi := func() {
		const mx int = 1e6
		phi := [mx + 1]int{}
		for i := 1; i <= mx; i++ {
			phi[i] = i
		}
		for i := 2; i <= mx; i++ {
			if phi[i] == i {
				for j := i; j <= mx; j += i {
					phi[j] = phi[j] / i * (i - 1)
				}
			}
		}
	}

	// 线性筛 https://oi-wiki.org/math/sieve/#_8
	// https://www.luogu.com.cn/discuss/show/213297
	sievePhi := func() {
		const mx int = 1e6
		phi := [mx + 1]int{1: 1}
		primes := []int{}
		vis := [mx + 1]bool{}
		for i := 2; i <= mx; i++ {
			if !vis[i] {
				phi[i] = i - 1
				primes = append(primes, i)
			}
			for _, p := range primes {
				v := p * i
				if v > mx {
					break
				}
				vis[v] = true
				if i%p == 0 {
					phi[v] = phi[i] * p
					break
				}
				phi[v] = phi[i] * (p - 1)
			}
		}
	}

	// 欧拉定理
	// 如果 gcd(a,n) = 1，则 a^φ(n) ≡ 1(mod n)
	// 推论：如果 gcd(a,n) = 1，则 a^x ≡ 1(mod n) 的最小正整数解是 φ(n) 的因子（证明见《算法竞赛进阶指南》）
	// LC1015 https://leetcode.cn/problems/smallest-integer-divisible-by-k/ http://poj.org/problem?id=3696
	// https://atcoder.jp/contests/abc222/tasks/abc222_g
	// https://oj.socoding.cn/p/1981

	// 扩展欧拉定理（降幂公式）
	// https://oi-wiki.org/math/fermat/#_5
	// https://zhuanlan.zhihu.com/p/42632291
	// https://blog.csdn.net/synapse7/article/details/19610361
	// a^b ≡ a^(b mod φ(m)) (mod m), gcd(a,m)=1
	// a^b ≡ a^(b mod φ(m) + φ(m)) (mod m), gcd(a,m)!=1 且 b>φ(m)
	// 模板题 https://www.luogu.com.cn/problem/P5091
	// 例题 https://codeforces.com/problemset/problem/615/D
	// https://atcoder.jp/contests/abc228/tasks/abc228_e
	// https://cses.fi/problemset/task/1712
	exPhi := func(a, b int64, m int) int64 {
		phi := int64(calcPhi(m))
		if b >= phi {
			b = b%phi + phi
		}
		return pow(a, b, int64(m))
	}

	// 原根
	// https://en.wikipedia.org/wiki/Primitive_root_modulo_n
	// https://oi-wiki.org/math/primitive-root/
	// https://cp-algorithms.com/algebra/primitive-root.html
	// todo 原根&离散对数相关 https://www.luogu.com.cn/blog/command-block/yuan-gen-li-san-dui-shuo-xiang-guan
	//      浅谈离散对数问题 https://www.luogu.com.cn/blog/328405/qian-tan-li-san-dui-shuo-wen-ti
	// https://oeis.org/A033948 Numbers that have a primitive root (the multiplicative group modulo n is cyclic)
	//     The sequence consists of 1, 2, 4 and numbers of the form p^i and 2p^i, where p is an odd prime and i > 0
	// https://oeis.org/A046144 Number of primitive roots modulo n
	//    https://oeis.org/A010554 a(n) = phi(phi(n))
	// https://oeis.org/A008330 Number of primitive roots of n-th prime = phi(p-1)
	// https://oeis.org/A046145 Smallest primitive root modulo n, or 0 if no root exists
	// https://oeis.org/A001918 Smallest primitive root of n-th prime
	// https://oeis.org/A046146 Largest primitive root (<n) modulo n, or 0 if no root exists
	// https://oeis.org/A071894 Largest primitive root (<p) of n-th prime p
	// https://oeis.org/A056619 Smallest prime with primitive root of n or 0 if no such prime exists
	// https://oeis.org/A023049 Smallest prime > n having primitive root n, or 0 if no such prime exists
	//
	// 从威尔逊定理到 Gauss's generalization: 与 n 互质的数的乘积模 n 的值
	// https://en.wikipedia.org/wiki/Wilson%27s_theorem#Gauss's_generalization
	// https://math.stackexchange.com/questions/441667/the-product-of-integers-relatively-prime-to-n-congruent-to-pm-1-pmod-n
	// 相关题目 https://codeforces.com/contest/1514/problem/C
	//
	// 模板题 https://www.luogu.com.cn/problem/U125141
	//
	// 返回 n 的最小的原根, n >= 2
	// 不存在时返回 -1
	// 由于有 phi(phi(n)) 个原根，密度足够大，最小原根可以很快找到，复杂度约为 O(n^0.25logn)
	primitiveRoot := func(n int64, calcPhi func(int64) int64) int64 {
		if n != 2 && n != 4 {
			x := n
			if x&1 == 0 {
				x /= 2
			}
			if x&1 == 0 || len(primeDivisors(x)) > 1 {
				return -1
			}
		}

		pn := calcPhi(n)
		ps := primeDivisors(pn)
	o:
		for g := int64(1); ; g++ {
			if gcd(g, n) > 1 {
				continue
			}
			for _, p := range ps {
				if pow(g, pn/p, n) == 1 {
					continue o
				}
			}
			return g
		}
	}

	// 返回 n 的所有原根
	// n 没有原根时返回空切片
	// 模板题 https://www.luogu.com.cn/problem/P6091
	primitiveRootsAll := func(n int, primitiveRoot func(int) int, gcd func(int, int) int) []int {
		rt0 := primitiveRoot(n)
		if rt0 < 0 {
			return nil
		}
		pn := calcPhi(n)
		ppn := calcPhi(pn)
		roots := make([]int, 0, ppn)
		for i, rtPow := 1, 1; len(roots) < ppn; i++ {
			rtPow = rtPow * rt0 % n
			if gcd(i, pn) == 1 {
				roots = append(roots, rtPow)
			}
		}
		sort.Ints(roots)
		return roots
	}

	/* 同余 逆元
	https://oeis.org/A006254 a(n) = 2^-1 mod p         Numbers n such that 2n-1 is prime
	https://oeis.org/A283419 3^-1 mod p
	https://oeis.org/A006093 (p-1)^-1 mod p     prime(n) - 1
	https://oeis.org/A040976 (p/2)^-1 mod p     prime(n) - 2
	https://oeis.org/A090938 Least multiple of n == 1 (mod prime(n))
	https://oeis.org/A090939 Least multiple of n == -1 (mod prime(n))
	https://oeis.org/A091185 a(n) = A090938(n)/n      n^-1 mod prime(n)
	*/
	{
		/* 涉及到 0 与逆元的题目（mod 为质数）
		例如 res = mod^k * x % mod，后面要除掉 mod^k，得到 x（如果前面直接取模会得到 0，没法保留 x 的信息）
		这个时候可以用二元组 (k, x) 来表示 res%mod（这里 k>=0）
		如果 k>0 那么 res=0
		乘法运算 (k1, x1) * (k2, x2) = (k1+k2, x1*x2)
		除法运算 (k1, x1) / (k2, x2) = (k1-k2, x1*inv(x2))  这里 k1>=k2
		加法运算（见下面的 add1）
		https://codeforces.com/contest/1848/problem/E
		https://codeforces.com/problemset/problem/543/D
		https://ac.nowcoder.com/acm/contest/39759/D
		*/
		type pair struct {
			k int
			x int64
		}
		// 计算 (k,x) + (0,1)
		// https://codeforces.com/problemset/problem/543/D
		add1 := func(p pair) pair {
			if p.k > 0 {
				return pair{0, 1}
			}
			if p.x == mod-1 {
				return pair{1, 1}
			}
			return pair{0, p.x + 1}
		}
		mul := func(p, q pair) pair { return pair{p.k + q.k, p.x * q.x % mod} }
		div := func(p, q pair) pair { return pair{p.k - q.k, p.x * pow(q.x, mod-2, mod) % mod} }
		// p%mod 的实际值
		val := func(p pair) int64 {
			if p.k > 0 {
				return 0
			}
			return p.x
		}

		_ = []any{add1, mul, div, val}
	}

	// 二元一次不定方程（线性丢番图方程中的一种）https://en.wikipedia.org/wiki/Diophantine_equation
	// exgcd solve equation ax+by=gcd(a,b)
	// 特解满足 |x|<=|b|, |y|<=|a|
	// https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
	var exgcd func(a, b int64) (gcd, x, y int64)
	exgcd = func(a, b int64) (gcd, x, y int64) {
		if b == 0 {
			return a, 1, 0
		}
		gcd, y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}

	// 任意非零模数逆元 ax ≡ 1 (mod m)，要求 |gcd(a,m)| = 1     注：不要求 m 为质数
	// 返回最小正整数解
	// 模板题 https://www.luogu.com.cn/problem/P1082
	// https://codeforces.com/problemset/problem/772/C
	invM := func(a, m int64) int64 {
		g, x, _ := exgcd(a, m)
		if g != 1 && g != -1 {
			return -1
		}
		return (x%m + m) % m
	}

	// ax ≡ b (mod m)，要求 gcd(a,m) | b       注：不要求 m 为质数
	// 或者，ax-b 是 m 的倍数，求最小非负整数 x
	// 或者，求 ax-km = b 的一个最小非负整数解
	// 示例代码 https://codeforces.com/contest/1748/submission/205834351
	invM2 := func(a, b, m int64) int64 {
		g, x, _ := exgcd(a, m)
		if b%g != 0 {
			return -1
		}
		x *= b / g
		m /= g
		return (x%m + m) % m
	}

	// a*x + b*y = c 的通解为
	// x = (c/g)*x0 + (b/g)*k
	// y = (c/g)*y0 - (a/g)*k
	// 其中 g = gcd(a,b) 且需要满足 g|c
	// x0 和 y0 是 ax+by=g 的一组特解（即 exgcd(a,b) 的返回值）
	//
	// 为方便讨论，这里要求输入的 a b c 必须为正整数
	// 注意若输入超过 1e9 可能要用高精
	// 返回：正整数解的个数（无解时为 -1，无正整数解时为 0）
	//      x 取最小正整数时的解 x1 y1，此时 y1 是最大正整数解
	//      y 取最小正整数时的解 x2 y2，此时 x2 是最大正整数解
	// 相关论文 THE NUMBER OF SOLUTIONS TO ax + by = n http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.376.403
	// 相关题目 https://www.luogu.com.cn/problem/P5656
	// 使非负解 x+y 尽量小 https://codeforces.com/problemset/problem/1244/C
	//    最简单的做法就是 min(x1+y1, x2+y2)
	// 需要转换一下符号 https://atcoder.jp/contests/abc186/tasks/abc186_e
	// https://codeforces.com/problemset/problem/1748/D
	solveLinearDiophantineEquations := func(a, b, c int64) (n, x1, y1, x2, y2 int64) {
		g, x0, y0 := exgcd(a, b)

		// 无解
		if c%g != 0 {
			n = -1
			return
		}

		a /= g
		b /= g
		c /= g
		x0 *= c
		y0 *= c

		x1 = x0 % b
		if x1 <= 0 { // 若要求的是非负整数解，去掉等号
			x1 += b
		}
		k1 := (x1 - x0) / b
		y1 = y0 - k1*a

		y2 = y0 % a
		if y2 <= 0 { // 若要求的是非负整数解，去掉等号
			y2 += a
		}
		k2 := (y0 - y2) / a
		x2 = x0 + k2*b

		// 无正整数解
		if y1 <= 0 {
			return
		}

		// k 越大 x 越大
		n = k2 - k1 + 1
		return
	}

	// 关于 ax+by+cz=n 的解的个数（三币种找零问题）
	// On the number of solutions of the Diophantine equation of Frobenius – General case https://core.ac.uk/download/pdf/14375587.pdf
	// The Number of Solutions to ax + by + cz = n and its Relation to Quadratic Residues https://cs.uwaterloo.ca/journals/JIS/VOL23/Binner/binner4.pdf
	// 上面这篇提出了一个 O(log max(a,b,c)) 的算法来求 N(a,b,c;n)
	// https://oeis.org/A001399 N(1,2,3;n) = round((n+3)^2/12)
	// https://oeis.org/A000115 N(1,2,5;n) = round((n+4)^2/20)
	// https://oeis.org/A008672 N(1,3,5;n) = round((n+3)*(n+6)/30)  =  floor((n^3+9n+30)/30)
	// https://oeis.org/A005044 N(2,3,4;n) = round(n^2/12)-floor(n/4)*floor((n+2)/4)      a(n) = round(n^2/48) if n is even; a(n) = round((n+3)^2/48) if n is odd
	// https://oeis.org/A025795 N(2,3,5;n)
	// https://oeis.org/A008680 N(3,4,5;n)

	// 费马小定理求质数逆元
	// ax ≡ 1 (mod p)
	// x^-1 ≡ a^(p-2) (mod p)
	// 滑窗 https://codeforces.com/contest/1833/problem/F
	invP := func(a, p int64) int64 {
		if a <= 0 {
			panic(-1)
		}
		return pow(a, p-2, p)
	}

	// 有理数取模
	// 模板题 https://www.luogu.com.cn/problem/P2613
	divM := func(a, b, m int64) int64 { return a * invM(b, m) % m }
	divP := func(a, b, p int64) int64 { return a * invP(b, p) % p }

	// 一种递归求单个逆元的方法
	// https://codeforces.com/blog/entry/84150?#comment-716585

	// 线性求逆元·其一
	// 求 1^-1, 2^-1, ..., (p−1)^-1 mod p
	// http://blog.miskcoo.com/2014/09/linear-find-all-invert
	// https://www.zhihu.com/question/59033693
	// 模板题 https://www.luogu.com.cn/problem/P3811
	{
		const mod = 998244353
		const mx int = 1e6
		inv := [mx + 1]int64{}
		inv[1] = 1
		for i := 2; i <= mx; i++ {
			inv[i] = int64(mod-mod/i) * inv[mod%i] % mod
		}
	}

	// 线性求逆元·其二（离线逆元）
	// 求 a1, a2, ..., an mod p 的逆元
	// 根据 ai^-1 ≡ Πai/ai * (Πai)^-1 (mod p)，求出 Πai 的前缀积和后缀积可以得到 Πai/ai，从而求出 ai^-1 mod p
	// https://zhuanlan.zhihu.com/p/86561431
	// 模板题 https://www.luogu.com.cn/problem/P5431
	calcAllInv := func(a []int64, p int64) []int64 {
		n := len(a)
		pre := make([]int64, n+1)
		pre[0] = 1
		for i, v := range a {
			pre[i+1] = pre[i] * v % p
		}
		invMulAll := invP(pre[n], p)
		suf := make([]int64, n+1)
		suf[n] = 1
		for i := len(a) - 1; i > 0; i-- { // i=0 不用求
			suf[i] = suf[i+1] * a[i] % p
		}
		inv := make([]int64, n)
		for i, pm := range pre[:n] {
			inv[i] = pm * suf[i+1] % p * invMulAll % p
		}
		return inv
	}

	// 模数两两互质的线性同余方程组 - 中国剩余定理 (CRT)
	// x ≡ bi (mod mi), bi 与 mi 互质且 Πmi <= 1e18
	// https://blog.csdn.net/synapse7/article/details/9946013
	// https://codeforces.com/blog/entry/61290
	// 模板题 https://www.luogu.com.cn/problem/P1495
	crt := func(B, M []int64) (x int64) {
		m := int64(1)
		for _, mi := range M {
			m *= mi
		}
		for i, mi := range M {
			Mi := m / mi
			_, inv, _ := exgcd(Mi, mi)
			x = (x + B[i]*Mi*inv) % m
		}
		x = (x + m) % m
		return
	}

	// 扩展中国剩余定理 (EXCRT)
	// ai * x ≡ bi (mod mi)
	// 解为 x ≡ b (mod m)
	// 有解时返回 (b, m)，无解时返回 (0, -1)
	// 推导过程见《挑战程序设计竞赛》P292
	// 注意乘法溢出的可能
	// 推荐 https://blog.csdn.net/niiick/article/details/80229217
	// todo 模板题 https://www.luogu.com.cn/problem/P4777 https://www.luogu.com.cn/problem/P4774
	// https://codeforces.com/contest/1500/problem/B
	excrt := func(A, B, M []int64) (x, m int64) {
		m = 1
		for i, mi := range M {
			a, b := A[i]*m, B[i]-A[i]*x
			d := gcd(a, mi)
			if b%d != 0 {
				return 0, -1
			}
			t := divM(b/d, a/d, mi/d)
			x += m * t
			m *= mi / d
		}
		x = (x%m + m) % m
		return
	}

	// 另一种写法，参考进阶指南 p.155
	// todo 待整理
	// excrt := func(a, m []int) (x int) {
	//	x = a[0]
	//	M := m[0]
	//	for i := 1; i < len(a); i++ {
	//		mi := m[i]
	//		c := (a[i] - x%mi + mi) % mi
	//		gcd, inv, _ := exgcd(M, mi)
	//		if c%gcd != 0 {
	//			return -1
	//		}
	//		c /= gcd
	//		mi /= gcd
	//		inv = inv * c % mi
	//		x += inv * M
	//		M *= mi
	//		x %= M
	//	}
	//	x = (x + M) % M
	//	return
	//}

	// 离散对数 - 小步大步算法 (BSGS)
	// a^x ≡ b (mod p)，a 和 p 互质
	// 有解时返回 x，无解时返回 -99，这样可以让 exBSGS 中的 +1 操作不影响无解的判断
	// 时间复杂度 O(√p)
	// 见进阶指南 p.155
	// todo https://www.luogu.com.cn/blog/command-block/yuan-gen-li-san-dui-shuo-xiang-guan
	//      http://blog.miskcoo.com/2015/05/discrete-logarithm-problem
	//      https://www.luogu.com.cn/blog/hzoiliuchang/shuo-lun-zhi-bsgs-suan-fa
	//
	// 模板题 https://www.luogu.com.cn/problem/P3846
	// https://atcoder.jp/contests/abc222/tasks/abc222_g
	// todo https://atcoder.jp/contests/abc270/tasks/abc270_g
	babyStepGiantStep := func(a, b, p, k int64) int64 { // 非 exBSGS 下 k=1
		b %= p
		t := int64(math.Sqrt(float64(p))) + 1
		mp := map[int64]int64{}
		for j, x := int64(0), b; j < t; j++ {
			mp[b] = j
			x = x * a % p
		}
		a = pow(a, t, p)
		if a == 0 {
			if b == 0 {
				return 1
			}
			return -99
		}
		for i, x := int64(0), k; i < t; i++ {
			if j, ok := mp[x]; ok && i*t >= j {
				return i*t - j
			}
			x = x * a % p
		}
		return -99
	}

	// 拓展大步小步算法
	// a^x ≡ b (mod m)，a 和 m 不一定互质
	// https://zhuanlan.zhihu.com/p/132603308
	// 模板题 https://www.luogu.com.cn/problem/P4195
	var _exBSGS func(a, b, m, k int64) int64
	_exBSGS = func(a, b, m, k int64) int64 {
		if b == 1 {
			return 0
		}
		if a == 0 && b == 0 {
			return 1
		}
		g := gcd(a, m)
		if b%g > 0 {
			return -99
		}
		if g == 1 {
			return babyStepGiantStep(a, b, m, k%m)
		}
		return 1 + _exBSGS(a, b/g, m/g, k*a/g%m)
	}

	exBSGS := func(a, b, m int64) int64 {
		x := _exBSGS(a%m, b%m, m, 1)
		phiM := int64(calcPhi(int(m)))
		if x > phiM {
			x = x%phiM + phiM
		}
		return x
	}

	// 二次剩余 x^2 ≡ a (mod p)       平方剩余
	// 一个数 a，如果不是 p 的倍数且模 p 同余于某个数的平方，则称 a 为模 p 的二次剩余
	// https://en.wikipedia.org/wiki/Quadratic_residue
	// https://en.wikipedia.org/wiki/Cipolla%27s_algorithm
	// https://oi-wiki.org/math/quad-residue/
	// https://blog.csdn.net/doyouseeman/article/details/52033204
	// Tonelli-Shanks https://www.luogu.com.cn/blog/242973/solution-p5491
	// 模板题 https://www.luogu.com.cn/problem/P5491
	modSqrt := func(x, p int64) []int64 { // p 必须是奇素数
		if x == 0 {
			return []int64{0}
		}
		x0 := new(big.Int).ModSqrt(big.NewInt(x), big.NewInt(p))
		if x0 == nil {
			return nil
		}
		// 如果要求小的在前，注意判断
		return []int64{x0.Int64(), p - x0.Int64()}
	}

	// 判断 a 是否为模 p 的二次剩余，p 必须是奇素数
	// Jacobi 符号为 -1
	isQuadraticResidue := func(a, p int64) bool {
		return big.Jacobi(big.NewInt(a), big.NewInt(p)) < 0
	}

	// todo N 次剩余 / 高次同余方程 x^a ≡ b (mod p)
	// todo 模板题 https://www.luogu.com.cn/problem/P5668

	// https://oeis.org/A072994 Number of solutions to x^n ≡ 1 (mod n), 1<=x<=n
	// Least k > 0 such that the number of solutions to x^k == 1 (mod k) 1 <= x <= k is equal to n, or 0 if no such k exists https://oeis.org/A072995

	// https://oeis.org/A182865 Minimal number of quadratic residues
	// a(n) is the least integer m such that any nonzero square is congruent (mod n) to one of the squares from 1 to m^2
	// 把这题的 1000 改成 i，则至多需要枚举到 a(i) https://ac.nowcoder.com/acm/contest/6489/A

	// https://oeis.org/A000224 Number of distinct squares residues mod n
	// Multiplicative with a(p^e) = floor(p^e/6) + 2 IF p = 2 ELSE floor(p^(e+1)/(2p + 2)) + 1
	// https://oeis.org/A046530 Number of distinct cubic residues mod n

	/* 阶乘 组合数/二项式系数 */

	// https://en.wikipedia.org/wiki/Factorial
	// https://oeis.org/A000142
	// https://en.wikipedia.org/wiki/Stirling%27s_approximation
	// n! ~ √(2πn)*(n/e)^n
	factorial := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800 /*10!*/, 39916800, 479001600}

	// https://oeis.org/A008904  n! 的最后一个非 0 数字  a(n) is the final nonzero digit of n!
	// https://math.stackexchange.com/questions/130352/last-non-zero-digit-of-a-factorial
	// https://blog.csdn.net/LuckilyYu/article/details/2078993
	// 1, 1, 2, 6, 4, 2, 2, 4, 2, 8, 8, 8, 6, 8, 2, 8, 8, 6, 8, 2, 4, 4, 8, 4, 6, 4, 4, 8, 4, 6, 8, 8, 6, 8, 2, 2, 2
	// 趣题 https://math.stackexchange.com/questions/3334779/what-is-the-last-non-zero-digit-of-dots2018-underset-text-occurs-1009
	/*
		def a(n: int) -> int:
		    if n <= 1:
		        return 1
		    return 6 * [1, 1, 2, 6, 4, 4, 4, 8, 4, 6][n % 10] * (3 ** (n // 5 % 4)) * a(n // 5) % 10
	*/

	// 等差数列的乘积转换成阶乘 https://atcoder.jp/contests/m-solutions2019/tasks/m_solutions2019_e

	// https://oeis.org/A003070 a(n) = ceiling(log_2(n!))
	// https://oeis.org/A067850 Highest power of 2 not exceeding n!
	// https://oeis.org/A049606 Largest odd divisor of n!
	// https://oeis.org/A240533 a(n) = numerators of n!/10^n

	calcFactorial := func(n int) int64 {
		res := int64(1) % mod
		for i := 2; i <= n; i++ {
			res = res * int64(i) % mod
		}
		return res
	}
	// n 小于 1 时返回 1
	calcFactorialBig := func(n int) *big.Int {
		return new(big.Int).MulRange(1, int64(n))
	}

	initFactorial := func() {
		const mx int = 1e6
		F := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			F[i] = F[i-1] * int64(i) % mod
		}
	}

	// 阶乘模质数（质数较小）
	// 时间复杂度 O(plogn)
	// todo 待整理 https://cp-algorithms.com/algebra/factorial-modulo.html
	// todo O(√n logn) 见 https://www.luogu.com.cn/problem/P5282
	_factorial := func(n, p int) int {
		res := 1
		for ; n > 1; n /= p {
			if n/p&1 == 1 {
				res = res * (p - 1) % p
			}
			for i := 2; i <= n%p; i++ {
				res = res * i % p
			}
		}
		return res
	}

	// 双阶乘
	// https://en.wikipedia.org/wiki/Double_factorial

	// 偶阶乘
	// https://oeis.org/A000165 Double factorial of even numbers: (2n)!! = 2^n*n!
	calcEvenFactorialBig := func(n int) *big.Int {
		return new(big.Int).Lsh(new(big.Int).MulRange(1, int64(n)), uint(n))
	}

	// 奇阶乘
	// https://oeis.org/A001147 Double factorial of odd numbers: (2*n-1)!! = 1*3*5*...*(2*n-1) = A(2*n,n) / 2^n
	// 1, 3, 15, 105, 945, 10395, 135135, 2027025, 34459425, 654729075, 13749310575, 316234143225, 7905853580625, ...
	// Number of ways to choose n disjoint pairs of items from 2*n items
	// Number of perfect matchings in the complete graph K(2n)
	// https://atcoder.jp/contests/abc236/tasks/abc236_d
	// LC1359 有效的快递序列数目 https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/
	// 奇阶乘模 2^64 http://acm.hdu.edu.cn/showproblem.php?pid=6481 https://www.90yang.com/hdu6481-a-math-problem/
	calcOddFactorialBig := func(n int) *big.Int {
		return new(big.Int).Rsh(new(big.Int).MulRange(int64(n+1), int64(2*n)), uint(n))
	}

	// https://oeis.org/A002109 Hyperfactorials: Product_{k=1..n} k^k

	// https://oeis.org/A010786 Floor-factorial numbers: a(n) = Product_{k=1..n} floor(n/k)
	// 1, 2, 3, 8, 10, 36, 42, 128, 216, 600, 660, 3456, 3744, 9408, 18900, 61440, 65280, 279936, 295488, 1152000, 2116800, 4878720, 5100480, 31850496, 41472000, 93450240, 163762560, 568995840, 589317120, 3265920000, 3374784000
	// https://oeis.org/A309912 a(n) = Product_{p prime, p <= n} floor(n/p)

	// GCD(C(n,1),C(n,2),...,C(n,n-1))
	// = p, n = p^k (p is a prime)
	// = 1, otherwise

	// binomial(n, floor(n/2)) https://oeis.org/A001405
	// a(n) ~ 2^n / sqrt(π * n/2)     O(2^n/sqrt(n))
	// a(2n) ~ 4^n / sqrt(πn)         O(4^n/sqrt(n))
	// 从一个大小为 n 的集合的子集中随机选一个，选到 floor(n/2) 大小的子集的概率约为 1 / sqrt(π * n/2)
	// Sperner's theorem says that this is the maximal number of subsets of an n-set such that no one contains another
	// EXTRA: https://oeis.org/A000984 Central binomial coefficients: binomial(2*n,n) = (2*n)!/(n!)^2
	// EXTRA: https://oeis.org/A100071 a(n) = n * A001405(n-1) = 1, 2, 6, 12, 30, 60, 140, 280, 630, 1260, ...
	//                                 a(n) = a(n-1) * n / floor(n/2)
	// EXTRA: https://oeis.org/A107373 a(n) = (n/2) * A001405(n-1) - 2^(n-2)
	combHalf := []int64{
		1, 1, 2, 3, 6, 10, 20, 35, 70, 126, // C(9,4)
		252, 462, 924, 1716, 3432, 6435, 12870, 24310, 48620, 92378, // C(19,9)
		184756, 352716, 705432, 1352078, 2704156, 5200300, 10400600, 20058300, 40116600, 77558760, // C(29,14)
		155117520, 300540195, 601080390, 1166803110, 2333606220, 4537567650, 9075135300, 17672631900, 35345263800, 68923264410, // C(39,19)
		137846528820, 269128937220, 538257874440, 1052049481860, 2104098963720, 4116715363800, 8233430727600, 16123801841550, 32247603683100, 63205303218876, // C(49,24)
		126410606437752, 247959266474052, 495918532948104, 973469712824056, 1946939425648112, 3824345300380220, 7648690600760440, 15033633249770520, 30067266499541040, 59132290782430712, // C(59,29)
		118264581564861424, 232714176627630544, 465428353255261088, 916312070471295267, 1832624140942590534, 3609714217008132870, 7219428434016265740, // C(66,33)
	}

	// 组合数/二项式系数
	// 不取模，仅适用于小范围的 n 和 k
	initComb := func() {
		const mx = 60
		C := [mx + 1][mx + 1]int64{}
		for i := 0; i <= mx; i++ {
			C[i][0], C[i][i] = 1, 1
			for j := 1; j < i; j++ {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}

	// O(n) 预处理阶乘及其逆元，O(1) 求组合数
	{
		const mod int64 = 1e9 + 7
		const mx int = 2e6
		F := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			F[i] = F[i-1] * int64(i) % mod
		}
		pow := func(x, n int64) int64 {
			//x %= mod
			res := int64(1)
			for ; n > 0; n >>= 1 {
				if n&1 == 1 {
					res = res * x % mod
				}
				x = x * x % mod
			}
			return res
		}
		invF := [...]int64{mx: pow(F[mx], mod-2)}
		for i := mx; i > 0; i-- {
			invF[i-1] = invF[i] * int64(i) % mod
		}
		C := func(n, k int) int64 {
			if k < 0 || k > n {
				return 0
			}
			return F[n] * invF[k] % mod * invF[n-k] % mod
		}
		P := func(n, k int) int64 {
			if k < 0 || k > n {
				return 0
			}
			return F[n] * invF[n-k] % mod
		}

		// EXTRA: 卢卡斯定理 https://en.wikipedia.org/wiki/Lucas%27s_theorem
		// https://yangty.blog.luogu.org/lucas-theorem-note
		// C(n,m)%p = C(n%p,m%p) * C(n/p,m/p) % p
		// 注意初始化 F 和 invF 时 mx 取 mod-1
		// 推论：n&m==m 时 C(n,m) 为奇数，否则为偶数
		// https://www.luogu.com.cn/problem/P3807
		// https://www.luogu.com.cn/problem/P7386
		var lucas func(n, k int64) int64
		lucas = func(n, k int64) int64 {
			if k == 0 {
				return 1
			}
			return C(int(n%mod), int(k%mod)) * lucas(n/mod, k/mod) % mod
		}

		// 库默尔定理 https://en.wikipedia.org/wiki/Kummer%27s_theorem
		// todo https://atcoder.jp/contests/arc137/tasks/arc137_d

		// 可重组合 https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition
		// 方案数 H(n,k)=C(n+k-1,k) https://oeis.org/A059481
		// 相当于把 k 个无区别的球放入 n 个有区别的盒子中，且允许空盒的方案数
		//		隔板法：把 n 个盒子当做 n-1 个隔板，这样相当于总共有 k+n-1个位置，从中选择 k 个位置放球，剩下的位置放隔板。这样就把 k 个球划分成了 n 份，放入对应的盒子中
		// NOTE: 若改成「至多放 k 个球」，则等价于多了一个盒子，用来放「不放入盒子的球」
		// NOTE: mx 要开两倍空间！
		H := func(n, k int) int64 { return C(n+k-1, k) }
		// 也相当于，给出元素取值种类数 kinds 和序列长度 length，求有多少种非降序列
		// 也可以理解成在 length * (kinds-1) 的网格上走单调路径
		H = func(kinds, length int) int64 { return C(kinds+length-1, length) }

		// 卡特兰数 Cn = C(2n,n)/(n+1) = C(2n,n)-C(2n,n-1)
		// C(n) = 2*(2*n-1)*C(n-1)/(n+1) with C(0) = 1
		// 证明见这个视频末尾 https://www.bilibili.com/video/BV1E8411f7Mu/?t=33m16s
		// https://en.wikipedia.org/wiki/Catalan_number
		// https://oeis.org/A000108
		// 从 n=0 开始：1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862, 16796, 58786, 208012, 742900, 2674440, 9694845, 35357670, 129644790
		// 所有在 n×n 格点中不越过对角线的单调路径的个数
		// Number of noncrossing partitions of the n-set (不相交握手问题) LC1259 https://leetcode-cn.com/problems/handshakes-that-dont-cross/
		// Dyck Path https://mathworld.wolfram.com/DyckPath.html
		// https://www.luogu.com.cn/problem/P1641
		// 
		// https://codeforces.com/problemset/problem/1830/C
		//
		// 将全部偶数提取一个 2，可得 (2n)! = 1*3*5*...*(2n-1) * (2^n) * (n!)
		// 故 C(2*n,n)/(n+1) = (2*n)!/(n!)/(n+1)! = 1*3*5*...*(2n-1)*(2^n)/(n+1)!
		// 又由于 n! 的 2 的因子个数 = n/2 + n/4 + ... + n/2^k <= n-1 当且仅当 n 为 2^k 时取到等号
		// 对比分子分母的 2 的因子个数，可以得出如下结论：
		//     当且仅当 n+1 为 2^k 时，卡特兰数为奇数
		//
		// EXTRA: 高维的情况 https://loj.ac/p/6051
		Catalan := func(n int) int64 { return F[2*n] * invF[n+1] % mod * invF[n] % mod }
		Catalan = func(n int) int64 { return new(big.Int).Rem(new(big.Int).Div(new(big.Int).Binomial(int64(2*n), int64(n)), big.NewInt(int64(n+1))), big.NewInt(mod)).Int64() }

		// 默慈金数 Motzkin number https://oeis.org/A001006
		// 从 (0,0) 移动到 (n,0) 的网格路径数，每步只能向右移动一格（可以向右上、右下、横向向右），并禁止移动到 y=0 以下的地方
		// M(n) = Sum_{i=0..n/2} C(n,2*i)*Catalan(i)
		// https://en.wikipedia.org/wiki/Motzkin_number
		// 包含生成函数 https://mathworld.wolfram.com/MotzkinNumber.html
		// 生成函数推导 https://zhuanlan.zhihu.com/p/187502941
		// https://blog.csdn.net/acdreamers/article/details/41213667
		// http://acm.hdu.edu.cn/showproblem.php?pid=3723
		Motzkin := func(n int) (res int64) {
			for i := 0; i <= n/2; i++ {
				res = (res + C(n, 2*i)*Catalan(i)) % mod
			}
			return
		}

		// EXTRA: 若仅限定起点为 (0,0)，终点可以是任意 (n,i) https://oeis.org/A005773
		// a(0)=1, a(n) = Sum_{k=0..n-1} M(k)*a(n-k-1)

		// EXTRA: 起点为 (0,i)，终点为 (n,j) https://oeis.org/A081113 Number of paths of length n-1 a king can take from one side of an n X n chessboard to the opposite side
		// a(n) = number of sequences (a_1,a_2,...,a_n) with 1<=a_i<=n for all i and |a_(i+1)-a_(i)|<=1 for 1<=i<=n-1
		// a(n) = Sum_{k=1..n} k*(n-k+1)*M(n-1, k-1) where M() is the Motzkin triangle https://oeis.org/A026300
		// 1, 4, 17, 68, 259, 950, 3387, 11814, 40503, 136946, 457795, 1515926, 4979777, 16246924, 52694573, 170028792, 546148863, 1747255194, 5569898331, 17698806798, 56076828573, 177208108824, 558658899825, 1757365514652

		// 那罗延数 Narayana number (Narayana triangle) https://oeis.org/A001263
		// 从 (0,0) 移动到 (2n,0) 且恰好有 k 个山峰的网格路径数，每步只能向右上或右下移动一格（不能横向），并禁止移动到 x 轴以下的地方
		// N(n,k) = C(n,k)*C(n,k-1)/n
		// https://en.wikipedia.org/wiki/Narayana_number

		// Fuss-Catalan 数、(m-1)-Dyck 路与 m 叉树 https://www.luogu.com.cn/blog/your-alpha1022/solution-p2767

		// 某些组合题可能用到
		pow2 := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			pow2[i] = pow2[i-1] << 1 % mod
		}

		_ = []interface{}{C, P, H, Catalan, Motzkin}
	}

	// 适用于 n 巨大但 k 或 n-k 较小的情况
	comb := func(n, k int) int64 {
		if k > n-k {
			k = n - k
		}
		var a, b int64 = 1, 1
		for i := 1; i <= k; i++ {
			a = a * int64(n) % mod
			n--
			b = b * int64(i) % mod
		}
		return divP(a, b, mod)
	}

	// 另类组合数求法
	{
		const mod int64 = 1e9 + 7
		var n, k int64
		// 注意当 n 为负数时，可能会算出非 0 的结果，这种情况要特判
		// 当 0 <= n < k 时结果为 0
		_ = new(big.Int).Binomial(n, k).Int64() // small
		_ = new(big.Int).Rem(new(big.Int).Binomial(n, k), big.NewInt(mod)).Int64()
		_ = int64(math.Round(math.Gamma(float64(n+1)) / math.Gamma(float64(k+1)) / math.Gamma(float64(n-k+1))))
	}

	// 扩展卢卡斯
	// todo https://blog.csdn.net/niiick/article/details/81064156
	// https://blog.csdn.net/skywalkert/article/details/52553048
	// https://blog.csdn.net/skywalkert/article/details/104681101
	// https://cp-algorithms.com/combinatorics/binomial-coefficients.html
	// 模板题 https://www.luogu.com.cn/problem/P4720
	// 古代猪文 https://www.luogu.com.cn/problem/P2480

	// 斯特林数（斯特林轮换数，斯特林子集数）
	// https://en.wikipedia.org/wiki/Stirling_number
	// https://oi-wiki.org/math/stirling/
	// todo 斯特林数的四种求法 https://www.luogu.com.cn/blog/command-block/si-te-lin-shuo-zong-jie
	// todo https://www.luogu.com.cn/blog/xzc/zu-ge-shuo-xue-hu-si-te-lin-shuo
	// https://blog.csdn.net/ACdreamers/article/details/8521134
	// 一个对称的斯特林数恒等式及其扩展 https://blog.csdn.net/EI_Captain/article/details/108806153
	// 【第一类斯特林数】s(n,k) https://oeis.org/A008275
	// 绝对值版本 https://oeis.org/A130534 coefficients of the polynomial (x+1)(x+2)...(x+n), expanded in increasing powers of x
	//    定义为对应递降阶乘展开式的各项系数
	//    将 n 个元素排成 k 个非空循环排列的方法数
	//    s(n,k) 的递推公式： s(n,k)=(n-1)*s(n-1,k)+s(n-1,k-1), 1<=k<=n-1
	//    边界条件：s(n,0)=0, n>=1    s(n,n)=1, n>=0
	//    LC1866 https://leetcode-cn.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/
	//    建筑师 https://www.luogu.com.cn/problem/P4609
	//    UVa1638 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4513
	//    todo https://www.luogu.com.cn/problem/P5408
	//         https://www.luogu.com.cn/problem/P5409
	//         https://codeforces.com/problemset/problem/1516/E
	// 【第二类斯特林数】S2(n,k) https://oeis.org/A008277
	//    将 n 个元素拆分为 k 个非空集的方法数
	//    用容斥计算单个项 S2(n,k) = (1/k!) * ∑{i=0..k} (-1)^(k-i)*C(k, i)*i^n
	//         https://codeforces.com/problemset/problem/1342/E
	//    S2(n,k) 的递推公式：S2(n,k)=k*S2(n-1,k)+S2(n-1,k-1), 1<=k<=n-1
	//    边界条件：S(n,0)=0, n>=1    S(n,n)=1, n>=0
	//    LC1692 https://leetcode.cn/problems/count-ways-to-distribute-candies/
	//    https://www.luogu.com.cn/problem/P5395
	//    todo https://www.luogu.com.cn/problem/P5396
	//    https://oeis.org/A019538 n 个位置，每个位置填 [1,k] 之间的数，要求每个数字至少出现一次 => k!*S2(n,k)
	// Generalized Stirling numbers: a(n) = n! * Sum_{k=0..n-1} (k+1)/(n-k) https://oeis.org/A001705
	// - 对 1~n 的每个排列，计算其后缀最小值数组去重后的元素之和
	// Unsigned Stirling numbers of first kind: s(n+1,2): a(n+1) = (n+1)*a(n) + n! https://oeis.org/A000254
	// todo 斯特林数，斯特林反演初探 https://www.yijan.co/si-te-lin-shu-si-te-lin-fan-yan-chu-tan/
	// todo https://codeforces.com/contest/1278/problem/F 洛谷有艹标算的题解
	stirling1 := func(n, k int) int64 {
		s := make([][]int64, n+1)
		for i := range s {
			s[i] = make([]int64, n+1) // K+1
		}
		s[0][0] = 1
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ { // j <= K
				s[i][j] = (s[i-1][j-1] + int64(i-1)*s[i-1][j]) % mod
			}
		}
		return s[n][k]
	}

	stirling2 := func(n, k int) int64 {
		s2 := make([][]int64, n+1)
		for i := range s2 {
			s2[i] = make([]int64, n+1)
		}
		s2[0][0] = 1
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				s2[i][j] = (s2[i-1][j-1] + int64(j)*s2[i-1][j]) % mod
			}
		}
		return s2[n][k]
	}

	// 第二类斯特林数·行
	// https://www.luogu.com.cn/problem/P5395
	stirling2RowPoly := func(n int) poly {
		F := make([]int64, n+1)
		F[0] = 1
		for i := 1; i <= n; i++ {
			F[i] = F[i-1] * int64(i) % P
		}
		invF := make(poly, n+1)
		invF[n] = _pow(F[n], P-2)
		for i := n; i > 0; i-- {
			invF[i-1] = invF[i] * int64(i) % P
		}
		a := make(poly, n+1)
		b := make(poly, n+1)
		for i, v := range invF {
			if i&1 == 0 {
				a[i] = v
			} else {
				a[i] = P - v
			}
			b[i] = _pow(int64(i), n) * v % P
		}
		return a.conv(b)[:n+1]
	}

	// 贝尔数：基数为 n 的集合的划分方法数 https://oeis.org/A000110
	// https://en.wikipedia.org/wiki/Bell_number
	// 1, 1, 2, 5, 15, 52, 203, 877, 4140, 21147, 115975, 678570, 4213597, 27644437, 190899322, 1382958545, ...
	// https://en.wikipedia.org/wiki/Bell_triangle https://oeis.org/A011971 Aitken's array
	// a(0,0)=1, a(n,0) = a(n-1,n-1), a(n,k) = a(n,k-1) + a(n-1,k-1)
	// 其他公式
	// B(n+1) = Sum_{k=0..n} C(n,k)*B(k)
	// B(n) = Sum_{k=1..n} S2(n,k)
	bellTriangle := func(n int) [][]int64 {
		b := make([][]int64, n+1) // 第一列为贝尔数
		b[0] = []int64{1}
		for i := 1; i <= n; i++ {
			b[i] = make([]int64, i+1)
			b[i][0] = b[i-1][i-1]
			for j := 1; j <= i; j++ {
				b[i][j] = (b[i][j-1] + b[i-1][j-1]) % mod
			}
		}
		return b
	}

	// 贝尔数的多项式求法
	// https://blog.csdn.net/a_forever_dream/article/details/106489066
	// https://www.luogu.com.cn/problem/P5748
	bellPoly := func(n int) poly {
		F := make([]int64, n+1)
		F[0] = 1
		for i := 1; i <= n; i++ {
			F[i] = F[i-1] * int64(i) % P
		}
		invF := make(poly, n+1)
		invF[n] = _pow(F[n], P-2)
		for i := n; i > 1; i-- { // 注意为了计算下面的 exp，invF[0] = 0
			invF[i-1] = invF[i] * int64(i) % P
		}

		b := invF.exp()
		for i, v := range b {
			b[i] = v * F[i] % P
		}
		return b
	}

	// 贝尔数 EXTRA：如何搜索所有集合划分
	// 相关题目：https://codeforces.com/contest/954/problem/I
	setPartition := func(n int) {
		groups := [][]int{} // 或者用一个 roots 数组表示集合的根节点（代表元）
		var f func(int)
		f = func(p int) {
			if p == n {
				// do groups ...

				return
			}
			groups = append(groups, []int{p})
			f(p + 1)
			groups = groups[:len(groups)-1]
			for i := range groups {
				groups[i] = append(groups[i], p)
				f(p + 1)
				groups[i] = groups[i][:len(groups[i])-1]
			}
		}
		f(0)
	}

	// 富比尼数（有序贝尔数）
	// Count the number of weak orderings on a set of n elements
	// 即允许平局下，n 人比赛的结果数
	// https://en.wikipedia.org/wiki/Ordered_Bell_number https://oeis.org/A000670
	// a(n) = Sum_{k=0..n} k! * S2(n,k)
	// a(n) = Sum_{k=1..n} C(n,k) * a(n-k), a(0) = 1
	// 相关题目：UVa12034 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=242&page=show_problem&problem=3185

	// 欧拉数 https://oeis.org/A008292
	// https://en.wikipedia.org/wiki/Eulerian_number
	// T(n, k) = k * T(n-1, k) + (n-k+1) * T(n-1, k-1), T(1, 1) = 1
	// T(n, k) = Sum_{j=0..k} (-1)^j * (k-j)^n * C(n+1, j)
	// todo 浅谈欧拉数 https://www.luogu.com.cn/blog/Karry5307/eulerian-numbers

	//

	// 莫比乌斯函数 Möbius function μ(n) https://oeis.org/A008683
	// μ(1)=1
	// μ(n)=0 if n 含有平方因子
	// μ(n)=(-1)^k, k=omega(n)
	// https://en.wikipedia.org/wiki/M%C3%B6bius_function
	// https://oi-wiki.org/math/mobius/#_13
	// φ(n) = Sum_{d|n} d*μ(n/d)
	// 部分φ(n) = 不超过 m 的数中与 n 互质的数的个数（m 可以大于 n）
	//          = Sum_{d|n} d*μ(m/d)    用 n 的因子来容斥
	//          https://codeforces.com/problemset/problem/920/G
	// 线性筛 https://oi-wiki.org/math/sieve/#_9
	// 前缀和 https://oeis.org/A002321 Mertens's function 梅滕斯函数
	//    https://en.wikipedia.org/wiki/Mertens_function
	//    |a(n)| = O(x^(1/2 + eps))
	//    零点 https://oeis.org/A028442
	// https://oeis.org/A013928 ∑μ^2(x) = Sum_{d = 1..floor(sqrt(n-1)} mu(d)*floor((n-1)/d^2)
	//                                  = 6*n/Pi^2 + O(sqrt(n))
	//                                  也就是说，随机一个正整数，squarefree 的概率 = 6/Pi^2 ≈ 0.6079
	// https://oeis.org/A030229 μ(x)=1 的数
	// https://oeis.org/A013929 μ(x)=0 的数
	// https://oeis.org/A030059 μ(x)=-1 的数
	// https://oeis.org/A005117 μ(x)!=0 的数（即 squarefree numbers）
	calcMu := func(n int) int {
		if n == 1 {
			return 1
		}
		mu := 1
		for i := 2; i*i <= n; i++ {
			e := 0
			for ; n%i == 0; n /= i {
				e++
			}
			if e > 1 {
				return 0
			}
			if e > 0 {
				mu = -mu
			}
		}
		if n > 1 {
			mu = -mu
		}
		return mu
	}

	// 线性筛写法
	sieveMu := func() {
		const mx int = 1e6
		mu := [mx + 1]int{1: 1} // int8
		primes := []int{}
		vis := [mx + 1]bool{}
		for i := 2; i <= mx; i++ {
			if !vis[i] {
				mu[i] = -1
				primes = append(primes, i)
			}
			for _, p := range primes {
				v := p * i
				if v > mx {
					break
				}
				vis[v] = true
				if i%p == 0 {
					mu[v] = 0
					break
				}
				mu[v] = -mu[i]
			}
		}
	}

	// 埃氏筛写法
	sieveMu = func() {
		const mx int = 1e6
		mu := [mx + 1]int{1: 1} // int8
		for i := 1; i <= mx; i++ {
			for j := i * 2; j <= mx; j += i {
				mu[j] -= mu[i]
			}
		}
	}

	// 狄利克雷卷积 Dirichlet convolution
	// https://en.wikipedia.org/wiki/Dirichlet_convolution
	// https://zhuanlan.zhihu.com/p/137619492

	// 莫比乌斯反演 Möbius inversion formula
	// https://en.wikipedia.org/wiki/M%C3%B6bius_inversion_formula
	// todo 题目 https://oi-wiki.org/math/mobius/
	// https://zhuanlan.zhihu.com/p/138038817
	// 莫比乌斯反演-让我们从基础开始 https://www.luogu.com.cn/blog/An-Amazing-Blog/mu-bi-wu-si-fan-yan-ji-ge-ji-miao-di-dong-xi
	// https://www.luogu.com.cn/blog/61088/jian-dan-shuo-lun-tian-keng
	// [Tutorial] Generalized Möbius Inversion on Posets https://codeforces.com/blog/entry/98413
	//
	// todo 专题练习[一些好玩的数学题] https://www.luogu.com.cn/training/1432
	// https://codeforces.com/problemset/problem/547/C
	// https://codeforces.com/problemset/problem/900/D
	// https://codeforces.com/problemset/problem/1559/E
	// todo https://www.luogu.com.cn/problem/P2257
	//  https://www.luogu.com.cn/problem/P2522
	//  https://www.luogu.com.cn/blog/203623/sol-jrksjr6D https://www.luogu.com.cn/blog/Silver187/qian-lian-di-shi-jie-ti-xie
	//  GCD=1 的子序列个数 https://codeforces.com/problemset/problem/803/F https://ac.nowcoder.com/acm/problem/112055
	//  GCD=1 的数对个数 * (i-j) https://codeforces.com/problemset/problem/1780/F

	// todo 推式子 https://ac.nowcoder.com/acm/contest/11171/E

	//

	// 数论分块/除法分块/整除分块
	// https://oi-wiki.org/math/mobius/#_3
	//     https://oeis.org/A006218
	//     a(n) = ∑{k=1..n} floor(n/k)
	//          = n * (log(n) + 2*gamma - 1) + O(sqrt(n))
	//          也就是说平均每项的贡献约为 log(n)
	//     a(n) = ∑{k=1..n} floor(n/k)
	//          = 2*(∑{i=1..floor(sqrt(n))} floor(n/i)) - floor(sqrt(n))^2
	//          因此 a(n) % 2 == floor(sqrt(n)) % 2
	//     a(n) 前缀和 = Sum_{k=1..n-1} Sum_{i=1..n-1} floor(k/i) https://oeis.org/A078567
	// 恒等式 n%i = n-(n/i)*i
	// ∑n/i https://www.luogu.com.cn/problem/P1403 n=1e18 的做法（整点计数问题）见 https://www.luogu.com.cn/problem/SP26073 https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/solution/by-hqztrue-lv4e/
	// ∑k%i 见下面的 floorLoopK
	// ∑(n/i)*(n%i) https://ac.nowcoder.com/acm/contest/9005/C
	// todo https://codeforces.com/contest/1202/problem/F
	// ∑∑(n%i)*(m%j) 代码见下面的 floorLoop2 https://www.luogu.com.cn/problem/P2260
	// [L,R] 内任意 k 个不同数字的 GCD 有多少种 https://ac.nowcoder.com/acm/contest/35232/C
	// https://codeforces.com/problemset/problem/1789/E
	// https://codeforces.com/problemset/problem/938/C
	// https://codeforces.com/problemset/problem/449/A
	// 数论分块优化 DP https://codeforces.com/problemset/problem/1603/C
	//
	// https://oeis.org/A257212           Least d>0 such that floor(n/d) - floor(n/(d+1)) <= 1
	// https://oeis.org/A257213 mex(n/i); Least d>0 such that floor(n/d) = floor(n/(d+1))
	//
	// EXTRA: 一些另类的求和
	// https://oeis.org/A116477 a(n) = Sum_{1<=k<=n, gcd(k,n)=1} floor(n/k)
	//                          sum{k|n} a(k) = sum{k=1 to n} d(k) = https://oeis.org/A006218
	// https://oeis.org/A013939 a(n) = Sum_{k = 1..n} floor(n/prime(k)) = omega(n!)
	//
	// EXTRA: n/k (k=1..n) 的不同数字的个数 https://oeis.org/A055086
	//        = floor(sqrt(4*n+1)) - 1 (注意用 int64)
	floorLoop := func(n int64) (sum int64) {
		for l, r := int64(1), int64(0); l <= n; l = r + 1 {
			h := n / l
			r = n / h
			w := r - l + 1
			sum += h * w // for all i in [l,r], floor(n/i) = floor(n/l)
		}
		return
	}

	// ∑x/i, i in [low,up]
	// 转换 https://codeforces.com/problemset/problem/1485/C
	floorLoopRange := func(low, up, x int64, min func(int64, int64) int64) (sum int64) {
		for l, r := low, int64(0); l <= up; l = r + 1 {
			h := x / l
			if h == 0 {
				break
			}
			r = min(x/h, up)
			w := r - l + 1
			sum += h * w
		}
		return
	}

	// 余数求和
	//   ∑k%i (i in [1,n])
	// = ∑k-(k/i)*i
	// = n*k-∑(k/i)*i
	// 对于 [l,r] 范围内的 i，k/i 不变，此时 ∑(k/i)*i = (k/i)*∑i = (k/i)*(l+r)*(r-l+1)/2
	// https://www.luogu.com.cn/problem/P2261
	// https://codeforces.com/problemset/problem/616/E
	// NEERC05，紫书例题 10-25，UVa1363 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=446&page=show_problem&problem=4109 https://codeforces.com/gym/101334 J
	floorLoopRem := func(n, k int64, min func(int64, int64) int64) int64 {
		sum := n * k
		for l, r := int64(1), int64(0); l <= n; l = r + 1 {
			h := k / l
			if h > 0 {
				r = min(k/h, n)
			} else {
				r = n
			}
			w := r - l + 1
			s := (l + r) * w / 2
			sum -= h * s
		}
		return sum
	}

	// 二维整除分块
	// ∑{i=1..min(n,m)} floor(n/i)*floor(m/i)
	// https://www.luogu.com.cn/blog/command-block/zheng-chu-fen-kuai-ru-men-xiao-ji
	// todo ∑∑(n%i)*(m%j) 模积和 https://www.luogu.com.cn/problem/P2260
	floorLoop2D := func(n, m int64, min func(int64, int64) int64) (sum int64) {
		for l, r := int64(1), int64(0); l <= min(n, m); l = r + 1 {
			hn, hm := n/l, m/l
			r = min(n/hn, m/hm)
			w := r - l + 1
			sum += hn * hm * w
		}
		return
	}

	// 杜教筛 - 积性函数前缀和
	// 复杂度 O(n^(2/3))
	// https://zhuanlan.zhihu.com/p/258336043
	// https://blog.csdn.net/weixin_43914593/article/details/104229700 算法竞赛专题解析（4）：杜教筛--以及积性函数的前世今生
	// https://www.luogu.com.cn/blog/command-block/du-jiao-shai
	// http://baihacker.github.io/main/
	// The prefix-sum of multiplicative function: the black algorithm http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_the_black_algorithm.html
	// The prefix-sum of multiplicative function: Dirichlet convolution http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_dirichlet_convolution.html
	// The prefix-sum of multiplicative function: powerful number sieve http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_powerful_number_sieve.html
	// todo 浅谈一类积性函数的前缀和 + 套题 https://blog.csdn.net/skywalkert/article/details/50500009
	// 模板题 https://www.luogu.com.cn/problem/P4213
	// todo ∑∑i*j*gcd(i,j) https://www.luogu.com.cn/problem/solution/P3768
	sieveDu := func() {
		const mx int = 1e6
		phi := [mx + 1]int64{1: 1}
		mu := [mx + 1]int{1: 1}
		primes := []int{}
		vis := [mx + 1]bool{}
		for i := 2; i <= mx; i++ {
			if !vis[i] {
				phi[i] = int64(i - 1)
				mu[i] = -1
				primes = append(primes, i)
			}
			for _, p := range primes {
				v := p * i
				if v > mx {
					break
				}
				vis[v] = true
				if i%p == 0 {
					phi[v] = phi[i] * int64(p)
					mu[v] = 0
					break
				}
				phi[v] = phi[i] * int64(p-1)
				mu[v] = -mu[i]
			}
		}
		for i := 0; i < mx; i++ {
			phi[i+1] += phi[i]
			mu[i+1] += mu[i]
		}

		cachePhi := map[int]int64{}
		var sumPhi func(int) int64
		sumPhi = func(n int) int64 {
			if n <= mx {
				return phi[n]
			}
			if s := cachePhi[n]; s > 0 {
				return s
			}
			m := int64(n)
			res := m * (m + 1) / 2
			for l, r := int64(2), int64(0); l <= m; l = r + 1 {
				h := m / l
				r = m / h
				res -= (r - l + 1) * sumPhi(int(h))
			}
			cachePhi[n] = res
			return res
		}

		cacheMu := map[int]int{}
		var sumMu func(int) int
		sumMu = func(n int) int {
			if n <= mx {
				return mu[n]
			}
			if s, has := cacheMu[n]; has {
				return s
			}
			res := 1
			for l, r := 2, 0; l <= n; l = r + 1 {
				h := n / l
				r = n / h
				res -= (r - l + 1) * sumMu(h)
			}
			cacheMu[n] = res
			return res
		}
	}

	// Min25 筛 - 积性函数前缀和
	// https://zhuanlan.zhihu.com/p/60378354
	// https://oi-wiki.org/math/min-25/
	// https://codeforces.com/blog/entry/92703
	// todo 模板题 https://www.luogu.com.cn/problem/P5325

	// 一篇新论文，复杂度为 O((nlogn)^(3/5))
	// Summing μ(n): a faster elementary algorithm
	// https://arxiv.org/pdf/2101.08773.pdf

	//

	// 埃及分数 - 不同的单位分数的和 (IDA*)
	// https://www.luogu.com.cn/problem/UVA12558
	// 贪婪算法：将一项分数分解成若干项单分子分数后的项数最少，称为第一种好算法；最大的分母数值最小，称为第二种好算法
	// https://en.wikipedia.org/wiki/Egyptian_fraction
	// https://oeis.org/A006585 number of solutions
	// https://oeis.org/A247765 Table of denominators in the Egyptian fraction representation of n/(n+1) by the greedy algorithm
	// https://oeis.org/A100678 Number of Egyptian fractions in the representation of n/(n+1) via the greedy algorithm
	// https://oeis.org/A100695	Largest denominator used in the Egyptian fraction representation of n/(n+1) by the greedy algorithm
	//
	// 		埃尔德什-施特劳斯猜想（Erdős–Straus conjecture）https://en.wikipedia.org/wiki/Erd%C5%91s%E2%80%93Straus_conjecture

	/* 斐波那契数列 F(n) https://oeis.org/A000045
	性质 https://oi-wiki.org/math/fibonacci/
	快速计算 F(1e18) https://codeforces.com/blog/entry/14516
	https://oeis.org/A000071 F(n) 前缀和 = F(n+2)-1
	∑k=[0,n]C(n,k)F(k) = F(2n)
	https://oeis.org/A007598 F^2(n)    a(n) = 2*a(n-1) + 2*a(n-2) - a(n-3), n > 2. a(0)=0, a(1)=1, a(2)=1
	                                   a(n) = (F(n)*F(n+4)-3*F(n)*F(n+1))/2
	https://oeis.org/A001690 补集
	https://oeis.org/A022307 F(n) 的不同的质因子个数
	https://oeis.org/A001175 π(n) = {F}%n 的周期（皮萨诺周期）  Pisano periods / Pisano numbers https://en.wikipedia.org/wiki/Pisano_period
	                         π(n) ≤ 6n, with equality if and only if n = 2 · 5r, for r ≥ 1
	https://oeis.org/A060305 π(p) = {F}%p 的周期 https://blog.csdn.net/acdreamers/article/details/10983813
	https://oeis.org/A003893 F(n)%10
	https://oeis.org/A001605 使 F(n) 为质数的 n
	https://oeis.org/A191797 C(F(n), 2)
	https://oeis.org/A001177 Fibonacci entry points: a(n) = least k >= 1 such that n divides Fibonacci number
	异或和 F(n) 1,0,2,1,4,12,1,20,54,1,88,200,33,344,826,225,1756,3268,7313,1788
	定义 f(m) 为最小的满足 F(i)+F(j) ≡ 0 (mod m) 的 i (j<=i)，f(m) 大概是 O(√m) 的
	todo https://codeforces.com/problemset/problem/226/C

	其他相关序列
	https://oeis.org/A000213 Tribonacci numbers: a(n)=a(n-1)+a(n-2)+a(n-3) with a(0)=a(1)=a(2)=1
	https://oeis.org/A000931 Padovan sequence (or Padovan numbers): a(n)=a(n-2)+a(n-3) with a(0)=1, a(1)=a(2)=0
	https://oeis.org/A001045 Jacobsthal sequence (or Jacobsthal numbers): a(n)=a(n-1)+2*a(n-2) = (2^n - (-1)^n)/3, with a(0) = 0, a(1) = 1
	https://oeis.org/A112387 This sequence originated from the Fibonacci sequence, but instead of adding the last two terms, you get the average.
	    https://codeforces.com/problemset/problem/696/C
	*/

	// https://oeis.org/A195264 Iterate x -> A080670(x) (replace x with the concatenation of the primes and exponents in its prime factorization)
	// starting at n until reach 1 or a prime; or -1 if a prime is never reached
	// https://www.zhihu.com/question/48612677/answer/487252829

	_ = []interface{}{
		primes, primes10k, primes10, primes10_,
		sqCheck, cubeCheck, sqrt, cbrt, bottomDiff,
		gcd, gcdPrefix, gcdSuffix, lcm, lcms, makeFrac, lessFrac, countDifferentSubsequenceGCDs, floorSum,
		isPrime, sieve, sieveEuler, sieveEulerTemplate, factorize, primeDivisors, primeDivisors2, powerOfFactorialPrimeDivisor, primeExponentsCountAll, primeExponentsCount,
		divisors, divisorsO1Space, oddDivisorsNum, maxSqrtDivisor, divisorsAll, primeFactorsAll, lpfAll, initSquarefreeNumbers, distinctPrimesCountAll,
		calcPhi, initPhi, sievePhi, exPhi,
		primitiveRoot, primitiveRootsAll,
		exgcd, solveLinearDiophantineEquations, invM, invM2, invP, divM, divP, calcAllInv,
		crt, excrt,
		babyStepGiantStep, exBSGS,
		modSqrt, isQuadraticResidue,
		factorial, calcFactorial, calcFactorialBig, initFactorial, _factorial, calcEvenFactorialBig, calcOddFactorialBig, combHalf, initComb, comb,
		stirling1, stirling2, stirling2RowPoly,
		bellTriangle, bellPoly, setPartition,
		calcMu, sieveMu,
		floorLoop, floorLoopRange, floorLoopRem, floorLoop2D,
		sieveDu,
	}
}

/* hack overflow
1<<32 + 95168 = 2^6 * 3 * 7^5 * 11^3
1<<32 + 66304 = 2^8 * 5^2 * 11 * 13^2 * 19^2
1<<32 + 48704 = 2^6 * 3^2 * 5^3 * 11^2 * 17 * 29
1<<32 - 49216 = 2^6 * 3^7 * 5 * 17 * 19^2

1<<32 - 49216 => https://github.com/LeetCode-Feedback/LeetCode-Feedback/issues/13613 hack Java
a := []int{2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 5, 17, 19, 19}
for i, n := 0, len(a); i < n/2; i++ {
	a[i], a[n-1-i] = a[n-1-i], a[i]
}

Print("[")
for i := range a {
	Print("[", i+1, ",", i+2, "],")
}
cur := len(a) + 2
for i, c := range a {
	for j := 1; j < c; j++ {
		Print("[", i+1, ",", cur, "],")
		cur++
	}
}
Print("]")
*/
