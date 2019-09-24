package copypasta

func mathCollection() {
	gcd := func(a, b int64) int64 {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	lcm := func(a, b int64) int64 {
		return a / gcd(a, b) * b
	}

	_ = []interface{}{lcm}
}
