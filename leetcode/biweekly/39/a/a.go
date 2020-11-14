package main

// github.com/EndlessCheng/codeforces-go
func decrypt(a []int, k int) (ans []int) {
	n := len(a)
	ans = make([]int, n)
	if k != 0 {
		for i := range ans {
			if k > 0 {
				for j := i + 1; j-i <= k; j++ {
					ans[i] += a[j%n]
				}
			} else {
				for j := i - 1; i-j <= -k; j-- {
					ans[i] += a[(j+n)%n]
				}
			}
		}
	}
	return
}
