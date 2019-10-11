package copypasta

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
		for i := int64(1); i*i <= n; i++ {
			if n%i == 0 {
				res = append(res, i)
				if d := n / i; d != i {
					res = append(res, d)
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
