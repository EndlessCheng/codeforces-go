package main

import "strings"

// https://space.bilibili.com/206214
// 在文本串 text 中查找模式串 pattern
func kmpSearch(text, pattern string) bool {
	m := len(pattern)
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		b := pattern[i]
		for cnt > 0 && pattern[cnt] != b {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == b {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for _, b := range text {
		for cnt > 0 && pattern[cnt] != byte(b) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(b) {
			cnt++
		}
		if cnt == m {
			return true
		}
	}
	return false
}

func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	mx := min(n, lenA-lenB+1)
	firsts := min(lcm(n, m), lenB)
	A := strings.Repeat(a, (mx+firsts-2)/n+1)
	B := strings.Repeat(b, firsts/m)
	return kmpSearch(A, B)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
