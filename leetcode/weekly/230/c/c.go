package main

// github.com/EndlessCheng/codeforces-go
func minOperations(a []int, b []int) (ans int) {
	f := func(a []int) (c [7]int, s int) {
		for _, v := range a {
			c[v]++
			s += v
		}
		return
	}
	ca, sumA := f(a)
	cb, sumB := f(b)
	if sumA > sumB {
		ca, cb, sumA, sumB = cb, ca, sumB, sumA
	}
	diff := sumB - sumA
	if diff == 0 {
		return
	}
	for i := 1; i < 6; i++ {
		c := ca[i] + cb[7-i]
		maxChange := c * (6 - i)
		if diff <= maxChange {
			ans += (diff-1)/(6-i) + 1
			return
		}
		ans += c
		diff -= maxChange
	}
	return -1
}
