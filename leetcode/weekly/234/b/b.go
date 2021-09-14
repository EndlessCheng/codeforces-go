package main

// github.com/EndlessCheng/codeforces-go
func reinitializePermutation(n int) (ans int) {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	b := append([]int(nil), a...)
	for {
		c := make([]int, n)
		for i := range c {
			if i%2 == 0 {
				c[i] = b[i/2]
			} else {
				c[i] = b[n/2+i/2]
			}
		}
		ans++
		if equal(c, a) {
			break
		}
		b = c
	}
	return
}

func equal(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
