package main

// github.com/EndlessCheng/codeforces-go
func checkZeroOnes(s string) bool {
	m1, m0 := 0, 0
	for i, n := 0, len(s); i < n; {
		st := i
		v := s[st]
		for ; i < n && s[i] == v; i++ {
		}
		if v == '1' {
			m1 = max(m1, i-st)
		} else {
			m0 = max(m0, i-st)
		}
	}
	return m1 > m0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
