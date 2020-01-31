package copypasta

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// NOTE: a%-b == a%b
func mathCollection() {
	const mod int64 = 1e9 + 7 // 998244353

	// 阶乘
	factorial := func(n int) int64 {
		x := int64(1)
		for i := int64(2); i <= int64(n); i++ {
			x = x * i % mod
		}
		return x
	}

	calcGCD := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	calcGCDN := func(nums ...int) (gcd int) {
		gcd = nums[0]
		for _, v := range nums[1:] {
			gcd = calcGCD(gcd, v)
		}
		return
	}
	calcLCM := func(a, b int) int {
		return a / calcGCD(a, b) * b
	}

	// 给定数组，统计所有区间的 GCD 值
	// 返回 map[GCD值]等于该值的区间个数
	cntRangeGCD := func(arr []int) map[int]int64 {
		n := len(arr)
		cntMp := map[int]int64{}
		gcd := make([]int, n)
		copy(gcd, arr)
		lPos := make([]int, n)
		for i, v := range arr {
			lPos[i] = i
			// 从当前位置 i 往左遍历，更新 gcd[j] 的同时维护等于 gcd[j] 的区间最左端位置
			for j := i; j >= 0; j = lPos[j] - 1 {
				gcd[j] = calcGCD(gcd[j], v)
				g := gcd[j]
				for lPos[j] > 0 && calcGCD(gcd[lPos[j]-1], v) == g {
					lPos[j] = lPos[lPos[j]-1]
				}
				// [lPos[j], j], [lPos[j]+1, j], ..., [j, j] 的区间 GCD 值均等于 gcd[j]
				cntMp[g] += int64(j - lPos[j] + 1)
			}
		}
		return cntMp
	}

	// 判断一个数是否为质数
	isPrime := func(n int64) bool {
		for i := int64(2); i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return n >= 2
	}

	// 预处理: [2,mx] 范围内的质数
	// 埃拉托斯特尼筛法 Sieve of Eratosthenes
	sieve := func() {
		const mx int = 1e6
		primes := make([]int, 0, mx/10) // need check
		isPrime := [mx + 1]bool{}
		for i := range isPrime {
			isPrime[i] = true
		}
		isPrime[0], isPrime[1] = false, false
		for i := 2; i <= mx; i++ {
			if isPrime[i] {
				primes = append(primes, i)
				for j := 2 * i; j <= mx; j += i {
					isPrime[j] = false
				}
			}
		}

		// EXTRA: pi(n), the number of primes <= n
		// https://oeis.org/A000720
		pi := [mx + 1]int{}
		for i := 2; i <= mx; i++ {
			pi[i] = pi[i-1]
			if isPrime[i] {
				pi[i]++
			}
		}
	}

	// 预处理: [2,mx] 范围内数的质因子（例如 factors[12] = [2,3]）
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
	lpfAll := func() {
		const mx int = 1e6
		lpf := [mx + 1]int{}
		lpf[1] = 1
		for i := 2; i <= mx; i++ {
			if lpf[i] == 0 {
				for j := i; j <= mx; j += i {
					if lpf[j] == 0 {
						lpf[j] = i
					}
				}
			}
		}

		// EXTRA: 分解 v
		var v int
		for v > 1 {
			p := lpf[v]
			// 处理 p ...
			for v /= p; lpf[v] == p; v /= p {
			}
		}
	}

	// GPF(n): greatest prime dividing n, for n >= 2; a(1)=1
	// https://oeis.org/A006530
	// 可以预处理出 LPF 然后再试除

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

	// 素因数分解 prime factorization
	primeFactors := func(n int64) (factors []int64, exponents []int) {
		for i := int64(2); i*i <= n; i++ {
			cnt := 0
			for ; n%i == 0; n /= i {
				cnt++
			}
			if cnt > 0 {
				factors = append(factors, i)
				exponents = append(exponents, cnt)
			}
		}
		if n > 1 {
			factors = append(factors, n)
			exponents = append(exponents, 1)
		}

		// EXTRA: 约数个数 d(n), also called tau(n) or sigma_0(n)
		// https://oeis.org/A000005

		// EXTRA: 约数之和 sigma(n), also called sigma_1(n)
		// https://oeis.org/A000203

		return
	}

	// 预处理: [2,mx] 范围内数的质因子个数
	// Number of distinct primes dividing n (also called omega(n))
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
	}

	// 预处理 [2,mx] 的质因数分解的系数和
	// Number of prime divisors of n counted with multiplicity (also called bigomega(n) or Omega(n))
	// https://oeis.org/A001222
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

		// EXTRA: 前缀和
		//for i := 3; i <= mx; i++ {
		//	cnt[i] += cnt[i-1]
		//}
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

	// n 的欧拉函数（互素的数的个数）Euler totient function
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

	// 预处理 [2,mx] 的欧拉函数（互素的数的个数）Euler totient function
	// https://oeis.org/A000010
	// NOTE: phi 的递归调用（指 phi[phi...[n]]）是 log 级别收敛的：奇数减一，偶数减半
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

	//

	// exgcd solve equation ax+by=gcd(a,b)
	// we have |x|<=b and |y|<=a in result (x,y)
	// https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
	var exgcd func(a, b int) (gcd, x, y int)
	exgcd = func(a, b int) (gcd, x, y int) {
		if b == 0 {
			return a, 1, 0
		}
		gcd, y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}

	// 逆元
	// ax ≡ 1 (mod m)
	modInverse := func(a, m int) int {
		_, x, _ := exgcd(a, m)
		return (x%m + m) % m
	}

	// 费马小定理
	// ax ≡ 1 (mod p)
	// x^-1 ≡ a^(p-2) (mod p)
	modInverseP := func(a, p int64) int64 {
		quickPow := func(x, n, p int64) int64 {
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
		return quickPow(a, p-2, p)
	}

	// 计算 a/b (mod m)
	modFrac := func(a, b, m int) int {
		return a * modInverse(b, m) % m
	}

	// 线性同余方程组
	// ai * x ≡ bi (mod mi)
	// 解为 x ≡ b (mod m)
	// 有解时返回 (b, m)，无解时返回 (0, -1)
	// 推导过程见《挑战程序设计竞赛》P292
	// 注意乘法溢出的可能
	solveLinearCongruence := func(A, B, M []int) (int, int) {
		x, m := 0, 1
		for i, mi := range M {
			a, b := A[i]*m, B[i]-A[i]*x
			d := calcGCD(a, mi)
			if b%d != 0 {
				return 0, -1
			}
			t := modFrac(b/d, a/d, mi/d)
			x += m * t
			m *= mi / d
		}
		return (x%m + m) % m, m
	}

	// a*b%mod，原理和 a^n%mod 类似
	quickMul := func(a, b, mod int64) (res int64) {
		for ; b > 0; b >>= 1 {
			if b&1 == 1 {
				res = (res + a) % mod
			}
			a = (a + a) % mod
		}
		return
	}

	// TODO: 扩展中国剩余定理 (EXCRT)
	// https://www.luogu.com.cn/problemnew/solution/P4777

	//

	// 数论分块/除法分块
	// https://oeis.org/A006218
	// a(n) = Sum_{k=1..n} floor(n/k)
	//      = 2*(Sum_{i=1..floor(sqrt(n))} floor(n/i)) - floor(sqrt(n))^2
	// thus, a(n) % 2 == floor(sqrt(n)) % 2

	//

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

	// 隔板法
	// https://zh.wikipedia.org/wiki/%E9%9A%94%E6%9D%BF%E6%B3%95

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

	_ = []interface{}{
		factorial, calcGCDN, calcLCM, cntRangeGCD,
		isPrime, sieve, primeFactorsAll, lpfAll, divisors, doDivisors, primeFactors, distinctPrimesCountAll, primeExponentsCountAll, calcPhi, phiAll,
		modInverseP, modFrac, solveLinearCongruence, quickMul,
		fractionToDecimal, decimalToFraction,
		moveToRange,
	}
}

