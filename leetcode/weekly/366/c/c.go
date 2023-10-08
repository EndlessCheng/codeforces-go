package main

// https://space.bilibili.com/206214
func minOperations(s1, s2 string, x int) int {
	if s1 == s2 {
		return 0
	}
	p := []int{}
	for i, c := range s1 {
		if byte(c) != s2[i] {
			p = append(p, i)
		}
	}
	if len(p)%2 > 0 {
		return -1
	}
	f0, f1 := 0, x
	for i := 1; i < len(p); i++ {
		f0, f1 = f1, min(f1+x, f0+(p[i]-p[i-1])*2)
	}
	return f1 / 2
}

func min(a, b int) int { if b < a { return b }; return a }
