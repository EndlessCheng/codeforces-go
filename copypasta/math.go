package copypasta

import (
	. "fmt"
	"math"
	"math/big"
	"math/bits"
	"math/rand"
)

/* 数论 组合数学 博弈论 趣味数学

https://en.wikipedia.org/wiki/List_of_recreational_number_theory_topics

NOTE: a%-b == a%b

AP: Sn = n*(2*a1+(n-1)*d)/2
GP: Sn = a1*(pow(q,n)-1)/(q-1), q!=1
       = a1*n, q==1

https://oeis.org/A001787 n * 2^(n-1) = 0~n-1 的 1 的个数之和
https://oeis.org/A000337 (n-1) * 2^n + 1
	Number of zeros in binary numbers 1 to 111..1 (n+1 bits)

https://oeis.org/A036289 n * 2^n
https://oeis.org/A036799 (n-1) * 2^(n+1) + 2

https://oeis.org/A027992 a(n) = 2^n*(3n-1)+2 = The total sum of squares of parts in all compositions of n （做 https://codeforces.com/problemset/problem/235/B 时找到的序列）
https://oeis.org/A271638 a(n) = (13*n-36)*2^(n-1)+6*n+18 = 	The total sum of the cubes of all parts of all compositions of n
	线性的情况是上面的 A001787

https://en.wikipedia.org/wiki/Faulhaber%27s_formula
https://oeis.org/A000330 平方和 = n*(n+1)*(2*n+1)/6
https://oeis.org/A000537 立方和 = (n*(n+1)/2)^2

https://oeis.org/A000295 Eulerian numbers: Sum_{k=0..n} (n-k)*2^k = 2^n - n - 1
	Number of permutations of {1,2,...,n} with exactly one descent
	Number of partitions of an n-set having exactly one block of size > 1
	a(n-1) is the number of subsets of {1..n} in which the largest element of the set exceeds by at least 2 the next largest element
		For example, for n = 5, a(4) = 11 and the 11 sets are {1,3}, {1,4}, {1,5}, {2,4}, {2,5}, {3,5}, {1,2,4}, {1,2,5}, {1,3,5}, {2,3,5}, {1,2,3,5}
	a(n-1) is also the number of subsets of {1..n} in which the second smallest element of the set exceeds by at least 2 the smallest element
		For example, for n = 5, a(4) = 11 and the 11 sets are {1,3}, {1,4}, {1,5}, {2,4}, {2,5}, {3,5}, {1,3,4}, {1,3,5}, {1,4,5}, {2,4,5}, {1,3,4,5}

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

椭圆曲线加密算法 https://ac.nowcoder.com/acm/contest/6916/C

挑战 2.6 节练习题
2429
1930
3126
3421
3292
3641
1995
4.1 节练习题（模运算的世界）
1150
1284
2115
3708
2720
GCJ Japan 2011 Final B

CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=number+theory
CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=combinatorics

*/

