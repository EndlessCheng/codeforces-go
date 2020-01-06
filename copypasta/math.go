package copypasta

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func mathCollection() {
	const mod int64 = 1e9 + 7 // 998244353
	factorial := func(n int) int64 {
		x := int64(1)
		for i := int64(2); i <= int64(n); i++ {
			x = x * i % mod
		}
		return x
	}

	calcGCD := func(a, b int64) int64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	calcGCDN := func(nums ...int64) (gcd int64) {
		gcd = nums[0]
		for _, v := range nums[1:] {
			gcd = calcGCD(gcd, v)
		}
		return
	}
	calcLCM := func(a, b int64) int64 {
		return a / calcGCD(a, b) * b
	}

	isPrime := func(n int64) bool {
		for i := int64(2); i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return n >= 2
	}

	sieve := func(n_ int) (primes []int, isPrime []bool) {
		primes = make([]int, 0, n_/10) // need check
		isPrime = make([]bool, n_+1)
		for i := range isPrime {
			isPrime[i] = true
		}
		isPrime[0], isPrime[1] = false, false
		for i := 2; i <= n_; i++ {
			if isPrime[i] {
				primes = append(primes, i)
				for j := 2 * i; j <= n_; j += i {
					isPrime[j] = false
				}
			}
		}
		// https://oeis.org/A000720
		pi := make([]int, n_+1)
		for i := 2; i <= n_; i++ {
			pi[i] = pi[i-1]
			if isPrime[i] {
				pi[i]++
			}
		}
		return
	}

	// for i>=2, primes[i][0] == i means i is prime
	primeFactorsAll := func(n_ int) (factors [][]int) {
		factors = make([][]int, n_+1)
		for i := 2; i <= n_; i++ {
			if len(factors[i]) == 0 {
				for j := i; j <= n_; j += i {
					factors[j] = append(factors[j], i)
				}
			}
		}
		return
	}

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
	doDivisors := func(n int, do func(div int)) {
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
		if n != 1 {
			factors = append(factors, n)
			exponents = append(exponents, 1)
		}
		return
	}

	// https://oeis.org/A001222
	// Number of prime divisors of n counted with multiplicity (also called bigomega(n) or Omega(n)).
	// 生成 2-n 的质因数分解的系数和
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
	primeExponentsCount := func(n int) []int {
		cnt := make([]int, n+1)
		primes := make([]int, 0, n/10) // need check
		for i := 2; i <= n; i++ {
			if cnt[i] == 0 {
				primes = append(primes, i)
				cnt[i] = 1
			}
			for _, p := range primes {
				if j := i * p; j <= n {
					cnt[j] = cnt[i] + 1
				} else {
					break
				}
			}
		}
		// 前缀和
		// for i := 3; i <= n; i++ {
		//		cnt[i] += cnt[i-1]
		// }
		return cnt
	}

	// https://oeis.org/A020639
	// Lpf(n): least prime dividing n (when n > 1); a(1) = 1.
	calcLPF := func(n int) []int {
		lpf := make([]int, n+1)
		lpf[1] = 1
		for i := 2; i <= n; i++ {
			if lpf[i] == 0 {
				for j := i; j <= n; j += i {
					if lpf[j] == 0 {
						lpf[j] = i
					}
				}
			}
		}
		return lpf
	}

	//

	// exgcd solve equation ax+by=gcd(a,b)
	// we have |x|<=b and |y|<=a in result (x,y)
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

	// ax ≡ 1 (mod m)
	modInverse := func(a, m int64) int64 {
		_, x, _ := exgcd(a, m)
		return (x%m + m) % m
	}

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
	// ax ≡ 1 (mod p)
	modInverseP := func(a, p int64) int64 {
		return quickPow(a, p-2, p)
	}

	// a/b % m
	modFrac := func(a, b, m int64) int64 {
		return a * modInverse(b, m) % m
	}

	//

	// https://oeis.org/A006218
	// a(n) = Sum_{k=1..n} floor(n/k)
	//      = 2*(Sum_{i=1..floor(sqrt(n))} floor(n/i)) - floor(sqrt(n))^2
	// thus, a(n) % 2 == floor(sqrt(n)) % 2

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
		factorial, calcGCDN, calcLCM,
		isPrime, sieve, primeFactorsAll, divisors, doDivisors, primeFactors, primeExponentsCount, calcLPF,
		modInverseP, modFrac,
		fractionToDecimal, decimalToFraction,
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
	//var sg []int
	//initSG := func(n int) {
	//	// TODO implement
	//}

	_ = []interface{}{nim}
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
