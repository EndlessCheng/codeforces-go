package main

// github.com/EndlessCheng/codeforces-go
func canConvert(s, t string) (ans bool) {
	if s == t {
		return true
	}
	g := [26]int{}
	for i := range g {
		g[i] = -1
	}
	deg := [26]int{}
	for i, b := range s {
		b -= 'a'
		if c := int(t[i] - 'a'); g[b] < 0 {
			g[b] = c
			deg[c]++
		} else if g[b] != c {
			return
		}
	}
	for _, d := range deg {
		if d == 0 {
			return true
		}
	}
	return
}
