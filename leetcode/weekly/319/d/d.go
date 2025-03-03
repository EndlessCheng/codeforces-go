package main

// https://space.bilibili.com/206214
func maxPalindromes(s string, k int) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		f[l+1] = max(f[l+1], f[l])
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 >= k {
				f[r+1] = max(f[r+1], f[l]+1)
				break
			}
			l--
			r++
		}
	}
	return f[n]
}
