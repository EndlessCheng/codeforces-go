package main

// github.com/EndlessCheng/codeforces-go
func turn(s string, k int) (a int) {
	for i := 0; i < k; i++ {
		c := [123]int{}
		for j := i; j < len(s); j += k {
			b := s[j]
			for r := byte('a'); r < b; r++ {
				a += c[r]
			}
			c[b]++
		}
	}
	return
}
