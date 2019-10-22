package copypasta

import "math"

func mathCollection() {
	calcGCD := func(a, b int64) int64 {
		for b > 0 {
			a, b = b, a%b
		}
		return a
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

	sieve := func(n int) (prime []int, isPrime []bool) {
		isPrime = make([]bool, n+1)
		for i := range isPrime {
			isPrime[i] = true
		}
		isPrime[0], isPrime[1] = false, false
		for i := 2; i <= n; i++ {
			if isPrime[i] {
				prime = append(prime, i)
				for j := 2 * i; j <= n; j += i {
					isPrime[j] = false
				}
			}
		}
		return
	}

	divisors := func(n int64) (res []int64) {
		for d := int64(1); d*d <= n; d++ {
			if n%d == 0 {
				res = append(res, d)
				if d2 := n / d; d2 != d {
					res = append(res, d2)
				}
			}
		}
		return
	}

	primeFactors := func(n int64) (factors []int64, exps []int) {
		for i := int64(2); i*i <= n; i++ {
			cnt := 0
			for ; n%i == 0; n /= i {
				cnt++
			}
			if cnt > 0 {
				factors = append(factors, i)
				exps = append(exps, cnt)
			}
		}
		if n != 1 {
			factors = append(factors, n)
			exps = append(exps, 1)
		}
		return
	}

	// ax ≡ 1 (mod m)
	modInverse := func(a, m int64) int64 {
		_, x, _ := exgcd(a, m)
		return (x%m + m) % m
	}

	_ = []interface{}{calcLCM, isPrime, sieve, divisors, primeFactors, modInverse}
}

// exgcd solve equation ax+by=gcd(a,b)
// we have |x|<=b and |y|<=a in result (x,y)
func exgcd(a, b int64) (gcd, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
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
func asr(a, b, eps float64, f mathF) float64 { return _asr(a, b, eps, simpson(a, b, f), f) }
