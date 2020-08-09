package main

// github.com/EndlessCheng/codeforces-go
func findKthBit(n int, k int) (ans byte) {
	// 也可以对 k 递归
	reverse := func(a []byte) []byte {
		n := len(a)
		r := make([]byte, n)
		for i, v := range a {
			r[n-1-i] = v ^ 1
		}
		return r
	}
	s := []byte{0}
	for i := 1; i < n; i++ {
		s = append(append(s, 1), reverse(s)...)
	}
	return s[k-1] + '0'
}
