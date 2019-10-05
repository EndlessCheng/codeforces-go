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

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	_ = []interface{}{calcLCM, isPrime}
}

// exgcd solve equation ax+by=gcd(a,b)
// |x|<=b
// |y|<=a
func exgcd(a, b int64) (gcd, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
}
