package copypasta

import (
	. "fmt"
	"math"
	"math/big"
	"math/bits"
	"regexp"
	"strconv"
	"strings"
)

/*
todo 待整理 https://math.stackexchange.com/questions/1955105/corectness-of-prime-factorization-over-a-range

CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=number+theory
CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=combinatorics

NOTE: a%-b == a%b
*/

func numberTheoryCollection() {
	const mod int64 = 1e9 + 7 // 998244353
	// https://oeis.org/A000040
	primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	sqCheck := func(a int64) bool { r := int64(math.Round(math.Sqrt(float64(a)))); return r*r == a }
	cubeCheck := func(a int64) bool { r := int64(math.Round(math.Cbrt(float64(a)))); return r*r*r == a }
	sqrt := func(a int64) int64 {
		r := int64(math.Round(math.Sqrt(float64(a))))
		if r*r == a {
			return r
		}
		return -1
	}
	cbrt := func(a int64) int64 {
		r := int64(math.Round(math.Cbrt(float64(a))))
		if r*r*r == a {
			return r
		}
		return -1
	}

	// Least k such that H(k) > n, where H(k) is the harmonic number sum_{i=1..k} 1/i
	// https://oeis.org/A002387
	// https://oeis.org/A004080
	// a(n) = largest m such that the harmonic number H(m)= Sum_{i=1..m} 1/i is < n
	// https://oeis.org/A115515

	// a(n) = smallest prime p such that Sum_{primes q = 2, ..., p} 1/q exceeds n
	// 5, 277, 5_195_977, 1801241230056600523
	// https://oeis.org/A016088
	// a(n) = largest prime p such that Sum_{primes q = 2, ..., p} 1/q does not exceed n
	// 3, 271, 5_195_969, 1801241230056600467
	// https://oeis.org/A223037

	// a*b%m，原理和 a^n%m 类似
	mul := func(a, b, m int64) (res int64) {
		for ; b > 0; b >>= 1 {
			if b&1 == 1 {
				res = (res + a) % m
			}
			a = (a + a) % m
		}
		return
	}
	muls := func(nums ...int64) int64 {
		res := nums[0]
		for _, v := range nums[1:] {
			res = res * v % mod
		}
		return res
	}

	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int64) int64 { return a / gcd(a, b) * b }

	// 前 n 个数的 LCM
	// https://oeis.org/A003418
	// Mangoldt Function
	// https://mathworld.wolfram.com/MangoldtFunction.html
	// https://oeis.org/A014963

	gcds := func(nums ...int64) (g int64) {
		g = nums[0]
		for _, v := range nums[1:] {
			g = gcd(g, v)
		}
		return
	}

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

	/* 质数
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

	// 哥德巴赫猜想 - 偶数分拆的最小质数
	// 由质数分布可知选到一对质数的概率是 O(1/ln^2(n))
	// https://oeis.org/A020481
	// n https://oeis.org/A025018
	// a(n) https://oeis.org/A025019
	// 1e9 内最大的为 a(721013438) = 1789
	// 2e9 内最大的为 a(1847133842) = 1861

	// 预处理: [2,mx] 范围内的质数
	// 埃拉托斯特尼筛法 Sieve of Eratosthenes
	// 也有线性时间的算法，见 https://oi-wiki.org/math/sieve/ 以及进阶指南 p.136-137
	// 素数个数 π(n) https://oeis.org/A000720
	sieve := func() {
		const mx int = 1e6
		primes := make([]int, 0, mx/10) // need check
		isP := [mx + 1]bool{}
		for i := range isP {
			isP[i] = true
		}
		isP[0], isP[1] = false, false
		for i := 2; i <= mx; i++ {
			if isP[i] {
				primes = append(primes, i)
				for j := 2 * i; j <= mx; j += i {
					isP[j] = false
				}
			}
		}

		// EXTRA: pi(n), the number of primes <= n
		// https://oeis.org/A000720
		pi := [mx + 1]int{}
		for i := 2; i <= mx; i++ {
			pi[i] = pi[i-1]
			if isP[i] {
				pi[i]++
			}
		}
	}

	// 区间筛法
	// 预处理 [2,√R] 的所有质数，去筛 [L,R] 之间的质数
	// EXTRA: 区间最大质数间隔
	// Prime gaps: differences between consecutive primes
	// https://oeis.org/A001223
	// Indices https://oeis.org/A005669
	// Records https://oeis.org/A005250

	// 质因数分解 prime factorization
	// 返回分解出的质数及其指数
	// https://mathworld.wolfram.com/PrimeFactorization.html
	// todo 更高效的算法 - Pollard's Rho
	primeFactorization := func(n int) (factors [][2]int) {
		for i := 2; i*i <= n; i++ {
			k := 0
			for ; n%i == 0; n /= i {
				k++
			}
			if k > 0 {
				factors = append(factors, [2]int{i, k})
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

	// 预处理: [2,mx] 的质因数分解的系数和 bigomega(n) or Omega(n)
	// Number of prime divisors of n counted with multiplicity
	// https://oeis.org/A001222
	//
	// Omega(n) - omega(n)
	// https://oeis.org/A046660
	// a(n) depends only on prime signature of n (cf. https://oeis.org/A025487)
	// So a(24) = a(375) since 24 = 2^3 * 3 and 375 = 3 * 5^3 both have prime signature (3, 1)
	// a(n) = 0 for squarefree n
	primeExponentsCountAll := func() {
		const mx int = 1e6
		cnts := make([]int, mx+1)
		primes := make([]int, 0, mx/10) // need check
		for i := 2; i <= mx; i++ {
			if cnts[i] == 0 {
				primes = append(primes, i)
				cnts[i] = 1
			}
			for _, p := range primes {
				if j := i * p; j <= mx {
					cnts[j] = cnts[i] + 1
				} else {
					break
				}
			}
		}

		// EXTRA: 前缀和，即 Omega(n!)
		// https://oeis.org/A022559
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

	/* 约数
	 */

	// 枚举一个数的全部约数
	divisors := func(n int64) (res []int64) {
		for d := int64(1); d*d <= n; d++ {
			if n%d == 0 {
				res = append(res, d)
				if d*d < n {
					res = append(res, n/d)
				}
			}
		}
		//sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
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
		return
	}
	doDivisors2 := func(n int, do func(d1, d2 int)) {
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				do(d, n/d)
			}
		}
		return
	}

	// 预处理: [1,mx] 范围内数的所有约数
	// 复杂度 O(nlogn)
	// NOTE: 1~n 的约数个数总和大约为 nlogn
	// NOTE: divisors[x] 为奇数 => x 为完全平方数 https://oeis.org/A000290
	// NOTE: halfDivisors(x) 为 ≤√x 的因数集合 https://oeis.org/A161906
	divisorsAll := func() {
		const mx int = 1e5
		divisors := [mx + 1][]int{}
		for i := 1; i <= mx; i++ {
			for j := i; j <= mx; j += i {
				divisors[j] = append(divisors[j], i)
			}
		}

		isSquareNumber := func(x int) bool { return len(divisors[x])&1 == 1 }
		halfDivisors := func(x int) []int { d := divisors[x]; return d[:(len(d)-1)/2+1] }

		_, _ = isSquareNumber, halfDivisors
	}

	// EXTRA: 约数个数 d(n), also called tau(n) or sigma_0(n) https://oeis.org/A000005
	//        约数个数的前缀和 a(n) = Sum_{k=1..n} floor(n/k) https://oeis.org/A006218
	//                            = 见后文「数论分块/除法分块」

	// EXTRA: 约数之和 sigma(n), also called sigma_1(n) https://oeis.org/A000203
	//        约数之和的前缀和 a(n) = Sum_{k=1..n} k*floor(n/k) https://oeis.org/A024916

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

	// LPF(n): least prime dividing n (when n > 1); a(1) = 1
	// 有时候数据范围比较大，用 primeFactorsAll 预处理会 MLE，这时候就要用 LPF 了（同样是预处理但是内存占用低）
	// 先预处理出 LPF，然后对要处理的数 v 不断地除 LPF(v) 直到等于 1
	// https://oeis.org/A020639
	//
	// GPF(n): greatest prime dividing n, for n >= 2; a(1)=1
	// https://oeis.org/A006530
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

	// 预处理: [2,mx] 的不同的质因子个数 omega(n)
	// Number of distinct primes dividing n
	// https://oeis.org/A001221
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

		// EXTRA: 前缀和，即 omega(n!)
		// https://oeis.org/A013939
		for i := 3; i <= mx; i++ {
			cnts[i] += cnts[i-1]
		}
	}

	// Squarefree numbers: numbers that are not divisible by a square greater than 1
	// Lim_{n->infinity} a(n)/n = Pi^2/6
	// https://oeis.org/A005117 介绍了一种筛法
	// Numbers that are not squarefree. Numbers that are divisible by a square greater than 1
	// https://oeis.org/A013929

	// Semiprimes (or biprimes): products of two primes
	// https://oeis.org/A001358
	// Squarefree semiprimes: Numbers that are the product of two distinct primes
	// https://oeis.org/A006881

	// Squarefree part of n: a(n) is the smallest positive number m such that n/m is a square
	// Also called core(n)
	// https://oeis.org/A007913

	// Largest squarefree number dividing n: the squarefree kernel of n, rad(n), radical of n
	// https://oeis.org/A007947

	// 高合成数/反质数 Highly Composite Numbers
	// 一个高合成数一定是由另一个高合成数乘某个质数得到
	// 见进阶指南 p.140-141
	// https://oeis.org/A002182
	// https://oeis.org/A002183
	// https://www.luogu.com.cn/problem/P1463

	// n 的欧拉函数（互质的数的个数）Euler totient function
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

	// 预处理 [2,mx] 的欧拉函数（互质的数的个数）Euler totient function
	// https://oeis.org/A000010
	// NOTE: phi 的迭代（指 phi[phi...[n]]）是 log 级别收敛的：奇数减一，偶数减半
	phiAll := func() {
		const mx int = 1e6
		phi := [mx + 1]int{}
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

	/* 同余
	 */

	// 二元一次不定方程
	// exgcd solve equation ax+by=gcd(a,b)
	// we have |x|<=b and |y|<=a in result (x,y)
	// https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
	// 模板题 https://www.luogu.com.cn/problem/P5656
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
	// 模板题 https://www.luogu.com.cn/problem/P1082 https://www.luogu.com.cn/problem/P3811
	invM := func(a, m int64) int64 { _, x, _ := exgcd(a, m); return (x%m + m) % m }

	pow := func(x, n, p int64) int64 {
		x %= p
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % p
			}
			x = x * x % p
		}
		return res
	}

	// 费马小定理求质数逆元
	// ax ≡ 1 (mod p)
	// x^-1 ≡ a^(p-2) (mod p)
	invP := func(a, p int64) int64 { return pow(a, p-2, p) }

	// 有理数取模
	// 模板题 https://www.luogu.com.cn/problem/P2613
	divM := func(a, b, m int64) int64 { return a * invM(b, m) % m }
	divP := func(a, b, p int64) int64 { return a * invP(b, p) % p }

	// 模数两两互质的线性同余方程组 - 中国剩余定理 (CRT)
	// https://blog.csdn.net/synapse7/article/details/9946013
	// todo https://codeforces.ml/blog/entry/61290
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
	// 证明见进阶指南 p.155
	// 推荐 https://blog.csdn.net/niiick/article/details/80229217
	// 模板题 https://www.luogu.com.cn/problemnew/solution/P4777
	// todo 整理 excrt := func(a, m []int) (x int) {
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

	// ai * x ≡ bi (mod mi)
	// 解为 x ≡ b (mod m)
	// 有解时返回 (b, m)，无解时返回 (0, -1)
	// 推导过程见《挑战程序设计竞赛》P292
	// 注意乘法溢出的可能
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

	// 高次同余方程 a^x ≡ b (mod p)，a 和 p 互质 - 小步大步算法 (BSGS)
	// 时间复杂度 O(√p)
	// 见进阶指南 p.155
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

	// 高次同余方程 x^a ≡ b (mod p) - N次剩余 - 原根
	// todo
	// 模板题 https://www.luogu.com.cn/problem/P5491 https://www.luogu.com.cn/problem/P5668

	/* 组合数；二项式系数
	 */

	// 阶乘
	factorial := func(n int) int64 {
		res := int64(1) % mod
		for i := 2; i <= n; i++ {
			res = res * int64(i) % mod
		}
		return res
	}

	initFactorial := func() {
		const mx int = 1e5
		F := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			F[i] = F[i-1] * int64(i) % mod
		}

		factorial := func(n int) int64 { return F[n] }
		_ = factorial
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

	// EXTRA: binomial(n, floor(n/2))
	// https://oeis.org/A001405
	combHalf := [...]int64{
		1, 1, 2, 3, 6, 10, 20, 35, 70, 126, // C(9,4)
		252, 462, 924, 1716, 3432, 6435, 12870, 24310, 48620, 92378, // C(19,9)
		184756, 352716, 705432, 1352078, 2704156, 5200300, 10400600, 20058300, 40116600, 77558760, // C(29,14)
		155117520, 300540195, 601080390, 1166803110, 2333606220, 4537567650, 9075135300, 17672631900, 35345263800, 68923264410, // C(39,19)
		137846528820, 269128937220, 538257874440, 1052049481860, 2104098963720, 4116715363800, 8233430727600, 16123801841550, 32247603683100, 63205303218876, // C(49,24)
		126410606437752, 247959266474052, 495918532948104, 973469712824056, 1946939425648112, 3824345300380220, 7648690600760440, 15033633249770520, 30067266499541040, 59132290782430712, // C(59,29)
		118264581564861424, 232714176627630544, 465428353255261088, 916312070471295267, 1832624140942590534, 3609714217008132870, 7219428434016265740, // C(66,33)
	}

	// EXTRA: Central binomial coefficients: binomial(2*n,n) = (2*n)!/(n!)^2
	// https://oeis.org/A000984

	// 仅适用于小范围的 n k
	// 更大范围的见下面的三种处理方式
	comb := func(n, k int) int64 {
		res := int64(1)
		for i := 1; i <= k; i++ {
			res = res * int64(n-k+i) / int64(i)
		}
		return res
		//return big.Int{}.Binomial(n, k).Int64()
	}

	// https://www.zhihu.com/question/26094736
	// 不推荐，因为逆元可以做到 O(nlogn) 预处理
	//{
	//	// O(n^2) 预处理，O(1) 求组合数
	//	const mod int64 = 1e9 + 7
	//	const mx int = 2e3
	//	C := [mx + 1][mx + 1]int64{}
	//	for i := 0; i <= mx; i++ {
	//		C[i][0] = 1
	//		for j := 1; j < i; j++ {
	//			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
	//		}
	//		C[i][i] = 1
	//	}
	//
	//	comb := func(n, k int) int64 { return C[n][k] }
	//	_ = comb
	//}

	{
		// O(n) 预处理，O(logn) 求组合数
		const mod int64 = 1e9 + 7
		const mx int = 1e5
		F := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			F[i] = F[i-1] * int64(i) % mod
		}
		pow := func(x, n int64) int64 {
			x %= mod
			res := int64(1)
			for ; n > 0; n >>= 1 {
				if n&1 == 1 {
					res = res * x % mod
				}
				x = x * x % mod
			}
			return res
		}
		inv := func(a int64) int64 { return pow(a, mod-2) }
		div := func(a, b int64) int64 { return a * inv(b) % mod }
		comb := func(n, k int64) int64 { return div(F[n], F[k]*F[n-k]%mod) }

		_ = comb
	}

	{
		// O(nlogn) 预处理，O(1) 求组合数
		const mod int64 = 1e9 + 7
		pow := func(x, n int64) int64 {
			x %= mod
			res := int64(1)
			for ; n > 0; n >>= 1 {
				if n&1 == 1 {
					res = res * x % mod
				}
				x = x * x % mod
			}
			return res
		}
		inv := func(a int64) int64 { return pow(a, mod-2) }
		const mx int = 1e5
		F := [mx + 1]int64{1}
		invF := [mx + 1]int64{inv(1)}
		for i := 1; i <= mx; i++ {
			F[i] = F[i-1] * int64(i) % mod
			invF[i] = inv(F[i])
		}
		comb := func(n, k int64) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

		// 卢卡斯定理
		var lucas func(n, k int64) int64
		lucas = func(n, k int64) int64 {
			if k == 0 {
				return 1
			}
			return comb(n%mod, k%mod) * lucas(n/mod, k/mod) % mod
		}

		_ = comb
	}

	// 扩展卢卡斯
	// todo https://blog.csdn.net/niiick/article/details/81064156
	// https://blog.csdn.net/skywalkert/article/details/52553048
	// https://blog.csdn.net/skywalkert/article/details/104681101
	// https://cp-algorithms.com/combinatorics/binomial-coefficients.html
	// 模板题 https://www.luogu.com.cn/problem/P4720
	// 古代猪文 https://www.luogu.com.cn/problem/P2480

	//

	// 莫比乌斯函数
	// todo

	// 莫比乌斯反演（岛娘推荐！https://zhuanlan.zhihu.com/p/133761303）
	// https://oi-wiki.org/math/mobius/

	//

	// 数论分块/除法分块
	// https://oeis.org/A006218
	// a(n) = Sum_{k=1..n} floor(n/k)
	//      = 2*(Sum_{i=1..floor(sqrt(n))} floor(n/i)) - floor(sqrt(n))^2
	// thus, a(n) % 2 == floor(sqrt(n)) % 2

	// 杜教筛 - 积性函数前缀和
	// todo 推荐 https://blog.csdn.net/weixin_43914593/article/details/104229700
	// todo http://baihacker.github.io/main/
	// The prefix-sum of multiplicative function: the black algorithm http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_the_black_algorithm.html
	// The prefix-sum of multiplicative function: Dirichlet convolution http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_dirichlet_convolution.html
	// The prefix-sum of multiplicative function: powerful number sieve http://baihacker.github.io/main/2020/The_prefix-sum_of_multiplicative_function_powerful_number_sieve.html
	// 浅谈一类积性函数的前缀和 + 套题 https://blog.csdn.net/skywalkert/article/details/50500009
	// 模板题 https://www.luogu.com.cn/problem/P4213

	//

	// Number of odd divisors of n
	// https://oeis.org/A001227
	consecutiveNumbersSum := func(n int) (ans int) {
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

	// 把 n 用 m 等分，得到 m-n%m 个 n/m 和 n%m 个 n/m+1
	partition := func(n, m int) (q, cntQ, cntQ1 int) {
		// m must > 0
		return n / m, m - n%m, n % m
	}

	_ = []interface{}{
		primes,
		sqCheck, cubeCheck, sqrt, cbrt,
		mul, muls, gcds, lcm, cntRangeGCD,
		isPrime, sieve, primeFactorization, primeDivisors, primeExponentsCountAll,
		divisors, doDivisors, doDivisors2, divisorsAll, primeFactorsAll, lpfAll, distinctPrimesCountAll, calcPhi, phiAll,
		exgcd, invM, invP, divM, divP, crt, excrt, babyStepGiantStep,
		factorial, initFactorial, _factorial, combHalf, comb,
		consecutiveNumbersSum, partition,
	}
}

