package main

// https://space.bilibili.com/206214
func applyOperations(a []int) []int {
	n := len(a)
	b := a[:0]
	for i := 0; i < n-1; i++ {
		if a[i] > 0 {
			if a[i] == a[i+1] {
				a[i] *= 2
				a[i+1] = 0
			}
			b = append(b, a[i])
		}
	}
	if a[n-1] > 0 {
		b = append(b, a[n-1])
	}
	for i := len(b); i < n; i++ {
		a[i] = 0
	}
	return a
}
