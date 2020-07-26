package main

func minNumberOperations(a []int) (ans int) {
	ans += a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			ans += a[i] - a[i-1]
		}
	}
	return
}
