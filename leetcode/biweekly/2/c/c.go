package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func expand(s string) (ans []string) {
	a := [][]byte{}
	for i, n := 0, len(s); i < n; i++ {
		if s[i] != '{' {
			a = append(a, []byte{s[i]})
			continue
		}
		b := []byte{}
		for i++; ; i++ {
			b = append(b, s[i])
			i++
			if s[i] == '}' {
				break
			}
		}
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		a = append(a, b)
	}
	t := make([]byte, len(a))
	var f func(int)
	f = func(p int) {
		if p == len(a) {
			ans = append(ans, string(t))
			return
		}
		for _, b := range a[p] {
			t[p] = b
			f(p + 1)
		}
	}
	f(0)
	return
}
