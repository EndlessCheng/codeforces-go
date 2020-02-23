package main

func closestDivisors(num int) (ans []int) {
	doDivisors := func(n int, do func(d1, d2 int)) {
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				do(d, n/d)
			}
		}
		return
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := int(2e9)
	doDivisors(num+1, func(d1, d2 int) {
		if d := abs(d1 - d2); d < min {
			min = d
			ans = []int{d1, d2}
		}
	})
	doDivisors(num+2, func(d1, d2 int) {
		if d := abs(d1 - d2); d < min {
			min = d
			ans = []int{d1, d2}
		}
	})
	return
}
