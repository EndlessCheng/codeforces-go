package main

// github.com/EndlessCheng/codeforces-go
func longestAwesome(s string) (ans int) {
	p := [1024]int{}
	for i := 1; i < 1024; i++ {
		p[i] = -1
	}
	v := 0
	for i, b := range s {
		i++
		v ^= 1 << (b & 15)
		if p[v] < 0 {
			p[v] = i
		}
		ans = max(ans, i-p[v])
		for j := 0; j < 10; j++ {
			if l := p[v^1<<j]; l >= 0 {
				ans = max(ans, i-l)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