// https://oi-wiki.org/math/game-theory/
func gameTheoryCollection() {
	// 异或和不为0零则先手必胜
	// https://blog.csdn.net/weixin_44023181/article/details/85619512
	nim := func(a []int) (firstWin bool) {
		sum := 0
		for _, v := range a {
			sum ^= v
		}
		return sum != 0
	}

	// https://cp-algorithms.com/game_theory/sprague-grundy-nim.html
	var sg []int
	initSG := func(mx int) {
		_ = sg
		// TODO: implement
	}

	_ = []interface{}{nim, initSG}
}

type mathF func(x float64) float64

// Simpson's 1/3 rule
// https://en.wikipedia.org/wiki/Simpson%27s_rule
// 证明过程 https://phqghume.github.io/2018/05/19/%E8%87%AA%E9%80%82%E5%BA%94%E8%BE%9B%E6%99%AE%E6%A3%AE%E6%B3%95/
func simpson(l, r float64, f mathF) float64 {
	h := (r - l) / 2
	return h * (f(l) + 4*f(l+h) + f(r)) / 3
}

// 不放心的话还可以设置一个最大递归深度 maxDeep
// 15eps 的证明过程 http://www2.math.umd.edu/~mariakc/teaching/adaptive.pdf
func _asr(l, r, eps, A float64, f mathF) float64 {
	mid := l + (r-l)/2
	L := simpson(l, mid, f)
	R := simpson(mid, r, f)
	if math.Abs(L+R-A) <= 15*eps {
		return L + R + (L+R-A)/15
	}
	return _asr(l, mid, eps/2, L, f) + _asr(mid, r, eps/2, R, f)
}

// Adaptive Simpson's Rule
// https://en.wikipedia.org/wiki/Adaptive_Simpson%27s_method
// https://cp-algorithms.com/num_methods/simpson-integration.html
func asr(a, b, eps float64, f mathF) float64 { return _asr(a, b, eps, simpson(a, b, f), f) }