// 组合、杂项
// 隔板法 https://zh.wikipedia.org/wiki/%E9%9A%94%E6%9D%BF%E6%B3%95
// 圆排列 https://zh.wikipedia.org/wiki/%E5%9C%86%E6%8E%92%E5%88%97
// 可重集排列
// 可重集组合 todo https://codeforces.ml/problemset/problem/451/E
// 错排 https://zh.wikipedia.org/wiki/%E9%94%99%E6%8E%92%E9%97%AE%E9%A2%98
// 范德蒙德恒等式 https://zh.wikipedia.org/wiki/%E8%8C%83%E5%BE%B7%E8%92%99%E6%81%92%E7%AD%89%E5%BC%8F
// 二阶递推数列通项 https://zhuanlan.zhihu.com/p/75096951
// 斯特林数 https://blog.csdn.net/ACdreamers/article/details/8521134
// Stirling numbers of the first kind, s(n,k) https://oeis.org/A008275
//    将 n 个元素排成 k 个非空循环排列的方法数
//    s(n,k) 的递推公式： s(n,k)=(n-1)*s(n-1,k)+s(n-1,k-1), 1<=k<=n-1
//    边界条件：s(n,0)=0, n>=1    s(n,n)=1, n>=0
// Stirling numbers of the second kind, S2(n,k) https://oeis.org/A008277
//    将 n 个元素拆分为 k 个非空集的方法数
//    S2(n, k) = (1/k!) * Sum_{i=0..k} (-1)^(k-i)*binomial(k, i)*i^n.
//    S2(n,k) 的递推公式：S2(n,k)=k*S2(n-1,k)+S2(n-1,k-1), 1<=k<=n-1
//    边界条件：S(n,0)=0, n>=1    S(n,n)=1, n>=0
func miscCollection() {
	// n married couples are seated in a row so that every wife is to the left of her husband
	// 若不考虑顺序，则所有排列的个数为 (2n)!
	// 考虑顺序可以发现，对于每一对夫妻来说，妻子在丈夫左侧的情况和在右侧的情况相同且不同对夫妻之间是独立的
	// 因此每有一对夫妻，符合条件的排列个数就减半
	// 所以结果为 a(n) = (2n)!/2^n
	// https://oeis.org/A000680
	// https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/

	// 容斥原理 Inclusion–exclusion principle
	// 参考《挑战程序设计竞赛》P296
	solveInclusionExclusion := func(arr []int) (ans int) {
		n := len(arr)
		for sub := uint(0); sub < 1<<n; sub++ {
			res := 0
			for i := 0; i < n; i++ {
				if sub>>i&1 == 1 {
					_ = arr[i]
				}
			}
			if bits.OnesCount(sub)&1 == 1 {
				ans += res
			} else {
				ans -= res
			}
		}
		return
	}

	// 从 st 跳到 [l,r]，每次跳 d 个单位长度，问首次到达的位置（或无法到达）
	moveToRange := func(st, d, l, r int) (firstPos int, ok bool) {
		switch {
		case st < l:
			if d <= 0 {
				return
			}
			return l + ((st-l)%d+d)%d, true
		case st <= r:
			return st, true
		default:
			if d >= 0 {
				return
			}
			return r + ((st-r)%d+d)%d, true
		}
	}

	// floatStr must contain a .
	// all decimal part must have same length
	// floatToInt("3.000100", 1e6) => 3000100
	// "3.0001" is not allowed
	floatToInt := func(floatStr string, shift10 int) int {
		splits := strings.SplitN(floatStr, ".", 2)
		i, _ := strconv.Atoi(splits[0])
		d, _ := strconv.Atoi(splits[1])
		return i*shift10 + d
	}

	// floatToRat("1.2", 1e1) => (6, 5)
	floatToRat := func(floatStr string, shift10 int) (m, n int) {
		m = floatToInt(floatStr, shift10)
		n = shift10
		var g int // g:= calcGCD(m, n)
		m /= g
		n /= g
		return
	}

	// https://en.wikipedia.org/wiki/Repeating_decimal
	// https://oeis.org/A051626 Period of decimal representation of 1/n, or 0 if 1/n terminates.
	// https://oeis.org/A036275 The periodic part of the decimal expansion of 1/n. Any initial 0's are to be placed at end of cycle.
	// 例如 (2, -3) => ("-0.", "6")
	// b must not be zero
	fractionToDecimal := func(a, b int64) (beforeCycle, cycle []byte) {
		if a == 0 {
			return []byte{'0'}, nil
		}
		var res []byte
		if a < 0 && b > 0 || a > 0 && b < 0 {
			res = []byte{'-'}
		}
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		res = append(res, strconv.FormatInt(a/b, 10)...)

		r := a % b
		if r == 0 {
			return res, nil
		}
		res = append(res, '.')

		posMap := map[int64]int{}
		for r != 0 {
			if pos, ok := posMap[r]; ok {
				return res[:pos], res[pos:]
			}
			posMap[r] = len(res)
			r *= 10
			res = append(res, '0'+byte(r/b))
			r %= b
		}
		return res, nil
	}

	// decimal like "2.15(376)", which means "2.15376376376..."
	// https://zh.wikipedia.org/wiki/%E5%BE%AA%E7%8E%AF%E5%B0%8F%E6%95%B0#%E5%8C%96%E7%82%BA%E5%88%86%E6%95%B8%E7%9A%84%E6%96%B9%E6%B3%95
	r := regexp.MustCompile(`(?P<integerPart>\d+)\.?(?P<nonRepeatingPart>\d*)\(?(?P<repeatingPart>\d*)\)?`)
	decimalToFraction := func(decimal string) (a, b int64) {
		match := r.FindStringSubmatch(decimal)
		integerPart, nonRepeatingPart, repeatingPart := match[1], match[2], match[3]
		intPartNum, _ := strconv.ParseInt(integerPart, 10, 64)
		if repeatingPart == "" {
			repeatingPart = "0"
		}
		b, _ = strconv.ParseInt(strings.Repeat("9", len(repeatingPart))+strings.Repeat("0", len(nonRepeatingPart)), 10, 64)
		a, _ = strconv.ParseInt(nonRepeatingPart+repeatingPart, 10, 64)
		if nonRepeatingPart != "" {
			v, _ := strconv.ParseInt(nonRepeatingPart, 10, 64)
			a -= v
		}
		a += intPartNum * b
		// 后续需要用 gcd 化简
		// 或者用 return big.NewRat(a, b)
		return
	}

	_ = []interface{}{
		solveInclusionExclusion,
		moveToRange,
		floatToRat, fractionToDecimal, decimalToFraction,
	}
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
	// https://oi-wiki.org/math/game-theory/#sg
	// https://en.wikipedia.org/wiki/Sprague%E2%80%93Grundy_theorem
	// https://cp-algorithms.com/game_theory/sprague-grundy-nim.html
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

/* 杂项 */

//func grayCode(length int) []int {
//	if length == 1 {
//		return []int{0, 1}
//	}
//	part0 := grayCode(length - 1)
//	part1 := make([]int, len(part0))
//	for i, v := range part0 {
//		part1[len(part0)-i-1] = v
//	}
//	for i, v := range part1 {
//		part1[i] = v | 1<<(length-1)
//	}
//	return append(part0, part1...)
//}
func grayCode(length int) []int {
	ans := make([]int, 1<<length)
	for i := range ans {
		ans[i] = i ^ i>>1
	}
	return ans
}

// https://oeis.org/A000127
// Maximal number of regions obtained by joining n points around a circle by straight lines.
// Also number of regions in 4-space formed by n-1 hyperplanes.
//
//     n*(n-1)*(n*n-5*n+18)/24+1

// https://leetcode-cn.com/contest/weekly-contest-139/problems/adding-two-negabinary-numbers/
func addNegabinary(a1, a2 []int) []int {
	if len(a1) < len(a2) {
		a1, a2 = a2, a1
	}
	for i, j := len(a1)-1, len(a2)-1; j >= 0; {
		a1[i] += a2[j]
		i--
		j--
	}
	ans := append(make([]int, 2), a1...)
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] >= 2 {
			ans[i] -= 2
			if ans[i-1] >= 1 {
				ans[i-1]--
			} else {
				ans[i-1]++
				ans[i-2]++
			}
		}
	}
	for i, v := range ans {
		if v != 0 {
			return ans[i:]
		}
	}
	return []int{0}
}

// https://leetcode.com/problems/convert-to-base-2/
func toNegabinary(n int) (res string) {
	if n == 0 {
		return "0"
	}
	for ; n != 0; n = -(n >> 1) {
		res = string('0'+n&1) + res
	}
	return
}
