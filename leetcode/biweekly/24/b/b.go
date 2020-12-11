package main

// github.com/EndlessCheng/codeforces-go
func findMinFibonacciNumbers(k int) (ans int) {
	f := []int{1, 1}
	for f[len(f)-1] < k {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}
	for k > 0 {
		for i := len(f) - 1; ; i-- {
			if f[i] <= k {
				k -= f[i]
				ans++
				break
			}
		}
	}
	return
}