func numberTheoryCollection() {
	const mod int64 = 1e9 + 7 // 998244353
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

	TIPS: 一般 LCM 的题目都需要用 LCM=x*y/GCD 转换成研究 GCD 的性质

	GCD 与质因子 https://codeforces.com/problemset/problem/264/B
	数组中最小的 LCM(ai,aj) https://codeforces.com/problemset/problem/1154/G
	分拆与 LCM  https://ac.nowcoder.com/acm/contest/5961/D https://ac.nowcoder.com/discuss/439005

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

	// 前 n 个数的 LCM https://oeis.org/A003418
	// a(n)/a(n-1) = https://oeis.org/A014963
	// Mangoldt Function
	// https://mathworld.wolfram.com/MangoldtFunction.html

	// GCD 性质统计相关
	// NOTE: 对于一任意非负序列，前 i 个数的 GCD 是非增序列，且至多有 O(logMax) 个不同值
	//       应用：https://codeforces.com/problemset/problem/1210/C
	// #{(a,b) | 1<=a<=b<=n, gcd(a,b)=1}   https://oeis.org/A002088
	//     = ∑phi(i)
	// #{(a,b) | 1<=a,b<=n, gcd(a,b)=1}   https://oeis.org/A018805
	//     = 2*(∑phi(i))-1
	//     = 2*A002088(n)-1
	// #{(a,b,c) | 1<=a,b,c<=n, gcd(a,b,c)=1}   https://oeis.org/A071778
	//     = ∑mu(i)*floor(n/i)^3
	// #{(a,b,c,d) | 1<=a,b,c,d<=n, gcd(a,b,c,d)=1}   https://oeis.org/A082540
	//     = ∑mu(i)*floor(n/i)^4

	// GCD 求和相关
	// ∑gcd(n,i) = ∑{d|n}d*phi(n/d)   https://oeis.org/A018804
	// ∑n/gcd(n,i) = ∑{d|n}d*phi(d)   https://oeis.org/A057660
	// ∑∑gcd(i,j) = ∑phi(i)*(floor(n/i))^2   https://oeis.org/A018806

	// LCM 求和相关
	// ∑lcm(n,i) = n*(1+∑{d|n}d*phi(d))/2 = n*(1+A057660(n))/2   https://oeis.org/A051193
	// ∑lcm(n,i)/n = A051193(n)/n = (1+∑{d|n}d*phi(d))/2 = (1+A057660(n))/2   https://oeis.org/A057661
	// ∑∑lcm(i,j)   https://oeis.org/A064951

	// 最简分数
	type pair struct{ x, y int64 }
	frac := func(a, b int64) pair { g := gcd(a, b); return pair{a / g, b / g} }

	// todo 待整理
	// 给定数组，统计所有区间的 GCD 值
	// 返回 map[GCD值]等于该值的区间个数
	cntRangeGCD := func(arr []int64) map[int64]int64 {
		n := len(arr)
		cntMp := map[int64]int64{}
		gcds := make([]int64, n)
		copy(gcds, arr)
		lPos := make([]int, n)
		for i, v := range arr {
			lPos[i] = i
			// 从当前位置 i 往左遍历，更新 gcd[j] 的同时维护等于 gcd[j] 的区间最左端位置
			for j := i; j >= 0; j = lPos[j] - 1 {
				gcds[j] = gcd(gcds[j], v)
				g := gcds[j]
				for lPos[j] > 0 && gcd(gcds[lPos[j]-1], v) == g {
					lPos[j] = lPos[lPos[j]-1]
				}
				// [lPos[j], j], [lPos[j]+1, j], ..., [j, j] 的区间 GCD 值均等于 gcd[j]
				cntMp[g] += int64(j - lPos[j] + 1)
			}
		}
		return cntMp
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
	bottomDiff := func(a []int) int {
		for ; len(a) > 1; a = a[:len(a)-1] {
			for i := 0; i+1 < len(a); i++ {
				a[i] = a[i+1] - a[i]
			}
		}
		return a[0]
	}

	/* 质数 质因子分解 */

	// n/2^k http://oeis.org/A000265

	// 质数表 https://oeis.org/A000040
	// primes[i]%10 http://oeis.org/A007652
	// 10-primes[i]%10 http://oeis.org/A072003
	// p-1 http://oeis.org/A006093
	// p+1 http://oeis.org/A008864
	// p^2+p+1 http://oeis.org/A060800 = sigma(p^2)
	primes := [...]int{
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

	// map{小于 10^n 的素数个数: 小于 10^n 的最大素数} http://oeis.org/A006880 http://oeis.org/A003618   10^n-a(n): http://oeis.org/A033874
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

	// 大于 10^n 的最小素数 http://oeis.org/A090226 http://oeis.org/A003617    a(n)-10^n: http://oeis.org/A033873
	primes10_ := [...]int64{
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

	质数的幂次组成的集合 {p^k} https://oeis.org/A000961
		补集 https://oeis.org/A024619
		Exponential of Mangoldt function https://oeis.org/A014963

	质数前缀和 http://oeis.org/A007504
	a(n)^2 - a(n-1)^2 = A034960(n)
	EXTRA: divide odd numbers into groups with prime(n) elements and add together http://oeis.org/A034960

	质数前缀积 prime(n)# https://oeis.org/A002110
	the least number with n distinct prime factors
	2, 6, 30, 210, 2310, 30030, 510510, 9699690, 223092870, /9/
	6469693230, 200560490130, 7420738134810, 304250263527210, 13082761331670030, 614889782588491410

	质数差分 Gap https://oeis.org/A001223
		Positions of records https://oeis.org/A005669
		Values of records https://oeis.org/A005250
		Gap 均值 https://oeis.org/A286888 a(n)= floor((prime(n) - 2)/(n - 1))
		相关题目 https://www.luogu.com.cn/problem/P6104 https://class.luogu.com.cn/classroom/lgr69
		Numbers whose distance to the closest prime number is a prime number http://oeis.org/A160666

	任意质数之差 https://oeis.org/A030173
	非任意质数之差 https://oeis.org/A007921

	质数的逆二项变换 Inverse binomial transform of primes https://oeis.org/A007442

	合数前缀和 https://oeis.org/A053767

	合数前缀积 Compositorial number https://oeis.org/A036691

	不与质数相邻的合数 http://oeis.org/A079364

	半素数 https://oeis.org/A001358 也叫双素数/二次殆素数 Semiprimes (or biprimes): products of two primes
		https://en.wikipedia.org/wiki/Semiprime
		https://en.wikipedia.org/wiki/Almost_prime
		非平方半素数 https://oeis.org/A006881 Squarefree semiprimes: Numbers that are the product of two distinct primes.

	绝对素数 https://oeis.org/A003459 各位数字可以任意交换位置，其结果仍为素数
		https://en.wikipedia.org/wiki/Permutable_prime

	哥德巴赫猜想 - 偶数分拆的最小质数 Goldbach’s conjecture https://oeis.org/A020481
	由质数分布可知选到一对质数的概率是 O(1/ln^2(n))
	https://en.wikipedia.org/wiki/Goldbach%27s_conjecture
		Positions of records https://oeis.org/A025018
		Values of records https://oeis.org/A025019
		1e9 内最大的为 a(721013438) = 1789
		2e9 内最大的为 a(1847133842) = 1861

	勒让德猜想 - 在两个相邻平方数之间，至少有一个质数 Legendre’s conjecture
	https://en.wikipedia.org/wiki/Legendre%27s_conjecture
	Number of primes between n^2 and (n+1)^2 https://oeis.org/A014085
	Number of primes between n^3 and (n+1)^3 https://oeis.org/A060199

	伯特兰-切比雪夫定理 - n ~ 2n 之间至少有一个质数 Bertrand's postulate
	https://en.wikipedia.org/wiki/Bertrand%27s_postulate
	Number of primes between n and 2n (inclusive) https://oeis.org/A035250
	Number of primes between n and 2n exclusive https://oeis.org/A060715

	Least k such that H(k) > n, where H(k) is the harmonic number Σ{i=1..k} 1/i
	https://oeis.org/A002387
	https://oeis.org/A004080

		a(n) = smallest prime p such that Σ{primes q = 2, ..., p} 1/q exceeds n
		5, 277, 5_195_977, 1801241230056600523
		https://oeis.org/A016088

	a(n) = largest m such that the harmonic number H(m)= Σ{i=1..m} 1/i is < n
	https://oeis.org/A115515

		a(n) = largest prime p such that Σ{primes q = 2, ..., p} 1/q does not exceed n
		3, 271, 5_195_969, 1801241230056600467
		https://oeis.org/A223037

	Exponent of highest power of 2 dividing n, a.k.a. the binary carry sequence, the ruler sequence, or the 2-adic valuation of n
	a(n) = 0 if n is odd, otherwise 1 + a(n/2)
	http://oeis.org/A007814

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
	isPrime = func(n int64) bool { return big.NewInt(n).ProbablyPrime(0) }

	// 预处理: [2,mx] 范围内的质数
	// 埃拉托斯特尼筛法 Sieve of Eratosthenes
	// 也有线性时间的算法，见 https://oi-wiki.org/math/sieve/ 以及进阶指南 p.136-137
	// 质数个数 π(n) https://oeis.org/A000720
	//         π(10^n) https://oeis.org/A006880
	//         4, 25, 168, 1229, 9592, 78498, 664579, 5761455, 50847534, /* 1e9 */
	//         455052511, 4118054813, 37607912018, 346065536839, 3204941750802, 29844570422669, 279238341033925, 2623557157654233, 24739954287740860, 234057667276344607,
	sieve := func() {
		const mx int = 1e6
		primes := []int{}
		pid := [mx + 1]int{-1, -1}
		for i := 2; i <= mx; i++ {
			if pid[i] == 0 {
				primes = append(primes, i)
				pid[i] = len(primes)
				for j := 2 * i; j <= mx; j += i {
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

	// 线性筛
	// 避免多次标记合数
	// https://www.luogu.com.cn/problem/solution/P3383
	// https://www.luogu.com.cn/problem/P3383
	sieveL := func() {
		const mx int = 1e7
		primes := []int{}
		pid := [mx + 1]int{-1, -1}
		for i := 2; i <= mx; i++ {
			if pid[i] == 0 {
				primes = append(primes, i)
				pid[i] = len(primes)
			}
			for _, p := range primes {
				if p*i > mx {
					break
				}
				pid[p*i] = -1
				if i%p == 0 {
					break
				}
			}
		}
	}

	// 区间筛法
	// 预处理 [2,√R] 的所有质数，去筛 [L,R] 之间的质数

	// 质因数分解 prime factorization
	// 返回分解出的质数及其指数
	// https://mathworld.wolfram.com/PrimeFactorization.html
	// todo 更高效的算法 - Pollard's Rho
	// n 的质因数分解中 2 的幂次 http://oeis.org/A007814
	// n 的质因数分解中非 2 的幂次之和 http://oeis.org/A087436
	primeFactorization := func(n int) (factors [][2]int) {
		for i := 2; i*i <= n; i++ {
			e := 0
			for ; n%i == 0; n /= i {
				e++
			}
			if e > 0 {
				factors = append(factors, [2]int{i, e})
			}
		}
		if n > 1 { // n 是质数
			factors = append(factors, [2]int{n, 1})
		}
		return
	}
	primeDivisors := func(n int) (primes []int) {
		for i := 2; i*i <= n; i++ {
			k := 0
			for ; n%i == 0; n /= i {
				k++
			}
			if k > 0 {
				primes = append(primes, i)
			}
		}
		if n > 1 { // n 是质数
			primes = append(primes, n)
		}
		return
	}

	// 阶乘的质因数分解 Finding Power of Factorial Divisor
	// 见进阶指南 p.138
	// https://cp-algorithms.com/algebra/factorial-divisors.html

	// 预处理: [2,mx] 的质因数分解的系数和 bigomega(n) or Omega(n) https://oeis.org/A001222
	// a(n) depends only on prime signature of n (cf. https://oeis.org/A025487)
	// So a(24) = a(375) since 24 = 2^3 * 3 and 375 = 3 * 5^3 both have prime signature (3, 1)
	//
	// 		Omega(n) - omega(n) https://oeis.org/A046660
	//
	// 另一种写法 https://math.stackexchange.com/questions/1955105/corectness-of-prime-factorization-over-a-range
	// 性质：Omega(nm)=Omega(n)+Omega(m)
	primeExponentsCountAll := func() {
		const mx int = 1e6
		cnts := [mx + 1]int{}
		primes := []int{}
		for i := 2; i <= mx; i++ {
			if cnts[i] == 0 {
				cnts[i] = 1
				primes = append(primes, i)
			}
			for _, p := range primes {
				if p*i > mx {
					break
				}
				cnts[p*i] = cnts[i] + 1
			}
		}

		// EXTRA: 前缀和，即 Omega(n!) https://oeis.org/A022559
		for i := 3; i <= mx; i++ {
			cnts[i] += cnts[i-1]
		}
	}
	//primeExponentsCount := func(n int) (count []int) {
	//	count = make([]int, n+1)
	//	left := make([]int, n+1)
	//	for i := range left {
	//		left[i] = i
	//	}
	//	i := 2
	//	for ; i*i <= n; i++ {
	//		if count[i] == 0 {
	//			for j := i; j <= n; j += i {
	//				count[j]++
	//				left[j] /= i
	//			}
	//		} else if left[i] > 1 { // i is non square-free
	//			count[i] += count[left[i]]
	//		}
	//	}
	//	for ; i <= n; i++ {
	//		if count[i] == 0 {
	//			count[i] = 1
	//		} else if left[i] > 1 { // i is non square-free
	//			count[i] += count[left[i]]
	//		}
	//	}
	//	return
	//}

	/* 因子/因数/约数

	n 的因子个数 d(n) = Π(ei+1), ei 为第 i 个质数的系数 https://oeis.org/A000005 d(n) 也写作 τ(n)
		Positions of records (高合成数，反素数) https://oeis.org/A002182
		Values of records https://oeis.org/A002183
		相关题目：范围内的最多约数个数 https://www.luogu.com.cn/problem/P1221

		d(n) 前缀和 = Σ{k=1..n} floor(n/k) https://oeis.org/A006218
	               = 见后文「数论分块/除法分块」

		n+d(n) https://oeis.org/A062249
		n-d(n) https://oeis.org/A049820   count https://oeis.org/A060990   前缀和 https://oeis.org/A161664
		n*d(n) https://oeis.org/A038040   前缀和 https://oeis.org/A143127 = Sum_{i=1..floor(√n)}i*(i+floor(n/i))*(floor(n/i)+1-i) - 平方和(floor(√n))
												https://atcoder.jp/contests/abc172/tasks/abc172_d
		d(n)|n https://oeis.org/A033950 refactorable numbers / tau numbers
			n/d(n) https://oeis.org/A036762
		n%d(n) https://oeis.org/A054008
		a(1)=1, a(n+1)=a(n)+d(a(n)) https://oeis.org/A064491
		Smallest number with d(x) = n https://oeis.org/A005179
			a(p) = 2^(p-1) for primes p
			相关题目 https://codeforces.com/problemset/problem/27/E https://www.luogu.com.cn/problem/P1128
			质数的情况 https://oeis.org/A061286

	n 的因子之和 σ(n) = Π(pi^(ei+1)-1)/(pi-1) https://oeis.org/A000203 todo 相关题目 https://www.luogu.com.cn/problem/P1593
		σ(n) 前缀和 = Σ{k=1..n} k*floor(n/k) https://oeis.org/A024916
		真因子之和 https://oeis.org/A001065 σ(n)-n
		完全数/完美数/完备数 https://oeis.org/A000396 Perfect numbers (σ(n) = 2n)
			https://en.wikipedia.org/wiki/Perfect_number
		过剩数/丰数/盈数 https://oeis.org/A005101 Abundant numbers (σ(n) > 2n)
			https://en.wikipedia.org/wiki/Abundant_number
		亏数/缺数/不足数 https://oeis.org/A005100 Deficient numbers (σ(n) < 2n)
			https://en.wikipedia.org/wiki/Deficient_number

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

	d(10^n) https://oeis.org/A066150
	方便估计复杂度 - 近似为开立方
	4, 12, 32, 64, 128, 240, 448, 768, 1344, /9/
	2304, 4032, 6720, 10752, 17280, 26880, 41472, 64512, 103680, 161280 /19/

		上面这些数对应的最小的 n https://oeis.org/A066151
		6, 60, 840, 7560, 83160, 720720, 8648640, 73513440, 735134400,
		6983776800, 97772875200, 963761198400, 9316358251200, 97821761637600, 866421317361600, 8086598962041600, 74801040398884800, 897612484786617600

	d(d(...d(n))) 迭代至 2 所需要的迭代次数
	0,0,1,0,2,0,2,1,2,0,3,0,2,2,1,0,3,0,3,2,2,0,3,1,2,2,3

	高合成数/反质数 Highly Composite Numbers https://oeis.org/A002182
	https://oi-wiki.org/math/prime/#_7
	性质：一个高合成数一定是由另一个高合成数乘一个质数得到
	见进阶指南 p.140-141
	Number of divisors of n-th highly composite number https://oeis.org/A002183
	Number of highly composite numbers not divisible by n https://oeis.org/A199337
	求出不超过 n 的最大的反质数 https://www.luogu.com.cn/problem/P1463

	Largest divisor of n having the form 2^i*5^j http://oeis.org/A132741
	a(n) = A006519(n)*A060904(n) = 2^A007814(n)*5^A112765(n)

	Squarefree numbers https://oeis.org/A005117 (介绍了一种筛法)
	Numbers that are not divisible by a square greater than 1
	Lim_{n->infinity} a(n)/n = Pi^2/6

		Numbers that are not squarefree https://oeis.org/A013929
		Numbers that are divisible by a square greater than 1

	Squarefree part of n (also called core(n)) https://oeis.org/A007913
	a(n) is the smallest positive number m such that n/m is a square

	Largest squarefree number dividing n https://oeis.org/A007947
	the squarefree kernel of n, rad(n), radical of n

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

	divisorPairs := func(n int64) (ds [][2]int64) {
		for d := int64(1); d*d <= n; d++ {
			if n%d == 0 {
				ds = append(ds, [2]int64{d, n / d})
			}
		}
		return
	}

	doDivisors := func(n int, do func(d int)) {
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				do(d)
				if d*d < n {
					do(n / d)
				}
			}
		}
	}
	doDivisors2 := func(n int, do func(d1, d2 int)) {
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				do(d, n/d)
			}
		}
	}

	// Number of odd divisors of n https://oeis.org/A001227
	// a(n) = d(2*n) - d(n)
	// 亦为整数 n 分拆成若干连续整数的方法数
	// Number of partitions of n into consecutive positive integers including the trivial partition of length 1
	// e.g. 9 = 2+3+4 or 4+5 or 9 so a(9)=3
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
			// 统计因子个数 d(n)
			// NOTE: 复杂度可以做到线性 https://codeforces.com/contest/920/submission/76859782
			// 相关 OEIS 见上面的注释块
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

		isSquareNumber := func(x int) bool { return len(divisors[x])&1 == 1 }
		halfDivisors := func(x int) []int { d := divisors[x]; return d[:(len(d)-1)/2+1] }

		_, _ = isSquareNumber, halfDivisors
	}

	// LPF(n): least prime dividing n (when n > 1); a(1) = 1 https://oeis.org/A020639
	// 有时候数据范围比较大，用 primeFactorsAll 预处理会 MLE，这时候就要用 LPF 了（同样是预处理但是内存占用低）
	// 先预处理出 LPF，然后对要处理的数 v 不断地除 LPF(v) 直到等于 1
	// 		LPF 前缀和 https://oeis.org/A046669 前缀积 https://oeis.org/A072486
	//		n+LPF(n) https://oeis.org/A061228 the smallest number greater than n which is not coprime to n
	// 		n-LPF(n) https://oeis.org/A046666
	//		n*LPF(n) https://oeis.org/A285109
	// 		n/LPF(n) https://oeis.org/A032742 即 n 的最大因子 = Max{gcd(n,j); j=n+1..2n-1}
	//
	//		只考虑奇质数 http://oeis.org/A078701
	//
	// GPF(n): greatest prime dividing n, for n >= 2; a(1)=1 https://oeis.org/A006530
	//		GPF(p-1) http://oeis.org/A023503
	//		GPF(p+1) http://oeis.org/A023509
	// 		GPF 前缀和 https://oeis.org/A046670 前缀积 https://oeis.org/A104350
	//		n+GPF(n) https://oeis.org/A070229 the next m>n such that GPF(n)|m
	// 		n-GPF(n) https://oeis.org/A076563
	// 		n*GPF(n) https://oeis.org/A253560
	// 		n/GPF(n) https://oeis.org/A052126
	//      a(1)=1, a(n+1)=a(n)+GPF(a(n)) https://oeis.org/A076271
	//
	// 		n/LPF(n)*GPF(n) https://oeis.org/A130064
	// 		n/GPF(n)*LPF(n) https://oeis.org/A130065
	//
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

		// EXTRA: 分解 v
		// lpf[v]==p 也可以写成 v%p==0
		var v int
		for v > 1 {
			p := lpf[v]
			e := 1
			for v /= p; lpf[v] == p; v /= p {
				e++
			}
			// do(p,e)
		}
	}

	// 预处理: [2,mx] 范围内数的不同质因子，例如 factors[12] = [2,3]
	// for i>=2, factors[i][0] == i means i is prime
	primeFactorsAll := func() {
		const mx int = 1e6
		factors := [mx + 1][]int{}
		for i := 2; i <= mx; i++ {
			if len(factors[i]) == 0 {
				for j := i; j <= mx; j += i {
					factors[j] = append(factors[j], i)
				}
			}
		}
	}

	// 预处理: [2,mx] 的不同的质因子个数 omega(n) https://oeis.org/A001221
	distinctPrimesCountAll := func() {
		const mx int = 1e6
		cnts := make([]int, mx+1)
		for i := 2; i <= mx; i++ {
			if cnts[i] == 0 {
				for j := i; j <= mx; j += i {
					cnts[i]++
				}
			}
		}

		// EXTRA: 前缀和，即 omega(n!) https://oeis.org/A013939
		for i := 3; i <= mx; i++ {
			cnts[i] += cnts[i-1]
		}
	}

	// 计算单个数 n 的欧拉函数（互质的数的个数）Euler totient function
	calcPhi := func(n int) int {
		ans := n
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				ans = ans / i * (i - 1)
				for ; n%i == 0; n /= i {
				}
			}
		}
		if n > 1 {
			ans = ans / n * (n - 1)
		}
		return ans
	}

	// 欧拉函数（互质的数的个数）Euler totient function https://oeis.org/A000010
	// https://en.wikipedia.org/wiki/Euler%27s_totient_function
	// 预处理 [1,mx] 欧拉函数
	// NOTE: φ(φ...(n)) 收敛到 1 的迭代次数是 log 级别的：奇数减一，偶数减半 https://oeis.org/A003434
	// GCD(n, φ(n)) https://oeis.org/A009195
	// n+φ(n) http://oeis.org/A121048
	// n-φ(n) https://oeis.org/A051953 called Cototient
	// n*φ(n) https://oeis.org/A002618 = φ(n^2)
	// φ(n)|n https://oeis.org/A007694 iff n = 1 or n = 2^w * 3^u for w > 0 and u >= 0
	// 		n/φ(n) https://oeis.org/A049237
	//      the quotients can take only 3 distinct values:
	//			n/φ(n) = 1 iff n = 1
	//			n/φ(n) = 2 iff n = 2^w, w >= 1
	//			n/φ(n) = 3 iff n = 2^w * 3^u, w >= 1, u >= 1
	// n%φ(n) https://oeis.org/A068494
	// a(1)=1, a(n+1)=a(n)+φ(a(n)) https://oeis.org/A074693
	// Least number k such that phi(k) = n https://oeis.org/A002181    Inverse of Euler totient function
	// Number of values of k such that phi(k) = n https://oeis.org/A058277
	// φ集合 https://oeis.org/A002202
	// φ补集 https://oeis.org/A007617
	initPhi := func() {
		const mx int = 1e6
		phi := [mx + 1]int{1: 1}
		for i := 2; i <= mx; i++ {
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

	// 扩展欧拉定理（降幂公式）
	// https://oi-wiki.org/math/fermat/#_5
	// https://zhuanlan.zhihu.com/p/42632291
	// https://blog.csdn.net/synapse7/article/details/19610361
	// a^b ≡ a^(b mod φ(m)) (mod m), gcd(a,m)=1
	// a^b ≡ a^(b mod φ(m) + φ(m)) (mod m), gcd(a,m)!=1 且 b>φ(m)
	// 模板题 https://www.luogu.com.cn/problem/P5091
	// 例题 https://codeforces.com/problemset/problem/615/D
	// https://cses.fi/problemset/task/1712
	exPhi := func(a, b int64, m int) int64 {
		phi := int64(calcPhi(m))
		if b >= phi {
			b = b%phi + phi
		}
		return pow(a, b, int64(m))
	}

	// phi 求和相关
	// ∑φ(i) https://oeis.org/A002088 #{(x,y): 1<=x<=y<=n, gcd(x,y)=1}
	// 近似为 n^2 / 3   相关题目 https://codeforces.com/problemset/problem/1009/D
	// 1, 2, 4, 6, 10, 12, 18, 22, 28, 32, 42, 46, 58, 64, 72, 80, 96, 102
	// ∑φ(i)-1 http://oeis.org/A015614 #{(x,y): 1<=x<y<=n, gcd(x,y)=1}
	// 0, 1, 3, 5, 9, 11, 17, 21, 27, 31, 41, 45, 57, 63, 71, 79, 95, 101
	// todo https://oi-wiki.org/math/min-25/#_7

	// Number of numbers "unrelated to n" http://oeis.org/A045763
	// m < n such that m is neither a divisor of n nor relatively prime to n
	// a(n) = n + 1 - d(n) - phi(n); where d(n) is the number of divisors of n

	// Unitary totient (or unitary phi) function uphi(n) http://oeis.org/A047994

	/* 同余 逆元
	http://oeis.org/A006254 a(n) = 2^-1 mod p         Numbers n such that 2n-1 is prime
	http://oeis.org/A283419 3^-1 mod p
	http://oeis.org/A006093 (p-1)^-1 mod p     prime(n) - 1
	http://oeis.org/A040976 (p/2)^-1 mod p     prime(n) - 2
	http://oeis.org/A090938 Least multiple of n == 1 (mod prime(n))
	http://oeis.org/A090939 Least multiple of n == -1 (mod prime(n))
	http://oeis.org/A091185 a(n) = A090938(n)/n      n^-1 mod prime(n)
	*/

	// 二元一次不定方程
	// exgcd solve equation ax+by=gcd(a,b)
	// we have |x|<=b and |y|<=a in result (x,y)
	// https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
	// 模板题 https://www.luogu.com.cn/problem/P5656
	// 使非负解 x+y 尽量小 https://codeforces.com/problemset/problem/1244/C
	var exgcd func(a, b int64) (gcd, x, y int64)
	exgcd = func(a, b int64) (gcd, x, y int64) {
		if b == 0 {
			return a, 1, 0
		}
		gcd, y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}

	// 任意非零模数逆元
	// ax ≡ 1 (mod m)
	// 模板题 https://www.luogu.com.cn/problem/P1082
	invM := func(a, m int64) int64 { _, x, _ := exgcd(a, m); return (x%m + m) % m }

	// 费马小定理求质数逆元
	// ax ≡ 1 (mod p)
	// x^-1 ≡ a^(p-2) (mod p)
	invP := func(a, p int64) int64 { return pow(a, p-2, p) }

	// 有理数取模
	// 模板题 https://www.luogu.com.cn/problem/P2613
	divM := func(a, b, m int64) int64 { return a * invM(b, m) % m }
	divP := func(a, b, p int64) int64 { return a * invP(b, p) % p }

	// 线性求逆元·其一
	// 求 1, 2, ..., p−1 mod p 的逆元
	// http://blog.miskcoo.com/2014/09/linear-find-all-invert
	// https://www.zhihu.com/question/59033693
	// 模板题 https://www.luogu.com.cn/problem/P3811
	initAllInv := func(p int) []int64 {
		inv := make([]int64, p)
		inv[1] = 1
		for i := 2; i < p; i++ {
			inv[i] = int64(p-p/i) * inv[p%i] % int64(p)
		}
		return inv
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
	// https://blog.csdn.net/synapse7/article/details/9946013
	// todo https://codeforces.com/blog/entry/61290
	// 模板题 https://www.luogu.com.cn/problem/P1495
	crt := func(a, m []int64) (x int64) {
		M := int64(1)
		for _, mi := range m {
			M *= mi
		}
		for i, mi := range m {
			Mi := M / mi
			_, inv, _ := exgcd(Mi, mi)
			x = (x + a[i]*Mi*inv) % M
		}
		x = (x + M) % M
		return
	}

	// 扩展中国剩余定理 (EXCRT)
	// ai * x ≡ bi (mod mi)
	// 解为 x ≡ b (mod m)
	// 有解时返回 (b, m)，无解时返回 (0, -1)
	// 推导过程见《挑战程序设计竞赛》P292
	// 注意乘法溢出的可能
	// 推荐 https://blog.csdn.net/niiick/article/details/80229217
	// 模板题 https://www.luogu.com.cn/problemnew/solution/P4777
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

	// 离散对数
	// a^x ≡ b (mod p)，a 和 p 互质 - 小步大步算法 (BSGS)
	// 时间复杂度 O(√p)
	// 见进阶指南 p.155
	// 扩展大步小步法解决离散对数问题 http://blog.miskcoo.com/2015/05/discrete-logarithm-problem
	babyStepGiantStep := func(a, b, p int64) int64 {
		hash := map[int64]int64{}
		b %= p
		t := int64(math.Sqrt(float64(p))) + 1
		for j := int64(0); j < t; j++ {
			v := b * pow(a, j, p) % p
			hash[v] = j
		}
		a = pow(a, t, p)
		if a == 0 {
			if b == 0 {
				return 1
			}
			return -1
		}
		for i := int64(0); i < t; i++ {
			v := pow(a, i, p)
			if j, ok := hash[v]; ok && i*t >= j {
				return i*t - j
			}
		}
		return -1
	}

	// 原根 / N次剩余 / 高次同余方程 x^a ≡ b (mod p)
	// todo
	// 模板题 https://www.luogu.com.cn/problem/P5491 https://www.luogu.com.cn/problem/P5668

	// Number of solutions to x^n ≡ 1 (mod n), 1<=x<=n https://oeis.org/A072994
	// Least k > 0 such that the number of solutions to x^k == 1 (mod k) 1 <= x <= k is equal to n, or 0 if no such k exists https://oeis.org/A072995

	// http://oeis.org/A182865 Minimal number of quadratic residues
	// a(n) is the least integer m such that any nonzero square is congruent (mod n) to one of the squares from 1 to m^2
	// 把这题的 1000 改成 i，则至多需要枚举到 a(i) https://ac.nowcoder.com/acm/contest/6489/A

	// https://oeis.org/A000224 Number of distinct squares residues mod n
	// Multiplicative with a(p^e) = floor(p^e/6) + 2 IF p = 2 ELSE floor(p^(e+1)/(2p + 2)) + 1
	// https://oeis.org/A046530 Number of distinct cubic residues mod n

	/* 阶乘 组合数/二项式系数 */

	// https://oeis.org/A000142
	factorial := [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800 /*10!*/, 39916800, 479001600}

	// https://oeis.org/A067850 Highest power of 2 not exceeding n!

	calcFactorial := func(n int) int64 {
		res := int64(1) % mod
		for i := 2; i <= n; i++ {
			res = res * int64(i) % mod
		}
		return res
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

	// EXTRA: binomial(n, floor(n/2)) https://oeis.org/A001405
	// Sperner's theorem says that this is the maximal number of subsets of an n-set such that no one contains another
	combHalf := [...]int64{
		1, 1, 2, 3, 6, 10, 20, 35, 70, 126, // C(9,4)
		252, 462, 924, 1716, 3432, 6435, 12870, 24310, 48620, 92378, // C(19,9)
		184756, 352716, 705432, 1352078, 2704156, 5200300, 10400600, 20058300, 40116600, 77558760, // C(29,14)
		155117520, 300540195, 601080390, 1166803110, 2333606220, 4537567650, 9075135300, 17672631900, 35345263800, 68923264410, // C(39,19)
		137846528820, 269128937220, 538257874440, 1052049481860, 2104098963720, 4116715363800, 8233430727600, 16123801841550, 32247603683100, 63205303218876, // C(49,24)
		126410606437752, 247959266474052, 495918532948104, 973469712824056, 1946939425648112, 3824345300380220, 7648690600760440, 15033633249770520, 30067266499541040, 59132290782430712, // C(59,29)
		118264581564861424, 232714176627630544, 465428353255261088, 916312070471295267, 1832624140942590534, 3609714217008132870, 7219428434016265740, // C(66,33)
	}

	// EXTRA: Central binomial coefficients: binomial(2*n,n) = (2*n)!/(n!)^2 https://oeis.org/A000984

	// 求组合数/二项式系数
	// 不取模，仅适用于小范围的 n 和 k
	// 更大范围的见线性求逆元
	comb := func(n, k int) int64 {
		if k > n-k {
			k = n - k
		}
		res := int64(1)
		for i := 1; i <= k; i++ {
			res = res * int64(n-k+i) / int64(i)
		}
		return res
		//return big.Int{}.Binomial(n, k).Int64()
	}

	// 取模，适用于 n 较大但 k 或 n-k 较小的情况
	comb = func(n, k int) int64 {
		if k > n-k {
			k = n - k
		}
		a, b := int64(1), int64(1)
		for i := 1; i <= k; i++ {
			a = a * int64(n) % mod
			n--
			b = b * int64(i) % mod
		}
		return divP(a, b, mod)
	}

	{
		// 初始化组合数
		// 不取模，仅适用于小范围的 n 和 k
		const mx = 60
		C := [mx + 1][mx + 1]int64{}
		for i := 0; i <= mx; i++ {
			C[i][0], C[i][i] = 1, 1
			for j := 1; j < i; j++ {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}

	// 不推荐，见后面的代码块
	//{
	//	// O(n) 预处理，O(logn) 求组合数
	//	const mod int64 = 1e9 + 7
	//	const mx = 3e5
	//	F := [mx + 1]int64{1}
	//	for i := 1; i <= mx; i++ {
	//		F[i] = F[i-1] * int64(i) % mod
	//	}
	//	pow := func(x, n int64) int64 {
	//		x %= mod
	//		res := int64(1)
	//		for ; n > 0; n >>= 1 {
	//			if n&1 == 1 {
	//				res = res * x % mod
	//			}
	//			x = x * x % mod
	//		}
	//		return res
	//	}
	//	inv := func(a int64) int64 { return pow(a, mod-2) }
	//	div := func(a, b int64) int64 { return a * inv(b) % mod }
	//	C := func(n, k int64) int64 { return div(F[n], F[k]*F[n-k]%mod) }
	//
	//	_ = C
	//}

	{
		// O(n) 预处理阶乘及其逆元，O(1) 求组合数
		const mod int64 = 1e9 + 7
		const mx int = 1e6
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

		// EXTRA: 卢卡斯定理
		var lucas func(n, k int64) int64
		lucas = func(n, k int64) int64 {
			if k == 0 {
				return 1
			}
			return C(int(n%mod), int(k%mod)) * lucas(n/mod, k/mod) % mod
		}

		// 可重组合 https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition
		// 方案数 H(n,k)=C(n+k-1,k) https://oeis.org/A059481
		// 相当于把 k 个无区别的球放入 n 个有区别的盒子中，且允许空盒的方案数
		//		隔板法：把 n 个盒子当做 n-1 个隔板，这样相当于总共有 k+n-1个位置，从中选择 k 个位置放球，剩下的位置放隔板。这样就把 k 个球划分成了 n 份，放入对应的盒子中
		// 相当于长度为 k，元素范围在 [1,n] 的非降序列的个数
		H := func(n, k int) int64 { return C(n+k-1, k) }

		pow2 := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			pow2[i] = pow2[i-1] << 1 % mod
		}

		_, _ = C, H
	}

	// 扩展卢卡斯
	// todo https://blog.csdn.net/niiick/article/details/81064156
	// https://blog.csdn.net/skywalkert/article/details/52553048
	// https://blog.csdn.net/skywalkert/article/details/104681101
	// https://cp-algorithms.com/combinatorics/binomial-coefficients.html
	// 模板题 https://www.luogu.com.cn/problem/P4720
	// 古代猪文 https://www.luogu.com.cn/problem/P2480

	// 原根
	// https://oeis.org/A033948 Numbers that have a primitive root (the multiplicative group modulo n is cyclic)
	//     The sequence consists of 1, 2, 4 and numbers of the form p^i and 2p^i, where p is an odd prime and i > 0
	// https://oeis.org/A046145 Smallest primitive root modulo n, or 0 if no root exists
	// https://oeis.org/A001918 Smallest primitive root of n-th prime
	// https://oeis.org/A046146 Largest primitive root (<n) modulo n, or 0 if no root exists
	// https://oeis.org/A071894 Largest primitive root (<p) of n-th prime p
	// https://oeis.org/A046144 Number of primitive roots modulo n
	// https://oeis.org/A008330 Number of primitive roots of n-th prime = phi(p-1)
	// https://oeis.org/A056619 Smallest prime with primitive root of n or 0 if no such prime exists
	// https://oeis.org/A023049 Smallest prime > n having primitive root n, or 0 if no such prime exists
	// todo https://cp-algorithms.com/algebra/primitive-root.html
	// todo 模板题 https://www.luogu.com.cn/problem/P6091

	// 莫比乌斯函数 mu https://oeis.org/A008683
	// 基于线性筛方法
	// todo https://oi-wiki.org/math/mobius/#_11
	// 前缀和 https://oi-wiki.org/math/min-25/#_6
	muInit := func() {
		const mx int = 1e6
		mu := [mx + 1]int{1: 1}
		primes := []int{}
		vis := [mx + 1]bool{}
		for i := 2; i <= mx; i++ {
			if !vis[i] {
				mu[i] = -1
				primes = append(primes, i)
			}
			for _, p := range primes {
				if p*i > mx {
					break
				}
				vis[p*i] = true
				if i%p == 0 {
					mu[p*i] = 0
					break
				}
				mu[p*i] = -mu[i]
			}
		}
	}

	// 莫比乌斯反演（岛娘推荐！https://zhuanlan.zhihu.com/p/133761303）
	// todo https://oi-wiki.org/math/mobius/

	//

	// 数论分块/除法分块
	// a(n) = Σ{k=1..n} floor(n/k) https://oeis.org/A006218
	//      = 2*( Σ{i=1..floor(sqrt(n))} floor(n/i) ) - floor(sqrt(n))^2
	// thus, a(n) % 2 == floor(sqrt(n)) % 2

	// 杜教筛 - 积性函数前缀和
	// todo 推荐 https://blog.csdn.net/weixin_43914593/article/details/104229700 算法竞赛专题解析（4）：杜教筛--以及积性函数的前世今生
	// todo http://baihacker.github.io/main/
	// The prefix-sum of multiplicative function: the black algorithm http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_the_black_algorithm.html
	// The prefix-sum of multiplicative function: Dirichlet convolution http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_dirichlet_convolution.html
	// The prefix-sum of multiplicative function: powerful number sieve http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_powerful_number_sieve.html
	// 浅谈一类积性函数的前缀和 + 套题 https://blog.csdn.net/skywalkert/article/details/50500009
	// 模板题 https://www.luogu.com.cn/problem/P4213

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

	/* 斐波那契数列 F(n) http://oeis.org/A000045
	http://oeis.org/A000071 F(n) 前缀和 = F(n)-1
	- 异或和 F(n) 1,0,2,1,4,12,1,20,54,1,88,200,33,344,826,225,1756,3268,7313,1788
	http://oeis.org/A007598 F^2(n)    a(n) = 2*a(n-1) + 2*a(n-2) - a(n-3), n > 2. a(0)=0, a(1)=1, a(2)=1
	                                  a(n) = (F(n)*F(n+4)-3*F(n)*F(n+1))/2
	http://oeis.org/A001690 补集
	http://oeis.org/A022307 F(n) 的不同的质因子个数
	http://oeis.org/A001175 N(m) = F%m 的周期    Pisano periods / Pisano numbers
	                        N(m) = LCM(N(p1^e1), ..., N(pk^ek))
	http://oeis.org/A060305 N(p) = F%p 的周期
	http://oeis.org/A003893 F(n)%10
	http://oeis.org/A001605 使 F(n) 为质数的 n

	其他相关序列
	http://oeis.org/A000213 Tribonacci numbers: a(n)=a(n-1)+a(n-2)+a(n-3) with a(0)=a(1)=a(2)=1
	http://oeis.org/A000931 Padovan sequence (or Padovan numbers): a(n)=a(n-2)+a(n-3) with a(0)=1, a(1)=a(2)=0
	*/

	// https://oeis.org/A195264 Iterate x -> A080670(x) (replace x with the concatenation of the primes and exponents in its prime factorization)
	// starting at n until reach 1 or a prime; or -1 if a prime is never reached
	// https://www.zhihu.com/question/48612677/answer/487252829

	_ = []interface{}{
		primes, primes10, primes10_,
		sqCheck, cubeCheck, sqrt, cbrt, bottomDiff,
		gcd, gcdPrefix, gcdSuffix, lcm, frac, cntRangeGCD,
		isPrime, sieve, sieveL, primeFactorization, primeDivisors, primeExponentsCountAll,
		divisors, divisorPairs, doDivisors, doDivisors2, oddDivisorsNum, maxSqrtDivisor, divisorsAll, primeFactorsAll, lpfAll, distinctPrimesCountAll,
		calcPhi, initPhi, exPhi,
		exgcd, invM, invP, divM, divP, initAllInv, calcAllInv,
		crt, excrt,
		babyStepGiantStep,
		factorial, calcFactorial, initFactorial, _factorial, combHalf, comb,
		muInit,
	}
}

/* 组合数学
https://en.wikipedia.org/wiki/Combinatorics
https://en.wikipedia.org/wiki/Enumerative_combinatorics

NOTE: 涉及到相邻的组合问题：可以考虑当前位置和左侧位置所满足的性质（例题 https://atcoder.jp/contests/abc167/tasks/abc167_e）

一些常用组合恒等式的解释 https://www.zhihu.com/question/26094736
       C(n, k-1) + C(n, k) = C(n+1, k)
       C(r, r) + C(r+1, r) + ... + C(n, r) = C(n+1, r+1)
上式亦为 C(n, 0) + C(n+1, 1) + ... + C(n+m, m) = C(n+m+1, m)（例题 https://atcoder.jp/contests/abc154/tasks/abc154_f）

隔三组合数 https://oeis.org/A024493 https://oeis.org/A024494 https://oeis.org/A024495 C(n,0) + C(n,3) + ... + C(n,3[n/3])
隔四组合数 https://oeis.org/A038503 https://oeis.org/A038504 https://oeis.org/A038505 https://oeis.org/A000749

二项式定理 https://en.wikipedia.org/wiki/Binomial_theorem

Tetrahedral (or triangular pyramidal) numbers: a(n) = C(n+2,3) = n*(n+1)*(n+2)/6 http://oeis.org/A000292
a(n) = Sum_{1<=i<=j<=n} j-i
a(n) = sum of all the possible products p*q where (p,q) are ordered pairs and p + q = n + 1
a(n) = 长度为 n 的字符串的所有子串长度之和

隔板法 https://zh.wikipedia.org/wiki/%E9%9A%94%E6%9D%BF%E6%B3%95
放球问题（总结得不错）https://baike.baidu.com/item/%E6%94%BE%E7%90%83%E9%97%AE%E9%A2%98
	扩展例题 https://codeforces.com/problemset/problem/893/E
圆排列 https://zh.wikipedia.org/wiki/%E5%9C%86%E6%8E%92%E5%88%97
可重集排列
可重集组合 todo https://codeforces.com/problemset/problem/451/E
https://oeis.org/A000166 错排 subfactorial numbers  a[n]=(n-1)*(a[n-1]+a[n-2])  https://zh.wikipedia.org/wiki/%E9%94%99%E6%8E%92%E9%97%AE%E9%A2%98
	https://oeis.org/A082491 n! * A000166(n)   a(n+2) = (n+2)*(n+1)*(a(n+1)+(n+1)*a(n))
	https://oeis.org/A000255 错排的比较对象的范围是 [1,n+1]  a(n) = n*a(n-1) + (n-1)*a(n-2), a(0) = a(1) = 1
	https://oeis.org/A000153 错排的比较对象的范围是 [1,n+2]  a(n) = n*a(n-1) + (n-2)*a(n-2), a(0) = 0, a(1) = 1
	https://oeis.org/A000261 错排的比较对象的范围是 [1,n+3]  a(n) = n*a(n-1) + (n-3)*a(n-2), a(1) = 0, a(2) = 1
	https://oeis.org/A001909 错排的比较对象的范围是 [1,n+4]  a(n) = n*a(n-1) + (n-4)*a(n-2), a(2) = 0, a(3) = 1
		https://atcoder.jp/contests/abc172/tasks/abc172_e
范德蒙德恒等式 https://zh.wikipedia.org/wiki/%E8%8C%83%E5%BE%B7%E8%92%99%E6%81%92%E7%AD%89%E5%BC%8F
斯特林数 https://blog.csdn.net/ACdreamers/article/details/8521134
Stirling numbers of the first kind, s(n,k) https://oeis.org/A008275
   将 n 个元素排成 k 个非空循环排列的方法数
   s(n,k) 的递推公式： s(n,k)=(n-1)*s(n-1,k)+s(n-1,k-1), 1<=k<=n-1
   边界条件：s(n,0)=0, n>=1    s(n,n)=1, n>=0
Stirling numbers of the second kind, S2(n,k) https://oeis.org/A008277
   将 n 个元素拆分为 k 个非空集的方法数
   S2(n, k) = (1/k!) * Σ{i=0..k} (-1)^(k-i)*binomial(k, i)*i^n.
   S2(n,k) 的递推公式：S2(n,k)=k*S2(n-1,k)+S2(n-1,k-1), 1<=k<=n-1
   边界条件：S(n,0)=0, n>=1    S(n,n)=1, n>=0

二阶递推数列通项 https://zhuanlan.zhihu.com/p/75096951
凯莱公式 Cayley’s formula: the number of trees on n labeled vertices is n^(n-2).
普吕弗序列 Prüfer sequence: 由树唯一地产生的序列
约瑟夫问题 Josephus Problem https://cp-algorithms.com/others/josephus_problem.html https://en.wikipedia.org/wiki/Josephus_problem
Stern-Brocot 树与 Farey 序列 https://oi-wiki.org/misc/stern-brocot/ https://cp-algorithms.com/others/stern_brocot_tree_farey_sequences.html

* 生成函数/母函数 *
https://en.wikipedia.org/wiki/Generating_function
整数分拆 https://oeis.org/A000041 https://en.wikipedia.org/wiki/Partition_(number_theory)
https://oeis.org/A104513 The number of consecutive integers > 1 beginning with A104512(n), the sum of which equals n, or 0 if impossible.
						a(n)=0 iff n=2^k
https://oeis.org/A069283 将 n 分拆成至少两个连续整数的方法数 = n 的奇因子数 - 1
						见上面的 oddDivisorsNum 函数
https://oeis.org/A018819 Binary partition function: number of partitions of n into powers of 2
	相关题目 https://www.luogu.com.cn/problem/P6065 http://poj.org/problem?id=2229

	质数分拆
	https://oeis.org/A061358 Number of ways of writing n=p+q with p, q primes and p>=q
	https://oeis.org/A067187 Numbers that can be expressed as the sum of two primes in exactly one way
	https://oeis.org/A068307 number of partitions of n into a sum of three primes
	https://oeis.org/A071335 Number of partitions of n into a sum of at most three primes

Maximum product of two integers whose sum is n https://oeis.org/A002620
Quarter-squares: floor(n/2)*ceiling(n/2). Equivalently, floor(n^2/4)

	Maximal product of three numbers with sum n: a(n) = max(r*s*t), n = r+s+t https://oeis.org/A006501
	a(n) = floor(n/3)*floor((n+1)/3)*floor((n+2)/3)
	Expansion of (1+x^2) / ( (1-x)^2 * (1-x^3)^2 )

	Maximal product of four nonnegative integers whose sum is n https://oeis.org/A008233
	a(n) = floor(n/4)*floor((n+1)/4)*floor((n+2)/4)*floor((n+3)/4)

	...

	相关题目 https://codeforces.com/problemset/problem/1368/B

记 A = [1,2,...,n]，A 的全排列中与 A 的最大差值为 n^2/2 https://oeis.org/A007590
Maximum sum of displacements of elements in a permutation of (1..n)
For example, with n = 9, permutation (5,6,7,8,9,1,2,3,4) has displacements (4,4,4,4,4,5,5,5,5) with maximal sum = 40

n married couples are seated in a row so that every wife is to the left of her husband
若不考虑顺序，则所有排列的个数为 (2n)!
考虑顺序可以发现，对于每一对夫妻来说，妻子在丈夫左侧的情况和在右侧的情况相同且不同对夫妻之间是独立的
因此每有一对夫妻，符合条件的排列个数就减半
所以结果为 a(n) = (2n)!/2^n https://oeis.org/A000680
或者见这道题目的背景 LC1359 https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/

NxN 大小的对称置换矩阵的个数 http://oeis.org/A000085
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
相关题目 LC1411/周赛184D https://leetcode-cn.com/problems/number-of-ways-to-paint-n-x-3-grid/ https://leetcode-cn.com/contest/weekly-contest-184/

男厕问题 / 电话问题 https://oeis.org/A185456
Assume that the first person to use a bank of payphones selects one at the end,
and all subsequent users select the phone which puts them farthest from the current phone users.
U(n) is the smallest number of phones such that n may be used without any two adjacent phones being used
https://www.zhihu.com/question/278361000/answer/1004606685

一些二进制的计数问题见 bits.go

CF 上的一些组合计数问题 http://blog.miskcoo.com/2015/06/codeforces-combinatorics-and-probabilities-problem
*/
func combinatoricsCollection() {
	// 容斥原理 (PIE, the principle of inclusion and exclusion)
	// 参考《挑战程序设计竞赛》P296
	// https://codeforces.com/blog/entry/64625
	// https://ac.nowcoder.com/acm/contest/6219/C
	solveInclusionExclusion := func(a []int) (ans int64) {
		n := len(a)
		const mod int64 = 1e9 + 7 // 998244353
		for sub := uint(0); sub < 1<<n; sub++ {
			res := int64(0)
			for i, v := range a {
				if sub>>i&1 == 1 {
					// 视情况而定，有时候包含元素 i 表示考虑这种情况，有时候表示不考虑这种情况
					_ = v // do v...

				}
			}
			if bits.OnesCount(sub)&1 == 1 { // 某些题目是 0
				res = -res
			}
			ans += res // mod
		}
		ans = (ans%mod + mod) % mod // 注意最后变为非负
		return
	}

	_ = []interface{}{solveInclusionExclusion}
}

// 博弈论
// 定义必胜状态为先手必胜的状态，必败状态为先手必败的状态
// 定理 1：没有后继状态的状态是必败状态
// 定理 2：一个状态是必胜状态当且仅当存在至少一个必败状态为它的后继状态
// 定理 3：一个状态是必败状态当且仅当它的所有后继状态均为必胜状态
// 对于定理 1，如果游戏进行不下去了，那么这个玩家就输掉了游戏
// 对于定理 2，如果该状态至少有一个后继状态为必败状态，那么玩家可以通过操作到该必败状态；
//           此时对手的状态为必败状态——对手必定是失败的，而相反地，自己就获得了胜利
// 对于定理 3，如果不存在一个后继状态为必败状态，那么无论如何，玩家只能操作到必胜状态；
//           此时对手的状态为必胜状态——对手必定是胜利的，自己就输掉了游戏
// The Sprague–Grundy theorem generalizes the strategy used in nim to all games that fulfil the following requirements:
// - There are two players who move alternately.
// - The game consists of states, and the possible moves in a state do not depend on whose turn it is.
// - The game ends when a player cannot make a move.
// - The game surely ends sooner or later.
// - The players have complete information about the states and allowed moves, and there is no randomness in the game.
// 推荐 https://blog.csdn.net/acm_cxlove/article/details/7854530
// https://oi-wiki.org/math/game-theory/
// 个人写的总结 https://github.com/SDU-ACM-ICPC/Qiki/blob/master/%E5%8D%9A%E5%BC%88%E8%AE%BA(Game%20Theory).md
// TODO: 题目推荐 https://blog.csdn.net/ACM_cxlove/article/details/7854526
func gameTheoryCollection() {
	{
		// CF 1194D 打表
		// 上面三定理的基础题目
		const mx = 1000
		const k = 4
		win := [mx]bool{}
		win[1] = true
		win[2] = true
		for i := 3; i < k; i++ {
			win[i] = !win[i-1] || !win[i-2]
		}
		win[k] = true
		for i := k + 1; i < mx; i++ {
			win[i] = !win[i-1] || !win[i-2] || !win[i-k]
		}
		for i := 0; i < mx; i++ {
			Println(i, win[i])
		}
	}

	// 异或和不为0零则先手必胜
	// https://blog.csdn.net/weixin_44023181/article/details/85619512
	// 模板题 https://www.luogu.com.cn/problem/P2197
	nim := func(a []int) (firstWin bool) {
		sum := 0
		for _, v := range a {
			sum ^= v
		}
		return sum != 0
	}

	// Sprague-Grundy theorem
	// 有向图游戏的某个局面必胜 <=> 该局面对应节点的 SG 函数值 > 0
	// 有向图游戏的某个局面必败 <=> 该局面对应节点的 SG 函数值 = 0
	// 推荐资料 Competitive Programmer’s Handbook Ch.25
	// https://oi-wiki.org/math/game-theory/#sg
	// https://en.wikipedia.org/wiki/Sprague%E2%80%93Grundy_theorem
	// https://cp-algorithms.com/game_theory/sprague-grundy-nim.html
	// 整数分拆博弈 https://codeforces.com/problemset/problem/87/C
	{
		// 剪纸博弈
		// http://poj.org/problem?id=2311
		var n, m int
		sg := make([][]int, n+1)
		for i := range sg {
			sg[i] = make([]int, m+1)
			for j := range sg[i] {
				sg[i][j] = -1
			}
		}
		var SG func(int, int) int
		SG = func(x, y int) (mex int) {
			ptr := &sg[x][y]
			if *ptr != -1 {
				return *ptr
			}
			defer func() { *ptr = mex }()
			has := map[int]bool{} // 若能确定 mex 上限可以用 bool 数组
			for i := 2; i <= x-i; i++ {
				has[SG(i, y)^SG(x-i, y)] = true
			}
			for i := 2; i <= y-i; i++ {
				has[SG(x, i)^SG(x, y-i)] = true
			}
			for ; has[mex]; mex++ {
			}
			return
		}

		// 设定一些初始必败局面
		sg[2][2] = 0
		sg[2][3] = 0
		sg[3][2] = 0
		// 计算有向图游戏的 SG 函数值
		ans := SG(n, m)
		Println(ans)
	}

	_ = []interface{}{nim}
}

// 数值分析
// https://zh.wikipedia.org/wiki/%E6%95%B0%E5%80%BC%E5%88%86%E6%9E%90
func numericalAnalysisCollection() {
	type mathF func(x float64) float64

	// Simpson's 1/3 rule
	// https://en.wikipedia.org/wiki/Simpson%27s_rule
	// 证明过程 https://phqghume.github.io/2018/05/19/%E8%87%AA%E9%80%82%E5%BA%94%E8%BE%9B%E6%99%AE%E6%A3%AE%E6%B3%95/
	simpson := func(l, r float64, f mathF) float64 { h := (r - l) / 2; return h * (f(l) + 4*f(l+h) + f(r)) / 3 }

	// 不放心的话还可以设置一个最大递归深度 maxDeep
	// 15eps 的证明过程 http://www2.math.umd.edu/~mariakc/teaching/adaptive.pdf
	var _asr func(l, r, eps, A float64, f mathF) float64
	_asr = func(l, r, eps, A float64, f mathF) float64 {
		mid := l + (r-l)/2
		L := simpson(l, mid, f)
		R := simpson(mid, r, f)
		if math.Abs(L+R-A) <= 15*eps {
			return L + R + (L+R-A)/15
		}
		return _asr(l, mid, eps/2, L, f) + _asr(mid, r, eps/2, R, f)
	}

	// 自适应辛普森积分 Adaptive Simpson's Rule
	// https://en.wikipedia.org/wiki/Adaptive_Simpson%27s_method
	// https://cp-algorithms.com/num_methods/simpson-integration.html
	// 模板题 https://www.luogu.com.cn/problem/P4525 https://www.luogu.com.cn/problem/P4526
	asr := func(a, b, eps float64, f mathF) float64 { return _asr(a, b, eps, simpson(a, b, f), f) }

	_ = []interface{}{asr}
}

/* 随机算法 */

// 模拟退火 (Simulated Annealing, SA)
// 基于 Metropolis 准则
// https://oi-wiki.org/misc/simulated-annealing/
// https://en.wikipedia.org/wiki/Simulated_annealing
// https://en.wikipedia.org/wiki/Metropolis%E2%80%93Hastings_algorithm
// https://www.luogu.com.cn/blog/Darth-Che/mu-ni-tui-huo-xue-xi-bi-ji
// https://zhuanlan.zhihu.com/p/47234502
// 技巧：可以在时限内重复跑 SA 取最优值，防止脸黑
// 模板题 https://www.luogu.com.cn/problem/P1337
// LC/周赛197D https://leetcode-cn.com/contest/weekly-contest-197/problems/best-position-for-a-service-centre/ http://poj.org/problem?id=2420 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=14&page=show_problem&problem=1169
func simulatedAnnealing(f func(x float64) float64) float64 {
	// 例：最小值
	x := 0.0
	ans := f(x)
	for t := 1e5; t > 1e-8; t *= 0.99 {
		xx := x + (2*rand.Float64()-1)*t
		v := f(xx)
		if v < ans || math.Exp((ans-v)/t) > rand.Float64() { // 最小直接取，或者以一定概率接受较大的值
			ans = v
			x = xx
		}
	}
	return ans
}

// 爬山
// https://leetcode-cn.com/problems/best-position-for-a-service-centre/solution/fu-wu-zhong-xin-de-zui-jia-wei-zhi-by-leetcode-sol/
